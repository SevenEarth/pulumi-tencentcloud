// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a monitor tmpTkeBasicConfig
type TmpTkeBasicConfig struct {
	pulumi.CustomResourceState

	// ID of cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Type of cluster.
	ClusterType pulumi.StringOutput `pulumi:"clusterType"`
	// Full configuration in yaml format.
	Config pulumi.StringOutput `pulumi:"config"`
	// config type, `serviceMonitors`, `podMonitors`, `rawJobs`.
	ConfigType pulumi.StringOutput `pulumi:"configType"`
	// ID of instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Configure the name of the metric to keep on.
	MetricsNames pulumi.StringArrayOutput `pulumi:"metricsNames"`
	// Name. The naming rule is: namespace/name. If you don&#39;t have any namespace, use the default namespace: kube-system, otherwise use the specified one.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewTmpTkeBasicConfig registers a new resource with the given unique name, arguments, and options.
func NewTmpTkeBasicConfig(ctx *pulumi.Context,
	name string, args *TmpTkeBasicConfigArgs, opts ...pulumi.ResourceOption) (*TmpTkeBasicConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.ClusterType == nil {
		return nil, errors.New("invalid value for required argument 'ClusterType'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.MetricsNames == nil {
		return nil, errors.New("invalid value for required argument 'MetricsNames'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TmpTkeBasicConfig
	err := ctx.RegisterResource("tencentcloud:Monitor/tmpTkeBasicConfig:TmpTkeBasicConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTmpTkeBasicConfig gets an existing TmpTkeBasicConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTmpTkeBasicConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TmpTkeBasicConfigState, opts ...pulumi.ResourceOption) (*TmpTkeBasicConfig, error) {
	var resource TmpTkeBasicConfig
	err := ctx.ReadResource("tencentcloud:Monitor/tmpTkeBasicConfig:TmpTkeBasicConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TmpTkeBasicConfig resources.
type tmpTkeBasicConfigState struct {
	// ID of cluster.
	ClusterId *string `pulumi:"clusterId"`
	// Type of cluster.
	ClusterType *string `pulumi:"clusterType"`
	// Full configuration in yaml format.
	Config *string `pulumi:"config"`
	// config type, `serviceMonitors`, `podMonitors`, `rawJobs`.
	ConfigType *string `pulumi:"configType"`
	// ID of instance.
	InstanceId *string `pulumi:"instanceId"`
	// Configure the name of the metric to keep on.
	MetricsNames []string `pulumi:"metricsNames"`
	// Name. The naming rule is: namespace/name. If you don&#39;t have any namespace, use the default namespace: kube-system, otherwise use the specified one.
	Name *string `pulumi:"name"`
}

type TmpTkeBasicConfigState struct {
	// ID of cluster.
	ClusterId pulumi.StringPtrInput
	// Type of cluster.
	ClusterType pulumi.StringPtrInput
	// Full configuration in yaml format.
	Config pulumi.StringPtrInput
	// config type, `serviceMonitors`, `podMonitors`, `rawJobs`.
	ConfigType pulumi.StringPtrInput
	// ID of instance.
	InstanceId pulumi.StringPtrInput
	// Configure the name of the metric to keep on.
	MetricsNames pulumi.StringArrayInput
	// Name. The naming rule is: namespace/name. If you don&#39;t have any namespace, use the default namespace: kube-system, otherwise use the specified one.
	Name pulumi.StringPtrInput
}

func (TmpTkeBasicConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpTkeBasicConfigState)(nil)).Elem()
}

type tmpTkeBasicConfigArgs struct {
	// ID of cluster.
	ClusterId string `pulumi:"clusterId"`
	// Type of cluster.
	ClusterType string `pulumi:"clusterType"`
	// ID of instance.
	InstanceId string `pulumi:"instanceId"`
	// Configure the name of the metric to keep on.
	MetricsNames []string `pulumi:"metricsNames"`
	// Name. The naming rule is: namespace/name. If you don&#39;t have any namespace, use the default namespace: kube-system, otherwise use the specified one.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a TmpTkeBasicConfig resource.
type TmpTkeBasicConfigArgs struct {
	// ID of cluster.
	ClusterId pulumi.StringInput
	// Type of cluster.
	ClusterType pulumi.StringInput
	// ID of instance.
	InstanceId pulumi.StringInput
	// Configure the name of the metric to keep on.
	MetricsNames pulumi.StringArrayInput
	// Name. The naming rule is: namespace/name. If you don&#39;t have any namespace, use the default namespace: kube-system, otherwise use the specified one.
	Name pulumi.StringPtrInput
}

func (TmpTkeBasicConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpTkeBasicConfigArgs)(nil)).Elem()
}

type TmpTkeBasicConfigInput interface {
	pulumi.Input

	ToTmpTkeBasicConfigOutput() TmpTkeBasicConfigOutput
	ToTmpTkeBasicConfigOutputWithContext(ctx context.Context) TmpTkeBasicConfigOutput
}

func (*TmpTkeBasicConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpTkeBasicConfig)(nil)).Elem()
}

func (i *TmpTkeBasicConfig) ToTmpTkeBasicConfigOutput() TmpTkeBasicConfigOutput {
	return i.ToTmpTkeBasicConfigOutputWithContext(context.Background())
}

func (i *TmpTkeBasicConfig) ToTmpTkeBasicConfigOutputWithContext(ctx context.Context) TmpTkeBasicConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpTkeBasicConfigOutput)
}

// TmpTkeBasicConfigArrayInput is an input type that accepts TmpTkeBasicConfigArray and TmpTkeBasicConfigArrayOutput values.
// You can construct a concrete instance of `TmpTkeBasicConfigArrayInput` via:
//
//	TmpTkeBasicConfigArray{ TmpTkeBasicConfigArgs{...} }
type TmpTkeBasicConfigArrayInput interface {
	pulumi.Input

	ToTmpTkeBasicConfigArrayOutput() TmpTkeBasicConfigArrayOutput
	ToTmpTkeBasicConfigArrayOutputWithContext(context.Context) TmpTkeBasicConfigArrayOutput
}

type TmpTkeBasicConfigArray []TmpTkeBasicConfigInput

func (TmpTkeBasicConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpTkeBasicConfig)(nil)).Elem()
}

func (i TmpTkeBasicConfigArray) ToTmpTkeBasicConfigArrayOutput() TmpTkeBasicConfigArrayOutput {
	return i.ToTmpTkeBasicConfigArrayOutputWithContext(context.Background())
}

func (i TmpTkeBasicConfigArray) ToTmpTkeBasicConfigArrayOutputWithContext(ctx context.Context) TmpTkeBasicConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpTkeBasicConfigArrayOutput)
}

// TmpTkeBasicConfigMapInput is an input type that accepts TmpTkeBasicConfigMap and TmpTkeBasicConfigMapOutput values.
// You can construct a concrete instance of `TmpTkeBasicConfigMapInput` via:
//
//	TmpTkeBasicConfigMap{ "key": TmpTkeBasicConfigArgs{...} }
type TmpTkeBasicConfigMapInput interface {
	pulumi.Input

	ToTmpTkeBasicConfigMapOutput() TmpTkeBasicConfigMapOutput
	ToTmpTkeBasicConfigMapOutputWithContext(context.Context) TmpTkeBasicConfigMapOutput
}

type TmpTkeBasicConfigMap map[string]TmpTkeBasicConfigInput

func (TmpTkeBasicConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpTkeBasicConfig)(nil)).Elem()
}

func (i TmpTkeBasicConfigMap) ToTmpTkeBasicConfigMapOutput() TmpTkeBasicConfigMapOutput {
	return i.ToTmpTkeBasicConfigMapOutputWithContext(context.Background())
}

func (i TmpTkeBasicConfigMap) ToTmpTkeBasicConfigMapOutputWithContext(ctx context.Context) TmpTkeBasicConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpTkeBasicConfigMapOutput)
}

type TmpTkeBasicConfigOutput struct{ *pulumi.OutputState }

func (TmpTkeBasicConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpTkeBasicConfig)(nil)).Elem()
}

func (o TmpTkeBasicConfigOutput) ToTmpTkeBasicConfigOutput() TmpTkeBasicConfigOutput {
	return o
}

func (o TmpTkeBasicConfigOutput) ToTmpTkeBasicConfigOutputWithContext(ctx context.Context) TmpTkeBasicConfigOutput {
	return o
}

// ID of cluster.
func (o TmpTkeBasicConfigOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpTkeBasicConfig) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Type of cluster.
func (o TmpTkeBasicConfigOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpTkeBasicConfig) pulumi.StringOutput { return v.ClusterType }).(pulumi.StringOutput)
}

// Full configuration in yaml format.
func (o TmpTkeBasicConfigOutput) Config() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpTkeBasicConfig) pulumi.StringOutput { return v.Config }).(pulumi.StringOutput)
}

// config type, `serviceMonitors`, `podMonitors`, `rawJobs`.
func (o TmpTkeBasicConfigOutput) ConfigType() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpTkeBasicConfig) pulumi.StringOutput { return v.ConfigType }).(pulumi.StringOutput)
}

// ID of instance.
func (o TmpTkeBasicConfigOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpTkeBasicConfig) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Configure the name of the metric to keep on.
func (o TmpTkeBasicConfigOutput) MetricsNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TmpTkeBasicConfig) pulumi.StringArrayOutput { return v.MetricsNames }).(pulumi.StringArrayOutput)
}

// Name. The naming rule is: namespace/name. If you don&#39;t have any namespace, use the default namespace: kube-system, otherwise use the specified one.
func (o TmpTkeBasicConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpTkeBasicConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type TmpTkeBasicConfigArrayOutput struct{ *pulumi.OutputState }

func (TmpTkeBasicConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpTkeBasicConfig)(nil)).Elem()
}

func (o TmpTkeBasicConfigArrayOutput) ToTmpTkeBasicConfigArrayOutput() TmpTkeBasicConfigArrayOutput {
	return o
}

func (o TmpTkeBasicConfigArrayOutput) ToTmpTkeBasicConfigArrayOutputWithContext(ctx context.Context) TmpTkeBasicConfigArrayOutput {
	return o
}

func (o TmpTkeBasicConfigArrayOutput) Index(i pulumi.IntInput) TmpTkeBasicConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TmpTkeBasicConfig {
		return vs[0].([]*TmpTkeBasicConfig)[vs[1].(int)]
	}).(TmpTkeBasicConfigOutput)
}

type TmpTkeBasicConfigMapOutput struct{ *pulumi.OutputState }

func (TmpTkeBasicConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpTkeBasicConfig)(nil)).Elem()
}

func (o TmpTkeBasicConfigMapOutput) ToTmpTkeBasicConfigMapOutput() TmpTkeBasicConfigMapOutput {
	return o
}

func (o TmpTkeBasicConfigMapOutput) ToTmpTkeBasicConfigMapOutputWithContext(ctx context.Context) TmpTkeBasicConfigMapOutput {
	return o
}

func (o TmpTkeBasicConfigMapOutput) MapIndex(k pulumi.StringInput) TmpTkeBasicConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TmpTkeBasicConfig {
		return vs[0].(map[string]*TmpTkeBasicConfig)[vs[1].(string)]
	}).(TmpTkeBasicConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TmpTkeBasicConfigInput)(nil)).Elem(), &TmpTkeBasicConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpTkeBasicConfigArrayInput)(nil)).Elem(), TmpTkeBasicConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpTkeBasicConfigMapInput)(nil)).Elem(), TmpTkeBasicConfigMap{})
	pulumi.RegisterOutputType(TmpTkeBasicConfigOutput{})
	pulumi.RegisterOutputType(TmpTkeBasicConfigArrayOutput{})
	pulumi.RegisterOutputType(TmpTkeBasicConfigMapOutput{})
}
