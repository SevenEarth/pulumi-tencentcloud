// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a vpc ipv6EniAddress
type Ipv6EniAddress struct {
	pulumi.CustomResourceState

	// The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
	Ipv6Addresses Ipv6EniAddressIpv6AddressArrayOutput `pulumi:"ipv6Addresses"`
	// ENI instance `ID`, in the form of `eni-m6dyj72l`.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// VPC `ID`, in the form of `vpc-m6dyj72l`.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewIpv6EniAddress registers a new resource with the given unique name, arguments, and options.
func NewIpv6EniAddress(ctx *pulumi.Context,
	name string, args *Ipv6EniAddressArgs, opts ...pulumi.ResourceOption) (*Ipv6EniAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkInterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaceId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Ipv6EniAddress
	err := ctx.RegisterResource("tencentcloud:Vpc/ipv6EniAddress:Ipv6EniAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpv6EniAddress gets an existing Ipv6EniAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpv6EniAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ipv6EniAddressState, opts ...pulumi.ResourceOption) (*Ipv6EniAddress, error) {
	var resource Ipv6EniAddress
	err := ctx.ReadResource("tencentcloud:Vpc/ipv6EniAddress:Ipv6EniAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ipv6EniAddress resources.
type ipv6EniAddressState struct {
	// The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
	Ipv6Addresses []Ipv6EniAddressIpv6Address `pulumi:"ipv6Addresses"`
	// ENI instance `ID`, in the form of `eni-m6dyj72l`.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// VPC `ID`, in the form of `vpc-m6dyj72l`.
	VpcId *string `pulumi:"vpcId"`
}

type Ipv6EniAddressState struct {
	// The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
	Ipv6Addresses Ipv6EniAddressIpv6AddressArrayInput
	// ENI instance `ID`, in the form of `eni-m6dyj72l`.
	NetworkInterfaceId pulumi.StringPtrInput
	// VPC `ID`, in the form of `vpc-m6dyj72l`.
	VpcId pulumi.StringPtrInput
}

func (Ipv6EniAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv6EniAddressState)(nil)).Elem()
}

type ipv6EniAddressArgs struct {
	// The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
	Ipv6Addresses []Ipv6EniAddressIpv6Address `pulumi:"ipv6Addresses"`
	// ENI instance `ID`, in the form of `eni-m6dyj72l`.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// VPC `ID`, in the form of `vpc-m6dyj72l`.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Ipv6EniAddress resource.
type Ipv6EniAddressArgs struct {
	// The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
	Ipv6Addresses Ipv6EniAddressIpv6AddressArrayInput
	// ENI instance `ID`, in the form of `eni-m6dyj72l`.
	NetworkInterfaceId pulumi.StringInput
	// VPC `ID`, in the form of `vpc-m6dyj72l`.
	VpcId pulumi.StringInput
}

func (Ipv6EniAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv6EniAddressArgs)(nil)).Elem()
}

type Ipv6EniAddressInput interface {
	pulumi.Input

	ToIpv6EniAddressOutput() Ipv6EniAddressOutput
	ToIpv6EniAddressOutputWithContext(ctx context.Context) Ipv6EniAddressOutput
}

func (*Ipv6EniAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6EniAddress)(nil)).Elem()
}

func (i *Ipv6EniAddress) ToIpv6EniAddressOutput() Ipv6EniAddressOutput {
	return i.ToIpv6EniAddressOutputWithContext(context.Background())
}

func (i *Ipv6EniAddress) ToIpv6EniAddressOutputWithContext(ctx context.Context) Ipv6EniAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6EniAddressOutput)
}

// Ipv6EniAddressArrayInput is an input type that accepts Ipv6EniAddressArray and Ipv6EniAddressArrayOutput values.
// You can construct a concrete instance of `Ipv6EniAddressArrayInput` via:
//
//	Ipv6EniAddressArray{ Ipv6EniAddressArgs{...} }
type Ipv6EniAddressArrayInput interface {
	pulumi.Input

	ToIpv6EniAddressArrayOutput() Ipv6EniAddressArrayOutput
	ToIpv6EniAddressArrayOutputWithContext(context.Context) Ipv6EniAddressArrayOutput
}

type Ipv6EniAddressArray []Ipv6EniAddressInput

func (Ipv6EniAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv6EniAddress)(nil)).Elem()
}

func (i Ipv6EniAddressArray) ToIpv6EniAddressArrayOutput() Ipv6EniAddressArrayOutput {
	return i.ToIpv6EniAddressArrayOutputWithContext(context.Background())
}

func (i Ipv6EniAddressArray) ToIpv6EniAddressArrayOutputWithContext(ctx context.Context) Ipv6EniAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6EniAddressArrayOutput)
}

// Ipv6EniAddressMapInput is an input type that accepts Ipv6EniAddressMap and Ipv6EniAddressMapOutput values.
// You can construct a concrete instance of `Ipv6EniAddressMapInput` via:
//
//	Ipv6EniAddressMap{ "key": Ipv6EniAddressArgs{...} }
type Ipv6EniAddressMapInput interface {
	pulumi.Input

	ToIpv6EniAddressMapOutput() Ipv6EniAddressMapOutput
	ToIpv6EniAddressMapOutputWithContext(context.Context) Ipv6EniAddressMapOutput
}

type Ipv6EniAddressMap map[string]Ipv6EniAddressInput

func (Ipv6EniAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv6EniAddress)(nil)).Elem()
}

func (i Ipv6EniAddressMap) ToIpv6EniAddressMapOutput() Ipv6EniAddressMapOutput {
	return i.ToIpv6EniAddressMapOutputWithContext(context.Background())
}

func (i Ipv6EniAddressMap) ToIpv6EniAddressMapOutputWithContext(ctx context.Context) Ipv6EniAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6EniAddressMapOutput)
}

type Ipv6EniAddressOutput struct{ *pulumi.OutputState }

func (Ipv6EniAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6EniAddress)(nil)).Elem()
}

func (o Ipv6EniAddressOutput) ToIpv6EniAddressOutput() Ipv6EniAddressOutput {
	return o
}

func (o Ipv6EniAddressOutput) ToIpv6EniAddressOutputWithContext(ctx context.Context) Ipv6EniAddressOutput {
	return o
}

// The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
func (o Ipv6EniAddressOutput) Ipv6Addresses() Ipv6EniAddressIpv6AddressArrayOutput {
	return o.ApplyT(func(v *Ipv6EniAddress) Ipv6EniAddressIpv6AddressArrayOutput { return v.Ipv6Addresses }).(Ipv6EniAddressIpv6AddressArrayOutput)
}

// ENI instance `ID`, in the form of `eni-m6dyj72l`.
func (o Ipv6EniAddressOutput) NetworkInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6EniAddress) pulumi.StringOutput { return v.NetworkInterfaceId }).(pulumi.StringOutput)
}

// VPC `ID`, in the form of `vpc-m6dyj72l`.
func (o Ipv6EniAddressOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6EniAddress) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type Ipv6EniAddressArrayOutput struct{ *pulumi.OutputState }

func (Ipv6EniAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv6EniAddress)(nil)).Elem()
}

func (o Ipv6EniAddressArrayOutput) ToIpv6EniAddressArrayOutput() Ipv6EniAddressArrayOutput {
	return o
}

func (o Ipv6EniAddressArrayOutput) ToIpv6EniAddressArrayOutputWithContext(ctx context.Context) Ipv6EniAddressArrayOutput {
	return o
}

func (o Ipv6EniAddressArrayOutput) Index(i pulumi.IntInput) Ipv6EniAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ipv6EniAddress {
		return vs[0].([]*Ipv6EniAddress)[vs[1].(int)]
	}).(Ipv6EniAddressOutput)
}

type Ipv6EniAddressMapOutput struct{ *pulumi.OutputState }

func (Ipv6EniAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv6EniAddress)(nil)).Elem()
}

func (o Ipv6EniAddressMapOutput) ToIpv6EniAddressMapOutput() Ipv6EniAddressMapOutput {
	return o
}

func (o Ipv6EniAddressMapOutput) ToIpv6EniAddressMapOutputWithContext(ctx context.Context) Ipv6EniAddressMapOutput {
	return o
}

func (o Ipv6EniAddressMapOutput) MapIndex(k pulumi.StringInput) Ipv6EniAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ipv6EniAddress {
		return vs[0].(map[string]*Ipv6EniAddress)[vs[1].(string)]
	}).(Ipv6EniAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6EniAddressInput)(nil)).Elem(), &Ipv6EniAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6EniAddressArrayInput)(nil)).Elem(), Ipv6EniAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6EniAddressMapInput)(nil)).Elem(), Ipv6EniAddressMap{})
	pulumi.RegisterOutputType(Ipv6EniAddressOutput{})
	pulumi.RegisterOutputType(Ipv6EniAddressArrayOutput{})
	pulumi.RegisterOutputType(Ipv6EniAddressMapOutput{})
}
