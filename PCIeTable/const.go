package pcietable

import (
	"fmt"

	sourcelist "github.com/grep-michael/goPCIe/SourceList"
)

var SysPCIeTabl = &PCITable{}

func init() {
	setUpTable(SysPCIeTabl)

	for _, source := range sourcelist.CommonSources {
		err := SysPCIeTabl.LoadSource(source)
		if err != nil {
			fmt.Printf("Failed to process source %s\n", source)
		}
	}
}
