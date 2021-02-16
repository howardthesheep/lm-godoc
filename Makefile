lm-godoc:
	go build -o lm-godoc.out main.go

all: lm-godoc

build: lm-godoc

clean:
	rm -f lm-godoc

run:
	go run lm-godoc
