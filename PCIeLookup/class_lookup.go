package pcielookup

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PCIeClassLookupFromSource(class, sourcePath string) (LookupResult, error) {
	class = strings.TrimPrefix(class, "0x")
	result := LookupResult{
		Source: sourcePath,
	}

	if len(class) != 6 {
		return result, fmt.Errorf("class string length incorrect: %q", class)
	}

	if _, err := os.Stat(sourcePath); err != nil {
		return result, err
	}

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return result, err
	}
	defer sourceFile.Close()

	classSegments := []string{class[0:2], class[2:4], class[4:6]}
	foundUpperClass := false

	matchClassLine := func(line, classSegment string) string {
		cleanedLine := strings.TrimSpace(line)
		if cleanedLine[0:2] == classSegment {
			return cleanedLine[4:]
		}
		return ""
	}

	scanner := bufio.NewScanner(sourceFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		switch {
		case strings.HasPrefix(line, "C"):
			if foundUpperClass {
				return result, nil
			}
			if classSegments[0] == line[2:4] {
				foundUpperClass = true
				result.Class = line[6:]
			}
		case foundUpperClass && strings.HasPrefix(line, "\t\t"):
			if name := matchClassLine(line, classSegments[2]); name != "" {
				result.Class = name
			}
		case foundUpperClass && strings.HasPrefix(line, "\t"):
			if name := matchClassLine(line, classSegments[1]); name != "" {
				result.Class = name
			}
		}
	}
	return result, fmt.Errorf("Failed to find class %s\n", class)
}
