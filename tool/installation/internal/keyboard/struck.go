package keyboard

type Keystr struct {
	keys map[string]byte
}

type Keyboard interface {
	FindKey(codename string) byte
}
