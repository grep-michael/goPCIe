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

	var lastVendor *Vendor
	var lastDevice *Device

	scanner := bufio.NewScanner(sourceFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		switch {
		case strings.HasPrefix(line, "C"):

		case !strings.HasPrefix(line, "\t"): //vendor
			vendor := lineToVendor(line)
			table.RegisterVendor(vendor)
			lastVendor = vendor
			lastDevice = nil

		case strings.HasPrefix(line, "\t\t"): //subdevice
			subDev := lineToSubDevice(line)
			lastDevice.addSubDevice(subDev)

		case strings.HasPrefix(line, "\t"): //device
			device := lineToDevice(line)
			device.Vendor = lastVendor
			device.VendorID = lastVendor.ID
			lastVendor.addDevice(device, table)
			lastDevice = device

		default:
			log.Println(line)
		}

	}

	return nil
}
