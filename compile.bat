@echo off
set ver=v0.1.2
set GOOS=linux
go build -ldflags="-n -s -X 'main.Version=%ver%'" -o eseguibili\

set GOOS=darwin
go build -ldflags="-n -s -X 'main.Version=%ver%'" -o eseguibili\goscannermac

set GOOS=windows
go build -ldflags="-n -s -X 'main.Version=%ver%'" -o eseguibili\