// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tat

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tat invocationCommandAttachment
type InvocationCommandAttachment struct {
	pulumi.CustomResourceState

	// Command ID.
	CommandId pulumi.StringOutput `pulumi:"commandId"`
	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName pulumi.StringPtrOutput `pulumi:"commandName"`
	// Command type. SHELL and POWERSHELL are supported. The default value is SHELL.
	CommandType pulumi.StringPtrOutput `pulumi:"commandType"`
	// Base64-encoded command. The maximum length is 64 KB.
	Content pulumi.StringOutput `pulumi:"content"`
	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If Parameters is not provided, the default values specified here are used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters pulumi.StringPtrOutput `pulumi:"defaultParameters"`
	// Command description. The maximum length is 120 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether to enable the custom parameter feature.This cannot be modified once created.Default value: false.
	EnableParameter pulumi.BoolPtrOutput `pulumi:"enableParameter"`
	// ID of instances about to execute commands. Supported instance types:  CVM  LIGHTHOUSE.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The COS bucket URL for uploading logs; The URL must start with https, such as https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
	OutputCosBucketUrl pulumi.StringPtrOutput `pulumi:"outputCosBucketUrl"`
	// The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive slashes.
	OutputCosKeyPrefix pulumi.StringPtrOutput `pulumi:"outputCosKeyPrefix"`
	// Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	Parameters pulumi.StringPtrOutput `pulumi:"parameters"`
	// Whether to save the command. Valid values:rue: SaveFalse:Do not saveThe default value is False.
	SaveCommand pulumi.BoolPtrOutput `pulumi:"saveCommand"`
	// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root is used to execute commands on Linux and the user System is used on Windows.
	Username pulumi.StringPtrOutput `pulumi:"username"`
	// Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for POWERSHELL commands.
	WorkingDirectory pulumi.StringPtrOutput `pulumi:"workingDirectory"`
}

// NewInvocationCommandAttachment registers a new resource with the given unique name, arguments, and options.
func NewInvocationCommandAttachment(ctx *pulumi.Context,
	name string, args *InvocationCommandAttachmentArgs, opts ...pulumi.ResourceOption) (*InvocationCommandAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource InvocationCommandAttachment
	err := ctx.RegisterResource("tencentcloud:Tat/invocationCommandAttachment:InvocationCommandAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInvocationCommandAttachment gets an existing InvocationCommandAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInvocationCommandAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InvocationCommandAttachmentState, opts ...pulumi.ResourceOption) (*InvocationCommandAttachment, error) {
	var resource InvocationCommandAttachment
	err := ctx.ReadResource("tencentcloud:Tat/invocationCommandAttachment:InvocationCommandAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InvocationCommandAttachment resources.
type invocationCommandAttachmentState struct {
	// Command ID.
	CommandId *string `pulumi:"commandId"`
	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName *string `pulumi:"commandName"`
	// Command type. SHELL and POWERSHELL are supported. The default value is SHELL.
	CommandType *string `pulumi:"commandType"`
	// Base64-encoded command. The maximum length is 64 KB.
	Content *string `pulumi:"content"`
	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If Parameters is not provided, the default values specified here are used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters *string `pulumi:"defaultParameters"`
	// Command description. The maximum length is 120 characters.
	Description *string `pulumi:"description"`
	// Whether to enable the custom parameter feature.This cannot be modified once created.Default value: false.
	EnableParameter *bool `pulumi:"enableParameter"`
	// ID of instances about to execute commands. Supported instance types:  CVM  LIGHTHOUSE.
	InstanceId *string `pulumi:"instanceId"`
	// The COS bucket URL for uploading logs; The URL must start with https, such as https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
	OutputCosBucketUrl *string `pulumi:"outputCosBucketUrl"`
	// The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive slashes.
	OutputCosKeyPrefix *string `pulumi:"outputCosKeyPrefix"`
	// Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	Parameters *string `pulumi:"parameters"`
	// Whether to save the command. Valid values:rue: SaveFalse:Do not saveThe default value is False.
	SaveCommand *bool `pulumi:"saveCommand"`
	// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
	Timeout *int `pulumi:"timeout"`
	// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root is used to execute commands on Linux and the user System is used on Windows.
	Username *string `pulumi:"username"`
	// Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for POWERSHELL commands.
	WorkingDirectory *string `pulumi:"workingDirectory"`
}

type InvocationCommandAttachmentState struct {
	// Command ID.
	CommandId pulumi.StringPtrInput
	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName pulumi.StringPtrInput
	// Command type. SHELL and POWERSHELL are supported. The default value is SHELL.
	CommandType pulumi.StringPtrInput
	// Base64-encoded command. The maximum length is 64 KB.
	Content pulumi.StringPtrInput
	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If Parameters is not provided, the default values specified here are used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters pulumi.StringPtrInput
	// Command description. The maximum length is 120 characters.
	Description pulumi.StringPtrInput
	// Whether to enable the custom parameter feature.This cannot be modified once created.Default value: false.
	EnableParameter pulumi.BoolPtrInput
	// ID of instances about to execute commands. Supported instance types:  CVM  LIGHTHOUSE.
	InstanceId pulumi.StringPtrInput
	// The COS bucket URL for uploading logs; The URL must start with https, such as https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
	OutputCosBucketUrl pulumi.StringPtrInput
	// The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive slashes.
	OutputCosKeyPrefix pulumi.StringPtrInput
	// Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	Parameters pulumi.StringPtrInput
	// Whether to save the command. Valid values:rue: SaveFalse:Do not saveThe default value is False.
	SaveCommand pulumi.BoolPtrInput
	// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
	Timeout pulumi.IntPtrInput
	// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root is used to execute commands on Linux and the user System is used on Windows.
	Username pulumi.StringPtrInput
	// Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for POWERSHELL commands.
	WorkingDirectory pulumi.StringPtrInput
}

func (InvocationCommandAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*invocationCommandAttachmentState)(nil)).Elem()
}

type invocationCommandAttachmentArgs struct {
	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName *string `pulumi:"commandName"`
	// Command type. SHELL and POWERSHELL are supported. The default value is SHELL.
	CommandType *string `pulumi:"commandType"`
	// Base64-encoded command. The maximum length is 64 KB.
	Content string `pulumi:"content"`
	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If Parameters is not provided, the default values specified here are used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters *string `pulumi:"defaultParameters"`
	// Command description. The maximum length is 120 characters.
	Description *string `pulumi:"description"`
	// Whether to enable the custom parameter feature.This cannot be modified once created.Default value: false.
	EnableParameter *bool `pulumi:"enableParameter"`
	// ID of instances about to execute commands. Supported instance types:  CVM  LIGHTHOUSE.
	InstanceId string `pulumi:"instanceId"`
	// The COS bucket URL for uploading logs; The URL must start with https, such as https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
	OutputCosBucketUrl *string `pulumi:"outputCosBucketUrl"`
	// The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive slashes.
	OutputCosKeyPrefix *string `pulumi:"outputCosKeyPrefix"`
	// Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	Parameters *string `pulumi:"parameters"`
	// Whether to save the command. Valid values:rue: SaveFalse:Do not saveThe default value is False.
	SaveCommand *bool `pulumi:"saveCommand"`
	// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
	Timeout *int `pulumi:"timeout"`
	// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root is used to execute commands on Linux and the user System is used on Windows.
	Username *string `pulumi:"username"`
	// Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for POWERSHELL commands.
	WorkingDirectory *string `pulumi:"workingDirectory"`
}

// The set of arguments for constructing a InvocationCommandAttachment resource.
type InvocationCommandAttachmentArgs struct {
	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName pulumi.StringPtrInput
	// Command type. SHELL and POWERSHELL are supported. The default value is SHELL.
	CommandType pulumi.StringPtrInput
	// Base64-encoded command. The maximum length is 64 KB.
	Content pulumi.StringInput
	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If Parameters is not provided, the default values specified here are used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters pulumi.StringPtrInput
	// Command description. The maximum length is 120 characters.
	Description pulumi.StringPtrInput
	// Whether to enable the custom parameter feature.This cannot be modified once created.Default value: false.
	EnableParameter pulumi.BoolPtrInput
	// ID of instances about to execute commands. Supported instance types:  CVM  LIGHTHOUSE.
	InstanceId pulumi.StringInput
	// The COS bucket URL for uploading logs; The URL must start with https, such as https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
	OutputCosBucketUrl pulumi.StringPtrInput
	// The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive slashes.
	OutputCosKeyPrefix pulumi.StringPtrInput
	// Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	Parameters pulumi.StringPtrInput
	// Whether to save the command. Valid values:rue: SaveFalse:Do not saveThe default value is False.
	SaveCommand pulumi.BoolPtrInput
	// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
	Timeout pulumi.IntPtrInput
	// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root is used to execute commands on Linux and the user System is used on Windows.
	Username pulumi.StringPtrInput
	// Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for POWERSHELL commands.
	WorkingDirectory pulumi.StringPtrInput
}

func (InvocationCommandAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*invocationCommandAttachmentArgs)(nil)).Elem()
}

type InvocationCommandAttachmentInput interface {
	pulumi.Input

	ToInvocationCommandAttachmentOutput() InvocationCommandAttachmentOutput
	ToInvocationCommandAttachmentOutputWithContext(ctx context.Context) InvocationCommandAttachmentOutput
}

func (*InvocationCommandAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**InvocationCommandAttachment)(nil)).Elem()
}

func (i *InvocationCommandAttachment) ToInvocationCommandAttachmentOutput() InvocationCommandAttachmentOutput {
	return i.ToInvocationCommandAttachmentOutputWithContext(context.Background())
}

func (i *InvocationCommandAttachment) ToInvocationCommandAttachmentOutputWithContext(ctx context.Context) InvocationCommandAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InvocationCommandAttachmentOutput)
}

// InvocationCommandAttachmentArrayInput is an input type that accepts InvocationCommandAttachmentArray and InvocationCommandAttachmentArrayOutput values.
// You can construct a concrete instance of `InvocationCommandAttachmentArrayInput` via:
//
//          InvocationCommandAttachmentArray{ InvocationCommandAttachmentArgs{...} }
type InvocationCommandAttachmentArrayInput interface {
	pulumi.Input

	ToInvocationCommandAttachmentArrayOutput() InvocationCommandAttachmentArrayOutput
	ToInvocationCommandAttachmentArrayOutputWithContext(context.Context) InvocationCommandAttachmentArrayOutput
}

type InvocationCommandAttachmentArray []InvocationCommandAttachmentInput

func (InvocationCommandAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InvocationCommandAttachment)(nil)).Elem()
}

func (i InvocationCommandAttachmentArray) ToInvocationCommandAttachmentArrayOutput() InvocationCommandAttachmentArrayOutput {
	return i.ToInvocationCommandAttachmentArrayOutputWithContext(context.Background())
}

func (i InvocationCommandAttachmentArray) ToInvocationCommandAttachmentArrayOutputWithContext(ctx context.Context) InvocationCommandAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InvocationCommandAttachmentArrayOutput)
}

// InvocationCommandAttachmentMapInput is an input type that accepts InvocationCommandAttachmentMap and InvocationCommandAttachmentMapOutput values.
// You can construct a concrete instance of `InvocationCommandAttachmentMapInput` via:
//
//          InvocationCommandAttachmentMap{ "key": InvocationCommandAttachmentArgs{...} }
type InvocationCommandAttachmentMapInput interface {
	pulumi.Input

	ToInvocationCommandAttachmentMapOutput() InvocationCommandAttachmentMapOutput
	ToInvocationCommandAttachmentMapOutputWithContext(context.Context) InvocationCommandAttachmentMapOutput
}

type InvocationCommandAttachmentMap map[string]InvocationCommandAttachmentInput

func (InvocationCommandAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InvocationCommandAttachment)(nil)).Elem()
}

func (i InvocationCommandAttachmentMap) ToInvocationCommandAttachmentMapOutput() InvocationCommandAttachmentMapOutput {
	return i.ToInvocationCommandAttachmentMapOutputWithContext(context.Background())
}

func (i InvocationCommandAttachmentMap) ToInvocationCommandAttachmentMapOutputWithContext(ctx context.Context) InvocationCommandAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InvocationCommandAttachmentMapOutput)
}

type InvocationCommandAttachmentOutput struct{ *pulumi.OutputState }

func (InvocationCommandAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InvocationCommandAttachment)(nil)).Elem()
}

func (o InvocationCommandAttachmentOutput) ToInvocationCommandAttachmentOutput() InvocationCommandAttachmentOutput {
	return o
}

func (o InvocationCommandAttachmentOutput) ToInvocationCommandAttachmentOutputWithContext(ctx context.Context) InvocationCommandAttachmentOutput {
	return o
}

// Command ID.
func (o InvocationCommandAttachmentOutput) CommandId() pulumi.StringOutput {
	return o.ApplyT(func(v *InvocationCommandAttachment) pulumi.StringOutput { return v.CommandId }).(pulumi.StringOutput)
}

// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
func (o InvocationCommandAttachmentOutput) CommandName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InvocationCommandAttachment) pulumi.StringPtrOutput { return v.CommandName }).(pulumi.StringPtrOutput)
}

// Command type. SHELL and POWERSHELL are supported. The default value is SHELL.
func (o InvocationCommandAttachmentOutput) CommandType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InvocationCommandAttachment) pulumi.StringPtrOutput { return v.CommandType }).(pulumi.StringPtrOutput)
}

// Base64-encoded command. The maximum length is 64 KB.
func (o InvocationCommandAttachmentOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *InvocationCommandAttachment) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If Parameters is not provided, the default values specified here are used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
func (o InvocationCommandAttachmentOutput) DefaultParameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InvocationCommandAttachment) pulumi.StringPtrOutput { return v.DefaultParameters }).(pulumi.StringPtrOutput)
}

// Command description. The maximum length is 120 characters.
func (o InvocationCommandAttachmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InvocationCommandAttachment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether to enable the custom parameter feature.This cannot be modified once created.Default value: false.
func (o InvocationCommandAttachmentOutput) EnableParameter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InvocationCommandAttachment) pulumi.BoolPtrOutput { return v.EnableParameter }).(pulumi.BoolPtrOutput)
}

// ID of instances about to execute commands. Supported instance types:  CVM  LIGHTHOUSE.
func (o InvocationCommandAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *InvocationCommandAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The COS bucket URL for uploading logs; The URL must start with https, such as https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
func (o InvocationCommandAttachmentOutput) OutputCosBucketUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InvocationCommandAttachment) pulumi.StringPtrOutput { return v.OutputCosBucketUrl }).(pulumi.StringPtrOutput)
}

// The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive slashes.
func (o InvocationCommandAttachmentOutput) OutputCosKeyPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InvocationCommandAttachment) pulumi.StringPtrOutput { return v.OutputCosKeyPrefix }).(pulumi.StringPtrOutput)
}

// Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
func (o InvocationCommandAttachmentOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InvocationCommandAttachment) pulumi.StringPtrOutput { return v.Parameters }).(pulumi.StringPtrOutput)
}

// Whether to save the command. Valid values:rue: SaveFalse:Do not saveThe default value is False.
func (o InvocationCommandAttachmentOutput) SaveCommand() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InvocationCommandAttachment) pulumi.BoolPtrOutput { return v.SaveCommand }).(pulumi.BoolPtrOutput)
}

// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
func (o InvocationCommandAttachmentOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InvocationCommandAttachment) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root is used to execute commands on Linux and the user System is used on Windows.
func (o InvocationCommandAttachmentOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InvocationCommandAttachment) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

// Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for POWERSHELL commands.
func (o InvocationCommandAttachmentOutput) WorkingDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InvocationCommandAttachment) pulumi.StringPtrOutput { return v.WorkingDirectory }).(pulumi.StringPtrOutput)
}

type InvocationCommandAttachmentArrayOutput struct{ *pulumi.OutputState }

func (InvocationCommandAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InvocationCommandAttachment)(nil)).Elem()
}

func (o InvocationCommandAttachmentArrayOutput) ToInvocationCommandAttachmentArrayOutput() InvocationCommandAttachmentArrayOutput {
	return o
}

func (o InvocationCommandAttachmentArrayOutput) ToInvocationCommandAttachmentArrayOutputWithContext(ctx context.Context) InvocationCommandAttachmentArrayOutput {
	return o
}

func (o InvocationCommandAttachmentArrayOutput) Index(i pulumi.IntInput) InvocationCommandAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InvocationCommandAttachment {
		return vs[0].([]*InvocationCommandAttachment)[vs[1].(int)]
	}).(InvocationCommandAttachmentOutput)
}

type InvocationCommandAttachmentMapOutput struct{ *pulumi.OutputState }

func (InvocationCommandAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InvocationCommandAttachment)(nil)).Elem()
}

func (o InvocationCommandAttachmentMapOutput) ToInvocationCommandAttachmentMapOutput() InvocationCommandAttachmentMapOutput {
	return o
}

func (o InvocationCommandAttachmentMapOutput) ToInvocationCommandAttachmentMapOutputWithContext(ctx context.Context) InvocationCommandAttachmentMapOutput {
	return o
}

func (o InvocationCommandAttachmentMapOutput) MapIndex(k pulumi.StringInput) InvocationCommandAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InvocationCommandAttachment {
		return vs[0].(map[string]*InvocationCommandAttachment)[vs[1].(string)]
	}).(InvocationCommandAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InvocationCommandAttachmentInput)(nil)).Elem(), &InvocationCommandAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*InvocationCommandAttachmentArrayInput)(nil)).Elem(), InvocationCommandAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InvocationCommandAttachmentMapInput)(nil)).Elem(), InvocationCommandAttachmentMap{})
	pulumi.RegisterOutputType(InvocationCommandAttachmentOutput{})
	pulumi.RegisterOutputType(InvocationCommandAttachmentArrayOutput{})
	pulumi.RegisterOutputType(InvocationCommandAttachmentMapOutput{})
}
