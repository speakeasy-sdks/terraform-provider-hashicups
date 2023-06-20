.PHONY: *

speakeasy:
	speakeasy generate sdk --lang terraform -o . -s hashicups.yaml
