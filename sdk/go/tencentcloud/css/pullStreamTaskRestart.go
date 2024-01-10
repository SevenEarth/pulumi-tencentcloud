// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package css

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a css restartPushTask
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
//			_, err := Css.NewPullStreamTaskRestart(ctx, "restartPushTask", &Css.PullStreamTaskRestartArgs{
//				Operator: pulumi.String("tf-test"),
//				TaskId:   pulumi.String("3573"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type PullStreamTaskRestart struct {
	pulumi.CustomResourceState

	// Task operator.
	Operator pulumi.StringOutput `pulumi:"operator"`
	// Task Id.
	TaskId pulumi.StringOutput `pulumi:"taskId"`
}

// NewPullStreamTaskRestart registers a new resource with the given unique name, arguments, and options.
func NewPullStreamTaskRestart(ctx *pulumi.Context,
	name string, args *PullStreamTaskRestartArgs, opts ...pulumi.ResourceOption) (*PullStreamTaskRestart, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Operator == nil {
		return nil, errors.New("invalid value for required argument 'Operator'")
	}
	if args.TaskId == nil {
		return nil, errors.New("invalid value for required argument 'TaskId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource PullStreamTaskRestart
	err := ctx.RegisterResource("tencentcloud:Css/pullStreamTaskRestart:PullStreamTaskRestart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPullStreamTaskRestart gets an existing PullStreamTaskRestart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPullStreamTaskRestart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PullStreamTaskRestartState, opts ...pulumi.ResourceOption) (*PullStreamTaskRestart, error) {
	var resource PullStreamTaskRestart
	err := ctx.ReadResource("tencentcloud:Css/pullStreamTaskRestart:PullStreamTaskRestart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PullStreamTaskRestart resources.
type pullStreamTaskRestartState struct {
	// Task operator.
	Operator *string `pulumi:"operator"`
	// Task Id.
	TaskId *string `pulumi:"taskId"`
}

type PullStreamTaskRestartState struct {
	// Task operator.
	Operator pulumi.StringPtrInput
	// Task Id.
	TaskId pulumi.StringPtrInput
}

func (PullStreamTaskRestartState) ElementType() reflect.Type {
	return reflect.TypeOf((*pullStreamTaskRestartState)(nil)).Elem()
}

type pullStreamTaskRestartArgs struct {
	// Task operator.
	Operator string `pulumi:"operator"`
	// Task Id.
	TaskId string `pulumi:"taskId"`
}

// The set of arguments for constructing a PullStreamTaskRestart resource.
type PullStreamTaskRestartArgs struct {
	// Task operator.
	Operator pulumi.StringInput
	// Task Id.
	TaskId pulumi.StringInput
}

func (PullStreamTaskRestartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pullStreamTaskRestartArgs)(nil)).Elem()
}

type PullStreamTaskRestartInput interface {
	pulumi.Input

	ToPullStreamTaskRestartOutput() PullStreamTaskRestartOutput
	ToPullStreamTaskRestartOutputWithContext(ctx context.Context) PullStreamTaskRestartOutput
}

func (*PullStreamTaskRestart) ElementType() reflect.Type {
	return reflect.TypeOf((**PullStreamTaskRestart)(nil)).Elem()
}

func (i *PullStreamTaskRestart) ToPullStreamTaskRestartOutput() PullStreamTaskRestartOutput {
	return i.ToPullStreamTaskRestartOutputWithContext(context.Background())
}

func (i *PullStreamTaskRestart) ToPullStreamTaskRestartOutputWithContext(ctx context.Context) PullStreamTaskRestartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PullStreamTaskRestartOutput)
}

// PullStreamTaskRestartArrayInput is an input type that accepts PullStreamTaskRestartArray and PullStreamTaskRestartArrayOutput values.
// You can construct a concrete instance of `PullStreamTaskRestartArrayInput` via:
//
//	PullStreamTaskRestartArray{ PullStreamTaskRestartArgs{...} }
type PullStreamTaskRestartArrayInput interface {
	pulumi.Input

	ToPullStreamTaskRestartArrayOutput() PullStreamTaskRestartArrayOutput
	ToPullStreamTaskRestartArrayOutputWithContext(context.Context) PullStreamTaskRestartArrayOutput
}

type PullStreamTaskRestartArray []PullStreamTaskRestartInput

func (PullStreamTaskRestartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PullStreamTaskRestart)(nil)).Elem()
}

func (i PullStreamTaskRestartArray) ToPullStreamTaskRestartArrayOutput() PullStreamTaskRestartArrayOutput {
	return i.ToPullStreamTaskRestartArrayOutputWithContext(context.Background())
}

func (i PullStreamTaskRestartArray) ToPullStreamTaskRestartArrayOutputWithContext(ctx context.Context) PullStreamTaskRestartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PullStreamTaskRestartArrayOutput)
}

// PullStreamTaskRestartMapInput is an input type that accepts PullStreamTaskRestartMap and PullStreamTaskRestartMapOutput values.
// You can construct a concrete instance of `PullStreamTaskRestartMapInput` via:
//
//	PullStreamTaskRestartMap{ "key": PullStreamTaskRestartArgs{...} }
type PullStreamTaskRestartMapInput interface {
	pulumi.Input

	ToPullStreamTaskRestartMapOutput() PullStreamTaskRestartMapOutput
	ToPullStreamTaskRestartMapOutputWithContext(context.Context) PullStreamTaskRestartMapOutput
}

type PullStreamTaskRestartMap map[string]PullStreamTaskRestartInput

func (PullStreamTaskRestartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PullStreamTaskRestart)(nil)).Elem()
}

func (i PullStreamTaskRestartMap) ToPullStreamTaskRestartMapOutput() PullStreamTaskRestartMapOutput {
	return i.ToPullStreamTaskRestartMapOutputWithContext(context.Background())
}

func (i PullStreamTaskRestartMap) ToPullStreamTaskRestartMapOutputWithContext(ctx context.Context) PullStreamTaskRestartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PullStreamTaskRestartMapOutput)
}

type PullStreamTaskRestartOutput struct{ *pulumi.OutputState }

func (PullStreamTaskRestartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PullStreamTaskRestart)(nil)).Elem()
}

func (o PullStreamTaskRestartOutput) ToPullStreamTaskRestartOutput() PullStreamTaskRestartOutput {
	return o
}

func (o PullStreamTaskRestartOutput) ToPullStreamTaskRestartOutputWithContext(ctx context.Context) PullStreamTaskRestartOutput {
	return o
}

// Task operator.
func (o PullStreamTaskRestartOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v *PullStreamTaskRestart) pulumi.StringOutput { return v.Operator }).(pulumi.StringOutput)
}

// Task Id.
func (o PullStreamTaskRestartOutput) TaskId() pulumi.StringOutput {
	return o.ApplyT(func(v *PullStreamTaskRestart) pulumi.StringOutput { return v.TaskId }).(pulumi.StringOutput)
}

type PullStreamTaskRestartArrayOutput struct{ *pulumi.OutputState }

func (PullStreamTaskRestartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PullStreamTaskRestart)(nil)).Elem()
}

func (o PullStreamTaskRestartArrayOutput) ToPullStreamTaskRestartArrayOutput() PullStreamTaskRestartArrayOutput {
	return o
}

func (o PullStreamTaskRestartArrayOutput) ToPullStreamTaskRestartArrayOutputWithContext(ctx context.Context) PullStreamTaskRestartArrayOutput {
	return o
}

func (o PullStreamTaskRestartArrayOutput) Index(i pulumi.IntInput) PullStreamTaskRestartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PullStreamTaskRestart {
		return vs[0].([]*PullStreamTaskRestart)[vs[1].(int)]
	}).(PullStreamTaskRestartOutput)
}

type PullStreamTaskRestartMapOutput struct{ *pulumi.OutputState }

func (PullStreamTaskRestartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PullStreamTaskRestart)(nil)).Elem()
}

func (o PullStreamTaskRestartMapOutput) ToPullStreamTaskRestartMapOutput() PullStreamTaskRestartMapOutput {
	return o
}

func (o PullStreamTaskRestartMapOutput) ToPullStreamTaskRestartMapOutputWithContext(ctx context.Context) PullStreamTaskRestartMapOutput {
	return o
}

func (o PullStreamTaskRestartMapOutput) MapIndex(k pulumi.StringInput) PullStreamTaskRestartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PullStreamTaskRestart {
		return vs[0].(map[string]*PullStreamTaskRestart)[vs[1].(string)]
	}).(PullStreamTaskRestartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PullStreamTaskRestartInput)(nil)).Elem(), &PullStreamTaskRestart{})
	pulumi.RegisterInputType(reflect.TypeOf((*PullStreamTaskRestartArrayInput)(nil)).Elem(), PullStreamTaskRestartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PullStreamTaskRestartMapInput)(nil)).Elem(), PullStreamTaskRestartMap{})
	pulumi.RegisterOutputType(PullStreamTaskRestartOutput{})
	pulumi.RegisterOutputType(PullStreamTaskRestartArrayOutput{})
	pulumi.RegisterOutputType(PullStreamTaskRestartMapOutput{})
}
