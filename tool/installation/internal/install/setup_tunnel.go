package install

import (
	"fmt"
	"installation/internal/helper"
)

func TunnelSetup() {

	bar.NewOption(0, 100, "Initilize")

	// enable tunnel Service
	execCommand := []string{"enable", "tunnel.service"}
	enabletunnelResult := InstallExecute(execSystemctl, execCommand...)
	content := fmt.Sprintf("endpoint enable %s", enabletunnelResult)
	helper.InstallLog("INFO", content)

	// start tunnel Service
	execCommand = []string{"start", "tunnel.service"}
	starttunnelResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("endpoint start %s", starttunnelResult)
	helper.InstallLog("INFO", content)

	bar.Add(int64(100))
}
