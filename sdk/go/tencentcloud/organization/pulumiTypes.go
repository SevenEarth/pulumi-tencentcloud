// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OrgMemberOrgPermission struct {
	// Permissions ID.
	Id *int `pulumi:"id"`
	// Member name.
	Name *string `pulumi:"name"`
}

// OrgMemberOrgPermissionInput is an input type that accepts OrgMemberOrgPermissionArgs and OrgMemberOrgPermissionOutput values.
// You can construct a concrete instance of `OrgMemberOrgPermissionInput` via:
//
//          OrgMemberOrgPermissionArgs{...}
type OrgMemberOrgPermissionInput interface {
	pulumi.Input

	ToOrgMemberOrgPermissionOutput() OrgMemberOrgPermissionOutput
	ToOrgMemberOrgPermissionOutputWithContext(context.Context) OrgMemberOrgPermissionOutput
}

type OrgMemberOrgPermissionArgs struct {
	// Permissions ID.
	Id pulumi.IntPtrInput `pulumi:"id"`
	// Member name.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (OrgMemberOrgPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrgMemberOrgPermission)(nil)).Elem()
}

func (i OrgMemberOrgPermissionArgs) ToOrgMemberOrgPermissionOutput() OrgMemberOrgPermissionOutput {
	return i.ToOrgMemberOrgPermissionOutputWithContext(context.Background())
}

func (i OrgMemberOrgPermissionArgs) ToOrgMemberOrgPermissionOutputWithContext(ctx context.Context) OrgMemberOrgPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgMemberOrgPermissionOutput)
}

// OrgMemberOrgPermissionArrayInput is an input type that accepts OrgMemberOrgPermissionArray and OrgMemberOrgPermissionArrayOutput values.
// You can construct a concrete instance of `OrgMemberOrgPermissionArrayInput` via:
//
//          OrgMemberOrgPermissionArray{ OrgMemberOrgPermissionArgs{...} }
type OrgMemberOrgPermissionArrayInput interface {
	pulumi.Input

	ToOrgMemberOrgPermissionArrayOutput() OrgMemberOrgPermissionArrayOutput
	ToOrgMemberOrgPermissionArrayOutputWithContext(context.Context) OrgMemberOrgPermissionArrayOutput
}

type OrgMemberOrgPermissionArray []OrgMemberOrgPermissionInput

func (OrgMemberOrgPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OrgMemberOrgPermission)(nil)).Elem()
}

func (i OrgMemberOrgPermissionArray) ToOrgMemberOrgPermissionArrayOutput() OrgMemberOrgPermissionArrayOutput {
	return i.ToOrgMemberOrgPermissionArrayOutputWithContext(context.Background())
}

func (i OrgMemberOrgPermissionArray) ToOrgMemberOrgPermissionArrayOutputWithContext(ctx context.Context) OrgMemberOrgPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgMemberOrgPermissionArrayOutput)
}

type OrgMemberOrgPermissionOutput struct{ *pulumi.OutputState }

func (OrgMemberOrgPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrgMemberOrgPermission)(nil)).Elem()
}

func (o OrgMemberOrgPermissionOutput) ToOrgMemberOrgPermissionOutput() OrgMemberOrgPermissionOutput {
	return o
}

func (o OrgMemberOrgPermissionOutput) ToOrgMemberOrgPermissionOutputWithContext(ctx context.Context) OrgMemberOrgPermissionOutput {
	return o
}

// Permissions ID.
func (o OrgMemberOrgPermissionOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OrgMemberOrgPermission) *int { return v.Id }).(pulumi.IntPtrOutput)
}

// Member name.
func (o OrgMemberOrgPermissionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrgMemberOrgPermission) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type OrgMemberOrgPermissionArrayOutput struct{ *pulumi.OutputState }

func (OrgMemberOrgPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OrgMemberOrgPermission)(nil)).Elem()
}

func (o OrgMemberOrgPermissionArrayOutput) ToOrgMemberOrgPermissionArrayOutput() OrgMemberOrgPermissionArrayOutput {
	return o
}

func (o OrgMemberOrgPermissionArrayOutput) ToOrgMemberOrgPermissionArrayOutputWithContext(ctx context.Context) OrgMemberOrgPermissionArrayOutput {
	return o
}

func (o OrgMemberOrgPermissionArrayOutput) Index(i pulumi.IntInput) OrgMemberOrgPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OrgMemberOrgPermission {
		return vs[0].([]OrgMemberOrgPermission)[vs[1].(int)]
	}).(OrgMemberOrgPermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgMemberOrgPermissionInput)(nil)).Elem(), OrgMemberOrgPermissionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgMemberOrgPermissionArrayInput)(nil)).Elem(), OrgMemberOrgPermissionArray{})
	pulumi.RegisterOutputType(OrgMemberOrgPermissionOutput{})
	pulumi.RegisterOutputType(OrgMemberOrgPermissionArrayOutput{})
}
