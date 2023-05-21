//dockerが起動した際にcmdフォルダ内のなにか？？？のコマンド？が使える？？

package main

import (
	"go_blog/internal/article"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", article.Index)
	http.HandleFunc("/show", article.Show)
	http.HandleFunc("/create", article.Create)
	http.HandleFunc("/edit", article.Edit)
	http.HandleFunc("/delete", article.Delete)
	//http.HandleFunc ...apiを作ってる？webページになるとは限らない

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
		// errはhttp.Li~~の返値でしかない
	}
}

// ListenAndServe　...サーバーのポートの解放
// サーバーを起動する
// 　第1引数にTCPのアドレス
// TCP
// 1対1のセッションによる信頼性の高い通信を行うためのプロトコル
// Transmission Control Protocol
//　 第2引数にhttp.Handler

// log.Fatal("ListenAndServe:", err)
// 記述　"ListenAndServe:～"
// errを返す処理が２つ：中身を記述したりlogに残すorパニック関数で処理終了
