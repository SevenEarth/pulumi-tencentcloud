// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eip

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetAddressQuotaQuotaSet struct {
	// Current count.
	QuotaCurrent int `pulumi:"quotaCurrent"`
	// Quota name: TOTAL_EIP_QUOTA,DAILY_EIP_APPLY,DAILY_PUBLIC_IP_ASSIGN.
	QuotaId string `pulumi:"quotaId"`
	// quota count.
	QuotaLimit int `pulumi:"quotaLimit"`
}

// GetAddressQuotaQuotaSetInput is an input type that accepts GetAddressQuotaQuotaSetArgs and GetAddressQuotaQuotaSetOutput values.
// You can construct a concrete instance of `GetAddressQuotaQuotaSetInput` via:
//
//          GetAddressQuotaQuotaSetArgs{...}
type GetAddressQuotaQuotaSetInput interface {
	pulumi.Input

	ToGetAddressQuotaQuotaSetOutput() GetAddressQuotaQuotaSetOutput
	ToGetAddressQuotaQuotaSetOutputWithContext(context.Context) GetAddressQuotaQuotaSetOutput
}

type GetAddressQuotaQuotaSetArgs struct {
	// Current count.
	QuotaCurrent pulumi.IntInput `pulumi:"quotaCurrent"`
	// Quota name: TOTAL_EIP_QUOTA,DAILY_EIP_APPLY,DAILY_PUBLIC_IP_ASSIGN.
	QuotaId pulumi.StringInput `pulumi:"quotaId"`
	// quota count.
	QuotaLimit pulumi.IntInput `pulumi:"quotaLimit"`
}

func (GetAddressQuotaQuotaSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddressQuotaQuotaSet)(nil)).Elem()
}

func (i GetAddressQuotaQuotaSetArgs) ToGetAddressQuotaQuotaSetOutput() GetAddressQuotaQuotaSetOutput {
	return i.ToGetAddressQuotaQuotaSetOutputWithContext(context.Background())
}

func (i GetAddressQuotaQuotaSetArgs) ToGetAddressQuotaQuotaSetOutputWithContext(ctx context.Context) GetAddressQuotaQuotaSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAddressQuotaQuotaSetOutput)
}

// GetAddressQuotaQuotaSetArrayInput is an input type that accepts GetAddressQuotaQuotaSetArray and GetAddressQuotaQuotaSetArrayOutput values.
// You can construct a concrete instance of `GetAddressQuotaQuotaSetArrayInput` via:
//
//          GetAddressQuotaQuotaSetArray{ GetAddressQuotaQuotaSetArgs{...} }
type GetAddressQuotaQuotaSetArrayInput interface {
	pulumi.Input

	ToGetAddressQuotaQuotaSetArrayOutput() GetAddressQuotaQuotaSetArrayOutput
	ToGetAddressQuotaQuotaSetArrayOutputWithContext(context.Context) GetAddressQuotaQuotaSetArrayOutput
}

type GetAddressQuotaQuotaSetArray []GetAddressQuotaQuotaSetInput

func (GetAddressQuotaQuotaSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAddressQuotaQuotaSet)(nil)).Elem()
}

func (i GetAddressQuotaQuotaSetArray) ToGetAddressQuotaQuotaSetArrayOutput() GetAddressQuotaQuotaSetArrayOutput {
	return i.ToGetAddressQuotaQuotaSetArrayOutputWithContext(context.Background())
}

func (i GetAddressQuotaQuotaSetArray) ToGetAddressQuotaQuotaSetArrayOutputWithContext(ctx context.Context) GetAddressQuotaQuotaSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAddressQuotaQuotaSetArrayOutput)
}

type GetAddressQuotaQuotaSetOutput struct{ *pulumi.OutputState }

func (GetAddressQuotaQuotaSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddressQuotaQuotaSet)(nil)).Elem()
}

func (o GetAddressQuotaQuotaSetOutput) ToGetAddressQuotaQuotaSetOutput() GetAddressQuotaQuotaSetOutput {
	return o
}

func (o GetAddressQuotaQuotaSetOutput) ToGetAddressQuotaQuotaSetOutputWithContext(ctx context.Context) GetAddressQuotaQuotaSetOutput {
	return o
}

// Current count.
func (o GetAddressQuotaQuotaSetOutput) QuotaCurrent() pulumi.IntOutput {
	return o.ApplyT(func(v GetAddressQuotaQuotaSet) int { return v.QuotaCurrent }).(pulumi.IntOutput)
}

// Quota name: TOTAL_EIP_QUOTA,DAILY_EIP_APPLY,DAILY_PUBLIC_IP_ASSIGN.
func (o GetAddressQuotaQuotaSetOutput) QuotaId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAddressQuotaQuotaSet) string { return v.QuotaId }).(pulumi.StringOutput)
}

// quota count.
func (o GetAddressQuotaQuotaSetOutput) QuotaLimit() pulumi.IntOutput {
	return o.ApplyT(func(v GetAddressQuotaQuotaSet) int { return v.QuotaLimit }).(pulumi.IntOutput)
}

type GetAddressQuotaQuotaSetArrayOutput struct{ *pulumi.OutputState }

func (GetAddressQuotaQuotaSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAddressQuotaQuotaSet)(nil)).Elem()
}

func (o GetAddressQuotaQuotaSetArrayOutput) ToGetAddressQuotaQuotaSetArrayOutput() GetAddressQuotaQuotaSetArrayOutput {
	return o
}

func (o GetAddressQuotaQuotaSetArrayOutput) ToGetAddressQuotaQuotaSetArrayOutputWithContext(ctx context.Context) GetAddressQuotaQuotaSetArrayOutput {
	return o
}

func (o GetAddressQuotaQuotaSetArrayOutput) Index(i pulumi.IntInput) GetAddressQuotaQuotaSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetAddressQuotaQuotaSet {
		return vs[0].([]GetAddressQuotaQuotaSet)[vs[1].(int)]
	}).(GetAddressQuotaQuotaSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetAddressQuotaQuotaSetInput)(nil)).Elem(), GetAddressQuotaQuotaSetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetAddressQuotaQuotaSetArrayInput)(nil)).Elem(), GetAddressQuotaQuotaSetArray{})
	pulumi.RegisterOutputType(GetAddressQuotaQuotaSetOutput{})
	pulumi.RegisterOutputType(GetAddressQuotaQuotaSetArrayOutput{})
}