// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package css

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a css enableOptimalSwitching
//
// > **NOTE:** This resource is only valid when the push stream. When the push stream ends, it will be deleted.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Css"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Css.NewEnableOptimalSwitching(ctx, "enableOptimalSwitching", &Css.EnableOptimalSwitchingArgs{
//				EnableSwitch:  pulumi.Int(1),
//				HostGroupName: pulumi.String("test-group"),
//				StreamName:    pulumi.String("1308919341_test"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// css domain can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Css/enableOptimalSwitching:EnableOptimalSwitching enable_optimal_switching streamName
//
// ```
type EnableOptimalSwitching struct {
	pulumi.CustomResourceState

	// `0`:disabled, `1`:enable.
	EnableSwitch pulumi.IntPtrOutput `pulumi:"enableSwitch"`
	// Group name.
	HostGroupName pulumi.StringPtrOutput `pulumi:"hostGroupName"`
	// Stream id.
	StreamName pulumi.StringOutput `pulumi:"streamName"`
}

// NewEnableOptimalSwitching registers a new resource with the given unique name, arguments, and options.
func NewEnableOptimalSwitching(ctx *pulumi.Context,
	name string, args *EnableOptimalSwitchingArgs, opts ...pulumi.ResourceOption) (*EnableOptimalSwitching, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StreamName == nil {
		return nil, errors.New("invalid value for required argument 'StreamName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource EnableOptimalSwitching
	err := ctx.RegisterResource("tencentcloud:Css/enableOptimalSwitching:EnableOptimalSwitching", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnableOptimalSwitching gets an existing EnableOptimalSwitching resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnableOptimalSwitching(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnableOptimalSwitchingState, opts ...pulumi.ResourceOption) (*EnableOptimalSwitching, error) {
	var resource EnableOptimalSwitching
	err := ctx.ReadResource("tencentcloud:Css/enableOptimalSwitching:EnableOptimalSwitching", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnableOptimalSwitching resources.
type enableOptimalSwitchingState struct {
	// `0`:disabled, `1`:enable.
	EnableSwitch *int `pulumi:"enableSwitch"`
	// Group name.
	HostGroupName *string `pulumi:"hostGroupName"`
	// Stream id.
	StreamName *string `pulumi:"streamName"`
}

type EnableOptimalSwitchingState struct {
	// `0`:disabled, `1`:enable.
	EnableSwitch pulumi.IntPtrInput
	// Group name.
	HostGroupName pulumi.StringPtrInput
	// Stream id.
	StreamName pulumi.StringPtrInput
}

func (EnableOptimalSwitchingState) ElementType() reflect.Type {
	return reflect.TypeOf((*enableOptimalSwitchingState)(nil)).Elem()
}

type enableOptimalSwitchingArgs struct {
	// `0`:disabled, `1`:enable.
	EnableSwitch *int `pulumi:"enableSwitch"`
	// Group name.
	HostGroupName *string `pulumi:"hostGroupName"`
	// Stream id.
	StreamName string `pulumi:"streamName"`
}

// The set of arguments for constructing a EnableOptimalSwitching resource.
type EnableOptimalSwitchingArgs struct {
	// `0`:disabled, `1`:enable.
	EnableSwitch pulumi.IntPtrInput
	// Group name.
	HostGroupName pulumi.StringPtrInput
	// Stream id.
	StreamName pulumi.StringInput
}

func (EnableOptimalSwitchingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enableOptimalSwitchingArgs)(nil)).Elem()
}

type EnableOptimalSwitchingInput interface {
	pulumi.Input

	ToEnableOptimalSwitchingOutput() EnableOptimalSwitchingOutput
	ToEnableOptimalSwitchingOutputWithContext(ctx context.Context) EnableOptimalSwitchingOutput
}

func (*EnableOptimalSwitching) ElementType() reflect.Type {
	return reflect.TypeOf((**EnableOptimalSwitching)(nil)).Elem()
}

func (i *EnableOptimalSwitching) ToEnableOptimalSwitchingOutput() EnableOptimalSwitchingOutput {
	return i.ToEnableOptimalSwitchingOutputWithContext(context.Background())
}

func (i *EnableOptimalSwitching) ToEnableOptimalSwitchingOutputWithContext(ctx context.Context) EnableOptimalSwitchingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnableOptimalSwitchingOutput)
}

// EnableOptimalSwitchingArrayInput is an input type that accepts EnableOptimalSwitchingArray and EnableOptimalSwitchingArrayOutput values.
// You can construct a concrete instance of `EnableOptimalSwitchingArrayInput` via:
//
//	EnableOptimalSwitchingArray{ EnableOptimalSwitchingArgs{...} }
type EnableOptimalSwitchingArrayInput interface {
	pulumi.Input

	ToEnableOptimalSwitchingArrayOutput() EnableOptimalSwitchingArrayOutput
	ToEnableOptimalSwitchingArrayOutputWithContext(context.Context) EnableOptimalSwitchingArrayOutput
}

type EnableOptimalSwitchingArray []EnableOptimalSwitchingInput

func (EnableOptimalSwitchingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnableOptimalSwitching)(nil)).Elem()
}

func (i EnableOptimalSwitchingArray) ToEnableOptimalSwitchingArrayOutput() EnableOptimalSwitchingArrayOutput {
	return i.ToEnableOptimalSwitchingArrayOutputWithContext(context.Background())
}

func (i EnableOptimalSwitchingArray) ToEnableOptimalSwitchingArrayOutputWithContext(ctx context.Context) EnableOptimalSwitchingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnableOptimalSwitchingArrayOutput)
}

// EnableOptimalSwitchingMapInput is an input type that accepts EnableOptimalSwitchingMap and EnableOptimalSwitchingMapOutput values.
// You can construct a concrete instance of `EnableOptimalSwitchingMapInput` via:
//
//	EnableOptimalSwitchingMap{ "key": EnableOptimalSwitchingArgs{...} }
type EnableOptimalSwitchingMapInput interface {
	pulumi.Input

	ToEnableOptimalSwitchingMapOutput() EnableOptimalSwitchingMapOutput
	ToEnableOptimalSwitchingMapOutputWithContext(context.Context) EnableOptimalSwitchingMapOutput
}

type EnableOptimalSwitchingMap map[string]EnableOptimalSwitchingInput

func (EnableOptimalSwitchingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnableOptimalSwitching)(nil)).Elem()
}

func (i EnableOptimalSwitchingMap) ToEnableOptimalSwitchingMapOutput() EnableOptimalSwitchingMapOutput {
	return i.ToEnableOptimalSwitchingMapOutputWithContext(context.Background())
}

func (i EnableOptimalSwitchingMap) ToEnableOptimalSwitchingMapOutputWithContext(ctx context.Context) EnableOptimalSwitchingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnableOptimalSwitchingMapOutput)
}

type EnableOptimalSwitchingOutput struct{ *pulumi.OutputState }

func (EnableOptimalSwitchingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnableOptimalSwitching)(nil)).Elem()
}

func (o EnableOptimalSwitchingOutput) ToEnableOptimalSwitchingOutput() EnableOptimalSwitchingOutput {
	return o
}

func (o EnableOptimalSwitchingOutput) ToEnableOptimalSwitchingOutputWithContext(ctx context.Context) EnableOptimalSwitchingOutput {
	return o
}

// `0`:disabled, `1`:enable.
func (o EnableOptimalSwitchingOutput) EnableSwitch() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EnableOptimalSwitching) pulumi.IntPtrOutput { return v.EnableSwitch }).(pulumi.IntPtrOutput)
}

// Group name.
func (o EnableOptimalSwitchingOutput) HostGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnableOptimalSwitching) pulumi.StringPtrOutput { return v.HostGroupName }).(pulumi.StringPtrOutput)
}

// Stream id.
func (o EnableOptimalSwitchingOutput) StreamName() pulumi.StringOutput {
	return o.ApplyT(func(v *EnableOptimalSwitching) pulumi.StringOutput { return v.StreamName }).(pulumi.StringOutput)
}

type EnableOptimalSwitchingArrayOutput struct{ *pulumi.OutputState }

func (EnableOptimalSwitchingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnableOptimalSwitching)(nil)).Elem()
}

func (o EnableOptimalSwitchingArrayOutput) ToEnableOptimalSwitchingArrayOutput() EnableOptimalSwitchingArrayOutput {
	return o
}

func (o EnableOptimalSwitchingArrayOutput) ToEnableOptimalSwitchingArrayOutputWithContext(ctx context.Context) EnableOptimalSwitchingArrayOutput {
	return o
}

func (o EnableOptimalSwitchingArrayOutput) Index(i pulumi.IntInput) EnableOptimalSwitchingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnableOptimalSwitching {
		return vs[0].([]*EnableOptimalSwitching)[vs[1].(int)]
	}).(EnableOptimalSwitchingOutput)
}

type EnableOptimalSwitchingMapOutput struct{ *pulumi.OutputState }

func (EnableOptimalSwitchingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnableOptimalSwitching)(nil)).Elem()
}

func (o EnableOptimalSwitchingMapOutput) ToEnableOptimalSwitchingMapOutput() EnableOptimalSwitchingMapOutput {
	return o
}

func (o EnableOptimalSwitchingMapOutput) ToEnableOptimalSwitchingMapOutputWithContext(ctx context.Context) EnableOptimalSwitchingMapOutput {
	return o
}

func (o EnableOptimalSwitchingMapOutput) MapIndex(k pulumi.StringInput) EnableOptimalSwitchingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnableOptimalSwitching {
		return vs[0].(map[string]*EnableOptimalSwitching)[vs[1].(string)]
	}).(EnableOptimalSwitchingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnableOptimalSwitchingInput)(nil)).Elem(), &EnableOptimalSwitching{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnableOptimalSwitchingArrayInput)(nil)).Elem(), EnableOptimalSwitchingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnableOptimalSwitchingMapInput)(nil)).Elem(), EnableOptimalSwitchingMap{})
	pulumi.RegisterOutputType(EnableOptimalSwitchingOutput{})
	pulumi.RegisterOutputType(EnableOptimalSwitchingArrayOutput{})
	pulumi.RegisterOutputType(EnableOptimalSwitchingMapOutput{})
}
