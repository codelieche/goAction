package main

import (
	"database/sql"

	"log"

	_ "github.com/lib/pq"
)

type Post struct {
	Id      int64  `json:"id"`      // 帖子ID
	Title   string `json:"title"`   // 帖子标题
	Content string `json:"content"` // 帖子内容
	Author  string `json:"author"`  // 作者
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres dbname=postgres password=''"+
		"sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		log.Println(Db)
	}
}

func createTable() {
	log.Println("创建表格开始")
	var sql string = "CREATE TABLE posts (id serial primary key, title varchar(108), author varchar(40), content text);"
	if result, err := Db.Exec(sql); err != nil {
		log.Println("创建表格失败！", err.Error())
		panic(err)
	} else {
		log.Println("创建表格成功")
		log.Println(result)
	}
}

func (p *Post) Create() (err error) {
	// 把Post插入到Db中
	statement := "INSERT INTO posts (title, author, content) VALUES($1, $2, $3) returning id"
	if stmt, err := Db.Prepare(statement); err != nil {
		panic(err)
	} else {
		defer stmt.Close()
		err = stmt.QueryRow(p.Title, p.Author, p.Content).Scan(&p.Id)
		return err
	}
}

func main() {
	//createTable()
	log.Println("Start...")
	// 创建帖子
	p := Post{Title: "Hello Golang", Author: "codelieche", Content: "## Golang \n > This is a good!"}
	log.Println(p)
	p.Create()
	log.Println(p)

}
