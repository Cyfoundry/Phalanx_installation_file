package install

import (
	"fmt"
	"installation/internal/adapter"
	"installation/internal/helper"
	"log"
	"os"
	"os/exec"
	"time"
)

func Setup() {
	bar.NewOption(0, 100, "Initilize")
	bar.Add(int64(1))
	execCommand := []string{"update", "-qq", "-y"}
	updateResult := InstallExecute(execAptGet, execCommand...)
	content := fmt.Sprintf("Update install: %s", updateResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(5))

	execCommand = []string{"clean", "-qq", "-y"}
	cleanAptResult := InstallExecute(execAptGet, execCommand...)
	content = fmt.Sprintf("Apt clean install: %s", cleanAptResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(10))

	execCommand = []string{"autoclean", "-qq", "-y"}
	cleanAptResult = InstallExecute(execAptGet, execCommand...)
	content = fmt.Sprintf("Apt clean install: %s", cleanAptResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(15))

	execCommand = []string{"autoremove", "--purge", "-qq", "-y"}
	cleanAptResult = InstallExecute(execAptGet, execCommand...)
	content = fmt.Sprintf("Apt clean install: %s", cleanAptResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(20))

	execCommand = []string{"upgrade", "-y", "-qq"}
	upgradeAptResult := InstallExecute(execAptGet, execCommand...)
	content = fmt.Sprintf("Apt upgrade: %s", upgradeAptResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(25))

	// create config
	helper.CheckDirC("/config/install")
	helper.CheckDirC("/config/setting")
	err := helper.CopyFileAll("./config/install/*", "/config/install")
	if err != nil {
		copyResult := err.Error()
		helper.InstallLog("ERROR", copyResult)
	}
	bar.Add(int64(30))

	// tool
	//"ipset", "ipset-persistent"
	//phalanx
	execCommand = []string{"install", "-y", "-qq", "vim", "bash-completion", "iproute2", "python3", "python3-pip", "rsyslog", "vnstat", "expect"}
	toolResult := InstallExecute(execAptGet, execCommand...)
	content = fmt.Sprintf("tool install: %s", toolResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(35))
	// network
	// execCommand = []string{"install", "-y", "-qq", "iperf", "iperf3", "iftop", "lsof", "ldnsutils", "ethtool", "iputils-ping"}
	//phalanx
	execCommand = []string{"install", "-y", "-qq", "iperf", "iperf3", "iftop", "lsof", "ldnsutils", "ethtool", "iputils-ping"}
	networkResult := InstallExecute(execAptGet, execCommand...)
	content = fmt.Sprintf("network install: %s", networkResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(40))
	// system
	//phalanx
	execCommand = []string{"install", "-y", "-qq", "curl", "htop", "lm-sensors", "tmux", "cron"}
	systemResult := InstallExecute(execAptGet, execCommand...)
	content = fmt.Sprintf("system install: %s", systemResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(45))

	// route
	//phalanx "nftables", "dnsmasq", "cpufrequtils", "intel-microcode", "amd64-microcode"
	execCommand = []string{"install", "-y", "-qq", "nftables", "dnsmasq", "cpufrequtils", "intel-microcode", "amd64-microcode"}
	routeResult := InstallExecute(execAptGet, execCommand...)
	content = fmt.Sprintf("route install: %s", routeResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(50))

	// Apt Install
	// 安裝基本工具
	execCommand = []string{"install", "-y", "-qq", "python3-pip", "pkg-config", "libcairo2-dev", "ca-certificates", "curl", "gnupg"}
	toolResult = InstallExecute(execAptGet, execCommand...)
	content = fmt.Sprintf("tool install: %s", toolResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(55))

	//install custom deb
	execCommand = []string{"-i", "/config/install/phalanx-installation_0.1_amd64.deb"}
	dpkginstallResult := InstallExecute(execDpkg, execCommand...)
	content = fmt.Sprintf("dpkg-installation: %s", dpkginstallResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(60))
	//reload service
	execCommand = []string{"daemon-reload"}
	daemonResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("systemctl daemon reload %s", daemonResult)
	helper.InstallLog("INFO", content)

	//stop check network timeout for spead
	//phalanx
	execCommand = []string{"stop", "systemd-networkd-wait-online.service"}
	timeoutsResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("network timeout: %s", timeoutsResult)
	helper.InstallLog("INFO", content)

	execCommand = []string{"disable", "systemd-networkd-wait-online.service"}
	timeoutdResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("network timeout: %s", timeoutdResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(80))
	//open mode
	//phalanx
	execCommand = []string{"ip_conntrack"}
	modeResult := InstallExecute(execMode, execCommand...)
	content = fmt.Sprintf("modprobe: %s", modeResult)
	helper.InstallLog("INFO", content)

	// sync
	//phalanx
	execCommand = []string{""}
	syncResult := InstallExecute(linuxsync, execCommand...)
	content = fmt.Sprintf("sync: %s", syncResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(70))
	bar.Add(int64(83))
	//enable nftable
	//phalanx
	execCommand = []string{"enable", "nftables"}
	nftResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("nftables: %s", nftResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(85))
	//disable dnsmasq
	//phalanx
	execCommand = []string{"disable", "dnsmasq"}
	dnsmasqDResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("dnsmasq: %s", dnsmasqDResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(89))

	//phalanx
	execCommand = []string{"/config/lshell/setup.py", "install", "--no-compile", "--install-scripts=/usr/bin/"}
	limitcmd := exec.Command(exePython3, execCommand...)
	limitcmd.Dir = "/config/lshell"
	out, err := limitcmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	content = fmt.Sprintf("limitResult: %s", string(out))
	helper.InstallLog("INFO", content)
	e := os.Remove("/etc/lshell.conf")
	if e != nil {
		log.Println(e)
	}
	limitconf := `[global]
logpath         : /var/log/lshell/
loglevel        : 2
[default]
allowed         : ['ls','echo','cd','ll','pwd','sudo']
intro           : "Hi, this is PhalanxOS\nRun help or ?\nlist commands\nUse lpath\nWatch Access Path"
forbidden       : ['>','<', '$(', '${']
warning_counter : 2
timer           : 0
path            : ['/usr']
env_path        : ':/usr/local/bin:/usr/sbin:/bin'
scp             : 0 # or 1
sftp            : 0 # or 1
overssh         : ['ls']
aliases         : {'ls':'ls --color=auto','ll':'ls -l','vi':'vim', 'mongodbUP':'sudo /usr/sbin/mongodbUP','changewanIP': 'sudo /usr/sbin/changewanIP'}
history_size    : 9999
[phalanx_user]
allowed         : ['mongodbUP','autoexe','changewanIP','ifconfig']
path            : ['/var/log','/home/phalanx_user']
env_path        : ':/usr/local/bin:/usr/sbin:/sbin:/bin:/usr/local/sbin:/ust/bin' 
home_path       : '/home/phalanx_user'
`
	helper.WriteFile("/etc/lshell.conf", limitconf)
	bar.Add(int64(90))
	//phalanx
	//cphalanx DHCP service
	execCommand = []string{"start", "phalanxDHCP"}
	phalanxDHCPSResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("DHCP %s", phalanxDHCPSResult)
	helper.InstallLog("INFO", content)
	//phalanx
	execCommand = []string{"enable", "phalanxDHCP"}
	phalanxDHCPEResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("DHCP %s", phalanxDHCPEResult)
	helper.InstallLog("INFO", content)

	//Install npm
	execCommand = []string{"install", "-y", "-qq", "npm"}
	npmInstallResult := InstallExecute(execAptGet, execCommand...)
	content = fmt.Sprintf("system install: %s", npmInstallResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(93))

	//Install pm2
	tmpNpm := execNpmUpdated
	if err := helper.CheckExist(tmpNpm); err != nil {
		tmpNpm = execNpm
	}
	execCommand = []string{"install", "--global", "pm2"}
	pm2installResult := InstallExecute(tmpNpm, execCommand...)
	content = fmt.Sprintf("tool install: %s", pm2installResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(94))

	//Install DB
	execCommand = []string{"install", "-y", "-qq", "postgresql-client"}
	dbinstallResult := InstallExecute(execAptGet, execCommand...)
	content = fmt.Sprintf("db install: %s", dbinstallResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(97))

	//install Nmap
	execCommand = []string{"-i", "/config/nmap_7.94-2_amd64.deb"}
	nmapResult := InstallExecute(execDpkg, execCommand...)
	content = fmt.Sprintf("dpkg-installation: %s", nmapResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(85))

	//Install SQLMap
	execCommand = []string{"install", "-y", "-qq", "sqlmap"}
	sqlmapResult := InstallExecute(execAptGet, execCommand...)
	content = fmt.Sprintf("sqlmap install: %s", sqlmapResult)
	helper.InstallLog("INFO", content)

	//create sn number
	execCommand = []string{""}
	SNResult := InstallExecute(execSN, execCommand...)
	content = fmt.Sprintf("Create Machine SN number: %s", SNResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(100))
}

func SetupNetwork() {
	bar.NewOption(0, 100, "Network")
	bar.Add(int64(1))
	// orgAdpater := helper.OriginalNameAdpater()
	// remove default network setting
	err := helper.RemoveGlob("/etc/netplan/*")
	if err != nil {
		log.Fatalf("Error removing files: %+v", err)
	}

	err = helper.RemoveGlob("/run/systemd/network/*")
	if err != nil {
		log.Fatalf("Error removing files: %+v", err)
	}

	execCommand := []string{"start", "systemd-networkd.service"}
	daemonResult := InstallExecute(execSystemctl, execCommand...)
	content := fmt.Sprintf("systemctl start systemd-networkd.service %s", daemonResult)
	helper.InstallLog("INFO", content)
	execCommand = []string{"enable", "systemd-networkd.service"}
	daemonResult = InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("systemctl enable systemd-networkd.service %s", daemonResult)
	helper.InstallLog("INFO", content)
	const workdir string = "/etc/systemd/network"
	mtubyte := 9612
	radpaters := helper.RenameAdpater()
	if len(radpaters) >= 4 {
		mtubyte = 1500
	}
	configlLink := `[Match]
MACAddress= {{.mac}}

[Link]
Name={{.adpater}}
`
	configWanNetwork := `[Match]
Name={{.adpater}}

[Link]
MTUBytes={{.mtubyte}}

[Network]
DHCP=yes
IPv6AcceptRA=yes
`
	configLanNetwork := `[Match]
Name={{.adpater}}

[Link]
MTUBytes={{.mtubyte}}

[Network]
Bridge=br0
`
	bar.Add(int64(10))
	// for _, orgrad := range orgAdpater {
	// 	execCommand = []string{"link", "set", orgrad.Name, "down"}
	// 	daemonResult = InstallExecute(execIP, execCommand...)
	// 	content = fmt.Sprintf("link %s (down): %s", orgrad.Name, daemonResult)
	// 	helper.InstallLog("INFO", content)
	// }
	bar.Add(int64(14))
	radpaters = helper.RenameAdpater()
	go func(radpaters []adapter.AdapterInfo, configlLink, configWanNetwork, configLanNetwork string) {
		for _, rad := range radpaters {
			// set link
			content := helper.FormatStr(configlLink, helper.Format{"mac": rad.HardwareAddr, "adpater": rad.Name, "mtubyte": mtubyte})
			savePath := fmt.Sprintf("%s/00-%s.link", workdir, rad.Name)
			helper.WriteFile(savePath, content)
			// set network
			if rad.Name == "wan0" {
				content := helper.FormatStr(configWanNetwork, helper.Format{"adpater": rad.Name, "mtubyte": mtubyte})
				savePath := fmt.Sprintf("%s/01-%s.network", workdir, rad.Name)
				helper.WriteFile(savePath, content)
			} else {
				content := helper.FormatStr(configLanNetwork, helper.Format{"adpater": rad.Name, "mtubyte": mtubyte})
				savePath := fmt.Sprintf("%s/01-%s.network", workdir, rad.Name)
				helper.WriteFile(savePath, content)
			}

		}
	}(radpaters, configlLink, configWanNetwork, configLanNetwork)
	bar.Add(int64(20))
	// set bridge
	configBridgedev := `[NetDev]
Name=br0
Kind=bridge
`
	savePath := fmt.Sprintf("%s/10-br0.netdev", workdir)
	helper.WriteFile(savePath, configBridgedev)
	bar.Add(int64(25))
	configBridgeNetwork := `[Match]
Name=br0

[Network]
IPv6AcceptRA=no
ConfigureWithoutCarrier=yes
LinkLocalAddressing=ipv6

[Address]
Address=192.168.20.1/24

[Address]
Address=fd10::1/64

[Link]
RequiredForOnline=no
`
	bar.Add(int64(30))
	savePath = fmt.Sprintf("%s/20-br0.network", workdir)
	helper.WriteFile(savePath, configBridgeNetwork)
	bar.Add(int64(40))
	configBindBridge := `[Match]
{{.bindnames}}
[Link]
MTUBytes={{.mtubyte}}

[Network]
Bridge=br0
ConfigureWithoutCarrier=true
`
	var bindnames string
	radpatersNum := len(radpaters) - 1
	if radpatersNum >= 5 {
		radpatersNum = radpatersNum - 1
	}

	//radpatersNum
	for i := 0; i <= radpatersNum; i++ {
		if radpaters[i].Name != "wan0" {
			bindnames = fmt.Sprintf("%sName=%s\n", bindnames, radpaters[i].Name)
		}
	}

	// for _, rad := range radpaters {
	// 	if rad.Name != "wan0" {
	// 		bindnames = fmt.Sprintf("%sName=%s\n", bindnames, rad.Name)
	// 	}
	// }

	content = helper.FormatStr(configBindBridge, helper.Format{"bindnames": bindnames, "mtubyte": mtubyte})
	savePath = fmt.Sprintf("%s/25-bridge-ports.network", workdir)
	helper.WriteFile(savePath, content)
	bar.Add(int64(50))
	configLoNetwork := `[Match]
Name=lo

# rfc6890
[Route]
Destination=0.0.0.0/8
Type=unreachable
[Route]
Destination=10.0.0.0/8
Type=unreachable
[Route]
Destination=100.64.0.0/10
Type=unreachable
[Route]
Destination=127.0.0.0/8
Type=unreachable
[Route]
Destination=169.254.0.0/16
Type=unreachable
[Route]
Destination=172.16.0.0/12
Type=unreachable
[Route]
Destination=192.0.0.0/24
Type=unreachable
[Route]
Destination=192.0.0.0/29
Type=unreachable
[Route]
Destination=192.0.2.0/24
Type=unreachable
[Route]
Destination=192.88.99.0/24
Type=unreachable
[Route]
Destination=192.168.0.0/16
Type=unreachable
[Route]
Destination=198.18.0.0/15
Type=unreachable
[Route]
Destination=198.51.100.0/24
Type=unreachable
[Route]
Destination=203.0.113.0/24
Type=unreachable
[Route]
Destination=240.0.0.0/4
Type=unreachable
[Route]
Destination=255.255.255.255/32
Type=unreachable
[Route]
Destination=::1/128
Type=unreachable
[Route]
Destination=::/128
Type=unreachable
[Route]
Destination=64:ff9b::/96
Type=unreachable
[Route]
Destination=64:ff9b::/96
Type=unreachable
[Route]
Destination=::ffff:0:0/96
Type=unreachable
[Route]
Destination=100::/64
Type=unreachable
[Route]
Destination=2001::/23
Type=unreachable
[Route]
Destination=2001::/32
Type=unreachable
[Route]
Destination=2001:2::/48
Type=unreachable
[Route]
Destination=2001:db8::/32
Type=unreachable
[Route]
Destination=2001:10::/28
Type=unreachable
[Route]
Destination=2002::/16
Type=unreachable
[Route]
Destination=fc00::/7
Type=unreachable
[Route]
Destination=fe80::/10
Type=unreachable

[Route]
Table=1024
Destination=fc00::/7
Type=throw
[Route]
Table=1024
Destination=::/0
Type=prohibit
[Route]
Table=1024
Destination=10.0.0.0/10
Type=throw
[Route]
Table=1024
Destination=172.16.0.0/12
Type=throw
[Route]
Table=1024
Destination=192.168.0.0/16
Type=throw
[Route]
Table=1024
Destination=0.0.0.0/0
Type=prohibit

[RoutingPolicyRule]
Priority=10000
From=fc00::/7
Table=1024
`
	savePath = fmt.Sprintf("%s/50-lo.network", workdir)
	helper.WriteFile(savePath, configLoNetwork)
	bar.Add(int64(60))
	configStaticOnu := `[Match]
Name=wan0

[Network]
LinkLocalAddressing=no

[Address]
Address=192.168.20.3/24

[Link]
RequiredForOnline=no
`
	savePath = fmt.Sprintf("%s/50-static-onu.network", workdir)
	helper.WriteFile(savePath, configStaticOnu)
	bar.Add(int64(70))

	configPpoeNetwork := `[Match]
Name=ppp0
Type=ppp

[Network]
LLMNR=no

[IPv6AcceptRA]
DHCPv6Client=always
UseDNS=no

[Link]
RequiredForOnline=no
`

	savePath = fmt.Sprintf("%s/75-pppoe.network", workdir)
	helper.WriteFile(savePath, configPpoeNetwork)
	bar.Add(int64(80))

	configOffload := `[Match]
OriginalName=*

[Link]
NamePolicy=keep kernel database onboard slot path
AlternativeNamesPolicy=database onboard slot path
MACAddressPolicy=persistent
GenericSegmentationOffload=no
GenericReceiveOffload=no
TCPSegmentationOffload=no
TCP6SegmentationOffload=no
LargeReceiveOffload=no
`
	savePath = fmt.Sprintf("%s/90-offload.link", workdir)
	helper.WriteFile(savePath, configOffload)
	bar.Add(int64(90))
	execCommand = []string{"restart", "systemd-networkd.service"}
	daemonResult = InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("systemctl restart systemd-networkd.service %s", daemonResult)
	helper.InstallLog("INFO", content)
	execCommand = []string{"enable", "systemd-networkd.service"}
	daemonResult = InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("systemctl enable systemd-networkd.service %s", daemonResult)
	helper.InstallLog("INFO", content)

	execCommand = []string{"stop", "systemd-networkd-wait-online.service"}
	timeoutsResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("network timeout: %s", timeoutsResult)
	helper.InstallLog("INFO", content)

	bar.Add(int64(100))
	bar.Finish()
}

func SetupHost() {
	bar.NewOption(0, 100, "Host")
	bar.Add(int64(1))
	execCommand := []string{"hostname", "plocal"}
	syncResult := InstallExecute(exeHostname, execCommand...)
	content := fmt.Sprintf("Set Hostname: %s", syncResult)
	bar.Add(int64(25))
	helper.InstallLog("INFO", content)

	content = `127.0.0.1 localhost
127.0.1.1 plocal
127.0.0.1 mongo1
# The following lines are desirable for IPv6 capable hosts
::1     ip6-localhost ip6-loopback
fe00::0 ip6-localnet
ff00::0 ip6-mcastprefix
ff02::1 ip6-allnodes
ff02::2 ip6-allrouters
`
	helper.WriteFile("/etc/hosts", content)
	bar.Add(int64(100))
	bar.Finish()
}

func SetupResolve() {
	bar.NewOption(0, 100, "Resolve")
	bar.Add(int64(1))

	e := os.Remove("/etc/resolv.conf")
	if e != nil {
		log.Println(e)
	}

	bar.Add(int64(20))
	content := `nameserver 127.0.0.1
nameserver 192.168.20.1
`
	bar.Add(int64(25))
	helper.WriteFile("/etc/resolv.conf", content)
	content = `nameserver 8.8.8.8
`
	helper.WriteFile("/etc/resolv.dnsmasq", content)
	bar.Add(int64(100))
	bar.Finish()
}

func SetupLimitUser() {
	bar.NewOption(0, 100, "Limit User")
	bar.Add(int64(1))
	// gernerate password
	var username string = "phalanx_user"
	username_n := fmt.Sprintf("%s\n", username)
	helper.WriteFile("/config/limit", username_n)
	time.Sleep(1000 * time.Millisecond)
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 1
	passwordLength := 8
	password := helper.GeneratePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
	bar.Add(int64(15))
	helper.AddUser("/usr/bin/lshell", username)
	bar.Add(int64(30))
	helper.ChPassword(username, password)
	bar.Add(int64(60))
	// os.Chmod("/etc/profile.d/02-limit.sh", 0755)
	// helper.ChangeOWN("/etc/profile.d/02-limit.sh", "root", "root")
	bar.Add(int64(100))
	fmt.Printf("\nLimit-User-password:%s", password)
	time.Sleep(30000 * time.Millisecond)
	bar.Finish()
}

func SetupModules() {
	bar.NewOption(0, 100, "Modules")
	bar.Add(int64(1))
	content := `# This configuration 
# Optimize netfilter related modules at system boot

nf_conntrack
`
	bar.Add(int64(25))
	helper.WriteFile("/etc/modules-load.d/server_modules.conf", content)
	bar.Add(int64(100))
	bar.Finish()
}

func SetupSysctl() {
	bar.NewOption(0, 100, "Sysctl")
	bar.Add(int64(1))
	content := `net.ipv4.ip_forward = 1
net.ipv4.conf.default.forwarding = 1
net.ipv4.conf.all.forwarding = 1
net.ipv4.tcp_retries2=8
`
	helper.WriteFile("/etc/sysctl.d/01-ipv4-forwarding.conf", content)
	// ###########################################################################
	content = `net.ipv6.conf.default.forwarding = 1
net.ipv6.conf.all.forwarding = 1
`
	helper.WriteFile("/etc/sysctl.d/01-ipv6-forwarding.conf", content)
	bar.Add(int64(15))
	// ###########################################################################
	content = `net.ipv4.conf.all.promote_secondaries = 1
`
	helper.WriteFile("/etc/sysctl.d/02-promote-secondaries.conf", content)
	bar.Add(int64(30))
	// ###########################################################################
	content = `# Set ARP sysctls as appropriate for a router.
#
# ARPs for each interface are answered based on whether the kernel would route
# a packet from the ARPed IP out that interface.
net.ipv4.conf.all.arp_filter=1
net.ipv4.conf.default.arp_filter=1

# When sending ARP requests, always use the best local address for the target
# as the source address.
net.ipv4.conf.all.arp_announce=2
net.ipv4.conf.default.arp_announce=2

# Reply to ARP only if the target IP address is a local address configured on
# the incoming interface. Otherwise the WANs will answer ARP for LAN IPs and
# LANs will answer ARP for different VLANs and WAN IPs.
net.ipv4.conf.all.arp_ignore=1
net.ipv4.conf.default.arp_ignore=1
`

	helper.WriteFile("/etc/sysctl.d/05-arp.conf", content)
	bar.Add(int64(45))
	// ###########################################################################

	content = `net.netfilter.nf_conntrack_acct=0
net.netfilter.nf_conntrack_tcp_timeout_established = 7440
net.netfilter.nf_conntrack_udp_timeout = 60
net.netfilter.nf_conntrack_udp_timeout_stream = 180
net.netfilter.nf_conntrack_buckets=32768
net.netfilter.nf_conntrack_expect_max=2048
net.netfilter.nf_conntrack_max=262144
net.netfilter.nf_conntrack_tcp_be_liberal=1
net.netfilter.nf_conntrack_tcp_timeout_max_retrans=3600
net.netfilter.nf_conntrack_tcp_timeout_unacknowledged=3600
`
	helper.WriteFile("/etc/sysctl.d/10-set-default_nf_conntrack.conf", content)
	bar.Add(int64(60))
	// ###########################################################################
	content = `net.ipv4.conf.all.accept_redirects = 0
net.ipv4.conf.all.rp_filter = 1
net.ipv4.conf.default.rp_filter = 1
net.ipv4.conf.all.log_martians = 1
net.ipv4.icmp_ratelimit = 100
net.ipv4.igmp_max_memberships = 100
net.ipv4.route.error_burst = 500
net.ipv4.route.error_cost = 100
net.ipv4.route.redirect_load = 2
net.ipv4.route.redirect_silence = 2048
net.ipv4.tcp_fin_timeout = 30
net.ipv4.tcp_keepalive_time = 120
net.ipv4.tcp_max_orphans = 4096
net.ipv4.tcp_max_tw_buckets = 4096
net.ipv4.tcp_syncookies = 1
net.ipv6.conf.all.accept_ra = 0
net.ipv6.conf.default.accept_ra = 0
net.ipv6.conf.all.use_tempaddr = 0
net.ipv6.conf.default.use_tempaddr = 0
`
	helper.WriteFile("/etc/sysctl.d/99-all.conf", content)
	bar.Add(int64(75))
	// ###########################################################################
	execCommand := []string{"-p"}
	syncResult := InstallExecute(exeSysctl, execCommand...)
	content = fmt.Sprintf("linux system setting apply: %s", syncResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(100))
	bar.Finish()
}

func SetupBanner() {
	bar.NewOption(0, 100, "Banner")
	bar.Add(int64(1))
	err := helper.RemoveGlob("/etc/update-motd.d/*")
	if err != nil {
		log.Fatalf("Error removing files: %+v", err)
	}
	bar.Add(int64(30))
	content := `#!/bin/sh
echo '                                                                            
              ...........               
         ...:..:.......  .::.           
       .:.....:     .:     .::..        
     .:.......:      . .    :...:.      
    :.........:    =*..++   . ...::     
   :...:::.    ..  .-=+-+. .   ...::    
  :...:.  ..::..     :=+-     .....:.   
  :..:   .::---:..      .    .:---:.:   
  :.:.  ..::--::.       .    ..::.. .   
  :.:                   .     ...   .   
  :.:                   .   .....:  .   
   ::      .            . :-:=+++=:.    
    ::       .          ..*.+++++++:    
     ::                 .:+:=+++--=:    
      .=::..              .---=    .    
       -==+++*+++=========++++=.   :    
       :....::::.....=----  :.:::.:..   
       :......:.     :::=.   :.:...:.   
       ::........            ..:...   ' 
`
	helper.WriteFile("/etc/update-motd.d/00-custom-welcome", content)
	execCommand := []string{"+x", "/etc/update-motd.d/00-custom-welcome"}
	updateerr := InstallExecute(execChmod, execCommand...)
	content = fmt.Sprintf("change  /etc/update-motd.d/00-custom-welcome: %s", updateerr)
	helper.InstallLog("INFO", content)
	time.Sleep(1 * time.Second)
	bar.Add(int64(40))
	os.Remove("/etc/issue")
	content = `Date: \d \t
Phalanx
`
	helper.WriteFile("/etc/issue", content)
	helper.InstallLog("INFO", "write /etc/issue")
	time.Sleep(1 * time.Second)
	bar.Add(int64(60))
	os.Remove("/etc/issue.net")
	content = `Date: \d \t
Phalanx
`
	helper.WriteFile("/etc/issue.net", content)
	helper.InstallLog("INFO", "write /etc/issue")
	time.Sleep(1 * time.Second)
	bar.Add(int64(100))
	bar.Finish()
}

func SetupDhcp() {
	bar.NewOption(0, 100, "DHCP")
	bar.Add(int64(1))
	bridge, err := adapter.FindAdapter("br0")
	if err != nil {
		content := fmt.Sprintf("Find Bridge Info: %s", err.Error())
		helper.InstallLog("ERROR", content)
		return
	}
	bar.Add(int64(20))
	configDHCP := `no-dhcp-interface=lo
no-ping
interface=br0
dhcp-host=set:net_Default_br0_192-168-10-0-24,{{.bridgeMac}},192.168.20.129
domain=local,192.168.20.0/24
dhcp-broadcast=tag:needs-broadcast
dhcp-range=set:net_Default_br0_192-168-10-0-24,192.168.20.6,192.168.20.254,255.255.255.0,86400
dhcp-option=tag:net_Default_br0_192-168-10-0-24,option:router,192.168.20.1
dhcp-option=tag:net_Default_br0_192-168-10-0-24,option:dns-server,192.168.20.1
dhcp-option=tag:net_Default_br0_192-168-10-0-24,option:classless-static-route,192.168.20.0/24,192.168.20.1
dhcp-option=tag:net_Default_br0_192-168-10-0-24,28,192.168.20.255
`
	bar.Add(int64(40))
	content := helper.FormatStr(configDHCP, helper.Format{"bridgeMac": bridge.HardwareAddr})
	savePath := "/etc/dnsmasq.d/dhcp/dhcp.dhcpServers-net_Default_br0_192-168-10-0-24.conf"
	helper.WriteFile(savePath, content)
	bar.Add(int64(100))
	bar.Finish()
}

func SetupNftable() {
	bar.NewOption(0, 100, "Nftables")
	bar.Add(int64(1))
	helper.CheckDirC("/var/log/nftables")
	helper.ChownR("/var/log/nftables", "syslog", "adm")
	var adapters string
	radpaters := helper.RenameAdpater()
	radpatersNum := len(radpaters) - 1
	if radpatersNum >= 5 {
		radpatersNum = radpatersNum - 1
	}

	for i := 0; i <= radpatersNum; i++ {
		if radpaters[i].Name != "br0" {
			if i == 0 {
				adapters = radpaters[i].Name
			} else {
				adapters = fmt.Sprintf("%s,%s", adapters, radpaters[i].Name)
			}
		}
	}

	//wan0,eth0,eth1,eth2
	bar.Add(int64(30))
	content := `#!/usr/sbin/nft -f
	flush ruleset
	table inet phalanx {
			#
			# Flowtable
			#

			flowtable ft {
					hook ingress priority filter;
					devices = { {{ .adapters}} };   
			}


			#
			# Defines
			#

			define local_dns_ipv4 = { 192.168.20.1}
			define local_dns_ipv6 = { fd10::1 }


			#
			# Filter rules
			#

			chain input {
					type filter hook input priority filter; policy accept;
					iifname "lo" accept comment "defconf: Accept traffic from loopback"
					ct state established,related counter accept comment "defconf: Allow inbound established and related flows"
					ct state invalid counter log prefix "Drop" counter drop comment "defconf: Drop flows with invalid conntrack state"
					tcp flags & (fin | syn | rst | ack) == syn counter jump syn_flood comment "defconf: Rate limit TCP syn packets"
					iifname "br0" jump input_lan comment "defconf: Handle lan IPv4/IPv6 input traffici"
					iifname "wan0" jump input_wan comment "defconf: Handle wan IPv4/IPv6 input traffic"
			}

			chain forward {
					type filter hook forward priority filter; policy accept;
					ct state established,related accept comment "defconf: Allow forwarded established and related flows"
					ct state invalid counter drop comment "defconf: Drop flows with invalid conntrack state"
					iifname "br0" jump forward_lan comment "defconf: Handle lan IPv4/IPv6 forward traffic"
					iifname "wan0" jump forward_wan comment "defconf: Handle wan IPv4/IPv6 forward traffic"
			}

			chain output {
					type filter hook output priority filter; policy accept;
					oifname "lo" accept comment "defconf: Accept traffic towards loopback"
					ct state established,related counter accept comment "defconf: Allow outbound established and related flows"
					ct state invalid counter log prefix "Drop"  drop comment "defconf: Drop flows with invalid conntrack state"
					oifname "br0" jump output_lan comment "defconf: Handle lan IPv4/IPv6 output traffic"
					oifname "wan0" jump output_wan comment "defconf: Handle wan IPv4/IPv6 output traffic"
			}

			chain prerouting {
					type filter hook prerouting priority filter; policy accept;
					iifname "br0" jump helper_lan comment "defconf: Handle lan IPv4/IPv6 helper assignment"
			}

			chain handle_reject {
					meta l4proto tcp reject with tcp reset comment "defconf: Reject TCP traffic"
					counter reject comment "defconf: Reject any other traffic"
			}

			chain syn_flood {
					limit rate 25/second burst 50 packets return comment "defconf: Accept SYN packets below rate-limit"
					counter log prefix "Drop" drop comment "defconf: Drop excess packets"
			}

			chain input_lan {
					ct status dnat counter accept comment "lanconf: Accept port redirections"
					jump accept_from_lan
			}

			chain output_lan {
					jump accept_to_lan
			}

			chain forward_lan {
					jump accept_to_wan comment "defconf: Accept lan to wan forwarding"
					ct status dnat counter accept comment "lanconf: Accept port forwards"
					jump accept_to_lan
			}

			chain helper_lan {
			}

			chain accept_from_lan {
					iifname "br0" counter accept comment "defconf: Accept lan IPv4/IPv6 traffic"
					#tcp dport { http, https, sip, snmp, ssh , ftp, telnet } iifname { "br0", "wan0" }  counter accept
			}

			chain accept_to_lan {
					oifname "br0" counter accept comment "defconf: Accept lan IPv4/IPv6 traffic"
			}

			chain input_wan {
					#tcp dport { http, https, sip, snmp, ssh , ftp, telnet } counter accept
					meta nfproto ipv4 tcp dport 1-65535 counter accept comment "defconf: All ipv4 tcp port"
				meta nfproto ipv4 udp dport 1-65535 counter accept comment "defconf: All ipv4 udp port"
				meta nfproto ipv6 udp dport 1-65535 counter accept comment "defconf: All ipv6 tcp port"
				meta nfproto ipv6 tcp dport 1-65535 counter accept comment "defconf: All ipv6 udp port"
					meta nfproto ipv4 udp dport 68 counter accept comment "defconf: Allow-DHCP-Renew"
					meta nfproto ipv4 icmp type echo-request counter accept comment "defconf: Drop-ICMP-Ping-Input"
					meta nfproto ipv6 icmpv6 type echo-request counter accept comment "defconf: Drop-ICMPv6-Ping-Input"
					meta nfproto ipv4 meta l4proto igmp counter accept comment "defconf: Allow-IGMP"
					meta nfproto ipv6 udp dport 546 counter accept comment "defconf: Allow-DHCPv6"
					ip6 saddr fe80::/10 icmpv6 type . icmpv6 code { mld-listener-query . no-route, mld-listener-report . no-route, mld-listener-done . no-route, mld2-listener-report . no-route } counter accept comment "defconf: Allow-MLD"
					meta nfproto ipv6 icmpv6 type { destination-unreachable, time-exceeded, echo-request, echo-reply, nd-router-solicit, nd-router-advert } limit rate 100/second counter accept comment "defconf: Allow-ICMPv6-Input"
					meta nfproto ipv6 icmpv6 type . icmpv6 code { packet-too-big . no-route, parameter-problem . no-route, nd-neighbor-solicit . no-route, nd-neighbor-advert . no-route, parameter-problem . admin-prohibited } limit rate 100/second counter accept comment "defconf: Allow-ICMPv6-Input"
					jump drop_from_wan
			}

			chain output_wan {
					jump accept_to_wan
			}

			chain forward_wan {
					meta nfproto ipv6 icmpv6 type echo-request counter drop comment "defconf: Drop-ICMPv6-Ping-Forward"
					meta nfproto ipv6 icmpv6 type { destination-unreachable, time-exceeded, echo-request, echo-reply } limit rate 100/second counter accept comment "defconf: Allow-ICMPv6-Forward"
					meta nfproto ipv6 icmpv6 type . icmpv6 code { packet-too-big . no-route, parameter-problem . no-route, parameter-problem . admin-prohibited } limit rate 100/second counter accept comment "defconf: Allow-ICMPv6-Forward"
					meta l4proto esp counter jump accept_to_lan comment "defconf: Allow-IPSec-ESP"
					udp dport 500 counter jump accept_to_lan comment "defconf: Allow-ISAKMP"
					jump drop_to_wan
			}

			chain accept_to_wan {
					oifname { "br0" , "wan0" } counter accept comment "defconf: Accept br0 IPv4/IPv6 traffic"
					oifname "wan0" counter accept comment "defconf: Accept wan IPv4/IPv6 traffic"
			}

			chain drop_from_wan {
					iifname "wan0" counter drop comment "defconf: Drop wan IPv4/IPv6 traffic"
			}

			chain drop_to_wan {
					oifname "wan0" counter drop comment "defconf: Drop wan IPv4/IPv6 traffic"
			}

			#
			# NAT rules
			#
			chain dstnat {
					type nat hook prerouting priority dstnat; policy accept;
					jump nat_block_redirect
					iifname "br0" meta l4proto { tcp, udp } th dport domain counter jump dstnat_lan comment "!fw4: Handle lan IPv4/IPv6 dstnat traffic"
			}

			chain srcnat {
					type nat hook postrouting priority srcnat; policy accept;
					oifname "wan0" jump srcnat_wan comment "defconf: Handle wan IPv4/IPv6 srcnat traffic"
			}

			chain dstnat_lan {
					ip saddr $local_dns_ipv4 meta l4proto { tcp, udp } th dport domain counter accept comment "lanconf: Accept lan dns IPv4 bootstrap query"
					ip6 saddr $local_dns_ipv6 meta l4proto { tcp, udp } th dport domain counter accept comment "lanconf: Accept lan dns IPv6 bootstrap query"
					meta l4proto { tcp, udp } th dport domain counter redirect to domain comment "lanconf: Lan dns redirect"
			}

			chain srcnat_wan {
					meta nfproto ipv4 counter masquerade comment "defconf: Masquerade IPv4 wan traffic"
			}

			chain nat_block_redirect {
			}

			#
			# Raw rules (notrack)
			#

			chain raw_prerouting {
					type filter hook prerouting priority raw; policy accept;
			}

			chain raw_output {
					type filter hook output priority raw; policy accept;
			}


			#
			# Mangle rules
			#

			chain mangle_prerouting {
					type filter hook prerouting priority mangle; policy accept;
			}

			chain mangle_postrouting {
					type filter hook postrouting priority mangle; policy accept;
			}

			chain mangle_input {
					type filter hook input priority mangle; policy accept;
			}


			chain mangle_output {
					type route hook output priority mangle; policy accept;
			}

			chain mangle_forward {
					type filter hook forward priority mangle; policy accept;
					iifname "wan0" tcp flags syn tcp option maxseg size set rt mtu comment "defconf: Zone wan IPv4/IPv6 ingress MTU fixing"
					oifname "wan0" tcp flags syn tcp option maxseg size set rt mtu comment "defconf: Zone wan IPv4/IPv6 egress MTU fixin"
			}
	}


`
	content = helper.FormatStr(content, helper.Format{"adapters": adapters})
	helper.WriteFile("/etc/nftables.conf", content)

	content = `:msg,regex,"invalid" -/var/log/nftables/invalid.log
:msg,regex,"defconf: Allow inbound" -/var/log/nftables/allowInbound.log
:msg,regex,"defconf: Allow forwarded" -/var/log/nftables/allowForwarded.log
:msg,regex,"defconf: Allow outbound" -/var/log/nftables/allowOutbound.log
:msg,regex,"Drop" -/var/log/nftables/drop.log
:msg,regex,"Accept" -/var/log/nftables/accept.log
:msg,regex,"IN=.*OUT=.*SRC=.*DST=.*" -/var/log/nftables/nftables_all.log
:msg,regex,"New SSH connection: "  -/var/log/nftables/ssh.log
`
	helper.WriteFile("/etc/rsyslog.d/nftables.conf", content)

	content = `/var/log/nftables/* { rotate 5 daily  dateext create 644 root root  maxsize 50M missingok notifempty delaycompress compress postrotate invoke-rc.d rsyslog rotate > /dev/null endscript }`
	helper.WriteFile("/etc/logrotate.d/nftables", content)
	bar.Add(int64(35))

	bar.Add(int64(40))
	content = `[Unit]
Description=nftables
Documentation=man:nft(8) http://wiki.nftables.org
After=systemd-networkd.service
Wants=network-pre.target
Before=network-online.target shutdown.target
#Before=network-pre.target shutdown.target
Conflicts=shutdown.target
DefaultDependencies=no

[Service]
Type=oneshot
RemainAfterExit=yes
Restart=on-failure
RestartSec=5
TimeoutStartSec=10
StandardInput=null
ProtectSystem=full
ProtectHome=true
ExecStart=/usr/sbin/nft -f /etc/nftables.conf
ExecReload=/usr/sbin/nft -f /etc/nftables.conf
ExecStop=/usr/sbin/nft flush ruleset

[Install]
WantedBy=sysinit.target
`
	bar.Add(int64(50))
	helper.WriteFile("/lib/systemd/system/nftables.service", content)
	execCommand := []string{"daemon-reload"}
	daemonResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("systemctl daemon reload %s", daemonResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(60))
	execCommand = []string{"restart", "rsyslog"}
	rsyslogResult := InstallExecute(execSystemctl, execCommand...)
	content = fmt.Sprintf("systemctl restart rsyslog %s", rsyslogResult)
	helper.InstallLog("INFO", content)
	bar.Add(int64(80))
	bar.Add(int64(100))
	bar.Finish()
}

func Reboot() {
	const execReboot = "/usr/sbin/reboot"
	execCommand := []string{""}
	rebootResult := InstallExecute(execReboot, execCommand...)
	content := fmt.Sprintf("Reboot System%s", rebootResult)
	helper.InstallLog("INFO", content)
}
