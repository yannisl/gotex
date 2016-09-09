// Tables and Tabulars
// Uses the pgf packages
// to render nice charts in various formats.
// Also uses tables for data, in various formats
package tabular

import (
	"fmt"
	"io/ioutil"

)

const (
	space        = " "
	beginTabular = "\\begin{tabular}"
	endTabular   = "\\end{tabular}"
	beginTable   = "\\begin{table}"
	endTable     = "\\end{table}"
)

// Insert is an interface of inserting a set of items in
// a TeX stream
type Insert interface{
	Insert(row ...interface{}) 
}

// MWE is an interface to create Minimum Working Examples
// of LaTeX documents
type MWE interface{
	MWE(latexOptions ...string)
}

type Style interface{}
type Data [][]string
type Row []string

//Type Caption holds fields for defining a caption
type Caption struct {
	Text          string
	Specification string
	Index		  bool 
	ShortIndex    string       
}

type Table struct {
	Specification        string
	DefaultSpecification string
}

// Tabular defines a structure for tabulated data
//
type Tabular struct {
	BeginCmd          string
	Header            string
	HasRules          bool
	HasHead           bool
	Cells             string
	Rows              interface{}
	Specification     string
	SpecificationList []string //similar to Specification but as a list
	// is more readable and can be iterated quicker
	Data    Data
	Get     interface{}
	Style   interface{}
	EndCmd  string
	Caption Caption
	Type    string
	Rendered string //stores the tex string
}

// Insert implements the Insert interface
func (t *Tabular) Insert(element ...interface{}) {
	var s []string
	for _, v := range element {
		 s = append(s, v.(string))
		 //fmt.Println(element[i].(string))
	}
	t.Data = append(t.Data, s)
}

func (t *Tabular) checkSpecificationList() bool {
	if t.SpecificationList != nil {
		//fmt.Println("Not Empty")
		return false
	} else {
		//fmt.Println("There is a spec")
		return true
	}
}

// ToTeX translates the contents of tabular data to a
// LaTeX longtable, tabular or other tabular like environment
// and returns the contents.
func (t *Tabular) ToTeX() bool {
	
	s := "\n" + beginTable + "[htbp]\n"
	//s = "\\centering"

	if t.checkSpecificationList()  {
		s += beginTabular + "{" + t.Specification + "}" + "\n"
	} else {
		s += beginTabular + "{" + t.Specification + "}" + "\n" //todo automate and center all
	}
	if t.HasRules {
		s += "\\toprule\n"
	}
	for i, v := range t.Data {
		//fmt.Print(v[0] + space)
		if t.HasHead && i == 1 {
			s += "\\midrule\n"
		}
		for j, val := range v {
			if j > 0 {
				s += "&" + val + space
			} else {
				s += val + space
			}
		}
		s += "\\\\\n"

	}
	if t.HasRules {
		s += "\\bottomrule\n"
	}
	s += endTabular + "\n"
	s += "\\caption{" + t.Caption.Text + "}" + "\n" 
	s += endTable + "\n"
	fmt.Println(s)
	// Save the string as well
	t.Rendered = s 
	return true
}

// Caption returns the rendered LaTeX caption. 
func (c *Caption) Caption (s string, idx bool) string {
	if idx {
		s = "\\caption{" +s + "}\n"
		} else {
	    	s = "\\caption*{" + s + "}\n"
		}
	return s
}

//MWE wraps a tabular environment, as a minimum working
//LaTeX example 
func (t *Tabular) CreateMWE(latexOptions ...string) {
   s:="\\documentclass{article}\n" + "\\usepackage{booktabs}\n" +
    "\\begin{document}\n"
    s +=  t.Rendered 
    s += "\\end{document}\n"
    d := []byte(s)
    _ = ioutil.WriteFile("mwe.tex", d, 0644)
    //check(err)
    //defer f.Close()
    fmt.Print(string(d))  
}