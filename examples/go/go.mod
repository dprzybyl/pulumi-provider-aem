module github.com/dprzybyl/pulumi-provider-aem/examples

go 1.20

replace github.com/dprzybyl/pulumi-provider-aem/sdk => ../../sdk

require (
	github.com/dprzybyl/pulumi-provider-aem/sdk v0.0.0-00010101000000-000000000000
	github.com/pulumi/pulumi/sdk/v3 v3.30.0
)