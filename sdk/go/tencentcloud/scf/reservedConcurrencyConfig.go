// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scf

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a scf reservedConcurrencyConfig
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Scf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Scf.NewReservedConcurrencyConfig(ctx, "reservedConcurrencyConfig", &Scf.ReservedConcurrencyConfigArgs{
//				FunctionName:           pulumi.String("keep-1676351130"),
//				Namespace:              pulumi.String("default"),
//				ReservedConcurrencyMem: pulumi.Int(128000),
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
// scf reserved_concurrency_config can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Scf/reservedConcurrencyConfig:ReservedConcurrencyConfig reserved_concurrency_config reserved_concurrency_config_id
// ```
type ReservedConcurrencyConfig struct {
	pulumi.CustomResourceState

	// Specifies the function of which you want to configure the reserved quota.
	FunctionName pulumi.StringOutput `pulumi:"functionName"`
	// Function namespace. Default value: default.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Reserved memory quota of the function. Note: the upper limit for the total reserved quota of the function is the user's total concurrency memory minus 12800.
	ReservedConcurrencyMem pulumi.IntOutput `pulumi:"reservedConcurrencyMem"`
}

// NewReservedConcurrencyConfig registers a new resource with the given unique name, arguments, and options.
func NewReservedConcurrencyConfig(ctx *pulumi.Context,
	name string, args *ReservedConcurrencyConfigArgs, opts ...pulumi.ResourceOption) (*ReservedConcurrencyConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FunctionName == nil {
		return nil, errors.New("invalid value for required argument 'FunctionName'")
	}
	if args.ReservedConcurrencyMem == nil {
		return nil, errors.New("invalid value for required argument 'ReservedConcurrencyMem'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReservedConcurrencyConfig
	err := ctx.RegisterResource("tencentcloud:Scf/reservedConcurrencyConfig:ReservedConcurrencyConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReservedConcurrencyConfig gets an existing ReservedConcurrencyConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReservedConcurrencyConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReservedConcurrencyConfigState, opts ...pulumi.ResourceOption) (*ReservedConcurrencyConfig, error) {
	var resource ReservedConcurrencyConfig
	err := ctx.ReadResource("tencentcloud:Scf/reservedConcurrencyConfig:ReservedConcurrencyConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReservedConcurrencyConfig resources.
type reservedConcurrencyConfigState struct {
	// Specifies the function of which you want to configure the reserved quota.
	FunctionName *string `pulumi:"functionName"`
	// Function namespace. Default value: default.
	Namespace *string `pulumi:"namespace"`
	// Reserved memory quota of the function. Note: the upper limit for the total reserved quota of the function is the user's total concurrency memory minus 12800.
	ReservedConcurrencyMem *int `pulumi:"reservedConcurrencyMem"`
}

type ReservedConcurrencyConfigState struct {
	// Specifies the function of which you want to configure the reserved quota.
	FunctionName pulumi.StringPtrInput
	// Function namespace. Default value: default.
	Namespace pulumi.StringPtrInput
	// Reserved memory quota of the function. Note: the upper limit for the total reserved quota of the function is the user's total concurrency memory minus 12800.
	ReservedConcurrencyMem pulumi.IntPtrInput
}

func (ReservedConcurrencyConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*reservedConcurrencyConfigState)(nil)).Elem()
}

type reservedConcurrencyConfigArgs struct {
	// Specifies the function of which you want to configure the reserved quota.
	FunctionName string `pulumi:"functionName"`
	// Function namespace. Default value: default.
	Namespace *string `pulumi:"namespace"`
	// Reserved memory quota of the function. Note: the upper limit for the total reserved quota of the function is the user's total concurrency memory minus 12800.
	ReservedConcurrencyMem int `pulumi:"reservedConcurrencyMem"`
}

// The set of arguments for constructing a ReservedConcurrencyConfig resource.
type ReservedConcurrencyConfigArgs struct {
	// Specifies the function of which you want to configure the reserved quota.
	FunctionName pulumi.StringInput
	// Function namespace. Default value: default.
	Namespace pulumi.StringPtrInput
	// Reserved memory quota of the function. Note: the upper limit for the total reserved quota of the function is the user's total concurrency memory minus 12800.
	ReservedConcurrencyMem pulumi.IntInput
}

func (ReservedConcurrencyConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reservedConcurrencyConfigArgs)(nil)).Elem()
}

type ReservedConcurrencyConfigInput interface {
	pulumi.Input

	ToReservedConcurrencyConfigOutput() ReservedConcurrencyConfigOutput
	ToReservedConcurrencyConfigOutputWithContext(ctx context.Context) ReservedConcurrencyConfigOutput
}

func (*ReservedConcurrencyConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ReservedConcurrencyConfig)(nil)).Elem()
}

func (i *ReservedConcurrencyConfig) ToReservedConcurrencyConfigOutput() ReservedConcurrencyConfigOutput {
	return i.ToReservedConcurrencyConfigOutputWithContext(context.Background())
}

func (i *ReservedConcurrencyConfig) ToReservedConcurrencyConfigOutputWithContext(ctx context.Context) ReservedConcurrencyConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservedConcurrencyConfigOutput)
}

// ReservedConcurrencyConfigArrayInput is an input type that accepts ReservedConcurrencyConfigArray and ReservedConcurrencyConfigArrayOutput values.
// You can construct a concrete instance of `ReservedConcurrencyConfigArrayInput` via:
//
//	ReservedConcurrencyConfigArray{ ReservedConcurrencyConfigArgs{...} }
type ReservedConcurrencyConfigArrayInput interface {
	pulumi.Input

	ToReservedConcurrencyConfigArrayOutput() ReservedConcurrencyConfigArrayOutput
	ToReservedConcurrencyConfigArrayOutputWithContext(context.Context) ReservedConcurrencyConfigArrayOutput
}

type ReservedConcurrencyConfigArray []ReservedConcurrencyConfigInput

func (ReservedConcurrencyConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReservedConcurrencyConfig)(nil)).Elem()
}

func (i ReservedConcurrencyConfigArray) ToReservedConcurrencyConfigArrayOutput() ReservedConcurrencyConfigArrayOutput {
	return i.ToReservedConcurrencyConfigArrayOutputWithContext(context.Background())
}

func (i ReservedConcurrencyConfigArray) ToReservedConcurrencyConfigArrayOutputWithContext(ctx context.Context) ReservedConcurrencyConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservedConcurrencyConfigArrayOutput)
}

// ReservedConcurrencyConfigMapInput is an input type that accepts ReservedConcurrencyConfigMap and ReservedConcurrencyConfigMapOutput values.
// You can construct a concrete instance of `ReservedConcurrencyConfigMapInput` via:
//
//	ReservedConcurrencyConfigMap{ "key": ReservedConcurrencyConfigArgs{...} }
type ReservedConcurrencyConfigMapInput interface {
	pulumi.Input

	ToReservedConcurrencyConfigMapOutput() ReservedConcurrencyConfigMapOutput
	ToReservedConcurrencyConfigMapOutputWithContext(context.Context) ReservedConcurrencyConfigMapOutput
}

type ReservedConcurrencyConfigMap map[string]ReservedConcurrencyConfigInput

func (ReservedConcurrencyConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReservedConcurrencyConfig)(nil)).Elem()
}

func (i ReservedConcurrencyConfigMap) ToReservedConcurrencyConfigMapOutput() ReservedConcurrencyConfigMapOutput {
	return i.ToReservedConcurrencyConfigMapOutputWithContext(context.Background())
}

func (i ReservedConcurrencyConfigMap) ToReservedConcurrencyConfigMapOutputWithContext(ctx context.Context) ReservedConcurrencyConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservedConcurrencyConfigMapOutput)
}

type ReservedConcurrencyConfigOutput struct{ *pulumi.OutputState }

func (ReservedConcurrencyConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReservedConcurrencyConfig)(nil)).Elem()
}

func (o ReservedConcurrencyConfigOutput) ToReservedConcurrencyConfigOutput() ReservedConcurrencyConfigOutput {
	return o
}

func (o ReservedConcurrencyConfigOutput) ToReservedConcurrencyConfigOutputWithContext(ctx context.Context) ReservedConcurrencyConfigOutput {
	return o
}

// Specifies the function of which you want to configure the reserved quota.
func (o ReservedConcurrencyConfigOutput) FunctionName() pulumi.StringOutput {
	return o.ApplyT(func(v *ReservedConcurrencyConfig) pulumi.StringOutput { return v.FunctionName }).(pulumi.StringOutput)
}

// Function namespace. Default value: default.
func (o ReservedConcurrencyConfigOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReservedConcurrencyConfig) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Reserved memory quota of the function. Note: the upper limit for the total reserved quota of the function is the user's total concurrency memory minus 12800.
func (o ReservedConcurrencyConfigOutput) ReservedConcurrencyMem() pulumi.IntOutput {
	return o.ApplyT(func(v *ReservedConcurrencyConfig) pulumi.IntOutput { return v.ReservedConcurrencyMem }).(pulumi.IntOutput)
}

type ReservedConcurrencyConfigArrayOutput struct{ *pulumi.OutputState }

func (ReservedConcurrencyConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReservedConcurrencyConfig)(nil)).Elem()
}

func (o ReservedConcurrencyConfigArrayOutput) ToReservedConcurrencyConfigArrayOutput() ReservedConcurrencyConfigArrayOutput {
	return o
}

func (o ReservedConcurrencyConfigArrayOutput) ToReservedConcurrencyConfigArrayOutputWithContext(ctx context.Context) ReservedConcurrencyConfigArrayOutput {
	return o
}

func (o ReservedConcurrencyConfigArrayOutput) Index(i pulumi.IntInput) ReservedConcurrencyConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReservedConcurrencyConfig {
		return vs[0].([]*ReservedConcurrencyConfig)[vs[1].(int)]
	}).(ReservedConcurrencyConfigOutput)
}

type ReservedConcurrencyConfigMapOutput struct{ *pulumi.OutputState }

func (ReservedConcurrencyConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReservedConcurrencyConfig)(nil)).Elem()
}

func (o ReservedConcurrencyConfigMapOutput) ToReservedConcurrencyConfigMapOutput() ReservedConcurrencyConfigMapOutput {
	return o
}

func (o ReservedConcurrencyConfigMapOutput) ToReservedConcurrencyConfigMapOutputWithContext(ctx context.Context) ReservedConcurrencyConfigMapOutput {
	return o
}

func (o ReservedConcurrencyConfigMapOutput) MapIndex(k pulumi.StringInput) ReservedConcurrencyConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReservedConcurrencyConfig {
		return vs[0].(map[string]*ReservedConcurrencyConfig)[vs[1].(string)]
	}).(ReservedConcurrencyConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReservedConcurrencyConfigInput)(nil)).Elem(), &ReservedConcurrencyConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReservedConcurrencyConfigArrayInput)(nil)).Elem(), ReservedConcurrencyConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReservedConcurrencyConfigMapInput)(nil)).Elem(), ReservedConcurrencyConfigMap{})
	pulumi.RegisterOutputType(ReservedConcurrencyConfigOutput{})
	pulumi.RegisterOutputType(ReservedConcurrencyConfigArrayOutput{})
	pulumi.RegisterOutputType(ReservedConcurrencyConfigMapOutput{})
}
