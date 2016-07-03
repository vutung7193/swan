package swan 

import (
	"gopkg.in/yaml.v2"
	"os"
	// "fmt"
	// "encoding/gob"
    // "bytes"
  	"io/ioutil"
)
type DBConfig struct {
	Adapter       string `yaml:"adapter"`
	Host          string `yaml:"host"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	Dbname        string `yaml:"dbname"`
	Encoding      string `yaml:"encoding"`
	Port          string `yaml:"port"`
}

var db *DBConfig;

func GetDb(s *SwanConfig) (*DBConfig, error){
	if(db == nil) {
		db = &DBConfig{};
		filename := s.GetDbConfigPath()
		if _, err := os.Stat(filename); err == nil {
			yamlFile, err := ioutil.ReadFile(filename)
		    if err != nil {
		        return nil, err
		    }
		    m := make(map[interface{}]interface{})
		    err = yaml.Unmarshal(yamlFile, &m)
		    if err != nil {
		    	return nil, err
		    }
		    
		    d, err := yaml.Marshal(m[s.GetEnv()])
		    if err != nil {
	    		return nil, err
		    }

		    err = yaml.Unmarshal(d, &db)
		    if err != nil {
		    	return nil, err
		    }
		} else {
			return nil, err
		}
	}
	return db, nil
}