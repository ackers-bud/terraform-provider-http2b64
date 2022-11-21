# Terraform Provider: http2b64

This provider requests an URL and will take whatever the URL serves saving it as a base64 encode string.  Written to be used 
with https://github.com/Ackers-BUD/terraform-provider-zip2b64 were the resource is a zip file. 

### Example Usage 

```
resource "http-b64" "provisioningzipfile" {
    provider = http2b64
    url      = "http://<url-to-a-zipfile>.zip"
}

resource "zip2b64" "extractedfile" {
    provider = zip2b64
    base64file = http-b64.provisioningzipfile.response_body_base64
    filename = "/ca.crt"
}
```

## http-b64 computed values

### "id"
Set as the URL of the Request. 

### "status_code" 
The Status code of the http request.  Useful if determining if file isn't found.  ALL request bodies regardless of http code are converted to base64. 

### "response_body_base64"
A base64 representation of the the responce. 


See example usage below

## Requirements

* [Terraform](https://www.terraform.io/downloads) (>= 0.12)
* [Go](https://go.dev/doc/install) (1.18)
* [golangci-lint](https://golangci-lint.run/usage/install/#local-installation) (optional)

## Development

### Building

1. `git clone` this repository and `cd` into its directory
2. `make` will trigger the Golang build

The provided Makefile defines additional commands generally useful during development,
like for running tests, generating documentation, code formatting and linting.
Taking a look at it's content is recommended.

### Testing

In order to test the provider, you can run

* `make test` to run provider tests
* `make testacc` to run provider acceptance tests

It's important to note that acceptance tests (`testacc`) will actually spawn
`terraform` and the provider. Read more about they work on the
[official page](https://www.terraform.io/plugin/sdkv2/testing/acceptance-tests).

### Generating documentation

This provider uses [terraform-plugin-docs](https://github.com/hashicorp/terraform-plugin-docs/)
to generate documentation and store it in the `docs/` directory.
Once a release is cut, the Terraform Registry will download the documentation from `docs/`
and associate it with the release version. Read more about how this works on the
[official page](https://www.terraform.io/registry/providers/docs).

Use `make generate` to ensure the documentation is regenerated with any changes.

### Using a development build

If [running tests and acceptance tests](#testing) isn't enough, it's possible to set up a local terraform configuration
to use a development builds of the provider. This can be achieved by leveraging the Terraform CLI
[configuration file development overrides](https://www.terraform.io/cli/config/config-file#development-overrides-for-provider-developers).

First, use `make install` to place a fresh development build of the provider in your
[`${GOBIN}`](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies)
(defaults to `${GOPATH}/bin` or `${HOME}/go/bin` if `${GOPATH}` is not set). Repeat
this every time you make changes to the provider locally.

Then, setup your environment following [these instructions](https://www.terraform.io/plugin/debugging#terraform-cli-development-overrides)
to make your local terraform use your local build.

### Testing GitHub Actions

This project uses [GitHub Actions](https://docs.github.com/en/actions/automating-builds-and-tests) to realize its CI.

Sometimes it might be helpful to locally reproduce the behaviour of those actions,
and for this we use [act](https://github.com/nektos/act). Once installed, you can _simulate_ the actions executed
when opening a PR with:

```shell
# List of workflows for the 'pull_request' action
$ act -l pull_request

# Execute the workflows associated with the `pull_request' action 
$ act pull_request
```

## Releasing

The release process is automated via GitHub Actions, and it's defined in the Workflow
[release.yml](./.github/workflows/release.yml).

Each release is cut by pushing a [semantically versioned](https://semver.org/) tag to the default branch.

## License

[Mozilla Public License v2.0](./LICENSE)
