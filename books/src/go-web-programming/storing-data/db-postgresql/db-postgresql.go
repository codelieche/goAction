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

func (p *Post) Delete() error {
	// 删除Post
	query := "DELETE FROM posts where id = $1"
	if result, err := Db.Exec(query, p.Id); err != nil {
		log.Println("删除Post出错:", err.Error())
		return err
	} else {
		log.Println(result)
		return nil
	}
}

func (p *Post) Update() (err error) {
	// 更新Post
	query := "UPDATE posts SET title=$2, author=$3, content=$4 WHERE id = $1"
	if result, err := Db.Exec(query, p.Id, p.Title, p.Author, p.Content); err != nil {
		panic(err)
	} else {
		log.Println(result)
		return nil
	}
}

func getPost(id int) (post Post, err error) {
	post = Post{}
	query := "SELECT id, title, author, content FROM posts WHERE id = $1"
	err = Db.QueryRow(query, id).Scan(&post.Id, &post.Title, &post.Author, &post.Content)
	return
}

func runCreate() {
	// 创建帖子
	p := Post{Title: "Hello Golang", Author: "codelieche", Content: "## Golang \n > This is a good!"}
	log.Println(p)
	p.Create()
	log.Println(p)
}

func runDelete(id int) {
	// 删除帖子
	// 1. 先获取到帖子
	if p, err := getPost(id); err != nil {
		panic(err)
	} else {
		log.Println(p)
		p.Delete()
	}
}

func runUpdate(id int) {
	// 更新帖子
	// 1. 先获取到帖子
	if p, err := getPost(id); err != nil {
		panic(err)
	} else {
		log.Println(p)
		p.Author = "update"
		p.Content = "Update Test"
		p.Update()
	}
}

func main() {
	//createTable()
	log.Println("Start...")

	//runCreate()

	//runDelete(10)

	runUpdate(9)
}
