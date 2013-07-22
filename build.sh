#!/bin/sh

version="$(git describe --always)"
name="ipfixcat-$version"

rm $GOPATH/bin/ipfixcat
go install github.com/calmh/ipfixcat
mkdir $name
cp $GOPATH/bin/ipfixcat $name
cp $GOPATH/src/github.com/calmh/ipfixcat/*-fields.ini $name
tar zcvf $name.tar.gz $name
rm -r $name
