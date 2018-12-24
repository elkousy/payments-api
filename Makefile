
.PHONY: install build test run test-cover

.PHONY: build
build:
	@CGO_ENABLED=0 GOOS=linux go build -o ./app -a -ldflags '-s' -installsuffix cgo main.go

.PHONY: install
install:
	@go get -u golang.org/x/lint/golint
	@go get -u github.com/golang/dep/cmd/dep
	@dep ensure -v

.PHONY: run
run: 
	@go run main.go

.PHONY: update-mocks
update-mocks:
	@go get github.com/vektra/mockery/.../
	@go list -f '{{.Dir}}' ./... | grep -v "payments-api$$" | xargs -n1 ${GOPATH}/bin/mockery -inpkg -case "underscore" -all -note "NOTE: run 'make update-mocks' from payments-api top folder to update this file and generate new ones." -dir || true

.PHONY: unit-test
unit-test:
	go test -v ./...

.PHONY:test-cover
test-cover:
	@go test `go list ./... | grep -v /vendor/` -cover -coverprofile=cover.out
	@go tool cover -html=cover.out

.PHONY: integration-tests
integration-tests: 
	cd newman && newman run payments-api.integration-test.json -e Dev.postman_environment.json

# .PHONY: test-gherkin
# test-gherkin:
# 	go get -u github.com/DATA-DOG/godog/cmd/godog
# 	cd features && godog .

fmt:
	@go fmt github.com/elkousy/payments-api/...

.PHONY: docker-compose-build
docker-compose-build:
	make build & \
	wait && \
	docker-compose build

.PHONY: docker-compose-up
docker-compose-up:
	make docker-compose-up-dep && \
   	NO_PROXY=* docker-compose up payments-api newman

.PHONY: docker-compose-up-dep
docker-compose-up-dep:
	NO_PROXY=* docker-compose up -d db