package main

import (
	_ "github.com/SecKill/SecProxy/router"
	"github.com/astaxie/beego"
)

func main() {
	err := initConfig()
	if err != nil {
		panic(err)
		return
	}

	err = initSec()
	if err != nil {
		panic(err)
		return
	}

	beego.Run()
}
