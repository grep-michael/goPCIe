package pcietable

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

// iterate pci.ids file

func (table *PCITable) LoadSource(path string) error {
	setUpTable(table)

	sourceFile, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	table.Sources = append(table.Sources, path)

	//var lastVendor *Vendor
	//var lastDevice *Device

	var parent Parent
	var lastChild Parent

	scanner := bufio.NewScanner(sourceFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		switch {
		case strings.HasPrefix(line, "C"): //class parent
			class := lineToClass(line)
			table.RegisterClass(class)
			parent = class
			lastChild = nil

		case !strings.HasPrefix(line, "\t"): //vendor parent
			vendor := lineToVendor(line)
			table.RegisterVendor(vendor)
			parent = vendor
			lastChild = nil

		case strings.HasPrefix(line, "\t\t"): //sub child
			lastChild.AddChild(line)

		case strings.HasPrefix(line, "\t"): //child
			device := parent.AddChild(line)
			lastChild = device

		default:
			log.Println(line)
		}

	}

	return nil
}
