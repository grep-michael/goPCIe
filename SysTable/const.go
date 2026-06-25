package systable

import (
	"fmt"

	pcietable "github.com/grep-michael/goPCIe/PCIeTable"
	sourcelist "github.com/grep-michael/goPCIe/SourceList"
)

var SysPCIeTabl = &pcietable.PCITable{}

func init() {
	for _, source := range sourcelist.CommonSources {
		err := SysPCIeTabl.LoadSource(source)
		if err != nil {
			fmt.Printf("Failed to process source %s\n", source)
		}
	}
}
