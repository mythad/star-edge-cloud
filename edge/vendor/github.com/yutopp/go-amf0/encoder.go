//
// Copyright (c) 2018- yutopp (yutopp@gmail.com)
//
// Distributed under the Boost Software License, Version 1.0. (See accompanying
// file LICENSE_1_0.txt or copy at  https://www.boost.org/LICENSE_1_0.txt)
//

package amf0

import (
	"encoding/binary"
	"github.com/pkg/errors"
	"io"
	"math"
	"reflect"
	"sort"
	"time"
)

// Encoder Encode objects in Golang into AMF0 and writes to the writer
type Encoder struct {
	w        io.Writer
	sortKeys bool
}

// NewEncoder Create a new instance of Encoder
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: w,
	}
}

// Encode Encode objects
func (enc *Encoder) Encode(v interface{}) error {
	rv := reflect.ValueOf(v)
	return enc.encode(rv)
}

// Reset Reset a state of the encoder
func (enc *Encoder) Reset(w io.Writer) {
	enc.w = w
}

func (enc *Encoder) encode(rv reflect.Value) error {
	switch rv.Kind() {
	case reflect.Ptr:
		return enc.encode(rv.Elem())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fallthrough
	case reflect.Float32, reflect.Float64:
		return enc.encodeNumber(rv)

	case reflect.Bool:
		return enc.encodeBoolean(rv)

	case reflect.String:
		return enc.encodeString(rv)

	case reflect.Map:
		return enc.encodeMap(rv)

	case reflect.Array, reflect.Slice:
		return enc.encodeStrictArray(rv)

	case reflect.Interface:
		if rv.IsNil() {
			return enc.encodeNull()
		}
		return enc.Encode(rv.Interface())

	case reflect.Invalid:
		return enc.encodeNull()

	case reflect.Struct:
		switch rv.Type() {
		case reflect.TypeOf(ObjectEnd):
			return enc.encodeObjectEnd()
		case reflect.TypeOf(time.Time{}):
			return enc.encodeDate(rv)
		default:
			return enc.encodeObject(rv)
		}

	default:
		return &UnexpectedValueError{
			Kind: rv.Kind(),
		}
	}
}

func (enc *Encoder) encodeMap(rv reflect.Value) error {
	if rv.Type() == reflect.TypeOf(ECMAArray{}) {
		return enc.encodeMapAsECMAArray(rv)
	}

	return enc.encodeMapAsObject(rv)
}

func (enc *Encoder) encodeObject(rv reflect.Value) error {
	if err := enc.writeU8(uint8(MarkerObject)); err != nil {
		return err
	}

	ty := rv.Type()
	numFields := rv.NumField()
	for i := 0; i < numFields; i++ {
		fieldType := ty.Field(i)

		key := fieldType.Tag.Get("amf0")
		if key == "" {
			key = fieldType.Name
		}

		if err := enc.writeUTF8(key); err != nil {
			return err
		}

		value := rv.Field(i)
		if err := enc.encode(value); err != nil {
			return err
		}
	}

	return enc.encodeObjectEnd()
}

func (enc *Encoder) encodeNumber(rv reflect.Value) error {
	if err := enc.writeU8(uint8(MarkerNumber)); err != nil {
		return err
	}

	var d float64
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		d = float64(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		d = float64(rv.Uint())
	case reflect.Float32, reflect.Float64:
		d = rv.Float()
	default:
		return &UnexpectedValueError{
			Kind: rv.Kind(),
		}
	}

	return enc.writeDouble(d)
}

func (enc *Encoder) encodeBoolean(rv reflect.Value) error {
	if err := enc.writeU8(uint8(MarkerBoolean)); err != nil {
		return err
	}

	b := uint8(0)
	if rv.Bool() {
		b = 1
	}

	return enc.writeU8(b)
}

func (enc *Encoder) encodeString(rv reflect.Value) error {
	s := rv.String()
	if len(s) > 65535 {
		return enc.encodeLongString(rv)
	}

	if err := enc.writeU8(uint8(MarkerString)); err != nil {
		return err
	}
	return enc.writeUTF8(s)
}

func (enc *Encoder) encodeMapAsObject(rv reflect.Value) error {
	if err := enc.writeU8(uint8(MarkerObject)); err != nil {
		return err
	}

	keys := rv.MapKeys()
	if enc.sortKeys {
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].String() < keys[j].String()
		})
	}

	for _, key := range keys {
		if key.Kind() != reflect.String {
			return &UnexpectedKeyTypeError{
				ActualKind: key.Kind(),
				ExpectKind: reflect.String,
			}
		}

		if err := enc.writeUTF8(key.String()); err != nil {
			return err
		}

		value := rv.MapIndex(key)
		if err := enc.encode(value); err != nil {
			return err
		}
	}

	return enc.encodeObjectEnd()
}

func (enc *Encoder) encodeMovieClip(rv reflect.Value) error {
	return errors.New("Not implemented: MovieClip")
}

func (enc *Encoder) encodeNull() error {
	return enc.writeU8(uint8(MarkerNull))
}

func (enc *Encoder) encodeUndefined(rv reflect.Value) error {
	return errors.New("Not implemented: Undefined")
}

func (enc *Encoder) encodeReference(rv reflect.Value) error {
	return errors.New("Not implemented: Reference")
}

func (enc *Encoder) encodeMapAsECMAArray(rv reflect.Value) error {
	if err := enc.writeU8(uint8(MarkerEcmaArray)); err != nil {
		return err
	}

	l := rv.Len()
	if err := enc.writeU32(uint32(l)); err != nil {
		return err
	}

	keys := rv.MapKeys()
	if enc.sortKeys {
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].String() < keys[j].String()
		})
	}

	for _, key := range keys {
		if err := enc.writeUTF8(key.String()); err != nil {
			return err
		}

		value := rv.MapIndex(key)
		if err := enc.encode(value); err != nil {
			return err
		}
	}

	return enc.encodeObjectEnd()
}

func (enc *Encoder) encodeObjectEnd() error {
	if err := enc.writeUTF8(""); err != nil { // utf-8-empty
		return err
	}

	return enc.writeU8(uint8(MarkerObjectEnd))
}

func (enc *Encoder) encodeStrictArray(rv reflect.Value) error {
	if err := enc.writeU8(uint8(MarkerStrictArray)); err != nil {
		return err
	}

	if err := enc.writeU32(uint32(rv.Len())); err != nil {
		return err
	}

	for i := 0; i < rv.Len(); i++ {
		if err := enc.encode(rv.Index(i)); err != nil {
			return err
		}
	}

	return nil
}

func (enc *Encoder) encodeDate(rv reflect.Value) error {
	t := rv.Interface().(time.Time)
	t = t.In(time.UTC) // Time zone is not supported yet, thus force convert to UTC. TODO: fix

	if t.UnixNano()%int64(time.Millisecond) != 0 {
		return errors.Errorf("Date time of nano sec is not supported: Expected = 0, Actual = %d",
			t.UnixNano()%int64(time.Millisecond),
		)
	}

	unixMs := float64(t.UnixNano() / int64(time.Millisecond))
	tz := int16(0x00)

	if err := enc.writeU8(uint8(MarkerDate)); err != nil {
		return err
	}

	if err := enc.writeDouble(unixMs); err != nil {
		return err
	}

	return enc.writeS16(tz)
}

func (enc *Encoder) encodeLongString(rv reflect.Value) error {
	return errors.New("Not implemented: LongString")
}

func (enc *Encoder) encodeUnsupported(rv reflect.Value) error {
	return errors.New("Not implemented: Unsupported")
}

func (enc *Encoder) encodeRecordSet(rv reflect.Value) error {
	return errors.New("Not implemented: RecordSet")
}

func (enc *Encoder) encodeXMLDocument(rv reflect.Value) error {
	return errors.New("Not implemented: XMLDocument")
}

func (enc *Encoder) encodeTypedObject(rv reflect.Value) error {
	return errors.New("Not implemented: TypedObject")
}

func (enc *Encoder) writeU8(num uint8) error {
	_, err := enc.w.Write([]byte{num}) // TODO: optimize
	return err
}

func (enc *Encoder) writeU16(num uint16) error {
	buf := make([]byte, 2) // TODO: optimize
	binary.BigEndian.PutUint16(buf, num)

	_, err := enc.w.Write(buf)
	return err
}

func (enc *Encoder) writeS16(num int16) error {
	return enc.writeU16(uint16(num))
}

func (enc *Encoder) writeU32(num uint32) error {
	buf := make([]byte, 4) // TODO: optimize
	binary.BigEndian.PutUint32(buf, num)

	_, err := enc.w.Write(buf)
	return err
}

func (enc *Encoder) writeDouble(f64 float64) error {
	buf := make([]byte, 8) // TODO: optimize
	u64 := math.Float64bits(f64)
	binary.BigEndian.PutUint64(buf, u64)

	_, err := enc.w.Write(buf)
	return err
}

func (enc *Encoder) writeUTF8(str string) error {
	l := uint16(len(str))
	if err := enc.writeU16(l); err != nil {
		return err
	}
	_, err := enc.w.Write([]byte(str))
	return err
}
