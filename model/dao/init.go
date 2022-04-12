package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kzfighting/config"
)

var DB *gorm.DB

func DBinit(con config.Config) {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?%v",
		con.Mysql.Username,
		con.Mysql.Password,
		con.Mysql.Host,
		con.Mysql.Port,
		con.Mysql.Dbname,
		con.Mysql.Param,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate()
	if err != nil {
		panic(err)
	}

	return
}
