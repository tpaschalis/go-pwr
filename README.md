# pwr

Yet another note-taking utility! `pwr` is a small project of mine to ease into the habit of taking structured notes using Markdown, and rendering them into a library-like page for easy reference.

You can find a small example of the final result [here](#).

## Functionality

First of all, you need to fill out `config.json` with the path where you want to store your pages, and your favorite editor. Here's my `config` in a Windows system
```
# config.json
{
	"storagePath" : "C://pwr-notes//",
	"templatesPath" : "C://Go//src//pwr//html_templates//", 
	"editor" : "code",
	"editorArgs" : ""
}
```

All command-line arguments are case-insensitive and pages are created using all-lowercase.

Let's check it out in action!

```
$ pwr 				# When `pwr` is called with no argument, or with argument "today", it will create or open
$ pwr today			# a 'page' with today's date as "YYYY-MM-DD.md", where you can take Markdown notes, 
					# or check your existing ones

$ pwr yesterday 	# Same thing for yesterday and tomorrow, really useful when you
$ pwr tomorrow		# need to remember where you left off, or what's on today's tasklist

$ pwr <subject>		# You can create and/or access a named page by specifying any other word
$ pwr linux			# Here you will create a new "linux.md" page, or access your previous notes
					# under the "linux.md" file

$ pwr -v 			# -v, short for view, will render all of your created 'pages', as well as a structured index, so you can access them from a browser
Access your library by opening the following link
/home/paschalis/pwr-notes/index.html
Rendering complete!

$ pwr --add <sub>	# Create a new, empty page "<sub>.md" if it doesn't exist.
$ pwr --del <sub> 	# Delete any existing "<sub>.md" pages, if they exist 
```


## Prerequisites
`pwr` should work fine in Windows, Linux, MacOS, and most Unix-y systems.

To build from source, you need need Go >=1.10 and the `blackfriday v2` module.

```
go get gopkg.in/russross/blackfriday.v2
git clone https://github.com/tpaschalis/pwr.git
cd pwr
export PWR_PATH_TO_CONFIG=$(pwd)/config.json
go build .
```

If you want to run `pwr` from anywhere in your system, you can either add the exeutable to your `$PATH` or use `go install`, but remember to set the `PWR_PATH_TO_CONFIG` variable to point to pwr's `config.json`.

If you prefer, there will be [releases](https://github.com/tpaschalis/pwr/releases) available, with precompiled binaries for the most popular systems, but for the time being you will *still* need to set `PWR_PATH_TO_CONFIG`.


## Roadmap
- Backup your notes/pages into a database
- Restore your notes/pages from a database

## Disclaimer

This is a work-in-progress project in its nascent stage. I don't know how the project will play out, if I will be the only one to use it, or if it gets even a couple of more people interested. I hope it's okay for you to expect small bumps here and there :) 

Contributions and issues are encouraged. I'd be *really* grateful if you'd reach out for critisicm and advice.
