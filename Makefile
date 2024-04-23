all: ch1 ch2

.PHONY: ch1
ch1:
	$(MAKE) -C ./ch1

.PHONY: ch2
ch2:
	$(MAKE) -C ./ch2

clean:
	rm ./bin/*
