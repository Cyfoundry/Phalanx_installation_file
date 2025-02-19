package install

import (
	"fmt"
	"installation/internal/helper"
)

func EndpointSetup() {
	bar.NewOption(0, 100, "Initilize")

	// enable endpoint Service
	execCommand := []string{"enable", "endpoint.service"}
	enableendpointResult := InstallExecute(execSystemctl, execCommand...)
	content := fmt.Sprintf("endpoint enable %s", enableendpointResult)
	helper.InstallLog("INFO", content)

	// start endpoint Service
	execCommand = []string{"start", "endpoint.service"}
	startendpointResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("endpoint start %s", startendpointResult)
	helper.InstallLog("INFO", content)

	bar.Add(int64(100))
}
