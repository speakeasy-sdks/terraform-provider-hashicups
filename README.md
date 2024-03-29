# Terraform Provider Hashicups

This repository is an example of generating an example provider called [hashicups](https://github.com/hashicorp/terraform-provider-hashicups-pf) from the OpenAPI specification in `hashicups.yaml`, with configuration in `gen.yaml`. No other inputs are required.

It works by using the [Speakeasy](https://speakeasyapi.dev) platform and associated CLI Tooling to generate code. This tooling parses an OpenAPI specification annotated with extensions that help inform it about terraform resources, and how to invoke API operations to create, update, read and destroy those resources. For information on terraform annotations that you can add to your OpenAPI spec please see [here](https://speakeasyapi.dev/docs/using-speakeasy/create-terraform/intro/). 

This provider is for demonstration purposes only. Do not publish this to the terraform registry.

<!-- Start SDK Installation [installation] -->
## SDK Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    hashicups = {
      source  = "speakeasy/hashicups"
      version = "0.14.1"
    }
  }
}

provider "hashicups" {
  # Configuration options
}
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Testing the provider locally

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

### Example

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations


<!-- End Available Resources and Operations [operations] -->



<!-- Placeholder for Future Speakeasy SDK Sections -->



### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
