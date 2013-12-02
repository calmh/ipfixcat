#!/bin/bash

BIN=ipfixcat
VER=$(git describe --always)

for GOARCH in amd64 386 ; do
	for GOOS in darwin linux freebsd freebsd ; do
		export GOOS
		export GOARCH

		NAME="$BIN-$GOOS-$GOARCH"
		rm -rf "$NAME" "$NAME.tar.gz" "$BIN" "$BIN.exe"
		go build -ldflags "-X main.ipfixcatVersion ${VER}"

		mkdir "$NAME"
		cp *.ini ipfixcat "$NAME"
		tar zcf "$NAME.tar.gz" "$NAME"
		rm -r "$NAME"
	done

	for GOOS in windows ; do
		export GOOS
		export GOARCH

		NAME="$BIN-$GOOS-$GOARCH"
		rm -rf "$NAME" "$NAME.tar.gz" "$BIN" "$BIN.exe"
		go build -ldflags "-X main.ipfixcatVersion ${VER}"

		mkdir "$NAME"
		cp *.ini ipfixcat.exe "$NAME"
		zip -r "$NAME.zip" "$NAME"
		rm -r "$NAME"
	done
done
