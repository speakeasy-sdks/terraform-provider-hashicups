terraform {
  required_providers {
    hashicups = {
      source  = "speakeasy/hashicups"
      version = "0.0.1"
    }
  }
}

provider "hashicups" {
  server_url = "http://localhost:35123/anything"
  api_key    = "example_api_key"
}