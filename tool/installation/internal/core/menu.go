package core

import (
	"fmt"
	"log"
	"os"

	"installation/internal/keyboard"

	"github.com/buger/goterm"
	"github.com/pkg/term"
)

// Raw input keycodes
const (
	up   byte = 65
	down byte = 66
)

var keys = map[byte]bool{
	up:   true,
	down: true,
}

type Menu struct {
	Prompt    string
	CursorPos int
	MenuItems []*MenuItem
	ChoiceStr string
}

type MenuItem struct {
	Text    string
	ID      string
	SubMenu *Menu
}

func NewMenu(prompt string, choicestr string) *Menu {
	return &Menu{
		Prompt:    prompt,
		MenuItems: make([]*MenuItem, 0),
		ChoiceStr: choicestr,
	}
}

func (m *Menu) AddItem(option string, id string) *Menu {
	menuItem := &MenuItem{
		Text: option,
		ID:   id,
	}

	m.MenuItems = append(m.MenuItems, menuItem)
	return m
}

func (m *Menu) renderMenuItems(redraw bool) {
	if redraw {

		fmt.Printf("\033[%dA", len(m.MenuItems)-1)
	}

	for index, menuItem := range m.MenuItems {
		var newline = "\n"
		if index == len(m.MenuItems)-1 {
			newline = ""
		}

		menuItemText := menuItem.Text
		cursor := "  "
		if index == m.CursorPos {
			// "> "
			cursor = goterm.Color(m.ChoiceStr, goterm.GREEN)
			menuItemText = goterm.Color(menuItemText, goterm.RED)
		}

		fmt.Printf("\r%s %s%s", cursor, menuItemText, newline)
	}
}

func (m *Menu) Clear() {
	fmt.Println("\033[2J")
}

func (m *Menu) Flush() {
	fmt.Println("\033[2J")
	fmt.Println("\033[H")
}

func (m *Menu) Display() string {
	defer func() {
		// Show cursor again.
		fmt.Printf("\033[?25h")
	}()

	fmt.Printf("%s\n", goterm.Color(goterm.Bold(m.Prompt)+":", goterm.CYAN))

	m.renderMenuItems(false)

	// Turn the terminal cursor off
	fmt.Printf("\033[?25l")

	kbc := keyboard.New()
	up := kbc.FindKey("up")
	down := kbc.FindKey("down")
	escape := kbc.FindKey("escape")
	enter := kbc.FindKey("enter")
	ctrlc := kbc.FindKey("ctrl-c")

	for {
		keyCode := getInput()
		if keyCode == escape {
			return ""
		} else if keyCode == enter {
			menuItem := m.MenuItems[m.CursorPos]
			fmt.Println("\r")
			return menuItem.ID
		} else if keyCode == ctrlc {
			m.Flush()
			GracefullExit()
		} else if keyCode == up {
			m.CursorPos = (m.CursorPos + len(m.MenuItems) - 1) % len(m.MenuItems)
			m.renderMenuItems(true)
		} else if keyCode == down {
			m.CursorPos = (m.CursorPos + 1) % len(m.MenuItems)
			m.renderMenuItems(true)
		}
	}
}

func GracefullExit() {
	fmt.Println("Start Exit...")
	fmt.Println("Execute Clean...")
	fmt.Println("End Exit...")
	os.Exit(0)
}

func getInput() byte {
	t, _ := term.Open("/dev/tty")

	err := term.RawMode(t)
	if err != nil {
		log.Fatal(err)
	}

	var read int
	readBytes := make([]byte, 3)
	read, err = t.Read(readBytes)

	if err != nil {
		log.Fatal(err)
	}

	t.Restore()
	t.Close()

	if read == 3 {
		if _, ok := keys[readBytes[2]]; ok {
			return readBytes[2]
		}
	} else {
		return readBytes[0]
	}

	return 0
}
