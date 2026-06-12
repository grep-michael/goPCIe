package lib

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLineToDev(t *testing.T) {
	line := "        0b60  NVMe DC SSD [Sentinel Rock Plus controller]"

	dev := lineToDevice(line)
	fmt.Printf("%+v\n", dev)
}

func TestLineToSubDev(t *testing.T) {
	line := "                025e d81d  NVMe DC SSD E1.L 9.5mm [D5-P5336]"
	dev := lineToSubDevice(line)
	fmt.Printf("%+v\n", dev)
}

func TestFile(t *testing.T) {
	table := &PCITable{}
	ProcessFile("/usr/share/misc/pci.ids", table)
	js, err := json.MarshalIndent(table, "", "   ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(js))
}
