package swan

import (
	// "fmt"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)
type Database struct {
	Default DBConfig    	`yaml:"default"`
	Production DBConfig    	`yaml:"production"`
	Development DBConfig    `yaml:"development"`
	Test DBConfig    		`yaml:"test"`
}
func CreateProject (projectName string) {
	//create project name
	os.Mkdir(projectName, 0777)
	// create app folder
	os.Mkdir(projectName + "/app", 0777)
	os.Mkdir(projectName + "/app/controllers", 0777)
	os.Mkdir(projectName + "/app/models", 0777)
	os.Mkdir(projectName + "/app/views", 0777)
	os.Mkdir(projectName + "/app/middleware", 0777)
	//create config
	os.Mkdir(projectName + "/config", 0777)
	//create database config
	yamlDB := Database{
		Default: DBConfig{
			Host: "localhost",
			Username: "root",
			Password: "root",
			Adapter: "mysql",
			Dbname: projectName,
			Encoding: "utf8",
			Port: "3306",
		},
		Production: DBConfig{
			Host: "localhost",
			Username: "root",
			Password: "root",
			Dbname: projectName,
		},
		Development: DBConfig{
			Host: "localhost",
			Username: "root",
			Password: "root",
			Dbname: projectName,
		},
		Test: DBConfig{
			Host: "localhost",
			Username: "root",
			Password: "root",
			Dbname: projectName,
		},
	}
	dataConfig, err := yaml.Marshal(&yamlDB)
	err = ioutil.WriteFile(projectName + "/config/database.yaml", dataConfig, 0755)
	if err != nil {
        panic(err)
    }
	//create static
	os.Mkdir(projectName + "/static", 0777)
	os.Mkdir(projectName + "/static/images", 0777)
	os.Mkdir(projectName + "/static/css", 0777)
	os.Mkdir(projectName + "/static/js", 0777)
	//create vendor
	os.Mkdir(projectName + "/vendor", 0777)
	//create Readme file
	dataMD := []byte("# " + projectName + "\n You are welcome")
    err = ioutil.WriteFile(projectName + "/README.md", dataMD, 0755)
    if err != nil {
        panic(err)
    }
}