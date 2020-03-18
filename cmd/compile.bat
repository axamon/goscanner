@echo off
set ver=v0.2.7

set GOOS=openbsd
go build -ldflags="-n -s -X 'main.Version=%ver%'" -o ..\eseguibili\goscanner-openbsd

set GOOS=linux
go build -ldflags="-n -s -X 'main.Version=%ver%'" -o ..\eseguibili\goscanner-linux

set GOOS=darwin
go build -ldflags="-n -s -X 'main.Version=%ver%'" -o ..\eseguibili\goscannermac

set GOOS=windows
go build -ldflags="-n -s -X 'main.Version=%ver%'" -o ..\eseguibili\goscanner.exe

git add .
git commit -m "Version: %ver%"