package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// SN碼組合：產品名稱-版號-年月日時分-uuid
func main() {
	var match bool
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	uuid := uuid.New()
	key := uuid.String()
	formatdate := strings.Replace(date, "-", "", -1)
	formatkey := strings.Replace(key, "-", "", -1)
	cmd := exec.Command("id", "-u")
	output, _ := cmd.Output()
	i, _ := strconv.Atoi(string(output[:len(output)-1]))
	if i == 0 {
		checkfile := "/etc/sn-id"
		if _, err := os.Stat(checkfile); err == nil {
			data, err := os.ReadFile(checkfile)
			check(err)
			match, _ = regexp.MatchString("Phalanx-v[0-9].[0-9]{1,}-[0-9]{8}-[a-z A-Z 0-9]{32}", string(data))
			sn := []byte(fmt.Sprintf("Phalanx-v0.1-%s-%s\n", formatdate, formatkey))
			writefont := []byte(string(sn))
			if !match {
				fmt.Printf("Phalanx-v0.1-%s-%s\n", formatdate, formatkey)
				err := os.WriteFile(checkfile, writefont, 0644)
				check(err)
			}

		} else {
			sn := fmt.Sprintf("Phalanx-v0.1-%s-%s\n", formatdate, formatkey)
			writefont := []byte(string(sn))
			fmt.Printf("Phalanx-v0.1-%s-%s\n", formatdate, formatkey)
			err := os.WriteFile(checkfile, writefont, 0644)
			check(err)
		}

	} else {
		fmt.Printf("This program must be run as root! (sudo)")
	}

}
