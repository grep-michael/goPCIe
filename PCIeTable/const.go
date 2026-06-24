package pcietable

import "fmt"

var SysPCIeTabl = &PCITable{}

func init() {
	setUpTable(SysPCIeTabl)

	commonSources := [2]string{"/usr/share/misc/pci.ids", "/usr/share/hwdata/pci.ids"}

	for _, source := range commonSources {
		err := SysPCIeTabl.LoadSource(source)
		if err != nil {
			fmt.Printf("Failed to process source %s\n", source)
		}
	}
}
