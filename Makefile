DEMOS := demo/hook.scanner.go demo/issue.scanner.go demo/user.scanner.go

all: demo

build:
	go build

clean:
	-rm sqlgen

clear_all:
	-rm $(DEMOS) sqlgen

demo: clean build $(DEMOS)

%.scanner.go : %.go
	./sqlgen -file $< -pkg demo -o $@

