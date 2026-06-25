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

	var lastVendor Parent
	var lastDevice Parent

	scanner := bufio.NewScanner(sourceFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "C") {
			break
		}

		switch {

		case !strings.HasPrefix(line, "\t"): //vendor
			vendor := lineToVendor(line)
			table.RegisterVendor(vendor)
			lastVendor = vendor
			lastDevice = nil

		case strings.HasPrefix(line, "\t\t"): //subdevice
			subDev := lineToSubDevice(line)
			lastDevice.AddChild(subDev)
		case strings.HasPrefix(line, "\t"): //device
			device := lineToDevice(line)
			lastVendor.AddChild(device)
			lastDevice = device

		default:
			log.Println(line)
		}

	}

	return nil
}
