// Package mwe creates and manipulates MWE for LaTeX
// Also can parse a .tex file and create
// minimum examples from the directories of file name
// can number them automatically.
// Can also replace them with rendered output.
//
package mwe

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

// MWEInterface is an interface to create Minimum Working Examples
// of LaTeX documents
type MWEInterface interface {
	CreateMWE(path string, latexOptions ...string)
}

// MWE is a struct for building Minimum Working Examples.
type MWE struct {
	Type                  string
	FileName              string
	FileNamePrefix        string
	FileNameNumberingAuto bool
	BasePath              string
	Version               string
}

var templ  = `% From Wikibooks
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

// CharacterTable as used in a dtx.
var CharacterTable = `%<*package>
%% \CharacterTable
%%  {Upper-case    \A\B\C\D\E\F\G\H\I\J\K\L\M\N\O\P\Q\R\S\T\U\V\W\X\Y\Z
%%   Lower-case    \a\b\c\d\e\f\g\h\i\j\k\l\m\n\o\p\q\r\s\t\u\v\w\x\y\z
%%   Digits        \0\1\2\3\4\5\6\7\8\9
%%   Exclamation   \!     Double quote  \"     Hash (number) \#
%%   Dollar        \$     Percent       \%     Ampersand     \&
%%   Acute accent  \’     Left paren    \(     Right paren   \)
%%   Asterisk      \*     Plus          \+     Comma         \,
%%   Minus         \-     Point         \.     Solidus       \/
%%   Colon         \:     Semicolon     \;     Less than     \<
%%   Equals        \=     Greater than  \>     Question mark \?
%%   Commercial at \@     Left bracket  \[     Backslash     \\
%%   Right bracket \]     Circumflex    \^     Underscore    \_
%%   Grave accent  \‘     Left brace    \{     Vertical bar  \|
%%   Right brace   \}     Tilde         \~}
%</package>`

// Test provides a string for a MWE.
var Test  = `\begin{document}
 				 \chapter{Test} 	
				 This is the body of the document
				 \begin{figure}
				 some body
				 \end{figure}
                 \end{document}`

// Sty  bprovides a structure for style files.
type Sty struct {
	Type        string
	FileName    string
	PackageName string
	PackageTitle string
	License     string
	Author      string
	Date        string 
}

// Struct for holding a LaTeX .cls file (.dtx are handled separately)
type LaTeXClass struct {
	BaseClass string
	Cls Sty
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

//CreateLaTeXClass assembles a LaTeX .cls file
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

//Example exports an example.
func Example() {
	t := new(MWE)
	t.CreateMWE("mwe1.tex", "code...")
}

// Extract extracts everything between a begin and end document.
func Extract() {
	Num:=1
	//gets anything between body
	//r, err := regexp.Compile("(?s)(\\\\begin{document})(.+)(\\\\end{document})")
	r, err := regexp.Compile("\\\\(.+)")
	if err != nil {
        fmt.Printf("There is a problem with your regexp.\n")
        return
    }
    // Will print 'Match'
    //r.ReplaceAllString("\\bc def ghi", "$2 $1") 
    fmt.Printf("\nMatch: %s\n",  r.ReplaceAllString("\\bc \\def ghi", "SLASH"+ strconv.Itoa(Num) +"$1") )
    fmt.Printf(Test)
    //} else {
      //  fmt.Printf("No match ")
    //..}

}

// Parse provides helper functions for parsing text and replacing with other items
// func Parse() {
//         src := []byte(Test)
//         search := regexp.MustCompile("(?s)\\\\(begin){figure}(.+)\\\\(end{figure})")
//         repl := []byte("$1")

//         i := 0
// 	//n := -10
// 	src = search.ReplaceAllFunc(src, 
// 		func(s []byte) []byte {
// 		i++
// 		tmp:= append(strconv.AppendInt(repl,int64(i),10))
// 		fmt.Println("\nSomething:",string(tmp), i, string(s),"fence")
// 		t2:= string(tmp) 
// 		return []byte(t2)
// 		}
// 	)

		
//     fmt.Println(string(src))
// }

func computedFrom(s string) string {
        return fmt.Sprintf("computedFrom(%s)", s)
}

// Parse2 parses a key value string.
func Parse2(){
        input := `b:foo="hop" b:bar="hu?"`
        r := regexp.MustCompile(`\b.:\w+="([^"]+)"`)
        r2 := regexp.MustCompile(`"([^"]+)"`)
        fmt.Println(r.ReplaceAllStringFunc(input, func(m string) string {
                match := string(r2.Find([]byte(m)))
                fmt.Println(m)
                return r2.ReplaceAllString(m, computedFrom(match))
        }))
}

// Parse3 encloses a figure.
func Parse3(){
        input := `\\begin{figure}
                   Something
                  \\end{figure}`
        r := regexp.MustCompile(`\\\\begin{figure}`)
        r2 := regexp.MustCompile(`(?s)(.+)\\\\end{figure}`)
        fmt.Println(r.ReplaceAllStringFunc(input, func(m string) string {
                match := string(r2.Find([]byte(m)))
                fmt.Println("\n M:",m)
                return r2.ReplaceAllString(m, computedFrom(match))
        }))
}
