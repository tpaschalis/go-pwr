package pwrf

import "os"
import "os/exec"
import "time"

func CheckFileExists(p string, n string) string {
	// Implement checking p as filesystem path using `os`
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

func OpenTodayNote(store, ed, edA string) {
	now := time.Now()
	dateStr := now.Format("2006-01-02")
	filepath := CheckFileExists(store, dateStr)
	cmd := exec.Command(ed, edA, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err)
}

func OpenYesterdayNote(store, ed, edA string) {
	yesterday := time.Now().AddDate(0, 0, -1)
	dateStr := yesterday.Format("2006-01-02")
	filepath := CheckFileExists(store, dateStr)
	cmd := exec.Command(ed, edA, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err)
}

func OpenNamedNote(store, name, ed, edA string) {
	// Not implemented yet
	filepath := CheckFileExists(store, name)
	cmd := exec.Command(ed, edA, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
