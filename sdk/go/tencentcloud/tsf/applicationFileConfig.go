// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a tsf applicationFileConfig
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tsf.NewApplicationFileConfig(ctx, "applicationFileConfig", &Tsf.ApplicationFileConfigArgs{
//				ApplicationId:     pulumi.String("application-a24x29xv"),
//				ConfigFileCode:    pulumi.String("UTF-8"),
//				ConfigFileName:    pulumi.String("application.yaml"),
//				ConfigFilePath:    pulumi.String("/etc/nginx"),
//				ConfigFileValue:   pulumi.String("test: 1"),
//				ConfigName:        pulumi.String("terraform-test"),
//				ConfigPostCmd:     pulumi.String("source .bashrc"),
//				ConfigVersion:     pulumi.String("1.0"),
//				ConfigVersionDesc: pulumi.String("1.0"),
//				EncodeWithBase64:  pulumi.Bool(true),
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
type ApplicationFileConfig struct {
	pulumi.CustomResourceState

	// Config file associated application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Configuration file encoding, utf-8 or gbk. Note: If you choose gbk, you need the support of a new version of tsf-consul-template (public cloud virtual machines need to use 1.32 tsf-agent, and containers need to obtain the latest tsf-consul-template-docker.tar.gz from the documentation).
	ConfigFileCode pulumi.StringPtrOutput `pulumi:"configFileCode"`
	// Config file name.
	ConfigFileName pulumi.StringOutput `pulumi:"configFileName"`
	// config release path.
	ConfigFilePath pulumi.StringOutput `pulumi:"configFilePath"`
	// Configuration file content (the original content encoding needs to be in utf-8 format, if the ConfigFileCode is gbk, it will be converted in the background).
	ConfigFileValue pulumi.StringOutput `pulumi:"configFileValue"`
	// Config Name.
	ConfigName pulumi.StringOutput `pulumi:"configName"`
	// post command.
	ConfigPostCmd pulumi.StringPtrOutput `pulumi:"configPostCmd"`
	// Config version.
	ConfigVersion pulumi.StringOutput `pulumi:"configVersion"`
	// config version description.
	ConfigVersionDesc pulumi.StringPtrOutput `pulumi:"configVersionDesc"`
	// the config value is encoded with base64 or not.
	EncodeWithBase64 pulumi.BoolPtrOutput `pulumi:"encodeWithBase64"`
	// datasource for auth.
	ProgramIdLists pulumi.StringArrayOutput `pulumi:"programIdLists"`
}

// NewApplicationFileConfig registers a new resource with the given unique name, arguments, and options.
func NewApplicationFileConfig(ctx *pulumi.Context,
	name string, args *ApplicationFileConfigArgs, opts ...pulumi.ResourceOption) (*ApplicationFileConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.ConfigFileName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigFileName'")
	}
	if args.ConfigFilePath == nil {
		return nil, errors.New("invalid value for required argument 'ConfigFilePath'")
	}
	if args.ConfigFileValue == nil {
		return nil, errors.New("invalid value for required argument 'ConfigFileValue'")
	}
	if args.ConfigName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigName'")
	}
	if args.ConfigVersion == nil {
		return nil, errors.New("invalid value for required argument 'ConfigVersion'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationFileConfig
	err := ctx.RegisterResource("tencentcloud:Tsf/applicationFileConfig:ApplicationFileConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationFileConfig gets an existing ApplicationFileConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationFileConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationFileConfigState, opts ...pulumi.ResourceOption) (*ApplicationFileConfig, error) {
	var resource ApplicationFileConfig
	err := ctx.ReadResource("tencentcloud:Tsf/applicationFileConfig:ApplicationFileConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationFileConfig resources.
type applicationFileConfigState struct {
	// Config file associated application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// Configuration file encoding, utf-8 or gbk. Note: If you choose gbk, you need the support of a new version of tsf-consul-template (public cloud virtual machines need to use 1.32 tsf-agent, and containers need to obtain the latest tsf-consul-template-docker.tar.gz from the documentation).
	ConfigFileCode *string `pulumi:"configFileCode"`
	// Config file name.
	ConfigFileName *string `pulumi:"configFileName"`
	// config release path.
	ConfigFilePath *string `pulumi:"configFilePath"`
	// Configuration file content (the original content encoding needs to be in utf-8 format, if the ConfigFileCode is gbk, it will be converted in the background).
	ConfigFileValue *string `pulumi:"configFileValue"`
	// Config Name.
	ConfigName *string `pulumi:"configName"`
	// post command.
	ConfigPostCmd *string `pulumi:"configPostCmd"`
	// Config version.
	ConfigVersion *string `pulumi:"configVersion"`
	// config version description.
	ConfigVersionDesc *string `pulumi:"configVersionDesc"`
	// the config value is encoded with base64 or not.
	EncodeWithBase64 *bool `pulumi:"encodeWithBase64"`
	// datasource for auth.
	ProgramIdLists []string `pulumi:"programIdLists"`
}

type ApplicationFileConfigState struct {
	// Config file associated application ID.
	ApplicationId pulumi.StringPtrInput
	// Configuration file encoding, utf-8 or gbk. Note: If you choose gbk, you need the support of a new version of tsf-consul-template (public cloud virtual machines need to use 1.32 tsf-agent, and containers need to obtain the latest tsf-consul-template-docker.tar.gz from the documentation).
	ConfigFileCode pulumi.StringPtrInput
	// Config file name.
	ConfigFileName pulumi.StringPtrInput
	// config release path.
	ConfigFilePath pulumi.StringPtrInput
	// Configuration file content (the original content encoding needs to be in utf-8 format, if the ConfigFileCode is gbk, it will be converted in the background).
	ConfigFileValue pulumi.StringPtrInput
	// Config Name.
	ConfigName pulumi.StringPtrInput
	// post command.
	ConfigPostCmd pulumi.StringPtrInput
	// Config version.
	ConfigVersion pulumi.StringPtrInput
	// config version description.
	ConfigVersionDesc pulumi.StringPtrInput
	// the config value is encoded with base64 or not.
	EncodeWithBase64 pulumi.BoolPtrInput
	// datasource for auth.
	ProgramIdLists pulumi.StringArrayInput
}

func (ApplicationFileConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationFileConfigState)(nil)).Elem()
}

type applicationFileConfigArgs struct {
	// Config file associated application ID.
	ApplicationId string `pulumi:"applicationId"`
	// Configuration file encoding, utf-8 or gbk. Note: If you choose gbk, you need the support of a new version of tsf-consul-template (public cloud virtual machines need to use 1.32 tsf-agent, and containers need to obtain the latest tsf-consul-template-docker.tar.gz from the documentation).
	ConfigFileCode *string `pulumi:"configFileCode"`
	// Config file name.
	ConfigFileName string `pulumi:"configFileName"`
	// config release path.
	ConfigFilePath string `pulumi:"configFilePath"`
	// Configuration file content (the original content encoding needs to be in utf-8 format, if the ConfigFileCode is gbk, it will be converted in the background).
	ConfigFileValue string `pulumi:"configFileValue"`
	// Config Name.
	ConfigName string `pulumi:"configName"`
	// post command.
	ConfigPostCmd *string `pulumi:"configPostCmd"`
	// Config version.
	ConfigVersion string `pulumi:"configVersion"`
	// config version description.
	ConfigVersionDesc *string `pulumi:"configVersionDesc"`
	// the config value is encoded with base64 or not.
	EncodeWithBase64 *bool `pulumi:"encodeWithBase64"`
	// datasource for auth.
	ProgramIdLists []string `pulumi:"programIdLists"`
}

// The set of arguments for constructing a ApplicationFileConfig resource.
type ApplicationFileConfigArgs struct {
	// Config file associated application ID.
	ApplicationId pulumi.StringInput
	// Configuration file encoding, utf-8 or gbk. Note: If you choose gbk, you need the support of a new version of tsf-consul-template (public cloud virtual machines need to use 1.32 tsf-agent, and containers need to obtain the latest tsf-consul-template-docker.tar.gz from the documentation).
	ConfigFileCode pulumi.StringPtrInput
	// Config file name.
	ConfigFileName pulumi.StringInput
	// config release path.
	ConfigFilePath pulumi.StringInput
	// Configuration file content (the original content encoding needs to be in utf-8 format, if the ConfigFileCode is gbk, it will be converted in the background).
	ConfigFileValue pulumi.StringInput
	// Config Name.
	ConfigName pulumi.StringInput
	// post command.
	ConfigPostCmd pulumi.StringPtrInput
	// Config version.
	ConfigVersion pulumi.StringInput
	// config version description.
	ConfigVersionDesc pulumi.StringPtrInput
	// the config value is encoded with base64 or not.
	EncodeWithBase64 pulumi.BoolPtrInput
	// datasource for auth.
	ProgramIdLists pulumi.StringArrayInput
}

func (ApplicationFileConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationFileConfigArgs)(nil)).Elem()
}

type ApplicationFileConfigInput interface {
	pulumi.Input

	ToApplicationFileConfigOutput() ApplicationFileConfigOutput
	ToApplicationFileConfigOutputWithContext(ctx context.Context) ApplicationFileConfigOutput
}

func (*ApplicationFileConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationFileConfig)(nil)).Elem()
}

func (i *ApplicationFileConfig) ToApplicationFileConfigOutput() ApplicationFileConfigOutput {
	return i.ToApplicationFileConfigOutputWithContext(context.Background())
}

func (i *ApplicationFileConfig) ToApplicationFileConfigOutputWithContext(ctx context.Context) ApplicationFileConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationFileConfigOutput)
}

// ApplicationFileConfigArrayInput is an input type that accepts ApplicationFileConfigArray and ApplicationFileConfigArrayOutput values.
// You can construct a concrete instance of `ApplicationFileConfigArrayInput` via:
//
//	ApplicationFileConfigArray{ ApplicationFileConfigArgs{...} }
type ApplicationFileConfigArrayInput interface {
	pulumi.Input

	ToApplicationFileConfigArrayOutput() ApplicationFileConfigArrayOutput
	ToApplicationFileConfigArrayOutputWithContext(context.Context) ApplicationFileConfigArrayOutput
}

type ApplicationFileConfigArray []ApplicationFileConfigInput

func (ApplicationFileConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationFileConfig)(nil)).Elem()
}

func (i ApplicationFileConfigArray) ToApplicationFileConfigArrayOutput() ApplicationFileConfigArrayOutput {
	return i.ToApplicationFileConfigArrayOutputWithContext(context.Background())
}

func (i ApplicationFileConfigArray) ToApplicationFileConfigArrayOutputWithContext(ctx context.Context) ApplicationFileConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationFileConfigArrayOutput)
}

// ApplicationFileConfigMapInput is an input type that accepts ApplicationFileConfigMap and ApplicationFileConfigMapOutput values.
// You can construct a concrete instance of `ApplicationFileConfigMapInput` via:
//
//	ApplicationFileConfigMap{ "key": ApplicationFileConfigArgs{...} }
type ApplicationFileConfigMapInput interface {
	pulumi.Input

	ToApplicationFileConfigMapOutput() ApplicationFileConfigMapOutput
	ToApplicationFileConfigMapOutputWithContext(context.Context) ApplicationFileConfigMapOutput
}

type ApplicationFileConfigMap map[string]ApplicationFileConfigInput

func (ApplicationFileConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationFileConfig)(nil)).Elem()
}

func (i ApplicationFileConfigMap) ToApplicationFileConfigMapOutput() ApplicationFileConfigMapOutput {
	return i.ToApplicationFileConfigMapOutputWithContext(context.Background())
}

func (i ApplicationFileConfigMap) ToApplicationFileConfigMapOutputWithContext(ctx context.Context) ApplicationFileConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationFileConfigMapOutput)
}

type ApplicationFileConfigOutput struct{ *pulumi.OutputState }

func (ApplicationFileConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationFileConfig)(nil)).Elem()
}

func (o ApplicationFileConfigOutput) ToApplicationFileConfigOutput() ApplicationFileConfigOutput {
	return o
}

func (o ApplicationFileConfigOutput) ToApplicationFileConfigOutputWithContext(ctx context.Context) ApplicationFileConfigOutput {
	return o
}

// Config file associated application ID.
func (o ApplicationFileConfigOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationFileConfig) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// Configuration file encoding, utf-8 or gbk. Note: If you choose gbk, you need the support of a new version of tsf-consul-template (public cloud virtual machines need to use 1.32 tsf-agent, and containers need to obtain the latest tsf-consul-template-docker.tar.gz from the documentation).
func (o ApplicationFileConfigOutput) ConfigFileCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationFileConfig) pulumi.StringPtrOutput { return v.ConfigFileCode }).(pulumi.StringPtrOutput)
}

// Config file name.
func (o ApplicationFileConfigOutput) ConfigFileName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationFileConfig) pulumi.StringOutput { return v.ConfigFileName }).(pulumi.StringOutput)
}

// config release path.
func (o ApplicationFileConfigOutput) ConfigFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationFileConfig) pulumi.StringOutput { return v.ConfigFilePath }).(pulumi.StringOutput)
}

// Configuration file content (the original content encoding needs to be in utf-8 format, if the ConfigFileCode is gbk, it will be converted in the background).
func (o ApplicationFileConfigOutput) ConfigFileValue() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationFileConfig) pulumi.StringOutput { return v.ConfigFileValue }).(pulumi.StringOutput)
}

// Config Name.
func (o ApplicationFileConfigOutput) ConfigName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationFileConfig) pulumi.StringOutput { return v.ConfigName }).(pulumi.StringOutput)
}

// post command.
func (o ApplicationFileConfigOutput) ConfigPostCmd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationFileConfig) pulumi.StringPtrOutput { return v.ConfigPostCmd }).(pulumi.StringPtrOutput)
}

// Config version.
func (o ApplicationFileConfigOutput) ConfigVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationFileConfig) pulumi.StringOutput { return v.ConfigVersion }).(pulumi.StringOutput)
}

// config version description.
func (o ApplicationFileConfigOutput) ConfigVersionDesc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationFileConfig) pulumi.StringPtrOutput { return v.ConfigVersionDesc }).(pulumi.StringPtrOutput)
}

// the config value is encoded with base64 or not.
func (o ApplicationFileConfigOutput) EncodeWithBase64() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationFileConfig) pulumi.BoolPtrOutput { return v.EncodeWithBase64 }).(pulumi.BoolPtrOutput)
}

// datasource for auth.
func (o ApplicationFileConfigOutput) ProgramIdLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationFileConfig) pulumi.StringArrayOutput { return v.ProgramIdLists }).(pulumi.StringArrayOutput)
}

type ApplicationFileConfigArrayOutput struct{ *pulumi.OutputState }

func (ApplicationFileConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationFileConfig)(nil)).Elem()
}

func (o ApplicationFileConfigArrayOutput) ToApplicationFileConfigArrayOutput() ApplicationFileConfigArrayOutput {
	return o
}

func (o ApplicationFileConfigArrayOutput) ToApplicationFileConfigArrayOutputWithContext(ctx context.Context) ApplicationFileConfigArrayOutput {
	return o
}

func (o ApplicationFileConfigArrayOutput) Index(i pulumi.IntInput) ApplicationFileConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationFileConfig {
		return vs[0].([]*ApplicationFileConfig)[vs[1].(int)]
	}).(ApplicationFileConfigOutput)
}

type ApplicationFileConfigMapOutput struct{ *pulumi.OutputState }

func (ApplicationFileConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationFileConfig)(nil)).Elem()
}

func (o ApplicationFileConfigMapOutput) ToApplicationFileConfigMapOutput() ApplicationFileConfigMapOutput {
	return o
}

func (o ApplicationFileConfigMapOutput) ToApplicationFileConfigMapOutputWithContext(ctx context.Context) ApplicationFileConfigMapOutput {
	return o
}

func (o ApplicationFileConfigMapOutput) MapIndex(k pulumi.StringInput) ApplicationFileConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationFileConfig {
		return vs[0].(map[string]*ApplicationFileConfig)[vs[1].(string)]
	}).(ApplicationFileConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationFileConfigInput)(nil)).Elem(), &ApplicationFileConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationFileConfigArrayInput)(nil)).Elem(), ApplicationFileConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationFileConfigMapInput)(nil)).Elem(), ApplicationFileConfigMap{})
	pulumi.RegisterOutputType(ApplicationFileConfigOutput{})
	pulumi.RegisterOutputType(ApplicationFileConfigArrayOutput{})
	pulumi.RegisterOutputType(ApplicationFileConfigMapOutput{})
}
