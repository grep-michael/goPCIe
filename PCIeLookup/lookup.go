package pcielookup

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	sourcelist "github.com/grep-michael/goPCIe/SourceList"
)

func PCIeLookup(vendor, device string) (LookupResult, error) {
	var errs []error
	for _, source := range sourcelist.CommonSources {
		result, err := PCIeLookupFromSource(vendor, device, source)
		if err == nil {
			return result, nil
		}
		errs = append(errs, err)
	}
	return LookupResult{}, errors.Join(errs...)
}

func PCIeLookupFromSource(vendor, device, sourcePath string) (LookupResult, error) {
	vendor = strings.TrimPrefix(vendor, "0x")
	device = strings.TrimPrefix(device, "0x")
	result := LookupResult{}

	if _, err := os.Stat(sourcePath); err != nil {
		return result, err
	}

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return result, err
	}
	defer sourceFile.Close()

	foundVendor := false

	scanner := bufio.NewScanner(sourceFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "C") {
			continue
		}

		segments := strings.SplitN(strings.TrimSpace(line), " ", 2)
		if len(segments) == 2 {
			segments[0] = strings.TrimSpace(segments[0])
			segments[1] = strings.TrimSpace(segments[1])
		}

		switch {
		case line == "" || strings.HasPrefix(line, "#"):
			continue
		case !strings.HasPrefix(line, "\t"):
			if foundVendor {
				return result, fmt.Errorf("Failed to find device in Vendor")
			}
			if segments[0] == vendor {
				foundVendor = true
				result.Vendor = PCIeEntitiy{
					ID:   segments[0],
					Name: segments[1],
				}
			}
		case foundVendor && strings.HasPrefix(line, "\t"):
			if segments[0] == device {
				result.Device = PCIeEntitiy{
					ID:   segments[0],
					Name: segments[1],
				}
				result.Source = sourcePath
				return result, nil
			}
		}

	}

	return result, nil
}
