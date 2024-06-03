// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a mysql backupEncryptionStatus
//
// ## Example Usage
//
// ### Enable encryption
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
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
//				SlaveDeployMode:  pulumi.Int(0),
//				AvailabilityZone: pulumi.String(zones.Zones[0].Name),
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
//				Tags: pulumi.Map{
//					"name": pulumi.Any("test"),
//				},
//				Parameters: pulumi.Map{
//					"character_set_server": pulumi.Any("utf8"),
//					"max_connections":      pulumi.Any("1000"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Mysql.NewBackupEncryptionStatus(ctx, "exampleBackupEncryptionStatus", &Mysql.BackupEncryptionStatusArgs{
//				InstanceId:       exampleInstance.ID(),
//				EncryptionStatus: pulumi.String("on"),
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
// ### Disable encryption
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mysql.NewBackupEncryptionStatus(ctx, "example", &Mysql.BackupEncryptionStatusArgs{
//				InstanceId:       pulumi.Any(tencentcloud_mysql_instance.Example.Id),
//				EncryptionStatus: pulumi.String("off"),
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
// mysql backup_encryption_status can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Mysql/backupEncryptionStatus:BackupEncryptionStatus backup_encryption_status backup_encryption_status_id
// ```
type BackupEncryptionStatus struct {
	pulumi.CustomResourceState

	// Whether physical backup encryption is enabled for the instance. Possible values are `on`, `off`.
	EncryptionStatus pulumi.StringOutput `pulumi:"encryptionStatus"`
	// Instance ID, in the format: cdb-XXXX. Same instance ID as displayed in the ApsaraDB for Console page.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewBackupEncryptionStatus registers a new resource with the given unique name, arguments, and options.
func NewBackupEncryptionStatus(ctx *pulumi.Context,
	name string, args *BackupEncryptionStatusArgs, opts ...pulumi.ResourceOption) (*BackupEncryptionStatus, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EncryptionStatus == nil {
		return nil, errors.New("invalid value for required argument 'EncryptionStatus'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackupEncryptionStatus
	err := ctx.RegisterResource("tencentcloud:Mysql/backupEncryptionStatus:BackupEncryptionStatus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupEncryptionStatus gets an existing BackupEncryptionStatus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupEncryptionStatus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupEncryptionStatusState, opts ...pulumi.ResourceOption) (*BackupEncryptionStatus, error) {
	var resource BackupEncryptionStatus
	err := ctx.ReadResource("tencentcloud:Mysql/backupEncryptionStatus:BackupEncryptionStatus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupEncryptionStatus resources.
type backupEncryptionStatusState struct {
	// Whether physical backup encryption is enabled for the instance. Possible values are `on`, `off`.
	EncryptionStatus *string `pulumi:"encryptionStatus"`
	// Instance ID, in the format: cdb-XXXX. Same instance ID as displayed in the ApsaraDB for Console page.
	InstanceId *string `pulumi:"instanceId"`
}

type BackupEncryptionStatusState struct {
	// Whether physical backup encryption is enabled for the instance. Possible values are `on`, `off`.
	EncryptionStatus pulumi.StringPtrInput
	// Instance ID, in the format: cdb-XXXX. Same instance ID as displayed in the ApsaraDB for Console page.
	InstanceId pulumi.StringPtrInput
}

func (BackupEncryptionStatusState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupEncryptionStatusState)(nil)).Elem()
}

type backupEncryptionStatusArgs struct {
	// Whether physical backup encryption is enabled for the instance. Possible values are `on`, `off`.
	EncryptionStatus string `pulumi:"encryptionStatus"`
	// Instance ID, in the format: cdb-XXXX. Same instance ID as displayed in the ApsaraDB for Console page.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a BackupEncryptionStatus resource.
type BackupEncryptionStatusArgs struct {
	// Whether physical backup encryption is enabled for the instance. Possible values are `on`, `off`.
	EncryptionStatus pulumi.StringInput
	// Instance ID, in the format: cdb-XXXX. Same instance ID as displayed in the ApsaraDB for Console page.
	InstanceId pulumi.StringInput
}

func (BackupEncryptionStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupEncryptionStatusArgs)(nil)).Elem()
}

type BackupEncryptionStatusInput interface {
	pulumi.Input

	ToBackupEncryptionStatusOutput() BackupEncryptionStatusOutput
	ToBackupEncryptionStatusOutputWithContext(ctx context.Context) BackupEncryptionStatusOutput
}

func (*BackupEncryptionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupEncryptionStatus)(nil)).Elem()
}

func (i *BackupEncryptionStatus) ToBackupEncryptionStatusOutput() BackupEncryptionStatusOutput {
	return i.ToBackupEncryptionStatusOutputWithContext(context.Background())
}

func (i *BackupEncryptionStatus) ToBackupEncryptionStatusOutputWithContext(ctx context.Context) BackupEncryptionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupEncryptionStatusOutput)
}

// BackupEncryptionStatusArrayInput is an input type that accepts BackupEncryptionStatusArray and BackupEncryptionStatusArrayOutput values.
// You can construct a concrete instance of `BackupEncryptionStatusArrayInput` via:
//
//	BackupEncryptionStatusArray{ BackupEncryptionStatusArgs{...} }
type BackupEncryptionStatusArrayInput interface {
	pulumi.Input

	ToBackupEncryptionStatusArrayOutput() BackupEncryptionStatusArrayOutput
	ToBackupEncryptionStatusArrayOutputWithContext(context.Context) BackupEncryptionStatusArrayOutput
}

type BackupEncryptionStatusArray []BackupEncryptionStatusInput

func (BackupEncryptionStatusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupEncryptionStatus)(nil)).Elem()
}

func (i BackupEncryptionStatusArray) ToBackupEncryptionStatusArrayOutput() BackupEncryptionStatusArrayOutput {
	return i.ToBackupEncryptionStatusArrayOutputWithContext(context.Background())
}

func (i BackupEncryptionStatusArray) ToBackupEncryptionStatusArrayOutputWithContext(ctx context.Context) BackupEncryptionStatusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupEncryptionStatusArrayOutput)
}

// BackupEncryptionStatusMapInput is an input type that accepts BackupEncryptionStatusMap and BackupEncryptionStatusMapOutput values.
// You can construct a concrete instance of `BackupEncryptionStatusMapInput` via:
//
//	BackupEncryptionStatusMap{ "key": BackupEncryptionStatusArgs{...} }
type BackupEncryptionStatusMapInput interface {
	pulumi.Input

	ToBackupEncryptionStatusMapOutput() BackupEncryptionStatusMapOutput
	ToBackupEncryptionStatusMapOutputWithContext(context.Context) BackupEncryptionStatusMapOutput
}

type BackupEncryptionStatusMap map[string]BackupEncryptionStatusInput

func (BackupEncryptionStatusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupEncryptionStatus)(nil)).Elem()
}

func (i BackupEncryptionStatusMap) ToBackupEncryptionStatusMapOutput() BackupEncryptionStatusMapOutput {
	return i.ToBackupEncryptionStatusMapOutputWithContext(context.Background())
}

func (i BackupEncryptionStatusMap) ToBackupEncryptionStatusMapOutputWithContext(ctx context.Context) BackupEncryptionStatusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupEncryptionStatusMapOutput)
}

type BackupEncryptionStatusOutput struct{ *pulumi.OutputState }

func (BackupEncryptionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupEncryptionStatus)(nil)).Elem()
}

func (o BackupEncryptionStatusOutput) ToBackupEncryptionStatusOutput() BackupEncryptionStatusOutput {
	return o
}

func (o BackupEncryptionStatusOutput) ToBackupEncryptionStatusOutputWithContext(ctx context.Context) BackupEncryptionStatusOutput {
	return o
}

// Whether physical backup encryption is enabled for the instance. Possible values are `on`, `off`.
func (o BackupEncryptionStatusOutput) EncryptionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupEncryptionStatus) pulumi.StringOutput { return v.EncryptionStatus }).(pulumi.StringOutput)
}

// Instance ID, in the format: cdb-XXXX. Same instance ID as displayed in the ApsaraDB for Console page.
func (o BackupEncryptionStatusOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupEncryptionStatus) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type BackupEncryptionStatusArrayOutput struct{ *pulumi.OutputState }

func (BackupEncryptionStatusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupEncryptionStatus)(nil)).Elem()
}

func (o BackupEncryptionStatusArrayOutput) ToBackupEncryptionStatusArrayOutput() BackupEncryptionStatusArrayOutput {
	return o
}

func (o BackupEncryptionStatusArrayOutput) ToBackupEncryptionStatusArrayOutputWithContext(ctx context.Context) BackupEncryptionStatusArrayOutput {
	return o
}

func (o BackupEncryptionStatusArrayOutput) Index(i pulumi.IntInput) BackupEncryptionStatusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupEncryptionStatus {
		return vs[0].([]*BackupEncryptionStatus)[vs[1].(int)]
	}).(BackupEncryptionStatusOutput)
}

type BackupEncryptionStatusMapOutput struct{ *pulumi.OutputState }

func (BackupEncryptionStatusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupEncryptionStatus)(nil)).Elem()
}

func (o BackupEncryptionStatusMapOutput) ToBackupEncryptionStatusMapOutput() BackupEncryptionStatusMapOutput {
	return o
}

func (o BackupEncryptionStatusMapOutput) ToBackupEncryptionStatusMapOutputWithContext(ctx context.Context) BackupEncryptionStatusMapOutput {
	return o
}

func (o BackupEncryptionStatusMapOutput) MapIndex(k pulumi.StringInput) BackupEncryptionStatusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupEncryptionStatus {
		return vs[0].(map[string]*BackupEncryptionStatus)[vs[1].(string)]
	}).(BackupEncryptionStatusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupEncryptionStatusInput)(nil)).Elem(), &BackupEncryptionStatus{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupEncryptionStatusArrayInput)(nil)).Elem(), BackupEncryptionStatusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupEncryptionStatusMapInput)(nil)).Elem(), BackupEncryptionStatusMap{})
	pulumi.RegisterOutputType(BackupEncryptionStatusOutput{})
	pulumi.RegisterOutputType(BackupEncryptionStatusArrayOutput{})
	pulumi.RegisterOutputType(BackupEncryptionStatusMapOutput{})
}
