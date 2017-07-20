package go_crud

import (
	"fmt"
	"testing"

	"path/filepath"

	"github.com/stretchr/testify/assert"
)

func TestIsTarget(t *testing.T) {
	factors := []struct {
		path     string
		out      OutFilter
		in       InFilter
		expected bool
	}{
		{
			path: "",
			out: OutFilter{
				DirFilter:  "",
				FileFilter: "",
			},
			in: InFilter{
				DirFilter:  "",
				FileFilter: "",
			},
			expected: true,
		},
	}

	for idx, f := range factors {
		fmt.Printf("No.%02d", idx+1)
		assert.Equal(t, f.expected, IsTarget(f.path, f.out, f.in))
	}
}

type Filter interface {
	Match(path string) bool
}

type OutFilter struct {
	DirFilter  string
	FileFilter string
}

func (f *OutFilter) Match(path string) bool {
	return false
}

type InFilter struct {
	DirFilter  string
	FileFilter string
}

func (f *InFilter) Match(path string) bool {
	return true
}

func IsTarget(path string, out OutFilter, in InFilter) bool {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return true
}
