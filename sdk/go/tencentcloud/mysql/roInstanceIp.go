// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mysql roInstanceIp
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			zones, err := Availability.GetZonesByProduct(ctx, &availability.GetZonesByProductArgs{
//				Product: "cdb",
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
//			_, err = Mysql.NewRoInstanceIp(ctx, "example", &Mysql.RoInstanceIpArgs{
//				InstanceId:   pulumi.String("cdbro-bdlvcfpj"),
//				UniqSubnetId: subnet.ID(),
//				UniqVpcId:    vpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type RoInstanceIp struct {
	pulumi.CustomResourceState

	// Read-only instance ID, in the format: cdbro-3i70uj0k, which is the same as the read-only instance ID displayed on the cloud database console page.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Intranet IP address of the read-only instance.
	RoVip pulumi.StringOutput `pulumi:"roVip"`
	// Intranet port number of the read-only instance.
	RoVport pulumi.IntOutput `pulumi:"roVport"`
	// Subnet descriptor, for example: subnet-1typ0s7d.
	UniqSubnetId pulumi.StringPtrOutput `pulumi:"uniqSubnetId"`
	// vpc descriptor, for example: vpc-a23yt67j, if this field is passed, UniqSubnetId must be passed.
	UniqVpcId pulumi.StringPtrOutput `pulumi:"uniqVpcId"`
}

// NewRoInstanceIp registers a new resource with the given unique name, arguments, and options.
func NewRoInstanceIp(ctx *pulumi.Context,
	name string, args *RoInstanceIpArgs, opts ...pulumi.ResourceOption) (*RoInstanceIp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RoInstanceIp
	err := ctx.RegisterResource("tencentcloud:Mysql/roInstanceIp:RoInstanceIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoInstanceIp gets an existing RoInstanceIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoInstanceIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoInstanceIpState, opts ...pulumi.ResourceOption) (*RoInstanceIp, error) {
	var resource RoInstanceIp
	err := ctx.ReadResource("tencentcloud:Mysql/roInstanceIp:RoInstanceIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoInstanceIp resources.
type roInstanceIpState struct {
	// Read-only instance ID, in the format: cdbro-3i70uj0k, which is the same as the read-only instance ID displayed on the cloud database console page.
	InstanceId *string `pulumi:"instanceId"`
	// Intranet IP address of the read-only instance.
	RoVip *string `pulumi:"roVip"`
	// Intranet port number of the read-only instance.
	RoVport *int `pulumi:"roVport"`
	// Subnet descriptor, for example: subnet-1typ0s7d.
	UniqSubnetId *string `pulumi:"uniqSubnetId"`
	// vpc descriptor, for example: vpc-a23yt67j, if this field is passed, UniqSubnetId must be passed.
	UniqVpcId *string `pulumi:"uniqVpcId"`
}

type RoInstanceIpState struct {
	// Read-only instance ID, in the format: cdbro-3i70uj0k, which is the same as the read-only instance ID displayed on the cloud database console page.
	InstanceId pulumi.StringPtrInput
	// Intranet IP address of the read-only instance.
	RoVip pulumi.StringPtrInput
	// Intranet port number of the read-only instance.
	RoVport pulumi.IntPtrInput
	// Subnet descriptor, for example: subnet-1typ0s7d.
	UniqSubnetId pulumi.StringPtrInput
	// vpc descriptor, for example: vpc-a23yt67j, if this field is passed, UniqSubnetId must be passed.
	UniqVpcId pulumi.StringPtrInput
}

func (RoInstanceIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*roInstanceIpState)(nil)).Elem()
}

type roInstanceIpArgs struct {
	// Read-only instance ID, in the format: cdbro-3i70uj0k, which is the same as the read-only instance ID displayed on the cloud database console page.
	InstanceId string `pulumi:"instanceId"`
	// Subnet descriptor, for example: subnet-1typ0s7d.
	UniqSubnetId *string `pulumi:"uniqSubnetId"`
	// vpc descriptor, for example: vpc-a23yt67j, if this field is passed, UniqSubnetId must be passed.
	UniqVpcId *string `pulumi:"uniqVpcId"`
}

// The set of arguments for constructing a RoInstanceIp resource.
type RoInstanceIpArgs struct {
	// Read-only instance ID, in the format: cdbro-3i70uj0k, which is the same as the read-only instance ID displayed on the cloud database console page.
	InstanceId pulumi.StringInput
	// Subnet descriptor, for example: subnet-1typ0s7d.
	UniqSubnetId pulumi.StringPtrInput
	// vpc descriptor, for example: vpc-a23yt67j, if this field is passed, UniqSubnetId must be passed.
	UniqVpcId pulumi.StringPtrInput
}

func (RoInstanceIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roInstanceIpArgs)(nil)).Elem()
}

type RoInstanceIpInput interface {
	pulumi.Input

	ToRoInstanceIpOutput() RoInstanceIpOutput
	ToRoInstanceIpOutputWithContext(ctx context.Context) RoInstanceIpOutput
}

func (*RoInstanceIp) ElementType() reflect.Type {
	return reflect.TypeOf((**RoInstanceIp)(nil)).Elem()
}

func (i *RoInstanceIp) ToRoInstanceIpOutput() RoInstanceIpOutput {
	return i.ToRoInstanceIpOutputWithContext(context.Background())
}

func (i *RoInstanceIp) ToRoInstanceIpOutputWithContext(ctx context.Context) RoInstanceIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoInstanceIpOutput)
}

// RoInstanceIpArrayInput is an input type that accepts RoInstanceIpArray and RoInstanceIpArrayOutput values.
// You can construct a concrete instance of `RoInstanceIpArrayInput` via:
//
//	RoInstanceIpArray{ RoInstanceIpArgs{...} }
type RoInstanceIpArrayInput interface {
	pulumi.Input

	ToRoInstanceIpArrayOutput() RoInstanceIpArrayOutput
	ToRoInstanceIpArrayOutputWithContext(context.Context) RoInstanceIpArrayOutput
}

type RoInstanceIpArray []RoInstanceIpInput

func (RoInstanceIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoInstanceIp)(nil)).Elem()
}

func (i RoInstanceIpArray) ToRoInstanceIpArrayOutput() RoInstanceIpArrayOutput {
	return i.ToRoInstanceIpArrayOutputWithContext(context.Background())
}

func (i RoInstanceIpArray) ToRoInstanceIpArrayOutputWithContext(ctx context.Context) RoInstanceIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoInstanceIpArrayOutput)
}

// RoInstanceIpMapInput is an input type that accepts RoInstanceIpMap and RoInstanceIpMapOutput values.
// You can construct a concrete instance of `RoInstanceIpMapInput` via:
//
//	RoInstanceIpMap{ "key": RoInstanceIpArgs{...} }
type RoInstanceIpMapInput interface {
	pulumi.Input

	ToRoInstanceIpMapOutput() RoInstanceIpMapOutput
	ToRoInstanceIpMapOutputWithContext(context.Context) RoInstanceIpMapOutput
}

type RoInstanceIpMap map[string]RoInstanceIpInput

func (RoInstanceIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoInstanceIp)(nil)).Elem()
}

func (i RoInstanceIpMap) ToRoInstanceIpMapOutput() RoInstanceIpMapOutput {
	return i.ToRoInstanceIpMapOutputWithContext(context.Background())
}

func (i RoInstanceIpMap) ToRoInstanceIpMapOutputWithContext(ctx context.Context) RoInstanceIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoInstanceIpMapOutput)
}

type RoInstanceIpOutput struct{ *pulumi.OutputState }

func (RoInstanceIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoInstanceIp)(nil)).Elem()
}

func (o RoInstanceIpOutput) ToRoInstanceIpOutput() RoInstanceIpOutput {
	return o
}

func (o RoInstanceIpOutput) ToRoInstanceIpOutputWithContext(ctx context.Context) RoInstanceIpOutput {
	return o
}

// Read-only instance ID, in the format: cdbro-3i70uj0k, which is the same as the read-only instance ID displayed on the cloud database console page.
func (o RoInstanceIpOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RoInstanceIp) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Intranet IP address of the read-only instance.
func (o RoInstanceIpOutput) RoVip() pulumi.StringOutput {
	return o.ApplyT(func(v *RoInstanceIp) pulumi.StringOutput { return v.RoVip }).(pulumi.StringOutput)
}

// Intranet port number of the read-only instance.
func (o RoInstanceIpOutput) RoVport() pulumi.IntOutput {
	return o.ApplyT(func(v *RoInstanceIp) pulumi.IntOutput { return v.RoVport }).(pulumi.IntOutput)
}

// Subnet descriptor, for example: subnet-1typ0s7d.
func (o RoInstanceIpOutput) UniqSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoInstanceIp) pulumi.StringPtrOutput { return v.UniqSubnetId }).(pulumi.StringPtrOutput)
}

// vpc descriptor, for example: vpc-a23yt67j, if this field is passed, UniqSubnetId must be passed.
func (o RoInstanceIpOutput) UniqVpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoInstanceIp) pulumi.StringPtrOutput { return v.UniqVpcId }).(pulumi.StringPtrOutput)
}

type RoInstanceIpArrayOutput struct{ *pulumi.OutputState }

func (RoInstanceIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoInstanceIp)(nil)).Elem()
}

func (o RoInstanceIpArrayOutput) ToRoInstanceIpArrayOutput() RoInstanceIpArrayOutput {
	return o
}

func (o RoInstanceIpArrayOutput) ToRoInstanceIpArrayOutputWithContext(ctx context.Context) RoInstanceIpArrayOutput {
	return o
}

func (o RoInstanceIpArrayOutput) Index(i pulumi.IntInput) RoInstanceIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RoInstanceIp {
		return vs[0].([]*RoInstanceIp)[vs[1].(int)]
	}).(RoInstanceIpOutput)
}

type RoInstanceIpMapOutput struct{ *pulumi.OutputState }

func (RoInstanceIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoInstanceIp)(nil)).Elem()
}

func (o RoInstanceIpMapOutput) ToRoInstanceIpMapOutput() RoInstanceIpMapOutput {
	return o
}

func (o RoInstanceIpMapOutput) ToRoInstanceIpMapOutputWithContext(ctx context.Context) RoInstanceIpMapOutput {
	return o
}

func (o RoInstanceIpMapOutput) MapIndex(k pulumi.StringInput) RoInstanceIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RoInstanceIp {
		return vs[0].(map[string]*RoInstanceIp)[vs[1].(string)]
	}).(RoInstanceIpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoInstanceIpInput)(nil)).Elem(), &RoInstanceIp{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoInstanceIpArrayInput)(nil)).Elem(), RoInstanceIpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoInstanceIpMapInput)(nil)).Elem(), RoInstanceIpMap{})
	pulumi.RegisterOutputType(RoInstanceIpOutput{})
	pulumi.RegisterOutputType(RoInstanceIpArrayOutput{})
	pulumi.RegisterOutputType(RoInstanceIpMapOutput{})
}
