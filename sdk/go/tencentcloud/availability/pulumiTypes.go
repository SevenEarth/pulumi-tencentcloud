// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package availability

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetRegionsRegion struct {
	// The description of the region, like `Guangzhou Region`.
	Description string `pulumi:"description"`
	// When specified, only the region with the exactly name match will be returned. `default` value means it consistent with the provider region.
	Name string `pulumi:"name"`
	// The state of the region, indicate availability using `AVAILABLE` and `UNAVAILABLE` values.
	State string `pulumi:"state"`
}

// GetRegionsRegionInput is an input type that accepts GetRegionsRegionArgs and GetRegionsRegionOutput values.
// You can construct a concrete instance of `GetRegionsRegionInput` via:
//
//          GetRegionsRegionArgs{...}
type GetRegionsRegionInput interface {
	pulumi.Input

	ToGetRegionsRegionOutput() GetRegionsRegionOutput
	ToGetRegionsRegionOutputWithContext(context.Context) GetRegionsRegionOutput
}

type GetRegionsRegionArgs struct {
	// The description of the region, like `Guangzhou Region`.
	Description pulumi.StringInput `pulumi:"description"`
	// When specified, only the region with the exactly name match will be returned. `default` value means it consistent with the provider region.
	Name pulumi.StringInput `pulumi:"name"`
	// The state of the region, indicate availability using `AVAILABLE` and `UNAVAILABLE` values.
	State pulumi.StringInput `pulumi:"state"`
}

func (GetRegionsRegionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegionsRegion)(nil)).Elem()
}

func (i GetRegionsRegionArgs) ToGetRegionsRegionOutput() GetRegionsRegionOutput {
	return i.ToGetRegionsRegionOutputWithContext(context.Background())
}

func (i GetRegionsRegionArgs) ToGetRegionsRegionOutputWithContext(ctx context.Context) GetRegionsRegionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetRegionsRegionOutput)
}

// GetRegionsRegionArrayInput is an input type that accepts GetRegionsRegionArray and GetRegionsRegionArrayOutput values.
// You can construct a concrete instance of `GetRegionsRegionArrayInput` via:
//
//          GetRegionsRegionArray{ GetRegionsRegionArgs{...} }
type GetRegionsRegionArrayInput interface {
	pulumi.Input

	ToGetRegionsRegionArrayOutput() GetRegionsRegionArrayOutput
	ToGetRegionsRegionArrayOutputWithContext(context.Context) GetRegionsRegionArrayOutput
}

type GetRegionsRegionArray []GetRegionsRegionInput

func (GetRegionsRegionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetRegionsRegion)(nil)).Elem()
}

func (i GetRegionsRegionArray) ToGetRegionsRegionArrayOutput() GetRegionsRegionArrayOutput {
	return i.ToGetRegionsRegionArrayOutputWithContext(context.Background())
}

func (i GetRegionsRegionArray) ToGetRegionsRegionArrayOutputWithContext(ctx context.Context) GetRegionsRegionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetRegionsRegionArrayOutput)
}

type GetRegionsRegionOutput struct{ *pulumi.OutputState }

func (GetRegionsRegionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegionsRegion)(nil)).Elem()
}

func (o GetRegionsRegionOutput) ToGetRegionsRegionOutput() GetRegionsRegionOutput {
	return o
}

func (o GetRegionsRegionOutput) ToGetRegionsRegionOutputWithContext(ctx context.Context) GetRegionsRegionOutput {
	return o
}

// The description of the region, like `Guangzhou Region`.
func (o GetRegionsRegionOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionsRegion) string { return v.Description }).(pulumi.StringOutput)
}

// When specified, only the region with the exactly name match will be returned. `default` value means it consistent with the provider region.
func (o GetRegionsRegionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionsRegion) string { return v.Name }).(pulumi.StringOutput)
}

// The state of the region, indicate availability using `AVAILABLE` and `UNAVAILABLE` values.
func (o GetRegionsRegionOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionsRegion) string { return v.State }).(pulumi.StringOutput)
}

type GetRegionsRegionArrayOutput struct{ *pulumi.OutputState }

func (GetRegionsRegionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetRegionsRegion)(nil)).Elem()
}

func (o GetRegionsRegionArrayOutput) ToGetRegionsRegionArrayOutput() GetRegionsRegionArrayOutput {
	return o
}

func (o GetRegionsRegionArrayOutput) ToGetRegionsRegionArrayOutputWithContext(ctx context.Context) GetRegionsRegionArrayOutput {
	return o
}

func (o GetRegionsRegionArrayOutput) Index(i pulumi.IntInput) GetRegionsRegionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetRegionsRegion {
		return vs[0].([]GetRegionsRegion)[vs[1].(int)]
	}).(GetRegionsRegionOutput)
}

type GetZonesByProductZone struct {
	// The description of the zone, like `Guangzhou Zone 3`.
	Description string `pulumi:"description"`
	// An internal id for the zone, like `200003`, usually not so useful.
	Id string `pulumi:"id"`
	// When specified, only the zone with the exactly name match will be returned.
	Name string `pulumi:"name"`
	// The state of the zone, indicate availability using `AVAILABLE` and `UNAVAILABLE` values.
	State string `pulumi:"state"`
}

// GetZonesByProductZoneInput is an input type that accepts GetZonesByProductZoneArgs and GetZonesByProductZoneOutput values.
// You can construct a concrete instance of `GetZonesByProductZoneInput` via:
//
//          GetZonesByProductZoneArgs{...}
type GetZonesByProductZoneInput interface {
	pulumi.Input

	ToGetZonesByProductZoneOutput() GetZonesByProductZoneOutput
	ToGetZonesByProductZoneOutputWithContext(context.Context) GetZonesByProductZoneOutput
}

type GetZonesByProductZoneArgs struct {
	// The description of the zone, like `Guangzhou Zone 3`.
	Description pulumi.StringInput `pulumi:"description"`
	// An internal id for the zone, like `200003`, usually not so useful.
	Id pulumi.StringInput `pulumi:"id"`
	// When specified, only the zone with the exactly name match will be returned.
	Name pulumi.StringInput `pulumi:"name"`
	// The state of the zone, indicate availability using `AVAILABLE` and `UNAVAILABLE` values.
	State pulumi.StringInput `pulumi:"state"`
}

func (GetZonesByProductZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZonesByProductZone)(nil)).Elem()
}

func (i GetZonesByProductZoneArgs) ToGetZonesByProductZoneOutput() GetZonesByProductZoneOutput {
	return i.ToGetZonesByProductZoneOutputWithContext(context.Background())
}

func (i GetZonesByProductZoneArgs) ToGetZonesByProductZoneOutputWithContext(ctx context.Context) GetZonesByProductZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetZonesByProductZoneOutput)
}

// GetZonesByProductZoneArrayInput is an input type that accepts GetZonesByProductZoneArray and GetZonesByProductZoneArrayOutput values.
// You can construct a concrete instance of `GetZonesByProductZoneArrayInput` via:
//
//          GetZonesByProductZoneArray{ GetZonesByProductZoneArgs{...} }
type GetZonesByProductZoneArrayInput interface {
	pulumi.Input

	ToGetZonesByProductZoneArrayOutput() GetZonesByProductZoneArrayOutput
	ToGetZonesByProductZoneArrayOutputWithContext(context.Context) GetZonesByProductZoneArrayOutput
}

type GetZonesByProductZoneArray []GetZonesByProductZoneInput

func (GetZonesByProductZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetZonesByProductZone)(nil)).Elem()
}

func (i GetZonesByProductZoneArray) ToGetZonesByProductZoneArrayOutput() GetZonesByProductZoneArrayOutput {
	return i.ToGetZonesByProductZoneArrayOutputWithContext(context.Background())
}

func (i GetZonesByProductZoneArray) ToGetZonesByProductZoneArrayOutputWithContext(ctx context.Context) GetZonesByProductZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetZonesByProductZoneArrayOutput)
}

type GetZonesByProductZoneOutput struct{ *pulumi.OutputState }

func (GetZonesByProductZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZonesByProductZone)(nil)).Elem()
}

func (o GetZonesByProductZoneOutput) ToGetZonesByProductZoneOutput() GetZonesByProductZoneOutput {
	return o
}

func (o GetZonesByProductZoneOutput) ToGetZonesByProductZoneOutputWithContext(ctx context.Context) GetZonesByProductZoneOutput {
	return o
}

// The description of the zone, like `Guangzhou Zone 3`.
func (o GetZonesByProductZoneOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetZonesByProductZone) string { return v.Description }).(pulumi.StringOutput)
}

// An internal id for the zone, like `200003`, usually not so useful.
func (o GetZonesByProductZoneOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetZonesByProductZone) string { return v.Id }).(pulumi.StringOutput)
}

// When specified, only the zone with the exactly name match will be returned.
func (o GetZonesByProductZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetZonesByProductZone) string { return v.Name }).(pulumi.StringOutput)
}

// The state of the zone, indicate availability using `AVAILABLE` and `UNAVAILABLE` values.
func (o GetZonesByProductZoneOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetZonesByProductZone) string { return v.State }).(pulumi.StringOutput)
}

type GetZonesByProductZoneArrayOutput struct{ *pulumi.OutputState }

func (GetZonesByProductZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetZonesByProductZone)(nil)).Elem()
}

func (o GetZonesByProductZoneArrayOutput) ToGetZonesByProductZoneArrayOutput() GetZonesByProductZoneArrayOutput {
	return o
}

func (o GetZonesByProductZoneArrayOutput) ToGetZonesByProductZoneArrayOutputWithContext(ctx context.Context) GetZonesByProductZoneArrayOutput {
	return o
}

func (o GetZonesByProductZoneArrayOutput) Index(i pulumi.IntInput) GetZonesByProductZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetZonesByProductZone {
		return vs[0].([]GetZonesByProductZone)[vs[1].(int)]
	}).(GetZonesByProductZoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetRegionsRegionInput)(nil)).Elem(), GetRegionsRegionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetRegionsRegionArrayInput)(nil)).Elem(), GetRegionsRegionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetZonesByProductZoneInput)(nil)).Elem(), GetZonesByProductZoneArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetZonesByProductZoneArrayInput)(nil)).Elem(), GetZonesByProductZoneArray{})
	pulumi.RegisterOutputType(GetRegionsRegionOutput{})
	pulumi.RegisterOutputType(GetRegionsRegionArrayOutput{})
	pulumi.RegisterOutputType(GetZonesByProductZoneOutput{})
	pulumi.RegisterOutputType(GetZonesByProductZoneArrayOutput{})
}
