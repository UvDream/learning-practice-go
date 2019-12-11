package main

import (
	config "config/conf"
	"fmt"
)

func main() {
	var c config.Conf
	conf := *c.GetConf()
	fmt.Println(conf.Host)
}
