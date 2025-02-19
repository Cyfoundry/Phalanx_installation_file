package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"installation/internal/adapter"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

func ChangeOWN(path string, username string, groupname string) error {

	owner, err := user.Lookup(username)
	if err != nil {
		log.Fatalf("Error Get Owner: %v", err)
		return err
	}

	group, err := user.LookupGroup(groupname)
	if err != nil {
		log.Fatalf("Error Get Group: %v", err)
		return err
	}

	uid, _ := strconv.Atoi(owner.Uid)
	gid, _ := strconv.Atoi(group.Gid)

	err = os.Chown(path, uid, gid)

	if err != nil {
		log.Fatalf("Error Change Owner: %v", err)
		return err
	}

	return nil
}

func ChownR(path string, username string, groupname string) error {
	owner, err := user.Lookup(username)
	if err != nil {
		log.Fatalf("Error Get Owner: %v", err)
		return err
	}

	group, err := user.LookupGroup(groupname)
	if err != nil {
		log.Fatalf("Error Get Group: %v", err)
		return err
	}

	uid, _ := strconv.Atoi(owner.Uid)
	gid, _ := strconv.Atoi(group.Gid)
	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chown(name, uid, gid)
		}
		return err
	})
}

func ChmodR(path string, mode fs.FileMode) error {
	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chmod(name, mode)
			log.Printf("Change: %s, mode: %s, Error: %v", name, mode, err)
		}
		return err
	})
}

func ChmodRwithFile(path string, filetype string, mode fs.FileMode) error {
	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {

		if err == nil {
			if !info.IsDir() && filepath.Ext(path) == filetype {
				err = os.Chmod(name, mode)
				log.Printf("Change: %s, mode: %s, Error: %v", name, mode, err)
			}
		}

		return err
	})
}

func RemoveGlob(path string) (err error) {
	contents, err := filepath.Glob(path)
	if err != nil {
		return
	}
	for _, item := range contents {
		err = os.RemoveAll(item)
		if err != nil {
			return
		}
	}
	return
}

type Format map[string]interface{}

func FormatStr(content string, format Format) string {
	temp := &bytes.Buffer{}
	template.Must(template.New("").Parse(content)).Execute(temp, format)

	return temp.String()
}

func WriteFile(path string, content string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	} else {
		file.WriteString(content)
		log.Printf("Write %s Done", path)
	}
	defer file.Close()
}

func ReadFile(path string) string {

	dat, err := os.ReadFile(path)
	log.Println("Read File Failed", err)

	return string(dat)
}

func CopyFileAll(src string, dst string) error {
	contents, err := filepath.Glob(src)
	if err != nil {
		return err
	}

	filelength := len(contents)
	for i := 0; i < (filelength); i++ {
		item := contents[i]
		path := item
		filename := filepath.Base(path)
		dstformat := fmt.Sprintf("%s/%s", dst, filename)
		bytesRead, err := ioutil.ReadFile(item)
		if err != nil {
			log.Fatal(err)
			fmt.Println(err.Error())
			return err
		}

		err = ioutil.WriteFile(dstformat, bytesRead, 0755)
		if err != nil {
			log.Fatal(err)
			fmt.Println(err.Error())
			return err
		}
	}

	return nil
}
func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		log.Printf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		log.Printf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		log.Printf("Writing to output file failed: %s", err)
	}
	err = os.Remove(sourcePath)
	if err != nil {
		log.Printf("Failed removing original file: %s", err)
	}
	return nil
}
func InstallLog(logtype string, content string) {
	if content != "" {
		log.Printf("[%s] %s", logtype, content)
	}
}

func UserExecute(commandPath string, execCommand ...string) (output string) {
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
	// fmt.Println(string(out))
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

func RenameAdpater() []adapter.AdapterInfo {
	var (
		wanNum int = 0
		lanNum int = 0
	)
	ads, _ := adapter.Adapters()
	adapters := ads.Adapters
	var temp []adapter.AdapterInfo
	for _, ad := range adapters {
		if ad.HardwareAddr != "" && ad.Name != "br0" {
			if (wanNum) != 1 {
				ad.Name = fmt.Sprintf("wan%d", wanNum)
				wanNum = wanNum + 1
			} else {
				ad.Name = fmt.Sprintf("lan%d", lanNum)
				lanNum = lanNum + 1
			}
			temp = append(temp, ad)
		}
	}

	return temp
}
