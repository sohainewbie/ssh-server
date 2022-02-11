.PHONY : format ssh-osx ssh-linux create

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

ssh-osx: main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@

ssh-linux: main.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@
