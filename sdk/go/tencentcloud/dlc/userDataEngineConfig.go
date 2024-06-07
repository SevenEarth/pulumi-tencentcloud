// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dlc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a dlc userDataEngineConfig
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dlc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dlc.NewUserDataEngineConfig(ctx, "userDataEngineConfig", &Dlc.UserDataEngineConfigArgs{
//				DataEngineConfigPairs: dlc.UserDataEngineConfigDataEngineConfigPairArray{
//					&dlc.UserDataEngineConfigDataEngineConfigPairArgs{
//						ConfigItem:  pulumi.String("qq"),
//						ConfigValue: pulumi.String("ff"),
//					},
//				},
//				DataEngineId: pulumi.String("DataEngine-cgkvbas6"),
//				SessionResourceTemplate: &dlc.UserDataEngineConfigSessionResourceTemplateArgs{
//					DriverSize:         pulumi.String("small"),
//					ExecutorMaxNumbers: pulumi.Int(1),
//					ExecutorNums:       pulumi.Int(1),
//					ExecutorSize:       pulumi.String("small"),
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
// dlc user_data_engine_config can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Dlc/userDataEngineConfig:UserDataEngineConfig user_data_engine_config user_data_engine_config_id
// ```
type UserDataEngineConfig struct {
	pulumi.CustomResourceState

	// Engine configuration items.
	DataEngineConfigPairs UserDataEngineConfigDataEngineConfigPairArrayOutput `pulumi:"dataEngineConfigPairs"`
	// Engine unique id.
	DataEngineId pulumi.StringOutput `pulumi:"dataEngineId"`
	// Job engine resource configuration template.
	SessionResourceTemplate UserDataEngineConfigSessionResourceTemplatePtrOutput `pulumi:"sessionResourceTemplate"`
}

// NewUserDataEngineConfig registers a new resource with the given unique name, arguments, and options.
func NewUserDataEngineConfig(ctx *pulumi.Context,
	name string, args *UserDataEngineConfigArgs, opts ...pulumi.ResourceOption) (*UserDataEngineConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataEngineId == nil {
		return nil, errors.New("invalid value for required argument 'DataEngineId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserDataEngineConfig
	err := ctx.RegisterResource("tencentcloud:Dlc/userDataEngineConfig:UserDataEngineConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserDataEngineConfig gets an existing UserDataEngineConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserDataEngineConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserDataEngineConfigState, opts ...pulumi.ResourceOption) (*UserDataEngineConfig, error) {
	var resource UserDataEngineConfig
	err := ctx.ReadResource("tencentcloud:Dlc/userDataEngineConfig:UserDataEngineConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserDataEngineConfig resources.
type userDataEngineConfigState struct {
	// Engine configuration items.
	DataEngineConfigPairs []UserDataEngineConfigDataEngineConfigPair `pulumi:"dataEngineConfigPairs"`
	// Engine unique id.
	DataEngineId *string `pulumi:"dataEngineId"`
	// Job engine resource configuration template.
	SessionResourceTemplate *UserDataEngineConfigSessionResourceTemplate `pulumi:"sessionResourceTemplate"`
}

type UserDataEngineConfigState struct {
	// Engine configuration items.
	DataEngineConfigPairs UserDataEngineConfigDataEngineConfigPairArrayInput
	// Engine unique id.
	DataEngineId pulumi.StringPtrInput
	// Job engine resource configuration template.
	SessionResourceTemplate UserDataEngineConfigSessionResourceTemplatePtrInput
}

func (UserDataEngineConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*userDataEngineConfigState)(nil)).Elem()
}

type userDataEngineConfigArgs struct {
	// Engine configuration items.
	DataEngineConfigPairs []UserDataEngineConfigDataEngineConfigPair `pulumi:"dataEngineConfigPairs"`
	// Engine unique id.
	DataEngineId string `pulumi:"dataEngineId"`
	// Job engine resource configuration template.
	SessionResourceTemplate *UserDataEngineConfigSessionResourceTemplate `pulumi:"sessionResourceTemplate"`
}

// The set of arguments for constructing a UserDataEngineConfig resource.
type UserDataEngineConfigArgs struct {
	// Engine configuration items.
	DataEngineConfigPairs UserDataEngineConfigDataEngineConfigPairArrayInput
	// Engine unique id.
	DataEngineId pulumi.StringInput
	// Job engine resource configuration template.
	SessionResourceTemplate UserDataEngineConfigSessionResourceTemplatePtrInput
}

func (UserDataEngineConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userDataEngineConfigArgs)(nil)).Elem()
}

type UserDataEngineConfigInput interface {
	pulumi.Input

	ToUserDataEngineConfigOutput() UserDataEngineConfigOutput
	ToUserDataEngineConfigOutputWithContext(ctx context.Context) UserDataEngineConfigOutput
}

func (*UserDataEngineConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDataEngineConfig)(nil)).Elem()
}

func (i *UserDataEngineConfig) ToUserDataEngineConfigOutput() UserDataEngineConfigOutput {
	return i.ToUserDataEngineConfigOutputWithContext(context.Background())
}

func (i *UserDataEngineConfig) ToUserDataEngineConfigOutputWithContext(ctx context.Context) UserDataEngineConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDataEngineConfigOutput)
}

// UserDataEngineConfigArrayInput is an input type that accepts UserDataEngineConfigArray and UserDataEngineConfigArrayOutput values.
// You can construct a concrete instance of `UserDataEngineConfigArrayInput` via:
//
//	UserDataEngineConfigArray{ UserDataEngineConfigArgs{...} }
type UserDataEngineConfigArrayInput interface {
	pulumi.Input

	ToUserDataEngineConfigArrayOutput() UserDataEngineConfigArrayOutput
	ToUserDataEngineConfigArrayOutputWithContext(context.Context) UserDataEngineConfigArrayOutput
}

type UserDataEngineConfigArray []UserDataEngineConfigInput

func (UserDataEngineConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserDataEngineConfig)(nil)).Elem()
}

func (i UserDataEngineConfigArray) ToUserDataEngineConfigArrayOutput() UserDataEngineConfigArrayOutput {
	return i.ToUserDataEngineConfigArrayOutputWithContext(context.Background())
}

func (i UserDataEngineConfigArray) ToUserDataEngineConfigArrayOutputWithContext(ctx context.Context) UserDataEngineConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDataEngineConfigArrayOutput)
}

// UserDataEngineConfigMapInput is an input type that accepts UserDataEngineConfigMap and UserDataEngineConfigMapOutput values.
// You can construct a concrete instance of `UserDataEngineConfigMapInput` via:
//
//	UserDataEngineConfigMap{ "key": UserDataEngineConfigArgs{...} }
type UserDataEngineConfigMapInput interface {
	pulumi.Input

	ToUserDataEngineConfigMapOutput() UserDataEngineConfigMapOutput
	ToUserDataEngineConfigMapOutputWithContext(context.Context) UserDataEngineConfigMapOutput
}

type UserDataEngineConfigMap map[string]UserDataEngineConfigInput

func (UserDataEngineConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserDataEngineConfig)(nil)).Elem()
}

func (i UserDataEngineConfigMap) ToUserDataEngineConfigMapOutput() UserDataEngineConfigMapOutput {
	return i.ToUserDataEngineConfigMapOutputWithContext(context.Background())
}

func (i UserDataEngineConfigMap) ToUserDataEngineConfigMapOutputWithContext(ctx context.Context) UserDataEngineConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDataEngineConfigMapOutput)
}

type UserDataEngineConfigOutput struct{ *pulumi.OutputState }

func (UserDataEngineConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDataEngineConfig)(nil)).Elem()
}

func (o UserDataEngineConfigOutput) ToUserDataEngineConfigOutput() UserDataEngineConfigOutput {
	return o
}

func (o UserDataEngineConfigOutput) ToUserDataEngineConfigOutputWithContext(ctx context.Context) UserDataEngineConfigOutput {
	return o
}

// Engine configuration items.
func (o UserDataEngineConfigOutput) DataEngineConfigPairs() UserDataEngineConfigDataEngineConfigPairArrayOutput {
	return o.ApplyT(func(v *UserDataEngineConfig) UserDataEngineConfigDataEngineConfigPairArrayOutput {
		return v.DataEngineConfigPairs
	}).(UserDataEngineConfigDataEngineConfigPairArrayOutput)
}

// Engine unique id.
func (o UserDataEngineConfigOutput) DataEngineId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDataEngineConfig) pulumi.StringOutput { return v.DataEngineId }).(pulumi.StringOutput)
}

// Job engine resource configuration template.
func (o UserDataEngineConfigOutput) SessionResourceTemplate() UserDataEngineConfigSessionResourceTemplatePtrOutput {
	return o.ApplyT(func(v *UserDataEngineConfig) UserDataEngineConfigSessionResourceTemplatePtrOutput {
		return v.SessionResourceTemplate
	}).(UserDataEngineConfigSessionResourceTemplatePtrOutput)
}

type UserDataEngineConfigArrayOutput struct{ *pulumi.OutputState }

func (UserDataEngineConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserDataEngineConfig)(nil)).Elem()
}

func (o UserDataEngineConfigArrayOutput) ToUserDataEngineConfigArrayOutput() UserDataEngineConfigArrayOutput {
	return o
}

func (o UserDataEngineConfigArrayOutput) ToUserDataEngineConfigArrayOutputWithContext(ctx context.Context) UserDataEngineConfigArrayOutput {
	return o
}

func (o UserDataEngineConfigArrayOutput) Index(i pulumi.IntInput) UserDataEngineConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserDataEngineConfig {
		return vs[0].([]*UserDataEngineConfig)[vs[1].(int)]
	}).(UserDataEngineConfigOutput)
}

type UserDataEngineConfigMapOutput struct{ *pulumi.OutputState }

func (UserDataEngineConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserDataEngineConfig)(nil)).Elem()
}

func (o UserDataEngineConfigMapOutput) ToUserDataEngineConfigMapOutput() UserDataEngineConfigMapOutput {
	return o
}

func (o UserDataEngineConfigMapOutput) ToUserDataEngineConfigMapOutputWithContext(ctx context.Context) UserDataEngineConfigMapOutput {
	return o
}

func (o UserDataEngineConfigMapOutput) MapIndex(k pulumi.StringInput) UserDataEngineConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserDataEngineConfig {
		return vs[0].(map[string]*UserDataEngineConfig)[vs[1].(string)]
	}).(UserDataEngineConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserDataEngineConfigInput)(nil)).Elem(), &UserDataEngineConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserDataEngineConfigArrayInput)(nil)).Elem(), UserDataEngineConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserDataEngineConfigMapInput)(nil)).Elem(), UserDataEngineConfigMap{})
	pulumi.RegisterOutputType(UserDataEngineConfigOutput{})
	pulumi.RegisterOutputType(UserDataEngineConfigArrayOutput{})
	pulumi.RegisterOutputType(UserDataEngineConfigMapOutput{})
}
