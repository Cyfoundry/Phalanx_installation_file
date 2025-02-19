package main

import (
	"fmt"
	"installation/internal/core"
	"installation/internal/helper"
	"installation/internal/install"
	"os"
	"time"
)

func main() {
	permission := helper.CheckPermission()
	if permission == 0 {
		helper.LoggingStart("/var/log/phalanx/install.log", "phalanx", "phalanx")
		menu := core.NewMenu("Phalanx Installation Menu", "=>")
		menu.AddItem("Installation", "setup")
		menu.AddItem("Install Jetstream", "setup_jetstream")
		// menu.AddItem("Install Endpoint", "setup_endpoint")
		// menu.AddItem("Install Jobmanager", "setup_jobmanager")
		// menu.AddItem("Install Crawler", "setup_crawler")
		menu.AddItem("Install Web", "setup_web")
		// menu.AddItem("Install Tunnel", "setup_tunnel")
		menu.AddItem("Install Subber", "setup_subber")
		menu.AddItem("Reboot", "reboot")
		menu.AddItem("Exit", "exit")
		for {
			menu.Flush()
			choice := menu.Display()
			menu.Flush()
			fmt.Printf("Choice: %s\n", choice)

			if choice == "setup" {
				install.Setup()
				install.SetupNetwork()
				install.SetupHost()
				install.SetupResolve()
				install.SetupLimitUser()
				install.SetupModules()
				install.SetupSysctl()
				install.SetupBanner()
				install.SetupDhcp()
				install.SetupNftable()
				core.Listensignal(func() {})
				time.Sleep(2000 * time.Millisecond)
				menu.Flush()
			}

			if choice == "setup_jetstream" {
				install.JetstreamSetup()
				install.EndpointSetup()
				install.JobmanagerSetup()
				install.CrawlerSetup()
				install.TunnelSetup()
				core.Listensignal(func() {})
				time.Sleep(2000 * time.Millisecond)
				menu.Flush()
			}

			// if choice == "setup_endpoint" {
			// 	install.EndpointSetup()
			// 	core.Listensignal(func() {})
			// 	time.Sleep(2000 * time.Millisecond)
			// 	menu.Flush()
			// }

			// if choice == "setup_jobmanager" {
			// 	install.JobmanagerSetup()
			// 	core.Listensignal(func() {})
			// 	time.Sleep(2000 * time.Millisecond)
			// 	menu.Flush()
			// }

			// if choice == "setup_crawler" {
			// 	install.CrawlerSetup()
			// 	core.Listensignal(func() {})
			// 	time.Sleep(2000 * time.Millisecond)
			// 	menu.Flush()
			// }

			if choice == "setup_web" {
				install.WebSetup()
				install.DBSetup()
				core.Listensignal(func() {})
				time.Sleep(2000 * time.Millisecond)
				menu.Flush()
			}

			// if choice == "setup_tunnel" {
			// 	install.TunnelSetup()
			// 	core.Listensignal(func() {})
			// 	time.Sleep(2000 * time.Millisecond)
			// 	menu.Flush()
			// }

			if choice == "setup_subber" {
				install.SubberSetup()
				core.Listensignal(func() {})
				time.Sleep(2000 * time.Millisecond)
				menu.Flush()
			}

			if choice == "reboot" {
				install.Reboot()
			}

			if choice == "exit" {
				menu.Flush()
				os.Exit(0)
			}
		}
	} else {
		fmt.Printf("This program must be run as root!\n")
	}
}
