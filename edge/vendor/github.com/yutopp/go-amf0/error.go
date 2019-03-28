//
// Copyright (c) 2018- yutopp (yutopp@gmail.com)
//
// Distributed under the Boost Software License, Version 1.0. (See accompanying
// file LICENSE_1_0.txt or copy at  https://www.boost.org/LICENSE_1_0.txt)
//

package amf0

import (
	"fmt"
	"reflect"
)

// UnexpectedMarkerError Occurs when an unexpected marker is passed to the decoder
type UnexpectedMarkerError struct {
	Marker uint8
}

// Error Returns a string representation of the error
func (e *UnexpectedMarkerError) Error() string {
	return fmt.Sprintf("Unexpected marker: Marker = %+v", e.Marker)
}

// UnsupportedError Occurs when decode unsupported messages
type UnsupportedError struct {
}

// Error Returns a string representation of the error
func (e *UnsupportedError) Error() string {
	return "Unsupported"
}

// UnexpectedValueError Occurs when an unexpected value is passed to the encoder
type UnexpectedValueError struct {
	Kind reflect.Kind
}

// Error Returns a string representation of the error
func (e *UnexpectedValueError) Error() string {
	return fmt.Sprintf("Unexpected value: Kind = %+v", e.Kind)
}

// UnexpectedKeyTypeError Occurs when an unexpected key type is passed to the encoder/decoder
type UnexpectedKeyTypeError struct {
	ActualKind reflect.Kind
	ExpectKind reflect.Kind
}

// Error Returns a string representation of the error
func (e *UnexpectedKeyTypeError) Error() string {
	return fmt.Sprintf("Unsupported key kind: %+v should be %+v", e.ActualKind.String(), e.ExpectKind.String())
}

// DecodeError Occurs when general errors are happen in the decoder
type DecodeError struct {
	Message string
	Dump    string
}

// Error Returns a string representation of the error
func (e *DecodeError) Error() string {
	return fmt.Sprintf("Message = %s, Dump = \n%s", e.Message, e.Dump)
}

// NotAssignableError Occurs when failed to assign a decoded value to the receiver value
type NotAssignableError struct {
	Message string
	Kind    reflect.Kind
	Type    reflect.Type
}

// Error Returns a string representation of the error
func (e *NotAssignableError) Error() string {
	return fmt.Sprintf("Not assignable to receiver value: Message=%+v, Kind=%s, Type=%s",
		e.Message,
		e.Kind.String(),
		e.Type.String(),
	)
}

// ErrObjectEndMarker ...
var ErrObjectEndMarker = fmt.Errorf("ObjectEndMarker")
