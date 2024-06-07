// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a sqlserver migration
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Security"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			zones, err := Availability.GetZonesByProduct(ctx, &availability.GetZonesByProductArgs{
//				Product: "sqlserver",
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
//				AvailabilityZone: pulumi.String(zones.Zones[4].Name),
//				VpcId:            vpc.ID(),
//				CidrBlock:        pulumi.String("10.0.0.0/16"),
//				IsMulticast:      pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			securityGroup, err := Security.NewGroup(ctx, "securityGroup", &Security.GroupArgs{
//				Description: pulumi.String("desc."),
//			})
//			if err != nil {
//				return err
//			}
//			srcExample, err := Sqlserver.NewBasicInstance(ctx, "srcExample", &Sqlserver.BasicInstanceArgs{
//				AvailabilityZone: pulumi.String(zones.Zones[4].Name),
//				ChargeType:       pulumi.String("POSTPAID_BY_HOUR"),
//				VpcId:            vpc.ID(),
//				SubnetId:         subnet.ID(),
//				ProjectId:        pulumi.Int(0),
//				Memory:           pulumi.Int(4),
//				Storage:          pulumi.Int(100),
//				Cpu:              pulumi.Int(2),
//				MachineType:      pulumi.String("CLOUD_PREMIUM"),
//				MaintenanceWeekSets: pulumi.IntArray{
//					pulumi.Int(1),
//					pulumi.Int(2),
//					pulumi.Int(3),
//				},
//				MaintenanceStartTime: pulumi.String("09:00"),
//				MaintenanceTimeSpan:  pulumi.Int(3),
//				SecurityGroups: pulumi.StringArray{
//					securityGroup.ID(),
//				},
//				Tags: pulumi.Map{
//					"test": pulumi.Any("test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			dstExample, err := Sqlserver.NewBasicInstance(ctx, "dstExample", &Sqlserver.BasicInstanceArgs{
//				AvailabilityZone: pulumi.String(zones.Zones[4].Name),
//				ChargeType:       pulumi.String("POSTPAID_BY_HOUR"),
//				VpcId:            vpc.ID(),
//				SubnetId:         subnet.ID(),
//				ProjectId:        pulumi.Int(0),
//				Memory:           pulumi.Int(4),
//				Storage:          pulumi.Int(100),
//				Cpu:              pulumi.Int(2),
//				MachineType:      pulumi.String("CLOUD_PREMIUM"),
//				MaintenanceWeekSets: pulumi.IntArray{
//					pulumi.Int(1),
//					pulumi.Int(2),
//					pulumi.Int(3),
//				},
//				MaintenanceStartTime: pulumi.String("09:00"),
//				MaintenanceTimeSpan:  pulumi.Int(3),
//				SecurityGroups: pulumi.StringArray{
//					securityGroup.ID(),
//				},
//				Tags: pulumi.Map{
//					"test": pulumi.Any("test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			srcDb, err := Sqlserver.NewDb(ctx, "srcDb", &Sqlserver.DbArgs{
//				InstanceId: srcExample.ID(),
//				Charset:    pulumi.String("Chinese_PRC_BIN"),
//				Remark:     pulumi.String("testACC-remark"),
//			})
//			if err != nil {
//				return err
//			}
//			dstDb, err := Sqlserver.NewDb(ctx, "dstDb", &Sqlserver.DbArgs{
//				InstanceId: dstExample.ID(),
//				Charset:    pulumi.String("Chinese_PRC_BIN"),
//				Remark:     pulumi.String("testACC-remark"),
//			})
//			if err != nil {
//				return err
//			}
//			srcAccount, err := Sqlserver.NewAccount(ctx, "srcAccount", &Sqlserver.AccountArgs{
//				InstanceId: srcExample.ID(),
//				Password:   pulumi.String("Qwer@234"),
//				IsAdmin:    pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			dstAccount, err := Sqlserver.NewAccount(ctx, "dstAccount", &Sqlserver.AccountArgs{
//				InstanceId: dstExample.ID(),
//				Password:   pulumi.String("Qwer@234"),
//				IsAdmin:    pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Sqlserver.NewAccountDbAttachment(ctx, "srcAccountDbAttachment", &Sqlserver.AccountDbAttachmentArgs{
//				InstanceId:  srcExample.ID(),
//				AccountName: srcAccount.Name,
//				DbName:      srcDb.Name,
//				Privilege:   pulumi.String("ReadWrite"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Sqlserver.NewAccountDbAttachment(ctx, "dstAccountDbAttachment", &Sqlserver.AccountDbAttachmentArgs{
//				InstanceId:  dstExample.ID(),
//				AccountName: dstAccount.Name,
//				DbName:      dstDb.Name,
//				Privilege:   pulumi.String("ReadWrite"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Sqlserver.NewMigration(ctx, "migration", &Sqlserver.MigrationArgs{
//				MigrateName: pulumi.String("tf_test_migration"),
//				MigrateType: pulumi.Int(1),
//				SourceType:  pulumi.Int(1),
//				Source: &sqlserver.MigrationSourceArgs{
//					InstanceId: srcExample.ID(),
//					UserName:   srcAccount.Name,
//					Password:   srcAccount.Password,
//				},
//				Target: &sqlserver.MigrationTargetArgs{
//					InstanceId: dstExample.ID(),
//					UserName:   dstAccount.Name,
//					Password:   dstAccount.Password,
//				},
//				MigrateDbSets: sqlserver.MigrationMigrateDbSetArray{
//					&sqlserver.MigrationMigrateDbSetArgs{
//						DbName: srcDb.Name,
//					},
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
// sqlserver migration can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Sqlserver/migration:Migration migration migration_id
// ```
type Migration struct {
	pulumi.CustomResourceState

	// Migrate DB objects. Offline migration is not used (SourceType=4 or SourceType=5).
	MigrateDbSets MigrationMigrateDbSetArrayOutput `pulumi:"migrateDbSets"`
	// Name of the migration task.
	MigrateName pulumi.StringOutput `pulumi:"migrateName"`
	// Migration type (1 structure migration 2 data migration 3 incremental synchronization).
	MigrateType pulumi.IntOutput `pulumi:"migrateType"`
	// Restore and rename the database in ReNameRestoreDatabase. If it is not filled in, the restored database will be named by default and all databases will be restored. Valid if SourceType=5.
	RenameRestores MigrationRenameRestoreArrayOutput `pulumi:"renameRestores"`
	// Migration source.
	Source MigrationSourceOutput `pulumi:"source"`
	// Type of migration source 1 TencentDB for SQLServer 2 Cloud server self-built SQLServer database 4 SQLServer backup and restore 5 SQLServer backup and restore (COS mode).
	SourceType pulumi.IntOutput `pulumi:"sourceType"`
	// Migration target.
	Target MigrationTargetOutput `pulumi:"target"`
}

// NewMigration registers a new resource with the given unique name, arguments, and options.
func NewMigration(ctx *pulumi.Context,
	name string, args *MigrationArgs, opts ...pulumi.ResourceOption) (*Migration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MigrateName == nil {
		return nil, errors.New("invalid value for required argument 'MigrateName'")
	}
	if args.MigrateType == nil {
		return nil, errors.New("invalid value for required argument 'MigrateType'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	if args.SourceType == nil {
		return nil, errors.New("invalid value for required argument 'SourceType'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Migration
	err := ctx.RegisterResource("tencentcloud:Sqlserver/migration:Migration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMigration gets an existing Migration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMigration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MigrationState, opts ...pulumi.ResourceOption) (*Migration, error) {
	var resource Migration
	err := ctx.ReadResource("tencentcloud:Sqlserver/migration:Migration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Migration resources.
type migrationState struct {
	// Migrate DB objects. Offline migration is not used (SourceType=4 or SourceType=5).
	MigrateDbSets []MigrationMigrateDbSet `pulumi:"migrateDbSets"`
	// Name of the migration task.
	MigrateName *string `pulumi:"migrateName"`
	// Migration type (1 structure migration 2 data migration 3 incremental synchronization).
	MigrateType *int `pulumi:"migrateType"`
	// Restore and rename the database in ReNameRestoreDatabase. If it is not filled in, the restored database will be named by default and all databases will be restored. Valid if SourceType=5.
	RenameRestores []MigrationRenameRestore `pulumi:"renameRestores"`
	// Migration source.
	Source *MigrationSource `pulumi:"source"`
	// Type of migration source 1 TencentDB for SQLServer 2 Cloud server self-built SQLServer database 4 SQLServer backup and restore 5 SQLServer backup and restore (COS mode).
	SourceType *int `pulumi:"sourceType"`
	// Migration target.
	Target *MigrationTarget `pulumi:"target"`
}

type MigrationState struct {
	// Migrate DB objects. Offline migration is not used (SourceType=4 or SourceType=5).
	MigrateDbSets MigrationMigrateDbSetArrayInput
	// Name of the migration task.
	MigrateName pulumi.StringPtrInput
	// Migration type (1 structure migration 2 data migration 3 incremental synchronization).
	MigrateType pulumi.IntPtrInput
	// Restore and rename the database in ReNameRestoreDatabase. If it is not filled in, the restored database will be named by default and all databases will be restored. Valid if SourceType=5.
	RenameRestores MigrationRenameRestoreArrayInput
	// Migration source.
	Source MigrationSourcePtrInput
	// Type of migration source 1 TencentDB for SQLServer 2 Cloud server self-built SQLServer database 4 SQLServer backup and restore 5 SQLServer backup and restore (COS mode).
	SourceType pulumi.IntPtrInput
	// Migration target.
	Target MigrationTargetPtrInput
}

func (MigrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*migrationState)(nil)).Elem()
}

type migrationArgs struct {
	// Migrate DB objects. Offline migration is not used (SourceType=4 or SourceType=5).
	MigrateDbSets []MigrationMigrateDbSet `pulumi:"migrateDbSets"`
	// Name of the migration task.
	MigrateName string `pulumi:"migrateName"`
	// Migration type (1 structure migration 2 data migration 3 incremental synchronization).
	MigrateType int `pulumi:"migrateType"`
	// Restore and rename the database in ReNameRestoreDatabase. If it is not filled in, the restored database will be named by default and all databases will be restored. Valid if SourceType=5.
	RenameRestores []MigrationRenameRestore `pulumi:"renameRestores"`
	// Migration source.
	Source MigrationSource `pulumi:"source"`
	// Type of migration source 1 TencentDB for SQLServer 2 Cloud server self-built SQLServer database 4 SQLServer backup and restore 5 SQLServer backup and restore (COS mode).
	SourceType int `pulumi:"sourceType"`
	// Migration target.
	Target MigrationTarget `pulumi:"target"`
}

// The set of arguments for constructing a Migration resource.
type MigrationArgs struct {
	// Migrate DB objects. Offline migration is not used (SourceType=4 or SourceType=5).
	MigrateDbSets MigrationMigrateDbSetArrayInput
	// Name of the migration task.
	MigrateName pulumi.StringInput
	// Migration type (1 structure migration 2 data migration 3 incremental synchronization).
	MigrateType pulumi.IntInput
	// Restore and rename the database in ReNameRestoreDatabase. If it is not filled in, the restored database will be named by default and all databases will be restored. Valid if SourceType=5.
	RenameRestores MigrationRenameRestoreArrayInput
	// Migration source.
	Source MigrationSourceInput
	// Type of migration source 1 TencentDB for SQLServer 2 Cloud server self-built SQLServer database 4 SQLServer backup and restore 5 SQLServer backup and restore (COS mode).
	SourceType pulumi.IntInput
	// Migration target.
	Target MigrationTargetInput
}

func (MigrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*migrationArgs)(nil)).Elem()
}

type MigrationInput interface {
	pulumi.Input

	ToMigrationOutput() MigrationOutput
	ToMigrationOutputWithContext(ctx context.Context) MigrationOutput
}

func (*Migration) ElementType() reflect.Type {
	return reflect.TypeOf((**Migration)(nil)).Elem()
}

func (i *Migration) ToMigrationOutput() MigrationOutput {
	return i.ToMigrationOutputWithContext(context.Background())
}

func (i *Migration) ToMigrationOutputWithContext(ctx context.Context) MigrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationOutput)
}

// MigrationArrayInput is an input type that accepts MigrationArray and MigrationArrayOutput values.
// You can construct a concrete instance of `MigrationArrayInput` via:
//
//	MigrationArray{ MigrationArgs{...} }
type MigrationArrayInput interface {
	pulumi.Input

	ToMigrationArrayOutput() MigrationArrayOutput
	ToMigrationArrayOutputWithContext(context.Context) MigrationArrayOutput
}

type MigrationArray []MigrationInput

func (MigrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Migration)(nil)).Elem()
}

func (i MigrationArray) ToMigrationArrayOutput() MigrationArrayOutput {
	return i.ToMigrationArrayOutputWithContext(context.Background())
}

func (i MigrationArray) ToMigrationArrayOutputWithContext(ctx context.Context) MigrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationArrayOutput)
}

// MigrationMapInput is an input type that accepts MigrationMap and MigrationMapOutput values.
// You can construct a concrete instance of `MigrationMapInput` via:
//
//	MigrationMap{ "key": MigrationArgs{...} }
type MigrationMapInput interface {
	pulumi.Input

	ToMigrationMapOutput() MigrationMapOutput
	ToMigrationMapOutputWithContext(context.Context) MigrationMapOutput
}

type MigrationMap map[string]MigrationInput

func (MigrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Migration)(nil)).Elem()
}

func (i MigrationMap) ToMigrationMapOutput() MigrationMapOutput {
	return i.ToMigrationMapOutputWithContext(context.Background())
}

func (i MigrationMap) ToMigrationMapOutputWithContext(ctx context.Context) MigrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationMapOutput)
}

type MigrationOutput struct{ *pulumi.OutputState }

func (MigrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Migration)(nil)).Elem()
}

func (o MigrationOutput) ToMigrationOutput() MigrationOutput {
	return o
}

func (o MigrationOutput) ToMigrationOutputWithContext(ctx context.Context) MigrationOutput {
	return o
}

// Migrate DB objects. Offline migration is not used (SourceType=4 or SourceType=5).
func (o MigrationOutput) MigrateDbSets() MigrationMigrateDbSetArrayOutput {
	return o.ApplyT(func(v *Migration) MigrationMigrateDbSetArrayOutput { return v.MigrateDbSets }).(MigrationMigrateDbSetArrayOutput)
}

// Name of the migration task.
func (o MigrationOutput) MigrateName() pulumi.StringOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringOutput { return v.MigrateName }).(pulumi.StringOutput)
}

// Migration type (1 structure migration 2 data migration 3 incremental synchronization).
func (o MigrationOutput) MigrateType() pulumi.IntOutput {
	return o.ApplyT(func(v *Migration) pulumi.IntOutput { return v.MigrateType }).(pulumi.IntOutput)
}

// Restore and rename the database in ReNameRestoreDatabase. If it is not filled in, the restored database will be named by default and all databases will be restored. Valid if SourceType=5.
func (o MigrationOutput) RenameRestores() MigrationRenameRestoreArrayOutput {
	return o.ApplyT(func(v *Migration) MigrationRenameRestoreArrayOutput { return v.RenameRestores }).(MigrationRenameRestoreArrayOutput)
}

// Migration source.
func (o MigrationOutput) Source() MigrationSourceOutput {
	return o.ApplyT(func(v *Migration) MigrationSourceOutput { return v.Source }).(MigrationSourceOutput)
}

// Type of migration source 1 TencentDB for SQLServer 2 Cloud server self-built SQLServer database 4 SQLServer backup and restore 5 SQLServer backup and restore (COS mode).
func (o MigrationOutput) SourceType() pulumi.IntOutput {
	return o.ApplyT(func(v *Migration) pulumi.IntOutput { return v.SourceType }).(pulumi.IntOutput)
}

// Migration target.
func (o MigrationOutput) Target() MigrationTargetOutput {
	return o.ApplyT(func(v *Migration) MigrationTargetOutput { return v.Target }).(MigrationTargetOutput)
}

type MigrationArrayOutput struct{ *pulumi.OutputState }

func (MigrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Migration)(nil)).Elem()
}

func (o MigrationArrayOutput) ToMigrationArrayOutput() MigrationArrayOutput {
	return o
}

func (o MigrationArrayOutput) ToMigrationArrayOutputWithContext(ctx context.Context) MigrationArrayOutput {
	return o
}

func (o MigrationArrayOutput) Index(i pulumi.IntInput) MigrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Migration {
		return vs[0].([]*Migration)[vs[1].(int)]
	}).(MigrationOutput)
}

type MigrationMapOutput struct{ *pulumi.OutputState }

func (MigrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Migration)(nil)).Elem()
}

func (o MigrationMapOutput) ToMigrationMapOutput() MigrationMapOutput {
	return o
}

func (o MigrationMapOutput) ToMigrationMapOutputWithContext(ctx context.Context) MigrationMapOutput {
	return o
}

func (o MigrationMapOutput) MapIndex(k pulumi.StringInput) MigrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Migration {
		return vs[0].(map[string]*Migration)[vs[1].(string)]
	}).(MigrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MigrationInput)(nil)).Elem(), &Migration{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrationArrayInput)(nil)).Elem(), MigrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrationMapInput)(nil)).Elem(), MigrationMap{})
	pulumi.RegisterOutputType(MigrationOutput{})
	pulumi.RegisterOutputType(MigrationArrayOutput{})
	pulumi.RegisterOutputType(MigrationMapOutput{})
}
