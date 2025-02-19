# Makefile

WORKPATH := $(shell pwd)


INSTALLATION_PATH = installationdir
DPKGSOURCE_PATH = dpkg
PACKAGENAME = phalanx.v1.0
DPKGDIRINSTALL_PATH = $(DPKGSOURCE_PATH)/$(PACKAGENAME)/phalanx-installation_0.1_amd64
DPKGDIRINSTALL_FILE = $(DPKGSOURCE_PATH)/$(PACKAGENAME)/phalanx-installation_0.1_amd64.deb
DPKUPDATEDB_PATH = $(DPKGSOURCE_PATH)/$(PACKAGENAME)/phalanx-update-db_0.1_amd64
DPKUPDATEDB_FILE = $(DPKGSOURCE_PATH)/$(PACKAGENAME)/phalanx-update-db_0.1_amd64.deb
# Update path
UPDATEINSTALL_PATH = $(DPKGDIRINSTALL_PATH)/usr/lib/phalanx

# Build path
BUILD_MENU_PATH 			 = $(WORKPATH)/tool/installation
BUILD_GENERATOR_PATH		 = $(WORKPATH)/tool/generator
# BUILD_NETWORK_PATH			 = $(WORKPATH)/tool/golang/network
# BUILD_NETWORK_SERVICE_PATH	 = $(WORKPATH)/tool/golang/network-service
# BUILD_NFT_PATH				 = $(WORKPATH)/tool/golang/nft
# BUILD_PPPOE_PATH			 = $(WORKPATH)/tool/golang/pppoe

# Build tool path
TOOL_PATH				 = $(WORKPATH)/build
TOOL_INSTALL 			 = $(TOOL_PATH)/install
TOOL_GENERATOR 			 = $(TOOL_PATH)/sn_generator
# TOOL_NETWORK			 = $(TOOL_PATH)/phalanxNetwork
# TOOL_NETWORK_SERVICE	 = $(TOOL_PATH)/phalanxDHCP

# Limit tool
LIMIT_TOOL_PATH = $(WORKPATH)/tool/lshell

.PHONY: all build clean

all:  buildinit clean build

build: buildinit build_tool build_installation  buildfinal


buildinit:
	mkdir -p build
	mkdir -p installationdir/config/install

buildfinal:
	mv $(TOOL_INSTALL) $(WORKPATH)/installationdir
	mv $(DPKGDIRINSTALL_FILE) $(WORKPATH)/installationdir/config/install
	mv $(DPKUPDATEDB_FILE) $(WORKPATH)/installationdir/config/install
	# cp -r $(LIMIT_TOOL_PATH) $(WORKPATH)/installationdir/config
	zip -r9q installation_$(version).zip ${INSTALLATION_PATH}


build_installation:
	mv $(TOOL_GENERATOR) $(UPDATEINSTALL_PATH)
	# mv $(TOOL_NETWORK) $(UPDATEINSTALL_PATH)
	# mv $(TOOL_NETWORK_SERVICE) $(UPDATEINSTALL_PATH)
	dpkg --build $(DPKGDIRINSTALL_PATH)
	dpkg --build $(DPKUPDATEDB_PATH)
	
	


build_tool:
	cd $(BUILD_MENU_PATH);GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o  $(TOOL_INSTALL)  			    		cmd/main.go
	cd $(BUILD_GENERATOR_PATH);GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o  $(TOOL_GENERATOR)   				main.go
	# cd $(BUILD_NETWORK_PATH);GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o  $(TOOL_NETWORK)  					cmd/main.go
	# cd $(BUILD_NETWORK_SERVICE_PATH);GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o  $(TOOL_NETWORK_SERVICE) 		main.go

clean: 
	find . -name ".DS_Store" -delete
	rm -f $(DPKGSOURCE_PATH)/$(PACKAGENAME)/*.deb
	rm -f installation*.zip
	rm -rf $(WORKPATH)/installationdir/config/install/*
	find $(WORKPATH)/installationdir/ -maxdepth 1 -type f -delete
	rm -f $(TOOL_PATH)/*
