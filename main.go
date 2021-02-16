package main

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

var lmlogger = Logger{}

func main() {
	var docPackages []doc.Package

	args := os.Args[1:]
	if len(args) != 1 {
		lmlogger.Debugf("%s", "Invalid Usage: Provide directory to document\n"+
			"\tUsage: ./godoc <project_dir>")
		return
	}

	docPackages, err := populateDocPackages(args)
	if err != nil {
		lmlogger.Errorf("%s", err)
		return
	}

	for _, docPackage := range docPackages {
		lmlogger.Noticef("%s\n%s", docPackage.Name, docPackage.Doc)
	}
}

// populateDocPackages iterates over the provided directory in args
// This directory and its contents are converted into an []doc.Package
func populateDocPackages(args []string) ([]doc.Package, error) {
	// Get Dir from command args
	baseDirStr := args[0]

	var pkgs []ast.Package

	// Parse baseDir
	fset := token.NewFileSet()
	packs, err := parser.ParseDir(fset, baseDirStr, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	// Append baseDir ast.Package to pkgs
	for _, pack := range packs {
		pkgs = append(pkgs, *pack)
	}

	// Get list of files in baseDir
	files, err := os.ReadDir(baseDirStr)
	if err != nil {
		return nil, err
	}

	// Get list of sub-directories (potential go packages)
	var dirs []os.DirEntry
	for _, file := range files {
		if file.IsDir() && !strings.Contains(file.Name(), ".") {
			dirs = append(dirs, file)
		}
	}

	// Convert sub-dirs to ast.Package
	for _, dir := range dirs {
		fset := token.NewFileSet()
		packs, err := parser.ParseDir(fset, baseDirStr+dir.Name(), nil, parser.AllErrors)
		if err != nil {
			return nil, err
		}

		// Append subdir as ast.Package to pkgs
		for _, pack := range packs {
			pkgs = append(pkgs, *pack)
		}
	}

	// For each package in the dir
	var docs []doc.Package
	for _, pkg := range pkgs {
		// Convert to doc.package and append to our list
		pack := doc.New(&pkg, "", doc.PreserveAST)
		docs = append(docs, *pack)
	}
	return docs, nil
}
