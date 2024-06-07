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
//			_, err = Mysql.NewBackupPolicy(ctx, "exampleBackupPolicy", &Mysql.BackupPolicyArgs{
//				MysqlId:             exampleInstance.ID(),
//				RetentionPeriod:     pulumi.Int(7),
//				BackupModel:         pulumi.String("physical"),
//				BackupTime:          pulumi.String("22:00-02:00"),
//				BinlogPeriod:        pulumi.Int(32),
//				EnableBinlogStandby: pulumi.String("off"),
//				BinlogStandbyDays:   pulumi.Int(31),
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
type BackupPolicy struct {
	pulumi.CustomResourceState

	// Backup method. Supported values include: `physical` - physical backup.
	BackupModel pulumi.StringPtrOutput `pulumi:"backupModel"`
	// Instance backup time, in the format of 'HH:mm-HH:mm'. Time setting interval is four hours. Default to `02:00-06:00`. The following value can be supported: `02:00-06:00`, `06:00-10:00`, `10:00-14:00`, `14:00-18:00`, `18:00-22:00`, and `22:00-02:00`.
	BackupTime pulumi.StringPtrOutput `pulumi:"backupTime"`
	// Binlog retention time, in days. The minimum value is 7 days and the maximum value is 1830 days. This value cannot be set greater than the backup file retention time.
	BinlogPeriod pulumi.IntOutput `pulumi:"binlogPeriod"`
	// The standard starting number of days for log backup storage. The log backup will be converted when it reaches the standard starting number of days for storage. The minimum is 30 days and must not be greater than the number of days for log backup retention.
	BinlogStandbyDays pulumi.IntOutput `pulumi:"binlogStandbyDays"`
	// Whether to enable the log backup standard storage policy, `off` - close, `on` - open, the default is off.
	EnableBinlogStandby pulumi.StringPtrOutput `pulumi:"enableBinlogStandby"`
	// Instance ID to which policies will be applied.
	MysqlId pulumi.StringOutput `pulumi:"mysqlId"`
	// The retention time of backup files, in days. The minimum value is 7 days and the maximum value is 1830 days. And default value is `7`.
	RetentionPeriod pulumi.IntPtrOutput `pulumi:"retentionPeriod"`
}

// NewBackupPolicy registers a new resource with the given unique name, arguments, and options.
func NewBackupPolicy(ctx *pulumi.Context,
	name string, args *BackupPolicyArgs, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MysqlId == nil {
		return nil, errors.New("invalid value for required argument 'MysqlId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackupPolicy
	err := ctx.RegisterResource("tencentcloud:Mysql/backupPolicy:BackupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupPolicy gets an existing BackupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupPolicyState, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	var resource BackupPolicy
	err := ctx.ReadResource("tencentcloud:Mysql/backupPolicy:BackupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupPolicy resources.
type backupPolicyState struct {
	// Backup method. Supported values include: `physical` - physical backup.
	BackupModel *string `pulumi:"backupModel"`
	// Instance backup time, in the format of 'HH:mm-HH:mm'. Time setting interval is four hours. Default to `02:00-06:00`. The following value can be supported: `02:00-06:00`, `06:00-10:00`, `10:00-14:00`, `14:00-18:00`, `18:00-22:00`, and `22:00-02:00`.
	BackupTime *string `pulumi:"backupTime"`
	// Binlog retention time, in days. The minimum value is 7 days and the maximum value is 1830 days. This value cannot be set greater than the backup file retention time.
	BinlogPeriod *int `pulumi:"binlogPeriod"`
	// The standard starting number of days for log backup storage. The log backup will be converted when it reaches the standard starting number of days for storage. The minimum is 30 days and must not be greater than the number of days for log backup retention.
	BinlogStandbyDays *int `pulumi:"binlogStandbyDays"`
	// Whether to enable the log backup standard storage policy, `off` - close, `on` - open, the default is off.
	EnableBinlogStandby *string `pulumi:"enableBinlogStandby"`
	// Instance ID to which policies will be applied.
	MysqlId *string `pulumi:"mysqlId"`
	// The retention time of backup files, in days. The minimum value is 7 days and the maximum value is 1830 days. And default value is `7`.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
}

type BackupPolicyState struct {
	// Backup method. Supported values include: `physical` - physical backup.
	BackupModel pulumi.StringPtrInput
	// Instance backup time, in the format of 'HH:mm-HH:mm'. Time setting interval is four hours. Default to `02:00-06:00`. The following value can be supported: `02:00-06:00`, `06:00-10:00`, `10:00-14:00`, `14:00-18:00`, `18:00-22:00`, and `22:00-02:00`.
	BackupTime pulumi.StringPtrInput
	// Binlog retention time, in days. The minimum value is 7 days and the maximum value is 1830 days. This value cannot be set greater than the backup file retention time.
	BinlogPeriod pulumi.IntPtrInput
	// The standard starting number of days for log backup storage. The log backup will be converted when it reaches the standard starting number of days for storage. The minimum is 30 days and must not be greater than the number of days for log backup retention.
	BinlogStandbyDays pulumi.IntPtrInput
	// Whether to enable the log backup standard storage policy, `off` - close, `on` - open, the default is off.
	EnableBinlogStandby pulumi.StringPtrInput
	// Instance ID to which policies will be applied.
	MysqlId pulumi.StringPtrInput
	// The retention time of backup files, in days. The minimum value is 7 days and the maximum value is 1830 days. And default value is `7`.
	RetentionPeriod pulumi.IntPtrInput
}

func (BackupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyState)(nil)).Elem()
}

type backupPolicyArgs struct {
	// Backup method. Supported values include: `physical` - physical backup.
	BackupModel *string `pulumi:"backupModel"`
	// Instance backup time, in the format of 'HH:mm-HH:mm'. Time setting interval is four hours. Default to `02:00-06:00`. The following value can be supported: `02:00-06:00`, `06:00-10:00`, `10:00-14:00`, `14:00-18:00`, `18:00-22:00`, and `22:00-02:00`.
	BackupTime *string `pulumi:"backupTime"`
	// Binlog retention time, in days. The minimum value is 7 days and the maximum value is 1830 days. This value cannot be set greater than the backup file retention time.
	BinlogPeriod *int `pulumi:"binlogPeriod"`
	// The standard starting number of days for log backup storage. The log backup will be converted when it reaches the standard starting number of days for storage. The minimum is 30 days and must not be greater than the number of days for log backup retention.
	BinlogStandbyDays *int `pulumi:"binlogStandbyDays"`
	// Whether to enable the log backup standard storage policy, `off` - close, `on` - open, the default is off.
	EnableBinlogStandby *string `pulumi:"enableBinlogStandby"`
	// Instance ID to which policies will be applied.
	MysqlId string `pulumi:"mysqlId"`
	// The retention time of backup files, in days. The minimum value is 7 days and the maximum value is 1830 days. And default value is `7`.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
}

// The set of arguments for constructing a BackupPolicy resource.
type BackupPolicyArgs struct {
	// Backup method. Supported values include: `physical` - physical backup.
	BackupModel pulumi.StringPtrInput
	// Instance backup time, in the format of 'HH:mm-HH:mm'. Time setting interval is four hours. Default to `02:00-06:00`. The following value can be supported: `02:00-06:00`, `06:00-10:00`, `10:00-14:00`, `14:00-18:00`, `18:00-22:00`, and `22:00-02:00`.
	BackupTime pulumi.StringPtrInput
	// Binlog retention time, in days. The minimum value is 7 days and the maximum value is 1830 days. This value cannot be set greater than the backup file retention time.
	BinlogPeriod pulumi.IntPtrInput
	// The standard starting number of days for log backup storage. The log backup will be converted when it reaches the standard starting number of days for storage. The minimum is 30 days and must not be greater than the number of days for log backup retention.
	BinlogStandbyDays pulumi.IntPtrInput
	// Whether to enable the log backup standard storage policy, `off` - close, `on` - open, the default is off.
	EnableBinlogStandby pulumi.StringPtrInput
	// Instance ID to which policies will be applied.
	MysqlId pulumi.StringInput
	// The retention time of backup files, in days. The minimum value is 7 days and the maximum value is 1830 days. And default value is `7`.
	RetentionPeriod pulumi.IntPtrInput
}

func (BackupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyArgs)(nil)).Elem()
}

type BackupPolicyInput interface {
	pulumi.Input

	ToBackupPolicyOutput() BackupPolicyOutput
	ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput
}

func (*BackupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicy)(nil)).Elem()
}

func (i *BackupPolicy) ToBackupPolicyOutput() BackupPolicyOutput {
	return i.ToBackupPolicyOutputWithContext(context.Background())
}

func (i *BackupPolicy) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyOutput)
}

// BackupPolicyArrayInput is an input type that accepts BackupPolicyArray and BackupPolicyArrayOutput values.
// You can construct a concrete instance of `BackupPolicyArrayInput` via:
//
//	BackupPolicyArray{ BackupPolicyArgs{...} }
type BackupPolicyArrayInput interface {
	pulumi.Input

	ToBackupPolicyArrayOutput() BackupPolicyArrayOutput
	ToBackupPolicyArrayOutputWithContext(context.Context) BackupPolicyArrayOutput
}

type BackupPolicyArray []BackupPolicyInput

func (BackupPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPolicy)(nil)).Elem()
}

func (i BackupPolicyArray) ToBackupPolicyArrayOutput() BackupPolicyArrayOutput {
	return i.ToBackupPolicyArrayOutputWithContext(context.Background())
}

func (i BackupPolicyArray) ToBackupPolicyArrayOutputWithContext(ctx context.Context) BackupPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyArrayOutput)
}

// BackupPolicyMapInput is an input type that accepts BackupPolicyMap and BackupPolicyMapOutput values.
// You can construct a concrete instance of `BackupPolicyMapInput` via:
//
//	BackupPolicyMap{ "key": BackupPolicyArgs{...} }
type BackupPolicyMapInput interface {
	pulumi.Input

	ToBackupPolicyMapOutput() BackupPolicyMapOutput
	ToBackupPolicyMapOutputWithContext(context.Context) BackupPolicyMapOutput
}

type BackupPolicyMap map[string]BackupPolicyInput

func (BackupPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPolicy)(nil)).Elem()
}

func (i BackupPolicyMap) ToBackupPolicyMapOutput() BackupPolicyMapOutput {
	return i.ToBackupPolicyMapOutputWithContext(context.Background())
}

func (i BackupPolicyMap) ToBackupPolicyMapOutputWithContext(ctx context.Context) BackupPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyMapOutput)
}

type BackupPolicyOutput struct{ *pulumi.OutputState }

func (BackupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicy)(nil)).Elem()
}

func (o BackupPolicyOutput) ToBackupPolicyOutput() BackupPolicyOutput {
	return o
}

func (o BackupPolicyOutput) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return o
}

// Backup method. Supported values include: `physical` - physical backup.
func (o BackupPolicyOutput) BackupModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringPtrOutput { return v.BackupModel }).(pulumi.StringPtrOutput)
}

// Instance backup time, in the format of 'HH:mm-HH:mm'. Time setting interval is four hours. Default to `02:00-06:00`. The following value can be supported: `02:00-06:00`, `06:00-10:00`, `10:00-14:00`, `14:00-18:00`, `18:00-22:00`, and `22:00-02:00`.
func (o BackupPolicyOutput) BackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringPtrOutput { return v.BackupTime }).(pulumi.StringPtrOutput)
}

// Binlog retention time, in days. The minimum value is 7 days and the maximum value is 1830 days. This value cannot be set greater than the backup file retention time.
func (o BackupPolicyOutput) BinlogPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.IntOutput { return v.BinlogPeriod }).(pulumi.IntOutput)
}

// The standard starting number of days for log backup storage. The log backup will be converted when it reaches the standard starting number of days for storage. The minimum is 30 days and must not be greater than the number of days for log backup retention.
func (o BackupPolicyOutput) BinlogStandbyDays() pulumi.IntOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.IntOutput { return v.BinlogStandbyDays }).(pulumi.IntOutput)
}

// Whether to enable the log backup standard storage policy, `off` - close, `on` - open, the default is off.
func (o BackupPolicyOutput) EnableBinlogStandby() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringPtrOutput { return v.EnableBinlogStandby }).(pulumi.StringPtrOutput)
}

// Instance ID to which policies will be applied.
func (o BackupPolicyOutput) MysqlId() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.MysqlId }).(pulumi.StringOutput)
}

// The retention time of backup files, in days. The minimum value is 7 days and the maximum value is 1830 days. And default value is `7`.
func (o BackupPolicyOutput) RetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.IntPtrOutput { return v.RetentionPeriod }).(pulumi.IntPtrOutput)
}

type BackupPolicyArrayOutput struct{ *pulumi.OutputState }

func (BackupPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPolicy)(nil)).Elem()
}

func (o BackupPolicyArrayOutput) ToBackupPolicyArrayOutput() BackupPolicyArrayOutput {
	return o
}

func (o BackupPolicyArrayOutput) ToBackupPolicyArrayOutputWithContext(ctx context.Context) BackupPolicyArrayOutput {
	return o
}

func (o BackupPolicyArrayOutput) Index(i pulumi.IntInput) BackupPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupPolicy {
		return vs[0].([]*BackupPolicy)[vs[1].(int)]
	}).(BackupPolicyOutput)
}

type BackupPolicyMapOutput struct{ *pulumi.OutputState }

func (BackupPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPolicy)(nil)).Elem()
}

func (o BackupPolicyMapOutput) ToBackupPolicyMapOutput() BackupPolicyMapOutput {
	return o
}

func (o BackupPolicyMapOutput) ToBackupPolicyMapOutputWithContext(ctx context.Context) BackupPolicyMapOutput {
	return o
}

func (o BackupPolicyMapOutput) MapIndex(k pulumi.StringInput) BackupPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupPolicy {
		return vs[0].(map[string]*BackupPolicy)[vs[1].(string)]
	}).(BackupPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyInput)(nil)).Elem(), &BackupPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyArrayInput)(nil)).Elem(), BackupPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyMapInput)(nil)).Elem(), BackupPolicyMap{})
	pulumi.RegisterOutputType(BackupPolicyOutput{})
	pulumi.RegisterOutputType(BackupPolicyArrayOutput{})
	pulumi.RegisterOutputType(BackupPolicyMapOutput{})
}
