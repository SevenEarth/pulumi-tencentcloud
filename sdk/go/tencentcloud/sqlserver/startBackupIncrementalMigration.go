// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a sqlserver startBackupIncrementalMigration
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Sqlserver.NewStartBackupIncrementalMigration(ctx, "startBackupIncrementalMigration", &Sqlserver.StartBackupIncrementalMigrationArgs{
//				BackupMigrationId:      pulumi.String("mssql-backup-migration-cg0ffgqt"),
//				IncrementalMigrationId: pulumi.String("mssql-incremental-migration-kp7bgv8p"),
//				InstanceId:             pulumi.String("mssql-i1z41iwd"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type StartBackupIncrementalMigration struct {
	pulumi.CustomResourceState

	// Backup import task ID, returned by the CreateBackupMigration interface.
	BackupMigrationId pulumi.StringOutput `pulumi:"backupMigrationId"`
	// Incremental backup import task ID.
	IncrementalMigrationId pulumi.StringOutput `pulumi:"incrementalMigrationId"`
	// ID of imported target instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewStartBackupIncrementalMigration registers a new resource with the given unique name, arguments, and options.
func NewStartBackupIncrementalMigration(ctx *pulumi.Context,
	name string, args *StartBackupIncrementalMigrationArgs, opts ...pulumi.ResourceOption) (*StartBackupIncrementalMigration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupMigrationId == nil {
		return nil, errors.New("invalid value for required argument 'BackupMigrationId'")
	}
	if args.IncrementalMigrationId == nil {
		return nil, errors.New("invalid value for required argument 'IncrementalMigrationId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource StartBackupIncrementalMigration
	err := ctx.RegisterResource("tencentcloud:Sqlserver/startBackupIncrementalMigration:StartBackupIncrementalMigration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStartBackupIncrementalMigration gets an existing StartBackupIncrementalMigration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStartBackupIncrementalMigration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StartBackupIncrementalMigrationState, opts ...pulumi.ResourceOption) (*StartBackupIncrementalMigration, error) {
	var resource StartBackupIncrementalMigration
	err := ctx.ReadResource("tencentcloud:Sqlserver/startBackupIncrementalMigration:StartBackupIncrementalMigration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StartBackupIncrementalMigration resources.
type startBackupIncrementalMigrationState struct {
	// Backup import task ID, returned by the CreateBackupMigration interface.
	BackupMigrationId *string `pulumi:"backupMigrationId"`
	// Incremental backup import task ID.
	IncrementalMigrationId *string `pulumi:"incrementalMigrationId"`
	// ID of imported target instance.
	InstanceId *string `pulumi:"instanceId"`
}

type StartBackupIncrementalMigrationState struct {
	// Backup import task ID, returned by the CreateBackupMigration interface.
	BackupMigrationId pulumi.StringPtrInput
	// Incremental backup import task ID.
	IncrementalMigrationId pulumi.StringPtrInput
	// ID of imported target instance.
	InstanceId pulumi.StringPtrInput
}

func (StartBackupIncrementalMigrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*startBackupIncrementalMigrationState)(nil)).Elem()
}

type startBackupIncrementalMigrationArgs struct {
	// Backup import task ID, returned by the CreateBackupMigration interface.
	BackupMigrationId string `pulumi:"backupMigrationId"`
	// Incremental backup import task ID.
	IncrementalMigrationId string `pulumi:"incrementalMigrationId"`
	// ID of imported target instance.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a StartBackupIncrementalMigration resource.
type StartBackupIncrementalMigrationArgs struct {
	// Backup import task ID, returned by the CreateBackupMigration interface.
	BackupMigrationId pulumi.StringInput
	// Incremental backup import task ID.
	IncrementalMigrationId pulumi.StringInput
	// ID of imported target instance.
	InstanceId pulumi.StringInput
}

func (StartBackupIncrementalMigrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*startBackupIncrementalMigrationArgs)(nil)).Elem()
}

type StartBackupIncrementalMigrationInput interface {
	pulumi.Input

	ToStartBackupIncrementalMigrationOutput() StartBackupIncrementalMigrationOutput
	ToStartBackupIncrementalMigrationOutputWithContext(ctx context.Context) StartBackupIncrementalMigrationOutput
}

func (*StartBackupIncrementalMigration) ElementType() reflect.Type {
	return reflect.TypeOf((**StartBackupIncrementalMigration)(nil)).Elem()
}

func (i *StartBackupIncrementalMigration) ToStartBackupIncrementalMigrationOutput() StartBackupIncrementalMigrationOutput {
	return i.ToStartBackupIncrementalMigrationOutputWithContext(context.Background())
}

func (i *StartBackupIncrementalMigration) ToStartBackupIncrementalMigrationOutputWithContext(ctx context.Context) StartBackupIncrementalMigrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartBackupIncrementalMigrationOutput)
}

// StartBackupIncrementalMigrationArrayInput is an input type that accepts StartBackupIncrementalMigrationArray and StartBackupIncrementalMigrationArrayOutput values.
// You can construct a concrete instance of `StartBackupIncrementalMigrationArrayInput` via:
//
//	StartBackupIncrementalMigrationArray{ StartBackupIncrementalMigrationArgs{...} }
type StartBackupIncrementalMigrationArrayInput interface {
	pulumi.Input

	ToStartBackupIncrementalMigrationArrayOutput() StartBackupIncrementalMigrationArrayOutput
	ToStartBackupIncrementalMigrationArrayOutputWithContext(context.Context) StartBackupIncrementalMigrationArrayOutput
}

type StartBackupIncrementalMigrationArray []StartBackupIncrementalMigrationInput

func (StartBackupIncrementalMigrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StartBackupIncrementalMigration)(nil)).Elem()
}

func (i StartBackupIncrementalMigrationArray) ToStartBackupIncrementalMigrationArrayOutput() StartBackupIncrementalMigrationArrayOutput {
	return i.ToStartBackupIncrementalMigrationArrayOutputWithContext(context.Background())
}

func (i StartBackupIncrementalMigrationArray) ToStartBackupIncrementalMigrationArrayOutputWithContext(ctx context.Context) StartBackupIncrementalMigrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartBackupIncrementalMigrationArrayOutput)
}

// StartBackupIncrementalMigrationMapInput is an input type that accepts StartBackupIncrementalMigrationMap and StartBackupIncrementalMigrationMapOutput values.
// You can construct a concrete instance of `StartBackupIncrementalMigrationMapInput` via:
//
//	StartBackupIncrementalMigrationMap{ "key": StartBackupIncrementalMigrationArgs{...} }
type StartBackupIncrementalMigrationMapInput interface {
	pulumi.Input

	ToStartBackupIncrementalMigrationMapOutput() StartBackupIncrementalMigrationMapOutput
	ToStartBackupIncrementalMigrationMapOutputWithContext(context.Context) StartBackupIncrementalMigrationMapOutput
}

type StartBackupIncrementalMigrationMap map[string]StartBackupIncrementalMigrationInput

func (StartBackupIncrementalMigrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StartBackupIncrementalMigration)(nil)).Elem()
}

func (i StartBackupIncrementalMigrationMap) ToStartBackupIncrementalMigrationMapOutput() StartBackupIncrementalMigrationMapOutput {
	return i.ToStartBackupIncrementalMigrationMapOutputWithContext(context.Background())
}

func (i StartBackupIncrementalMigrationMap) ToStartBackupIncrementalMigrationMapOutputWithContext(ctx context.Context) StartBackupIncrementalMigrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartBackupIncrementalMigrationMapOutput)
}

type StartBackupIncrementalMigrationOutput struct{ *pulumi.OutputState }

func (StartBackupIncrementalMigrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StartBackupIncrementalMigration)(nil)).Elem()
}

func (o StartBackupIncrementalMigrationOutput) ToStartBackupIncrementalMigrationOutput() StartBackupIncrementalMigrationOutput {
	return o
}

func (o StartBackupIncrementalMigrationOutput) ToStartBackupIncrementalMigrationOutputWithContext(ctx context.Context) StartBackupIncrementalMigrationOutput {
	return o
}

// Backup import task ID, returned by the CreateBackupMigration interface.
func (o StartBackupIncrementalMigrationOutput) BackupMigrationId() pulumi.StringOutput {
	return o.ApplyT(func(v *StartBackupIncrementalMigration) pulumi.StringOutput { return v.BackupMigrationId }).(pulumi.StringOutput)
}

// Incremental backup import task ID.
func (o StartBackupIncrementalMigrationOutput) IncrementalMigrationId() pulumi.StringOutput {
	return o.ApplyT(func(v *StartBackupIncrementalMigration) pulumi.StringOutput { return v.IncrementalMigrationId }).(pulumi.StringOutput)
}

// ID of imported target instance.
func (o StartBackupIncrementalMigrationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *StartBackupIncrementalMigration) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type StartBackupIncrementalMigrationArrayOutput struct{ *pulumi.OutputState }

func (StartBackupIncrementalMigrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StartBackupIncrementalMigration)(nil)).Elem()
}

func (o StartBackupIncrementalMigrationArrayOutput) ToStartBackupIncrementalMigrationArrayOutput() StartBackupIncrementalMigrationArrayOutput {
	return o
}

func (o StartBackupIncrementalMigrationArrayOutput) ToStartBackupIncrementalMigrationArrayOutputWithContext(ctx context.Context) StartBackupIncrementalMigrationArrayOutput {
	return o
}

func (o StartBackupIncrementalMigrationArrayOutput) Index(i pulumi.IntInput) StartBackupIncrementalMigrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StartBackupIncrementalMigration {
		return vs[0].([]*StartBackupIncrementalMigration)[vs[1].(int)]
	}).(StartBackupIncrementalMigrationOutput)
}

type StartBackupIncrementalMigrationMapOutput struct{ *pulumi.OutputState }

func (StartBackupIncrementalMigrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StartBackupIncrementalMigration)(nil)).Elem()
}

func (o StartBackupIncrementalMigrationMapOutput) ToStartBackupIncrementalMigrationMapOutput() StartBackupIncrementalMigrationMapOutput {
	return o
}

func (o StartBackupIncrementalMigrationMapOutput) ToStartBackupIncrementalMigrationMapOutputWithContext(ctx context.Context) StartBackupIncrementalMigrationMapOutput {
	return o
}

func (o StartBackupIncrementalMigrationMapOutput) MapIndex(k pulumi.StringInput) StartBackupIncrementalMigrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StartBackupIncrementalMigration {
		return vs[0].(map[string]*StartBackupIncrementalMigration)[vs[1].(string)]
	}).(StartBackupIncrementalMigrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StartBackupIncrementalMigrationInput)(nil)).Elem(), &StartBackupIncrementalMigration{})
	pulumi.RegisterInputType(reflect.TypeOf((*StartBackupIncrementalMigrationArrayInput)(nil)).Elem(), StartBackupIncrementalMigrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StartBackupIncrementalMigrationMapInput)(nil)).Elem(), StartBackupIncrementalMigrationMap{})
	pulumi.RegisterOutputType(StartBackupIncrementalMigrationOutput{})
	pulumi.RegisterOutputType(StartBackupIncrementalMigrationArrayOutput{})
	pulumi.RegisterOutputType(StartBackupIncrementalMigrationMapOutput{})
}
