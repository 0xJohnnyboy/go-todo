build:
	go build -o bin/todo .

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/todo-linux .

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/todo.exe .

clean:
	rm -f bin/*
