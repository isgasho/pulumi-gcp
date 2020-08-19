module github.com/pulumi/pulumi-gcp/provider/v3

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.11.0
	github.com/hashicorp/terraform-provider-google-beta v0.0.0-20200309221941-5fc1579be217
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.5.0
	github.com/pulumi/pulumi/sdk/v2 v2.7.1
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/terraform-provider-google-beta => github.com/pulumi/terraform-provider-google-beta v0.0.0-20200817191203-097ec6fba670
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	google.golang.org/api => google.golang.org/api v0.29.0
)
