// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Auto scaling group for kubernetes cluster (offlined).
//
// > **NOTE:**  This resource was offline no longer suppored.
type AsScalingGroup struct {
	pulumi.CustomResourceState

	// Auto scaling config parameters.
	AutoScalingConfig AsScalingGroupAutoScalingConfigOutput `pulumi:"autoScalingConfig"`
	// Auto scaling group parameters.
	AutoScalingGroup AsScalingGroupAutoScalingGroupOutput `pulumi:"autoScalingGroup"`
	// ID of the cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Custom parameter information related to the node.
	ExtraArgs pulumi.StringArrayOutput `pulumi:"extraArgs"`
	// Labels of kubernetes AS Group created nodes.
	Labels pulumi.MapOutput `pulumi:"labels"`
	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	Unschedulable pulumi.IntPtrOutput `pulumi:"unschedulable"`
}

// NewAsScalingGroup registers a new resource with the given unique name, arguments, and options.
func NewAsScalingGroup(ctx *pulumi.Context,
	name string, args *AsScalingGroupArgs, opts ...pulumi.ResourceOption) (*AsScalingGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoScalingConfig == nil {
		return nil, errors.New("invalid value for required argument 'AutoScalingConfig'")
	}
	if args.AutoScalingGroup == nil {
		return nil, errors.New("invalid value for required argument 'AutoScalingGroup'")
	}
	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AsScalingGroup
	err := ctx.RegisterResource("tencentcloud:Kubernetes/asScalingGroup:AsScalingGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAsScalingGroup gets an existing AsScalingGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAsScalingGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AsScalingGroupState, opts ...pulumi.ResourceOption) (*AsScalingGroup, error) {
	var resource AsScalingGroup
	err := ctx.ReadResource("tencentcloud:Kubernetes/asScalingGroup:AsScalingGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AsScalingGroup resources.
type asScalingGroupState struct {
	// Auto scaling config parameters.
	AutoScalingConfig *AsScalingGroupAutoScalingConfig `pulumi:"autoScalingConfig"`
	// Auto scaling group parameters.
	AutoScalingGroup *AsScalingGroupAutoScalingGroup `pulumi:"autoScalingGroup"`
	// ID of the cluster.
	ClusterId *string `pulumi:"clusterId"`
	// Custom parameter information related to the node.
	ExtraArgs []string `pulumi:"extraArgs"`
	// Labels of kubernetes AS Group created nodes.
	Labels map[string]interface{} `pulumi:"labels"`
	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	Unschedulable *int `pulumi:"unschedulable"`
}

type AsScalingGroupState struct {
	// Auto scaling config parameters.
	AutoScalingConfig AsScalingGroupAutoScalingConfigPtrInput
	// Auto scaling group parameters.
	AutoScalingGroup AsScalingGroupAutoScalingGroupPtrInput
	// ID of the cluster.
	ClusterId pulumi.StringPtrInput
	// Custom parameter information related to the node.
	ExtraArgs pulumi.StringArrayInput
	// Labels of kubernetes AS Group created nodes.
	Labels pulumi.MapInput
	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	Unschedulable pulumi.IntPtrInput
}

func (AsScalingGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*asScalingGroupState)(nil)).Elem()
}

type asScalingGroupArgs struct {
	// Auto scaling config parameters.
	AutoScalingConfig AsScalingGroupAutoScalingConfig `pulumi:"autoScalingConfig"`
	// Auto scaling group parameters.
	AutoScalingGroup AsScalingGroupAutoScalingGroup `pulumi:"autoScalingGroup"`
	// ID of the cluster.
	ClusterId string `pulumi:"clusterId"`
	// Custom parameter information related to the node.
	ExtraArgs []string `pulumi:"extraArgs"`
	// Labels of kubernetes AS Group created nodes.
	Labels map[string]interface{} `pulumi:"labels"`
	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	Unschedulable *int `pulumi:"unschedulable"`
}

// The set of arguments for constructing a AsScalingGroup resource.
type AsScalingGroupArgs struct {
	// Auto scaling config parameters.
	AutoScalingConfig AsScalingGroupAutoScalingConfigInput
	// Auto scaling group parameters.
	AutoScalingGroup AsScalingGroupAutoScalingGroupInput
	// ID of the cluster.
	ClusterId pulumi.StringInput
	// Custom parameter information related to the node.
	ExtraArgs pulumi.StringArrayInput
	// Labels of kubernetes AS Group created nodes.
	Labels pulumi.MapInput
	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	Unschedulable pulumi.IntPtrInput
}

func (AsScalingGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*asScalingGroupArgs)(nil)).Elem()
}

type AsScalingGroupInput interface {
	pulumi.Input

	ToAsScalingGroupOutput() AsScalingGroupOutput
	ToAsScalingGroupOutputWithContext(ctx context.Context) AsScalingGroupOutput
}

func (*AsScalingGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**AsScalingGroup)(nil)).Elem()
}

func (i *AsScalingGroup) ToAsScalingGroupOutput() AsScalingGroupOutput {
	return i.ToAsScalingGroupOutputWithContext(context.Background())
}

func (i *AsScalingGroup) ToAsScalingGroupOutputWithContext(ctx context.Context) AsScalingGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsScalingGroupOutput)
}

// AsScalingGroupArrayInput is an input type that accepts AsScalingGroupArray and AsScalingGroupArrayOutput values.
// You can construct a concrete instance of `AsScalingGroupArrayInput` via:
//
//	AsScalingGroupArray{ AsScalingGroupArgs{...} }
type AsScalingGroupArrayInput interface {
	pulumi.Input

	ToAsScalingGroupArrayOutput() AsScalingGroupArrayOutput
	ToAsScalingGroupArrayOutputWithContext(context.Context) AsScalingGroupArrayOutput
}

type AsScalingGroupArray []AsScalingGroupInput

func (AsScalingGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AsScalingGroup)(nil)).Elem()
}

func (i AsScalingGroupArray) ToAsScalingGroupArrayOutput() AsScalingGroupArrayOutput {
	return i.ToAsScalingGroupArrayOutputWithContext(context.Background())
}

func (i AsScalingGroupArray) ToAsScalingGroupArrayOutputWithContext(ctx context.Context) AsScalingGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsScalingGroupArrayOutput)
}

// AsScalingGroupMapInput is an input type that accepts AsScalingGroupMap and AsScalingGroupMapOutput values.
// You can construct a concrete instance of `AsScalingGroupMapInput` via:
//
//	AsScalingGroupMap{ "key": AsScalingGroupArgs{...} }
type AsScalingGroupMapInput interface {
	pulumi.Input

	ToAsScalingGroupMapOutput() AsScalingGroupMapOutput
	ToAsScalingGroupMapOutputWithContext(context.Context) AsScalingGroupMapOutput
}

type AsScalingGroupMap map[string]AsScalingGroupInput

func (AsScalingGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AsScalingGroup)(nil)).Elem()
}

func (i AsScalingGroupMap) ToAsScalingGroupMapOutput() AsScalingGroupMapOutput {
	return i.ToAsScalingGroupMapOutputWithContext(context.Background())
}

func (i AsScalingGroupMap) ToAsScalingGroupMapOutputWithContext(ctx context.Context) AsScalingGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsScalingGroupMapOutput)
}

type AsScalingGroupOutput struct{ *pulumi.OutputState }

func (AsScalingGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AsScalingGroup)(nil)).Elem()
}

func (o AsScalingGroupOutput) ToAsScalingGroupOutput() AsScalingGroupOutput {
	return o
}

func (o AsScalingGroupOutput) ToAsScalingGroupOutputWithContext(ctx context.Context) AsScalingGroupOutput {
	return o
}

// Auto scaling config parameters.
func (o AsScalingGroupOutput) AutoScalingConfig() AsScalingGroupAutoScalingConfigOutput {
	return o.ApplyT(func(v *AsScalingGroup) AsScalingGroupAutoScalingConfigOutput { return v.AutoScalingConfig }).(AsScalingGroupAutoScalingConfigOutput)
}

// Auto scaling group parameters.
func (o AsScalingGroupOutput) AutoScalingGroup() AsScalingGroupAutoScalingGroupOutput {
	return o.ApplyT(func(v *AsScalingGroup) AsScalingGroupAutoScalingGroupOutput { return v.AutoScalingGroup }).(AsScalingGroupAutoScalingGroupOutput)
}

// ID of the cluster.
func (o AsScalingGroupOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *AsScalingGroup) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Custom parameter information related to the node.
func (o AsScalingGroupOutput) ExtraArgs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AsScalingGroup) pulumi.StringArrayOutput { return v.ExtraArgs }).(pulumi.StringArrayOutput)
}

// Labels of kubernetes AS Group created nodes.
func (o AsScalingGroupOutput) Labels() pulumi.MapOutput {
	return o.ApplyT(func(v *AsScalingGroup) pulumi.MapOutput { return v.Labels }).(pulumi.MapOutput)
}

// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
func (o AsScalingGroupOutput) Unschedulable() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AsScalingGroup) pulumi.IntPtrOutput { return v.Unschedulable }).(pulumi.IntPtrOutput)
}

type AsScalingGroupArrayOutput struct{ *pulumi.OutputState }

func (AsScalingGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AsScalingGroup)(nil)).Elem()
}

func (o AsScalingGroupArrayOutput) ToAsScalingGroupArrayOutput() AsScalingGroupArrayOutput {
	return o
}

func (o AsScalingGroupArrayOutput) ToAsScalingGroupArrayOutputWithContext(ctx context.Context) AsScalingGroupArrayOutput {
	return o
}

func (o AsScalingGroupArrayOutput) Index(i pulumi.IntInput) AsScalingGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AsScalingGroup {
		return vs[0].([]*AsScalingGroup)[vs[1].(int)]
	}).(AsScalingGroupOutput)
}

type AsScalingGroupMapOutput struct{ *pulumi.OutputState }

func (AsScalingGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AsScalingGroup)(nil)).Elem()
}

func (o AsScalingGroupMapOutput) ToAsScalingGroupMapOutput() AsScalingGroupMapOutput {
	return o
}

func (o AsScalingGroupMapOutput) ToAsScalingGroupMapOutputWithContext(ctx context.Context) AsScalingGroupMapOutput {
	return o
}

func (o AsScalingGroupMapOutput) MapIndex(k pulumi.StringInput) AsScalingGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AsScalingGroup {
		return vs[0].(map[string]*AsScalingGroup)[vs[1].(string)]
	}).(AsScalingGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AsScalingGroupInput)(nil)).Elem(), &AsScalingGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*AsScalingGroupArrayInput)(nil)).Elem(), AsScalingGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AsScalingGroupMapInput)(nil)).Elem(), AsScalingGroupMap{})
	pulumi.RegisterOutputType(AsScalingGroupOutput{})
	pulumi.RegisterOutputType(AsScalingGroupArrayOutput{})
	pulumi.RegisterOutputType(AsScalingGroupMapOutput{})
}
