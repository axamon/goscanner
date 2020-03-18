set ver=v0.2.11

rem Forza lo stile di scrittura del codice su tutto il progetto.
gofmt -w .\

rem Crea eseguibile di test e lo lancia per test funzionali e verifica della coverage.
rem Se ci sono problemi esce senza proseguire con la build.
go test -v -cover -o goscannerTest.exe || goto exit

set GOOS=openbsd
go build .\cmd -ldflags="-n -s -X 'main.Version=%ver%'" -o ..\eseguibili\goscanner-openbsd

set GOOS=linux
go build .\cmd -ldflags="-n -s -X 'main.Version=%ver%'" -o ..\eseguibili\goscanner-linux

set GOOS=darwin
go build .\cmd -ldflags="-n -s -X 'main.Version=%ver%'" -o ..\eseguibili\goscannermac

set GOOS=windows
go build .\cmd -ldflags="-n -s -X 'main.Version=%ver%'" -o ..\eseguibili\goscanner.exe

git checkout -b %ver%
git add .
git commit -m "Version: %ver%"

:exit