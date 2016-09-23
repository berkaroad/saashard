@echo off

go get github.com/go-yaml/yaml

cd %~dp0sqlparser
go tool yacc -o yacc.go -v yacc.output yacc.y && gofmt -w yacc.go
cd %~dp0

go install github.com/berkaroad/saashard/utils/golog
go install github.com/berkaroad/saashard/config
go install github.com/berkaroad/saashard/errors
go install github.com/berkaroad/saashard/sqlparser/sqltypes
go install github.com/berkaroad/saashard/sqlparser
go install github.com/berkaroad/saashard/backend/mysql
go install github.com/berkaroad/saashard/backend
go install github.com/berkaroad/saashard/net/mysql
go install github.com/berkaroad/saashard/route
go install github.com/berkaroad/saashard/proxy
go install github.com/berkaroad/saashard/admin
go install github.com/berkaroad/saashard/statistic
go install github.com/berkaroad/saashard/server

go build -o %~dp0bin\saashard.exe github.com/berkaroad/saashard/cmd/saashard
%~dp0bin\saashard.exe -config %~dp0conf\ss.yaml

@pause