---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "hashicups_order Resource - terraform-provider-hashicups"
subcategory: ""
description: |-
  Order Resource
---

# hashicups_order (Resource)

Order Resource

## Example Usage

```terraform
resource "hashicups_order" "my_order" {
  image  = "...my_image..."
  name   = "Terrence Rau"
  price  = 85.8
  teaser = "...my_teaser..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `image` (String) URI for an image of the coffee.
- `name` (String) Product name of the coffee.
- `price` (Number) Suggested cost of the coffee.
- `teaser` (String) Fun tagline for the coffee.

### Optional

- `description` (String) Product description of the coffee.

### Read-Only

- `id` (Number) Order ID

