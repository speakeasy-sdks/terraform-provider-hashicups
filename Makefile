.PHONY: all
all: docs

speakeasy:
	speakeasy generate sdk --lang terraform -o . -s hashicups.yaml

docs: speakeasy
	go generate ./...

