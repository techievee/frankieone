all:    ## clean, format, build and unit test
	make clean-all
	make build
	make test

clean-all:  ## remove all generated artifacts and clean all build artifacts
	go clean -i ./...
	rm -fr rpc

test:
	go test -v ./... -short

test-it:
	go test -v ./...

test-bench: ## run benchmark tests
	go test -bench ./...

check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	GO111MODULE=on go mod vendor  && GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

serve-swagger: check-swagger
	swagger serve -F=swagger swagger.yaml

