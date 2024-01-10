// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a vpc ipv6CidrBlock
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			vpc, err := Vpc.NewInstance(ctx, "vpc", &Vpc.InstanceArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Vpc.NewIpv6CidrBlock(ctx, "example", &Vpc.Ipv6CidrBlockArgs{
//				VpcId: vpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// vpc ipv6_cidr_block can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Vpc/ipv6CidrBlock:Ipv6CidrBlock ipv6_cidr_block vpc_id
//
// ```
type Ipv6CidrBlock struct {
	pulumi.CustomResourceState

	// ipv6 cidr block.
	Ipv6CidrBlock pulumi.StringOutput `pulumi:"ipv6CidrBlock"`
	// `VPC` instance `ID`, in the form of `vpc-f49l6u0z`.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewIpv6CidrBlock registers a new resource with the given unique name, arguments, and options.
func NewIpv6CidrBlock(ctx *pulumi.Context,
	name string, args *Ipv6CidrBlockArgs, opts ...pulumi.ResourceOption) (*Ipv6CidrBlock, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Ipv6CidrBlock
	err := ctx.RegisterResource("tencentcloud:Vpc/ipv6CidrBlock:Ipv6CidrBlock", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpv6CidrBlock gets an existing Ipv6CidrBlock resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpv6CidrBlock(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ipv6CidrBlockState, opts ...pulumi.ResourceOption) (*Ipv6CidrBlock, error) {
	var resource Ipv6CidrBlock
	err := ctx.ReadResource("tencentcloud:Vpc/ipv6CidrBlock:Ipv6CidrBlock", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ipv6CidrBlock resources.
type ipv6CidrBlockState struct {
	// ipv6 cidr block.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// `VPC` instance `ID`, in the form of `vpc-f49l6u0z`.
	VpcId *string `pulumi:"vpcId"`
}

type Ipv6CidrBlockState struct {
	// ipv6 cidr block.
	Ipv6CidrBlock pulumi.StringPtrInput
	// `VPC` instance `ID`, in the form of `vpc-f49l6u0z`.
	VpcId pulumi.StringPtrInput
}

func (Ipv6CidrBlockState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv6CidrBlockState)(nil)).Elem()
}

type ipv6CidrBlockArgs struct {
	// `VPC` instance `ID`, in the form of `vpc-f49l6u0z`.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Ipv6CidrBlock resource.
type Ipv6CidrBlockArgs struct {
	// `VPC` instance `ID`, in the form of `vpc-f49l6u0z`.
	VpcId pulumi.StringInput
}

func (Ipv6CidrBlockArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv6CidrBlockArgs)(nil)).Elem()
}

type Ipv6CidrBlockInput interface {
	pulumi.Input

	ToIpv6CidrBlockOutput() Ipv6CidrBlockOutput
	ToIpv6CidrBlockOutputWithContext(ctx context.Context) Ipv6CidrBlockOutput
}

func (*Ipv6CidrBlock) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6CidrBlock)(nil)).Elem()
}

func (i *Ipv6CidrBlock) ToIpv6CidrBlockOutput() Ipv6CidrBlockOutput {
	return i.ToIpv6CidrBlockOutputWithContext(context.Background())
}

func (i *Ipv6CidrBlock) ToIpv6CidrBlockOutputWithContext(ctx context.Context) Ipv6CidrBlockOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6CidrBlockOutput)
}

// Ipv6CidrBlockArrayInput is an input type that accepts Ipv6CidrBlockArray and Ipv6CidrBlockArrayOutput values.
// You can construct a concrete instance of `Ipv6CidrBlockArrayInput` via:
//
//	Ipv6CidrBlockArray{ Ipv6CidrBlockArgs{...} }
type Ipv6CidrBlockArrayInput interface {
	pulumi.Input

	ToIpv6CidrBlockArrayOutput() Ipv6CidrBlockArrayOutput
	ToIpv6CidrBlockArrayOutputWithContext(context.Context) Ipv6CidrBlockArrayOutput
}

type Ipv6CidrBlockArray []Ipv6CidrBlockInput

func (Ipv6CidrBlockArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv6CidrBlock)(nil)).Elem()
}

func (i Ipv6CidrBlockArray) ToIpv6CidrBlockArrayOutput() Ipv6CidrBlockArrayOutput {
	return i.ToIpv6CidrBlockArrayOutputWithContext(context.Background())
}

func (i Ipv6CidrBlockArray) ToIpv6CidrBlockArrayOutputWithContext(ctx context.Context) Ipv6CidrBlockArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6CidrBlockArrayOutput)
}

// Ipv6CidrBlockMapInput is an input type that accepts Ipv6CidrBlockMap and Ipv6CidrBlockMapOutput values.
// You can construct a concrete instance of `Ipv6CidrBlockMapInput` via:
//
//	Ipv6CidrBlockMap{ "key": Ipv6CidrBlockArgs{...} }
type Ipv6CidrBlockMapInput interface {
	pulumi.Input

	ToIpv6CidrBlockMapOutput() Ipv6CidrBlockMapOutput
	ToIpv6CidrBlockMapOutputWithContext(context.Context) Ipv6CidrBlockMapOutput
}

type Ipv6CidrBlockMap map[string]Ipv6CidrBlockInput

func (Ipv6CidrBlockMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv6CidrBlock)(nil)).Elem()
}

func (i Ipv6CidrBlockMap) ToIpv6CidrBlockMapOutput() Ipv6CidrBlockMapOutput {
	return i.ToIpv6CidrBlockMapOutputWithContext(context.Background())
}

func (i Ipv6CidrBlockMap) ToIpv6CidrBlockMapOutputWithContext(ctx context.Context) Ipv6CidrBlockMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6CidrBlockMapOutput)
}

type Ipv6CidrBlockOutput struct{ *pulumi.OutputState }

func (Ipv6CidrBlockOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6CidrBlock)(nil)).Elem()
}

func (o Ipv6CidrBlockOutput) ToIpv6CidrBlockOutput() Ipv6CidrBlockOutput {
	return o
}

func (o Ipv6CidrBlockOutput) ToIpv6CidrBlockOutputWithContext(ctx context.Context) Ipv6CidrBlockOutput {
	return o
}

// ipv6 cidr block.
func (o Ipv6CidrBlockOutput) Ipv6CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6CidrBlock) pulumi.StringOutput { return v.Ipv6CidrBlock }).(pulumi.StringOutput)
}

// `VPC` instance `ID`, in the form of `vpc-f49l6u0z`.
func (o Ipv6CidrBlockOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6CidrBlock) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type Ipv6CidrBlockArrayOutput struct{ *pulumi.OutputState }

func (Ipv6CidrBlockArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv6CidrBlock)(nil)).Elem()
}

func (o Ipv6CidrBlockArrayOutput) ToIpv6CidrBlockArrayOutput() Ipv6CidrBlockArrayOutput {
	return o
}

func (o Ipv6CidrBlockArrayOutput) ToIpv6CidrBlockArrayOutputWithContext(ctx context.Context) Ipv6CidrBlockArrayOutput {
	return o
}

func (o Ipv6CidrBlockArrayOutput) Index(i pulumi.IntInput) Ipv6CidrBlockOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ipv6CidrBlock {
		return vs[0].([]*Ipv6CidrBlock)[vs[1].(int)]
	}).(Ipv6CidrBlockOutput)
}

type Ipv6CidrBlockMapOutput struct{ *pulumi.OutputState }

func (Ipv6CidrBlockMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv6CidrBlock)(nil)).Elem()
}

func (o Ipv6CidrBlockMapOutput) ToIpv6CidrBlockMapOutput() Ipv6CidrBlockMapOutput {
	return o
}

func (o Ipv6CidrBlockMapOutput) ToIpv6CidrBlockMapOutputWithContext(ctx context.Context) Ipv6CidrBlockMapOutput {
	return o
}

func (o Ipv6CidrBlockMapOutput) MapIndex(k pulumi.StringInput) Ipv6CidrBlockOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ipv6CidrBlock {
		return vs[0].(map[string]*Ipv6CidrBlock)[vs[1].(string)]
	}).(Ipv6CidrBlockOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6CidrBlockInput)(nil)).Elem(), &Ipv6CidrBlock{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6CidrBlockArrayInput)(nil)).Elem(), Ipv6CidrBlockArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6CidrBlockMapInput)(nil)).Elem(), Ipv6CidrBlockMap{})
	pulumi.RegisterOutputType(Ipv6CidrBlockOutput{})
	pulumi.RegisterOutputType(Ipv6CidrBlockArrayOutput{})
	pulumi.RegisterOutputType(Ipv6CidrBlockMapOutput{})
}
