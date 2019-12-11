package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

//profile variables
type Conf struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}

func (c *Conf) GetConf() *Conf {
	path, err := filepath.Abs("config.yaml")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}

	fmt.Println("path: ", path)

	yamlFile, err1 := ioutil.ReadFile(path)
	if err1 != nil {
		fmt.Println(err1.Error())
		return nil
	}

	err1 = yaml.Unmarshal(yamlFile, c)
	if err1 != nil {
		fmt.Println(err1.Error())
		return nil
	}

	return c
}
