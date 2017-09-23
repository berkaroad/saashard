all: build


build: saashard
	go build -o ./bin/saashard ./cmd/saashard

build-all: build-linux build-windows build-darwin

build-darwin: saashard
	export GOOS=darwin && go build -ldflags "-s -w" -o ./bin/saashard_darwin ./cmd/saashard

build-linux: saashard
	export GOOS=linux && go build -ldflags "-s -w" -o ./bin/saashard_linux ./cmd/saashard

build-windows: saashard
	export GOOS=windows && go build -ldflags "-s -w" -o ./bin/saashard.exe ./cmd/saashard

saashard:
	go get github.com/go-yaml/yaml
	# cd ./sqlparser && go tool yacc -o yacc.go -v yacc.output yacc.y && gofmt -w yacc.go
	@bash genver.sh

run: build
	./bin/saashard --config=conf/ss.yaml --cpuprofile=bin/saashard.cpuprof --memprofile=bin/saashard.memprof

dev: build
	./bin/saashard --config=conf/dev.yaml --cpuprofile=bin/saashard.cpuprof --memprofile=bin/saashard.memprof

test: saashard
	go install github.com/berkaroad/saashard/utils/simplelog
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

	go test github.com/berkaroad/saashard/sqlparser

clean:
	@rm -rf bin
	@rm -f ./sqlparse/yacc.output
