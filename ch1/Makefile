# Chapter 1 Examples

DIST = ../bin

all: $(DIST)/helloworld \
     $(DIST)/echo1 $(DIST)/echo2 $(DIST)/echo3 \
     $(DIST)/dup1 $(DIST)/dup2 $(DIST)/dup3 \
     $(DIST)/lissajous \
     $(DIST)/fetch $(DIST)/fetchall \
     $(DIST)/server1 $(DIST)/server2 $(DIST)/server3 \

# 1.1 Hello, world!

$(DIST)/helloworld:
	go build -o $(DIST)/helloworld ./helloworld

.PHONY: helloworld
helloworld:
	$(DIST)/helloworld

# 1.2 Command-Line Arguments

$(DIST)/echo1:
	go build -o $(DIST)/echo1 ./echo1

$(DIST)/echo2:
	go build -o $(DIST)/echo2 ./echo2

$(DIST)/echo3:
	go build -o $(DIST)/echo3 ./echo3

# 1.3 Finding Duplicate Lines

$(DIST)/dup1:
	go build -o $(DIST)/dup1 ./dup1

.PHONY: dup1
dup1: $(DIST)/dup1
	cat data/hobbits.txt | $(DIST)/dup1

$(DIST)/dup2:
	go build -o $(DIST)/dup2 ./dup2

.PHONY: dup2
dup2: $(DIST)/dup2
	cat data/hobbits.txt | $(DIST)/dup2
	$(DIST)/dup2 ./data/hobbits.txt ./data/humans.txt ./data/fantasyfolk.txt

$(DIST)/dup3:
	go build -o $(DIST)/dup3 ./dup3

.PHONY: dup3
dup3: $(DIST)/dup3
	$(DIST)/dup3 ./data/hobbits.txt ./data/humans.txt ./data/fantasyfolk.txt

# 1.4 Animated GIFs

$(DIST)/lissajous:
	go build -o $(DIST)/lissajous ./lissajous

.PHONY: lissajous
lissajous: $(DIST)/lissajous
	$(DIST)/lissajous > /mnt/d/Users/dshan/Pictures/liss.gif

# 1.5 Fetching a URL

$(DIST)/fetch:
	go build -o $(DIST)/fetch ./fetch

.PHONY: fetch
fetch: $(DIST)/fetch
	$(DIST)/fetch https://google.com news.ycombinator.co
	
# 1.6 Fetching URLs Concurrently

$(DIST)/fetchall:
	go build -o $(DIST)/fetchall ./fetchall

.PHONY: fetchall
fetchall: $(DIST)/fetchall
	$(DIST)/fetchall https://golang.org http://gopl.io https://godoc.org

# 1.7 A Web Server

$(DIST)/server1:
	go build -o $(DIST)/server1 ./server1

.PHONY: server1
server1: $(DIST)/server1
	$(DIST)/server1

$(DIST)/server2:
	go build -o $(DIST)/server2 ./server2

.PHONY: server2
server2: $(DIST)/server2
	$(DIST)/server2

$(DIST)/server3:
	go build -o $(DIST)/server3 ./server3

.PHONY: server3
server3: $(DIST)/server3
	$(DIST)/server3
