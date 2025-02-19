package helper

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
)

type Logger struct {
	CLogger *log.Logger
}

type LoggerMethon interface {
	LoggerAll(logtype string, content string)
}

func LoggingStart(path string, username string, groupname string) *Logger {
	parentdir := filepath.Dir(path)
	CheckDirC(parentdir)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		// return err
	}
	// defer f.Close()

	owner, err := user.Lookup(username)
	if err != nil {
		log.Fatalf("Error Get Owner: %v", err)
		// return err
	}

	user.Current()

	group, err := user.LookupGroup(groupname)
	if err != nil {
		log.Fatalf("Error Get Group: %v", err)
		// return err
	}

	uid, _ := strconv.Atoi(owner.Uid)
	gid, _ := strconv.Atoi(group.Gid)

	err = os.Chown(path, uid, gid)

	if err != nil {
		log.Fatalf("Error Change Owner: %v", err)
		// return err
	}
	log.SetOutput(f)
	clogger := log.New(f, "", log.Ldate|log.Ltime|log.Lshortfile)
	return &Logger{
		CLogger: clogger,
	}
}

func (l *Logger) LoggerAll(logtype string, content string) {
	if content != "" {
		l.CLogger.Printf("[%s] %s", logtype, content)
	}
}
