// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mariadb backupTime
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Mariadb.NewBackupTime(ctx, "backupTime", &Mariadb.BackupTimeArgs{
// 			EndBackupTime:   pulumi.String("04:00"),
// 			InstanceId:      pulumi.String("tdsql-9vqvls95"),
// 			StartBackupTime: pulumi.String("01:00"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// mariadb backup_time can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Mariadb/backupTime:BackupTime backup_time backup_time_id
// ```
type BackupTime struct {
	pulumi.CustomResourceState

	// End time of daily backup window in the format of `mm:ss`, such as 23:59.
	EndBackupTime pulumi.StringOutput `pulumi:"endBackupTime"`
	// instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Start time of daily backup window in the format of `mm:ss`, such as 22:00.
	StartBackupTime pulumi.StringOutput `pulumi:"startBackupTime"`
}

// NewBackupTime registers a new resource with the given unique name, arguments, and options.
func NewBackupTime(ctx *pulumi.Context,
	name string, args *BackupTimeArgs, opts ...pulumi.ResourceOption) (*BackupTime, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndBackupTime == nil {
		return nil, errors.New("invalid value for required argument 'EndBackupTime'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.StartBackupTime == nil {
		return nil, errors.New("invalid value for required argument 'StartBackupTime'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BackupTime
	err := ctx.RegisterResource("tencentcloud:Mariadb/backupTime:BackupTime", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupTime gets an existing BackupTime resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupTime(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupTimeState, opts ...pulumi.ResourceOption) (*BackupTime, error) {
	var resource BackupTime
	err := ctx.ReadResource("tencentcloud:Mariadb/backupTime:BackupTime", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupTime resources.
type backupTimeState struct {
	// End time of daily backup window in the format of `mm:ss`, such as 23:59.
	EndBackupTime *string `pulumi:"endBackupTime"`
	// instance id.
	InstanceId *string `pulumi:"instanceId"`
	// Start time of daily backup window in the format of `mm:ss`, such as 22:00.
	StartBackupTime *string `pulumi:"startBackupTime"`
}

type BackupTimeState struct {
	// End time of daily backup window in the format of `mm:ss`, such as 23:59.
	EndBackupTime pulumi.StringPtrInput
	// instance id.
	InstanceId pulumi.StringPtrInput
	// Start time of daily backup window in the format of `mm:ss`, such as 22:00.
	StartBackupTime pulumi.StringPtrInput
}

func (BackupTimeState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupTimeState)(nil)).Elem()
}

type backupTimeArgs struct {
	// End time of daily backup window in the format of `mm:ss`, such as 23:59.
	EndBackupTime string `pulumi:"endBackupTime"`
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// Start time of daily backup window in the format of `mm:ss`, such as 22:00.
	StartBackupTime string `pulumi:"startBackupTime"`
}

// The set of arguments for constructing a BackupTime resource.
type BackupTimeArgs struct {
	// End time of daily backup window in the format of `mm:ss`, such as 23:59.
	EndBackupTime pulumi.StringInput
	// instance id.
	InstanceId pulumi.StringInput
	// Start time of daily backup window in the format of `mm:ss`, such as 22:00.
	StartBackupTime pulumi.StringInput
}

func (BackupTimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupTimeArgs)(nil)).Elem()
}

type BackupTimeInput interface {
	pulumi.Input

	ToBackupTimeOutput() BackupTimeOutput
	ToBackupTimeOutputWithContext(ctx context.Context) BackupTimeOutput
}

func (*BackupTime) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupTime)(nil)).Elem()
}

func (i *BackupTime) ToBackupTimeOutput() BackupTimeOutput {
	return i.ToBackupTimeOutputWithContext(context.Background())
}

func (i *BackupTime) ToBackupTimeOutputWithContext(ctx context.Context) BackupTimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupTimeOutput)
}

// BackupTimeArrayInput is an input type that accepts BackupTimeArray and BackupTimeArrayOutput values.
// You can construct a concrete instance of `BackupTimeArrayInput` via:
//
//          BackupTimeArray{ BackupTimeArgs{...} }
type BackupTimeArrayInput interface {
	pulumi.Input

	ToBackupTimeArrayOutput() BackupTimeArrayOutput
	ToBackupTimeArrayOutputWithContext(context.Context) BackupTimeArrayOutput
}

type BackupTimeArray []BackupTimeInput

func (BackupTimeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupTime)(nil)).Elem()
}

func (i BackupTimeArray) ToBackupTimeArrayOutput() BackupTimeArrayOutput {
	return i.ToBackupTimeArrayOutputWithContext(context.Background())
}

func (i BackupTimeArray) ToBackupTimeArrayOutputWithContext(ctx context.Context) BackupTimeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupTimeArrayOutput)
}

// BackupTimeMapInput is an input type that accepts BackupTimeMap and BackupTimeMapOutput values.
// You can construct a concrete instance of `BackupTimeMapInput` via:
//
//          BackupTimeMap{ "key": BackupTimeArgs{...} }
type BackupTimeMapInput interface {
	pulumi.Input

	ToBackupTimeMapOutput() BackupTimeMapOutput
	ToBackupTimeMapOutputWithContext(context.Context) BackupTimeMapOutput
}

type BackupTimeMap map[string]BackupTimeInput

func (BackupTimeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupTime)(nil)).Elem()
}

func (i BackupTimeMap) ToBackupTimeMapOutput() BackupTimeMapOutput {
	return i.ToBackupTimeMapOutputWithContext(context.Background())
}

func (i BackupTimeMap) ToBackupTimeMapOutputWithContext(ctx context.Context) BackupTimeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupTimeMapOutput)
}

type BackupTimeOutput struct{ *pulumi.OutputState }

func (BackupTimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupTime)(nil)).Elem()
}

func (o BackupTimeOutput) ToBackupTimeOutput() BackupTimeOutput {
	return o
}

func (o BackupTimeOutput) ToBackupTimeOutputWithContext(ctx context.Context) BackupTimeOutput {
	return o
}

// End time of daily backup window in the format of `mm:ss`, such as 23:59.
func (o BackupTimeOutput) EndBackupTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupTime) pulumi.StringOutput { return v.EndBackupTime }).(pulumi.StringOutput)
}

// instance id.
func (o BackupTimeOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupTime) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Start time of daily backup window in the format of `mm:ss`, such as 22:00.
func (o BackupTimeOutput) StartBackupTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupTime) pulumi.StringOutput { return v.StartBackupTime }).(pulumi.StringOutput)
}

type BackupTimeArrayOutput struct{ *pulumi.OutputState }

func (BackupTimeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupTime)(nil)).Elem()
}

func (o BackupTimeArrayOutput) ToBackupTimeArrayOutput() BackupTimeArrayOutput {
	return o
}

func (o BackupTimeArrayOutput) ToBackupTimeArrayOutputWithContext(ctx context.Context) BackupTimeArrayOutput {
	return o
}

func (o BackupTimeArrayOutput) Index(i pulumi.IntInput) BackupTimeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupTime {
		return vs[0].([]*BackupTime)[vs[1].(int)]
	}).(BackupTimeOutput)
}

type BackupTimeMapOutput struct{ *pulumi.OutputState }

func (BackupTimeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupTime)(nil)).Elem()
}

func (o BackupTimeMapOutput) ToBackupTimeMapOutput() BackupTimeMapOutput {
	return o
}

func (o BackupTimeMapOutput) ToBackupTimeMapOutputWithContext(ctx context.Context) BackupTimeMapOutput {
	return o
}

func (o BackupTimeMapOutput) MapIndex(k pulumi.StringInput) BackupTimeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupTime {
		return vs[0].(map[string]*BackupTime)[vs[1].(string)]
	}).(BackupTimeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupTimeInput)(nil)).Elem(), &BackupTime{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupTimeArrayInput)(nil)).Elem(), BackupTimeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupTimeMapInput)(nil)).Elem(), BackupTimeMap{})
	pulumi.RegisterOutputType(BackupTimeOutput{})
	pulumi.RegisterOutputType(BackupTimeArrayOutput{})
	pulumi.RegisterOutputType(BackupTimeMapOutput{})
}
