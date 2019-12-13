package main

import (
	"fmt"
	"net/http"
)

func receiver(w http.ResponseWriter, r *http.Request){
	// クエリパラメータ取得してみる
	fmt.Fprintf(w, "クエリ：%s\n", r.URL.RawQuery)

	// Bodyデータを扱う場合には、事前にパースを行う
	r.ParseForm()

	// Formデータを取得.
	form := r.PostForm
	fmt.Fprintf(w, "フォーム：\n%v\n", form)
	fmt.Fprintf(w, "フォーム：\n%v\n", form["password"])



	// または、クエリパラメータも含めて全部.
	params := r.Form
	fmt.Fprintf(w, "フォーム2：\n%v\n", params)
}



func main() {
	http.Handle("/get/", http.StripPrefix("/get/", http.FileServer(http.Dir("./html/get/"))))
	http.Handle("/post/", http.StripPrefix("/post/", http.FileServer(http.Dir("./html/post/"))))
	http.HandleFunc("/receive", receiver)
	http.ListenAndServe(":8080", nil)
}