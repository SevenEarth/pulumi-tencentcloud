// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tem

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tem logConfig
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tem"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Tem.NewLogConfig(ctx, "logConfig", &Tem.LogConfigArgs{
// 			EnvironmentId: pulumi.String("en-o5edaepv"),
// 			ApplicationId: pulumi.String("app-3j29aa2p"),
// 			WorkloadId:    pulumi.Any(resource.Tencentcloud_tem_workload.Workload.Id),
// 			LogsetId:      pulumi.String("b5824781-8d5b-4029-a2f7-d03c37f72bdf"),
// 			TopicId:       pulumi.String("5a85bb6d-8e41-4e04-b7bd-c05e04782f94"),
// 			InputType:     pulumi.String("container_stdout"),
// 			LogType:       pulumi.String("minimalist_log"),
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
// tem logConfig can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Tem/logConfig:LogConfig logConfig environmentId#applicationId#name
// ```
type LogConfig struct {
	pulumi.CustomResourceState

	// application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// regex pattern.
	BeginningRegex pulumi.StringPtrOutput `pulumi:"beginningRegex"`
	// environment ID.
	EnvironmentId pulumi.StringOutput `pulumi:"environmentId"`
	// file name pattern if container_file.
	FilePattern pulumi.StringPtrOutput `pulumi:"filePattern"`
	// container_stdout or container_file.
	InputType pulumi.StringOutput `pulumi:"inputType"`
	// directory if container_file.
	LogPath pulumi.StringPtrOutput `pulumi:"logPath"`
	// minimalist_log or multiline_log.
	LogType pulumi.StringOutput `pulumi:"logType"`
	// logset.
	LogsetId pulumi.StringOutput `pulumi:"logsetId"`
	// appConfig name.
	Name pulumi.StringOutput `pulumi:"name"`
	// topic.
	TopicId pulumi.StringOutput `pulumi:"topicId"`
	// application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
	WorkloadId pulumi.StringOutput `pulumi:"workloadId"`
}

// NewLogConfig registers a new resource with the given unique name, arguments, and options.
func NewLogConfig(ctx *pulumi.Context,
	name string, args *LogConfigArgs, opts ...pulumi.ResourceOption) (*LogConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.InputType == nil {
		return nil, errors.New("invalid value for required argument 'InputType'")
	}
	if args.LogType == nil {
		return nil, errors.New("invalid value for required argument 'LogType'")
	}
	if args.LogsetId == nil {
		return nil, errors.New("invalid value for required argument 'LogsetId'")
	}
	if args.TopicId == nil {
		return nil, errors.New("invalid value for required argument 'TopicId'")
	}
	if args.WorkloadId == nil {
		return nil, errors.New("invalid value for required argument 'WorkloadId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource LogConfig
	err := ctx.RegisterResource("tencentcloud:Tem/logConfig:LogConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogConfig gets an existing LogConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogConfigState, opts ...pulumi.ResourceOption) (*LogConfig, error) {
	var resource LogConfig
	err := ctx.ReadResource("tencentcloud:Tem/logConfig:LogConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogConfig resources.
type logConfigState struct {
	// application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// regex pattern.
	BeginningRegex *string `pulumi:"beginningRegex"`
	// environment ID.
	EnvironmentId *string `pulumi:"environmentId"`
	// file name pattern if container_file.
	FilePattern *string `pulumi:"filePattern"`
	// container_stdout or container_file.
	InputType *string `pulumi:"inputType"`
	// directory if container_file.
	LogPath *string `pulumi:"logPath"`
	// minimalist_log or multiline_log.
	LogType *string `pulumi:"logType"`
	// logset.
	LogsetId *string `pulumi:"logsetId"`
	// appConfig name.
	Name *string `pulumi:"name"`
	// topic.
	TopicId *string `pulumi:"topicId"`
	// application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
	WorkloadId *string `pulumi:"workloadId"`
}

type LogConfigState struct {
	// application ID.
	ApplicationId pulumi.StringPtrInput
	// regex pattern.
	BeginningRegex pulumi.StringPtrInput
	// environment ID.
	EnvironmentId pulumi.StringPtrInput
	// file name pattern if container_file.
	FilePattern pulumi.StringPtrInput
	// container_stdout or container_file.
	InputType pulumi.StringPtrInput
	// directory if container_file.
	LogPath pulumi.StringPtrInput
	// minimalist_log or multiline_log.
	LogType pulumi.StringPtrInput
	// logset.
	LogsetId pulumi.StringPtrInput
	// appConfig name.
	Name pulumi.StringPtrInput
	// topic.
	TopicId pulumi.StringPtrInput
	// application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
	WorkloadId pulumi.StringPtrInput
}

func (LogConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*logConfigState)(nil)).Elem()
}

type logConfigArgs struct {
	// application ID.
	ApplicationId string `pulumi:"applicationId"`
	// regex pattern.
	BeginningRegex *string `pulumi:"beginningRegex"`
	// environment ID.
	EnvironmentId string `pulumi:"environmentId"`
	// file name pattern if container_file.
	FilePattern *string `pulumi:"filePattern"`
	// container_stdout or container_file.
	InputType string `pulumi:"inputType"`
	// directory if container_file.
	LogPath *string `pulumi:"logPath"`
	// minimalist_log or multiline_log.
	LogType string `pulumi:"logType"`
	// logset.
	LogsetId string `pulumi:"logsetId"`
	// appConfig name.
	Name *string `pulumi:"name"`
	// topic.
	TopicId string `pulumi:"topicId"`
	// application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
	WorkloadId string `pulumi:"workloadId"`
}

// The set of arguments for constructing a LogConfig resource.
type LogConfigArgs struct {
	// application ID.
	ApplicationId pulumi.StringInput
	// regex pattern.
	BeginningRegex pulumi.StringPtrInput
	// environment ID.
	EnvironmentId pulumi.StringInput
	// file name pattern if container_file.
	FilePattern pulumi.StringPtrInput
	// container_stdout or container_file.
	InputType pulumi.StringInput
	// directory if container_file.
	LogPath pulumi.StringPtrInput
	// minimalist_log or multiline_log.
	LogType pulumi.StringInput
	// logset.
	LogsetId pulumi.StringInput
	// appConfig name.
	Name pulumi.StringPtrInput
	// topic.
	TopicId pulumi.StringInput
	// application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
	WorkloadId pulumi.StringInput
}

func (LogConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logConfigArgs)(nil)).Elem()
}

type LogConfigInput interface {
	pulumi.Input

	ToLogConfigOutput() LogConfigOutput
	ToLogConfigOutputWithContext(ctx context.Context) LogConfigOutput
}

func (*LogConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**LogConfig)(nil)).Elem()
}

func (i *LogConfig) ToLogConfigOutput() LogConfigOutput {
	return i.ToLogConfigOutputWithContext(context.Background())
}

func (i *LogConfig) ToLogConfigOutputWithContext(ctx context.Context) LogConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogConfigOutput)
}

// LogConfigArrayInput is an input type that accepts LogConfigArray and LogConfigArrayOutput values.
// You can construct a concrete instance of `LogConfigArrayInput` via:
//
//          LogConfigArray{ LogConfigArgs{...} }
type LogConfigArrayInput interface {
	pulumi.Input

	ToLogConfigArrayOutput() LogConfigArrayOutput
	ToLogConfigArrayOutputWithContext(context.Context) LogConfigArrayOutput
}

type LogConfigArray []LogConfigInput

func (LogConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogConfig)(nil)).Elem()
}

func (i LogConfigArray) ToLogConfigArrayOutput() LogConfigArrayOutput {
	return i.ToLogConfigArrayOutputWithContext(context.Background())
}

func (i LogConfigArray) ToLogConfigArrayOutputWithContext(ctx context.Context) LogConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogConfigArrayOutput)
}

// LogConfigMapInput is an input type that accepts LogConfigMap and LogConfigMapOutput values.
// You can construct a concrete instance of `LogConfigMapInput` via:
//
//          LogConfigMap{ "key": LogConfigArgs{...} }
type LogConfigMapInput interface {
	pulumi.Input

	ToLogConfigMapOutput() LogConfigMapOutput
	ToLogConfigMapOutputWithContext(context.Context) LogConfigMapOutput
}

type LogConfigMap map[string]LogConfigInput

func (LogConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogConfig)(nil)).Elem()
}

func (i LogConfigMap) ToLogConfigMapOutput() LogConfigMapOutput {
	return i.ToLogConfigMapOutputWithContext(context.Background())
}

func (i LogConfigMap) ToLogConfigMapOutputWithContext(ctx context.Context) LogConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogConfigMapOutput)
}

type LogConfigOutput struct{ *pulumi.OutputState }

func (LogConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogConfig)(nil)).Elem()
}

func (o LogConfigOutput) ToLogConfigOutput() LogConfigOutput {
	return o
}

func (o LogConfigOutput) ToLogConfigOutputWithContext(ctx context.Context) LogConfigOutput {
	return o
}

// application ID.
func (o LogConfigOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogConfig) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// regex pattern.
func (o LogConfigOutput) BeginningRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogConfig) pulumi.StringPtrOutput { return v.BeginningRegex }).(pulumi.StringPtrOutput)
}

// environment ID.
func (o LogConfigOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogConfig) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

// file name pattern if container_file.
func (o LogConfigOutput) FilePattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogConfig) pulumi.StringPtrOutput { return v.FilePattern }).(pulumi.StringPtrOutput)
}

// container_stdout or container_file.
func (o LogConfigOutput) InputType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogConfig) pulumi.StringOutput { return v.InputType }).(pulumi.StringOutput)
}

// directory if container_file.
func (o LogConfigOutput) LogPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogConfig) pulumi.StringPtrOutput { return v.LogPath }).(pulumi.StringPtrOutput)
}

// minimalist_log or multiline_log.
func (o LogConfigOutput) LogType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogConfig) pulumi.StringOutput { return v.LogType }).(pulumi.StringOutput)
}

// logset.
func (o LogConfigOutput) LogsetId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogConfig) pulumi.StringOutput { return v.LogsetId }).(pulumi.StringOutput)
}

// appConfig name.
func (o LogConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LogConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// topic.
func (o LogConfigOutput) TopicId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogConfig) pulumi.StringOutput { return v.TopicId }).(pulumi.StringOutput)
}

// application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
func (o LogConfigOutput) WorkloadId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogConfig) pulumi.StringOutput { return v.WorkloadId }).(pulumi.StringOutput)
}

type LogConfigArrayOutput struct{ *pulumi.OutputState }

func (LogConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogConfig)(nil)).Elem()
}

func (o LogConfigArrayOutput) ToLogConfigArrayOutput() LogConfigArrayOutput {
	return o
}

func (o LogConfigArrayOutput) ToLogConfigArrayOutputWithContext(ctx context.Context) LogConfigArrayOutput {
	return o
}

func (o LogConfigArrayOutput) Index(i pulumi.IntInput) LogConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogConfig {
		return vs[0].([]*LogConfig)[vs[1].(int)]
	}).(LogConfigOutput)
}

type LogConfigMapOutput struct{ *pulumi.OutputState }

func (LogConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogConfig)(nil)).Elem()
}

func (o LogConfigMapOutput) ToLogConfigMapOutput() LogConfigMapOutput {
	return o
}

func (o LogConfigMapOutput) ToLogConfigMapOutputWithContext(ctx context.Context) LogConfigMapOutput {
	return o
}

func (o LogConfigMapOutput) MapIndex(k pulumi.StringInput) LogConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogConfig {
		return vs[0].(map[string]*LogConfig)[vs[1].(string)]
	}).(LogConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogConfigInput)(nil)).Elem(), &LogConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogConfigArrayInput)(nil)).Elem(), LogConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogConfigMapInput)(nil)).Elem(), LogConfigMap{})
	pulumi.RegisterOutputType(LogConfigOutput{})
	pulumi.RegisterOutputType(LogConfigArrayOutput{})
	pulumi.RegisterOutputType(LogConfigMapOutput{})
}
