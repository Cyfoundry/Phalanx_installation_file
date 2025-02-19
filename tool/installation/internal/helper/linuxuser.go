package helper

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

const (
	execAdduser string = "/usr/sbin/useradd"
	execPasswd  string = "/usr/sbin/chpasswd"
	execEcho    string = "/usr/bin/echo"
	userFile    string = "/etc/passwd"
)

func ReadEtcPasswd(f string) (list []string) {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := bufio.NewScanner(file)

	for r.Scan() {
		lines := r.Text()
		parts := strings.Split(lines, ":")
		list = append(list, parts[0])
	}
	return list
}

func CheckUserExist(username string) bool {
	userList := ReadEtcPasswd(userFile)
	for _, w := range userList {
		if username == w {
			return true
		}
	}
	return false
}

func AddUser(shell string, username string) bool {
	checkstatus := CheckUserExist(username)
	if !checkstatus {
		execCommand := []string{"-s", shell, "-m", username}
		adduserResult := UserExecute(execAdduser, execCommand...)
		content := fmt.Sprintf("Linux Add user: %s action: %s", username, adduserResult)
		InstallLog("INFO", content)
		return true
	} else {
		content := fmt.Sprintf("Linux user %s exist", username)
		InstallLog("Warning", content)
		return false
	}
}

func ChPassword(username string, password string) {
	execCommand := fmt.Sprintf("%s:%s\n", username, password)
	cmd := exec.Command("chpasswd")
	stdin, _ := cmd.StdinPipe()

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, execCommand)
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		content := fmt.Sprintf("Change Password:", err.Error())
		InstallLog("Error", content)
	} else {
		content := fmt.Sprintf("Change Password:", out)
		InstallLog("Info", content)
	}
}
