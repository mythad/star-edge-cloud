package common

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

// func ExecCommand(path string, exeFileName string) string {
// 	dir :=.GetCurrentDirectory()
// 	defer os.Chdir(dir)

// 	os.Chdir(fmt.Sprintf("./plugins/device/%[1]s/", dev.ID))
// 	cmd := exec.Command("/bin/bash", "-c", str)
// 	// cmd.SysProcAttr = &syscall.SysProcAttr{Foreground: false}
// 	if err := cmd.Start(); err != nil {
// 		return err
// 	}
// }

// ExecDeamonCommand -
// func ExecDeamonCommand(path string, commandName string, params ...string) string {
// 	dir := GetCurrentDirectory()
// 	defer os.Chdir(dir)

// 	os.Chdir(path)
// 	os.Chmod(commandName, os.ModePerm)
// 	cmd := exec.Command("./"+commandName, params...)
// 	//显示运行的命令
// 	// log.Println(cmd.Args)
// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		// log.Println(err)
// 		return err.Error()
// 	}
// 	if err := cmd.Start(); err != nil {
// 		// log.Println(err)
// 		return err.Error()
// 	}
// 	reader := bufio.NewReader(stdout)
// 	// return the first line
// 	// for {
// 	line, err := reader.ReadString('\n')
// 	if io.EOF == err {
// 		// log.Println(err)
// 		return ""
// 	}

// 	return line
// 	// }
// 	// cmd.Wait()
// }

// ExecDeamonCommand -
func ExecDeamonCommand(workingdir string, path string, commandName string, params ...string) string {
	os.Chdir(path)
	os.Chmod(commandName, os.ModePerm)
	cmd := exec.Command("./"+commandName, params...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		// os.Chdir(workingdir)
		// log.Println(err.Error())
	}
	str := out.String()
	os.Chdir(workingdir)
	return str
}

// ExecCheckStatus -
func ExecCheckStatus(workingdir string, path string, commandName string, params ...string) string {
	os.Chdir(path)
	os.Chmod(commandName, os.ModePerm)
	cmd := exec.Command("./"+commandName, params...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		// os.Chdir(workingdir)
		// log.Println(err.Error())
	}
	str := out.String()
	os.Chdir(workingdir)
	return str
}

// StatusCovert -
func StatusCovert(str string) int {
	if strings.Contains(str, "running") {
		return 2
	}

	// Service is stopped
	if strings.Contains(str, "stopped") {
		return 1
	}

	// Service is not installed
	if strings.Contains(str, "installed") {
		return 0
	}

	return -1
}
