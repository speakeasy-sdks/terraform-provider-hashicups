.PHONY: all docs
all: speakeasy docs

speakeasy:
	speakeasy generate sdk --lang terraform -o . -s hashicups.yaml

docs:
	go generate ./...
