create:
	mkdir mymodule

cd:
	cd mymodule

goinit:
	go mod init github.com/Ruslan0056/go-level2/mymodule

gofile:
	echo "package geekbrains_go_module import _ "github.com/valyala/fasthttp"" > main.go

fmt:
	go fmt main.go

tidy:
	go mod tidy

commit:
	git commit -a -m 'initial commit.'
	git push
	git tag -a v1.0.0 -m 'release'
	git push tegs

createv2:
	mkdir v2
	cd v2
	cp main.go ./v2/
	echo "package geekbrains_go_module 

	import (
		github.com/valyala/fasthttp
		"fmt"
	)

	func main() {
		fmt.Println("Hello!")
	}" > main.go
	go fmt main.go
	cp go.mod ./v2/
	go mod edit -module github.com/Ruslan0056/go-level2/mymodule/v2 v2/go.mod
	git commit -a -m 'added version 2.'
	git push
	git tag -a v2.0.0 -m 'version 2.'
	git push --tags
	go mod tidy




