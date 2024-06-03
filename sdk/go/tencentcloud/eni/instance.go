// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eni

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create an ENI.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Eni"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Security"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			zones, err := Availability.GetZonesByProduct(ctx, &availability.GetZonesByProductArgs{
//				Product: "vpc",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vpc, err := Vpc.NewInstance(ctx, "vpc", &Vpc.InstanceArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			subnet, err := Subnet.NewInstance(ctx, "subnet", &Subnet.InstanceArgs{
//				AvailabilityZone: pulumi.String(zones.Zones[0].Name),
//				VpcId:            vpc.ID(),
//				CidrBlock:        pulumi.String("10.0.0.0/16"),
//				IsMulticast:      pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			example1, err := Security.NewGroup(ctx, "example1", &Security.GroupArgs{
//				Description: pulumi.String("sg desc."),
//				ProjectId:   pulumi.Int(0),
//				Tags: pulumi.Map{
//					"example": pulumi.Any("test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			example2, err := Security.NewGroup(ctx, "example2", &Security.GroupArgs{
//				Description: pulumi.String("sg desc."),
//				ProjectId:   pulumi.Int(0),
//				Tags: pulumi.Map{
//					"example": pulumi.Any("test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Eni.NewInstance(ctx, "example", &Eni.InstanceArgs{
//				VpcId:       vpc.ID(),
//				SubnetId:    subnet.ID(),
//				Description: pulumi.String("eni desc."),
//				Ipv4Count:   pulumi.Int(1),
//				SecurityGroups: pulumi.StringArray{
//					example1.ID(),
//					example2.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// ENI can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Eni/instance:Instance tencentcloud_eni.foo eni-qka182br
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Creation time of the ENI.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the IP, maximum length 25.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The number of intranet IPv4s. When it is greater than 1, there is only one primary intranet IP. The others are auxiliary intranet IPs, which conflict with `ipv4s`.
	Ipv4Count pulumi.IntPtrOutput `pulumi:"ipv4Count"`
	// An information list of IPv4s. Each element contains the following attributes:
	Ipv4Infos InstanceIpv4InfoArrayOutput `pulumi:"ipv4Infos"`
	// Applying for intranet IPv4s collection, conflict with `ipv4Count`. When there are multiple ipv4s, can only be one primary IP, and the maximum length of the array is 30. Each element contains the following attributes:
	Ipv4s InstanceIpv4ArrayOutput `pulumi:"ipv4s"`
	// MAC address.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// Name of the ENI, maximum length 60.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates whether the IP is primary.
	Primary pulumi.BoolOutput `pulumi:"primary"`
	// A set of security group IDs.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// State of the ENI.
	State pulumi.StringOutput `pulumi:"state"`
	// ID of the subnet within this vpc.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Tags of the ENI.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// ID of the vpc.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("tencentcloud:Eni/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("tencentcloud:Eni/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Creation time of the ENI.
	CreateTime *string `pulumi:"createTime"`
	// Description of the IP, maximum length 25.
	Description *string `pulumi:"description"`
	// The number of intranet IPv4s. When it is greater than 1, there is only one primary intranet IP. The others are auxiliary intranet IPs, which conflict with `ipv4s`.
	Ipv4Count *int `pulumi:"ipv4Count"`
	// An information list of IPv4s. Each element contains the following attributes:
	Ipv4Infos []InstanceIpv4Info `pulumi:"ipv4Infos"`
	// Applying for intranet IPv4s collection, conflict with `ipv4Count`. When there are multiple ipv4s, can only be one primary IP, and the maximum length of the array is 30. Each element contains the following attributes:
	Ipv4s []InstanceIpv4 `pulumi:"ipv4s"`
	// MAC address.
	Mac *string `pulumi:"mac"`
	// Name of the ENI, maximum length 60.
	Name *string `pulumi:"name"`
	// Indicates whether the IP is primary.
	Primary *bool `pulumi:"primary"`
	// A set of security group IDs.
	SecurityGroups []string `pulumi:"securityGroups"`
	// State of the ENI.
	State *string `pulumi:"state"`
	// ID of the subnet within this vpc.
	SubnetId *string `pulumi:"subnetId"`
	// Tags of the ENI.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the vpc.
	VpcId *string `pulumi:"vpcId"`
}

type InstanceState struct {
	// Creation time of the ENI.
	CreateTime pulumi.StringPtrInput
	// Description of the IP, maximum length 25.
	Description pulumi.StringPtrInput
	// The number of intranet IPv4s. When it is greater than 1, there is only one primary intranet IP. The others are auxiliary intranet IPs, which conflict with `ipv4s`.
	Ipv4Count pulumi.IntPtrInput
	// An information list of IPv4s. Each element contains the following attributes:
	Ipv4Infos InstanceIpv4InfoArrayInput
	// Applying for intranet IPv4s collection, conflict with `ipv4Count`. When there are multiple ipv4s, can only be one primary IP, and the maximum length of the array is 30. Each element contains the following attributes:
	Ipv4s InstanceIpv4ArrayInput
	// MAC address.
	Mac pulumi.StringPtrInput
	// Name of the ENI, maximum length 60.
	Name pulumi.StringPtrInput
	// Indicates whether the IP is primary.
	Primary pulumi.BoolPtrInput
	// A set of security group IDs.
	SecurityGroups pulumi.StringArrayInput
	// State of the ENI.
	State pulumi.StringPtrInput
	// ID of the subnet within this vpc.
	SubnetId pulumi.StringPtrInput
	// Tags of the ENI.
	Tags pulumi.MapInput
	// ID of the vpc.
	VpcId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Description of the IP, maximum length 25.
	Description *string `pulumi:"description"`
	// The number of intranet IPv4s. When it is greater than 1, there is only one primary intranet IP. The others are auxiliary intranet IPs, which conflict with `ipv4s`.
	Ipv4Count *int `pulumi:"ipv4Count"`
	// Applying for intranet IPv4s collection, conflict with `ipv4Count`. When there are multiple ipv4s, can only be one primary IP, and the maximum length of the array is 30. Each element contains the following attributes:
	Ipv4s []InstanceIpv4 `pulumi:"ipv4s"`
	// Name of the ENI, maximum length 60.
	Name *string `pulumi:"name"`
	// A set of security group IDs.
	SecurityGroups []string `pulumi:"securityGroups"`
	// ID of the subnet within this vpc.
	SubnetId string `pulumi:"subnetId"`
	// Tags of the ENI.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the vpc.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Description of the IP, maximum length 25.
	Description pulumi.StringPtrInput
	// The number of intranet IPv4s. When it is greater than 1, there is only one primary intranet IP. The others are auxiliary intranet IPs, which conflict with `ipv4s`.
	Ipv4Count pulumi.IntPtrInput
	// Applying for intranet IPv4s collection, conflict with `ipv4Count`. When there are multiple ipv4s, can only be one primary IP, and the maximum length of the array is 30. Each element contains the following attributes:
	Ipv4s InstanceIpv4ArrayInput
	// Name of the ENI, maximum length 60.
	Name pulumi.StringPtrInput
	// A set of security group IDs.
	SecurityGroups pulumi.StringArrayInput
	// ID of the subnet within this vpc.
	SubnetId pulumi.StringInput
	// Tags of the ENI.
	Tags pulumi.MapInput
	// ID of the vpc.
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

// Creation time of the ENI.
func (o InstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the IP, maximum length 25.
func (o InstanceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The number of intranet IPv4s. When it is greater than 1, there is only one primary intranet IP. The others are auxiliary intranet IPs, which conflict with `ipv4s`.
func (o InstanceOutput) Ipv4Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.Ipv4Count }).(pulumi.IntPtrOutput)
}

// An information list of IPv4s. Each element contains the following attributes:
func (o InstanceOutput) Ipv4Infos() InstanceIpv4InfoArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceIpv4InfoArrayOutput { return v.Ipv4Infos }).(InstanceIpv4InfoArrayOutput)
}

// Applying for intranet IPv4s collection, conflict with `ipv4Count`. When there are multiple ipv4s, can only be one primary IP, and the maximum length of the array is 30. Each element contains the following attributes:
func (o InstanceOutput) Ipv4s() InstanceIpv4ArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceIpv4ArrayOutput { return v.Ipv4s }).(InstanceIpv4ArrayOutput)
}

// MAC address.
func (o InstanceOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Mac }).(pulumi.StringOutput)
}

// Name of the ENI, maximum length 60.
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Indicates whether the IP is primary.
func (o InstanceOutput) Primary() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.Primary }).(pulumi.BoolOutput)
}

// A set of security group IDs.
func (o InstanceOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// State of the ENI.
func (o InstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// ID of the subnet within this vpc.
func (o InstanceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Tags of the ENI.
func (o InstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Instance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// ID of the vpc.
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
