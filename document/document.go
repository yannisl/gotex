// document
// A document can be any paper or electronic document
// transmitted

// This package was an attempt for me to use text/template to make
// writing LaTeX code easier.
//
// It is also IDeal for reports and repetitive LaTeX tasks.
// Nothing that is here cannot be programmed in TeX itself, well
// almost.
//
package document

import (
	"bytes"
	"fmt"
	"github.com/mcuadros/go-defaults"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"
	"github.com/spf13/afero"
)

// A constant with some text
const body = `
	This is a test body for the template. It will be inserted at the point the variable
	body is inserted. 
`

type paths struct {
	defaultSectionsDir string
	defaultMainFile    string
	projectRoot        string
}

// non standard settings: these are to be obtaned from TOML files
type latexSettings struct {
	useHyperRef  bool `default:true`
	colorLinks   string
	author       string
	makeStyle    bool
	useListings  bool
	useTcolorbox bool
	}

type hyper struct {
	breakLinks   bool  `default:true`
	bookmarksOpen bool `default:true`
	bookmarks     bool `default:true`
}


// DocI defines the interface for all type of documents.
// It provides methods to render a document through templates
// and to store it either locally or on the cloud.
type DocI interface {
	Options()
	Wrap()
	Compile()
	Render()
	Upload()
	Mail() 
	Paths() 
	Show()
}

// Doc describes the structure of a publication. A publication can be anything from
// a letter, a poem, a report, an article, proceedings or even an encyclopedia. Anything that
// can be printed as a pdf. It can also be a website!
type Doc struct {
	//scriptContext   *script.Context
	DocumentID      string `default:"DocumentID"`
	DocumentType    string `default:"DocumentType"`
	Bucket          string
	tags            []string
	meta            map[string]string
	UserID          string // user capturing the record
	Title           string `default:"You need a Title"` // don't leave spaces after var
	Body            string `default:"\\lipsum"`
	HTMLTemplate    string
	LatexTemplate   string `default:"book.tex"`
	Language        string `default:"EN"`
	Cover           interface{}
	DateCreated     string
	DateLastRevised string
	Author          string `default:"Dr. Y. Lazarides"`
	Maintainer      string
	Settings        latexSettings
}

func Show (c *Doc)(){

}

// Render the document through a set of templates
// document.Render()
//http://stackoverflow.com/questions/11467731/is-it-possible-to-have-nested-templates-in-go-using-the-standard-library-googl/11468132#11468132
//https://elithrar.github.io/article/approximating-html-template-inheritance/
func Render(c *Doc) {

	fmt.Println(c.LatexTemplate)

	files, _ := filepath.Glob("templates/*.tex")
	fmt.Println(files)

	t, err := template.New(c.LatexTemplate).Delims("[[", "]]").ParseFiles("templates/" + c.LatexTemplate,
		"templates/partials/header.tex", "templates/partials/base.tex", "sections/main.tex") //templates/doc
	checkError(err)

	// get a file handle io.Writer object
	fileHandle, _ := os.Create("templates/output.tex")

	// Template parameters are in doc
	err = t.Execute(fileHandle, c)
	defer fileHandle.Close()

	// write to console also, Writer can do anything
	// we can write to an S3 bucket if we want or a socket
	err = t.Execute(os.Stdout, c)

	checkError(err)
}

// Domostration of the init function, it always runs on its own
// This got me at first...
//
func init() {
	var buf bytes.Buffer

	//setup logging
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Println("Hi")
	fmt.Print(&buf)
	// Open the file to write
	f, err := os.OpenFile("logger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	// set the logger output
	log.SetOutput(f)
	// print something to see we are alive
	log.Println("PHD logger running...!")
	log.Println("We are live...")
	// get and compile the template here
}

// SetDefaults function provides a ...
func SetDefaults(doc *Doc) {
	defaults.SetDefaults(doc) //changes only if empty
}



// New Creates a new Document type in memory
func New(latexClass string) *Doc {
	p := new(Doc)
	SetDefaults(p)
	p.LatexTemplate = latexClass + ".tex"
	return p
}

// PartialLatexFiles is a function that can get a number of files
// and include them in the templates or process them indivIDually
func PartialLatexFiles(path string) {
	//read all files in directory sections
	// this must be settleable and discoverable
	var counter int

	type Content struct {
		fileName string
		Contents string
	}

	var contentsList map[string]string
	contentsList = make(map[string]string)

	List := new(Content)

	//append(s []T, x ...T)
	files, _ := ioutil.ReadDir("./sections")
	for _, f := range files {
		if f.Name() == "main.tex" {
			fmt.Println("We found a main file")
		}
		fmt.Println(f.Name())
		S1, err := ioutil.ReadFile("./sections/" + f.Name())
		if err != nil {
			fmt.Println(err)
		}

		List.fileName = f.Name()
		List.Contents = string(S1)
		fmt.Println(string(S1))

		contentsList[f.Name()] = string(S1)
		counter++

	}
	fmt.Println("TEST", contentsList["main.tex"])
	//inFile, _ := ioutil.ReadFile(path)
	//fmt.Println("CONCATENATION:", Y.contents)
	//fmt.Printf("Found %v files", counter)
}

// TestWriteFile Just a test to see if I can use the Afero FS
// Maybe better to go to minio or both. Minio uses a db
// and maybe best to avoID it, although don't see why it could
// be used with S3
func TestWriteFile() {

	testFS:= new(afero.OsFs)
	//testFS = &fs.MemMapFs{}
	fsutil := &afero.Afero{Fs: testFS}

	f, err := fsutil.TempFile("temp/", "ioutil-test")
	if err != nil {
		fmt.Println(err)
	}
	filename := f.Name()
	data := "Programming today is a race between software engineers striving to " +
		"build bigger and better IDiot-proof programs, and the Universe trying " +
		"to produce bigger and better IDiots. So far, the Universe is winning."

	if err := fsutil.WriteFile(filename, []byte(data), 0644); err != nil {
		fmt.Println("WriteFile %s: %v", filename, err)
	}

	contents, err := fsutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile %s: %v", filename, err)
	}
	fmt.Println(string(contents))
	if string(contents) != data {
		log.Fatalf("contents = %q\nexpected = %q", string(contents), data)
	}

	// cleanup
	f.Close()
	//testFS.Remove(filename) // ignore error
}
// Program starts here
func Run() {
	//TestWriteFile()
	//a, _ := AppFs.Open("report.bat")
	//fs:= new(afero.OsFs)
	//err := fs.Mkdir("temp1", 777)
	//g, err:=fs.IsEmpty("temp/ioutil-test")
	//fmt.Println(fs)
	//fmt.Println(afero.DirExists(fs, "temp1"))

	//return
	doc := New("book")
	//doc1 := New("article")
	//Example(doc)
	today := today()
	doc.DateLastRevised = today

	Render(doc)
	//Render(doc1)
	// initialize the paths for sections directories
	paths := paths{"./sections", "main.tex", "C:/Projects/Go/src/gotex/"}

	PartialLatexFiles(paths.defaultSectionsDir + paths.defaultMainFile)

	// latex the files
	cmd := exec.Command("lualatex.exe", "output.tex", "--interaction=nonstopmode")
	// sets the directory we are operating in
	//
	cmd.Dir = "c:/Projects/Go/src/gotex/templates"
	out, err := cmd.CombinedOutput()
	fmt.Printf("OUT = %v\n", string(out))
	checkError(err)
	fmt.Println(today)
}

// utility to set the time to now at the format 
//we want
func today() string {
	// layout shows by example how the reference time should be represented.
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	return time.Now().Format("Jul 2, 2006")
}

// Checks for errors
func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
}
