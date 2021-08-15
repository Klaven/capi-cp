

##@ Build

build:
	go build -o bin/capi-cp main.go

clean:
	go clean -i -x -r