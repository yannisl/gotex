// Tables and Tabulars
// Uses the pgf packages
// to render nice charts in various formats.
// Also uses tables for data, in various formats
package enviro

import (
//"fmt"
//"io/ioutil"

)

const (
	space       = " "
	beginEnviro = "\\begin{"
	endEnviro   = "\\end{"
)

// CodeWrap provides an interface for wrapping
// code within a block such as \begin{} code \end{}
type CodeWrap interface {
	Wrap(env string, code string) string
	Name() bool
}

// LateXEnviro holds the commands for a LaTeX environment
type LateXEnviro struct {
	ID                  string
	HasOptionalArgument bool
}

type TeXEnviro struct {
	ID string
}

type HTMLEnviro struct {
	ID string
}

// Wrap implements the Enviro interface
func (e LateXEnviro) Wrap(env, code string) string {
	return beginEnviro + env + "}" + code + endEnviro + env + "}"
}

// Name implements sets the Name of the Environment
func (e LateXEnviro) Name(name string) bool {
	e.ID = name
	return true
}

// NewEnvironment creates a new LaTeXEnvro object
func NewEnvironment() *LateXEnviro {
	e := new(LateXEnviro)
	return e
}

// NewEnvironment creates a new LaTeXEnvro object
func NewTeXEnvironment() *TeXEnviro {
	e := new(TeXEnviro)
	return e
}

// Implement TeX Environment
// Wrap implements the Enviro interface
func (e TeXEnviro) Wrap(env, code string) string {
	return "\\" + env + "\n" + code + "\\end" + env
}

// Name implements the CodeWrap interface and sets the Name of the Environment
func (e TeXEnviro) Name(name string) bool {
	e.ID = name
	return true
}

// Implement TeX Environment
// Wrap implements the Enviro interface
func (e HTMLEnviro) Wrap(env, code string) string {
	return "<" + env + ">" + code + "</" + env + ">"
}

// Name implements the CodeWrap interface and sets the Name of the Environment
func (e HTMLEnviro) Name(name string) bool {
	e.ID = name
	return true
}
