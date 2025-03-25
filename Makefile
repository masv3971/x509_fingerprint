.PHONY : gosec staticcheck build vscode

LDFLAGS                 := -ldflags "-w -s --extldflags '-static'"

gosec:
	$(info Run gosec)
	gosec -color -nosec -tests ./...

staticcheck:
	$(info Run staticcheck)
	staticcheck ./...

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/x509_fingerprint ${LDFLAGS} ./main.go

vscode:
	$(info Install go packages)
	go install golang.org/x/tools/cmd/deadcode@latest && \
	go install github.com/securego/gosec/v2/cmd/gosec@latest && \
	go install honnef.co/go/tools/cmd/staticcheck@latest