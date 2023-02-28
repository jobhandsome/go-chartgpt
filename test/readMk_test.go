package test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	getwd, _ := os.Executable()
	fmt.Println(getwd)
	fmt.Println(filepath.Dir(filepath.Dir(getwd)))
}
