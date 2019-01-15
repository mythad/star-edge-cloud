package common

import (
	"strconv"
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
