package main

import (
	"os"

	"github.com/dropdevrahul/go-module-generator/generator"
)

func main() {
	// arg[0] skip program name
	// arg[1] target_dir
	// arg[2] full package name
	// arg[3] module name
	generator.GenerateModule(os.Args[1], os.Args[2], os.Args[3])
}
