#### Example Golang Project Layout
This is a simple demo, cmd line (and GUI) app shows a free TCP port on machine.

See: https://github.com/golang-standards/project-layout

How to run
```
cd .
go run cmd/example-go-cmd/main.go --gui
```
or

Doesn't work?
```
go get github.com/zvodd/example-go-command
example-go-command
```

How to setup go module
```
go mod init github.com/$username/$reponame
go get $dependant-packages
```
