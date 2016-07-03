package swan

import (
	"gopkg.in/yaml.v2"
	"sync"
	"os"
  	"path/filepath"
  	"io/ioutil"
)
type SwanConfig struct {
	DbConfigPath 	string `yaml:"database_config_path"`
	MigratePath 	string `yaml:"migrate_path"`
	ControllerPath 	string `yaml:"controller_path"`
	ModelPath		string `yaml:"model_path"`
	ViewPath		string `yaml:"view_path"`
	Env				string `yaml:"environment"`
}

var cwd string

func (c SwanConfig) GetDbConfigPath() string {
	return cwd + c.DbConfigPath
}

func (c SwanConfig) GetMigratePath() string {
	return cwd + c.MigratePath
}

func (c SwanConfig) GetControllerPath() string {
	return cwd + c.ControllerPath
}

func (c SwanConfig) GetModelPath() string {
	return cwd + c.ModelPath
}
func (c SwanConfig) GetViewPath() string {
	return cwd + c.ViewPath
}

func (c SwanConfig) GetEnv() string {
	return c.Env
}
var scf *SwanConfig
var once sync.Once


func GetSwanConfig() *SwanConfig {
  once.Do(func() {
  	cwd, _ = os.Getwd()
  	cwd = cwd + "/"

    scf = &SwanConfig{
    	DbConfigPath: "/config/database.yml",
    	MigratePath: "/migrate",
    	ControllerPath: "/app/controllers",
    	ModelPath: "/app/models",
    	ViewPath: "/app/views",
    	Env: "development",
    }
    
    filename, _ := filepath.Abs("./config/swan.yml")
    if _, err := os.Stat(filename); err == nil {
	  	yamlFile, err := ioutil.ReadFile(filename)

	    if err != nil {
	        panic(err)
	    }

	    err = yaml.Unmarshal(yamlFile, &scf)
	    if err != nil {
	        panic(err)
	    }
	}

  })
  return scf
}