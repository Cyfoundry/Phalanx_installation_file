package adapter

type AdapterInfo struct {
	Index        int
	Name         string
	HardwareAddr string
}

type AdaptersInfo struct {
	Adapters []AdapterInfo
}

type Adapter interface {
	Adapters() (*AdaptersInfo, error)
}
