package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Printf("Please enter username.")
		return
	}

	// POSTメソッド
	// url.Values{}でPOSTで送信する入れ物を準備
	ps := url.Values{}

	// Add()でPOSTで送信するデータを作成
	ps.Add("name", args[0])
	// 特殊文字や日本語をエンコード
	fmt.Println(ps.Encode())

	// http.PostForm()でPOSTメソッドを発行
	res, err := http.PostForm("http://localhost:8080/user", ps)
	if err != nil {
		log.Fatal(err)
	}
	// deferでクローズ処理
	defer res.Body.Close()
	// Bodyの内容を読み込む
	body, _ := io.ReadAll(res.Body)
	// Bodyの内容を出力する
	fmt.Print(string(body))
}
