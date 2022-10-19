// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a CFS access rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cfs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cfs.NewAccessRule(ctx, "foo", &Cfs.AccessRuleArgs{
//				AccessGroupId:  pulumi.String("pgroup-7nx89k7l"),
//				AuthClientIp:   pulumi.String("10.10.1.0/24"),
//				Priority:       pulumi.Int(1),
//				RwPermission:   pulumi.String("RO"),
//				UserPermission: pulumi.String("root_squash"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type AccessRule struct {
	pulumi.CustomResourceState

	// ID of a access group.
	AccessGroupId pulumi.StringOutput `pulumi:"accessGroupId"`
	// A single IP or a single IP address range such as 10.1.10.11 or 10.10.1.0/24 indicates that all IPs are allowed. Please note that the IP entered should be CVM's private IP.
	AuthClientIp pulumi.StringOutput `pulumi:"authClientIp"`
	// The priority level of rule. Valid value ranges: (1~100). `1` indicates the highest priority.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Read and write permissions. Valid values are `RO` and `RW`. and default is `RO`.
	RwPermission pulumi.StringPtrOutput `pulumi:"rwPermission"`
	// The permissions of accessing users. Valid values are `allSquash`, `noAllSquash`, `rootSquash` and `noRootSquash`. and default is `rootSquash`. `allSquash` indicates that all access users are mapped as anonymous users or user groups; `noAllSquash` indicates that access users will match local users first and be mapped to anonymous users or user groups after matching failed; `rootSquash` indicates that map access root users to anonymous users or user groups; `noRootSquash` indicates that access root users keep root account permission.
	UserPermission pulumi.StringPtrOutput `pulumi:"userPermission"`
}

// NewAccessRule registers a new resource with the given unique name, arguments, and options.
func NewAccessRule(ctx *pulumi.Context,
	name string, args *AccessRuleArgs, opts ...pulumi.ResourceOption) (*AccessRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessGroupId == nil {
		return nil, errors.New("invalid value for required argument 'AccessGroupId'")
	}
	if args.AuthClientIp == nil {
		return nil, errors.New("invalid value for required argument 'AuthClientIp'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	var resource AccessRule
	err := ctx.RegisterResource("tencentcloud:Cfs/accessRule:AccessRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessRule gets an existing AccessRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessRuleState, opts ...pulumi.ResourceOption) (*AccessRule, error) {
	var resource AccessRule
	err := ctx.ReadResource("tencentcloud:Cfs/accessRule:AccessRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessRule resources.
type accessRuleState struct {
	// ID of a access group.
	AccessGroupId *string `pulumi:"accessGroupId"`
	// A single IP or a single IP address range such as 10.1.10.11 or 10.10.1.0/24 indicates that all IPs are allowed. Please note that the IP entered should be CVM's private IP.
	AuthClientIp *string `pulumi:"authClientIp"`
	// The priority level of rule. Valid value ranges: (1~100). `1` indicates the highest priority.
	Priority *int `pulumi:"priority"`
	// Read and write permissions. Valid values are `RO` and `RW`. and default is `RO`.
	RwPermission *string `pulumi:"rwPermission"`
	// The permissions of accessing users. Valid values are `allSquash`, `noAllSquash`, `rootSquash` and `noRootSquash`. and default is `rootSquash`. `allSquash` indicates that all access users are mapped as anonymous users or user groups; `noAllSquash` indicates that access users will match local users first and be mapped to anonymous users or user groups after matching failed; `rootSquash` indicates that map access root users to anonymous users or user groups; `noRootSquash` indicates that access root users keep root account permission.
	UserPermission *string `pulumi:"userPermission"`
}

type AccessRuleState struct {
	// ID of a access group.
	AccessGroupId pulumi.StringPtrInput
	// A single IP or a single IP address range such as 10.1.10.11 or 10.10.1.0/24 indicates that all IPs are allowed. Please note that the IP entered should be CVM's private IP.
	AuthClientIp pulumi.StringPtrInput
	// The priority level of rule. Valid value ranges: (1~100). `1` indicates the highest priority.
	Priority pulumi.IntPtrInput
	// Read and write permissions. Valid values are `RO` and `RW`. and default is `RO`.
	RwPermission pulumi.StringPtrInput
	// The permissions of accessing users. Valid values are `allSquash`, `noAllSquash`, `rootSquash` and `noRootSquash`. and default is `rootSquash`. `allSquash` indicates that all access users are mapped as anonymous users or user groups; `noAllSquash` indicates that access users will match local users first and be mapped to anonymous users or user groups after matching failed; `rootSquash` indicates that map access root users to anonymous users or user groups; `noRootSquash` indicates that access root users keep root account permission.
	UserPermission pulumi.StringPtrInput
}

func (AccessRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessRuleState)(nil)).Elem()
}

type accessRuleArgs struct {
	// ID of a access group.
	AccessGroupId string `pulumi:"accessGroupId"`
	// A single IP or a single IP address range such as 10.1.10.11 or 10.10.1.0/24 indicates that all IPs are allowed. Please note that the IP entered should be CVM's private IP.
	AuthClientIp string `pulumi:"authClientIp"`
	// The priority level of rule. Valid value ranges: (1~100). `1` indicates the highest priority.
	Priority int `pulumi:"priority"`
	// Read and write permissions. Valid values are `RO` and `RW`. and default is `RO`.
	RwPermission *string `pulumi:"rwPermission"`
	// The permissions of accessing users. Valid values are `allSquash`, `noAllSquash`, `rootSquash` and `noRootSquash`. and default is `rootSquash`. `allSquash` indicates that all access users are mapped as anonymous users or user groups; `noAllSquash` indicates that access users will match local users first and be mapped to anonymous users or user groups after matching failed; `rootSquash` indicates that map access root users to anonymous users or user groups; `noRootSquash` indicates that access root users keep root account permission.
	UserPermission *string `pulumi:"userPermission"`
}

// The set of arguments for constructing a AccessRule resource.
type AccessRuleArgs struct {
	// ID of a access group.
	AccessGroupId pulumi.StringInput
	// A single IP or a single IP address range such as 10.1.10.11 or 10.10.1.0/24 indicates that all IPs are allowed. Please note that the IP entered should be CVM's private IP.
	AuthClientIp pulumi.StringInput
	// The priority level of rule. Valid value ranges: (1~100). `1` indicates the highest priority.
	Priority pulumi.IntInput
	// Read and write permissions. Valid values are `RO` and `RW`. and default is `RO`.
	RwPermission pulumi.StringPtrInput
	// The permissions of accessing users. Valid values are `allSquash`, `noAllSquash`, `rootSquash` and `noRootSquash`. and default is `rootSquash`. `allSquash` indicates that all access users are mapped as anonymous users or user groups; `noAllSquash` indicates that access users will match local users first and be mapped to anonymous users or user groups after matching failed; `rootSquash` indicates that map access root users to anonymous users or user groups; `noRootSquash` indicates that access root users keep root account permission.
	UserPermission pulumi.StringPtrInput
}

func (AccessRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessRuleArgs)(nil)).Elem()
}

type AccessRuleInput interface {
	pulumi.Input

	ToAccessRuleOutput() AccessRuleOutput
	ToAccessRuleOutputWithContext(ctx context.Context) AccessRuleOutput
}

func (*AccessRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessRule)(nil)).Elem()
}

func (i *AccessRule) ToAccessRuleOutput() AccessRuleOutput {
	return i.ToAccessRuleOutputWithContext(context.Background())
}

func (i *AccessRule) ToAccessRuleOutputWithContext(ctx context.Context) AccessRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessRuleOutput)
}

// AccessRuleArrayInput is an input type that accepts AccessRuleArray and AccessRuleArrayOutput values.
// You can construct a concrete instance of `AccessRuleArrayInput` via:
//
//	AccessRuleArray{ AccessRuleArgs{...} }
type AccessRuleArrayInput interface {
	pulumi.Input

	ToAccessRuleArrayOutput() AccessRuleArrayOutput
	ToAccessRuleArrayOutputWithContext(context.Context) AccessRuleArrayOutput
}

type AccessRuleArray []AccessRuleInput

func (AccessRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessRule)(nil)).Elem()
}

func (i AccessRuleArray) ToAccessRuleArrayOutput() AccessRuleArrayOutput {
	return i.ToAccessRuleArrayOutputWithContext(context.Background())
}

func (i AccessRuleArray) ToAccessRuleArrayOutputWithContext(ctx context.Context) AccessRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessRuleArrayOutput)
}

// AccessRuleMapInput is an input type that accepts AccessRuleMap and AccessRuleMapOutput values.
// You can construct a concrete instance of `AccessRuleMapInput` via:
//
//	AccessRuleMap{ "key": AccessRuleArgs{...} }
type AccessRuleMapInput interface {
	pulumi.Input

	ToAccessRuleMapOutput() AccessRuleMapOutput
	ToAccessRuleMapOutputWithContext(context.Context) AccessRuleMapOutput
}

type AccessRuleMap map[string]AccessRuleInput

func (AccessRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessRule)(nil)).Elem()
}

func (i AccessRuleMap) ToAccessRuleMapOutput() AccessRuleMapOutput {
	return i.ToAccessRuleMapOutputWithContext(context.Background())
}

func (i AccessRuleMap) ToAccessRuleMapOutputWithContext(ctx context.Context) AccessRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessRuleMapOutput)
}

type AccessRuleOutput struct{ *pulumi.OutputState }

func (AccessRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessRule)(nil)).Elem()
}

func (o AccessRuleOutput) ToAccessRuleOutput() AccessRuleOutput {
	return o
}

func (o AccessRuleOutput) ToAccessRuleOutputWithContext(ctx context.Context) AccessRuleOutput {
	return o
}

// ID of a access group.
func (o AccessRuleOutput) AccessGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessRule) pulumi.StringOutput { return v.AccessGroupId }).(pulumi.StringOutput)
}

// A single IP or a single IP address range such as 10.1.10.11 or 10.10.1.0/24 indicates that all IPs are allowed. Please note that the IP entered should be CVM's private IP.
func (o AccessRuleOutput) AuthClientIp() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessRule) pulumi.StringOutput { return v.AuthClientIp }).(pulumi.StringOutput)
}

// The priority level of rule. Valid value ranges: (1~100). `1` indicates the highest priority.
func (o AccessRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *AccessRule) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// Read and write permissions. Valid values are `RO` and `RW`. and default is `RO`.
func (o AccessRuleOutput) RwPermission() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessRule) pulumi.StringPtrOutput { return v.RwPermission }).(pulumi.StringPtrOutput)
}

// The permissions of accessing users. Valid values are `allSquash`, `noAllSquash`, `rootSquash` and `noRootSquash`. and default is `rootSquash`. `allSquash` indicates that all access users are mapped as anonymous users or user groups; `noAllSquash` indicates that access users will match local users first and be mapped to anonymous users or user groups after matching failed; `rootSquash` indicates that map access root users to anonymous users or user groups; `noRootSquash` indicates that access root users keep root account permission.
func (o AccessRuleOutput) UserPermission() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessRule) pulumi.StringPtrOutput { return v.UserPermission }).(pulumi.StringPtrOutput)
}

type AccessRuleArrayOutput struct{ *pulumi.OutputState }

func (AccessRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessRule)(nil)).Elem()
}

func (o AccessRuleArrayOutput) ToAccessRuleArrayOutput() AccessRuleArrayOutput {
	return o
}

func (o AccessRuleArrayOutput) ToAccessRuleArrayOutputWithContext(ctx context.Context) AccessRuleArrayOutput {
	return o
}

func (o AccessRuleArrayOutput) Index(i pulumi.IntInput) AccessRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessRule {
		return vs[0].([]*AccessRule)[vs[1].(int)]
	}).(AccessRuleOutput)
}

type AccessRuleMapOutput struct{ *pulumi.OutputState }

func (AccessRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessRule)(nil)).Elem()
}

func (o AccessRuleMapOutput) ToAccessRuleMapOutput() AccessRuleMapOutput {
	return o
}

func (o AccessRuleMapOutput) ToAccessRuleMapOutputWithContext(ctx context.Context) AccessRuleMapOutput {
	return o
}

func (o AccessRuleMapOutput) MapIndex(k pulumi.StringInput) AccessRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessRule {
		return vs[0].(map[string]*AccessRule)[vs[1].(string)]
	}).(AccessRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessRuleInput)(nil)).Elem(), &AccessRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessRuleArrayInput)(nil)).Elem(), AccessRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessRuleMapInput)(nil)).Elem(), AccessRuleMap{})
	pulumi.RegisterOutputType(AccessRuleOutput{})
	pulumi.RegisterOutputType(AccessRuleArrayOutput{})
	pulumi.RegisterOutputType(AccessRuleMapOutput{})
}
