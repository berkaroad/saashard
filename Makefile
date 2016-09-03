all: build

build:
	cd ./sqlparser && go tool yacc -o ./yacc.go ./yacc.y && gofmt -w ./yacc.go
	@bash genver.sh
	go build -o ./bin/saashard ./cmd/saashard

run: build
	./bin/saashard -config conf/ss.yaml

dev: build
	./bin/saashard -config conf/dev.yaml

test:
	go test github.com/berkaroad/saashard/sqlparser

clean:
	@rm -rf bin
	@rm -f ./sqlparse/y.output
