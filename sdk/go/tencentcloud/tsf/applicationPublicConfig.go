// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tsf applicationPublicConfig
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Tsf.NewApplicationPublicConfig(ctx, "applicationPublicConfig", &Tsf.ApplicationPublicConfigArgs{
// 			ConfigName:        pulumi.String("my_config"),
// 			ConfigType:        pulumi.String("P"),
// 			ConfigValue:       pulumi.String("test: 1"),
// 			ConfigVersion:     pulumi.String("1.0"),
// 			ConfigVersionDesc: pulumi.String("product version"),
// 			EncodeWithBase64:  pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ApplicationPublicConfig struct {
	pulumi.CustomResourceState

	// Config Name.
	ConfigName pulumi.StringOutput `pulumi:"configName"`
	// Config type.
	ConfigType pulumi.StringPtrOutput `pulumi:"configType"`
	// config value, only yaml file allowed.
	ConfigValue pulumi.StringOutput `pulumi:"configValue"`
	// config version.
	ConfigVersion pulumi.StringOutput `pulumi:"configVersion"`
	// Config version description.
	ConfigVersionDesc pulumi.StringPtrOutput `pulumi:"configVersionDesc"`
	// the config value is encoded with base64 or not.
	EncodeWithBase64 pulumi.BoolPtrOutput `pulumi:"encodeWithBase64"`
	// datasource for auth.
	ProgramIdLists pulumi.StringArrayOutput `pulumi:"programIdLists"`
}

// NewApplicationPublicConfig registers a new resource with the given unique name, arguments, and options.
func NewApplicationPublicConfig(ctx *pulumi.Context,
	name string, args *ApplicationPublicConfigArgs, opts ...pulumi.ResourceOption) (*ApplicationPublicConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigName'")
	}
	if args.ConfigValue == nil {
		return nil, errors.New("invalid value for required argument 'ConfigValue'")
	}
	if args.ConfigVersion == nil {
		return nil, errors.New("invalid value for required argument 'ConfigVersion'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ApplicationPublicConfig
	err := ctx.RegisterResource("tencentcloud:Tsf/applicationPublicConfig:ApplicationPublicConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationPublicConfig gets an existing ApplicationPublicConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationPublicConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationPublicConfigState, opts ...pulumi.ResourceOption) (*ApplicationPublicConfig, error) {
	var resource ApplicationPublicConfig
	err := ctx.ReadResource("tencentcloud:Tsf/applicationPublicConfig:ApplicationPublicConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationPublicConfig resources.
type applicationPublicConfigState struct {
	// Config Name.
	ConfigName *string `pulumi:"configName"`
	// Config type.
	ConfigType *string `pulumi:"configType"`
	// config value, only yaml file allowed.
	ConfigValue *string `pulumi:"configValue"`
	// config version.
	ConfigVersion *string `pulumi:"configVersion"`
	// Config version description.
	ConfigVersionDesc *string `pulumi:"configVersionDesc"`
	// the config value is encoded with base64 or not.
	EncodeWithBase64 *bool `pulumi:"encodeWithBase64"`
	// datasource for auth.
	ProgramIdLists []string `pulumi:"programIdLists"`
}

type ApplicationPublicConfigState struct {
	// Config Name.
	ConfigName pulumi.StringPtrInput
	// Config type.
	ConfigType pulumi.StringPtrInput
	// config value, only yaml file allowed.
	ConfigValue pulumi.StringPtrInput
	// config version.
	ConfigVersion pulumi.StringPtrInput
	// Config version description.
	ConfigVersionDesc pulumi.StringPtrInput
	// the config value is encoded with base64 or not.
	EncodeWithBase64 pulumi.BoolPtrInput
	// datasource for auth.
	ProgramIdLists pulumi.StringArrayInput
}

func (ApplicationPublicConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationPublicConfigState)(nil)).Elem()
}

type applicationPublicConfigArgs struct {
	// Config Name.
	ConfigName string `pulumi:"configName"`
	// Config type.
	ConfigType *string `pulumi:"configType"`
	// config value, only yaml file allowed.
	ConfigValue string `pulumi:"configValue"`
	// config version.
	ConfigVersion string `pulumi:"configVersion"`
	// Config version description.
	ConfigVersionDesc *string `pulumi:"configVersionDesc"`
	// the config value is encoded with base64 or not.
	EncodeWithBase64 *bool `pulumi:"encodeWithBase64"`
	// datasource for auth.
	ProgramIdLists []string `pulumi:"programIdLists"`
}

// The set of arguments for constructing a ApplicationPublicConfig resource.
type ApplicationPublicConfigArgs struct {
	// Config Name.
	ConfigName pulumi.StringInput
	// Config type.
	ConfigType pulumi.StringPtrInput
	// config value, only yaml file allowed.
	ConfigValue pulumi.StringInput
	// config version.
	ConfigVersion pulumi.StringInput
	// Config version description.
	ConfigVersionDesc pulumi.StringPtrInput
	// the config value is encoded with base64 or not.
	EncodeWithBase64 pulumi.BoolPtrInput
	// datasource for auth.
	ProgramIdLists pulumi.StringArrayInput
}

func (ApplicationPublicConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationPublicConfigArgs)(nil)).Elem()
}

type ApplicationPublicConfigInput interface {
	pulumi.Input

	ToApplicationPublicConfigOutput() ApplicationPublicConfigOutput
	ToApplicationPublicConfigOutputWithContext(ctx context.Context) ApplicationPublicConfigOutput
}

func (*ApplicationPublicConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPublicConfig)(nil)).Elem()
}

func (i *ApplicationPublicConfig) ToApplicationPublicConfigOutput() ApplicationPublicConfigOutput {
	return i.ToApplicationPublicConfigOutputWithContext(context.Background())
}

func (i *ApplicationPublicConfig) ToApplicationPublicConfigOutputWithContext(ctx context.Context) ApplicationPublicConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPublicConfigOutput)
}

// ApplicationPublicConfigArrayInput is an input type that accepts ApplicationPublicConfigArray and ApplicationPublicConfigArrayOutput values.
// You can construct a concrete instance of `ApplicationPublicConfigArrayInput` via:
//
//          ApplicationPublicConfigArray{ ApplicationPublicConfigArgs{...} }
type ApplicationPublicConfigArrayInput interface {
	pulumi.Input

	ToApplicationPublicConfigArrayOutput() ApplicationPublicConfigArrayOutput
	ToApplicationPublicConfigArrayOutputWithContext(context.Context) ApplicationPublicConfigArrayOutput
}

type ApplicationPublicConfigArray []ApplicationPublicConfigInput

func (ApplicationPublicConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationPublicConfig)(nil)).Elem()
}

func (i ApplicationPublicConfigArray) ToApplicationPublicConfigArrayOutput() ApplicationPublicConfigArrayOutput {
	return i.ToApplicationPublicConfigArrayOutputWithContext(context.Background())
}

func (i ApplicationPublicConfigArray) ToApplicationPublicConfigArrayOutputWithContext(ctx context.Context) ApplicationPublicConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPublicConfigArrayOutput)
}

// ApplicationPublicConfigMapInput is an input type that accepts ApplicationPublicConfigMap and ApplicationPublicConfigMapOutput values.
// You can construct a concrete instance of `ApplicationPublicConfigMapInput` via:
//
//          ApplicationPublicConfigMap{ "key": ApplicationPublicConfigArgs{...} }
type ApplicationPublicConfigMapInput interface {
	pulumi.Input

	ToApplicationPublicConfigMapOutput() ApplicationPublicConfigMapOutput
	ToApplicationPublicConfigMapOutputWithContext(context.Context) ApplicationPublicConfigMapOutput
}

type ApplicationPublicConfigMap map[string]ApplicationPublicConfigInput

func (ApplicationPublicConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationPublicConfig)(nil)).Elem()
}

func (i ApplicationPublicConfigMap) ToApplicationPublicConfigMapOutput() ApplicationPublicConfigMapOutput {
	return i.ToApplicationPublicConfigMapOutputWithContext(context.Background())
}

func (i ApplicationPublicConfigMap) ToApplicationPublicConfigMapOutputWithContext(ctx context.Context) ApplicationPublicConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPublicConfigMapOutput)
}

type ApplicationPublicConfigOutput struct{ *pulumi.OutputState }

func (ApplicationPublicConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPublicConfig)(nil)).Elem()
}

func (o ApplicationPublicConfigOutput) ToApplicationPublicConfigOutput() ApplicationPublicConfigOutput {
	return o
}

func (o ApplicationPublicConfigOutput) ToApplicationPublicConfigOutputWithContext(ctx context.Context) ApplicationPublicConfigOutput {
	return o
}

// Config Name.
func (o ApplicationPublicConfigOutput) ConfigName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPublicConfig) pulumi.StringOutput { return v.ConfigName }).(pulumi.StringOutput)
}

// Config type.
func (o ApplicationPublicConfigOutput) ConfigType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPublicConfig) pulumi.StringPtrOutput { return v.ConfigType }).(pulumi.StringPtrOutput)
}

// config value, only yaml file allowed.
func (o ApplicationPublicConfigOutput) ConfigValue() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPublicConfig) pulumi.StringOutput { return v.ConfigValue }).(pulumi.StringOutput)
}

// config version.
func (o ApplicationPublicConfigOutput) ConfigVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPublicConfig) pulumi.StringOutput { return v.ConfigVersion }).(pulumi.StringOutput)
}

// Config version description.
func (o ApplicationPublicConfigOutput) ConfigVersionDesc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPublicConfig) pulumi.StringPtrOutput { return v.ConfigVersionDesc }).(pulumi.StringPtrOutput)
}

// the config value is encoded with base64 or not.
func (o ApplicationPublicConfigOutput) EncodeWithBase64() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationPublicConfig) pulumi.BoolPtrOutput { return v.EncodeWithBase64 }).(pulumi.BoolPtrOutput)
}

// datasource for auth.
func (o ApplicationPublicConfigOutput) ProgramIdLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationPublicConfig) pulumi.StringArrayOutput { return v.ProgramIdLists }).(pulumi.StringArrayOutput)
}

type ApplicationPublicConfigArrayOutput struct{ *pulumi.OutputState }

func (ApplicationPublicConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationPublicConfig)(nil)).Elem()
}

func (o ApplicationPublicConfigArrayOutput) ToApplicationPublicConfigArrayOutput() ApplicationPublicConfigArrayOutput {
	return o
}

func (o ApplicationPublicConfigArrayOutput) ToApplicationPublicConfigArrayOutputWithContext(ctx context.Context) ApplicationPublicConfigArrayOutput {
	return o
}

func (o ApplicationPublicConfigArrayOutput) Index(i pulumi.IntInput) ApplicationPublicConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationPublicConfig {
		return vs[0].([]*ApplicationPublicConfig)[vs[1].(int)]
	}).(ApplicationPublicConfigOutput)
}

type ApplicationPublicConfigMapOutput struct{ *pulumi.OutputState }

func (ApplicationPublicConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationPublicConfig)(nil)).Elem()
}

func (o ApplicationPublicConfigMapOutput) ToApplicationPublicConfigMapOutput() ApplicationPublicConfigMapOutput {
	return o
}

func (o ApplicationPublicConfigMapOutput) ToApplicationPublicConfigMapOutputWithContext(ctx context.Context) ApplicationPublicConfigMapOutput {
	return o
}

func (o ApplicationPublicConfigMapOutput) MapIndex(k pulumi.StringInput) ApplicationPublicConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationPublicConfig {
		return vs[0].(map[string]*ApplicationPublicConfig)[vs[1].(string)]
	}).(ApplicationPublicConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPublicConfigInput)(nil)).Elem(), &ApplicationPublicConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPublicConfigArrayInput)(nil)).Elem(), ApplicationPublicConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPublicConfigMapInput)(nil)).Elem(), ApplicationPublicConfigMap{})
	pulumi.RegisterOutputType(ApplicationPublicConfigOutput{})
	pulumi.RegisterOutputType(ApplicationPublicConfigArrayOutput{})
	pulumi.RegisterOutputType(ApplicationPublicConfigMapOutput{})
}
