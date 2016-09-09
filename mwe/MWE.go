// Package mwe
// creates and manipulates MWE for LaTeX
// Also can parse a .tex file and create
// minimum examples from the directories of file name
// can number them automatically.
// Can also replace them with rendered output.
//
package mwe

import (
	"fmt"
	"io/ioutil"
)

// MWE is an interface to create Minimum Working Examples
// of LaTeX documents
type MWEInterface interface {
	CreateMWE(path string, latexOptions ...string)
}

type MWE struct {
	Type                  string
	FileName              string
	FileNamePrefix        string
	FileNameNumberingAuto bool
	BasePath              string
	Version               string
}

var templ string = `% From Wikibooks
%% https://en.wikibooks.org/wiki/LaTeX/Creating_Packages
\NeedsTeXFormat{LaTeX2e}[1994/06/01]
\ProvidesPackage{custom}[2013/01/13 Custom Package]
\RequirePackage{lmodern}
%% 'sans serif' option
\DeclareOption{sans}{
  \renewcommand{\familydefault}{\sfdefault}
}
%% 'roman' option
\DeclareOption{roman}{
  \renewcommand{\familydefault}{\rmdefault}
}
%% Global indentation option
\newif\if@neverindent\@neverindentfalse
\DeclareOption{neverindent}{
  \@neverindenttrue
}
\ExecuteOptions{roman}
\ProcessOptions\relax
%% Traditional LaTeX or TeX follows...
\endinput`

type Sty struct {
	Type        string
	FileName    string
	PackageName string
	License     string
	Author      string
}

// Struct for holding a LaTeX .cls file (.dtx are handled separately)
type LaTeXClass struct {
	Type        string
	FileName    string
	PackageName string
	License     string
	Author      string
}

//MWE wraps a tabular environment, as a minimum working
//LaTeX example
func (t *MWE) CreateMWE(path string, body string, latexOptions ...string) {
	s := "\\documentclass{article}\n" + "\\usepackage{booktabs}\n" +
		"\\begin{document}\n"
	s += "A document"
	s += "\\end{document}\n"
	d := []byte(s)
	_ = ioutil.WriteFile(path, d, 0644) //change to writer
	//defer f.Close()
	fmt.Print(string(d))
}

//CreateLaTeXStyle creates a package from a template file
//Inserts basic information at this stage. In later versions
//we will extract a preamble, create a Package file, install
//modify the original file and run
func (t *Sty) CreateLaTeXStyle(path string, body string, latexOptions ...string) {
	s := templ
	d := []byte(s)
	_ = ioutil.WriteFile(path, d, 0644) //change to writer
	//defer f.Close()
	fmt.Print(string(d))
}

//CreateLaTeXStyle wraps a tabular environment, as a minimum working
//LaTeX example
func (t *LaTeXClass) CreateLaTeXClass(path string, body string, latexOptions ...string) {
	s := "\\documentclass{article}\n" + "\\usepackage{booktabs}\n" +
		"\\begin{document}\n"
	s += "A document"
	s += "\\end{document}\n"
	d := []byte(s)
	_ = ioutil.WriteFile(path, d, 0644) //change to writer
	//defer f.Close()
	fmt.Print(string(d))
}

//Usage Example
func Example() {
	t := new(MWE)
	t.CreateMWE("mwe1.tex", "code...")
}
