package go_crud

import (
	"os"

	"path/filepath"

	"regexp"
)

// FilterManager ...
type FilterManager struct {
	Filter *FilterConfig
}

// IsTarget ...
func (m *FilterManager) IsTarget(path string, info os.FileInfo) bool {
	if info.IsDir() {
		return false
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return false
	}

	outDirs := m.Filter.OutDir
	for _, outDir := range outDirs {
		outDirExp, err := regexp.Compile(outDir)
		if err != nil {
			return false
		}
		if outDirExp.MatchString(absPath) {
			return false
		}
	}

	outFiles := m.Filter.OutFile
	for _, outFile := range outFiles {
		outFileExp, err := regexp.Compile(outFile)
		if err != nil {
			return false
		}
		if outFileExp.MatchString(path) {
			return false
		}
	}

	noMatch := true

	inDirs := m.Filter.InDir
	for _, inDir := range inDirs {
		inDirExp, err := regexp.Compile(inDir)
		if err != nil {
			return false
		}
		if inDirExp.MatchString(absPath) {
			noMatch = false
		}
	}

	inFiles := m.Filter.InFile
	for _, inFile := range inFiles {
		inFileExp, err := regexp.Compile(inFile)
		if err != nil {
			return false
		}
		if inFileExp.MatchString(path) {
			noMatch = false
		}
	}
	if noMatch {
		return false
	}

	return true
}
