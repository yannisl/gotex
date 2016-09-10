// Tables
// Uses the pgf packages
// to render nice charts in various formats.
// Also uses tables for data, in various formats
package main

import (
	//"fmt"
	"fmt"
	"gotex/MWE"
	"gotex/enviro"
	"gotex/tabular"
)

func main() {
	t := new(tabular.Tabular)
	t.HasRules = true
	t.HasHead = true
	t.Specification = "l l r r p{3.5cm}"
	t.SpecificationList = []string{"l", "l", "r", "r", "r"}
	t.Insert("Qty", "Apr", "May", "Jun", "Jul")
	t.Insert("Qty", "100", "200", "300.2", "-256.00")
	t.Insert("Usage", "100", "200", "300.2", "-256.00")
	t.Insert("Usage", "100", "200", "300.2", "-256.00")
	t.Insert("Usage", "100", "200", "300.2", "-256.00\\footnotemark\\footnotetext{This is a footnote}")
	t.Caption.Text = "This is the caption of the table."
	t.Caption.Index = false
	t.ToTeX()
	t1 := new(mwe.MWE)

	// Wrap the table body in a MWE
	t1.CreateMWE("MWE1.tex", t.Rendered)
	// TODO better method for data definition maybe from a plain string which get split?
	//  Qty Apr May Jun Jul
	//  100 200 300 400 500

	z := enviro.NewEnvironment()
	y := z.Wrap("enumerate", "\\item This is first item")
	fmt.Println(y)

	tex := new(enviro.TeXEnviro)
	y1 := tex.Wrap("enumerate", "\\item This is first item\n")
	fmt.Println(y1)

	t3 := new(mwe.Sty)
	t3.CreateLaTeXStyle("test.sty", "")

	mwe.Parse3()
}
