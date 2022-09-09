package generator

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func handlePanicError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GeneratePackage a new folder as a golang module with the given
// package name
// packName is the default package name
// moduleName is the full name of the module e.g.
// github.com/dropdevrahul/gocache
func GenerateModule(packName string, moduleName string) {
	if _, err := os.Stat(packName); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(packName, os.ModePerm)
		handlePanicError(err)
	}

	os.Chdir(packName)

	if _, err := os.Stat(packName); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(packName, os.ModePerm)
		handlePanicError(err)
	}
	f, err := os.Create(filepath.Join(packName, packName+".go"))
	handlePanicError(err)

	defer f.Close()
	f.WriteString(fmt.Sprintf("package %s", packName))

	// create main.go file
	fm, err := os.Create("main.go")
	handlePanicError(err)
	defer fm.Close()
	fm.WriteString("package main")
	fm.WriteString("\n")
	fm.WriteString("import (")
	fm.WriteString(fmt.Sprintf("  \"%s/%s\"", moduleName, packName))
	fm.WriteString(")")
	fm.WriteString("\n")
	fm.WriteString("func main() {")
	fm.WriteString("\n")
	fm.WriteString("}")

	cmd := exec.Command("go", "mod", "init", moduleName)
	_, err = cmd.Output()

	handlePanicError(err)
}
