package main

import (
	"flag"
	"fmt"
	// "path/filepath"
  	// "io/ioutil"
  	"time"
	"os"
	// "strings"
	"text/template"
	"github.com/dieuhd/swan/lib/swan"
)

var dbPath = flag.String("path", "config", "folder containing database config")
var env = flag.String("env", "development", "which DB environment to use")

func loading(c chan int) {
	for {
		select {
        case <- c:
            return
        default:
            fmt.Print(".")
        }
        time.Sleep (time.Millisecond * 500)
	}
}
func main () {
	flag.Usage = usage
	flag.Parse()
	
	// var swanConfig = swan.GetSwanConfig()
	// dbConfig, err := swan.GetDb(swanConfig)
	// if(dbConfig == nil) {
	// 	fmt.Println("db config invalid")
	// 	fmt.Println(err)
	// }

	args := flag.Args()
	// fmt.Println(args)
	if len(args) == 0 || args[0] == "-h" {
		flag.Usage()
		return
	}

	firstParam := flag.Arg(0)
	if firstParam == "new" {
		if len(args) > 2 {
			fmt.Println("invalid param");
			return
		} else if(len(args) == 1) {
			fmt.Println("Please enter project name");
			return
		}
		projectName := flag.Arg(1)
		fmt.Printf("Creating project golang with name `%s`\n", projectName)
		
		cwd, _ := os.Getwd()
		existed, _ := swan.Exists(cwd + "/" + projectName)
		if existed == true {
			fmt.Println("Project existed")
			var answer string
			for {
				fmt.Print("The project name existed. Do you want remove and create project(yes/no)?:")
				fmt.Scanf("%s", &answer)
				if answer == "no" {
					return
				} else if answer == "yes" {
					//remove this folder
					os.RemoveAll(projectName)
					break
				}
			}
		}

		quit := make(chan int)
		go loading(quit)
		swan.CreateProject(projectName)		
		time.Sleep (time.Millisecond * 4000)
	}
}

func usage() {
	fmt.Print(usagePrefix)
	flag.PrintDefaults()
}

var usagePrefix = `
swan is a crud management system for Go projects.

Usage:
    swan [options] <subcommand> [subcommand options]

Options:
`
var usageTmpl = template.Must(template.New("usage").Parse(
	`
Commands:{{range .}}
    {{.Name | printf "%-10s"}} {{.Summary}}{{end}}
`))