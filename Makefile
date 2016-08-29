all: build

build:
	@bash genver.sh
	@go build -o ./bin/saashard ./cmd/saashard
run: build
	./bin/saashard -config conf/ss.yaml
clean:
	@rm -rf bin
