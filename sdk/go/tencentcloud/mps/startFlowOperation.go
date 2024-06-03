// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mps

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a mps startFlowOperation
//
// ## Example Usage
//
// ### Start flow
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mps"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mps.NewStartFlowOperation(ctx, "operation", &Mps.StartFlowOperationArgs{
//				FlowId: pulumi.Any(tencentcloud_mps_flow.Flow_rtp.Id),
//				Start:  pulumi.Bool(true),
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
// ### Stop flow
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mps"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mps.NewStartFlowOperation(ctx, "operation", &Mps.StartFlowOperationArgs{
//				FlowId: pulumi.Any(tencentcloud_mps_flow.Flow_rtp.Id),
//				Start:  pulumi.Bool(false),
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
type StartFlowOperation struct {
	pulumi.CustomResourceState

	// Flow Id.
	FlowId pulumi.StringOutput `pulumi:"flowId"`
	// `true`: start mps stream link flow; `false`: stop.
	Start pulumi.BoolOutput `pulumi:"start"`
}

// NewStartFlowOperation registers a new resource with the given unique name, arguments, and options.
func NewStartFlowOperation(ctx *pulumi.Context,
	name string, args *StartFlowOperationArgs, opts ...pulumi.ResourceOption) (*StartFlowOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FlowId == nil {
		return nil, errors.New("invalid value for required argument 'FlowId'")
	}
	if args.Start == nil {
		return nil, errors.New("invalid value for required argument 'Start'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StartFlowOperation
	err := ctx.RegisterResource("tencentcloud:Mps/startFlowOperation:StartFlowOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStartFlowOperation gets an existing StartFlowOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStartFlowOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StartFlowOperationState, opts ...pulumi.ResourceOption) (*StartFlowOperation, error) {
	var resource StartFlowOperation
	err := ctx.ReadResource("tencentcloud:Mps/startFlowOperation:StartFlowOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StartFlowOperation resources.
type startFlowOperationState struct {
	// Flow Id.
	FlowId *string `pulumi:"flowId"`
	// `true`: start mps stream link flow; `false`: stop.
	Start *bool `pulumi:"start"`
}

type StartFlowOperationState struct {
	// Flow Id.
	FlowId pulumi.StringPtrInput
	// `true`: start mps stream link flow; `false`: stop.
	Start pulumi.BoolPtrInput
}

func (StartFlowOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*startFlowOperationState)(nil)).Elem()
}

type startFlowOperationArgs struct {
	// Flow Id.
	FlowId string `pulumi:"flowId"`
	// `true`: start mps stream link flow; `false`: stop.
	Start bool `pulumi:"start"`
}

// The set of arguments for constructing a StartFlowOperation resource.
type StartFlowOperationArgs struct {
	// Flow Id.
	FlowId pulumi.StringInput
	// `true`: start mps stream link flow; `false`: stop.
	Start pulumi.BoolInput
}

func (StartFlowOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*startFlowOperationArgs)(nil)).Elem()
}

type StartFlowOperationInput interface {
	pulumi.Input

	ToStartFlowOperationOutput() StartFlowOperationOutput
	ToStartFlowOperationOutputWithContext(ctx context.Context) StartFlowOperationOutput
}

func (*StartFlowOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**StartFlowOperation)(nil)).Elem()
}

func (i *StartFlowOperation) ToStartFlowOperationOutput() StartFlowOperationOutput {
	return i.ToStartFlowOperationOutputWithContext(context.Background())
}

func (i *StartFlowOperation) ToStartFlowOperationOutputWithContext(ctx context.Context) StartFlowOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartFlowOperationOutput)
}

// StartFlowOperationArrayInput is an input type that accepts StartFlowOperationArray and StartFlowOperationArrayOutput values.
// You can construct a concrete instance of `StartFlowOperationArrayInput` via:
//
//	StartFlowOperationArray{ StartFlowOperationArgs{...} }
type StartFlowOperationArrayInput interface {
	pulumi.Input

	ToStartFlowOperationArrayOutput() StartFlowOperationArrayOutput
	ToStartFlowOperationArrayOutputWithContext(context.Context) StartFlowOperationArrayOutput
}

type StartFlowOperationArray []StartFlowOperationInput

func (StartFlowOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StartFlowOperation)(nil)).Elem()
}

func (i StartFlowOperationArray) ToStartFlowOperationArrayOutput() StartFlowOperationArrayOutput {
	return i.ToStartFlowOperationArrayOutputWithContext(context.Background())
}

func (i StartFlowOperationArray) ToStartFlowOperationArrayOutputWithContext(ctx context.Context) StartFlowOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartFlowOperationArrayOutput)
}

// StartFlowOperationMapInput is an input type that accepts StartFlowOperationMap and StartFlowOperationMapOutput values.
// You can construct a concrete instance of `StartFlowOperationMapInput` via:
//
//	StartFlowOperationMap{ "key": StartFlowOperationArgs{...} }
type StartFlowOperationMapInput interface {
	pulumi.Input

	ToStartFlowOperationMapOutput() StartFlowOperationMapOutput
	ToStartFlowOperationMapOutputWithContext(context.Context) StartFlowOperationMapOutput
}

type StartFlowOperationMap map[string]StartFlowOperationInput

func (StartFlowOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StartFlowOperation)(nil)).Elem()
}

func (i StartFlowOperationMap) ToStartFlowOperationMapOutput() StartFlowOperationMapOutput {
	return i.ToStartFlowOperationMapOutputWithContext(context.Background())
}

func (i StartFlowOperationMap) ToStartFlowOperationMapOutputWithContext(ctx context.Context) StartFlowOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartFlowOperationMapOutput)
}

type StartFlowOperationOutput struct{ *pulumi.OutputState }

func (StartFlowOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StartFlowOperation)(nil)).Elem()
}

func (o StartFlowOperationOutput) ToStartFlowOperationOutput() StartFlowOperationOutput {
	return o
}

func (o StartFlowOperationOutput) ToStartFlowOperationOutputWithContext(ctx context.Context) StartFlowOperationOutput {
	return o
}

// Flow Id.
func (o StartFlowOperationOutput) FlowId() pulumi.StringOutput {
	return o.ApplyT(func(v *StartFlowOperation) pulumi.StringOutput { return v.FlowId }).(pulumi.StringOutput)
}

// `true`: start mps stream link flow; `false`: stop.
func (o StartFlowOperationOutput) Start() pulumi.BoolOutput {
	return o.ApplyT(func(v *StartFlowOperation) pulumi.BoolOutput { return v.Start }).(pulumi.BoolOutput)
}

type StartFlowOperationArrayOutput struct{ *pulumi.OutputState }

func (StartFlowOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StartFlowOperation)(nil)).Elem()
}

func (o StartFlowOperationArrayOutput) ToStartFlowOperationArrayOutput() StartFlowOperationArrayOutput {
	return o
}

func (o StartFlowOperationArrayOutput) ToStartFlowOperationArrayOutputWithContext(ctx context.Context) StartFlowOperationArrayOutput {
	return o
}

func (o StartFlowOperationArrayOutput) Index(i pulumi.IntInput) StartFlowOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StartFlowOperation {
		return vs[0].([]*StartFlowOperation)[vs[1].(int)]
	}).(StartFlowOperationOutput)
}

type StartFlowOperationMapOutput struct{ *pulumi.OutputState }

func (StartFlowOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StartFlowOperation)(nil)).Elem()
}

func (o StartFlowOperationMapOutput) ToStartFlowOperationMapOutput() StartFlowOperationMapOutput {
	return o
}

func (o StartFlowOperationMapOutput) ToStartFlowOperationMapOutputWithContext(ctx context.Context) StartFlowOperationMapOutput {
	return o
}

func (o StartFlowOperationMapOutput) MapIndex(k pulumi.StringInput) StartFlowOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StartFlowOperation {
		return vs[0].(map[string]*StartFlowOperation)[vs[1].(string)]
	}).(StartFlowOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StartFlowOperationInput)(nil)).Elem(), &StartFlowOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*StartFlowOperationArrayInput)(nil)).Elem(), StartFlowOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StartFlowOperationMapInput)(nil)).Elem(), StartFlowOperationMap{})
	pulumi.RegisterOutputType(StartFlowOperationOutput{})
	pulumi.RegisterOutputType(StartFlowOperationArrayOutput{})
	pulumi.RegisterOutputType(StartFlowOperationMapOutput{})
}
