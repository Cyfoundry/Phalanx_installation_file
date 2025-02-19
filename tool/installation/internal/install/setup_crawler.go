package install

import (
	"fmt"
	"installation/internal/helper"
)

func CrawlerSetup() {

	bar.NewOption(0, 100, "Initilize")

	// enable sqlmap_api Service
	execCommand := []string{"enable", "sqlmap_api.service"}
	enablesqlmapapiResult := InstallExecute(execSystemctl, execCommand...)
	content := fmt.Sprintf("sqlmap_api enable %s", enablesqlmapapiResult)
	helper.InstallLog("INFO", content)
	// start sqlmap_api Service
	execCommand = []string{"start", "sqlmap_api.service"}
	startsqlmapapiResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("sqlmap_api start %s", startsqlmapapiResult)
	helper.InstallLog("INFO", content)

	// enable sqlmap_server Service
	execCommand = []string{"enable", "sqlmap_server.service"}
	enablesqlmapserverResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("sqlmap_server enable %s", enablesqlmapserverResult)
	helper.InstallLog("INFO", content)

	// start sqlmap_server Service
	execCommand = []string{"start", "sqlmap_server.service"}
	startsqlmapserverResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("sqlmap_server start %s", startsqlmapserverResult)
	helper.InstallLog("INFO", content)

	// enable crawler Service
	execCommand = []string{"enable", "crawler.service"}
	enablecrawlerResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("sqlmap_server enable %s", enablecrawlerResult)
	helper.InstallLog("INFO", content)
	// start crawler Service
	execCommand = []string{"start", "crawler.service"}
	startcrawlerResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("sqlmap_server start %s", startcrawlerResult)
	helper.InstallLog("INFO", content)

	// enable sqlidetector Service
	execCommand = []string{"enable", "sqlidetector.service"}
	enablesqlidetectorResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("sqlidetector enable %s", enablesqlidetectorResult)
	helper.InstallLog("INFO", content)

	// start sqlidetector Service
	execCommand = []string{"start", "sqlidetector.service"}
	startesqlidetectorResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("sqlidetector start %s", startesqlidetectorResult)
	helper.InstallLog("INFO", content)

	// enable sqlmap Service
	execCommand = []string{"enable", "sqlmap.service"}
	enablesqlmapResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("sqlmap enable %s", enablesqlmapResult)
	helper.InstallLog("INFO", content)

	// start sqlmap Service
	execCommand = []string{"start", "sqlmap.service"}
	startsqlmapResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("sqlmap start %s", startsqlmapResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(100))
}
