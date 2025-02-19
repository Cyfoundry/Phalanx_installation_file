package install

import (
	"fmt"
	"installation/internal/helper"
	"os/exec"
)

func WebSetup() {
	var tmpNpm string

	/* 安裝mongodb
	   sudo openssl rand -base64 128 > keyFile
	   sudo chmod 600 keyFile
	   sudo chown 999:999 keyFile

	   ##　psql安裝、pm2安裝
	   sudo apt-get update
	   sudo apt-get install postgresql-client
	   sudo apt install npm
	   sudo npm install --global pm2

	   ##　開啟postgres跟mongodb
	   sudo docker-compose up -d

	   ##　mongoDB 初始化
	   /usr/sbin/mongoInit
	*/
	bar.NewOption(0, 100, "Initilize")
	// exec

	/*
	   cd ~/05-web/standalone/
	   PORT=80 sudo pm2 start "npm run start" --name server
	   cd ~/05-web/socket_server
	   sudo pm2 start "npm run start" --name socket
	   cd ~/05-web/phalanx_schedule_ts
	   sudo pm2 start "npm run start" --name schedule

	   sudo pm2 save
	   sudo pm2 startup
	*/
	/*
		sudo npm install n -g
		sudo n lts
		sudo npm install next -g
		cd ~/xlsx_mongodb/
		pip install -r requirements.txt
		python3 main.py
	*/
	tmpNpm = execNpmUpdated
	if err := helper.CheckExist(tmpNpm); err != nil {
		tmpNpm = execNpm
	}

	execCommand := []string{"install", "n", "-g"}
	installResult := InstallExecute(tmpNpm, execCommand...)
	content := fmt.Sprintf("tool install: %s", installResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(10))

	execCommand = []string{"-c", "n lts"}
	toolResult := InstallExecute(execbash, execCommand...)
	content = fmt.Sprintf("tool install: %s", toolResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(15))

	execCommand = []string{"install", "semver", "-g"}
	installResult = InstallExecuteNoExit(tmpNpm, execCommand...)
	content = fmt.Sprintf("tool install: %s", installResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(17))

	execCommand = []string{"install", "next", "-g"}
	installResult = InstallExecuteNoExit(tmpNpm, execCommand...)
	content = fmt.Sprintf("tool install: %s", installResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(20))

	execCommand = []string{"install", "next", "-g"}
	installResult = InstallExecuteNoExit(tmpNpm, execCommand...)
	content = fmt.Sprintf("tool install: %s", installResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(21))

	/*
	   //install CommunityID
	   	execCommand = []string{"-u", "router", "./configure", "--with-bifcl=/opt/zeek/bin/bifcl"}
	   	confResult := InstallExecuteWithPath(execSUDO, "/config/zeek-community-id", execCommand...)
	   	content = fmt.Sprintf("install CommunityID: %s", confResult)
	   	helper.InstallLog("INFO", content)
	   	bar.Add(int64(50))
	*/

	execCommand = []string{"-c", "PORT=80 pm2 start \"/usr/local/bin/npm run start\" --name server"}
	webinstallResult := InstallExecuteWithPath(execbash, "/opt/web/standalone", execCommand...)
	content = fmt.Sprintf("tool install: %s", webinstallResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(30))

	// execCommand = []string{"PORT=80", "pm2", "start", "\"/usr/local/bin/npm run start\"", "--name", "server"}
	// webinstallResult := InstallExecuteWithPath(execSUDO, "/opt/web/standalone", execCommand...)
	// content = fmt.Sprintf("start: %s", webinstallResult)
	// helper.InstallLog("INFO", content)
	// bar.Add(int64(30))

	execCommand = []string{"install", "-r", "/config/xlsx_mongodb/requirements.txt"}
	PipinstallResult := InstallExecute(execPip, execCommand...)
	content = fmt.Sprintf("Pip  install %s", PipinstallResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(40))

	execCommand = []string{"/config/xlsx_mongodb/main.py"}
	cmd := exec.Command(exePython3, execCommand...)
	content = fmt.Sprintf("Pip  install %s", cmd)
	helper.InstallLog("INFO", content)
	bar.Add(int64(50))

	execCommand = []string{"-c", "pm2 start \"/usr/local/bin/npm run start\" --name socket"}
	socketinstallResult := InstallExecuteWithPath(execbash, "/opt/web/phalanx_socket_ts", execCommand...)
	content = fmt.Sprintf("tool install: %s", socketinstallResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(60))

	// execCommand = []string{"start", "\"/usr/local/bin/npm run start\"", "--name", "socket"}
	// socketinstallResult := InstallExecuteWithPath(execPm2, "/opt/web/socket_server", execCommand...)
	// content = fmt.Sprintf("Pm2 start: %s", socketinstallResult)
	// helper.InstallLog("INFO", content)
	// bar.Add(int64(60))
	// sudo pm2 start "NODE_ENV=production node dist/app.js" --name schedule
	execCommand = []string{"-c", "pm2 start \"NODE_ENV=production node dist/app.js\" --name schedule"}
	scheduleinstallResult := InstallExecuteWithPath(execbash, "/opt/web/phalanx_schedule_ts", execCommand...)
	content = fmt.Sprintf("tool install: %s", scheduleinstallResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(70))

	// execCommand = []string{"start", "\"/usr/local/bin/npm run start\"", "--name", "schedule"}
	// scheduleinstallResult := InstallExecuteWithPath(execPm2, "/opt/web/phalanx_schedule_ts", execCommand...)
	// content = fmt.Sprintf("Pm2 start: %s", scheduleinstallResult)
	// helper.InstallLog("INFO", content)
	// bar.Add(int64(70))

	execCommand = []string{"save"}
	saveResult := InstallExecuteWithPath(execPm2, "/opt/web/standalone", execCommand...)
	content = fmt.Sprintf("Pm2 save: %s", saveResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(80))

	execCommand = []string{"startup"}
	startupResult := InstallExecuteWithPath(execPm2, "/opt/web/standalone", execCommand...)
	content = fmt.Sprintf("Pm2 startup: %s", startupResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(95))

	//npx prisma generate
	execCommand = []string{"-c", "npx prisma generate"}
	prismaResult := InstallExecuteWithPath(execbash, "/opt/web/phalanx_socket_ts", execCommand...)
	content = fmt.Sprintf("prisma generate install: %s", prismaResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(96))

	//npx prisma generate
	execCommand = []string{"-c", "npx prisma generate"}
	prismaResult = InstallExecuteWithPath(execbash, "/opt/web/phalanx_schedule_ts", execCommand...)
	content = fmt.Sprintf("prisma generate install: %s", prismaResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(100))

}

func DBSetup() {
	bar.NewOption(0, 100, "Initilize")
	// 	bar.NewOption(0, 100, "Initilize")

	// 	execCommand := []string{"-i", "/config/install/phalanx-update-db_0.1_amd64.deb"}
	// 	dpkginstallResult := InstallExecute(execDpkg, execCommand...)
	// 	content := fmt.Sprintf("dpkg-installation: %s", dpkginstallResult)
	// 	helper.InstallLog("INFO", content)
	// 	bar.Add(int64(5))
	// 	// apt-key
	// 	execCommand = []string{"add", "/config/db/mongo.asc"}
	// 	aptkResult := InstallExecute(execAptKey, execCommand...)
	// 	content = fmt.Sprintf("apt-key: %s", aptkResult)
	// 	helper.InstallLog("INFO", content)
	// 	bar.Add(int64(10))

	// 	// update
	// 	execCommand = []string{"update", "-qq", "-y"}
	// 	updateResult := InstallExecute(execAptGet, execCommand...)
	// 	content = fmt.Sprintf("Update install: %s", updateResult)
	// 	helper.InstallLog("INFO", content)
	// 	bar.Add(int64(15))

	// 	//libssl
	// 	execCommand = []string{"-i", "/config/install/db/libssl1.1_1.1.1f-1ubuntu2_amd64.deb"}
	// 	dpkgupdateResult := InstallExecute(execDpkg, execCommand...)
	// 	content = fmt.Sprintf("dpkg-db: %s", dpkgupdateResult)
	// 	helper.InstallLog("INFO", content)
	// 	bar.Add(int64(20))

	// 	// mongo
	// 	execCommand = []string{"install", "-y", "-qq", "mongodb-org=4.4.15", "mongodb-org-server=4.4.15", " mongodb-org-shell=4.4.15", "mongodb-org-mongos=4.4.15", "mongodb-org-tools=4.4.15"}
	// 	toolResult := InstallExecute(execAptGet, execCommand...)
	// 	content = fmt.Sprintf("tool install: %s", toolResult)
	// 	helper.InstallLog("INFO", content)
	// 	bar.Add(int64(30))

	// 	//start service
	// 	execCommand = []string{"start", "mongod.service"}
	// 	mongoService := InstallExecute(execSystemctl, execCommand...)
	// 	content = fmt.Sprintf("network timeout: %s", mongoService)
	// 	helper.InstallLog("INFO", content)
	// 	bar.Add(int64(50))

	// 	execCommand = []string{"enable", "mongod.service"}
	// 	mongoEnable := InstallExecute(execSystemctl, execCommand...)
	// 	content = fmt.Sprintf("network timeout: %s", mongoEnable)
	// 	helper.InstallLog("INFO", content)
	// 	bar.Add(int64(50))

	// sudo usermod -aG docker $USER
	// newgrp docker

	// ## 安裝mongodb
	// sudo openssl rand -base64 128 > keyFile
	// sudo chmod 600 keyFile
	// sudo chown 999:999 keyFile

	// ##　psql安裝、pm2安裝
	// sudo apt-get update
	// sudo apt-get install postgresql-client
	// sudo apt install npm
	// sudo npm install --global pm2

	// ##　開啟postgres跟mongodb
	// sudo docker-compose up -d

	// ##　mongoDB 初始化
	// /usr/sbin/mongoInit

	// docker isntall source
	execCommand := []string{"-c", "/config/install_docker.sh"}
	docakersourceResult := InstallExecute(execbash, execCommand...)
	content := fmt.Sprintf("Docker Source install: %s", docakersourceResult)

	// Install Docker
	execCommand = []string{"install", "-y", "-qq", "docker-ce", "docker-ce-cli", "containerd.io", "docker-buildx-plugin", "docker-compose", "docker-compose-plugin"}
	dockerInstallResult := InstallExecute(execAptGet, execCommand...)
	content = fmt.Sprintf("network install: %s", dockerInstallResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(5))

	/*
		sudo groupadd docker
		sudo usermod -aG docker $USER
		newgrp docker
	*/

	// // groupadd
	// execCommand = []string{"-c", "groupadd docker"}
	// docakergroupaddResult := InstallExecute(execbash, execCommand...)
	// content = fmt.Sprintf("Groupadd install: %s", docakergroupaddResult)

	// usermod
	execCommand = []string{"-c", "usermod -aG docker $USER"}
	docakerusermodResult := InstallExecute(execbash, execCommand...)
	content = fmt.Sprintf("usermod install: %s", docakerusermodResult)

	// newgrp
	execCommand = []string{"-c", "newgrp docker"}
	docakernewgrpResult := InstallExecute(execbash, execCommand...)
	content = fmt.Sprintf("newgrp install: %s", docakernewgrpResult)

	opensslCommand := []string{"-c", "openssl rand -base64 128 > keyFile"}
	toolResult := InstallExecuteWithPath(execbash, "/opt/web/mongodb", opensslCommand...)
	content = fmt.Sprintf("tool install: %s", toolResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(10))

	chmodCommand := []string{"-c", "chmod 600 keyFile"}
	toolResult = InstallExecuteWithPath(execbash, "/opt/web/mongodb", chmodCommand...)
	content = fmt.Sprintf("tool install: %s", toolResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(20))

	chownCommand := []string{"-c", "chown 999:999 keyFile"}
	toolResult = InstallExecuteWithPath(execbash, "/opt/web/mongodb", chownCommand...)
	content = fmt.Sprintf("tool install: %s", toolResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(30))

	execCommand = []string{"docker-compose", "up", "-d"}
	mongodbcomposeResult := InstallExecuteWithPath(execSUDO, "/opt/web/mongodb", execCommand...)
	content = fmt.Sprintf("mongodb compose: %s", mongodbcomposeResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(40))

	execCommand = []string{"docker-compose", "up", "-d"}
	postgresqlcomposeResult := InstallExecuteWithPath(execSUDO, "/opt/web/postgresql", execCommand...)
	content = fmt.Sprintf("postgresql compose: %s", postgresqlcomposeResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(60))

	//mongo init
	execCommand = []string{"-c", "docker exec mongo1 /scripts/mongoinit.sh"}
	initResult := InstallExecuteWithPath(execbash, "/config/", execCommand...)
	content = fmt.Sprintf("tool install: %s", initResult)
	helper.InstallLog("INFO", content)

	//mongodb import openvas
	execCommand = []string{"-c", "python3 main.py"}
	importResult := InstallExecuteWithPath(execbash, "/config/xlsx_mongodb", execCommand...)
	content = fmt.Sprintf("tool install: %s", importResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(100))
}
