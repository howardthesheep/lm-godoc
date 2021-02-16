
GOFILES := $(shell find . -type f -name "*.go")

lm-godoc:
	go build -o godoc.out main.go log.go

all: lm-godoc

build: lm-godoc

clean:
	rm -f *.out

run:
	go run lm-godoc
