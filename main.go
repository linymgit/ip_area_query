package main

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
	"ip_area_query/config"
	"ip_area_query/model"
	"log"
)

func main() {

	config.InitConfigure()

	model.InitMysql()

	region, err := ip2region.New("./ip2region.db")
	defer region.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		logLogins := model.ListNullRegion()
		for _, login := range logLogins {
			ip, err := region.MemorySearch(login.IP)
			if err != nil {
				log.Print(err, login)
			} else {
				login.Region = ip.Country
				login.UpdateRegion()
			}
		}
		if len(logLogins) < 100 {
			break
		}
	}

}
