package adapter

import (
	"fmt"
	"log"
	"net"
)

func FindAdapter(adapter string) (*AdapterInfo, error) {
	adempty := AdapterInfo{}
	var temp AdapterInfo
	ai, err := GetAdaptersList()
	if err != nil {
		return &adempty, err
	}

	for _, v := range ai.Adapters {
		if v.Name == adapter {
			temp = v
			break
		}
	}

	return &temp, nil
}

func Adapters() (*AdaptersInfo, error) {
	adempty := AdaptersInfo{}
	ai, err := GetAdaptersList()

	if err != nil {
		return &adempty, err
	}

	return &ai, nil
}

func GetAdaptersList() (AdaptersInfo, error) {
	var temp []AdapterInfo
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("%+v\n", err.Error()))
	}

	for _, i := range ifaces {
		adapter := AdapterInfo{
			Index:        i.Index,
			Name:         i.Name,
			HardwareAddr: i.HardwareAddr.String(),
		}
		temp = append(temp, adapter)

	}

	return AdaptersInfo{
		Adapters: temp,
	}, nil
}
