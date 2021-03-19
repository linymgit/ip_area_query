package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	Mysql struct {
		Dsn             string
		MaxIdleConns    int
		OpenConns       int
		ConnMaxLifetime int
	}
}

var C *Conf

func InitConfigure() {

	bytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	C = new(Conf)
	err = yaml.Unmarshal(bytes, C)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%#v", C))

}
