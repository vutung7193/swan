package main

import (
	"flag"
	"fmt"
	// "path/filepath"
  	// "io/ioutil"
	// "os"
	// "strings"
	// "text/template"
	"github.com/dieuhd/swan/lib/swan"
)

var dbPath = flag.String("path", "config", "folder containing database config")
var env = flag.String("env", "development", "which DB environment to use")

func main () {
	
	var swanConfig = swan.GetSwanConfig()
	dbConfig, err := swan.GetDb(swanConfig)
	if(dbConfig == nil) {
		fmt.Println("db config invalid")
		fmt.Println(err)
	}
}