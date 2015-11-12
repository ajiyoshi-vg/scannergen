DEMOS := demo/hook.scanner.go demo/issue.scanner.go demo/user.scanner.go

all: demo

build:
	go build

clean:
	-rm scannergen

clear_all:
	-rm $(DEMOS) sqlgen

demo: clean build $(DEMOS)

%.scanner.go : %.go
	./scannergen -file $< -pkg demo -o $@

