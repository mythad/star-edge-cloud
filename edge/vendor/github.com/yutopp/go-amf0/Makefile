all: pre test

pre: fmt lint vet

fmt:
	go fmt ./...

lint:
	golint $$(go list ./... | grep -v /vendor/)

vet:
	go vet $$(go list ./... | grep -v /vendor/)

test:
	go test -v -race -cover ./...

dep-init:
	dep ensure

dep-update:
	dep ensure -update

.PHONY: all pre fmt vet lint test dep-init dep-update
