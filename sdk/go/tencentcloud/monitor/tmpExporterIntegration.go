// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a monitor tmpExporterIntegration
type TmpExporterIntegration struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Integration config.
	Content pulumi.StringOutput `pulumi:"content"`
	// Instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Type.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Integration config.
	KubeType pulumi.IntOutput `pulumi:"kubeType"`
}

// NewTmpExporterIntegration registers a new resource with the given unique name, arguments, and options.
func NewTmpExporterIntegration(ctx *pulumi.Context,
	name string, args *TmpExporterIntegrationArgs, opts ...pulumi.ResourceOption) (*TmpExporterIntegration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.KubeType == nil {
		return nil, errors.New("invalid value for required argument 'KubeType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource TmpExporterIntegration
	err := ctx.RegisterResource("tencentcloud:Monitor/tmpExporterIntegration:TmpExporterIntegration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTmpExporterIntegration gets an existing TmpExporterIntegration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTmpExporterIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TmpExporterIntegrationState, opts ...pulumi.ResourceOption) (*TmpExporterIntegration, error) {
	var resource TmpExporterIntegration
	err := ctx.ReadResource("tencentcloud:Monitor/tmpExporterIntegration:TmpExporterIntegration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TmpExporterIntegration resources.
type tmpExporterIntegrationState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Integration config.
	Content *string `pulumi:"content"`
	// Instance id.
	InstanceId *string `pulumi:"instanceId"`
	// Type.
	Kind *string `pulumi:"kind"`
	// Integration config.
	KubeType *int `pulumi:"kubeType"`
}

type TmpExporterIntegrationState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Integration config.
	Content pulumi.StringPtrInput
	// Instance id.
	InstanceId pulumi.StringPtrInput
	// Type.
	Kind pulumi.StringPtrInput
	// Integration config.
	KubeType pulumi.IntPtrInput
}

func (TmpExporterIntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpExporterIntegrationState)(nil)).Elem()
}

type tmpExporterIntegrationArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Integration config.
	Content string `pulumi:"content"`
	// Instance id.
	InstanceId string `pulumi:"instanceId"`
	// Type.
	Kind string `pulumi:"kind"`
	// Integration config.
	KubeType int `pulumi:"kubeType"`
}

// The set of arguments for constructing a TmpExporterIntegration resource.
type TmpExporterIntegrationArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Integration config.
	Content pulumi.StringInput
	// Instance id.
	InstanceId pulumi.StringInput
	// Type.
	Kind pulumi.StringInput
	// Integration config.
	KubeType pulumi.IntInput
}

func (TmpExporterIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpExporterIntegrationArgs)(nil)).Elem()
}

type TmpExporterIntegrationInput interface {
	pulumi.Input

	ToTmpExporterIntegrationOutput() TmpExporterIntegrationOutput
	ToTmpExporterIntegrationOutputWithContext(ctx context.Context) TmpExporterIntegrationOutput
}

func (*TmpExporterIntegration) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpExporterIntegration)(nil)).Elem()
}

func (i *TmpExporterIntegration) ToTmpExporterIntegrationOutput() TmpExporterIntegrationOutput {
	return i.ToTmpExporterIntegrationOutputWithContext(context.Background())
}

func (i *TmpExporterIntegration) ToTmpExporterIntegrationOutputWithContext(ctx context.Context) TmpExporterIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpExporterIntegrationOutput)
}

// TmpExporterIntegrationArrayInput is an input type that accepts TmpExporterIntegrationArray and TmpExporterIntegrationArrayOutput values.
// You can construct a concrete instance of `TmpExporterIntegrationArrayInput` via:
//
//	TmpExporterIntegrationArray{ TmpExporterIntegrationArgs{...} }
type TmpExporterIntegrationArrayInput interface {
	pulumi.Input

	ToTmpExporterIntegrationArrayOutput() TmpExporterIntegrationArrayOutput
	ToTmpExporterIntegrationArrayOutputWithContext(context.Context) TmpExporterIntegrationArrayOutput
}

type TmpExporterIntegrationArray []TmpExporterIntegrationInput

func (TmpExporterIntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpExporterIntegration)(nil)).Elem()
}

func (i TmpExporterIntegrationArray) ToTmpExporterIntegrationArrayOutput() TmpExporterIntegrationArrayOutput {
	return i.ToTmpExporterIntegrationArrayOutputWithContext(context.Background())
}

func (i TmpExporterIntegrationArray) ToTmpExporterIntegrationArrayOutputWithContext(ctx context.Context) TmpExporterIntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpExporterIntegrationArrayOutput)
}

// TmpExporterIntegrationMapInput is an input type that accepts TmpExporterIntegrationMap and TmpExporterIntegrationMapOutput values.
// You can construct a concrete instance of `TmpExporterIntegrationMapInput` via:
//
//	TmpExporterIntegrationMap{ "key": TmpExporterIntegrationArgs{...} }
type TmpExporterIntegrationMapInput interface {
	pulumi.Input

	ToTmpExporterIntegrationMapOutput() TmpExporterIntegrationMapOutput
	ToTmpExporterIntegrationMapOutputWithContext(context.Context) TmpExporterIntegrationMapOutput
}

type TmpExporterIntegrationMap map[string]TmpExporterIntegrationInput

func (TmpExporterIntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpExporterIntegration)(nil)).Elem()
}

func (i TmpExporterIntegrationMap) ToTmpExporterIntegrationMapOutput() TmpExporterIntegrationMapOutput {
	return i.ToTmpExporterIntegrationMapOutputWithContext(context.Background())
}

func (i TmpExporterIntegrationMap) ToTmpExporterIntegrationMapOutputWithContext(ctx context.Context) TmpExporterIntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpExporterIntegrationMapOutput)
}

type TmpExporterIntegrationOutput struct{ *pulumi.OutputState }

func (TmpExporterIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpExporterIntegration)(nil)).Elem()
}

func (o TmpExporterIntegrationOutput) ToTmpExporterIntegrationOutput() TmpExporterIntegrationOutput {
	return o
}

func (o TmpExporterIntegrationOutput) ToTmpExporterIntegrationOutputWithContext(ctx context.Context) TmpExporterIntegrationOutput {
	return o
}

// Cluster ID.
func (o TmpExporterIntegrationOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpExporterIntegration) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Integration config.
func (o TmpExporterIntegrationOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpExporterIntegration) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// Instance id.
func (o TmpExporterIntegrationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpExporterIntegration) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Type.
func (o TmpExporterIntegrationOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpExporterIntegration) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Integration config.
func (o TmpExporterIntegrationOutput) KubeType() pulumi.IntOutput {
	return o.ApplyT(func(v *TmpExporterIntegration) pulumi.IntOutput { return v.KubeType }).(pulumi.IntOutput)
}

type TmpExporterIntegrationArrayOutput struct{ *pulumi.OutputState }

func (TmpExporterIntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpExporterIntegration)(nil)).Elem()
}

func (o TmpExporterIntegrationArrayOutput) ToTmpExporterIntegrationArrayOutput() TmpExporterIntegrationArrayOutput {
	return o
}

func (o TmpExporterIntegrationArrayOutput) ToTmpExporterIntegrationArrayOutputWithContext(ctx context.Context) TmpExporterIntegrationArrayOutput {
	return o
}

func (o TmpExporterIntegrationArrayOutput) Index(i pulumi.IntInput) TmpExporterIntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TmpExporterIntegration {
		return vs[0].([]*TmpExporterIntegration)[vs[1].(int)]
	}).(TmpExporterIntegrationOutput)
}

type TmpExporterIntegrationMapOutput struct{ *pulumi.OutputState }

func (TmpExporterIntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpExporterIntegration)(nil)).Elem()
}

func (o TmpExporterIntegrationMapOutput) ToTmpExporterIntegrationMapOutput() TmpExporterIntegrationMapOutput {
	return o
}

func (o TmpExporterIntegrationMapOutput) ToTmpExporterIntegrationMapOutputWithContext(ctx context.Context) TmpExporterIntegrationMapOutput {
	return o
}

func (o TmpExporterIntegrationMapOutput) MapIndex(k pulumi.StringInput) TmpExporterIntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TmpExporterIntegration {
		return vs[0].(map[string]*TmpExporterIntegration)[vs[1].(string)]
	}).(TmpExporterIntegrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TmpExporterIntegrationInput)(nil)).Elem(), &TmpExporterIntegration{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpExporterIntegrationArrayInput)(nil)).Elem(), TmpExporterIntegrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpExporterIntegrationMapInput)(nil)).Elem(), TmpExporterIntegrationMap{})
	pulumi.RegisterOutputType(TmpExporterIntegrationOutput{})
	pulumi.RegisterOutputType(TmpExporterIntegrationArrayOutput{})
	pulumi.RegisterOutputType(TmpExporterIntegrationMapOutput{})
}
