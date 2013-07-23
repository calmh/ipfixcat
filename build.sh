#!/bin/sh

version="$(git describe --always)"
name="ipfixcat-$version"

echo "package main" > $GOPATH/src/github.com/calmh/ipfixcat/version.go
echo "var ipfixcatVersion = \"$version\"" >> $GOPATH/src/github.com/calmh/ipfixcat/version.go
rm $GOPATH/bin/ipfixcat
go install github.com/calmh/ipfixcat
mkdir $name
cp $GOPATH/bin/ipfixcat $name
cp $GOPATH/src/github.com/calmh/ipfixcat/*-fields.ini $name
cp $GOPATH/src/github.com/calmh/ipfixcat/README.md $name
tar zcvf $name.tar.gz $name
rm -r $name
