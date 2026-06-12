package lib

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

// iterate pci.ids file

func ProcessFile(path string, table *PCITable) error {
	setUpTable(table)

	idFileBytes, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer idFileBytes.Close()
	table.sources = append(table.sources, path)

	var lastVendor *Vendor
	var lastDevice *Device

	scanner := bufio.NewScanner(idFileBytes)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "C") {
			break
		}

		switch {
		case line == "" || strings.HasPrefix(line, "#"):
			continue
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
