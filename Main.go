package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"./sqlutil"
	_ "github.com/lib/pq"
)

var gConn *sql.DB

func main() {
	//設定資料庫
	gConn = sqlutil.Conn("host=127.0.0.1 dbname=Resumes user=postgres password=1234 sslmode=disable")
	if gConn == nil {
		return
	}
	defer gConn.Close()

	//設定路由/路徑
	fmt.Println("start")
	http.HandleFunc("view/Home", Home)
	http.HandleFunc("Resumes", Resumes)
	http.HandleFunc("/Teach", Teach)
	http.HandleFunc("/selectMenu", selectMenu)
	http.HandleFunc("/SelectTeach", SelectTeach)
	http.Handle("/css/", http.FileServer(http.Dir("")))
	http.Handle("/js/", http.FileServer(http.Dir("")))
	http.Handle("/image/", http.FileServer(http.Dir("")))
	http.Handle("/view/", http.FileServer(http.Dir("")))
	//設定Server
	err := http.ListenAndServe(":443", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("Home.html")
	t.Execute(w, nil)
}

func Resumes(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("Resumes.html")
	t.Execute(w, nil)
}

func Teach(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("Teach.html")
	t.Execute(w, nil)
}

//查詢Menu項目
func selectMenu(w http.ResponseWriter, r *http.Request) {
	rows, err := gConn.Query("SELECT * FROM \"home_item\" ORDER BY \"index\" ASC")
	if err != nil {
		fmt.Println(err)
	}

	var index string
	var chinese_name string
	var english_name string
	var html_url string
	m := []map[string]interface{}{}
	for rows.Next() {
		err = rows.Scan(&index, &chinese_name, &english_name, &html_url)
		if err != nil {
			fmt.Println(err)
		}
		temp := map[string]interface{}{
			"index":        index,
			"chinese_name": chinese_name,
			"english_name": english_name,
			"html_url":     html_url}

		m = append(m, temp)
	}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Fprint(w, string(b))

}

func SelectTeach(w http.ResponseWriter, r *http.Request) {
	rows, err := gConn.Query("SELECT * FROM \"teach_item\" ORDER BY \"index\" ASC")
	if err != nil {
		fmt.Println(err)
	}

	var index string
	var url string
	var title string
	var content string
	var image_url string
	m := []map[string]interface{}{}
	for rows.Next() {
		err = rows.Scan(&index, &url, &title, &content, &image_url)
		if err != nil {
			fmt.Println(err)
		}
		temp := map[string]interface{}{
			"index":     index,
			"url":       url,
			"title":     title,
			"content":   content,
			"image_url": image_url}

		m = append(m, temp)
	}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Fprint(w, string(b))
}

func insertMarquee(w http.ResponseWriter, r *http.Request) {
	rows, err := gConn.Query("SELECT Max(index) as maximum FROM marquee")
	if err != nil {
		log.Println(err)
	}

	var maximum int
	for rows.Next() {
		err = rows.Scan(&maximum)
		if err != nil {
			log.Println(err)
		}
	}
	tx, _ := gConn.Begin()
	_, err = tx.Exec(fmt.Sprintf("INSERT INTO marquee VALUES (%v, '%v', '%v', '%v')", maximum+1, r.FormValue("text"), r.FormValue("startTime"), r.FormValue("endTime")))
	if err != nil {
		log.Println(err)
		tx.Rollback()
		fmt.Fprint(w, "新增失敗")
	} else {
		tx.Commit()
		fmt.Fprint(w, "新增成功")
	}
}

func updateMarquee(w http.ResponseWriter, r *http.Request) {
	tx, _ := gConn.Begin()
	_, err := tx.Exec(fmt.Sprintf("UPDATE marquee SET text='%v',start_time='%v', end_time='%v' WHERE index=%v;", r.FormValue("text"), r.FormValue("startTime"), r.FormValue("endTime"), r.FormValue("index")))
	if err != nil {
		log.Println(err)
		tx.Rollback()
		fmt.Fprint(w, "修改失敗")
	} else {
		tx.Commit()
		fmt.Fprint(w, "修改成功")
	}
}
