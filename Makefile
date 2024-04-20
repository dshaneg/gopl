# Chapter 1 Examples

.PHONY: ch1/lissajous
ch1/lissajous:
	go run ch1/lissajous/main.go > /mnt/d/Users/dshan/Pictures/liss.gif.PHONY: ch1/fetch

.PHONY: ch1/fetch
ch1/fetch:
	go run ./ch1/fetch https://google.com news.ycombinator.co
	
.PHONY: ch1/fetchall
ch1/fetchall:
	go run ./ch1/fetchall https://golang.org http://gopl.io https://godoc.org

.PHONY: ch1/server1
ch1/server1:
	go run ./ch1/server1

.PHONY: ch1/server2
ch1/server2:
	go run ./ch1/server2

.PHONY: ch1/server3
ch1/server3:
	go run ./ch1/server3

# Chapter 2 Examples

.PHONY: ch2/boiling
ch2/boiling:
	go run ./ch2/boiling

.PHONY: ch2/ftoc
ch2/ftoc:
	go run ./ch2/ftoc

# need to build this one then run it from the executable because the command line arguments didn't play nice with go run (even with --)
.PHONY: ch2/echo4-build
ch2/echo4-build:
	go build -o ./bin/echo4 ./ch2/echo4

.PHONY: ch2/tempconv0
ch2/tempconv0:
	go run ./ch2/tempconv0
