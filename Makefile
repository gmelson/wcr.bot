all:
	GOPATH=`pwd` go install nl/
	GOPATH=`pwd` go build src/main.go


