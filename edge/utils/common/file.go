package common

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ReadString -读取文件内容
func ReadString(filename string) string {
	if contents, err := ioutil.ReadFile(filename); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		// result := strings.Replace(string(contents), "\n", "", 1)
		return string(contents[:])
	}

	return ""
}

// GetCurrentDirectory - 获取程序运行路径
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err.Error())
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// AppendToFile - fileName:文件名字(带全路径)
// content: 写入的内容
func AppendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		return err
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	return err
}
