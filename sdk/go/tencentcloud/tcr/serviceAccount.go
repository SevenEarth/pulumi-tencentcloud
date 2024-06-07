// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcr

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a tcr service account.
//
// ## Example Usage
//
// ### Create custom account with specified duration days
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tcr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleInstance, err := Tcr.NewInstance(ctx, "exampleInstance", &Tcr.InstanceArgs{
//				InstanceType: pulumi.String("basic"),
//				DeleteBucket: pulumi.Bool(true),
//				Tags: pulumi.Map{
//					"createdBy": pulumi.Any("terraform"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleNamespace, err := Tcr.NewNamespace(ctx, "exampleNamespace", &Tcr.NamespaceArgs{
//				InstanceId:   exampleInstance.ID(),
//				IsPublic:     pulumi.Bool(true),
//				IsAutoScan:   pulumi.Bool(true),
//				IsPreventVul: pulumi.Bool(true),
//				Severity:     pulumi.String("medium"),
//				CveWhitelistItems: tcr.NamespaceCveWhitelistItemArray{
//					&tcr.NamespaceCveWhitelistItemArgs{
//						CveId: pulumi.String("tf_example_cve_id"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Tcr.NewServiceAccount(ctx, "exampleServiceAccount", &Tcr.ServiceAccountArgs{
//				RegistryId: exampleInstance.ID(),
//				Permissions: tcr.ServiceAccountPermissionArray{
//					&tcr.ServiceAccountPermissionArgs{
//						Resource: exampleNamespace.Name,
//						Actions: pulumi.StringArray{
//							pulumi.String("tcr:PushRepository"),
//							pulumi.String("tcr:PullRepository"),
//						},
//					},
//				},
//				Description: pulumi.String("tf example for tcr custom account"),
//				Duration:    pulumi.Int(10),
//				Disable:     pulumi.Bool(false),
//				Tags: pulumi.Map{
//					"createdBy": pulumi.Any("terraform"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### With specified expiration time
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tcr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tcr.NewServiceAccount(ctx, "example", &Tcr.ServiceAccountArgs{
//				RegistryId: pulumi.Any(tencentcloud_tcr_instance.Example.Id),
//				Permissions: tcr.ServiceAccountPermissionArray{
//					&tcr.ServiceAccountPermissionArgs{
//						Resource: pulumi.Any(tencentcloud_tcr_namespace.Example.Name),
//						Actions: pulumi.StringArray{
//							pulumi.String("tcr:PushRepository"),
//							pulumi.String("tcr:PullRepository"),
//						},
//					},
//				},
//				Description: pulumi.String("tf example for tcr custom account"),
//				ExpiresAt:   pulumi.Int(1676897989000),
//				Disable:     pulumi.Bool(false),
//				Tags: pulumi.Map{
//					"createdBy": pulumi.Any("terraform"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// tcr service_account can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Tcr/serviceAccount:ServiceAccount service_account registry_id#account_name
// ```
type ServiceAccount struct {
	pulumi.CustomResourceState

	// Service account description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// whether to disable Service accounts.
	Disable pulumi.BoolPtrOutput `pulumi:"disable"`
	// expiration date (unit: day), calculated from the current time, priority is higher than ExpiresAt Service account description.
	Duration pulumi.IntPtrOutput `pulumi:"duration"`
	// Service account expiration time (time stamp, unit: milliseconds).
	ExpiresAt pulumi.IntOutput `pulumi:"expiresAt"`
	// Service account name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password of the service account.
	Password pulumi.StringOutput `pulumi:"password"`
	// strategy list.
	Permissions ServiceAccountPermissionArrayOutput `pulumi:"permissions"`
	// instance id.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// Tag description list.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewServiceAccount registers a new resource with the given unique name, arguments, and options.
func NewServiceAccount(ctx *pulumi.Context,
	name string, args *ServiceAccountArgs, opts ...pulumi.ResourceOption) (*ServiceAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.RegistryId == nil {
		return nil, errors.New("invalid value for required argument 'RegistryId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceAccount
	err := ctx.RegisterResource("tencentcloud:Tcr/serviceAccount:ServiceAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceAccount gets an existing ServiceAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceAccountState, opts ...pulumi.ResourceOption) (*ServiceAccount, error) {
	var resource ServiceAccount
	err := ctx.ReadResource("tencentcloud:Tcr/serviceAccount:ServiceAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceAccount resources.
type serviceAccountState struct {
	// Service account description.
	Description *string `pulumi:"description"`
	// whether to disable Service accounts.
	Disable *bool `pulumi:"disable"`
	// expiration date (unit: day), calculated from the current time, priority is higher than ExpiresAt Service account description.
	Duration *int `pulumi:"duration"`
	// Service account expiration time (time stamp, unit: milliseconds).
	ExpiresAt *int `pulumi:"expiresAt"`
	// Service account name.
	Name *string `pulumi:"name"`
	// Password of the service account.
	Password *string `pulumi:"password"`
	// strategy list.
	Permissions []ServiceAccountPermission `pulumi:"permissions"`
	// instance id.
	RegistryId *string `pulumi:"registryId"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
}

type ServiceAccountState struct {
	// Service account description.
	Description pulumi.StringPtrInput
	// whether to disable Service accounts.
	Disable pulumi.BoolPtrInput
	// expiration date (unit: day), calculated from the current time, priority is higher than ExpiresAt Service account description.
	Duration pulumi.IntPtrInput
	// Service account expiration time (time stamp, unit: milliseconds).
	ExpiresAt pulumi.IntPtrInput
	// Service account name.
	Name pulumi.StringPtrInput
	// Password of the service account.
	Password pulumi.StringPtrInput
	// strategy list.
	Permissions ServiceAccountPermissionArrayInput
	// instance id.
	RegistryId pulumi.StringPtrInput
	// Tag description list.
	Tags pulumi.MapInput
}

func (ServiceAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountState)(nil)).Elem()
}

type serviceAccountArgs struct {
	// Service account description.
	Description *string `pulumi:"description"`
	// whether to disable Service accounts.
	Disable *bool `pulumi:"disable"`
	// expiration date (unit: day), calculated from the current time, priority is higher than ExpiresAt Service account description.
	Duration *int `pulumi:"duration"`
	// Service account expiration time (time stamp, unit: milliseconds).
	ExpiresAt *int `pulumi:"expiresAt"`
	// Service account name.
	Name *string `pulumi:"name"`
	// strategy list.
	Permissions []ServiceAccountPermission `pulumi:"permissions"`
	// instance id.
	RegistryId string `pulumi:"registryId"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a ServiceAccount resource.
type ServiceAccountArgs struct {
	// Service account description.
	Description pulumi.StringPtrInput
	// whether to disable Service accounts.
	Disable pulumi.BoolPtrInput
	// expiration date (unit: day), calculated from the current time, priority is higher than ExpiresAt Service account description.
	Duration pulumi.IntPtrInput
	// Service account expiration time (time stamp, unit: milliseconds).
	ExpiresAt pulumi.IntPtrInput
	// Service account name.
	Name pulumi.StringPtrInput
	// strategy list.
	Permissions ServiceAccountPermissionArrayInput
	// instance id.
	RegistryId pulumi.StringInput
	// Tag description list.
	Tags pulumi.MapInput
}

func (ServiceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountArgs)(nil)).Elem()
}

type ServiceAccountInput interface {
	pulumi.Input

	ToServiceAccountOutput() ServiceAccountOutput
	ToServiceAccountOutputWithContext(ctx context.Context) ServiceAccountOutput
}

func (*ServiceAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccount)(nil)).Elem()
}

func (i *ServiceAccount) ToServiceAccountOutput() ServiceAccountOutput {
	return i.ToServiceAccountOutputWithContext(context.Background())
}

func (i *ServiceAccount) ToServiceAccountOutputWithContext(ctx context.Context) ServiceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountOutput)
}

// ServiceAccountArrayInput is an input type that accepts ServiceAccountArray and ServiceAccountArrayOutput values.
// You can construct a concrete instance of `ServiceAccountArrayInput` via:
//
//	ServiceAccountArray{ ServiceAccountArgs{...} }
type ServiceAccountArrayInput interface {
	pulumi.Input

	ToServiceAccountArrayOutput() ServiceAccountArrayOutput
	ToServiceAccountArrayOutputWithContext(context.Context) ServiceAccountArrayOutput
}

type ServiceAccountArray []ServiceAccountInput

func (ServiceAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceAccount)(nil)).Elem()
}

func (i ServiceAccountArray) ToServiceAccountArrayOutput() ServiceAccountArrayOutput {
	return i.ToServiceAccountArrayOutputWithContext(context.Background())
}

func (i ServiceAccountArray) ToServiceAccountArrayOutputWithContext(ctx context.Context) ServiceAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountArrayOutput)
}

// ServiceAccountMapInput is an input type that accepts ServiceAccountMap and ServiceAccountMapOutput values.
// You can construct a concrete instance of `ServiceAccountMapInput` via:
//
//	ServiceAccountMap{ "key": ServiceAccountArgs{...} }
type ServiceAccountMapInput interface {
	pulumi.Input

	ToServiceAccountMapOutput() ServiceAccountMapOutput
	ToServiceAccountMapOutputWithContext(context.Context) ServiceAccountMapOutput
}

type ServiceAccountMap map[string]ServiceAccountInput

func (ServiceAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceAccount)(nil)).Elem()
}

func (i ServiceAccountMap) ToServiceAccountMapOutput() ServiceAccountMapOutput {
	return i.ToServiceAccountMapOutputWithContext(context.Background())
}

func (i ServiceAccountMap) ToServiceAccountMapOutputWithContext(ctx context.Context) ServiceAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountMapOutput)
}

type ServiceAccountOutput struct{ *pulumi.OutputState }

func (ServiceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccount)(nil)).Elem()
}

func (o ServiceAccountOutput) ToServiceAccountOutput() ServiceAccountOutput {
	return o
}

func (o ServiceAccountOutput) ToServiceAccountOutputWithContext(ctx context.Context) ServiceAccountOutput {
	return o
}

// Service account description.
func (o ServiceAccountOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAccount) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// whether to disable Service accounts.
func (o ServiceAccountOutput) Disable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceAccount) pulumi.BoolPtrOutput { return v.Disable }).(pulumi.BoolPtrOutput)
}

// expiration date (unit: day), calculated from the current time, priority is higher than ExpiresAt Service account description.
func (o ServiceAccountOutput) Duration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceAccount) pulumi.IntPtrOutput { return v.Duration }).(pulumi.IntPtrOutput)
}

// Service account expiration time (time stamp, unit: milliseconds).
func (o ServiceAccountOutput) ExpiresAt() pulumi.IntOutput {
	return o.ApplyT(func(v *ServiceAccount) pulumi.IntOutput { return v.ExpiresAt }).(pulumi.IntOutput)
}

// Service account name.
func (o ServiceAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password of the service account.
func (o ServiceAccountOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAccount) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// strategy list.
func (o ServiceAccountOutput) Permissions() ServiceAccountPermissionArrayOutput {
	return o.ApplyT(func(v *ServiceAccount) ServiceAccountPermissionArrayOutput { return v.Permissions }).(ServiceAccountPermissionArrayOutput)
}

// instance id.
func (o ServiceAccountOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAccount) pulumi.StringOutput { return v.RegistryId }).(pulumi.StringOutput)
}

// Tag description list.
func (o ServiceAccountOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *ServiceAccount) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type ServiceAccountArrayOutput struct{ *pulumi.OutputState }

func (ServiceAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceAccount)(nil)).Elem()
}

func (o ServiceAccountArrayOutput) ToServiceAccountArrayOutput() ServiceAccountArrayOutput {
	return o
}

func (o ServiceAccountArrayOutput) ToServiceAccountArrayOutputWithContext(ctx context.Context) ServiceAccountArrayOutput {
	return o
}

func (o ServiceAccountArrayOutput) Index(i pulumi.IntInput) ServiceAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceAccount {
		return vs[0].([]*ServiceAccount)[vs[1].(int)]
	}).(ServiceAccountOutput)
}

type ServiceAccountMapOutput struct{ *pulumi.OutputState }

func (ServiceAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceAccount)(nil)).Elem()
}

func (o ServiceAccountMapOutput) ToServiceAccountMapOutput() ServiceAccountMapOutput {
	return o
}

func (o ServiceAccountMapOutput) ToServiceAccountMapOutputWithContext(ctx context.Context) ServiceAccountMapOutput {
	return o
}

func (o ServiceAccountMapOutput) MapIndex(k pulumi.StringInput) ServiceAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceAccount {
		return vs[0].(map[string]*ServiceAccount)[vs[1].(string)]
	}).(ServiceAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountInput)(nil)).Elem(), &ServiceAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountArrayInput)(nil)).Elem(), ServiceAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountMapInput)(nil)).Elem(), ServiceAccountMap{})
	pulumi.RegisterOutputType(ServiceAccountOutput{})
	pulumi.RegisterOutputType(ServiceAccountArrayOutput{})
	pulumi.RegisterOutputType(ServiceAccountMapOutput{})
}
