package model

import (
	"fmt"
	"ip_area_query/config"
	"testing"
)

func TestMain(m *testing.M) {
	config.InitConfigure()
	InitMysql()
	m.Run()
}

func TestListNullRegion(t *testing.T) {
	region := ListNullRegion()
	fmt.Println(fmt.Sprintf("%#v", region))
}
