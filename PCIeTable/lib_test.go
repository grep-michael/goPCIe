package pcietable

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
	table.LoadSource("/usr/share/misc/pci.ids")
	class, ok := table.FindClass("0c0010")
	fmt.Println(ok)

	js, err := json.MarshalIndent(class, "", "   ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(js))
}
