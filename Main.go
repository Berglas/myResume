package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

//var gConn *sql.DB

func main() {
	//設定資料庫
	// gConn = sqlutil.Conn("host=127.0.0.1 port=5433 dbname=resume user=postgres password=1234 sslmode=disable")
	// if gConn == nil {
	// 	return
	// }
	// defer gConn.Close()

	//設定路由/路徑
	fmt.Println("start")
	http.HandleFunc("/Home", Home)
	http.Handle("/css/", http.FileServer(http.Dir("")))
	http.Handle("/js/", http.FileServer(http.Dir("")))
	http.Handle("/image/", http.FileServer(http.Dir("")))
	http.Handle("/fonts/", http.FileServer(http.Dir("")))
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
