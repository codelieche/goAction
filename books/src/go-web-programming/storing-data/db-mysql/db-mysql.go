package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/**
go get -u github.com/go-sql-driver/mysql
*/

var Db *sql.DB

func init() {
	var err error
	dbSourceName := "root:root@127.0.0.1：3307/study?charset=UTF8"
	Db, err = sql.Open("mysql", dbSourceName)
	if err != nil {
		panic(err)
	} else {
		log.Println(Db)
	}
}

func createTable() {
	log.Println("创建表格开始")
	var sql string = "CREATE TABLE `posts2` (`id` int(11) unsigned NOT NULL AUTO_INCREMENT," +
		"`title` varchar(108) NOT NULL DEFAULT '', `author` varchar(40) DEFAULT NULL," +
		"`content` text NOT NULL,PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"
	if result, err := Db.Exec(sql); err != nil {
		log.Println("创建表格失败！", err.Error())
		panic(err)
	} else {
		log.Println("创建表格成功")
		log.Println(result)
	}
}

func main() {
	createTable()
}
