package common

import (
	"strconv"
	"unsafe"
)

// // IntToBytes -
// func IntToBytes(x int) []byte {
// 	bbuf := bytes.NewBuffer([]byte{})
// 	binary.Write(bbuf, binary.BigEndian, x)
// 	return bbuf.Bytes()
// }

// // BytesToInt -
// func BytesToInt(b []byte) int {
// 	bbuf := bytes.NewBuffer(b)
// 	var x int
// 	binary.Read(bbuf, binary.BigEndian, &x)
// 	return x
// }

// // IntToBytes -
// func IntToBytes(x int) []byte {
// 	bs := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(bs, 31415926)
// 	return bs
// }

// // BytesToInt -
// func BytesToInt(b []byte) int {
// 	bs := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(bs, 31415926)
// }

// StringToInt -
func StringToInt(str string) (int, error) {
	i, err := strconv.Atoi(str)
	return i, err
}

// IntToString -
func IntToString(i int) string {
	str := strconv.Itoa(i)
	return str
}

// Int2Bytes -
func Int2Bytes(data int) (ret []byte) {
	len := unsafe.Sizeof(data)
	ret = make([]byte, len)
	tmp := 0xff
	var index uint = 0
	for index = 0; index < uint(len); index++ {
		ret[index] = byte((tmp << (index * 8) & data) >> (index * 8))
	}
	return ret
}

// Bytes2Int -
func Bytes2Int(data []byte) int {
	ret := 0
	len := len(data)
	var i uint = 0
	for i = 0; i < uint(len); i++ {
		ret = ret | (int(data[i]) << (i * 8))
	}
	return ret
}
