all: build/day1

.PHONY: clean
clean:
	rm -Rf build

build/day1:
	go build -o build/day1 cmd/day1/main.go