// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package reserved

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

var _ = internal.GetEnvOrDefault

type GetInstanceConfigsConfigList struct {
	// The available zone that the reserved instance locates at.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Configuration ID of the purchasable reserved instance.
	ConfigId string `pulumi:"configId"`
	// Settlement currency of the reserved instance, which is a standard currency code as listed in ISO 4217.
	CurrencyCode string `pulumi:"currencyCode"`
	// Validity period of the reserved instance. Valid values are `31536000`(1 year) and `94608000`(3 years).
	Duration int `pulumi:"duration"`
	// The type of reserved instance.
	InstanceType string `pulumi:"instanceType"`
	// Filter by Payment Type. Such as All Upfront.
	OfferingType string `pulumi:"offeringType"`
	// Platform of the reserved instance.
	Platform string `pulumi:"platform"`
	// Purchase price of the reserved instance.
	Price float64 `pulumi:"price"`
	// UsagePrice of the reserved instance.
	UsagePrice float64 `pulumi:"usagePrice"`
}

// GetInstanceConfigsConfigListInput is an input type that accepts GetInstanceConfigsConfigListArgs and GetInstanceConfigsConfigListOutput values.
// You can construct a concrete instance of `GetInstanceConfigsConfigListInput` via:
//
//	GetInstanceConfigsConfigListArgs{...}
type GetInstanceConfigsConfigListInput interface {
	pulumi.Input

	ToGetInstanceConfigsConfigListOutput() GetInstanceConfigsConfigListOutput
	ToGetInstanceConfigsConfigListOutputWithContext(context.Context) GetInstanceConfigsConfigListOutput
}

type GetInstanceConfigsConfigListArgs struct {
	// The available zone that the reserved instance locates at.
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	// Configuration ID of the purchasable reserved instance.
	ConfigId pulumi.StringInput `pulumi:"configId"`
	// Settlement currency of the reserved instance, which is a standard currency code as listed in ISO 4217.
	CurrencyCode pulumi.StringInput `pulumi:"currencyCode"`
	// Validity period of the reserved instance. Valid values are `31536000`(1 year) and `94608000`(3 years).
	Duration pulumi.IntInput `pulumi:"duration"`
	// The type of reserved instance.
	InstanceType pulumi.StringInput `pulumi:"instanceType"`
	// Filter by Payment Type. Such as All Upfront.
	OfferingType pulumi.StringInput `pulumi:"offeringType"`
	// Platform of the reserved instance.
	Platform pulumi.StringInput `pulumi:"platform"`
	// Purchase price of the reserved instance.
	Price pulumi.Float64Input `pulumi:"price"`
	// UsagePrice of the reserved instance.
	UsagePrice pulumi.Float64Input `pulumi:"usagePrice"`
}

func (GetInstanceConfigsConfigListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceConfigsConfigList)(nil)).Elem()
}

func (i GetInstanceConfigsConfigListArgs) ToGetInstanceConfigsConfigListOutput() GetInstanceConfigsConfigListOutput {
	return i.ToGetInstanceConfigsConfigListOutputWithContext(context.Background())
}

func (i GetInstanceConfigsConfigListArgs) ToGetInstanceConfigsConfigListOutputWithContext(ctx context.Context) GetInstanceConfigsConfigListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceConfigsConfigListOutput)
}

// GetInstanceConfigsConfigListArrayInput is an input type that accepts GetInstanceConfigsConfigListArray and GetInstanceConfigsConfigListArrayOutput values.
// You can construct a concrete instance of `GetInstanceConfigsConfigListArrayInput` via:
//
//	GetInstanceConfigsConfigListArray{ GetInstanceConfigsConfigListArgs{...} }
type GetInstanceConfigsConfigListArrayInput interface {
	pulumi.Input

	ToGetInstanceConfigsConfigListArrayOutput() GetInstanceConfigsConfigListArrayOutput
	ToGetInstanceConfigsConfigListArrayOutputWithContext(context.Context) GetInstanceConfigsConfigListArrayOutput
}

type GetInstanceConfigsConfigListArray []GetInstanceConfigsConfigListInput

func (GetInstanceConfigsConfigListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceConfigsConfigList)(nil)).Elem()
}

func (i GetInstanceConfigsConfigListArray) ToGetInstanceConfigsConfigListArrayOutput() GetInstanceConfigsConfigListArrayOutput {
	return i.ToGetInstanceConfigsConfigListArrayOutputWithContext(context.Background())
}

func (i GetInstanceConfigsConfigListArray) ToGetInstanceConfigsConfigListArrayOutputWithContext(ctx context.Context) GetInstanceConfigsConfigListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceConfigsConfigListArrayOutput)
}

type GetInstanceConfigsConfigListOutput struct{ *pulumi.OutputState }

func (GetInstanceConfigsConfigListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceConfigsConfigList)(nil)).Elem()
}

func (o GetInstanceConfigsConfigListOutput) ToGetInstanceConfigsConfigListOutput() GetInstanceConfigsConfigListOutput {
	return o
}

func (o GetInstanceConfigsConfigListOutput) ToGetInstanceConfigsConfigListOutputWithContext(ctx context.Context) GetInstanceConfigsConfigListOutput {
	return o
}

// The available zone that the reserved instance locates at.
func (o GetInstanceConfigsConfigListOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceConfigsConfigList) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Configuration ID of the purchasable reserved instance.
func (o GetInstanceConfigsConfigListOutput) ConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceConfigsConfigList) string { return v.ConfigId }).(pulumi.StringOutput)
}

// Settlement currency of the reserved instance, which is a standard currency code as listed in ISO 4217.
func (o GetInstanceConfigsConfigListOutput) CurrencyCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceConfigsConfigList) string { return v.CurrencyCode }).(pulumi.StringOutput)
}

// Validity period of the reserved instance. Valid values are `31536000`(1 year) and `94608000`(3 years).
func (o GetInstanceConfigsConfigListOutput) Duration() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstanceConfigsConfigList) int { return v.Duration }).(pulumi.IntOutput)
}

// The type of reserved instance.
func (o GetInstanceConfigsConfigListOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceConfigsConfigList) string { return v.InstanceType }).(pulumi.StringOutput)
}

// Filter by Payment Type. Such as All Upfront.
func (o GetInstanceConfigsConfigListOutput) OfferingType() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceConfigsConfigList) string { return v.OfferingType }).(pulumi.StringOutput)
}

// Platform of the reserved instance.
func (o GetInstanceConfigsConfigListOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceConfigsConfigList) string { return v.Platform }).(pulumi.StringOutput)
}

// Purchase price of the reserved instance.
func (o GetInstanceConfigsConfigListOutput) Price() pulumi.Float64Output {
	return o.ApplyT(func(v GetInstanceConfigsConfigList) float64 { return v.Price }).(pulumi.Float64Output)
}

// UsagePrice of the reserved instance.
func (o GetInstanceConfigsConfigListOutput) UsagePrice() pulumi.Float64Output {
	return o.ApplyT(func(v GetInstanceConfigsConfigList) float64 { return v.UsagePrice }).(pulumi.Float64Output)
}

type GetInstanceConfigsConfigListArrayOutput struct{ *pulumi.OutputState }

func (GetInstanceConfigsConfigListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceConfigsConfigList)(nil)).Elem()
}

func (o GetInstanceConfigsConfigListArrayOutput) ToGetInstanceConfigsConfigListArrayOutput() GetInstanceConfigsConfigListArrayOutput {
	return o
}

func (o GetInstanceConfigsConfigListArrayOutput) ToGetInstanceConfigsConfigListArrayOutputWithContext(ctx context.Context) GetInstanceConfigsConfigListArrayOutput {
	return o
}

func (o GetInstanceConfigsConfigListArrayOutput) Index(i pulumi.IntInput) GetInstanceConfigsConfigListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetInstanceConfigsConfigList {
		return vs[0].([]GetInstanceConfigsConfigList)[vs[1].(int)]
	}).(GetInstanceConfigsConfigListOutput)
}

type GetInstancesReservedInstanceList struct {
	// The available zone that the reserved instance locates at.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Expiry time of the reserved instance.
	EndTime string `pulumi:"endTime"`
	// Number of reserved instance.
	InstanceCount int `pulumi:"instanceCount"`
	// The type of reserved instance.
	InstanceType string `pulumi:"instanceType"`
	// ID of the reserved instance to be query.
	ReservedInstanceId string `pulumi:"reservedInstanceId"`
	// Start time of the reserved instance.
	StartTime string `pulumi:"startTime"`
	// Status of the reserved instance.
	Status string `pulumi:"status"`
}

// GetInstancesReservedInstanceListInput is an input type that accepts GetInstancesReservedInstanceListArgs and GetInstancesReservedInstanceListOutput values.
// You can construct a concrete instance of `GetInstancesReservedInstanceListInput` via:
//
//	GetInstancesReservedInstanceListArgs{...}
type GetInstancesReservedInstanceListInput interface {
	pulumi.Input

	ToGetInstancesReservedInstanceListOutput() GetInstancesReservedInstanceListOutput
	ToGetInstancesReservedInstanceListOutputWithContext(context.Context) GetInstancesReservedInstanceListOutput
}

type GetInstancesReservedInstanceListArgs struct {
	// The available zone that the reserved instance locates at.
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	// Expiry time of the reserved instance.
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// Number of reserved instance.
	InstanceCount pulumi.IntInput `pulumi:"instanceCount"`
	// The type of reserved instance.
	InstanceType pulumi.StringInput `pulumi:"instanceType"`
	// ID of the reserved instance to be query.
	ReservedInstanceId pulumi.StringInput `pulumi:"reservedInstanceId"`
	// Start time of the reserved instance.
	StartTime pulumi.StringInput `pulumi:"startTime"`
	// Status of the reserved instance.
	Status pulumi.StringInput `pulumi:"status"`
}

func (GetInstancesReservedInstanceListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesReservedInstanceList)(nil)).Elem()
}

func (i GetInstancesReservedInstanceListArgs) ToGetInstancesReservedInstanceListOutput() GetInstancesReservedInstanceListOutput {
	return i.ToGetInstancesReservedInstanceListOutputWithContext(context.Background())
}

func (i GetInstancesReservedInstanceListArgs) ToGetInstancesReservedInstanceListOutputWithContext(ctx context.Context) GetInstancesReservedInstanceListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstancesReservedInstanceListOutput)
}

// GetInstancesReservedInstanceListArrayInput is an input type that accepts GetInstancesReservedInstanceListArray and GetInstancesReservedInstanceListArrayOutput values.
// You can construct a concrete instance of `GetInstancesReservedInstanceListArrayInput` via:
//
//	GetInstancesReservedInstanceListArray{ GetInstancesReservedInstanceListArgs{...} }
type GetInstancesReservedInstanceListArrayInput interface {
	pulumi.Input

	ToGetInstancesReservedInstanceListArrayOutput() GetInstancesReservedInstanceListArrayOutput
	ToGetInstancesReservedInstanceListArrayOutputWithContext(context.Context) GetInstancesReservedInstanceListArrayOutput
}

type GetInstancesReservedInstanceListArray []GetInstancesReservedInstanceListInput

func (GetInstancesReservedInstanceListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstancesReservedInstanceList)(nil)).Elem()
}

func (i GetInstancesReservedInstanceListArray) ToGetInstancesReservedInstanceListArrayOutput() GetInstancesReservedInstanceListArrayOutput {
	return i.ToGetInstancesReservedInstanceListArrayOutputWithContext(context.Background())
}

func (i GetInstancesReservedInstanceListArray) ToGetInstancesReservedInstanceListArrayOutputWithContext(ctx context.Context) GetInstancesReservedInstanceListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstancesReservedInstanceListArrayOutput)
}

type GetInstancesReservedInstanceListOutput struct{ *pulumi.OutputState }

func (GetInstancesReservedInstanceListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesReservedInstanceList)(nil)).Elem()
}

func (o GetInstancesReservedInstanceListOutput) ToGetInstancesReservedInstanceListOutput() GetInstancesReservedInstanceListOutput {
	return o
}

func (o GetInstancesReservedInstanceListOutput) ToGetInstancesReservedInstanceListOutputWithContext(ctx context.Context) GetInstancesReservedInstanceListOutput {
	return o
}

// The available zone that the reserved instance locates at.
func (o GetInstancesReservedInstanceListOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesReservedInstanceList) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Expiry time of the reserved instance.
func (o GetInstancesReservedInstanceListOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesReservedInstanceList) string { return v.EndTime }).(pulumi.StringOutput)
}

// Number of reserved instance.
func (o GetInstancesReservedInstanceListOutput) InstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstancesReservedInstanceList) int { return v.InstanceCount }).(pulumi.IntOutput)
}

// The type of reserved instance.
func (o GetInstancesReservedInstanceListOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesReservedInstanceList) string { return v.InstanceType }).(pulumi.StringOutput)
}

// ID of the reserved instance to be query.
func (o GetInstancesReservedInstanceListOutput) ReservedInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesReservedInstanceList) string { return v.ReservedInstanceId }).(pulumi.StringOutput)
}

// Start time of the reserved instance.
func (o GetInstancesReservedInstanceListOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesReservedInstanceList) string { return v.StartTime }).(pulumi.StringOutput)
}

// Status of the reserved instance.
func (o GetInstancesReservedInstanceListOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesReservedInstanceList) string { return v.Status }).(pulumi.StringOutput)
}

type GetInstancesReservedInstanceListArrayOutput struct{ *pulumi.OutputState }

func (GetInstancesReservedInstanceListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstancesReservedInstanceList)(nil)).Elem()
}

func (o GetInstancesReservedInstanceListArrayOutput) ToGetInstancesReservedInstanceListArrayOutput() GetInstancesReservedInstanceListArrayOutput {
	return o
}

func (o GetInstancesReservedInstanceListArrayOutput) ToGetInstancesReservedInstanceListArrayOutputWithContext(ctx context.Context) GetInstancesReservedInstanceListArrayOutput {
	return o
}

func (o GetInstancesReservedInstanceListArrayOutput) Index(i pulumi.IntInput) GetInstancesReservedInstanceListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetInstancesReservedInstanceList {
		return vs[0].([]GetInstancesReservedInstanceList)[vs[1].(int)]
	}).(GetInstancesReservedInstanceListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetInstanceConfigsConfigListInput)(nil)).Elem(), GetInstanceConfigsConfigListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInstanceConfigsConfigListArrayInput)(nil)).Elem(), GetInstanceConfigsConfigListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInstancesReservedInstanceListInput)(nil)).Elem(), GetInstancesReservedInstanceListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInstancesReservedInstanceListArrayInput)(nil)).Elem(), GetInstancesReservedInstanceListArray{})
	pulumi.RegisterOutputType(GetInstanceConfigsConfigListOutput{})
	pulumi.RegisterOutputType(GetInstanceConfigsConfigListArrayOutput{})
	pulumi.RegisterOutputType(GetInstancesReservedInstanceListOutput{})
	pulumi.RegisterOutputType(GetInstancesReservedInstanceListArrayOutput{})
}
