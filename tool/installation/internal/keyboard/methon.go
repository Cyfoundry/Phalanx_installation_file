package keyboard

import (
	"log"
)

const (
	up     byte = 65
	down   byte = 66
	escape byte = 27
	enter  byte = 13
	ctrlc  byte = 3
	ctrlD  byte = 4
)

func New() Keyboard {
	keys := map[string]byte{
		"up":     up,
		"down":   down,
		"escape": escape,
		"enter":  enter,
		"ctrl-c": ctrlc,
		"ctrl-d": ctrlD,
	}

	return &Keystr{
		keys: keys,
	}
}

func (k *Keystr) FindKey(codename string) byte {
	if val, ok := k.keys[codename]; ok {
		return val
	} else {
		log.Fatalln("Can't Find Key Code from list")
		return 0
	}

}
