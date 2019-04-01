package pwrf

import "fmt"
import "gopkg.in/russross/blackfriday.v2"
import "os"
import "io/ioutil"

type note struct {
	title string
	body string
}
func BuildIndex(store string) {
	f, err := os.Open(store)
	check(err)
	fileInfo, err := f.Readdir(-1)
	f.Close()
	check(err)

	var allNotes []string
	for _, file := range fileInfo {
		if file.IsDir() && noteExists(store, file.Name()) == nil {
			currentLink := store + file.Name() + "/" + file.Name() + ".html"
			allNotes = append(allNotes, currentLink)
		}
	}
	fmt.Println(allNotes)
}

func RenderNotes(store string) {

	f, err := os.Open(store)
	check(err)
	fileInfo, err := f.Readdir(-1)
	f.Close()
	check(err)

	for _, file := range fileInfo {
		if file.IsDir() {
			//fmt.Println(store+file.Name())
			os.Chdir(store+file.Name())

			data, err := ioutil.ReadFile(file.Name()+".md")
			check(err)
			html := blackfriday.Run(data)

			f, err := os.Create(file.Name() + ".html")
			check(err)
			defer f.Close()
			_, err = f.Write(html)
			check(err)
		}
	}

	fmt.Println("Exiting renderNotes()")
}

func sample() {

	input := []byte(`# Title
## Sub title
### Sub sub title
This is some text *with italics* and **bolds**
`)

	html := string(blackfriday.Run(input))
	fmt.Println(html)
}

func noteExists(p, fn string) error {
	os.Chdir(p+fn)
	_, err := os.OpenFile(p+fn+"/"+fn+".md", os.O_RDWR, 0666)
	return err
}
