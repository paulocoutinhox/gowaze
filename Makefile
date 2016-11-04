GOFMT=gofmt -w
GODEPS=go get -u
GOTEST=go test

format:
	${GOFMT} address/address.go
	${GOFMT} address/address_test.go
	${GOFMT} domain/WazeAddress.go
	${GOFMT} domain/WazeRouteDirections.go
	${GOFMT} route/route.go
	${GOFMT} route/route_test.go

test:
	${GOTEST} ./address
	${GOTEST} ./route

deps:
	${GODEPS} github.com/parnurzeal/gorequest