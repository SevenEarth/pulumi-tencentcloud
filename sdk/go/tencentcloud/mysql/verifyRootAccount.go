// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mysql verifyRootAccount
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Security"
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
//			securityGroup, err := Security.NewGroup(ctx, "securityGroup", &Security.GroupArgs{
//				Description: pulumi.String("mysql test"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleInstance, err := Mysql.NewInstance(ctx, "exampleInstance", &Mysql.InstanceArgs{
//				InternetService:  pulumi.Int(1),
//				EngineVersion:    pulumi.String("5.7"),
//				ChargeType:       pulumi.String("POSTPAID"),
//				RootPassword:     pulumi.String("PassWord123"),
//				SlaveDeployMode:  pulumi.Int(1),
//				AvailabilityZone: pulumi.String(zones.Zones[0].Name),
//				FirstSlaveZone:   pulumi.String(zones.Zones[1].Name),
//				SlaveSyncMode:    pulumi.Int(1),
//				InstanceName:     pulumi.String("tf-example-mysql"),
//				MemSize:          pulumi.Int(4000),
//				VolumeSize:       pulumi.Int(200),
//				VpcId:            vpc.ID(),
//				SubnetId:         subnet.ID(),
//				IntranetPort:     pulumi.Int(3306),
//				SecurityGroups: pulumi.StringArray{
//					securityGroup.ID(),
//				},
//				Tags: pulumi.AnyMap{
//					"name": pulumi.Any("test"),
//				},
//				Parameters: pulumi.AnyMap{
//					"character_set_server": pulumi.Any("utf8"),
//					"max_connections":      pulumi.Any("1000"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Mysql.NewVerifyRootAccount(ctx, "exampleVerifyRootAccount", &Mysql.VerifyRootAccountArgs{
//				InstanceId: exampleInstance.ID(),
//				Password:   pulumi.String("Qwer@234"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type VerifyRootAccount struct {
	pulumi.CustomResourceState

	// The instance ID, in the format: cdb-c1nl9rpv, is the same as the instance ID displayed on the cloud database console page.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The password of the ROOT account of the instance.
	Password pulumi.StringOutput `pulumi:"password"`
}

// NewVerifyRootAccount registers a new resource with the given unique name, arguments, and options.
func NewVerifyRootAccount(ctx *pulumi.Context,
	name string, args *VerifyRootAccountArgs, opts ...pulumi.ResourceOption) (*VerifyRootAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource VerifyRootAccount
	err := ctx.RegisterResource("tencentcloud:Mysql/verifyRootAccount:VerifyRootAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVerifyRootAccount gets an existing VerifyRootAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVerifyRootAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VerifyRootAccountState, opts ...pulumi.ResourceOption) (*VerifyRootAccount, error) {
	var resource VerifyRootAccount
	err := ctx.ReadResource("tencentcloud:Mysql/verifyRootAccount:VerifyRootAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VerifyRootAccount resources.
type verifyRootAccountState struct {
	// The instance ID, in the format: cdb-c1nl9rpv, is the same as the instance ID displayed on the cloud database console page.
	InstanceId *string `pulumi:"instanceId"`
	// The password of the ROOT account of the instance.
	Password *string `pulumi:"password"`
}

type VerifyRootAccountState struct {
	// The instance ID, in the format: cdb-c1nl9rpv, is the same as the instance ID displayed on the cloud database console page.
	InstanceId pulumi.StringPtrInput
	// The password of the ROOT account of the instance.
	Password pulumi.StringPtrInput
}

func (VerifyRootAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*verifyRootAccountState)(nil)).Elem()
}

type verifyRootAccountArgs struct {
	// The instance ID, in the format: cdb-c1nl9rpv, is the same as the instance ID displayed on the cloud database console page.
	InstanceId string `pulumi:"instanceId"`
	// The password of the ROOT account of the instance.
	Password string `pulumi:"password"`
}

// The set of arguments for constructing a VerifyRootAccount resource.
type VerifyRootAccountArgs struct {
	// The instance ID, in the format: cdb-c1nl9rpv, is the same as the instance ID displayed on the cloud database console page.
	InstanceId pulumi.StringInput
	// The password of the ROOT account of the instance.
	Password pulumi.StringInput
}

func (VerifyRootAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*verifyRootAccountArgs)(nil)).Elem()
}

type VerifyRootAccountInput interface {
	pulumi.Input

	ToVerifyRootAccountOutput() VerifyRootAccountOutput
	ToVerifyRootAccountOutputWithContext(ctx context.Context) VerifyRootAccountOutput
}

func (*VerifyRootAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**VerifyRootAccount)(nil)).Elem()
}

func (i *VerifyRootAccount) ToVerifyRootAccountOutput() VerifyRootAccountOutput {
	return i.ToVerifyRootAccountOutputWithContext(context.Background())
}

func (i *VerifyRootAccount) ToVerifyRootAccountOutputWithContext(ctx context.Context) VerifyRootAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VerifyRootAccountOutput)
}

// VerifyRootAccountArrayInput is an input type that accepts VerifyRootAccountArray and VerifyRootAccountArrayOutput values.
// You can construct a concrete instance of `VerifyRootAccountArrayInput` via:
//
//	VerifyRootAccountArray{ VerifyRootAccountArgs{...} }
type VerifyRootAccountArrayInput interface {
	pulumi.Input

	ToVerifyRootAccountArrayOutput() VerifyRootAccountArrayOutput
	ToVerifyRootAccountArrayOutputWithContext(context.Context) VerifyRootAccountArrayOutput
}

type VerifyRootAccountArray []VerifyRootAccountInput

func (VerifyRootAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VerifyRootAccount)(nil)).Elem()
}

func (i VerifyRootAccountArray) ToVerifyRootAccountArrayOutput() VerifyRootAccountArrayOutput {
	return i.ToVerifyRootAccountArrayOutputWithContext(context.Background())
}

func (i VerifyRootAccountArray) ToVerifyRootAccountArrayOutputWithContext(ctx context.Context) VerifyRootAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VerifyRootAccountArrayOutput)
}

// VerifyRootAccountMapInput is an input type that accepts VerifyRootAccountMap and VerifyRootAccountMapOutput values.
// You can construct a concrete instance of `VerifyRootAccountMapInput` via:
//
//	VerifyRootAccountMap{ "key": VerifyRootAccountArgs{...} }
type VerifyRootAccountMapInput interface {
	pulumi.Input

	ToVerifyRootAccountMapOutput() VerifyRootAccountMapOutput
	ToVerifyRootAccountMapOutputWithContext(context.Context) VerifyRootAccountMapOutput
}

type VerifyRootAccountMap map[string]VerifyRootAccountInput

func (VerifyRootAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VerifyRootAccount)(nil)).Elem()
}

func (i VerifyRootAccountMap) ToVerifyRootAccountMapOutput() VerifyRootAccountMapOutput {
	return i.ToVerifyRootAccountMapOutputWithContext(context.Background())
}

func (i VerifyRootAccountMap) ToVerifyRootAccountMapOutputWithContext(ctx context.Context) VerifyRootAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VerifyRootAccountMapOutput)
}

type VerifyRootAccountOutput struct{ *pulumi.OutputState }

func (VerifyRootAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VerifyRootAccount)(nil)).Elem()
}

func (o VerifyRootAccountOutput) ToVerifyRootAccountOutput() VerifyRootAccountOutput {
	return o
}

func (o VerifyRootAccountOutput) ToVerifyRootAccountOutputWithContext(ctx context.Context) VerifyRootAccountOutput {
	return o
}

// The instance ID, in the format: cdb-c1nl9rpv, is the same as the instance ID displayed on the cloud database console page.
func (o VerifyRootAccountOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VerifyRootAccount) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The password of the ROOT account of the instance.
func (o VerifyRootAccountOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *VerifyRootAccount) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

type VerifyRootAccountArrayOutput struct{ *pulumi.OutputState }

func (VerifyRootAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VerifyRootAccount)(nil)).Elem()
}

func (o VerifyRootAccountArrayOutput) ToVerifyRootAccountArrayOutput() VerifyRootAccountArrayOutput {
	return o
}

func (o VerifyRootAccountArrayOutput) ToVerifyRootAccountArrayOutputWithContext(ctx context.Context) VerifyRootAccountArrayOutput {
	return o
}

func (o VerifyRootAccountArrayOutput) Index(i pulumi.IntInput) VerifyRootAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VerifyRootAccount {
		return vs[0].([]*VerifyRootAccount)[vs[1].(int)]
	}).(VerifyRootAccountOutput)
}

type VerifyRootAccountMapOutput struct{ *pulumi.OutputState }

func (VerifyRootAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VerifyRootAccount)(nil)).Elem()
}

func (o VerifyRootAccountMapOutput) ToVerifyRootAccountMapOutput() VerifyRootAccountMapOutput {
	return o
}

func (o VerifyRootAccountMapOutput) ToVerifyRootAccountMapOutputWithContext(ctx context.Context) VerifyRootAccountMapOutput {
	return o
}

func (o VerifyRootAccountMapOutput) MapIndex(k pulumi.StringInput) VerifyRootAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VerifyRootAccount {
		return vs[0].(map[string]*VerifyRootAccount)[vs[1].(string)]
	}).(VerifyRootAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VerifyRootAccountInput)(nil)).Elem(), &VerifyRootAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*VerifyRootAccountArrayInput)(nil)).Elem(), VerifyRootAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VerifyRootAccountMapInput)(nil)).Elem(), VerifyRootAccountMap{})
	pulumi.RegisterOutputType(VerifyRootAccountOutput{})
	pulumi.RegisterOutputType(VerifyRootAccountArrayOutput{})
	pulumi.RegisterOutputType(VerifyRootAccountMapOutput{})
}
