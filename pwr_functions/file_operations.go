package pwrf

import "os"
import "os/exec"
import "time"

func CreateOpenFile(p string, n string) string {
	// todo : Implement checking p as filesystem path using `os`
	// Also, check out https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go/22483001#22483001
	foldername := p + n + "/"
	filename := foldername + n + ".md"
	err := os.MkdirAll(foldername, 0700)
	check(err)
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0600)
	check(err)
	defer f.Close()

	fi, err := f.Stat()
	check(err)
	if fi.Size() == 0 {
		_, err = f.WriteString("# " + n + "\n\n\n")
		check(err)
	}
	f.Sync()
	return filename
}

func OpenTodayPage(store, ed, edA string) {
	now := time.Now()
	dateStr := now.Format("2006-01-02")
	filepath := CreateOpenFile(store, dateStr)
	cmd := exec.Command(ed, edA, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err)
}

func OpenYesterdayPage(store, ed, edA string) {
	yesterday := time.Now().AddDate(0, 0, -1)
	dateStr := yesterday.Format("2006-01-02")
	filepath := CreateOpenFile(store, dateStr)
	cmd := exec.Command(ed, edA, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err)
}


func OpenTomorrowPage(store, ed, edA string) {
	tomorrow := time.Now().AddDate(0, 0, +1)
	dateStr := tomorrow.Format("2006-01-02")
	filepath := CreateOpenFile(store, dateStr)
	cmd := exec.Command(ed, edA, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err)
}

func OpenNamedPage(store, name, ed, edA string) {
	filepath := CreateOpenFile(store, name)
	cmd := exec.Command(ed, edA, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err)
}

func CreateEmptyPage(store, name string) {
	_ = CreateOpenFile(store, name)
}

func DeleteNamedPage(store, name string) {

	var pageName string
	switch (name) {
		case "today" :
			today:= time.Now()
			pageName = today.Format("2006-01-02")
		case "yesterday":
			yesterday := time.Now().AddDate(0, 0, -1)
			pageName = yesterday.Format("2006-01-02")
		case "tomorrow":
			tomorrow := time.Now().AddDate(0, 0, +1)
			pageName = tomorrow.Format("2006-01-02")
		default:
			pageName = name
	}
	if _, err := os.Stat(store+"/"+pageName); !os.IsNotExist(err) {
		os.RemoveAll(store+pageName)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
