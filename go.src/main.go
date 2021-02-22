// This is a main package comment!
package main

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"./web"
)

// Package-specific Logger
var lmlogger = Logger{}

// HTTP Web Server Location
const URL = "http://localhost:6969/"

type htmlTemplateData struct {
	RootDir          string
	DocumentPackages []doc.Package
}

// main
func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		lmlogger.Debugf("%s", "Invalid Usage: Provide directory to document\n"+
			"\tUsage: ./godoc <project_dir>")
		return
	}

	// Get Dir from command args
	baseDirStr := args[0]

	// Find all directories containing Go Files
	goPackages := findGoPackages(baseDirStr)

	// Get []doc.Packages from our Go packages
	var documentPackages []doc.Package
	documentPackages, err := createDocPackages(goPackages, baseDirStr)
	if err != nil {
		lmlogger.Errorf("%s", err)
		return
	}

	// Load our html templates
	err = web.InitializeTemplates()
	if err != nil {
		lmlogger.Errorf("%s", err)
		return
	}

	//TODO: Move HandleFunc code to web/routes.go

	// Load HTML Template Data
	templateData := htmlTemplateData{
		RootDir:          baseDirStr,
		DocumentPackages: documentPackages,
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, URL+"home", http.StatusPermanentRedirect)
	})

	// Handler for serving indexTemplate on request to root path (Ex. http://localhost:6969/)
	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		err = web.IndexTemplate.Execute(writer, templateData)
		if err != nil {
			lmlogger.Errorf("%s", err)
		}
	})

	http.HandleFunc("/packages", func(writer http.ResponseWriter, request *http.Request) {
		err = web.PackagesTemplate.Execute(writer, templateData)
		if err != nil {
			lmlogger.Errorf("%s", err)
		}
	})

	http.HandleFunc("/files", func(writer http.ResponseWriter, request *http.Request) {
		err = web.FilesTemplate.Execute(writer, templateData)
		if err != nil {
			lmlogger.Errorf("%s", err)
		}
	})

	// Handle requests to CSS dir by HTML pages
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./www/css/"))))

	// Open User's Web Browser and direct to URL
	//command := exec.Command("xdg-open", URL)
	//_ = command.Start()

	// Start the HTTP Web Server
	_ = http.ListenAndServe(":6969", nil)
}

// findGoPackages recursively searches baseDir directory and its subdirectories for Golang
// Packages and returns what it finds
func findGoPackages(baseDir string) []string {
	var goPackages []string

	// Get list of files in baseDir
	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return nil
	}

	// Get dirs and files
	var dirs []os.FileInfo
	var fles []os.FileInfo
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file)
		}
		fles = append(fles, file)
	}

	// Search Dirs
	for _, dir := range dirs {
		// Recursive search sub-dirs
		packages := findGoPackages(baseDir + dir.Name() + "/")

		// Append our findings
		goPackages = append(goPackages, packages...)
	}

	// Search Files
	for _, fle := range fles {
		if strings.Contains(fle.Name(), ".go") {
			// If we have .go files, we know we're in a go package
			goPackages = append(goPackages, baseDir)
			break
		}
	}

	return goPackages
}

// createDocPackages takes a list of Go Package file paths and converts them to []doc.Package
func createDocPackages(packages []string, baseDir string) ([]doc.Package, error) {
	// Convert packages to ast.Package
	var pkgs []ast.Package
	for _, dir := range packages {
		fset := token.NewFileSet()
		packs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}

		for _, pack := range packs {
			// Pass Import Dir via pack.Name var
			pack.Name += "|" + dir[len(baseDir)-1:]
			pkgs = append(pkgs, *pack)
		}
	}

	// Convert []ast.Package to []doc.package and return to user
	var docs []doc.Package
	for _, pkg := range pkgs {
		var packName string
		var packPath string

		tokens := strings.Split(pkg.Name, "|")
		packName = tokens[0]
		packPath = tokens[1]

		// Get ImportDir out of pkg.Name and reset pkg.Name
		pkg.Name = packName

		// get doc.Package
		pack := doc.New(&pkg, "", doc.PreserveAST)

		// Update Import Path
		pack.ImportPath = packPath

		docs = append(docs, *pack)
	}
	return docs, nil
}
