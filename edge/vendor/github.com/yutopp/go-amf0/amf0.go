//
// Copyright (c) 2018- yutopp (yutopp@gmail.com)
//
// Distributed under the Boost Software License, Version 1.0. (See accompanying
// file LICENSE_1_0.txt or copy at  https://www.boost.org/LICENSE_1_0.txt)
//

package amf0

// Marker Represents AMF0 object types
type Marker byte

const (
	// MarkerNumber A marker for Number types
	MarkerNumber Marker = 0x00
	// MarkerBoolean A marker for Boolean types
	MarkerBoolean Marker = 0x01
	// MarkerString A marker for String types
	MarkerString Marker = 0x02
	// MarkerObject A marker for Object types
	MarkerObject Marker = 0x03
	// MarkerMovieclip A marker for ovieclip types
	MarkerMovieclip Marker = 0x04 // reserved, not supported
	// MarkerNull A marker for Null types
	MarkerNull Marker = 0x05
	// MarkerUndefined A marker for Undefined types
	MarkerUndefined Marker = 0x06
	// MarkerReference A marker for Reference types
	MarkerReference Marker = 0x07
	// MarkerEcmaArray A marker for EcmaArray types
	MarkerEcmaArray Marker = 0x08
	// MarkerObjectEnd A marker for ObjectEnd types
	MarkerObjectEnd Marker = 0x09
	// MarkerStrictArray A marker for StrictArray types
	MarkerStrictArray Marker = 0x0A
	// MarkerDate A marker for Date types
	MarkerDate Marker = 0x0B
	// MarkerLongString A marker for LongString types
	MarkerLongString Marker = 0x0C
	// MarkerUnsupported A marker for Unsupported types
	MarkerUnsupported Marker = 0x0D
	// MarkerRecordSet A marker for RecordSet types
	MarkerRecordSet Marker = 0x0E // reserved, not supported
	// MarkerXMLDocument A marker for XMLDocument types
	MarkerXMLDocument Marker = 0x0F
	// MarkerTypedObject A marker for TypedObject types
	MarkerTypedObject Marker = 0x10
)

// ECMAArray EcmaArray representation in Golang
type ECMAArray map[string]interface{}

// ObjectEnd ObjectEnd representation in Golang
var ObjectEnd = struct{}{}
