package lib

import (
	"strings"
)

func setUpTable(table *PCITable) {
	if table.Vendors == nil {
		table.Vendors = make(map[string]*Vendor)
	}
	if table.Devices == nil {
		table.Devices = make(map[string][]*Device)
	}
	if table.Sources == nil {
		table.Sources = make([]string, 0)
	}
}

func lineToDevice(line string) *Device {
	line = strings.TrimSpace(line)
	lines := strings.SplitN(line, "  ", 2)
	return &Device{
		ID:         lines[0],
		Name:       lines[1],
		SubDevices: make([]*Device, 0),
	}
}
func lineToSubDevice(line string) *Device {
	line = strings.TrimSpace(line)
	parts := strings.SplitN(line, "  ", 2)
	ids := strings.Split(parts[0], " ")
	return &Device{
		ID:         ids[1],
		VendorID:   ids[0],
		Name:       parts[1],
		SubDevices: make([]*Device, 0),
	}
}

func lineToVendor(line string) *Vendor {
	line = strings.TrimSpace(line)
	lines := strings.SplitN(line, " ", 2)
	return &Vendor{
		Name:    lines[1],
		ID:      lines[0],
		Devices: make(map[string]*Device),
	}
}
