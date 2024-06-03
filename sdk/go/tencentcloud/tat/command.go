// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tat

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a tat command
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tat"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tat.NewCommand(ctx, "command", &Tat.CommandArgs{
//				CommandName: pulumi.String("ls"),
//				CommandType: pulumi.String("SHELL"),
//				Content:     pulumi.String("bHM="),
//				Description: pulumi.String("xxx"),
//				Tags: tat.CommandTagArray{
//					&tat.CommandTagArgs{
//						Key:   pulumi.String(""),
//						Value: pulumi.String(""),
//					},
//				},
//				Timeout:          pulumi.Int(50),
//				Username:         pulumi.String("root"),
//				WorkingDirectory: pulumi.String("/root"),
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
// tat command can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Tat/command:Command command cmd-6fydo27j
// ```
type Command struct {
	pulumi.CustomResourceState

	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName pulumi.StringOutput `pulumi:"commandName"`
	// Command type. `SHELL` and `POWERSHELL` are supported. The default value is `SHELL`.
	CommandType pulumi.StringPtrOutput `pulumi:"commandType"`
	// Command. The maximum length of Base64 encoding is 64KB.
	Content pulumi.StringOutput `pulumi:"content"`
	// Command creator. `TAT` indicates a public command and `USER` indicates a personal command.
	CreatedBy pulumi.StringOutput `pulumi:"createdBy"`
	// Command creation time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {&amp;#39;varA&amp;#39;: &amp;#39;222&amp;#39;}.`key` is the name of the custom parameter and value is the default value. Both `key` and `value` are strings.If no parameter value is provided in the `InvokeCommand` API, the default value is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters pulumi.StringPtrOutput `pulumi:"defaultParameters"`
	// Command description. The maximum length is 120 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether to enable the custom parameter feature.This cannot be modified once created.Default value: `false`.
	EnableParameter pulumi.BoolPtrOutput `pulumi:"enableParameter"`
	// Formatted description of the command. This parameter is an empty string for user commands and contains values for public commands.
	FormattedDescription pulumi.StringOutput `pulumi:"formattedDescription"`
	// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
	OutputCosBucketUrl pulumi.StringPtrOutput `pulumi:"outputCosBucketUrl"`
	// The COS bucket directory where the logs are saved. Check below for the rules of the directory name.1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.2. Use a slash (/) to create a subdirectory.3. Consecutive dots (.) and slashes (/) are not allowed. It can not start with a slash (/).
	OutputCosKeyPrefix pulumi.StringPtrOutput `pulumi:"outputCosKeyPrefix"`
	// Tags bound to the command. At most 10 tags are allowed.
	Tags CommandTagArrayOutput `pulumi:"tags"`
	// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// Command update time.
	UpdatedTime pulumi.StringOutput `pulumi:"updatedTime"`
	// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the root user is used to execute commands on Linux and the System user is used on Windows.
	Username pulumi.StringPtrOutput `pulumi:"username"`
	// Command execution path. The default value is /root for `SHELL` commands and C:/Program Files/qcloudtat_agent/workdir for `POWERSHELL` commands.
	WorkingDirectory pulumi.StringPtrOutput `pulumi:"workingDirectory"`
}

// NewCommand registers a new resource with the given unique name, arguments, and options.
func NewCommand(ctx *pulumi.Context,
	name string, args *CommandArgs, opts ...pulumi.ResourceOption) (*Command, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CommandName == nil {
		return nil, errors.New("invalid value for required argument 'CommandName'")
	}
	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Command
	err := ctx.RegisterResource("tencentcloud:Tat/command:Command", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommand gets an existing Command resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommand(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommandState, opts ...pulumi.ResourceOption) (*Command, error) {
	var resource Command
	err := ctx.ReadResource("tencentcloud:Tat/command:Command", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Command resources.
type commandState struct {
	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName *string `pulumi:"commandName"`
	// Command type. `SHELL` and `POWERSHELL` are supported. The default value is `SHELL`.
	CommandType *string `pulumi:"commandType"`
	// Command. The maximum length of Base64 encoding is 64KB.
	Content *string `pulumi:"content"`
	// Command creator. `TAT` indicates a public command and `USER` indicates a personal command.
	CreatedBy *string `pulumi:"createdBy"`
	// Command creation time.
	CreatedTime *string `pulumi:"createdTime"`
	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {&amp;#39;varA&amp;#39;: &amp;#39;222&amp;#39;}.`key` is the name of the custom parameter and value is the default value. Both `key` and `value` are strings.If no parameter value is provided in the `InvokeCommand` API, the default value is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters *string `pulumi:"defaultParameters"`
	// Command description. The maximum length is 120 characters.
	Description *string `pulumi:"description"`
	// Whether to enable the custom parameter feature.This cannot be modified once created.Default value: `false`.
	EnableParameter *bool `pulumi:"enableParameter"`
	// Formatted description of the command. This parameter is an empty string for user commands and contains values for public commands.
	FormattedDescription *string `pulumi:"formattedDescription"`
	// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
	OutputCosBucketUrl *string `pulumi:"outputCosBucketUrl"`
	// The COS bucket directory where the logs are saved. Check below for the rules of the directory name.1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.2. Use a slash (/) to create a subdirectory.3. Consecutive dots (.) and slashes (/) are not allowed. It can not start with a slash (/).
	OutputCosKeyPrefix *string `pulumi:"outputCosKeyPrefix"`
	// Tags bound to the command. At most 10 tags are allowed.
	Tags []CommandTag `pulumi:"tags"`
	// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
	Timeout *int `pulumi:"timeout"`
	// Command update time.
	UpdatedTime *string `pulumi:"updatedTime"`
	// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the root user is used to execute commands on Linux and the System user is used on Windows.
	Username *string `pulumi:"username"`
	// Command execution path. The default value is /root for `SHELL` commands and C:/Program Files/qcloudtat_agent/workdir for `POWERSHELL` commands.
	WorkingDirectory *string `pulumi:"workingDirectory"`
}

type CommandState struct {
	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName pulumi.StringPtrInput
	// Command type. `SHELL` and `POWERSHELL` are supported. The default value is `SHELL`.
	CommandType pulumi.StringPtrInput
	// Command. The maximum length of Base64 encoding is 64KB.
	Content pulumi.StringPtrInput
	// Command creator. `TAT` indicates a public command and `USER` indicates a personal command.
	CreatedBy pulumi.StringPtrInput
	// Command creation time.
	CreatedTime pulumi.StringPtrInput
	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {&amp;#39;varA&amp;#39;: &amp;#39;222&amp;#39;}.`key` is the name of the custom parameter and value is the default value. Both `key` and `value` are strings.If no parameter value is provided in the `InvokeCommand` API, the default value is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters pulumi.StringPtrInput
	// Command description. The maximum length is 120 characters.
	Description pulumi.StringPtrInput
	// Whether to enable the custom parameter feature.This cannot be modified once created.Default value: `false`.
	EnableParameter pulumi.BoolPtrInput
	// Formatted description of the command. This parameter is an empty string for user commands and contains values for public commands.
	FormattedDescription pulumi.StringPtrInput
	// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
	OutputCosBucketUrl pulumi.StringPtrInput
	// The COS bucket directory where the logs are saved. Check below for the rules of the directory name.1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.2. Use a slash (/) to create a subdirectory.3. Consecutive dots (.) and slashes (/) are not allowed. It can not start with a slash (/).
	OutputCosKeyPrefix pulumi.StringPtrInput
	// Tags bound to the command. At most 10 tags are allowed.
	Tags CommandTagArrayInput
	// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
	Timeout pulumi.IntPtrInput
	// Command update time.
	UpdatedTime pulumi.StringPtrInput
	// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the root user is used to execute commands on Linux and the System user is used on Windows.
	Username pulumi.StringPtrInput
	// Command execution path. The default value is /root for `SHELL` commands and C:/Program Files/qcloudtat_agent/workdir for `POWERSHELL` commands.
	WorkingDirectory pulumi.StringPtrInput
}

func (CommandState) ElementType() reflect.Type {
	return reflect.TypeOf((*commandState)(nil)).Elem()
}

type commandArgs struct {
	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName string `pulumi:"commandName"`
	// Command type. `SHELL` and `POWERSHELL` are supported. The default value is `SHELL`.
	CommandType *string `pulumi:"commandType"`
	// Command. The maximum length of Base64 encoding is 64KB.
	Content string `pulumi:"content"`
	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {&amp;#39;varA&amp;#39;: &amp;#39;222&amp;#39;}.`key` is the name of the custom parameter and value is the default value. Both `key` and `value` are strings.If no parameter value is provided in the `InvokeCommand` API, the default value is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters *string `pulumi:"defaultParameters"`
	// Command description. The maximum length is 120 characters.
	Description *string `pulumi:"description"`
	// Whether to enable the custom parameter feature.This cannot be modified once created.Default value: `false`.
	EnableParameter *bool `pulumi:"enableParameter"`
	// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
	OutputCosBucketUrl *string `pulumi:"outputCosBucketUrl"`
	// The COS bucket directory where the logs are saved. Check below for the rules of the directory name.1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.2. Use a slash (/) to create a subdirectory.3. Consecutive dots (.) and slashes (/) are not allowed. It can not start with a slash (/).
	OutputCosKeyPrefix *string `pulumi:"outputCosKeyPrefix"`
	// Tags bound to the command. At most 10 tags are allowed.
	Tags []CommandTag `pulumi:"tags"`
	// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
	Timeout *int `pulumi:"timeout"`
	// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the root user is used to execute commands on Linux and the System user is used on Windows.
	Username *string `pulumi:"username"`
	// Command execution path. The default value is /root for `SHELL` commands and C:/Program Files/qcloudtat_agent/workdir for `POWERSHELL` commands.
	WorkingDirectory *string `pulumi:"workingDirectory"`
}

// The set of arguments for constructing a Command resource.
type CommandArgs struct {
	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName pulumi.StringInput
	// Command type. `SHELL` and `POWERSHELL` are supported. The default value is `SHELL`.
	CommandType pulumi.StringPtrInput
	// Command. The maximum length of Base64 encoding is 64KB.
	Content pulumi.StringInput
	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {&amp;#39;varA&amp;#39;: &amp;#39;222&amp;#39;}.`key` is the name of the custom parameter and value is the default value. Both `key` and `value` are strings.If no parameter value is provided in the `InvokeCommand` API, the default value is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters pulumi.StringPtrInput
	// Command description. The maximum length is 120 characters.
	Description pulumi.StringPtrInput
	// Whether to enable the custom parameter feature.This cannot be modified once created.Default value: `false`.
	EnableParameter pulumi.BoolPtrInput
	// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
	OutputCosBucketUrl pulumi.StringPtrInput
	// The COS bucket directory where the logs are saved. Check below for the rules of the directory name.1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.2. Use a slash (/) to create a subdirectory.3. Consecutive dots (.) and slashes (/) are not allowed. It can not start with a slash (/).
	OutputCosKeyPrefix pulumi.StringPtrInput
	// Tags bound to the command. At most 10 tags are allowed.
	Tags CommandTagArrayInput
	// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
	Timeout pulumi.IntPtrInput
	// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the root user is used to execute commands on Linux and the System user is used on Windows.
	Username pulumi.StringPtrInput
	// Command execution path. The default value is /root for `SHELL` commands and C:/Program Files/qcloudtat_agent/workdir for `POWERSHELL` commands.
	WorkingDirectory pulumi.StringPtrInput
}

func (CommandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*commandArgs)(nil)).Elem()
}

type CommandInput interface {
	pulumi.Input

	ToCommandOutput() CommandOutput
	ToCommandOutputWithContext(ctx context.Context) CommandOutput
}

func (*Command) ElementType() reflect.Type {
	return reflect.TypeOf((**Command)(nil)).Elem()
}

func (i *Command) ToCommandOutput() CommandOutput {
	return i.ToCommandOutputWithContext(context.Background())
}

func (i *Command) ToCommandOutputWithContext(ctx context.Context) CommandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandOutput)
}

// CommandArrayInput is an input type that accepts CommandArray and CommandArrayOutput values.
// You can construct a concrete instance of `CommandArrayInput` via:
//
//	CommandArray{ CommandArgs{...} }
type CommandArrayInput interface {
	pulumi.Input

	ToCommandArrayOutput() CommandArrayOutput
	ToCommandArrayOutputWithContext(context.Context) CommandArrayOutput
}

type CommandArray []CommandInput

func (CommandArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Command)(nil)).Elem()
}

func (i CommandArray) ToCommandArrayOutput() CommandArrayOutput {
	return i.ToCommandArrayOutputWithContext(context.Background())
}

func (i CommandArray) ToCommandArrayOutputWithContext(ctx context.Context) CommandArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandArrayOutput)
}

// CommandMapInput is an input type that accepts CommandMap and CommandMapOutput values.
// You can construct a concrete instance of `CommandMapInput` via:
//
//	CommandMap{ "key": CommandArgs{...} }
type CommandMapInput interface {
	pulumi.Input

	ToCommandMapOutput() CommandMapOutput
	ToCommandMapOutputWithContext(context.Context) CommandMapOutput
}

type CommandMap map[string]CommandInput

func (CommandMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Command)(nil)).Elem()
}

func (i CommandMap) ToCommandMapOutput() CommandMapOutput {
	return i.ToCommandMapOutputWithContext(context.Background())
}

func (i CommandMap) ToCommandMapOutputWithContext(ctx context.Context) CommandMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandMapOutput)
}

type CommandOutput struct{ *pulumi.OutputState }

func (CommandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Command)(nil)).Elem()
}

func (o CommandOutput) ToCommandOutput() CommandOutput {
	return o
}

func (o CommandOutput) ToCommandOutputWithContext(ctx context.Context) CommandOutput {
	return o
}

// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
func (o CommandOutput) CommandName() pulumi.StringOutput {
	return o.ApplyT(func(v *Command) pulumi.StringOutput { return v.CommandName }).(pulumi.StringOutput)
}

// Command type. `SHELL` and `POWERSHELL` are supported. The default value is `SHELL`.
func (o CommandOutput) CommandType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Command) pulumi.StringPtrOutput { return v.CommandType }).(pulumi.StringPtrOutput)
}

// Command. The maximum length of Base64 encoding is 64KB.
func (o CommandOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *Command) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// Command creator. `TAT` indicates a public command and `USER` indicates a personal command.
func (o CommandOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Command) pulumi.StringOutput { return v.CreatedBy }).(pulumi.StringOutput)
}

// Command creation time.
func (o CommandOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Command) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {&amp;#39;varA&amp;#39;: &amp;#39;222&amp;#39;}.`key` is the name of the custom parameter and value is the default value. Both `key` and `value` are strings.If no parameter value is provided in the `InvokeCommand` API, the default value is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
func (o CommandOutput) DefaultParameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Command) pulumi.StringPtrOutput { return v.DefaultParameters }).(pulumi.StringPtrOutput)
}

// Command description. The maximum length is 120 characters.
func (o CommandOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Command) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether to enable the custom parameter feature.This cannot be modified once created.Default value: `false`.
func (o CommandOutput) EnableParameter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Command) pulumi.BoolPtrOutput { return v.EnableParameter }).(pulumi.BoolPtrOutput)
}

// Formatted description of the command. This parameter is an empty string for user commands and contains values for public commands.
func (o CommandOutput) FormattedDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *Command) pulumi.StringOutput { return v.FormattedDescription }).(pulumi.StringOutput)
}

// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
func (o CommandOutput) OutputCosBucketUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Command) pulumi.StringPtrOutput { return v.OutputCosBucketUrl }).(pulumi.StringPtrOutput)
}

// The COS bucket directory where the logs are saved. Check below for the rules of the directory name.1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.2. Use a slash (/) to create a subdirectory.3. Consecutive dots (.) and slashes (/) are not allowed. It can not start with a slash (/).
func (o CommandOutput) OutputCosKeyPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Command) pulumi.StringPtrOutput { return v.OutputCosKeyPrefix }).(pulumi.StringPtrOutput)
}

// Tags bound to the command. At most 10 tags are allowed.
func (o CommandOutput) Tags() CommandTagArrayOutput {
	return o.ApplyT(func(v *Command) CommandTagArrayOutput { return v.Tags }).(CommandTagArrayOutput)
}

// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
func (o CommandOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Command) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

// Command update time.
func (o CommandOutput) UpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Command) pulumi.StringOutput { return v.UpdatedTime }).(pulumi.StringOutput)
}

// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the root user is used to execute commands on Linux and the System user is used on Windows.
func (o CommandOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Command) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

// Command execution path. The default value is /root for `SHELL` commands and C:/Program Files/qcloudtat_agent/workdir for `POWERSHELL` commands.
func (o CommandOutput) WorkingDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Command) pulumi.StringPtrOutput { return v.WorkingDirectory }).(pulumi.StringPtrOutput)
}

type CommandArrayOutput struct{ *pulumi.OutputState }

func (CommandArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Command)(nil)).Elem()
}

func (o CommandArrayOutput) ToCommandArrayOutput() CommandArrayOutput {
	return o
}

func (o CommandArrayOutput) ToCommandArrayOutputWithContext(ctx context.Context) CommandArrayOutput {
	return o
}

func (o CommandArrayOutput) Index(i pulumi.IntInput) CommandOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Command {
		return vs[0].([]*Command)[vs[1].(int)]
	}).(CommandOutput)
}

type CommandMapOutput struct{ *pulumi.OutputState }

func (CommandMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Command)(nil)).Elem()
}

func (o CommandMapOutput) ToCommandMapOutput() CommandMapOutput {
	return o
}

func (o CommandMapOutput) ToCommandMapOutputWithContext(ctx context.Context) CommandMapOutput {
	return o
}

func (o CommandMapOutput) MapIndex(k pulumi.StringInput) CommandOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Command {
		return vs[0].(map[string]*Command)[vs[1].(string)]
	}).(CommandOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CommandInput)(nil)).Elem(), &Command{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommandArrayInput)(nil)).Elem(), CommandArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommandMapInput)(nil)).Elem(), CommandMap{})
	pulumi.RegisterOutputType(CommandOutput{})
	pulumi.RegisterOutputType(CommandArrayOutput{})
	pulumi.RegisterOutputType(CommandMapOutput{})
}
