// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eni

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

var _ = internal.GetEnvOrDefault

type InstanceIpv4 struct {
	// Description of the IP, maximum length 25.
	Description *string `pulumi:"description"`
	// Intranet IP.
	Ip string `pulumi:"ip"`
	// Indicates whether the IP is primary.
	Primary bool `pulumi:"primary"`
}

// InstanceIpv4Input is an input type that accepts InstanceIpv4Args and InstanceIpv4Output values.
// You can construct a concrete instance of `InstanceIpv4Input` via:
//
//	InstanceIpv4Args{...}
type InstanceIpv4Input interface {
	pulumi.Input

	ToInstanceIpv4Output() InstanceIpv4Output
	ToInstanceIpv4OutputWithContext(context.Context) InstanceIpv4Output
}

type InstanceIpv4Args struct {
	// Description of the IP, maximum length 25.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Intranet IP.
	Ip pulumi.StringInput `pulumi:"ip"`
	// Indicates whether the IP is primary.
	Primary pulumi.BoolInput `pulumi:"primary"`
}

func (InstanceIpv4Args) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIpv4)(nil)).Elem()
}

func (i InstanceIpv4Args) ToInstanceIpv4Output() InstanceIpv4Output {
	return i.ToInstanceIpv4OutputWithContext(context.Background())
}

func (i InstanceIpv4Args) ToInstanceIpv4OutputWithContext(ctx context.Context) InstanceIpv4Output {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIpv4Output)
}

// InstanceIpv4ArrayInput is an input type that accepts InstanceIpv4Array and InstanceIpv4ArrayOutput values.
// You can construct a concrete instance of `InstanceIpv4ArrayInput` via:
//
//	InstanceIpv4Array{ InstanceIpv4Args{...} }
type InstanceIpv4ArrayInput interface {
	pulumi.Input

	ToInstanceIpv4ArrayOutput() InstanceIpv4ArrayOutput
	ToInstanceIpv4ArrayOutputWithContext(context.Context) InstanceIpv4ArrayOutput
}

type InstanceIpv4Array []InstanceIpv4Input

func (InstanceIpv4Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceIpv4)(nil)).Elem()
}

func (i InstanceIpv4Array) ToInstanceIpv4ArrayOutput() InstanceIpv4ArrayOutput {
	return i.ToInstanceIpv4ArrayOutputWithContext(context.Background())
}

func (i InstanceIpv4Array) ToInstanceIpv4ArrayOutputWithContext(ctx context.Context) InstanceIpv4ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIpv4ArrayOutput)
}

type InstanceIpv4Output struct{ *pulumi.OutputState }

func (InstanceIpv4Output) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIpv4)(nil)).Elem()
}

func (o InstanceIpv4Output) ToInstanceIpv4Output() InstanceIpv4Output {
	return o
}

func (o InstanceIpv4Output) ToInstanceIpv4OutputWithContext(ctx context.Context) InstanceIpv4Output {
	return o
}

// Description of the IP, maximum length 25.
func (o InstanceIpv4Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceIpv4) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Intranet IP.
func (o InstanceIpv4Output) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceIpv4) string { return v.Ip }).(pulumi.StringOutput)
}

// Indicates whether the IP is primary.
func (o InstanceIpv4Output) Primary() pulumi.BoolOutput {
	return o.ApplyT(func(v InstanceIpv4) bool { return v.Primary }).(pulumi.BoolOutput)
}

type InstanceIpv4ArrayOutput struct{ *pulumi.OutputState }

func (InstanceIpv4ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceIpv4)(nil)).Elem()
}

func (o InstanceIpv4ArrayOutput) ToInstanceIpv4ArrayOutput() InstanceIpv4ArrayOutput {
	return o
}

func (o InstanceIpv4ArrayOutput) ToInstanceIpv4ArrayOutputWithContext(ctx context.Context) InstanceIpv4ArrayOutput {
	return o
}

func (o InstanceIpv4ArrayOutput) Index(i pulumi.IntInput) InstanceIpv4Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceIpv4 {
		return vs[0].([]InstanceIpv4)[vs[1].(int)]
	}).(InstanceIpv4Output)
}

type InstanceIpv4Info struct {
	// Description of the IP, maximum length 25.
	Description *string `pulumi:"description"`
	// Intranet IP.
	Ip *string `pulumi:"ip"`
	// Indicates whether the IP is primary.
	Primary *bool `pulumi:"primary"`
}

// InstanceIpv4InfoInput is an input type that accepts InstanceIpv4InfoArgs and InstanceIpv4InfoOutput values.
// You can construct a concrete instance of `InstanceIpv4InfoInput` via:
//
//	InstanceIpv4InfoArgs{...}
type InstanceIpv4InfoInput interface {
	pulumi.Input

	ToInstanceIpv4InfoOutput() InstanceIpv4InfoOutput
	ToInstanceIpv4InfoOutputWithContext(context.Context) InstanceIpv4InfoOutput
}

type InstanceIpv4InfoArgs struct {
	// Description of the IP, maximum length 25.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Intranet IP.
	Ip pulumi.StringPtrInput `pulumi:"ip"`
	// Indicates whether the IP is primary.
	Primary pulumi.BoolPtrInput `pulumi:"primary"`
}

func (InstanceIpv4InfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIpv4Info)(nil)).Elem()
}

func (i InstanceIpv4InfoArgs) ToInstanceIpv4InfoOutput() InstanceIpv4InfoOutput {
	return i.ToInstanceIpv4InfoOutputWithContext(context.Background())
}

func (i InstanceIpv4InfoArgs) ToInstanceIpv4InfoOutputWithContext(ctx context.Context) InstanceIpv4InfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIpv4InfoOutput)
}

// InstanceIpv4InfoArrayInput is an input type that accepts InstanceIpv4InfoArray and InstanceIpv4InfoArrayOutput values.
// You can construct a concrete instance of `InstanceIpv4InfoArrayInput` via:
//
//	InstanceIpv4InfoArray{ InstanceIpv4InfoArgs{...} }
type InstanceIpv4InfoArrayInput interface {
	pulumi.Input

	ToInstanceIpv4InfoArrayOutput() InstanceIpv4InfoArrayOutput
	ToInstanceIpv4InfoArrayOutputWithContext(context.Context) InstanceIpv4InfoArrayOutput
}

type InstanceIpv4InfoArray []InstanceIpv4InfoInput

func (InstanceIpv4InfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceIpv4Info)(nil)).Elem()
}

func (i InstanceIpv4InfoArray) ToInstanceIpv4InfoArrayOutput() InstanceIpv4InfoArrayOutput {
	return i.ToInstanceIpv4InfoArrayOutputWithContext(context.Background())
}

func (i InstanceIpv4InfoArray) ToInstanceIpv4InfoArrayOutputWithContext(ctx context.Context) InstanceIpv4InfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIpv4InfoArrayOutput)
}

type InstanceIpv4InfoOutput struct{ *pulumi.OutputState }

func (InstanceIpv4InfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIpv4Info)(nil)).Elem()
}

func (o InstanceIpv4InfoOutput) ToInstanceIpv4InfoOutput() InstanceIpv4InfoOutput {
	return o
}

func (o InstanceIpv4InfoOutput) ToInstanceIpv4InfoOutputWithContext(ctx context.Context) InstanceIpv4InfoOutput {
	return o
}

// Description of the IP, maximum length 25.
func (o InstanceIpv4InfoOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceIpv4Info) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Intranet IP.
func (o InstanceIpv4InfoOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceIpv4Info) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

// Indicates whether the IP is primary.
func (o InstanceIpv4InfoOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v InstanceIpv4Info) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

type InstanceIpv4InfoArrayOutput struct{ *pulumi.OutputState }

func (InstanceIpv4InfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceIpv4Info)(nil)).Elem()
}

func (o InstanceIpv4InfoArrayOutput) ToInstanceIpv4InfoArrayOutput() InstanceIpv4InfoArrayOutput {
	return o
}

func (o InstanceIpv4InfoArrayOutput) ToInstanceIpv4InfoArrayOutputWithContext(ctx context.Context) InstanceIpv4InfoArrayOutput {
	return o
}

func (o InstanceIpv4InfoArrayOutput) Index(i pulumi.IntInput) InstanceIpv4InfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceIpv4Info {
		return vs[0].([]InstanceIpv4Info)[vs[1].(int)]
	}).(InstanceIpv4InfoOutput)
}

type Ipv4AddressPrivateIpAddress struct {
	// EIP instance ID, such as `eip-11112222`.
	AddressId *string `pulumi:"addressId"`
	// Private IP description.
	Description *string `pulumi:"description"`
	// Whether the public IP is blocked.
	IsWanIpBlocked *bool `pulumi:"isWanIpBlocked"`
	// Whether it is a primary IP.
	Primary *bool `pulumi:"primary"`
	// Private IP address.
	PrivateIpAddress string `pulumi:"privateIpAddress"`
	// Public IP address.
	PublicIpAddress *string `pulumi:"publicIpAddress"`
	// IP service level. Values: PT`(Gold),`AU`(Silver),`AG `(Bronze) and DEFAULT` (Default).
	QosLevel *string `pulumi:"qosLevel"`
	// IP status: `PENDING`: Creating, `MIGRATING`: Migrating, `DELETING`: Deleting, `AVAILABLE`: Available.
	State *string `pulumi:"state"`
}

// Ipv4AddressPrivateIpAddressInput is an input type that accepts Ipv4AddressPrivateIpAddressArgs and Ipv4AddressPrivateIpAddressOutput values.
// You can construct a concrete instance of `Ipv4AddressPrivateIpAddressInput` via:
//
//	Ipv4AddressPrivateIpAddressArgs{...}
type Ipv4AddressPrivateIpAddressInput interface {
	pulumi.Input

	ToIpv4AddressPrivateIpAddressOutput() Ipv4AddressPrivateIpAddressOutput
	ToIpv4AddressPrivateIpAddressOutputWithContext(context.Context) Ipv4AddressPrivateIpAddressOutput
}

type Ipv4AddressPrivateIpAddressArgs struct {
	// EIP instance ID, such as `eip-11112222`.
	AddressId pulumi.StringPtrInput `pulumi:"addressId"`
	// Private IP description.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Whether the public IP is blocked.
	IsWanIpBlocked pulumi.BoolPtrInput `pulumi:"isWanIpBlocked"`
	// Whether it is a primary IP.
	Primary pulumi.BoolPtrInput `pulumi:"primary"`
	// Private IP address.
	PrivateIpAddress pulumi.StringInput `pulumi:"privateIpAddress"`
	// Public IP address.
	PublicIpAddress pulumi.StringPtrInput `pulumi:"publicIpAddress"`
	// IP service level. Values: PT`(Gold),`AU`(Silver),`AG `(Bronze) and DEFAULT` (Default).
	QosLevel pulumi.StringPtrInput `pulumi:"qosLevel"`
	// IP status: `PENDING`: Creating, `MIGRATING`: Migrating, `DELETING`: Deleting, `AVAILABLE`: Available.
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (Ipv4AddressPrivateIpAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv4AddressPrivateIpAddress)(nil)).Elem()
}

func (i Ipv4AddressPrivateIpAddressArgs) ToIpv4AddressPrivateIpAddressOutput() Ipv4AddressPrivateIpAddressOutput {
	return i.ToIpv4AddressPrivateIpAddressOutputWithContext(context.Background())
}

func (i Ipv4AddressPrivateIpAddressArgs) ToIpv4AddressPrivateIpAddressOutputWithContext(ctx context.Context) Ipv4AddressPrivateIpAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4AddressPrivateIpAddressOutput)
}

// Ipv4AddressPrivateIpAddressArrayInput is an input type that accepts Ipv4AddressPrivateIpAddressArray and Ipv4AddressPrivateIpAddressArrayOutput values.
// You can construct a concrete instance of `Ipv4AddressPrivateIpAddressArrayInput` via:
//
//	Ipv4AddressPrivateIpAddressArray{ Ipv4AddressPrivateIpAddressArgs{...} }
type Ipv4AddressPrivateIpAddressArrayInput interface {
	pulumi.Input

	ToIpv4AddressPrivateIpAddressArrayOutput() Ipv4AddressPrivateIpAddressArrayOutput
	ToIpv4AddressPrivateIpAddressArrayOutputWithContext(context.Context) Ipv4AddressPrivateIpAddressArrayOutput
}

type Ipv4AddressPrivateIpAddressArray []Ipv4AddressPrivateIpAddressInput

func (Ipv4AddressPrivateIpAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ipv4AddressPrivateIpAddress)(nil)).Elem()
}

func (i Ipv4AddressPrivateIpAddressArray) ToIpv4AddressPrivateIpAddressArrayOutput() Ipv4AddressPrivateIpAddressArrayOutput {
	return i.ToIpv4AddressPrivateIpAddressArrayOutputWithContext(context.Background())
}

func (i Ipv4AddressPrivateIpAddressArray) ToIpv4AddressPrivateIpAddressArrayOutputWithContext(ctx context.Context) Ipv4AddressPrivateIpAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4AddressPrivateIpAddressArrayOutput)
}

type Ipv4AddressPrivateIpAddressOutput struct{ *pulumi.OutputState }

func (Ipv4AddressPrivateIpAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv4AddressPrivateIpAddress)(nil)).Elem()
}

func (o Ipv4AddressPrivateIpAddressOutput) ToIpv4AddressPrivateIpAddressOutput() Ipv4AddressPrivateIpAddressOutput {
	return o
}

func (o Ipv4AddressPrivateIpAddressOutput) ToIpv4AddressPrivateIpAddressOutputWithContext(ctx context.Context) Ipv4AddressPrivateIpAddressOutput {
	return o
}

// EIP instance ID, such as `eip-11112222`.
func (o Ipv4AddressPrivateIpAddressOutput) AddressId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv4AddressPrivateIpAddress) *string { return v.AddressId }).(pulumi.StringPtrOutput)
}

// Private IP description.
func (o Ipv4AddressPrivateIpAddressOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv4AddressPrivateIpAddress) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether the public IP is blocked.
func (o Ipv4AddressPrivateIpAddressOutput) IsWanIpBlocked() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Ipv4AddressPrivateIpAddress) *bool { return v.IsWanIpBlocked }).(pulumi.BoolPtrOutput)
}

// Whether it is a primary IP.
func (o Ipv4AddressPrivateIpAddressOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Ipv4AddressPrivateIpAddress) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

// Private IP address.
func (o Ipv4AddressPrivateIpAddressOutput) PrivateIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v Ipv4AddressPrivateIpAddress) string { return v.PrivateIpAddress }).(pulumi.StringOutput)
}

// Public IP address.
func (o Ipv4AddressPrivateIpAddressOutput) PublicIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv4AddressPrivateIpAddress) *string { return v.PublicIpAddress }).(pulumi.StringPtrOutput)
}

// IP service level. Values: PT`(Gold),`AU`(Silver),`AG `(Bronze) and DEFAULT` (Default).
func (o Ipv4AddressPrivateIpAddressOutput) QosLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv4AddressPrivateIpAddress) *string { return v.QosLevel }).(pulumi.StringPtrOutput)
}

// IP status: `PENDING`: Creating, `MIGRATING`: Migrating, `DELETING`: Deleting, `AVAILABLE`: Available.
func (o Ipv4AddressPrivateIpAddressOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv4AddressPrivateIpAddress) *string { return v.State }).(pulumi.StringPtrOutput)
}

type Ipv4AddressPrivateIpAddressArrayOutput struct{ *pulumi.OutputState }

func (Ipv4AddressPrivateIpAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ipv4AddressPrivateIpAddress)(nil)).Elem()
}

func (o Ipv4AddressPrivateIpAddressArrayOutput) ToIpv4AddressPrivateIpAddressArrayOutput() Ipv4AddressPrivateIpAddressArrayOutput {
	return o
}

func (o Ipv4AddressPrivateIpAddressArrayOutput) ToIpv4AddressPrivateIpAddressArrayOutputWithContext(ctx context.Context) Ipv4AddressPrivateIpAddressArrayOutput {
	return o
}

func (o Ipv4AddressPrivateIpAddressArrayOutput) Index(i pulumi.IntInput) Ipv4AddressPrivateIpAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ipv4AddressPrivateIpAddress {
		return vs[0].([]Ipv4AddressPrivateIpAddress)[vs[1].(int)]
	}).(Ipv4AddressPrivateIpAddressOutput)
}

type Ipv6AddressIpv6Address struct {
	// `IPv6` address, in the form of: `3402:4e00:20:100:0:8cd9:2a67:71f3`.
	Address string `pulumi:"address"`
	// `EIP` instance `ID`, such as:`eip-hxlqja90`.
	AddressId *string `pulumi:"addressId"`
	// Description.
	Description *string `pulumi:"description"`
	// Whether the public network IP is blocked.
	IsWanIpBlocked *bool `pulumi:"isWanIpBlocked"`
	// Whether to master `IP`.
	Primary *bool `pulumi:"primary"`
	// `IPv6` address status: `PENDING`: pending, `MIGRATING`: migrating, `DELETING`: deleting, `AVAILABLE`: available.
	State *string `pulumi:"state"`
}

// Ipv6AddressIpv6AddressInput is an input type that accepts Ipv6AddressIpv6AddressArgs and Ipv6AddressIpv6AddressOutput values.
// You can construct a concrete instance of `Ipv6AddressIpv6AddressInput` via:
//
//	Ipv6AddressIpv6AddressArgs{...}
type Ipv6AddressIpv6AddressInput interface {
	pulumi.Input

	ToIpv6AddressIpv6AddressOutput() Ipv6AddressIpv6AddressOutput
	ToIpv6AddressIpv6AddressOutputWithContext(context.Context) Ipv6AddressIpv6AddressOutput
}

type Ipv6AddressIpv6AddressArgs struct {
	// `IPv6` address, in the form of: `3402:4e00:20:100:0:8cd9:2a67:71f3`.
	Address pulumi.StringInput `pulumi:"address"`
	// `EIP` instance `ID`, such as:`eip-hxlqja90`.
	AddressId pulumi.StringPtrInput `pulumi:"addressId"`
	// Description.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Whether the public network IP is blocked.
	IsWanIpBlocked pulumi.BoolPtrInput `pulumi:"isWanIpBlocked"`
	// Whether to master `IP`.
	Primary pulumi.BoolPtrInput `pulumi:"primary"`
	// `IPv6` address status: `PENDING`: pending, `MIGRATING`: migrating, `DELETING`: deleting, `AVAILABLE`: available.
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (Ipv6AddressIpv6AddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv6AddressIpv6Address)(nil)).Elem()
}

func (i Ipv6AddressIpv6AddressArgs) ToIpv6AddressIpv6AddressOutput() Ipv6AddressIpv6AddressOutput {
	return i.ToIpv6AddressIpv6AddressOutputWithContext(context.Background())
}

func (i Ipv6AddressIpv6AddressArgs) ToIpv6AddressIpv6AddressOutputWithContext(ctx context.Context) Ipv6AddressIpv6AddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6AddressIpv6AddressOutput)
}

// Ipv6AddressIpv6AddressArrayInput is an input type that accepts Ipv6AddressIpv6AddressArray and Ipv6AddressIpv6AddressArrayOutput values.
// You can construct a concrete instance of `Ipv6AddressIpv6AddressArrayInput` via:
//
//	Ipv6AddressIpv6AddressArray{ Ipv6AddressIpv6AddressArgs{...} }
type Ipv6AddressIpv6AddressArrayInput interface {
	pulumi.Input

	ToIpv6AddressIpv6AddressArrayOutput() Ipv6AddressIpv6AddressArrayOutput
	ToIpv6AddressIpv6AddressArrayOutputWithContext(context.Context) Ipv6AddressIpv6AddressArrayOutput
}

type Ipv6AddressIpv6AddressArray []Ipv6AddressIpv6AddressInput

func (Ipv6AddressIpv6AddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ipv6AddressIpv6Address)(nil)).Elem()
}

func (i Ipv6AddressIpv6AddressArray) ToIpv6AddressIpv6AddressArrayOutput() Ipv6AddressIpv6AddressArrayOutput {
	return i.ToIpv6AddressIpv6AddressArrayOutputWithContext(context.Background())
}

func (i Ipv6AddressIpv6AddressArray) ToIpv6AddressIpv6AddressArrayOutputWithContext(ctx context.Context) Ipv6AddressIpv6AddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6AddressIpv6AddressArrayOutput)
}

type Ipv6AddressIpv6AddressOutput struct{ *pulumi.OutputState }

func (Ipv6AddressIpv6AddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv6AddressIpv6Address)(nil)).Elem()
}

func (o Ipv6AddressIpv6AddressOutput) ToIpv6AddressIpv6AddressOutput() Ipv6AddressIpv6AddressOutput {
	return o
}

func (o Ipv6AddressIpv6AddressOutput) ToIpv6AddressIpv6AddressOutputWithContext(ctx context.Context) Ipv6AddressIpv6AddressOutput {
	return o
}

// `IPv6` address, in the form of: `3402:4e00:20:100:0:8cd9:2a67:71f3`.
func (o Ipv6AddressIpv6AddressOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v Ipv6AddressIpv6Address) string { return v.Address }).(pulumi.StringOutput)
}

// `EIP` instance `ID`, such as:`eip-hxlqja90`.
func (o Ipv6AddressIpv6AddressOutput) AddressId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6AddressIpv6Address) *string { return v.AddressId }).(pulumi.StringPtrOutput)
}

// Description.
func (o Ipv6AddressIpv6AddressOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6AddressIpv6Address) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether the public network IP is blocked.
func (o Ipv6AddressIpv6AddressOutput) IsWanIpBlocked() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Ipv6AddressIpv6Address) *bool { return v.IsWanIpBlocked }).(pulumi.BoolPtrOutput)
}

// Whether to master `IP`.
func (o Ipv6AddressIpv6AddressOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Ipv6AddressIpv6Address) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

// `IPv6` address status: `PENDING`: pending, `MIGRATING`: migrating, `DELETING`: deleting, `AVAILABLE`: available.
func (o Ipv6AddressIpv6AddressOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6AddressIpv6Address) *string { return v.State }).(pulumi.StringPtrOutput)
}

type Ipv6AddressIpv6AddressArrayOutput struct{ *pulumi.OutputState }

func (Ipv6AddressIpv6AddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ipv6AddressIpv6Address)(nil)).Elem()
}

func (o Ipv6AddressIpv6AddressArrayOutput) ToIpv6AddressIpv6AddressArrayOutput() Ipv6AddressIpv6AddressArrayOutput {
	return o
}

func (o Ipv6AddressIpv6AddressArrayOutput) ToIpv6AddressIpv6AddressArrayOutputWithContext(ctx context.Context) Ipv6AddressIpv6AddressArrayOutput {
	return o
}

func (o Ipv6AddressIpv6AddressArrayOutput) Index(i pulumi.IntInput) Ipv6AddressIpv6AddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ipv6AddressIpv6Address {
		return vs[0].([]Ipv6AddressIpv6Address)[vs[1].(int)]
	}).(Ipv6AddressIpv6AddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIpv4Input)(nil)).Elem(), InstanceIpv4Args{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIpv4ArrayInput)(nil)).Elem(), InstanceIpv4Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIpv4InfoInput)(nil)).Elem(), InstanceIpv4InfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIpv4InfoArrayInput)(nil)).Elem(), InstanceIpv4InfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4AddressPrivateIpAddressInput)(nil)).Elem(), Ipv4AddressPrivateIpAddressArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4AddressPrivateIpAddressArrayInput)(nil)).Elem(), Ipv4AddressPrivateIpAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6AddressIpv6AddressInput)(nil)).Elem(), Ipv6AddressIpv6AddressArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6AddressIpv6AddressArrayInput)(nil)).Elem(), Ipv6AddressIpv6AddressArray{})
	pulumi.RegisterOutputType(InstanceIpv4Output{})
	pulumi.RegisterOutputType(InstanceIpv4ArrayOutput{})
	pulumi.RegisterOutputType(InstanceIpv4InfoOutput{})
	pulumi.RegisterOutputType(InstanceIpv4InfoArrayOutput{})
	pulumi.RegisterOutputType(Ipv4AddressPrivateIpAddressOutput{})
	pulumi.RegisterOutputType(Ipv4AddressPrivateIpAddressArrayOutput{})
	pulumi.RegisterOutputType(Ipv6AddressIpv6AddressOutput{})
	pulumi.RegisterOutputType(Ipv6AddressIpv6AddressArrayOutput{})
}
