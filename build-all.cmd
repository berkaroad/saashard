@echo off

go get github.com/go-yaml/yaml

go tool yacc -o %~dp0sqlparser\yacc.go -v %~dp0sqlparser\yacc.output %~dp0sqlparser\yacc.y && gofmt -w %~dp0sqlparser\yacc.go

rem build-darwin
set GOOS=darwin
go build -ldflags "-s -w" -o %~dp0bin/saashard_darwin github.com/berkaroad/saashard/cmd/saashard

rem build-linux
set GOOS=linux
go build -ldflags "-s -w" -o %~dp0bin/saashard_linux github.com/berkaroad/saashard/cmd/saashard

rem build-windows
set GOOS=windows
go build -ldflags "-s -w" -o %~dp0bin/saashard.exe github.com/berkaroad/saashard/cmd/saashard

@pause