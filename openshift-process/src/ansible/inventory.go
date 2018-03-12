package ansible

import (
	"os"
	"strings"
)

type Inventory struct {
	filename string

	sections []InventorySection
}

type InventorySection struct {
	name string

	lines []string
}

func NewInventory(filename string) *Inventory {
	inventory := &Inventory{
		filename: filename,
	}

	return inventory
}

func (inventory *Inventory) AddSection(name string, lines []string) *Inventory {
	section := InventorySection{
		name: name,
		lines: lines,
	}

	inventory.sections = append(inventory.sections, section)

	return inventory
}

func (inventory *Inventory) Write() {
	f, err := os.Create(inventory.filename)
	if err != nil {
		panic(nil)
	}
	defer f.Close()

	f.WriteString(inventory.ToString())
	f.Sync()
}

func (inventory *Inventory) ToString() string {
	var s string

	for _, section := range inventory.sections {
		s += "[" + section.name + "]\n"
		s += strings.Join(section.lines, "\n")
		s+= "\n\n"
	}

	return s
}