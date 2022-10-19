// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package audit

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetCosRegionsAuditCosRegionList struct {
	// Cos region.
	CosRegion string `pulumi:"cosRegion"`
	// Cos region chinese name.
	CosRegionName string `pulumi:"cosRegionName"`
}

// GetCosRegionsAuditCosRegionListInput is an input type that accepts GetCosRegionsAuditCosRegionListArgs and GetCosRegionsAuditCosRegionListOutput values.
// You can construct a concrete instance of `GetCosRegionsAuditCosRegionListInput` via:
//
//	GetCosRegionsAuditCosRegionListArgs{...}
type GetCosRegionsAuditCosRegionListInput interface {
	pulumi.Input

	ToGetCosRegionsAuditCosRegionListOutput() GetCosRegionsAuditCosRegionListOutput
	ToGetCosRegionsAuditCosRegionListOutputWithContext(context.Context) GetCosRegionsAuditCosRegionListOutput
}

type GetCosRegionsAuditCosRegionListArgs struct {
	// Cos region.
	CosRegion pulumi.StringInput `pulumi:"cosRegion"`
	// Cos region chinese name.
	CosRegionName pulumi.StringInput `pulumi:"cosRegionName"`
}

func (GetCosRegionsAuditCosRegionListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCosRegionsAuditCosRegionList)(nil)).Elem()
}

func (i GetCosRegionsAuditCosRegionListArgs) ToGetCosRegionsAuditCosRegionListOutput() GetCosRegionsAuditCosRegionListOutput {
	return i.ToGetCosRegionsAuditCosRegionListOutputWithContext(context.Background())
}

func (i GetCosRegionsAuditCosRegionListArgs) ToGetCosRegionsAuditCosRegionListOutputWithContext(ctx context.Context) GetCosRegionsAuditCosRegionListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCosRegionsAuditCosRegionListOutput)
}

// GetCosRegionsAuditCosRegionListArrayInput is an input type that accepts GetCosRegionsAuditCosRegionListArray and GetCosRegionsAuditCosRegionListArrayOutput values.
// You can construct a concrete instance of `GetCosRegionsAuditCosRegionListArrayInput` via:
//
//	GetCosRegionsAuditCosRegionListArray{ GetCosRegionsAuditCosRegionListArgs{...} }
type GetCosRegionsAuditCosRegionListArrayInput interface {
	pulumi.Input

	ToGetCosRegionsAuditCosRegionListArrayOutput() GetCosRegionsAuditCosRegionListArrayOutput
	ToGetCosRegionsAuditCosRegionListArrayOutputWithContext(context.Context) GetCosRegionsAuditCosRegionListArrayOutput
}

type GetCosRegionsAuditCosRegionListArray []GetCosRegionsAuditCosRegionListInput

func (GetCosRegionsAuditCosRegionListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCosRegionsAuditCosRegionList)(nil)).Elem()
}

func (i GetCosRegionsAuditCosRegionListArray) ToGetCosRegionsAuditCosRegionListArrayOutput() GetCosRegionsAuditCosRegionListArrayOutput {
	return i.ToGetCosRegionsAuditCosRegionListArrayOutputWithContext(context.Background())
}

func (i GetCosRegionsAuditCosRegionListArray) ToGetCosRegionsAuditCosRegionListArrayOutputWithContext(ctx context.Context) GetCosRegionsAuditCosRegionListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCosRegionsAuditCosRegionListArrayOutput)
}

type GetCosRegionsAuditCosRegionListOutput struct{ *pulumi.OutputState }

func (GetCosRegionsAuditCosRegionListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCosRegionsAuditCosRegionList)(nil)).Elem()
}

func (o GetCosRegionsAuditCosRegionListOutput) ToGetCosRegionsAuditCosRegionListOutput() GetCosRegionsAuditCosRegionListOutput {
	return o
}

func (o GetCosRegionsAuditCosRegionListOutput) ToGetCosRegionsAuditCosRegionListOutputWithContext(ctx context.Context) GetCosRegionsAuditCosRegionListOutput {
	return o
}

// Cos region.
func (o GetCosRegionsAuditCosRegionListOutput) CosRegion() pulumi.StringOutput {
	return o.ApplyT(func(v GetCosRegionsAuditCosRegionList) string { return v.CosRegion }).(pulumi.StringOutput)
}

// Cos region chinese name.
func (o GetCosRegionsAuditCosRegionListOutput) CosRegionName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCosRegionsAuditCosRegionList) string { return v.CosRegionName }).(pulumi.StringOutput)
}

type GetCosRegionsAuditCosRegionListArrayOutput struct{ *pulumi.OutputState }

func (GetCosRegionsAuditCosRegionListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCosRegionsAuditCosRegionList)(nil)).Elem()
}

func (o GetCosRegionsAuditCosRegionListArrayOutput) ToGetCosRegionsAuditCosRegionListArrayOutput() GetCosRegionsAuditCosRegionListArrayOutput {
	return o
}

func (o GetCosRegionsAuditCosRegionListArrayOutput) ToGetCosRegionsAuditCosRegionListArrayOutputWithContext(ctx context.Context) GetCosRegionsAuditCosRegionListArrayOutput {
	return o
}

func (o GetCosRegionsAuditCosRegionListArrayOutput) Index(i pulumi.IntInput) GetCosRegionsAuditCosRegionListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetCosRegionsAuditCosRegionList {
		return vs[0].([]GetCosRegionsAuditCosRegionList)[vs[1].(int)]
	}).(GetCosRegionsAuditCosRegionListOutput)
}

type GetKeyAliasAuditKeyAliasList struct {
	// Key alias.
	KeyAlias string `pulumi:"keyAlias"`
	// Key ID.
	KeyId string `pulumi:"keyId"`
}

// GetKeyAliasAuditKeyAliasListInput is an input type that accepts GetKeyAliasAuditKeyAliasListArgs and GetKeyAliasAuditKeyAliasListOutput values.
// You can construct a concrete instance of `GetKeyAliasAuditKeyAliasListInput` via:
//
//	GetKeyAliasAuditKeyAliasListArgs{...}
type GetKeyAliasAuditKeyAliasListInput interface {
	pulumi.Input

	ToGetKeyAliasAuditKeyAliasListOutput() GetKeyAliasAuditKeyAliasListOutput
	ToGetKeyAliasAuditKeyAliasListOutputWithContext(context.Context) GetKeyAliasAuditKeyAliasListOutput
}

type GetKeyAliasAuditKeyAliasListArgs struct {
	// Key alias.
	KeyAlias pulumi.StringInput `pulumi:"keyAlias"`
	// Key ID.
	KeyId pulumi.StringInput `pulumi:"keyId"`
}

func (GetKeyAliasAuditKeyAliasListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeyAliasAuditKeyAliasList)(nil)).Elem()
}

func (i GetKeyAliasAuditKeyAliasListArgs) ToGetKeyAliasAuditKeyAliasListOutput() GetKeyAliasAuditKeyAliasListOutput {
	return i.ToGetKeyAliasAuditKeyAliasListOutputWithContext(context.Background())
}

func (i GetKeyAliasAuditKeyAliasListArgs) ToGetKeyAliasAuditKeyAliasListOutputWithContext(ctx context.Context) GetKeyAliasAuditKeyAliasListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetKeyAliasAuditKeyAliasListOutput)
}

// GetKeyAliasAuditKeyAliasListArrayInput is an input type that accepts GetKeyAliasAuditKeyAliasListArray and GetKeyAliasAuditKeyAliasListArrayOutput values.
// You can construct a concrete instance of `GetKeyAliasAuditKeyAliasListArrayInput` via:
//
//	GetKeyAliasAuditKeyAliasListArray{ GetKeyAliasAuditKeyAliasListArgs{...} }
type GetKeyAliasAuditKeyAliasListArrayInput interface {
	pulumi.Input

	ToGetKeyAliasAuditKeyAliasListArrayOutput() GetKeyAliasAuditKeyAliasListArrayOutput
	ToGetKeyAliasAuditKeyAliasListArrayOutputWithContext(context.Context) GetKeyAliasAuditKeyAliasListArrayOutput
}

type GetKeyAliasAuditKeyAliasListArray []GetKeyAliasAuditKeyAliasListInput

func (GetKeyAliasAuditKeyAliasListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetKeyAliasAuditKeyAliasList)(nil)).Elem()
}

func (i GetKeyAliasAuditKeyAliasListArray) ToGetKeyAliasAuditKeyAliasListArrayOutput() GetKeyAliasAuditKeyAliasListArrayOutput {
	return i.ToGetKeyAliasAuditKeyAliasListArrayOutputWithContext(context.Background())
}

func (i GetKeyAliasAuditKeyAliasListArray) ToGetKeyAliasAuditKeyAliasListArrayOutputWithContext(ctx context.Context) GetKeyAliasAuditKeyAliasListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetKeyAliasAuditKeyAliasListArrayOutput)
}

type GetKeyAliasAuditKeyAliasListOutput struct{ *pulumi.OutputState }

func (GetKeyAliasAuditKeyAliasListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeyAliasAuditKeyAliasList)(nil)).Elem()
}

func (o GetKeyAliasAuditKeyAliasListOutput) ToGetKeyAliasAuditKeyAliasListOutput() GetKeyAliasAuditKeyAliasListOutput {
	return o
}

func (o GetKeyAliasAuditKeyAliasListOutput) ToGetKeyAliasAuditKeyAliasListOutputWithContext(ctx context.Context) GetKeyAliasAuditKeyAliasListOutput {
	return o
}

// Key alias.
func (o GetKeyAliasAuditKeyAliasListOutput) KeyAlias() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeyAliasAuditKeyAliasList) string { return v.KeyAlias }).(pulumi.StringOutput)
}

// Key ID.
func (o GetKeyAliasAuditKeyAliasListOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeyAliasAuditKeyAliasList) string { return v.KeyId }).(pulumi.StringOutput)
}

type GetKeyAliasAuditKeyAliasListArrayOutput struct{ *pulumi.OutputState }

func (GetKeyAliasAuditKeyAliasListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetKeyAliasAuditKeyAliasList)(nil)).Elem()
}

func (o GetKeyAliasAuditKeyAliasListArrayOutput) ToGetKeyAliasAuditKeyAliasListArrayOutput() GetKeyAliasAuditKeyAliasListArrayOutput {
	return o
}

func (o GetKeyAliasAuditKeyAliasListArrayOutput) ToGetKeyAliasAuditKeyAliasListArrayOutputWithContext(ctx context.Context) GetKeyAliasAuditKeyAliasListArrayOutput {
	return o
}

func (o GetKeyAliasAuditKeyAliasListArrayOutput) Index(i pulumi.IntInput) GetKeyAliasAuditKeyAliasListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetKeyAliasAuditKeyAliasList {
		return vs[0].([]GetKeyAliasAuditKeyAliasList)[vs[1].(int)]
	}).(GetKeyAliasAuditKeyAliasListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetCosRegionsAuditCosRegionListInput)(nil)).Elem(), GetCosRegionsAuditCosRegionListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCosRegionsAuditCosRegionListArrayInput)(nil)).Elem(), GetCosRegionsAuditCosRegionListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetKeyAliasAuditKeyAliasListInput)(nil)).Elem(), GetKeyAliasAuditKeyAliasListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetKeyAliasAuditKeyAliasListArrayInput)(nil)).Elem(), GetKeyAliasAuditKeyAliasListArray{})
	pulumi.RegisterOutputType(GetCosRegionsAuditCosRegionListOutput{})
	pulumi.RegisterOutputType(GetCosRegionsAuditCosRegionListArrayOutput{})
	pulumi.RegisterOutputType(GetKeyAliasAuditKeyAliasListOutput{})
	pulumi.RegisterOutputType(GetKeyAliasAuditKeyAliasListArrayOutput{})
}
