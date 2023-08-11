// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mysql instanceEncryptionOperation
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Security"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		zones, err := Availability.GetZonesByProduct(ctx, &availability.GetZonesByProductArgs{
// 			Product: "cdb",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		vpc, err := Vpc.NewInstance(ctx, "vpc", &Vpc.InstanceArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		subnet, err := Subnet.NewInstance(ctx, "subnet", &Subnet.InstanceArgs{
// 			AvailabilityZone: pulumi.String(zones.Zones[0].Name),
// 			VpcId:            vpc.ID(),
// 			CidrBlock:        pulumi.String("10.0.0.0/16"),
// 			IsMulticast:      pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		securityGroup, err := Security.NewGroup(ctx, "securityGroup", &Security.GroupArgs{
// 			Description: pulumi.String("mysql test"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleInstance, err := Mysql.NewInstance(ctx, "exampleInstance", &Mysql.InstanceArgs{
// 			InternetService:  pulumi.Int(1),
// 			EngineVersion:    pulumi.String("5.7"),
// 			ChargeType:       pulumi.String("POSTPAID"),
// 			RootPassword:     pulumi.String("PassWord123"),
// 			SlaveDeployMode:  pulumi.Int(0),
// 			AvailabilityZone: pulumi.String(zones.Zones[0].Name),
// 			SlaveSyncMode:    pulumi.Int(1),
// 			InstanceName:     pulumi.String("tf-example-mysql"),
// 			MemSize:          pulumi.Int(4000),
// 			VolumeSize:       pulumi.Int(200),
// 			VpcId:            vpc.ID(),
// 			SubnetId:         subnet.ID(),
// 			IntranetPort:     pulumi.Int(3306),
// 			SecurityGroups: pulumi.StringArray{
// 				securityGroup.ID(),
// 			},
// 			Tags: pulumi.AnyMap{
// 				"name": pulumi.Any("test"),
// 			},
// 			Parameters: pulumi.AnyMap{
// 				"character_set_server": pulumi.Any("utf8"),
// 				"max_connections":      pulumi.Any("1000"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Mysql.NewInstanceEncryptionOperation(ctx, "exampleInstanceEncryptionOperation", &Mysql.InstanceEncryptionOperationArgs{
// 			InstanceId: exampleInstance.ID(),
// 			KeyId:      pulumi.String("KMS-CDB"),
// 			KeyRegion:  pulumi.String("ap-guangzhou"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type InstanceEncryptionOperation struct {
	pulumi.CustomResourceState

	// TencentDB instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Custom key ID, which is the unique CMK ID. If this value is empty, the key KMS-CDB auto-generated by Tencent Cloud will be used.
	KeyId pulumi.StringPtrOutput `pulumi:"keyId"`
	// Custom storage region, such as ap-guangzhou. When `KeyId` is not empty, this parameter is required.
	KeyRegion pulumi.StringPtrOutput `pulumi:"keyRegion"`
}

// NewInstanceEncryptionOperation registers a new resource with the given unique name, arguments, and options.
func NewInstanceEncryptionOperation(ctx *pulumi.Context,
	name string, args *InstanceEncryptionOperationArgs, opts ...pulumi.ResourceOption) (*InstanceEncryptionOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource InstanceEncryptionOperation
	err := ctx.RegisterResource("tencentcloud:Mysql/instanceEncryptionOperation:InstanceEncryptionOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceEncryptionOperation gets an existing InstanceEncryptionOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceEncryptionOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceEncryptionOperationState, opts ...pulumi.ResourceOption) (*InstanceEncryptionOperation, error) {
	var resource InstanceEncryptionOperation
	err := ctx.ReadResource("tencentcloud:Mysql/instanceEncryptionOperation:InstanceEncryptionOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceEncryptionOperation resources.
type instanceEncryptionOperationState struct {
	// TencentDB instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// Custom key ID, which is the unique CMK ID. If this value is empty, the key KMS-CDB auto-generated by Tencent Cloud will be used.
	KeyId *string `pulumi:"keyId"`
	// Custom storage region, such as ap-guangzhou. When `KeyId` is not empty, this parameter is required.
	KeyRegion *string `pulumi:"keyRegion"`
}

type InstanceEncryptionOperationState struct {
	// TencentDB instance ID.
	InstanceId pulumi.StringPtrInput
	// Custom key ID, which is the unique CMK ID. If this value is empty, the key KMS-CDB auto-generated by Tencent Cloud will be used.
	KeyId pulumi.StringPtrInput
	// Custom storage region, such as ap-guangzhou. When `KeyId` is not empty, this parameter is required.
	KeyRegion pulumi.StringPtrInput
}

func (InstanceEncryptionOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceEncryptionOperationState)(nil)).Elem()
}

type instanceEncryptionOperationArgs struct {
	// TencentDB instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Custom key ID, which is the unique CMK ID. If this value is empty, the key KMS-CDB auto-generated by Tencent Cloud will be used.
	KeyId *string `pulumi:"keyId"`
	// Custom storage region, such as ap-guangzhou. When `KeyId` is not empty, this parameter is required.
	KeyRegion *string `pulumi:"keyRegion"`
}

// The set of arguments for constructing a InstanceEncryptionOperation resource.
type InstanceEncryptionOperationArgs struct {
	// TencentDB instance ID.
	InstanceId pulumi.StringInput
	// Custom key ID, which is the unique CMK ID. If this value is empty, the key KMS-CDB auto-generated by Tencent Cloud will be used.
	KeyId pulumi.StringPtrInput
	// Custom storage region, such as ap-guangzhou. When `KeyId` is not empty, this parameter is required.
	KeyRegion pulumi.StringPtrInput
}

func (InstanceEncryptionOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceEncryptionOperationArgs)(nil)).Elem()
}

type InstanceEncryptionOperationInput interface {
	pulumi.Input

	ToInstanceEncryptionOperationOutput() InstanceEncryptionOperationOutput
	ToInstanceEncryptionOperationOutputWithContext(ctx context.Context) InstanceEncryptionOperationOutput
}

func (*InstanceEncryptionOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceEncryptionOperation)(nil)).Elem()
}

func (i *InstanceEncryptionOperation) ToInstanceEncryptionOperationOutput() InstanceEncryptionOperationOutput {
	return i.ToInstanceEncryptionOperationOutputWithContext(context.Background())
}

func (i *InstanceEncryptionOperation) ToInstanceEncryptionOperationOutputWithContext(ctx context.Context) InstanceEncryptionOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceEncryptionOperationOutput)
}

// InstanceEncryptionOperationArrayInput is an input type that accepts InstanceEncryptionOperationArray and InstanceEncryptionOperationArrayOutput values.
// You can construct a concrete instance of `InstanceEncryptionOperationArrayInput` via:
//
//          InstanceEncryptionOperationArray{ InstanceEncryptionOperationArgs{...} }
type InstanceEncryptionOperationArrayInput interface {
	pulumi.Input

	ToInstanceEncryptionOperationArrayOutput() InstanceEncryptionOperationArrayOutput
	ToInstanceEncryptionOperationArrayOutputWithContext(context.Context) InstanceEncryptionOperationArrayOutput
}

type InstanceEncryptionOperationArray []InstanceEncryptionOperationInput

func (InstanceEncryptionOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceEncryptionOperation)(nil)).Elem()
}

func (i InstanceEncryptionOperationArray) ToInstanceEncryptionOperationArrayOutput() InstanceEncryptionOperationArrayOutput {
	return i.ToInstanceEncryptionOperationArrayOutputWithContext(context.Background())
}

func (i InstanceEncryptionOperationArray) ToInstanceEncryptionOperationArrayOutputWithContext(ctx context.Context) InstanceEncryptionOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceEncryptionOperationArrayOutput)
}

// InstanceEncryptionOperationMapInput is an input type that accepts InstanceEncryptionOperationMap and InstanceEncryptionOperationMapOutput values.
// You can construct a concrete instance of `InstanceEncryptionOperationMapInput` via:
//
//          InstanceEncryptionOperationMap{ "key": InstanceEncryptionOperationArgs{...} }
type InstanceEncryptionOperationMapInput interface {
	pulumi.Input

	ToInstanceEncryptionOperationMapOutput() InstanceEncryptionOperationMapOutput
	ToInstanceEncryptionOperationMapOutputWithContext(context.Context) InstanceEncryptionOperationMapOutput
}

type InstanceEncryptionOperationMap map[string]InstanceEncryptionOperationInput

func (InstanceEncryptionOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceEncryptionOperation)(nil)).Elem()
}

func (i InstanceEncryptionOperationMap) ToInstanceEncryptionOperationMapOutput() InstanceEncryptionOperationMapOutput {
	return i.ToInstanceEncryptionOperationMapOutputWithContext(context.Background())
}

func (i InstanceEncryptionOperationMap) ToInstanceEncryptionOperationMapOutputWithContext(ctx context.Context) InstanceEncryptionOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceEncryptionOperationMapOutput)
}

type InstanceEncryptionOperationOutput struct{ *pulumi.OutputState }

func (InstanceEncryptionOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceEncryptionOperation)(nil)).Elem()
}

func (o InstanceEncryptionOperationOutput) ToInstanceEncryptionOperationOutput() InstanceEncryptionOperationOutput {
	return o
}

func (o InstanceEncryptionOperationOutput) ToInstanceEncryptionOperationOutputWithContext(ctx context.Context) InstanceEncryptionOperationOutput {
	return o
}

// TencentDB instance ID.
func (o InstanceEncryptionOperationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceEncryptionOperation) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Custom key ID, which is the unique CMK ID. If this value is empty, the key KMS-CDB auto-generated by Tencent Cloud will be used.
func (o InstanceEncryptionOperationOutput) KeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceEncryptionOperation) pulumi.StringPtrOutput { return v.KeyId }).(pulumi.StringPtrOutput)
}

// Custom storage region, such as ap-guangzhou. When `KeyId` is not empty, this parameter is required.
func (o InstanceEncryptionOperationOutput) KeyRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceEncryptionOperation) pulumi.StringPtrOutput { return v.KeyRegion }).(pulumi.StringPtrOutput)
}

type InstanceEncryptionOperationArrayOutput struct{ *pulumi.OutputState }

func (InstanceEncryptionOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceEncryptionOperation)(nil)).Elem()
}

func (o InstanceEncryptionOperationArrayOutput) ToInstanceEncryptionOperationArrayOutput() InstanceEncryptionOperationArrayOutput {
	return o
}

func (o InstanceEncryptionOperationArrayOutput) ToInstanceEncryptionOperationArrayOutputWithContext(ctx context.Context) InstanceEncryptionOperationArrayOutput {
	return o
}

func (o InstanceEncryptionOperationArrayOutput) Index(i pulumi.IntInput) InstanceEncryptionOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceEncryptionOperation {
		return vs[0].([]*InstanceEncryptionOperation)[vs[1].(int)]
	}).(InstanceEncryptionOperationOutput)
}

type InstanceEncryptionOperationMapOutput struct{ *pulumi.OutputState }

func (InstanceEncryptionOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceEncryptionOperation)(nil)).Elem()
}

func (o InstanceEncryptionOperationMapOutput) ToInstanceEncryptionOperationMapOutput() InstanceEncryptionOperationMapOutput {
	return o
}

func (o InstanceEncryptionOperationMapOutput) ToInstanceEncryptionOperationMapOutputWithContext(ctx context.Context) InstanceEncryptionOperationMapOutput {
	return o
}

func (o InstanceEncryptionOperationMapOutput) MapIndex(k pulumi.StringInput) InstanceEncryptionOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceEncryptionOperation {
		return vs[0].(map[string]*InstanceEncryptionOperation)[vs[1].(string)]
	}).(InstanceEncryptionOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceEncryptionOperationInput)(nil)).Elem(), &InstanceEncryptionOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceEncryptionOperationArrayInput)(nil)).Elem(), InstanceEncryptionOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceEncryptionOperationMapInput)(nil)).Elem(), InstanceEncryptionOperationMap{})
	pulumi.RegisterOutputType(InstanceEncryptionOperationOutput{})
	pulumi.RegisterOutputType(InstanceEncryptionOperationArrayOutput{})
	pulumi.RegisterOutputType(InstanceEncryptionOperationMapOutput{})
}