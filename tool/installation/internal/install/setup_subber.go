package install

import (
	"fmt"
	"installation/internal/helper"
)

func SubberSetup() {

	bar.NewOption(0, 100, "Initilize")

	execCommand := []string{"install", "simplejson"}
	PipinstallResult := InstallExecute(execPip, execCommand...)
	content := fmt.Sprintf("Pip3  install %s", PipinstallResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(40))

	// enable nmap_1 Service
	execCommand = []string{"enable", "nmap_subber@1.service"}
	enablenmapResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("Nmap_1 enable %s", enablenmapResult)
	helper.InstallLog("INFO", content)

	// enable nmap_2 Service
	execCommand = []string{"enable", "nmap_subber@2.service"}
	enablenmapResult = InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("Nmap_2 enable %s", enablenmapResult)
	helper.InstallLog("INFO", content)

	// start nmap_1 Service
	execCommand = []string{"start", "nmap_subber@1.service"}
	startnmapResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("Nmap_1 start %s", startnmapResult)
	helper.InstallLog("INFO", content)

	// start nmap_2 Service
	execCommand = []string{"start", "nmap_subber@1.service"}
	startnmapResult = InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("Nmap_2 start %s", startnmapResult)
	helper.InstallLog("INFO", content)

	// enable nuclei_1 Service
	execCommand = []string{"enable", "nuclei_subber@1.service"}
	enablenucleiResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("nuclei_1 enable %s", enablenucleiResult)
	helper.InstallLog("INFO", content)

	// enable nuclei_2 Service
	execCommand = []string{"enable", "nuclei_subber@2.service"}
	enablenucleiResult = InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("Nmap_2 enable %s", enablenucleiResult)
	helper.InstallLog("INFO", content)

	// start nuclei_1 Service
	execCommand = []string{"start", "nuclei_subber@1.service"}
	startnucleiResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("Nmap_1 start %s", startnucleiResult)
	helper.InstallLog("INFO", content)

	// start nuclei_2 Service
	execCommand = []string{"start", "nuclei_subber@2.service"}
	startnucleiResult = InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("Nmap_2 start %s", startnucleiResult)
	helper.InstallLog("INFO", content)

	// install metasploit
	execCommand = []string{"-i", "/config/metasploit-framework_6.4.10_amd64.deb"}
	msfinstallResult := InstallExecute(execDpkg, execCommand...)
	content = fmt.Sprintf("dpkg-installation: %s", msfinstallResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(60))

	//init msf
	execCommand = []string{"-c", "sudo -u phalanx /config/init_msf.exp"}
	toolResult := InstallExecute(execbash, execCommand...)
	content = fmt.Sprintf("msf install %s", toolResult)
	helper.InstallLog("INFO", content)

	// enable mrpc Service
	execCommand = []string{"enable", "mrpc.service"}
	enablemrpcResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("msf enable %s", enablemrpcResult)
	helper.InstallLog("INFO", content)

	// start mrpc Service
	execCommand = []string{"start", "mrpc.service"}
	startmrpcResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("msf start %s", startmrpcResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(100))

	// enable msf Service
	execCommand = []string{"enable", "msf.service"}
	enablemsfResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("msf enable %s", enablemsfResult)
	helper.InstallLog("INFO", content)

	// start msf Service
	execCommand = []string{"start", "msf.service"}
	startmsfResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("msf start %s", startmsfResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(100))
}
