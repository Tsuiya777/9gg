package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println("Hello -9gg- World")

	// main関数の引数の取得
	flag.Parse()
	args := flag.Args()
	// fmt.Println(args)

	// 引数の個数が１個であることを確認
	// ※Goでは実行ファイル自身の名前は渡されない
	if len(args) != 1 {
		fmt.Printf("引数の個数が正しくありません。len=%d\n", len(args))
		return
	}

	s, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("引数が整数ではありません。value=%s\n", args[0])
		return
	}

	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".global main\n")
	fmt.Printf("main:\n")
	fmt.Printf("	mov rax, %d\n", s)
	fmt.Printf("	ret\n")

	return
}
