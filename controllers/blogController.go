package controllers

import (
"html/template"
"net/http"
"strconv"
	"excel/models"
	"fmt"
	"log"
	"encoding/json"
)

const upload_dir  = "/tmp"

type BlogController struct {
	baseController
}


//http://localhost:3000
func (c BlogController) Home(w http.ResponseWriter,r *http.Request)  {

	fmt.Fprintln(w,"home! access /public/")
}


//http://localhost:3000/index
func (c BlogController) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		showError(w, "异常", "非法请求，服务器无法响应")
	} else {
		if r.URL.Path == "/index" {
			tags, err := models.QueryAll()
			/*see := http.Cookie{
							Name:"tag",
							Value:"frist",
							HttpOnly:true,
						}
						http.SetCookie(w,&see)*/
			log.Println(tags)
			if err != nil {
				log.Println(err)
				showError(w, "异常", "查询异常")
				return
			}
			//t, err := template.ParseFiles("views/home.html")
			//if err != nil {
			//	showError(w, "异常", "页面渲染异常")
			//	return
			//}
			data := make(map[string][]models.Tag)
			data["list"] = tags

			//json.Marshal(data)
			rlt,_ := json.Marshal(data)
			fmt.Fprintln(w,string(rlt))
			//t.Execute(w, data)
		} else {
			// 404页面，路由不到的都会到这里
			showError(w, "404", "页面不存在")
		}
	}
}

func NewTag(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		showError(w, "异常", "非法请求")
	} else {
		title := r.FormValue("title")
		local_path := r.FormValue("img_url")
		img_url := upload_dir+local_path
		id, err := models.InsertTag(title,img_url)
		if err != nil || id <= 0 {
			showError(w, "异常", "插入数据异常")
			return
		}
		// 重定向到主界面
		http.Redirect(w, r, "/", http.StatusSeeOther)
		// 没有return，没有效果，重定向不过去
		return
	}
}

func DeleteTag(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		showError(w, "异常", "非法请求")
	} else {
		id := r.FormValue("id")
		intId, _ := strconv.ParseInt(id, 10, 64)
		_, err := models.DeleteTag(intId)
		if err != nil {
			showError(w, "异常", "删除失败")
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

// 错误处理
func showError(w http.ResponseWriter, title string, message string) {
	t, _ := template.ParseFiles("views/error.html")
	data := make(map[string]string)
	data["title"] = title
	data["message"] = message
	t.Execute(w, data)
}


