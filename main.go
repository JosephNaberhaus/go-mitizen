package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/JosephNaberhaus/go-mitizen/commit"
	"github.com/JosephNaberhaus/go-mitizen/git"
	"io/ioutil"
	"log"
	"os"
)

var installFlag = flag.Bool("install", false, "install this executable as a git subcommand runnable with \"git cz\"")
var logFlag = flag.Bool("log", false, "write program logs to \"logs.txt\" in the working directory")
var dryRun = flag.Bool("dry", false, "print the commit message without performing the commit")

func Usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "\nUsage: %s [-log] [-install]\n", os.Args[0])
	flag.PrintDefaults()
}

func main()  {
	flag.Usage = Usage
	flag.Parse()

	if *logFlag {
		f, err := os.OpenFile("logs.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
	} else {
		log.SetOutput(ioutil.Discard)
	}

	if *installFlag {
		err := git.InstallAsSubcommand("cz")
		if err != nil {
			if errors.Is(err, os.ErrPermission) {
				println("Install failed: permission denied. Try running with sudo")
			} else {
				println("Install failed")
			}

			log.Fatal(err)
		}

		println("Installed subcommand. Run with \"git cz\"")
		return
	}

	err := commit.Commit(*dryRun)
	if err != nil {
		println("No commit made")
		log.Fatal(err)
	}
}