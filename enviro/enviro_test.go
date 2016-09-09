// Tables and Tabulars
// Uses the pgf packages
// to render nice charts in various formats.
// Also uses tables for data, in various formats
package enviro

import (
	"testing"
)

var z = NewEnvironment()

func TestWrap(t *testing.T) {
	expected := `\begin{enumerate}test\end{enumerate}`
	actual := z.Wrap("enumerate", "test")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestWrapTeX(t *testing.T) {
	z := NewTeXEnvironment()
	expected := `\enumerate` + "\n" + `test\endenumerate`
	actual := z.Wrap("enumerate", "test")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestWrapDoc(t *testing.T) {
	z1 := NewEnvironment()
	expected := `\begin{document}` + `test\end{document}`
	actual := z1.Wrap("document", "test")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

// This is mostly to test my understanding of interfaces
func TestWrapHTML(t *testing.T) {
	z1 := new(HTMLEnviro)
	expected := `<h3>` + `test</h3>`
	actual := z1.Wrap("h3", "test")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
