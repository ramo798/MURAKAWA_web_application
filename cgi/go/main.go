package main

import (
	"net/http"
	"strings"

	"text/template"
)

type Query struct {
	Method string
	Id string
	Password string
	Anonymous string
	Kikazaru string
	Mizaru string
	Syaberazaru string
}

func receiver(w http.ResponseWriter, r *http.Request){
	//構造体の宣言
	query := new(Query)

	//メソッドチェック
	if r.Method == http.MethodPost {
		query.Method = "POST"
	}else{
		query.Method = "GET"
	}

	// Formデータを取得。mapに変換してから使用している。
	//map[anonymous:[on] id:[root] kikazaru:[1] mizaru:[1] password:[root] syaberazaru:[1]]
	r.ParseForm()
	parameter := r.Form

	//[]string型だったのでstringにキャストしてる
	query.Id = strings.Join(parameter["id"],"")
	query.Password = strings.Join(parameter["password"],"")
	query.Anonymous = strings.Join(parameter["anonymous"],"")

	if strings.Join(parameter["kikazaru"],"") == "1" {
		query.Kikazaru = "checked='checked'"
	}else{
		query.Kikazaru = ""
	}

	if strings.Join(parameter["mizaru"],"") == "1" {
		query.Mizaru = "checked='checked'"
	}else{
		query.Mizaru = ""
	}

	if strings.Join(parameter["syaberazaru"],"") == "1" {
		query.Syaberazaru = "checked='checked'"
	}else{
		query.Syaberazaru = ""
	}


	//htmlに埋め込んで返す
	tmpl, err := template.ParseFiles("html/receive/index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, query)
	if err != nil {
		panic(err)
	}
}

func main() {
	//ハンドラー
	//静的ファイル返す
	http.Handle("/get/", http.StripPrefix("/get/", http.FileServer(http.Dir("./html/get/"))))
	http.Handle("/post/", http.StripPrefix("/post/", http.FileServer(http.Dir("./html/post/"))))
	//func receiverの処理させる
	http.HandleFunc("/receive", receiver)

	//8080でリスニング
	http.ListenAndServe(":8080", nil)
}