# Go言語であるためコメントアウト
# CFLAGS=-std=c11 -g -static

# 9cc -> 9gg に変更
# .c -> .go に変更
9gg: 9gg.go
	go build 9gg.go

# 9cc -> 9gg に変更
test: 9gg
	./test.sh

# 9cc -> 9gg に変更
clean:
	rm -f 9gg *.o *~ tmp*

.PHONY: test clean