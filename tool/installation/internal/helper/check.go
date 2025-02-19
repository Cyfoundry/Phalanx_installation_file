package helper

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"

	"github.com/3th1nk/cidr"
	"github.com/jessevdk/go-flags"
)

func CheckPermission() int {
	cmd := exec.Command("id", "-u")
	output, _ := cmd.Output()
	i, _ := strconv.Atoi(string(output[:len(output)-1]))

	return i
}

func CheckPErr(err error) {
	if err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}

func CheckIPAddress(ip string) error {
	if net.ParseIP(ip) == nil {
		log.Printf("IP Address: %s - Invalid\n", ip)
		parseErr := fmt.Sprintf("IP Address: %s - Invalid\n", ip)
		err := errors.New(parseErr)
		return err
	} else {
		// fmt.Printf("IP Address: %s - Valid\n", ip)
		return nil
	}
}

func CheckCIDR(cidrStr string) error {
	_, err := cidr.Parse(cidrStr)
	if err != nil {
		log.Fatal(err)
		return err
	} else {
		return nil
	}
}

func CheckExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("File: %s not exisit", path)
		return err
	}

	return nil
}

func CheckDirC(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)

		if err != nil {
			log.Fatalf("Create directory %s Failed", path)
		}
	}
}
