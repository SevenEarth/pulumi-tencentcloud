// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

type DhcpAssociateAddress struct {
	pulumi.CustomResourceState

	// Elastic public network `IP`. Must be `EIP` not bound to `DhcpIp`.
	AddressIp pulumi.StringOutput `pulumi:"addressIp"`
	// `DhcpIp` unique `ID`, like: `dhcpip-9o233uri`. Must be a `DhcpIp` that is not bound to `EIP`.
	DhcpIpId pulumi.StringOutput `pulumi:"dhcpIpId"`
}

// NewDhcpAssociateAddress registers a new resource with the given unique name, arguments, and options.
func NewDhcpAssociateAddress(ctx *pulumi.Context,
	name string, args *DhcpAssociateAddressArgs, opts ...pulumi.ResourceOption) (*DhcpAssociateAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressIp == nil {
		return nil, errors.New("invalid value for required argument 'AddressIp'")
	}
	if args.DhcpIpId == nil {
		return nil, errors.New("invalid value for required argument 'DhcpIpId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DhcpAssociateAddress
	err := ctx.RegisterResource("tencentcloud:Vpc/dhcpAssociateAddress:DhcpAssociateAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDhcpAssociateAddress gets an existing DhcpAssociateAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDhcpAssociateAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DhcpAssociateAddressState, opts ...pulumi.ResourceOption) (*DhcpAssociateAddress, error) {
	var resource DhcpAssociateAddress
	err := ctx.ReadResource("tencentcloud:Vpc/dhcpAssociateAddress:DhcpAssociateAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DhcpAssociateAddress resources.
type dhcpAssociateAddressState struct {
	// Elastic public network `IP`. Must be `EIP` not bound to `DhcpIp`.
	AddressIp *string `pulumi:"addressIp"`
	// `DhcpIp` unique `ID`, like: `dhcpip-9o233uri`. Must be a `DhcpIp` that is not bound to `EIP`.
	DhcpIpId *string `pulumi:"dhcpIpId"`
}

type DhcpAssociateAddressState struct {
	// Elastic public network `IP`. Must be `EIP` not bound to `DhcpIp`.
	AddressIp pulumi.StringPtrInput
	// `DhcpIp` unique `ID`, like: `dhcpip-9o233uri`. Must be a `DhcpIp` that is not bound to `EIP`.
	DhcpIpId pulumi.StringPtrInput
}

func (DhcpAssociateAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpAssociateAddressState)(nil)).Elem()
}

type dhcpAssociateAddressArgs struct {
	// Elastic public network `IP`. Must be `EIP` not bound to `DhcpIp`.
	AddressIp string `pulumi:"addressIp"`
	// `DhcpIp` unique `ID`, like: `dhcpip-9o233uri`. Must be a `DhcpIp` that is not bound to `EIP`.
	DhcpIpId string `pulumi:"dhcpIpId"`
}

// The set of arguments for constructing a DhcpAssociateAddress resource.
type DhcpAssociateAddressArgs struct {
	// Elastic public network `IP`. Must be `EIP` not bound to `DhcpIp`.
	AddressIp pulumi.StringInput
	// `DhcpIp` unique `ID`, like: `dhcpip-9o233uri`. Must be a `DhcpIp` that is not bound to `EIP`.
	DhcpIpId pulumi.StringInput
}

func (DhcpAssociateAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpAssociateAddressArgs)(nil)).Elem()
}

type DhcpAssociateAddressInput interface {
	pulumi.Input

	ToDhcpAssociateAddressOutput() DhcpAssociateAddressOutput
	ToDhcpAssociateAddressOutputWithContext(ctx context.Context) DhcpAssociateAddressOutput
}

func (*DhcpAssociateAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpAssociateAddress)(nil)).Elem()
}

func (i *DhcpAssociateAddress) ToDhcpAssociateAddressOutput() DhcpAssociateAddressOutput {
	return i.ToDhcpAssociateAddressOutputWithContext(context.Background())
}

func (i *DhcpAssociateAddress) ToDhcpAssociateAddressOutputWithContext(ctx context.Context) DhcpAssociateAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpAssociateAddressOutput)
}

// DhcpAssociateAddressArrayInput is an input type that accepts DhcpAssociateAddressArray and DhcpAssociateAddressArrayOutput values.
// You can construct a concrete instance of `DhcpAssociateAddressArrayInput` via:
//
//	DhcpAssociateAddressArray{ DhcpAssociateAddressArgs{...} }
type DhcpAssociateAddressArrayInput interface {
	pulumi.Input

	ToDhcpAssociateAddressArrayOutput() DhcpAssociateAddressArrayOutput
	ToDhcpAssociateAddressArrayOutputWithContext(context.Context) DhcpAssociateAddressArrayOutput
}

type DhcpAssociateAddressArray []DhcpAssociateAddressInput

func (DhcpAssociateAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpAssociateAddress)(nil)).Elem()
}

func (i DhcpAssociateAddressArray) ToDhcpAssociateAddressArrayOutput() DhcpAssociateAddressArrayOutput {
	return i.ToDhcpAssociateAddressArrayOutputWithContext(context.Background())
}

func (i DhcpAssociateAddressArray) ToDhcpAssociateAddressArrayOutputWithContext(ctx context.Context) DhcpAssociateAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpAssociateAddressArrayOutput)
}

// DhcpAssociateAddressMapInput is an input type that accepts DhcpAssociateAddressMap and DhcpAssociateAddressMapOutput values.
// You can construct a concrete instance of `DhcpAssociateAddressMapInput` via:
//
//	DhcpAssociateAddressMap{ "key": DhcpAssociateAddressArgs{...} }
type DhcpAssociateAddressMapInput interface {
	pulumi.Input

	ToDhcpAssociateAddressMapOutput() DhcpAssociateAddressMapOutput
	ToDhcpAssociateAddressMapOutputWithContext(context.Context) DhcpAssociateAddressMapOutput
}

type DhcpAssociateAddressMap map[string]DhcpAssociateAddressInput

func (DhcpAssociateAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpAssociateAddress)(nil)).Elem()
}

func (i DhcpAssociateAddressMap) ToDhcpAssociateAddressMapOutput() DhcpAssociateAddressMapOutput {
	return i.ToDhcpAssociateAddressMapOutputWithContext(context.Background())
}

func (i DhcpAssociateAddressMap) ToDhcpAssociateAddressMapOutputWithContext(ctx context.Context) DhcpAssociateAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpAssociateAddressMapOutput)
}

type DhcpAssociateAddressOutput struct{ *pulumi.OutputState }

func (DhcpAssociateAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpAssociateAddress)(nil)).Elem()
}

func (o DhcpAssociateAddressOutput) ToDhcpAssociateAddressOutput() DhcpAssociateAddressOutput {
	return o
}

func (o DhcpAssociateAddressOutput) ToDhcpAssociateAddressOutputWithContext(ctx context.Context) DhcpAssociateAddressOutput {
	return o
}

// Elastic public network `IP`. Must be `EIP` not bound to `DhcpIp`.
func (o DhcpAssociateAddressOutput) AddressIp() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpAssociateAddress) pulumi.StringOutput { return v.AddressIp }).(pulumi.StringOutput)
}

// `DhcpIp` unique `ID`, like: `dhcpip-9o233uri`. Must be a `DhcpIp` that is not bound to `EIP`.
func (o DhcpAssociateAddressOutput) DhcpIpId() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpAssociateAddress) pulumi.StringOutput { return v.DhcpIpId }).(pulumi.StringOutput)
}

type DhcpAssociateAddressArrayOutput struct{ *pulumi.OutputState }

func (DhcpAssociateAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpAssociateAddress)(nil)).Elem()
}

func (o DhcpAssociateAddressArrayOutput) ToDhcpAssociateAddressArrayOutput() DhcpAssociateAddressArrayOutput {
	return o
}

func (o DhcpAssociateAddressArrayOutput) ToDhcpAssociateAddressArrayOutputWithContext(ctx context.Context) DhcpAssociateAddressArrayOutput {
	return o
}

func (o DhcpAssociateAddressArrayOutput) Index(i pulumi.IntInput) DhcpAssociateAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DhcpAssociateAddress {
		return vs[0].([]*DhcpAssociateAddress)[vs[1].(int)]
	}).(DhcpAssociateAddressOutput)
}

type DhcpAssociateAddressMapOutput struct{ *pulumi.OutputState }

func (DhcpAssociateAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpAssociateAddress)(nil)).Elem()
}

func (o DhcpAssociateAddressMapOutput) ToDhcpAssociateAddressMapOutput() DhcpAssociateAddressMapOutput {
	return o
}

func (o DhcpAssociateAddressMapOutput) ToDhcpAssociateAddressMapOutputWithContext(ctx context.Context) DhcpAssociateAddressMapOutput {
	return o
}

func (o DhcpAssociateAddressMapOutput) MapIndex(k pulumi.StringInput) DhcpAssociateAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DhcpAssociateAddress {
		return vs[0].(map[string]*DhcpAssociateAddress)[vs[1].(string)]
	}).(DhcpAssociateAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpAssociateAddressInput)(nil)).Elem(), &DhcpAssociateAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpAssociateAddressArrayInput)(nil)).Elem(), DhcpAssociateAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpAssociateAddressMapInput)(nil)).Elem(), DhcpAssociateAddressMap{})
	pulumi.RegisterOutputType(DhcpAssociateAddressOutput{})
	pulumi.RegisterOutputType(DhcpAssociateAddressArrayOutput{})
	pulumi.RegisterOutputType(DhcpAssociateAddressMapOutput{})
}
