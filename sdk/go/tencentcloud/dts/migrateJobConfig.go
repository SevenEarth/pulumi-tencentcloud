// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dts

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a dts migrateJobConfig
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dts"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dts"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		service, err := Dts.NewMigrateService(ctx, "service", &Dts.MigrateServiceArgs{
// 			SrcDatabaseType: pulumi.String("mysql"),
// 			DstDatabaseType: pulumi.String("cynosdbmysql"),
// 			SrcRegion:       pulumi.String("ap-guangzhou"),
// 			DstRegion:       pulumi.String("ap-guangzhou"),
// 			InstanceClass:   pulumi.String("small"),
// 			JobName:         pulumi.String("tf_test_xxx"),
// 			Tags: dts.MigrateServiceTagArray{
// 				&dts.MigrateServiceTagArgs{
// 					TagKey:   pulumi.String("aaa"),
// 					TagValue: pulumi.String("bbb"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		job, err := Dts.NewMigrateJob(ctx, "job", &Dts.MigrateJobArgs{
// 			ServiceId: service.ID(),
// 			RunMode:   pulumi.String("immediate"),
// 			MigrateOption: &dts.MigrateJobMigrateOptionArgs{
// 				DatabaseTable: &dts.MigrateJobMigrateOptionDatabaseTableArgs{
// 					ObjectMode: pulumi.String("partial"),
// 					Databases: dts.MigrateJobMigrateOptionDatabaseTableDatabaseArray{
// 						&dts.MigrateJobMigrateOptionDatabaseTableDatabaseArgs{
// 							DbName:    pulumi.String("tf_ci_test"),
// 							DbMode:    pulumi.String("partial"),
// 							TableMode: pulumi.String("partial"),
// 							Tables: dts.MigrateJobMigrateOptionDatabaseTableDatabaseTableArray{
// 								&dts.MigrateJobMigrateOptionDatabaseTableDatabaseTableArgs{
// 									TableName:     pulumi.String("test"),
// 									NewTableName:  pulumi.String("test_xxx"),
// 									TableEditMode: pulumi.String("rename"),
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			SrcInfo: &dts.MigrateJobSrcInfoArgs{
// 				Region:       pulumi.String("ap-guangzhou"),
// 				AccessType:   pulumi.String("cdb"),
// 				DatabaseType: pulumi.String("mysql"),
// 				NodeType:     pulumi.String("simple"),
// 				Infos: dts.MigrateJobSrcInfoInfoArray{
// 					&dts.MigrateJobSrcInfoInfoArgs{
// 						User:       pulumi.String("root"),
// 						Password:   pulumi.String("xxx"),
// 						InstanceId: pulumi.String("id"),
// 					},
// 				},
// 			},
// 			DstInfo: &dts.MigrateJobDstInfoArgs{
// 				Region:       pulumi.String("ap-guangzhou"),
// 				AccessType:   pulumi.String("cdb"),
// 				DatabaseType: pulumi.String("cynosdbmysql"),
// 				NodeType:     pulumi.String("simple"),
// 				Infos: dts.MigrateJobDstInfoInfoArray{
// 					&dts.MigrateJobDstInfoInfoArgs{
// 						User:       pulumi.String("user"),
// 						Password:   pulumi.String("xxx"),
// 						InstanceId: pulumi.String("id"),
// 					},
// 				},
// 			},
// 			AutoRetryTimeRangeMinutes: pulumi.Int(0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		start, err := Dts.NewMigrateJobStartOperation(ctx, "start", &Dts.MigrateJobStartOperationArgs{
// 			JobId: job.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Dts.NewMigrateJobConfig(ctx, "config", &Dts.MigrateJobConfigArgs{
// 			JobId:  start.ID(),
// 			Action: pulumi.String("pause"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Continue the a migration job
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dts"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dts.NewMigrateJobConfig(ctx, "config", &Dts.MigrateJobConfigArgs{
// 			JobId:  pulumi.Any(tencentcloud_dts_migrate_job_start_operation.Start.Id),
// 			Action: pulumi.String("continue"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Complete a migration job when the status is readyComplete
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dts"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dts.NewMigrateJobConfig(ctx, "config", &Dts.MigrateJobConfigArgs{
// 			JobId:  pulumi.Any(tencentcloud_dts_migrate_job_start_operation.Start.Id),
// 			Action: pulumi.String("continue"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Stop a running migration job
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dts"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dts.NewMigrateJobConfig(ctx, "config", &Dts.MigrateJobConfigArgs{
// 			JobId:  pulumi.Any(tencentcloud_dts_migrate_job_start_operation.Start.Id),
// 			Action: pulumi.String("stop"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Isolate a stopped/canceled migration job
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dts"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dts.NewMigrateJobConfig(ctx, "config", &Dts.MigrateJobConfigArgs{
// 			JobId:  pulumi.Any(tencentcloud_dts_migrate_job_start_operation.Start.Id),
// 			Action: pulumi.String("isolate"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Recover a isolated migration job
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dts"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dts.NewMigrateJobConfig(ctx, "config", &Dts.MigrateJobConfigArgs{
// 			JobId:  pulumi.Any(tencentcloud_dts_migrate_job_start_operation.Start.Id),
// 			Action: pulumi.String("recover"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type MigrateJobConfig struct {
	pulumi.CustomResourceState

	// The operation want to perform. Valid values are: `pause`, `continue`, `complete`, `recover`,`stop`.
	Action pulumi.StringOutput `pulumi:"action"`
	// complete mode, optional value is waitForSync or immediately.
	CompleteMode pulumi.StringPtrOutput `pulumi:"completeMode"`
	// job id.
	JobId pulumi.StringOutput `pulumi:"jobId"`
}

// NewMigrateJobConfig registers a new resource with the given unique name, arguments, and options.
func NewMigrateJobConfig(ctx *pulumi.Context,
	name string, args *MigrateJobConfigArgs, opts ...pulumi.ResourceOption) (*MigrateJobConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.JobId == nil {
		return nil, errors.New("invalid value for required argument 'JobId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MigrateJobConfig
	err := ctx.RegisterResource("tencentcloud:Dts/migrateJobConfig:MigrateJobConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMigrateJobConfig gets an existing MigrateJobConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMigrateJobConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MigrateJobConfigState, opts ...pulumi.ResourceOption) (*MigrateJobConfig, error) {
	var resource MigrateJobConfig
	err := ctx.ReadResource("tencentcloud:Dts/migrateJobConfig:MigrateJobConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MigrateJobConfig resources.
type migrateJobConfigState struct {
	// The operation want to perform. Valid values are: `pause`, `continue`, `complete`, `recover`,`stop`.
	Action *string `pulumi:"action"`
	// complete mode, optional value is waitForSync or immediately.
	CompleteMode *string `pulumi:"completeMode"`
	// job id.
	JobId *string `pulumi:"jobId"`
}

type MigrateJobConfigState struct {
	// The operation want to perform. Valid values are: `pause`, `continue`, `complete`, `recover`,`stop`.
	Action pulumi.StringPtrInput
	// complete mode, optional value is waitForSync or immediately.
	CompleteMode pulumi.StringPtrInput
	// job id.
	JobId pulumi.StringPtrInput
}

func (MigrateJobConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*migrateJobConfigState)(nil)).Elem()
}

type migrateJobConfigArgs struct {
	// The operation want to perform. Valid values are: `pause`, `continue`, `complete`, `recover`,`stop`.
	Action string `pulumi:"action"`
	// complete mode, optional value is waitForSync or immediately.
	CompleteMode *string `pulumi:"completeMode"`
	// job id.
	JobId string `pulumi:"jobId"`
}

// The set of arguments for constructing a MigrateJobConfig resource.
type MigrateJobConfigArgs struct {
	// The operation want to perform. Valid values are: `pause`, `continue`, `complete`, `recover`,`stop`.
	Action pulumi.StringInput
	// complete mode, optional value is waitForSync or immediately.
	CompleteMode pulumi.StringPtrInput
	// job id.
	JobId pulumi.StringInput
}

func (MigrateJobConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*migrateJobConfigArgs)(nil)).Elem()
}

type MigrateJobConfigInput interface {
	pulumi.Input

	ToMigrateJobConfigOutput() MigrateJobConfigOutput
	ToMigrateJobConfigOutputWithContext(ctx context.Context) MigrateJobConfigOutput
}

func (*MigrateJobConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateJobConfig)(nil)).Elem()
}

func (i *MigrateJobConfig) ToMigrateJobConfigOutput() MigrateJobConfigOutput {
	return i.ToMigrateJobConfigOutputWithContext(context.Background())
}

func (i *MigrateJobConfig) ToMigrateJobConfigOutputWithContext(ctx context.Context) MigrateJobConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateJobConfigOutput)
}

// MigrateJobConfigArrayInput is an input type that accepts MigrateJobConfigArray and MigrateJobConfigArrayOutput values.
// You can construct a concrete instance of `MigrateJobConfigArrayInput` via:
//
//          MigrateJobConfigArray{ MigrateJobConfigArgs{...} }
type MigrateJobConfigArrayInput interface {
	pulumi.Input

	ToMigrateJobConfigArrayOutput() MigrateJobConfigArrayOutput
	ToMigrateJobConfigArrayOutputWithContext(context.Context) MigrateJobConfigArrayOutput
}

type MigrateJobConfigArray []MigrateJobConfigInput

func (MigrateJobConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MigrateJobConfig)(nil)).Elem()
}

func (i MigrateJobConfigArray) ToMigrateJobConfigArrayOutput() MigrateJobConfigArrayOutput {
	return i.ToMigrateJobConfigArrayOutputWithContext(context.Background())
}

func (i MigrateJobConfigArray) ToMigrateJobConfigArrayOutputWithContext(ctx context.Context) MigrateJobConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateJobConfigArrayOutput)
}

// MigrateJobConfigMapInput is an input type that accepts MigrateJobConfigMap and MigrateJobConfigMapOutput values.
// You can construct a concrete instance of `MigrateJobConfigMapInput` via:
//
//          MigrateJobConfigMap{ "key": MigrateJobConfigArgs{...} }
type MigrateJobConfigMapInput interface {
	pulumi.Input

	ToMigrateJobConfigMapOutput() MigrateJobConfigMapOutput
	ToMigrateJobConfigMapOutputWithContext(context.Context) MigrateJobConfigMapOutput
}

type MigrateJobConfigMap map[string]MigrateJobConfigInput

func (MigrateJobConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MigrateJobConfig)(nil)).Elem()
}

func (i MigrateJobConfigMap) ToMigrateJobConfigMapOutput() MigrateJobConfigMapOutput {
	return i.ToMigrateJobConfigMapOutputWithContext(context.Background())
}

func (i MigrateJobConfigMap) ToMigrateJobConfigMapOutputWithContext(ctx context.Context) MigrateJobConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateJobConfigMapOutput)
}

type MigrateJobConfigOutput struct{ *pulumi.OutputState }

func (MigrateJobConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateJobConfig)(nil)).Elem()
}

func (o MigrateJobConfigOutput) ToMigrateJobConfigOutput() MigrateJobConfigOutput {
	return o
}

func (o MigrateJobConfigOutput) ToMigrateJobConfigOutputWithContext(ctx context.Context) MigrateJobConfigOutput {
	return o
}

// The operation want to perform. Valid values are: `pause`, `continue`, `complete`, `recover`,`stop`.
func (o MigrateJobConfigOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *MigrateJobConfig) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// complete mode, optional value is waitForSync or immediately.
func (o MigrateJobConfigOutput) CompleteMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateJobConfig) pulumi.StringPtrOutput { return v.CompleteMode }).(pulumi.StringPtrOutput)
}

// job id.
func (o MigrateJobConfigOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v *MigrateJobConfig) pulumi.StringOutput { return v.JobId }).(pulumi.StringOutput)
}

type MigrateJobConfigArrayOutput struct{ *pulumi.OutputState }

func (MigrateJobConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MigrateJobConfig)(nil)).Elem()
}

func (o MigrateJobConfigArrayOutput) ToMigrateJobConfigArrayOutput() MigrateJobConfigArrayOutput {
	return o
}

func (o MigrateJobConfigArrayOutput) ToMigrateJobConfigArrayOutputWithContext(ctx context.Context) MigrateJobConfigArrayOutput {
	return o
}

func (o MigrateJobConfigArrayOutput) Index(i pulumi.IntInput) MigrateJobConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MigrateJobConfig {
		return vs[0].([]*MigrateJobConfig)[vs[1].(int)]
	}).(MigrateJobConfigOutput)
}

type MigrateJobConfigMapOutput struct{ *pulumi.OutputState }

func (MigrateJobConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MigrateJobConfig)(nil)).Elem()
}

func (o MigrateJobConfigMapOutput) ToMigrateJobConfigMapOutput() MigrateJobConfigMapOutput {
	return o
}

func (o MigrateJobConfigMapOutput) ToMigrateJobConfigMapOutputWithContext(ctx context.Context) MigrateJobConfigMapOutput {
	return o
}

func (o MigrateJobConfigMapOutput) MapIndex(k pulumi.StringInput) MigrateJobConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MigrateJobConfig {
		return vs[0].(map[string]*MigrateJobConfig)[vs[1].(string)]
	}).(MigrateJobConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MigrateJobConfigInput)(nil)).Elem(), &MigrateJobConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrateJobConfigArrayInput)(nil)).Elem(), MigrateJobConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrateJobConfigMapInput)(nil)).Elem(), MigrateJobConfigMap{})
	pulumi.RegisterOutputType(MigrateJobConfigOutput{})
	pulumi.RegisterOutputType(MigrateJobConfigArrayOutput{})
	pulumi.RegisterOutputType(MigrateJobConfigMapOutput{})
}
