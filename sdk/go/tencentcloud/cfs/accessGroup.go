// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a CFS access group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cfs"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cfs.NewAccessGroup(ctx, "example", &Cfs.AccessGroupArgs{
// 			Description: pulumi.String("desc."),
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
// CFS access group can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Cfs/accessGroup:AccessGroup example pgroup-7nx89k7l
// ```
type AccessGroup struct {
	pulumi.CustomResourceState

	// Create time of the access group.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the access group, and max length is 255.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the access group, and max length is 64.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAccessGroup registers a new resource with the given unique name, arguments, and options.
func NewAccessGroup(ctx *pulumi.Context,
	name string, args *AccessGroupArgs, opts ...pulumi.ResourceOption) (*AccessGroup, error) {
	if args == nil {
		args = &AccessGroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource AccessGroup
	err := ctx.RegisterResource("tencentcloud:Cfs/accessGroup:AccessGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessGroup gets an existing AccessGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessGroupState, opts ...pulumi.ResourceOption) (*AccessGroup, error) {
	var resource AccessGroup
	err := ctx.ReadResource("tencentcloud:Cfs/accessGroup:AccessGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessGroup resources.
type accessGroupState struct {
	// Create time of the access group.
	CreateTime *string `pulumi:"createTime"`
	// Description of the access group, and max length is 255.
	Description *string `pulumi:"description"`
	// Name of the access group, and max length is 64.
	Name *string `pulumi:"name"`
}

type AccessGroupState struct {
	// Create time of the access group.
	CreateTime pulumi.StringPtrInput
	// Description of the access group, and max length is 255.
	Description pulumi.StringPtrInput
	// Name of the access group, and max length is 64.
	Name pulumi.StringPtrInput
}

func (AccessGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessGroupState)(nil)).Elem()
}

type accessGroupArgs struct {
	// Description of the access group, and max length is 255.
	Description *string `pulumi:"description"`
	// Name of the access group, and max length is 64.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a AccessGroup resource.
type AccessGroupArgs struct {
	// Description of the access group, and max length is 255.
	Description pulumi.StringPtrInput
	// Name of the access group, and max length is 64.
	Name pulumi.StringPtrInput
}

func (AccessGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessGroupArgs)(nil)).Elem()
}

type AccessGroupInput interface {
	pulumi.Input

	ToAccessGroupOutput() AccessGroupOutput
	ToAccessGroupOutputWithContext(ctx context.Context) AccessGroupOutput
}

func (*AccessGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessGroup)(nil)).Elem()
}

func (i *AccessGroup) ToAccessGroupOutput() AccessGroupOutput {
	return i.ToAccessGroupOutputWithContext(context.Background())
}

func (i *AccessGroup) ToAccessGroupOutputWithContext(ctx context.Context) AccessGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessGroupOutput)
}

// AccessGroupArrayInput is an input type that accepts AccessGroupArray and AccessGroupArrayOutput values.
// You can construct a concrete instance of `AccessGroupArrayInput` via:
//
//          AccessGroupArray{ AccessGroupArgs{...} }
type AccessGroupArrayInput interface {
	pulumi.Input

	ToAccessGroupArrayOutput() AccessGroupArrayOutput
	ToAccessGroupArrayOutputWithContext(context.Context) AccessGroupArrayOutput
}

type AccessGroupArray []AccessGroupInput

func (AccessGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessGroup)(nil)).Elem()
}

func (i AccessGroupArray) ToAccessGroupArrayOutput() AccessGroupArrayOutput {
	return i.ToAccessGroupArrayOutputWithContext(context.Background())
}

func (i AccessGroupArray) ToAccessGroupArrayOutputWithContext(ctx context.Context) AccessGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessGroupArrayOutput)
}

// AccessGroupMapInput is an input type that accepts AccessGroupMap and AccessGroupMapOutput values.
// You can construct a concrete instance of `AccessGroupMapInput` via:
//
//          AccessGroupMap{ "key": AccessGroupArgs{...} }
type AccessGroupMapInput interface {
	pulumi.Input

	ToAccessGroupMapOutput() AccessGroupMapOutput
	ToAccessGroupMapOutputWithContext(context.Context) AccessGroupMapOutput
}

type AccessGroupMap map[string]AccessGroupInput

func (AccessGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessGroup)(nil)).Elem()
}

func (i AccessGroupMap) ToAccessGroupMapOutput() AccessGroupMapOutput {
	return i.ToAccessGroupMapOutputWithContext(context.Background())
}

func (i AccessGroupMap) ToAccessGroupMapOutputWithContext(ctx context.Context) AccessGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessGroupMapOutput)
}

type AccessGroupOutput struct{ *pulumi.OutputState }

func (AccessGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessGroup)(nil)).Elem()
}

func (o AccessGroupOutput) ToAccessGroupOutput() AccessGroupOutput {
	return o
}

func (o AccessGroupOutput) ToAccessGroupOutputWithContext(ctx context.Context) AccessGroupOutput {
	return o
}

// Create time of the access group.
func (o AccessGroupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessGroup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the access group, and max length is 255.
func (o AccessGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the access group, and max length is 64.
func (o AccessGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type AccessGroupArrayOutput struct{ *pulumi.OutputState }

func (AccessGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessGroup)(nil)).Elem()
}

func (o AccessGroupArrayOutput) ToAccessGroupArrayOutput() AccessGroupArrayOutput {
	return o
}

func (o AccessGroupArrayOutput) ToAccessGroupArrayOutputWithContext(ctx context.Context) AccessGroupArrayOutput {
	return o
}

func (o AccessGroupArrayOutput) Index(i pulumi.IntInput) AccessGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessGroup {
		return vs[0].([]*AccessGroup)[vs[1].(int)]
	}).(AccessGroupOutput)
}

type AccessGroupMapOutput struct{ *pulumi.OutputState }

func (AccessGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessGroup)(nil)).Elem()
}

func (o AccessGroupMapOutput) ToAccessGroupMapOutput() AccessGroupMapOutput {
	return o
}

func (o AccessGroupMapOutput) ToAccessGroupMapOutputWithContext(ctx context.Context) AccessGroupMapOutput {
	return o
}

func (o AccessGroupMapOutput) MapIndex(k pulumi.StringInput) AccessGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessGroup {
		return vs[0].(map[string]*AccessGroup)[vs[1].(string)]
	}).(AccessGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessGroupInput)(nil)).Elem(), &AccessGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessGroupArrayInput)(nil)).Elem(), AccessGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessGroupMapInput)(nil)).Elem(), AccessGroupMap{})
	pulumi.RegisterOutputType(AccessGroupOutput{})
	pulumi.RegisterOutputType(AccessGroupArrayOutput{})
	pulumi.RegisterOutputType(AccessGroupMapOutput{})
}
