# Chapter 2 Examples

DIST = ../bin

all: $(DIST)/boiling $(DIST)/ftoc \
     $(DIST)/echo4 \
	 $(DIST)/tempconv0

# 2.2 Declarations

$(DIST)/boiling:
	go build -o $(DIST)/boiling ./boiling

.PHONY: boiling
boiling: $(DIST)/boiling
	$(DIST)/boiling

$(DIST)/ftoc:
	go build -o $(DIST)/ftoc ./ftoc

.PHONY: ftoc
ftoc: $(DIST)/ftoc
	$(DIST)/ftoc

# 2.3 Variables

$(DIST)/echo4:
	go build -o $(DIST)/echo4 ./echo4

.PHONY: echo4
echo4: $(DIST)/echo4
	$(DIST)/echo4 -s / a bc def

# 2.5 Type Declarations

$(DIST)/tempconv0:
	go build -o $(DIST)/tempconv0 ./tempconv0

.PHONY: tempconv0
tempconv0: $(DIST)/tempconv0
	$(DIST)/tempconv0

# 2.6 Packages and Files

$(DIST)/cf:
	go build -o $(DIST)/cf ./cf

.PHONY: cf
cf: $(DIST)/cf
	$(DIST)/cf 32
	$(DIST)/cf 212
	$(DIST)/cf -40

popcount:
	go test -v ./internal/popcount

popcount-bench:
	go test -v -bench=. -benchmem ./internal/popcount