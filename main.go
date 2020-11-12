package main

import (
	"github.com/davecgh/go-spew/spew"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fileSet := token.NewFileSet()

	walker := func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
			node, err := parser.ParseFile(fileSet, path, nil, parser.ParseComments)
			if err != nil {
				log.Fatal(err)
			}
			spew.Dump(node)
		}
		return nil
	}
	filepath.Walk("./internal", walker)

}
