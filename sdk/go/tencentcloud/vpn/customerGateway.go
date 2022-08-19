// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a VPN customer gateway.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpn"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Vpn.NewCustomerGateway(ctx, "foo", &Vpn.CustomerGatewayArgs{
// 			PublicIpAddress: pulumi.String("1.1.1.1"),
// 			Tags: pulumi.AnyMap{
// 				"tag": pulumi.Any("test"),
// 			},
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
// VPN customer gateway can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Vpn/customerGateway:CustomerGateway foo cgw-xfqag
// ```
type CustomerGateway struct {
	pulumi.CustomResourceState

	// Create time of the customer gateway.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Name of the customer gateway. The length of character is limited to 1-60.
	Name pulumi.StringOutput `pulumi:"name"`
	// Public IP of the customer gateway.
	PublicIpAddress pulumi.StringOutput `pulumi:"publicIpAddress"`
	// A list of tags used to associate different resources.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewCustomerGateway registers a new resource with the given unique name, arguments, and options.
func NewCustomerGateway(ctx *pulumi.Context,
	name string, args *CustomerGatewayArgs, opts ...pulumi.ResourceOption) (*CustomerGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublicIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'PublicIpAddress'")
	}
	var resource CustomerGateway
	err := ctx.RegisterResource("tencentcloud:Vpn/customerGateway:CustomerGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomerGateway gets an existing CustomerGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomerGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerGatewayState, opts ...pulumi.ResourceOption) (*CustomerGateway, error) {
	var resource CustomerGateway
	err := ctx.ReadResource("tencentcloud:Vpn/customerGateway:CustomerGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomerGateway resources.
type customerGatewayState struct {
	// Create time of the customer gateway.
	CreateTime *string `pulumi:"createTime"`
	// Name of the customer gateway. The length of character is limited to 1-60.
	Name *string `pulumi:"name"`
	// Public IP of the customer gateway.
	PublicIpAddress *string `pulumi:"publicIpAddress"`
	// A list of tags used to associate different resources.
	Tags map[string]interface{} `pulumi:"tags"`
}

type CustomerGatewayState struct {
	// Create time of the customer gateway.
	CreateTime pulumi.StringPtrInput
	// Name of the customer gateway. The length of character is limited to 1-60.
	Name pulumi.StringPtrInput
	// Public IP of the customer gateway.
	PublicIpAddress pulumi.StringPtrInput
	// A list of tags used to associate different resources.
	Tags pulumi.MapInput
}

func (CustomerGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*customerGatewayState)(nil)).Elem()
}

type customerGatewayArgs struct {
	// Name of the customer gateway. The length of character is limited to 1-60.
	Name *string `pulumi:"name"`
	// Public IP of the customer gateway.
	PublicIpAddress string `pulumi:"publicIpAddress"`
	// A list of tags used to associate different resources.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a CustomerGateway resource.
type CustomerGatewayArgs struct {
	// Name of the customer gateway. The length of character is limited to 1-60.
	Name pulumi.StringPtrInput
	// Public IP of the customer gateway.
	PublicIpAddress pulumi.StringInput
	// A list of tags used to associate different resources.
	Tags pulumi.MapInput
}

func (CustomerGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customerGatewayArgs)(nil)).Elem()
}

type CustomerGatewayInput interface {
	pulumi.Input

	ToCustomerGatewayOutput() CustomerGatewayOutput
	ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput
}

func (*CustomerGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerGateway)(nil)).Elem()
}

func (i *CustomerGateway) ToCustomerGatewayOutput() CustomerGatewayOutput {
	return i.ToCustomerGatewayOutputWithContext(context.Background())
}

func (i *CustomerGateway) ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayOutput)
}

// CustomerGatewayArrayInput is an input type that accepts CustomerGatewayArray and CustomerGatewayArrayOutput values.
// You can construct a concrete instance of `CustomerGatewayArrayInput` via:
//
//          CustomerGatewayArray{ CustomerGatewayArgs{...} }
type CustomerGatewayArrayInput interface {
	pulumi.Input

	ToCustomerGatewayArrayOutput() CustomerGatewayArrayOutput
	ToCustomerGatewayArrayOutputWithContext(context.Context) CustomerGatewayArrayOutput
}

type CustomerGatewayArray []CustomerGatewayInput

func (CustomerGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomerGateway)(nil)).Elem()
}

func (i CustomerGatewayArray) ToCustomerGatewayArrayOutput() CustomerGatewayArrayOutput {
	return i.ToCustomerGatewayArrayOutputWithContext(context.Background())
}

func (i CustomerGatewayArray) ToCustomerGatewayArrayOutputWithContext(ctx context.Context) CustomerGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayArrayOutput)
}

// CustomerGatewayMapInput is an input type that accepts CustomerGatewayMap and CustomerGatewayMapOutput values.
// You can construct a concrete instance of `CustomerGatewayMapInput` via:
//
//          CustomerGatewayMap{ "key": CustomerGatewayArgs{...} }
type CustomerGatewayMapInput interface {
	pulumi.Input

	ToCustomerGatewayMapOutput() CustomerGatewayMapOutput
	ToCustomerGatewayMapOutputWithContext(context.Context) CustomerGatewayMapOutput
}

type CustomerGatewayMap map[string]CustomerGatewayInput

func (CustomerGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomerGateway)(nil)).Elem()
}

func (i CustomerGatewayMap) ToCustomerGatewayMapOutput() CustomerGatewayMapOutput {
	return i.ToCustomerGatewayMapOutputWithContext(context.Background())
}

func (i CustomerGatewayMap) ToCustomerGatewayMapOutputWithContext(ctx context.Context) CustomerGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayMapOutput)
}

type CustomerGatewayOutput struct{ *pulumi.OutputState }

func (CustomerGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerGateway)(nil)).Elem()
}

func (o CustomerGatewayOutput) ToCustomerGatewayOutput() CustomerGatewayOutput {
	return o
}

func (o CustomerGatewayOutput) ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput {
	return o
}

// Create time of the customer gateway.
func (o CustomerGatewayOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Name of the customer gateway. The length of character is limited to 1-60.
func (o CustomerGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Public IP of the customer gateway.
func (o CustomerGatewayOutput) PublicIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.PublicIpAddress }).(pulumi.StringOutput)
}

// A list of tags used to associate different resources.
func (o CustomerGatewayOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type CustomerGatewayArrayOutput struct{ *pulumi.OutputState }

func (CustomerGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomerGateway)(nil)).Elem()
}

func (o CustomerGatewayArrayOutput) ToCustomerGatewayArrayOutput() CustomerGatewayArrayOutput {
	return o
}

func (o CustomerGatewayArrayOutput) ToCustomerGatewayArrayOutputWithContext(ctx context.Context) CustomerGatewayArrayOutput {
	return o
}

func (o CustomerGatewayArrayOutput) Index(i pulumi.IntInput) CustomerGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomerGateway {
		return vs[0].([]*CustomerGateway)[vs[1].(int)]
	}).(CustomerGatewayOutput)
}

type CustomerGatewayMapOutput struct{ *pulumi.OutputState }

func (CustomerGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomerGateway)(nil)).Elem()
}

func (o CustomerGatewayMapOutput) ToCustomerGatewayMapOutput() CustomerGatewayMapOutput {
	return o
}

func (o CustomerGatewayMapOutput) ToCustomerGatewayMapOutputWithContext(ctx context.Context) CustomerGatewayMapOutput {
	return o
}

func (o CustomerGatewayMapOutput) MapIndex(k pulumi.StringInput) CustomerGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomerGateway {
		return vs[0].(map[string]*CustomerGateway)[vs[1].(string)]
	}).(CustomerGatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerGatewayInput)(nil)).Elem(), &CustomerGateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerGatewayArrayInput)(nil)).Elem(), CustomerGatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerGatewayMapInput)(nil)).Elem(), CustomerGatewayMap{})
	pulumi.RegisterOutputType(CustomerGatewayOutput{})
	pulumi.RegisterOutputType(CustomerGatewayArrayOutput{})
	pulumi.RegisterOutputType(CustomerGatewayMapOutput{})
}
