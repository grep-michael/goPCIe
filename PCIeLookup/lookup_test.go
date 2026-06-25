package pcielookup

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLookupSource(t *testing.T) {
	vendor := "10de"
	device := "13f1"

	result, err := PCIeLookupFromSource(vendor, device, "/usr/share/misc/pci.ids")
	if err != nil {
		t.Error(err)
	}
	js, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println(string(js))
}

func TestLookup(t *testing.T) {
	vendor := "10de"
	device := "13f1"

	result, err := PCIeLookup(vendor, device)
	if err != nil {
		t.Error(err)
	}
	js, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println(string(js))
}

func TestLookupClass(t *testing.T) {
	class := "0x0c0340"

	result, err := PCIeClassLookupFromSource(class, "/usr/share/misc/pci.ids")
	if err != nil {
		t.Error(err)
	}
	js, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println(string(js))
}
