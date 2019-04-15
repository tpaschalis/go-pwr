package pwrf

import "fmt"
import "gopkg.in/russross/blackfriday.v2"
import "os"
import "io/ioutil"
import "html/template"
import "regexp"


type Page struct {
	Title string
	Body  template.HTML
}

type Index struct {
	Title string
	Calendar []Link
	Misc []Link
}

type Link struct {
	Name string
	Address string
}

func BuildIndex(store, templates string) {
	f, err := os.Open(store)
	check(err)
	fileInfo, err := f.Readdir(-1)
	f.Close()
	check(err)

	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	//fmt.Printf("Pattern: %v\n", re.String())

	var indexData Index
	indexData.Title = "pwr index"
	for _, file := range fileInfo {

		if file.IsDir() && pageExists(store, file.Name()) == nil {
			address := file.Name() + "/" + file.Name() + ".html"
			if re.MatchString(file.Name()) {
				indexData.Calendar = append(indexData.Calendar, Link{file.Name(), address})
			} else {
				indexData.Misc = append(indexData.Misc, Link{file.Name(), address})
			}
			//currentLink := store + file.Name() + "/" + file.Name() + ".html"
			//allNotes = append(allNotes, currentLink)
		}
	}
	//fmt.Println(indexData)
	tmpl := template.Must(template.ParseFiles(templates+"index.html"))

	os.Chdir(store)
	f, err = os.Create("index.html")
	check(err)
	defer f.Close()

	err = tmpl.Execute(f, indexData)
	check(err)
	fmt.Printf("Access your library by opening the following link\n%sindex.html\n", store)
}

func RenderPages(store, templates string) {

	f, err := os.Open(store)
	check(err)
	fileInfo, err := f.Readdir(-1)
	f.Close()
	check(err)

	tmpl := template.Must(template.ParseFiles(templates+"note.html"))

	customRenderer := blackfriday.NewHTMLRenderer(
		blackfriday.HTMLRendererParameters{
		Flags : blackfriday.CommonHTMLFlags|blackfriday.TOC,
		})

	for _, file := range fileInfo {
		if file.IsDir() {
			//fmt.Println(store+file.Name())
			os.Chdir(store + file.Name())
			renderPage(store, file.Name(), tmpl, customRenderer)

		}
	}

	fmt.Println("Rendering complete!")
}

func renderPage(store, filename string, tmpl *template.Template, pageRenderer blackfriday.Renderer) {
	data, err := ioutil.ReadFile(filename + ".md")
	check(err)

	f, err := os.Create(filename + ".html")
	check(err)
	defer f.Close()

	current := Page{Title: filename, Body: template.HTML(blackfriday.Run(data, blackfriday.WithRenderer(pageRenderer)))}
	err = tmpl.Execute(f, current)
	check(err)
}

func pageExists(p, fn string) error {
	os.Chdir(p + fn)
	_, err := os.OpenFile(p+fn+"/"+fn+".md", os.O_RDWR, 0666)
	return err
}
