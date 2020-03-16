@echo off
set ver=v0.2.1
set GOOS=linux
go build -ldflags="-n -s -X 'main.Version=%ver%'" -o ..\eseguibili\goscanner-linux

set GOOS=darwin
go build -ldflags="-n -s -X 'main.Version=%ver%'" -o ..\eseguibili\goscannermac

set GOOS=windows
go build -ldflags="-n -s -X 'main.Version=%ver%'" -o ..\eseguibili\goscanner.exe