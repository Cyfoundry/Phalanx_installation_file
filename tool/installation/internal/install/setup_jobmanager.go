package install

import (
	"fmt"
	"installation/internal/helper"
)

func JobmanagerSetup() {
	bar.NewOption(0, 100, "Initilize")

	// enable jobmanager Service
	execCommand := []string{"enable", "jobmanager.service"}
	enablejobmanagerResult := InstallExecute(execSystemctl, execCommand...)
	content := fmt.Sprintf("jobmanager enable %s", enablejobmanagerResult)
	helper.InstallLog("INFO", content)

	// start jobmanager Service
	execCommand = []string{"start", "jobmanager.service"}
	startjobmanagerResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("jobmanager start %s", startjobmanagerResult)
	helper.InstallLog("INFO", content)

	bar.Add(int64(100))
}
