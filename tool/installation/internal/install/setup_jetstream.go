package install

import (
	"fmt"
	"installation/internal/helper"
)

func JetstreamSetup() {
	bar.NewOption(0, 100, "Initilize")
	execCommand := []string{"-i", "/config/install/nats-server-v2.9.16-amd64.deb"}
	dpkgnatsResult := InstallExecute(execDpkg, execCommand...)
	content := fmt.Sprintf("dpkg-nats: %s", dpkgnatsResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(30))

	// enable Nats Service
	execCommand = []string{"enable", "nats-server.service"}
	enablenatsResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("NAD enable %s", enablenatsResult)
	helper.InstallLog("INFO", content)

	// start Nats Service
	execCommand = []string{"start", "nats-server.service"}
	startnatsResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("NAD enable %s", startnatsResult)
	helper.InstallLog("INFO", content)

	bar.Add(int64(100))
}
