---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "hashicups Provider"
subcategory: ""
description: |-
  Hashicups: Example Hashicups through Speakeasy
---

# hashicups Provider

Hashicups: Example Hashicups through Speakeasy

## Example Usage

```terraform
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

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `api_key` (String, Sensitive)
- `server_url` (String) Server URL (defaults to https://example.com)
