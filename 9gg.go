package main

import (
	"flag"
	"fmt"
	"strconv"
)

func lenOfDigits(input string) int {
	if input == "" {
		fmt.Println("lenOfDigits:引数の長さが０です")
		return 0
	}

	// ASCIIコードの０(0x30)～９(0x39)が何文字文(Length)あるか確認
	var length int
	// fmt.Println("lenOfDigits:len=", len(input))
	for i := 0; i < len(input); i++ {
		// fmt.Println("lenOfDigits:num=", string(input[i]))
		if 0x30 <= input[i] && input[i] <= 0x39 {
			length++
		} else {
			break
		}
	}
	return length
}

func byte2int(input string) (int, int) {

	// 数値が何文字あるか調査
	var length = lenOfDigits(input)

	if length == 0 {
		fmt.Printf("byte2int:整数が入力されていません")
		return 0, -1
	}

	// lenghtの長さのByte[]を作成
	s := input[0:length]
	output, err := strconv.Atoi(string(s))
	if err != nil {
		fmt.Printf("byte2int:整数に変換できません")
		return 0, -1
	}

	return output, length
}

func main() {
	// fmt.Println("Hello -9gg- World")

	// main関数の引数の取得
	flag.Parse()
	args := flag.Args()
	// fmt.Println(args)

	// 引数の個数が１個であることを確認
	// ※Goでは実行ファイル自身の名前は渡されない
	if len(args) != 1 {
		fmt.Printf("main:引数の個数が正しくありません。len=%d\n", len(args))
		return
	}

	var input = args[0]

	// fmt.Println(string(input[0]))
	// num, err := strconv.Atoi(string(input[0]))
	// if err != nil {
	// 	fmt.Printf("main:引数が整数ではありません。value=%s\n", args[0])
	// 	return
	// }

	var temp, length = byte2int(input)
	if length == -1 {
		fmt.Printf("main:引数が整数ではありません。value=%s\n", string(input[0]))
	}

	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".global main\n")
	fmt.Printf("main:\n")
	fmt.Printf("	mov rax, %d\n", temp)

	// inputが数字として読み込んだ文字数以下であれば
	if len(input) <= length {
		fmt.Printf("	ret\n")
		return
	}

	// 数字として読み込んだlength分ずらして読み始める
	for i := length; i < len(input); i++ {
		// fmt.Println("main:input[i]=", string(input[i]))
		if input[i] == '+' {
			i++
			temp, length = byte2int(input[i:])
			fmt.Printf("	add rax, %d\n", temp)
			i += length - 1
			continue
		}

		if input[i] == '-' {
			i++
			temp, length = byte2int(input[i:])
			fmt.Printf("	sub rax, %d\n", temp)
			i += length - 1
			continue
		}

		fmt.Printf("main:予期しない文字です： '%s'\n", string(input[i]))
		return
	}

	fmt.Printf("	ret\n")
	return
}
