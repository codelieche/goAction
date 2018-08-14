package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

/**
把帖子数据存储到内存中
api说明：
1. /post/create： POST创建帖子
2. /post/list 获取所有的帖子
3. /post/:id 获取帖子的内容
*/

type Post struct {
	Id      int64  `json:"id"`      // 帖子ID
	Title   string `json:"title"`   // 帖子标题
	Content string `json:"content"` // 帖子内容
	Author  string `json:"author"`  // 作者
}

func (p *Post) toString() string {
	return fmt.Sprintf("ID: %d: Author: %s: Title: %s", p.Id, p.Author, p.Title)
}

var PostById = make(map[int64]*Post)         // 根据ID找到帖子
var PostsByAuthor = make(map[string][]*Post) // 根据作者名字获取帖子列表

func storePost(post Post) {
	// 存储帖子数据
	// 方便后续通过id、author查询
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

// Web服务的处理函数
var currentPostId int64 = 0

func handlePostCreate(w http.ResponseWriter, r *http.Request) {
	// 不做权限判断，谁都可以创建帖子
	r.ParseMultipartForm(r.ContentLength)
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("title")
	author := r.PostForm.Get("author")
	log.Print(r.Form)
	if title == "" || content == "" || author == "" {
		http.Error(w, "传入的参数有误", 400)
		return
	} else {
		currentPostId += 1
		post := Post{currentPostId, title, content, author}
		storePost(post)
		if data, err := json.Marshal(post); err == nil {
			w.Header().Add("Content-Type", "application/json")
			w.Write(data)
		} else {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}

func handlePostDetail(w http.ResponseWriter, r *http.Request) {
	// 通过get获取id
	r.ParseForm()
	idStr := r.Form.Get("id")
	if id, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		http.Error(w, "参数有误", 400)
		return
	} else {
		p := PostById[id]
		if p != nil {
			if data, err := json.Marshal(*p); err == nil {
				w.Header().Add("Content-Type", "application/json")
				w.Write(data)
			} else {
				http.Error(w, err.Error(), 400)
				return
			}
		} else {
			http.Error(w, "Page Not Found", 400)
			return
		}
	}

}

func main() {

	server := http.Server{
		Addr: ":9000",
	}
	http.HandleFunc("/post/create", handlePostCreate)
	http.HandleFunc("/post", handlePostDetail)

	server.ListenAndServe()
}
