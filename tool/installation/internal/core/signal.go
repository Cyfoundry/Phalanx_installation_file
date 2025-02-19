package core

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Listensignal(clean func()) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				// fmt.Println("\nProgram Exit...", s)
				fmt.Println("\nStart Exit...")
				fmt.Println("Execute Clean...")
				clean()
				fmt.Println("End Exit...")
				os.Exit(0)
			case syscall.SIGUSR1:
				fmt.Println("usr1 signal")
			case syscall.SIGUSR2:
				fmt.Println("usr2 signal")
			default:
				fmt.Println("other signal")
			}
		}
	}()
}
