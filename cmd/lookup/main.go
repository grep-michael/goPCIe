package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/grep-michael/goPCIe/lib"
)

func main() {
	vendor := flag.String("vendor", "", "Vendor to look up")
	device := flag.String("device", "", "device lookup")

	sources := flag.String("sources", "/usr/share/misc/pci.ids,/usr/share/hwdata/pci.ids,/usr/share/misc/pci.ids", "comma separated list of source files")
	flag.Parse()

	table := &lib.PCITable{}

	for _, source := range strings.Split(*sources, ",") {
		err := lib.ProcessFile(source, table)
		if err != nil {
			fmt.Printf("Failed to process source %s\n", source)
		}
	}

	var ven *lib.Vendor
	var dev *lib.Device

	status := ""

	if *vendor == "" && *device == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *vendor != "" {
		var found bool
		ven, found = table.FindVendor(*vendor)
		if !found {
			fmt.Printf("Failed to find vendor with id \"%s\"\n", *vendor)
			os.Exit(1)
		}
		status += fmt.Sprintf("Vendor: %s (ID:%s)\n", ven.Name, ven.ID)
	}

	if *device != "" {
		if ven == nil {
			findDevices(*device, table)
		} else {
			var found bool
			dev, found = ven.FindDevice(*device)
			if !found {
				fmt.Printf("Failed to find device with id \"%s\" in vendor %s(ID: %s)\n", *device, ven.Name, ven.ID)
				os.Exit(1)
			}
			status += fmt.Sprintf("\tDevice: %s(ID:%s)\n", dev.Name, dev.ID)
		}
	}

	fmt.Println(status)
}

func findDevices(deviceId string, table *lib.PCITable) {
	devices, found := table.FindDevice(deviceId)
	if !found {
		fmt.Printf("Failed to find any devices with id \"%s\"\n", deviceId)
		os.Exit(1)
	}
	js, er := json.MarshalIndent(devices, "", "    ")
	if er != nil {
		fmt.Printf("Failed Marshal device list: %+v\n", er)
		os.Exit(1)
	}
	fmt.Println(string(js))
	os.Exit(0)
}
