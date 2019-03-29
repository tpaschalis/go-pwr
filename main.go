package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

type config struct {
	storagePath string
	editor      string
}

func main() {
	setupFlags(flag.CommandLine)
	confFlag := flag.String("c", "", "-c					Change path for config.json")

	flag.Parse()
	args := flag.Args()

	confFn := os.Getenv("PWR_PATH_TO_CONFIG")
	if confFn == "" {
		confFn = "config.json"
	}
	byt, err := ioutil.ReadFile(confFn)
	check(err)

	var conf map[string]interface{}
	err = json.Unmarshal(byt, &conf)
	check(err)
	storagePath := conf["storagePath"].(string)
	editor := conf["editor"].(string)
	editorArgs := conf["editorArgs"].(string)
	fmt.Println(editor, editorArgs)

	//err = checkEarlyExit(args, *confFlag)
	//check(err)

	if len(args) == 0 || strings.ToLower(args[0]) == "today" {
		openTodayNote(storagePath, editor, editorArgs)
	}

	if len(args) != 0 && strings.ToLower(args[0]) != "today" {
		openNamedNote(storagePath, strings.ToLower(args[0]), editor, editorArgs)
	}

	_ = confFlag
	fmt.Println("")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func checkFileExists(p string, n string) string {
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

func openTodayNote(store, ed, edA string) {
	now := time.Now()
	dateStr := now.Format("2006-01-02")
	filepath := checkFileExists(store, dateStr)
	cmd := exec.Command(ed, edA, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err)
}

func openNamedNote(store, name, ed, edA string) {
	// Not implemented yet
	filepath := checkFileExists(store, name)
	cmd := exec.Command(ed, edA, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err)
}

func setupFlags(f *flag.FlagSet) {
	f.Usage = func() {
		fmt.Println("\nSome Message")

		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		f.PrintDefaults()

		fmt.Println("\nExaplanatory stuff")
	}
}

func checkEarlyExit(args []string, c string) (err error) {
	if len(args) > 1 {
		fmt.Println("Couldn't understand what you're trying to do")
		os.Exit(1)
		flag.Usage()
	}
	//validArg := strings.ToLower(args[0])
	//switch validArg {
	//case
	//    "today",
	//    "tomorrow",
	//    "yesterday",
	//    "todo":
	//    return nil
	//}
	return fmt.Errorf("No such argument is available")
}
