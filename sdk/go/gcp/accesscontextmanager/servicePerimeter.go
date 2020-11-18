// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ServicePerimeter describes a set of GCP resources which can freely import
// and export data amongst themselves, but not export outside of the
// ServicePerimeter. If a request with a source within this ServicePerimeter
// has a target outside of the ServicePerimeter, the request will be blocked.
// Otherwise the request is allowed. There are two types of Service Perimeter
// - Regular and Bridge. Regular Service Perimeters cannot overlap, a single
//   GCP project can only belong to a single regular Service Perimeter. Service
//   Perimeter Bridges can contain only GCP projects as members, a single GCP
//   project may belong to multiple Service Perimeter Bridges.
//
// To get more information about ServicePerimeter, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters)
// * How-to Guides
//     * [Service Perimeter Quickstart](https://cloud.google.com/vpc-service-controls/docs/quickstart)
//
// > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
// you must specify a `billingProject` and set `userProjectOverride` to true
// in the provider configuration. Otherwise the ACM API will return a 403 error.
// Your account must have the `serviceusage.services.use` permission on the
// `billingProject` you defined.
//
// ## Example Usage
// ### Access Context Manager Service Perimeter Basic
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/accesscontextmanager"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := accesscontextmanager.NewAccessPolicy(ctx, "access_policy", &accesscontextmanager.AccessPolicyArgs{
// 			Parent: pulumi.String("organizations/123456789"),
// 			Title:  pulumi.String("my policy"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = accesscontextmanager.NewServicePerimeter(ctx, "service_perimeter", &accesscontextmanager.ServicePerimeterArgs{
// 			Parent: access_policy.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v", "accessPolicies/", name), nil
// 			}).(pulumi.StringOutput),
// 			Status: &accesscontextmanager.ServicePerimeterStatusArgs{
// 				RestrictedServices: pulumi.StringArray{
// 					pulumi.String("storage.googleapis.com"),
// 				},
// 			},
// 			Title: pulumi.String("restrict_storage"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = accesscontextmanager.NewAccessLevel(ctx, "access_level", &accesscontextmanager.AccessLevelArgs{
// 			Basic: &accesscontextmanager.AccessLevelBasicArgs{
// 				Conditions: accesscontextmanager.AccessLevelBasicConditionArray{
// 					&accesscontextmanager.AccessLevelBasicConditionArgs{
// 						DevicePolicy: &accesscontextmanager.AccessLevelBasicConditionDevicePolicyArgs{
// 							OsConstraints: accesscontextmanager.AccessLevelBasicConditionDevicePolicyOsConstraintArray{
// 								&accesscontextmanager.AccessLevelBasicConditionDevicePolicyOsConstraintArgs{
// 									OsType: pulumi.String("DESKTOP_CHROME_OS"),
// 								},
// 							},
// 							RequireScreenLock: pulumi.Bool(false),
// 						},
// 						Regions: pulumi.StringArray{
// 							pulumi.String("CH"),
// 							pulumi.String("IT"),
// 							pulumi.String("US"),
// 						},
// 					},
// 				},
// 			},
// 			Parent: access_policy.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v", "accessPolicies/", name), nil
// 			}).(pulumi.StringOutput),
// 			Title: pulumi.String("chromeos_no_lock"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Access Context Manager Service Perimeter Dry Run
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/accesscontextmanager"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := accesscontextmanager.NewAccessPolicy(ctx, "access_policy", &accesscontextmanager.AccessPolicyArgs{
// 			Parent: pulumi.String("organizations/123456789"),
// 			Title:  pulumi.String("my policy"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = accesscontextmanager.NewServicePerimeter(ctx, "service_perimeter", &accesscontextmanager.ServicePerimeterArgs{
// 			Parent: access_policy.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v", "accessPolicies/", name), nil
// 			}).(pulumi.StringOutput),
// 			Spec: &accesscontextmanager.ServicePerimeterSpecArgs{
// 				RestrictedServices: pulumi.StringArray{
// 					pulumi.String("storage.googleapis.com"),
// 				},
// 			},
// 			Status: &accesscontextmanager.ServicePerimeterStatusArgs{
// 				RestrictedServices: pulumi.StringArray{
// 					pulumi.String("bigquery.googleapis.com"),
// 				},
// 			},
// 			Title:                 pulumi.String("restrict_bigquery_dryrun_storage"),
// 			UseExplicitDryRunSpec: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ServicePerimeter can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:accesscontextmanager/servicePerimeter:ServicePerimeter default {{name}}
// ```
type ServicePerimeter struct {
	pulumi.CustomResourceState

	// Time the AccessPolicy was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the ServicePerimeter and its use. Does not affect
	// behavior.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource name for the ServicePerimeter. The shortName component must
	// begin with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
	Name pulumi.StringOutput `pulumi:"name"`
	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Specifies the type of the Perimeter. There are two types: regular and
	// bridge. Regular Service Perimeter contains resources, access levels,
	// and restricted services. Every resource can be in at most
	// ONE regular Service Perimeter.
	// In addition to being in a regular service perimeter, a resource can also
	// be in zero or more perimeter bridges. A perimeter bridge only contains
	// resources. Cross project operations are permitted if all effected
	// resources share some perimeter (whether bridge or regular). Perimeter
	// Bridge does not contain access levels or services: those are governed
	// entirely by the regular perimeter that resource is in.
	// Perimeter Bridges are typically useful when building more complex
	// topologies with many independent perimeters that need to share some data
	// with a common perimeter, but should not be able to share data among
	// themselves.
	// Default value is `PERIMETER_TYPE_REGULAR`.
	// Possible values are `PERIMETER_TYPE_REGULAR` and `PERIMETER_TYPE_BRIDGE`.
	PerimeterType pulumi.StringPtrOutput `pulumi:"perimeterType"`
	// Proposed (or dry run) ServicePerimeter configuration.
	// This configuration allows to specify and test ServicePerimeter configuration
	// without enforcing actual access restrictions. Only allowed to be set when
	// the `useExplicitDryRunSpec` flag is set.
	// Structure is documented below.
	Spec ServicePerimeterSpecPtrOutput `pulumi:"spec"`
	// ServicePerimeter configuration. Specifies sets of resources,
	// restricted services and access levels that determine
	// perimeter content and boundaries.
	// Structure is documented below.
	Status ServicePerimeterStatusPtrOutput `pulumi:"status"`
	// Human readable title. Must be unique within the Policy.
	Title pulumi.StringOutput `pulumi:"title"`
	// Time the AccessPolicy was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
	// for all Service Perimeters, and that spec is identical to the status for those
	// Service Perimeters. When this flag is set, it inhibits the generation of the
	// implicit spec, thereby allowing the user to explicitly provide a
	// configuration ("spec") to use in a dry-run version of the Service Perimeter.
	// This allows the user to test changes to the enforced config ("status") without
	// actually enforcing them. This testing is done through analyzing the differences
	// between currently enforced and suggested restrictions. useExplicitDryRunSpec must
	// bet set to True if any of the fields in the spec are set to non-default values.
	UseExplicitDryRunSpec pulumi.BoolPtrOutput `pulumi:"useExplicitDryRunSpec"`
}

// NewServicePerimeter registers a new resource with the given unique name, arguments, and options.
func NewServicePerimeter(ctx *pulumi.Context,
	name string, args *ServicePerimeterArgs, opts ...pulumi.ResourceOption) (*ServicePerimeter, error) {
	if args == nil || args.Parent == nil {
		return nil, errors.New("missing required argument 'Parent'")
	}
	if args == nil || args.Title == nil {
		return nil, errors.New("missing required argument 'Title'")
	}
	if args == nil {
		args = &ServicePerimeterArgs{}
	}
	var resource ServicePerimeter
	err := ctx.RegisterResource("gcp:accesscontextmanager/servicePerimeter:ServicePerimeter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePerimeter gets an existing ServicePerimeter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePerimeter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePerimeterState, opts ...pulumi.ResourceOption) (*ServicePerimeter, error) {
	var resource ServicePerimeter
	err := ctx.ReadResource("gcp:accesscontextmanager/servicePerimeter:ServicePerimeter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePerimeter resources.
type servicePerimeterState struct {
	// Time the AccessPolicy was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// Description of the ServicePerimeter and its use. Does not affect
	// behavior.
	Description *string `pulumi:"description"`
	// Resource name for the ServicePerimeter. The shortName component must
	// begin with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
	Name *string `pulumi:"name"`
	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent *string `pulumi:"parent"`
	// Specifies the type of the Perimeter. There are two types: regular and
	// bridge. Regular Service Perimeter contains resources, access levels,
	// and restricted services. Every resource can be in at most
	// ONE regular Service Perimeter.
	// In addition to being in a regular service perimeter, a resource can also
	// be in zero or more perimeter bridges. A perimeter bridge only contains
	// resources. Cross project operations are permitted if all effected
	// resources share some perimeter (whether bridge or regular). Perimeter
	// Bridge does not contain access levels or services: those are governed
	// entirely by the regular perimeter that resource is in.
	// Perimeter Bridges are typically useful when building more complex
	// topologies with many independent perimeters that need to share some data
	// with a common perimeter, but should not be able to share data among
	// themselves.
	// Default value is `PERIMETER_TYPE_REGULAR`.
	// Possible values are `PERIMETER_TYPE_REGULAR` and `PERIMETER_TYPE_BRIDGE`.
	PerimeterType *string `pulumi:"perimeterType"`
	// Proposed (or dry run) ServicePerimeter configuration.
	// This configuration allows to specify and test ServicePerimeter configuration
	// without enforcing actual access restrictions. Only allowed to be set when
	// the `useExplicitDryRunSpec` flag is set.
	// Structure is documented below.
	Spec *ServicePerimeterSpec `pulumi:"spec"`
	// ServicePerimeter configuration. Specifies sets of resources,
	// restricted services and access levels that determine
	// perimeter content and boundaries.
	// Structure is documented below.
	Status *ServicePerimeterStatus `pulumi:"status"`
	// Human readable title. Must be unique within the Policy.
	Title *string `pulumi:"title"`
	// Time the AccessPolicy was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
	// Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
	// for all Service Perimeters, and that spec is identical to the status for those
	// Service Perimeters. When this flag is set, it inhibits the generation of the
	// implicit spec, thereby allowing the user to explicitly provide a
	// configuration ("spec") to use in a dry-run version of the Service Perimeter.
	// This allows the user to test changes to the enforced config ("status") without
	// actually enforcing them. This testing is done through analyzing the differences
	// between currently enforced and suggested restrictions. useExplicitDryRunSpec must
	// bet set to True if any of the fields in the spec are set to non-default values.
	UseExplicitDryRunSpec *bool `pulumi:"useExplicitDryRunSpec"`
}

type ServicePerimeterState struct {
	// Time the AccessPolicy was created in UTC.
	CreateTime pulumi.StringPtrInput
	// Description of the ServicePerimeter and its use. Does not affect
	// behavior.
	Description pulumi.StringPtrInput
	// Resource name for the ServicePerimeter. The shortName component must
	// begin with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
	Name pulumi.StringPtrInput
	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringPtrInput
	// Specifies the type of the Perimeter. There are two types: regular and
	// bridge. Regular Service Perimeter contains resources, access levels,
	// and restricted services. Every resource can be in at most
	// ONE regular Service Perimeter.
	// In addition to being in a regular service perimeter, a resource can also
	// be in zero or more perimeter bridges. A perimeter bridge only contains
	// resources. Cross project operations are permitted if all effected
	// resources share some perimeter (whether bridge or regular). Perimeter
	// Bridge does not contain access levels or services: those are governed
	// entirely by the regular perimeter that resource is in.
	// Perimeter Bridges are typically useful when building more complex
	// topologies with many independent perimeters that need to share some data
	// with a common perimeter, but should not be able to share data among
	// themselves.
	// Default value is `PERIMETER_TYPE_REGULAR`.
	// Possible values are `PERIMETER_TYPE_REGULAR` and `PERIMETER_TYPE_BRIDGE`.
	PerimeterType pulumi.StringPtrInput
	// Proposed (or dry run) ServicePerimeter configuration.
	// This configuration allows to specify and test ServicePerimeter configuration
	// without enforcing actual access restrictions. Only allowed to be set when
	// the `useExplicitDryRunSpec` flag is set.
	// Structure is documented below.
	Spec ServicePerimeterSpecPtrInput
	// ServicePerimeter configuration. Specifies sets of resources,
	// restricted services and access levels that determine
	// perimeter content and boundaries.
	// Structure is documented below.
	Status ServicePerimeterStatusPtrInput
	// Human readable title. Must be unique within the Policy.
	Title pulumi.StringPtrInput
	// Time the AccessPolicy was updated in UTC.
	UpdateTime pulumi.StringPtrInput
	// Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
	// for all Service Perimeters, and that spec is identical to the status for those
	// Service Perimeters. When this flag is set, it inhibits the generation of the
	// implicit spec, thereby allowing the user to explicitly provide a
	// configuration ("spec") to use in a dry-run version of the Service Perimeter.
	// This allows the user to test changes to the enforced config ("status") without
	// actually enforcing them. This testing is done through analyzing the differences
	// between currently enforced and suggested restrictions. useExplicitDryRunSpec must
	// bet set to True if any of the fields in the spec are set to non-default values.
	UseExplicitDryRunSpec pulumi.BoolPtrInput
}

func (ServicePerimeterState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePerimeterState)(nil)).Elem()
}

type servicePerimeterArgs struct {
	// Description of the ServicePerimeter and its use. Does not affect
	// behavior.
	Description *string `pulumi:"description"`
	// Resource name for the ServicePerimeter. The shortName component must
	// begin with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
	Name *string `pulumi:"name"`
	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent string `pulumi:"parent"`
	// Specifies the type of the Perimeter. There are two types: regular and
	// bridge. Regular Service Perimeter contains resources, access levels,
	// and restricted services. Every resource can be in at most
	// ONE regular Service Perimeter.
	// In addition to being in a regular service perimeter, a resource can also
	// be in zero or more perimeter bridges. A perimeter bridge only contains
	// resources. Cross project operations are permitted if all effected
	// resources share some perimeter (whether bridge or regular). Perimeter
	// Bridge does not contain access levels or services: those are governed
	// entirely by the regular perimeter that resource is in.
	// Perimeter Bridges are typically useful when building more complex
	// topologies with many independent perimeters that need to share some data
	// with a common perimeter, but should not be able to share data among
	// themselves.
	// Default value is `PERIMETER_TYPE_REGULAR`.
	// Possible values are `PERIMETER_TYPE_REGULAR` and `PERIMETER_TYPE_BRIDGE`.
	PerimeterType *string `pulumi:"perimeterType"`
	// Proposed (or dry run) ServicePerimeter configuration.
	// This configuration allows to specify and test ServicePerimeter configuration
	// without enforcing actual access restrictions. Only allowed to be set when
	// the `useExplicitDryRunSpec` flag is set.
	// Structure is documented below.
	Spec *ServicePerimeterSpec `pulumi:"spec"`
	// ServicePerimeter configuration. Specifies sets of resources,
	// restricted services and access levels that determine
	// perimeter content and boundaries.
	// Structure is documented below.
	Status *ServicePerimeterStatus `pulumi:"status"`
	// Human readable title. Must be unique within the Policy.
	Title string `pulumi:"title"`
	// Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
	// for all Service Perimeters, and that spec is identical to the status for those
	// Service Perimeters. When this flag is set, it inhibits the generation of the
	// implicit spec, thereby allowing the user to explicitly provide a
	// configuration ("spec") to use in a dry-run version of the Service Perimeter.
	// This allows the user to test changes to the enforced config ("status") without
	// actually enforcing them. This testing is done through analyzing the differences
	// between currently enforced and suggested restrictions. useExplicitDryRunSpec must
	// bet set to True if any of the fields in the spec are set to non-default values.
	UseExplicitDryRunSpec *bool `pulumi:"useExplicitDryRunSpec"`
}

// The set of arguments for constructing a ServicePerimeter resource.
type ServicePerimeterArgs struct {
	// Description of the ServicePerimeter and its use. Does not affect
	// behavior.
	Description pulumi.StringPtrInput
	// Resource name for the ServicePerimeter. The shortName component must
	// begin with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
	Name pulumi.StringPtrInput
	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringInput
	// Specifies the type of the Perimeter. There are two types: regular and
	// bridge. Regular Service Perimeter contains resources, access levels,
	// and restricted services. Every resource can be in at most
	// ONE regular Service Perimeter.
	// In addition to being in a regular service perimeter, a resource can also
	// be in zero or more perimeter bridges. A perimeter bridge only contains
	// resources. Cross project operations are permitted if all effected
	// resources share some perimeter (whether bridge or regular). Perimeter
	// Bridge does not contain access levels or services: those are governed
	// entirely by the regular perimeter that resource is in.
	// Perimeter Bridges are typically useful when building more complex
	// topologies with many independent perimeters that need to share some data
	// with a common perimeter, but should not be able to share data among
	// themselves.
	// Default value is `PERIMETER_TYPE_REGULAR`.
	// Possible values are `PERIMETER_TYPE_REGULAR` and `PERIMETER_TYPE_BRIDGE`.
	PerimeterType pulumi.StringPtrInput
	// Proposed (or dry run) ServicePerimeter configuration.
	// This configuration allows to specify and test ServicePerimeter configuration
	// without enforcing actual access restrictions. Only allowed to be set when
	// the `useExplicitDryRunSpec` flag is set.
	// Structure is documented below.
	Spec ServicePerimeterSpecPtrInput
	// ServicePerimeter configuration. Specifies sets of resources,
	// restricted services and access levels that determine
	// perimeter content and boundaries.
	// Structure is documented below.
	Status ServicePerimeterStatusPtrInput
	// Human readable title. Must be unique within the Policy.
	Title pulumi.StringInput
	// Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
	// for all Service Perimeters, and that spec is identical to the status for those
	// Service Perimeters. When this flag is set, it inhibits the generation of the
	// implicit spec, thereby allowing the user to explicitly provide a
	// configuration ("spec") to use in a dry-run version of the Service Perimeter.
	// This allows the user to test changes to the enforced config ("status") without
	// actually enforcing them. This testing is done through analyzing the differences
	// between currently enforced and suggested restrictions. useExplicitDryRunSpec must
	// bet set to True if any of the fields in the spec are set to non-default values.
	UseExplicitDryRunSpec pulumi.BoolPtrInput
}

func (ServicePerimeterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePerimeterArgs)(nil)).Elem()
}

type ServicePerimeterInput interface {
	pulumi.Input

	ToServicePerimeterOutput() ServicePerimeterOutput
	ToServicePerimeterOutputWithContext(ctx context.Context) ServicePerimeterOutput
}

func (ServicePerimeter) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePerimeter)(nil)).Elem()
}

func (i ServicePerimeter) ToServicePerimeterOutput() ServicePerimeterOutput {
	return i.ToServicePerimeterOutputWithContext(context.Background())
}

func (i ServicePerimeter) ToServicePerimeterOutputWithContext(ctx context.Context) ServicePerimeterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePerimeterOutput)
}

type ServicePerimeterOutput struct {
	*pulumi.OutputState
}

func (ServicePerimeterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePerimeterOutput)(nil)).Elem()
}

func (o ServicePerimeterOutput) ToServicePerimeterOutput() ServicePerimeterOutput {
	return o
}

func (o ServicePerimeterOutput) ToServicePerimeterOutputWithContext(ctx context.Context) ServicePerimeterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServicePerimeterOutput{})
}
