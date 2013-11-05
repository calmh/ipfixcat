BIN=ipfixcat
SRC=$(wildcard *.go)
VER=$(shell git describe --always)
NAME=ipfixcat-$(VER)
TGZ=$(NAME).tar.gz

all: $(TGZ)

$(TGZ): $(BIN) $(wildcard *.ini) README.md
	rm -rf $(NAME)
	mkdir $(NAME)
	cp $^ $(NAME)
	tar zcvf $@ $(NAME)

$(BIN): $(SRC)
	go build -ldflags "-X main.ipfixcatVersion ${VER}"

clean:
	rm -rf $(NAME) $(TGZ) $(BIN)

