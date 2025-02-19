package install

import (
	"installation/internal/core"
	"log"
	"os/exec"
	"strings"
)

const (
	execSystemctl   string = "/usr/bin/systemctl"
	execAptGet      string = "/usr/bin/apt-get"
	execChattr      string = "/usr/bin/chattr"
	execApt         string = "/usr/bin/apt"
	execDpkg        string = "/usr/bin/dpkg"
	execMode        string = "/usr/sbin/modprobe"
	execSN          string = "/usr/sbin/sn_generator"
	execUPgrub      string = "/usr/sbin/update-grub"
	execAdduser     string = "/usr/sbin/adduser"
	execUsermod     string = "/usr/sbin/usermod"
	exeHostname     string = "/usr/bin/hostnamectl"
	execGrubconfig  string = "/usr/sbin/grub-mkconfig"
	exeSysctl       string = "/usr/sbin/sysctl"
	exePython3      string = "/usr/bin/python3"
	autointeractive string = "DEBIAN_FRONTEND=noninteractive"
	linuxsync       string = "/usr/bin/sync"
	execAptKey      string = "/usr/bin/apt-key"
	execEcho        string = "/usr/bin/echo"
	execMake        string = "/usr/bin/make"
	execGpg         string = "/usr/bin/gpg"
	execbash        string = "/bin/bash"
	execFind        string = "/usr/bin/find"
	execZeek        string = "/opt/zeek/bin/zeekctl"
	execSetcap      string = "/usr/sbin/setcap"
	execPG          string = "/usr/bin/psql"
	execSUDO        string = "/usr/bin/sudo"
	execChmod       string = "/usr/bin/chmod"
	execNpm         string = "/usr/bin/npm"
	execPm2         string = "/usr/local/bin/pm2"
	execKillAll     string = "/usr/bin/killall"
	execN           string = "/usr/local/bin/n"
	execNpmUpdated  string = "/usr/local/bin/npm"
	execPip         string = "/usr/bin/pip"
	execPip3        string = "/usr/bin/pip3"
)

var bar core.Bar

func InstallExecute(commandPath string, execCommand ...string) (output string) {
	log.Printf("ExecutePath: %s ## ExecuteCommand: %s", commandPath, execCommand)
	var cmd *exec.Cmd
	if execCommand[0] != "" {
		cmd = exec.Command(commandPath, execCommand...)
	} else {
		cmd = exec.Command(commandPath)
	}

	// err := cmd.Run()

	out, err := cmd.CombinedOutput()

	// log.Fatal(string(out))
	log.Println(string(out))
	// fmt.Println(err)
	result := string(out)
	result = strings.TrimSpace(result)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(result)
	return result
}

func InstallExecuteNoExit(commandPath string, execCommand ...string) (output string) {
	log.Printf("ExecutePath: %s ## ExecuteCommand: %s", commandPath, execCommand)
	var cmd *exec.Cmd
	if execCommand[0] != "" {
		cmd = exec.Command(commandPath, execCommand...)
	} else {
		cmd = exec.Command(commandPath)
	}

	// err := cmd.Run()

	out, err := cmd.CombinedOutput()

	// log.Fatal(string(out))
	log.Println(string(out))
	// fmt.Println(err)
	result := string(out)
	result = strings.TrimSpace(result)
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	log.Println(result)
	return result
}

func InstallExecuteWithPath(commandPath string, path string, execCommand ...string) (output string) {
	log.Printf("ExecutePath: %s ## ExecuteCommand: %s", commandPath, execCommand)
	var cmd *exec.Cmd
	if execCommand[0] != "" {
		cmd = exec.Command(commandPath, execCommand...)
	} else {
		cmd = exec.Command(commandPath)
	}

	cmd.Dir = path

	out, err := cmd.CombinedOutput()

	log.Println(string(out))
	result := string(out)
	result = strings.TrimSpace(result)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(result)
	return result
}

func InstallExecuteWithPathNoExit(commandPath string, path string, execCommand ...string) (output string) {
	log.Printf("ExecutePath: %s ## ExecuteCommand: %s", commandPath, execCommand)
	var cmd *exec.Cmd
	if execCommand[0] != "" {
		cmd = exec.Command(commandPath, execCommand...)
	} else {
		cmd = exec.Command(commandPath)
	}

	cmd.Dir = path

	out, err := cmd.CombinedOutput()

	log.Println(string(out))
	result := string(out)
	result = strings.TrimSpace(result)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(result)
	return result
}
