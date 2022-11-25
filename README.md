# Cloudflare Terraform Provider

## Developing locally

### Initial setup

Run `dev clone terraform-provider-cloudflare` and place the following snippet in `~/.terraformrc`:
```
provider_installation {
  dev_overrides {
    "cloudflare/cloudflare" = "/Users/<your_name>/src/github.com/Shopify/terraform-provider-cloudflare"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

### Testing your changes

After making your changes, run `make build-dev` and run `terraform init && terraform plan` in the terraform project using the module.

## Quickstarts

- [Getting started with Cloudflare and Terraform](https://developers.cloudflare.com/terraform/installing)
- [Developing the provider](contributing/development.md)

## Documentation

Full, comprehensive documentation is available on the [Terraform Registry](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs). [API documentation](https://api.cloudflare.com) and [Developer documentation](https://developers.cloudflare.com) is also available
for non-Terraform or service specific information.

## Migrating to Terraform from using the Dashboard

Do you have an existing Cloudflare account (or many!) that you'd like to transition
to be managed via Terraform? Check out [cf-terraforming](https://github.com/cloudflare/cf-terraforming)
which is a tool Cloudflare has built to help dump the existing resources and
import them into Terraform.

## Contributing

To contribute, please read the [contribution guidelines](contributing/README.md).

## Feedback

If you would like to provide feedback (not a bug or feature request) on the Cloudflare Terraform provider, you're welcome to via [this form](https://forms.gle/6ofUoRY2QmPMSqoR6).
