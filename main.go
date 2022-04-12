package main

import (
	"kzfighting/config"
	"kzfighting/controller"
	"kzfighting/model/dao"
)

var con config.Config

func main() {
	config.Initial(&con)
	dao.DBinit(con)

	controller.Run()
}
