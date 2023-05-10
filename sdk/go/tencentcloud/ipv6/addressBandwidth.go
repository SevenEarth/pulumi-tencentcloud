// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ipv6

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a ipv6AddressBandwidth
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ipv6"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Ipv6.NewAddressBandwidth(ctx, "ipv6AddressBandwidth", &Ipv6.AddressBandwidthArgs{
// 			InternetChargeType:      pulumi.String("TRAFFIC_POSTPAID_BY_HOUR"),
// 			InternetMaxBandwidthOut: pulumi.Int(6),
// 			Ipv6Address:             pulumi.String("2402:4e00:1019:9400:0:9905:a90b:2ef0"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AddressBandwidth struct {
	pulumi.CustomResourceState

	// The bandwidth package id, the Legacy account and the ipv6 address to apply for the bandwidth package charge type need to be passed in.
	BandwidthPackageId pulumi.StringPtrOutput `pulumi:"bandwidthPackageId"`
	// Network billing mode. IPV6 currently supports: `TRAFFIC_POSTPAID_BY_HOUR`, for standard account types; `BANDWIDTH_PACKAGE`, for traditional account types. The default network billing mode is: `TRAFFIC_POSTPAID_BY_HOUR`.
	InternetChargeType pulumi.StringPtrOutput `pulumi:"internetChargeType"`
	// Bandwidth, in Mbps. The default is 1Mbps.
	InternetMaxBandwidthOut pulumi.IntPtrOutput `pulumi:"internetMaxBandwidthOut"`
	// IPV6 address that needs to be enabled for public network access.
	Ipv6Address pulumi.StringOutput `pulumi:"ipv6Address"`
}

// NewAddressBandwidth registers a new resource with the given unique name, arguments, and options.
func NewAddressBandwidth(ctx *pulumi.Context,
	name string, args *AddressBandwidthArgs, opts ...pulumi.ResourceOption) (*AddressBandwidth, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ipv6Address == nil {
		return nil, errors.New("invalid value for required argument 'Ipv6Address'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AddressBandwidth
	err := ctx.RegisterResource("tencentcloud:Ipv6/addressBandwidth:AddressBandwidth", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddressBandwidth gets an existing AddressBandwidth resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddressBandwidth(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddressBandwidthState, opts ...pulumi.ResourceOption) (*AddressBandwidth, error) {
	var resource AddressBandwidth
	err := ctx.ReadResource("tencentcloud:Ipv6/addressBandwidth:AddressBandwidth", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AddressBandwidth resources.
type addressBandwidthState struct {
	// The bandwidth package id, the Legacy account and the ipv6 address to apply for the bandwidth package charge type need to be passed in.
	BandwidthPackageId *string `pulumi:"bandwidthPackageId"`
	// Network billing mode. IPV6 currently supports: `TRAFFIC_POSTPAID_BY_HOUR`, for standard account types; `BANDWIDTH_PACKAGE`, for traditional account types. The default network billing mode is: `TRAFFIC_POSTPAID_BY_HOUR`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// Bandwidth, in Mbps. The default is 1Mbps.
	InternetMaxBandwidthOut *int `pulumi:"internetMaxBandwidthOut"`
	// IPV6 address that needs to be enabled for public network access.
	Ipv6Address *string `pulumi:"ipv6Address"`
}

type AddressBandwidthState struct {
	// The bandwidth package id, the Legacy account and the ipv6 address to apply for the bandwidth package charge type need to be passed in.
	BandwidthPackageId pulumi.StringPtrInput
	// Network billing mode. IPV6 currently supports: `TRAFFIC_POSTPAID_BY_HOUR`, for standard account types; `BANDWIDTH_PACKAGE`, for traditional account types. The default network billing mode is: `TRAFFIC_POSTPAID_BY_HOUR`.
	InternetChargeType pulumi.StringPtrInput
	// Bandwidth, in Mbps. The default is 1Mbps.
	InternetMaxBandwidthOut pulumi.IntPtrInput
	// IPV6 address that needs to be enabled for public network access.
	Ipv6Address pulumi.StringPtrInput
}

func (AddressBandwidthState) ElementType() reflect.Type {
	return reflect.TypeOf((*addressBandwidthState)(nil)).Elem()
}

type addressBandwidthArgs struct {
	// The bandwidth package id, the Legacy account and the ipv6 address to apply for the bandwidth package charge type need to be passed in.
	BandwidthPackageId *string `pulumi:"bandwidthPackageId"`
	// Network billing mode. IPV6 currently supports: `TRAFFIC_POSTPAID_BY_HOUR`, for standard account types; `BANDWIDTH_PACKAGE`, for traditional account types. The default network billing mode is: `TRAFFIC_POSTPAID_BY_HOUR`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// Bandwidth, in Mbps. The default is 1Mbps.
	InternetMaxBandwidthOut *int `pulumi:"internetMaxBandwidthOut"`
	// IPV6 address that needs to be enabled for public network access.
	Ipv6Address string `pulumi:"ipv6Address"`
}

// The set of arguments for constructing a AddressBandwidth resource.
type AddressBandwidthArgs struct {
	// The bandwidth package id, the Legacy account and the ipv6 address to apply for the bandwidth package charge type need to be passed in.
	BandwidthPackageId pulumi.StringPtrInput
	// Network billing mode. IPV6 currently supports: `TRAFFIC_POSTPAID_BY_HOUR`, for standard account types; `BANDWIDTH_PACKAGE`, for traditional account types. The default network billing mode is: `TRAFFIC_POSTPAID_BY_HOUR`.
	InternetChargeType pulumi.StringPtrInput
	// Bandwidth, in Mbps. The default is 1Mbps.
	InternetMaxBandwidthOut pulumi.IntPtrInput
	// IPV6 address that needs to be enabled for public network access.
	Ipv6Address pulumi.StringInput
}

func (AddressBandwidthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addressBandwidthArgs)(nil)).Elem()
}

type AddressBandwidthInput interface {
	pulumi.Input

	ToAddressBandwidthOutput() AddressBandwidthOutput
	ToAddressBandwidthOutputWithContext(ctx context.Context) AddressBandwidthOutput
}

func (*AddressBandwidth) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressBandwidth)(nil)).Elem()
}

func (i *AddressBandwidth) ToAddressBandwidthOutput() AddressBandwidthOutput {
	return i.ToAddressBandwidthOutputWithContext(context.Background())
}

func (i *AddressBandwidth) ToAddressBandwidthOutputWithContext(ctx context.Context) AddressBandwidthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressBandwidthOutput)
}

// AddressBandwidthArrayInput is an input type that accepts AddressBandwidthArray and AddressBandwidthArrayOutput values.
// You can construct a concrete instance of `AddressBandwidthArrayInput` via:
//
//          AddressBandwidthArray{ AddressBandwidthArgs{...} }
type AddressBandwidthArrayInput interface {
	pulumi.Input

	ToAddressBandwidthArrayOutput() AddressBandwidthArrayOutput
	ToAddressBandwidthArrayOutputWithContext(context.Context) AddressBandwidthArrayOutput
}

type AddressBandwidthArray []AddressBandwidthInput

func (AddressBandwidthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AddressBandwidth)(nil)).Elem()
}

func (i AddressBandwidthArray) ToAddressBandwidthArrayOutput() AddressBandwidthArrayOutput {
	return i.ToAddressBandwidthArrayOutputWithContext(context.Background())
}

func (i AddressBandwidthArray) ToAddressBandwidthArrayOutputWithContext(ctx context.Context) AddressBandwidthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressBandwidthArrayOutput)
}

// AddressBandwidthMapInput is an input type that accepts AddressBandwidthMap and AddressBandwidthMapOutput values.
// You can construct a concrete instance of `AddressBandwidthMapInput` via:
//
//          AddressBandwidthMap{ "key": AddressBandwidthArgs{...} }
type AddressBandwidthMapInput interface {
	pulumi.Input

	ToAddressBandwidthMapOutput() AddressBandwidthMapOutput
	ToAddressBandwidthMapOutputWithContext(context.Context) AddressBandwidthMapOutput
}

type AddressBandwidthMap map[string]AddressBandwidthInput

func (AddressBandwidthMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AddressBandwidth)(nil)).Elem()
}

func (i AddressBandwidthMap) ToAddressBandwidthMapOutput() AddressBandwidthMapOutput {
	return i.ToAddressBandwidthMapOutputWithContext(context.Background())
}

func (i AddressBandwidthMap) ToAddressBandwidthMapOutputWithContext(ctx context.Context) AddressBandwidthMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressBandwidthMapOutput)
}

type AddressBandwidthOutput struct{ *pulumi.OutputState }

func (AddressBandwidthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressBandwidth)(nil)).Elem()
}

func (o AddressBandwidthOutput) ToAddressBandwidthOutput() AddressBandwidthOutput {
	return o
}

func (o AddressBandwidthOutput) ToAddressBandwidthOutputWithContext(ctx context.Context) AddressBandwidthOutput {
	return o
}

// The bandwidth package id, the Legacy account and the ipv6 address to apply for the bandwidth package charge type need to be passed in.
func (o AddressBandwidthOutput) BandwidthPackageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddressBandwidth) pulumi.StringPtrOutput { return v.BandwidthPackageId }).(pulumi.StringPtrOutput)
}

// Network billing mode. IPV6 currently supports: `TRAFFIC_POSTPAID_BY_HOUR`, for standard account types; `BANDWIDTH_PACKAGE`, for traditional account types. The default network billing mode is: `TRAFFIC_POSTPAID_BY_HOUR`.
func (o AddressBandwidthOutput) InternetChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddressBandwidth) pulumi.StringPtrOutput { return v.InternetChargeType }).(pulumi.StringPtrOutput)
}

// Bandwidth, in Mbps. The default is 1Mbps.
func (o AddressBandwidthOutput) InternetMaxBandwidthOut() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AddressBandwidth) pulumi.IntPtrOutput { return v.InternetMaxBandwidthOut }).(pulumi.IntPtrOutput)
}

// IPV6 address that needs to be enabled for public network access.
func (o AddressBandwidthOutput) Ipv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressBandwidth) pulumi.StringOutput { return v.Ipv6Address }).(pulumi.StringOutput)
}

type AddressBandwidthArrayOutput struct{ *pulumi.OutputState }

func (AddressBandwidthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AddressBandwidth)(nil)).Elem()
}

func (o AddressBandwidthArrayOutput) ToAddressBandwidthArrayOutput() AddressBandwidthArrayOutput {
	return o
}

func (o AddressBandwidthArrayOutput) ToAddressBandwidthArrayOutputWithContext(ctx context.Context) AddressBandwidthArrayOutput {
	return o
}

func (o AddressBandwidthArrayOutput) Index(i pulumi.IntInput) AddressBandwidthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AddressBandwidth {
		return vs[0].([]*AddressBandwidth)[vs[1].(int)]
	}).(AddressBandwidthOutput)
}

type AddressBandwidthMapOutput struct{ *pulumi.OutputState }

func (AddressBandwidthMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AddressBandwidth)(nil)).Elem()
}

func (o AddressBandwidthMapOutput) ToAddressBandwidthMapOutput() AddressBandwidthMapOutput {
	return o
}

func (o AddressBandwidthMapOutput) ToAddressBandwidthMapOutputWithContext(ctx context.Context) AddressBandwidthMapOutput {
	return o
}

func (o AddressBandwidthMapOutput) MapIndex(k pulumi.StringInput) AddressBandwidthOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AddressBandwidth {
		return vs[0].(map[string]*AddressBandwidth)[vs[1].(string)]
	}).(AddressBandwidthOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddressBandwidthInput)(nil)).Elem(), &AddressBandwidth{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressBandwidthArrayInput)(nil)).Elem(), AddressBandwidthArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressBandwidthMapInput)(nil)).Elem(), AddressBandwidthMap{})
	pulumi.RegisterOutputType(AddressBandwidthOutput{})
	pulumi.RegisterOutputType(AddressBandwidthArrayOutput{})
	pulumi.RegisterOutputType(AddressBandwidthMapOutput{})
}