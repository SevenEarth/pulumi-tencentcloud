// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package subnet

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a resource to create a VPC subnet.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			availabilityZone := "ap-guangzhou-3"
//			if param := cfg.Get("availabilityZone"); param != "" {
//				availabilityZone = param
//			}
//			foo, err := Vpc.NewInstance(ctx, "foo", &Vpc.InstanceArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Subnet.NewInstance(ctx, "subnet", &Subnet.InstanceArgs{
//				AvailabilityZone: pulumi.String(availabilityZone),
//				VpcId:            foo.ID(),
//				CidrBlock:        pulumi.String("10.0.20.0/28"),
//				IsMulticast:      pulumi.Bool(false),
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
// Vpc subnet instance can be imported, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Subnet/instance:Instance test subnet_id
//
// ```
type Instance struct {
	pulumi.CustomResourceState

	// The availability zone within which the subnet should be created.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The number of available IPs.
	AvailableIpCount pulumi.IntOutput `pulumi:"availableIpCount"`
	// A network address block of the subnet.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// Creation time of subnet resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Indicates whether it is the default VPC for this region.
	IsDefault pulumi.BoolOutput `pulumi:"isDefault"`
	// Indicates whether multicast is enabled. The default value is 'true'.
	IsMulticast pulumi.BoolPtrOutput `pulumi:"isMulticast"`
	// The name of subnet to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of a routing table to which the subnet should be associated.
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	// Tags of the subnet.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// ID of the VPC to be associated.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	if args.CidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'CidrBlock'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("tencentcloud:Subnet/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("tencentcloud:Subnet/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The availability zone within which the subnet should be created.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The number of available IPs.
	AvailableIpCount *int `pulumi:"availableIpCount"`
	// A network address block of the subnet.
	CidrBlock *string `pulumi:"cidrBlock"`
	// Creation time of subnet resource.
	CreateTime *string `pulumi:"createTime"`
	// Indicates whether it is the default VPC for this region.
	IsDefault *bool `pulumi:"isDefault"`
	// Indicates whether multicast is enabled. The default value is 'true'.
	IsMulticast *bool `pulumi:"isMulticast"`
	// The name of subnet to be created.
	Name *string `pulumi:"name"`
	// ID of a routing table to which the subnet should be associated.
	RouteTableId *string `pulumi:"routeTableId"`
	// Tags of the subnet.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the VPC to be associated.
	VpcId *string `pulumi:"vpcId"`
}

type InstanceState struct {
	// The availability zone within which the subnet should be created.
	AvailabilityZone pulumi.StringPtrInput
	// The number of available IPs.
	AvailableIpCount pulumi.IntPtrInput
	// A network address block of the subnet.
	CidrBlock pulumi.StringPtrInput
	// Creation time of subnet resource.
	CreateTime pulumi.StringPtrInput
	// Indicates whether it is the default VPC for this region.
	IsDefault pulumi.BoolPtrInput
	// Indicates whether multicast is enabled. The default value is 'true'.
	IsMulticast pulumi.BoolPtrInput
	// The name of subnet to be created.
	Name pulumi.StringPtrInput
	// ID of a routing table to which the subnet should be associated.
	RouteTableId pulumi.StringPtrInput
	// Tags of the subnet.
	Tags pulumi.MapInput
	// ID of the VPC to be associated.
	VpcId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The availability zone within which the subnet should be created.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// A network address block of the subnet.
	CidrBlock string `pulumi:"cidrBlock"`
	// Indicates whether multicast is enabled. The default value is 'true'.
	IsMulticast *bool `pulumi:"isMulticast"`
	// The name of subnet to be created.
	Name *string `pulumi:"name"`
	// ID of a routing table to which the subnet should be associated.
	RouteTableId *string `pulumi:"routeTableId"`
	// Tags of the subnet.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the VPC to be associated.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The availability zone within which the subnet should be created.
	AvailabilityZone pulumi.StringInput
	// A network address block of the subnet.
	CidrBlock pulumi.StringInput
	// Indicates whether multicast is enabled. The default value is 'true'.
	IsMulticast pulumi.BoolPtrInput
	// The name of subnet to be created.
	Name pulumi.StringPtrInput
	// ID of a routing table to which the subnet should be associated.
	RouteTableId pulumi.StringPtrInput
	// Tags of the subnet.
	Tags pulumi.MapInput
	// ID of the VPC to be associated.
	VpcId pulumi.StringInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// The availability zone within which the subnet should be created.
func (o InstanceOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// The number of available IPs.
func (o InstanceOutput) AvailableIpCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.AvailableIpCount }).(pulumi.IntOutput)
}

// A network address block of the subnet.
func (o InstanceOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CidrBlock }).(pulumi.StringOutput)
}

// Creation time of subnet resource.
func (o InstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Indicates whether it is the default VPC for this region.
func (o InstanceOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.IsDefault }).(pulumi.BoolOutput)
}

// Indicates whether multicast is enabled. The default value is 'true'.
func (o InstanceOutput) IsMulticast() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.IsMulticast }).(pulumi.BoolPtrOutput)
}

// The name of subnet to be created.
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of a routing table to which the subnet should be associated.
func (o InstanceOutput) RouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.RouteTableId }).(pulumi.StringOutput)
}

// Tags of the subnet.
func (o InstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Instance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// ID of the VPC to be associated.
func (o InstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
