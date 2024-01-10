// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tdmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tdmqRocketmq cluster
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tdmq"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tdmq.NewRocketmqCluster(ctx, "example", &Tdmq.RocketmqClusterArgs{
//				ClusterName: pulumi.String("tf_example"),
//				Remark:      pulumi.String("remark."),
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
// tdmqRocketmq cluster can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Tdmq/rocketmqCluster:RocketmqCluster cluster cluster_id
//
// ```
type RocketmqCluster struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Creation time in milliseconds.
	CreateTime pulumi.IntOutput `pulumi:"createTime"`
	// Whether it is an exclusive instance.
	IsVip pulumi.BoolOutput `pulumi:"isVip"`
	// Public network access address.
	PublicEndPoint pulumi.StringOutput `pulumi:"publicEndPoint"`
	// Region information.
	Region pulumi.StringOutput `pulumi:"region"`
	// Cluster description (up to 128 characters).
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// Rocketmq cluster identification.
	RocketMQFlag pulumi.BoolOutput `pulumi:"rocketMQFlag"`
	// Whether the namespace access point is supported.
	SupportNamespaceEndpoint pulumi.BoolOutput `pulumi:"supportNamespaceEndpoint"`
	// VPC access address.
	VpcEndPoint pulumi.StringOutput `pulumi:"vpcEndPoint"`
	// Vpc list.
	Vpcs RocketmqClusterVpcArrayOutput `pulumi:"vpcs"`
}

// NewRocketmqCluster registers a new resource with the given unique name, arguments, and options.
func NewRocketmqCluster(ctx *pulumi.Context,
	name string, args *RocketmqClusterArgs, opts ...pulumi.ResourceOption) (*RocketmqCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RocketmqCluster
	err := ctx.RegisterResource("tencentcloud:Tdmq/rocketmqCluster:RocketmqCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRocketmqCluster gets an existing RocketmqCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRocketmqCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RocketmqClusterState, opts ...pulumi.ResourceOption) (*RocketmqCluster, error) {
	var resource RocketmqCluster
	err := ctx.ReadResource("tencentcloud:Tdmq/rocketmqCluster:RocketmqCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RocketmqCluster resources.
type rocketmqClusterState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
	ClusterName *string `pulumi:"clusterName"`
	// Creation time in milliseconds.
	CreateTime *int `pulumi:"createTime"`
	// Whether it is an exclusive instance.
	IsVip *bool `pulumi:"isVip"`
	// Public network access address.
	PublicEndPoint *string `pulumi:"publicEndPoint"`
	// Region information.
	Region *string `pulumi:"region"`
	// Cluster description (up to 128 characters).
	Remark *string `pulumi:"remark"`
	// Rocketmq cluster identification.
	RocketMQFlag *bool `pulumi:"rocketMQFlag"`
	// Whether the namespace access point is supported.
	SupportNamespaceEndpoint *bool `pulumi:"supportNamespaceEndpoint"`
	// VPC access address.
	VpcEndPoint *string `pulumi:"vpcEndPoint"`
	// Vpc list.
	Vpcs []RocketmqClusterVpc `pulumi:"vpcs"`
}

type RocketmqClusterState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
	ClusterName pulumi.StringPtrInput
	// Creation time in milliseconds.
	CreateTime pulumi.IntPtrInput
	// Whether it is an exclusive instance.
	IsVip pulumi.BoolPtrInput
	// Public network access address.
	PublicEndPoint pulumi.StringPtrInput
	// Region information.
	Region pulumi.StringPtrInput
	// Cluster description (up to 128 characters).
	Remark pulumi.StringPtrInput
	// Rocketmq cluster identification.
	RocketMQFlag pulumi.BoolPtrInput
	// Whether the namespace access point is supported.
	SupportNamespaceEndpoint pulumi.BoolPtrInput
	// VPC access address.
	VpcEndPoint pulumi.StringPtrInput
	// Vpc list.
	Vpcs RocketmqClusterVpcArrayInput
}

func (RocketmqClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*rocketmqClusterState)(nil)).Elem()
}

type rocketmqClusterArgs struct {
	// Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
	ClusterName string `pulumi:"clusterName"`
	// Cluster description (up to 128 characters).
	Remark *string `pulumi:"remark"`
}

// The set of arguments for constructing a RocketmqCluster resource.
type RocketmqClusterArgs struct {
	// Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
	ClusterName pulumi.StringInput
	// Cluster description (up to 128 characters).
	Remark pulumi.StringPtrInput
}

func (RocketmqClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rocketmqClusterArgs)(nil)).Elem()
}

type RocketmqClusterInput interface {
	pulumi.Input

	ToRocketmqClusterOutput() RocketmqClusterOutput
	ToRocketmqClusterOutputWithContext(ctx context.Context) RocketmqClusterOutput
}

func (*RocketmqCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**RocketmqCluster)(nil)).Elem()
}

func (i *RocketmqCluster) ToRocketmqClusterOutput() RocketmqClusterOutput {
	return i.ToRocketmqClusterOutputWithContext(context.Background())
}

func (i *RocketmqCluster) ToRocketmqClusterOutputWithContext(ctx context.Context) RocketmqClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RocketmqClusterOutput)
}

// RocketmqClusterArrayInput is an input type that accepts RocketmqClusterArray and RocketmqClusterArrayOutput values.
// You can construct a concrete instance of `RocketmqClusterArrayInput` via:
//
//	RocketmqClusterArray{ RocketmqClusterArgs{...} }
type RocketmqClusterArrayInput interface {
	pulumi.Input

	ToRocketmqClusterArrayOutput() RocketmqClusterArrayOutput
	ToRocketmqClusterArrayOutputWithContext(context.Context) RocketmqClusterArrayOutput
}

type RocketmqClusterArray []RocketmqClusterInput

func (RocketmqClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RocketmqCluster)(nil)).Elem()
}

func (i RocketmqClusterArray) ToRocketmqClusterArrayOutput() RocketmqClusterArrayOutput {
	return i.ToRocketmqClusterArrayOutputWithContext(context.Background())
}

func (i RocketmqClusterArray) ToRocketmqClusterArrayOutputWithContext(ctx context.Context) RocketmqClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RocketmqClusterArrayOutput)
}

// RocketmqClusterMapInput is an input type that accepts RocketmqClusterMap and RocketmqClusterMapOutput values.
// You can construct a concrete instance of `RocketmqClusterMapInput` via:
//
//	RocketmqClusterMap{ "key": RocketmqClusterArgs{...} }
type RocketmqClusterMapInput interface {
	pulumi.Input

	ToRocketmqClusterMapOutput() RocketmqClusterMapOutput
	ToRocketmqClusterMapOutputWithContext(context.Context) RocketmqClusterMapOutput
}

type RocketmqClusterMap map[string]RocketmqClusterInput

func (RocketmqClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RocketmqCluster)(nil)).Elem()
}

func (i RocketmqClusterMap) ToRocketmqClusterMapOutput() RocketmqClusterMapOutput {
	return i.ToRocketmqClusterMapOutputWithContext(context.Background())
}

func (i RocketmqClusterMap) ToRocketmqClusterMapOutputWithContext(ctx context.Context) RocketmqClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RocketmqClusterMapOutput)
}

type RocketmqClusterOutput struct{ *pulumi.OutputState }

func (RocketmqClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RocketmqCluster)(nil)).Elem()
}

func (o RocketmqClusterOutput) ToRocketmqClusterOutput() RocketmqClusterOutput {
	return o
}

func (o RocketmqClusterOutput) ToRocketmqClusterOutputWithContext(ctx context.Context) RocketmqClusterOutput {
	return o
}

// Cluster ID.
func (o RocketmqClusterOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqCluster) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
func (o RocketmqClusterOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqCluster) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// Creation time in milliseconds.
func (o RocketmqClusterOutput) CreateTime() pulumi.IntOutput {
	return o.ApplyT(func(v *RocketmqCluster) pulumi.IntOutput { return v.CreateTime }).(pulumi.IntOutput)
}

// Whether it is an exclusive instance.
func (o RocketmqClusterOutput) IsVip() pulumi.BoolOutput {
	return o.ApplyT(func(v *RocketmqCluster) pulumi.BoolOutput { return v.IsVip }).(pulumi.BoolOutput)
}

// Public network access address.
func (o RocketmqClusterOutput) PublicEndPoint() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqCluster) pulumi.StringOutput { return v.PublicEndPoint }).(pulumi.StringOutput)
}

// Region information.
func (o RocketmqClusterOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqCluster) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Cluster description (up to 128 characters).
func (o RocketmqClusterOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RocketmqCluster) pulumi.StringPtrOutput { return v.Remark }).(pulumi.StringPtrOutput)
}

// Rocketmq cluster identification.
func (o RocketmqClusterOutput) RocketMQFlag() pulumi.BoolOutput {
	return o.ApplyT(func(v *RocketmqCluster) pulumi.BoolOutput { return v.RocketMQFlag }).(pulumi.BoolOutput)
}

// Whether the namespace access point is supported.
func (o RocketmqClusterOutput) SupportNamespaceEndpoint() pulumi.BoolOutput {
	return o.ApplyT(func(v *RocketmqCluster) pulumi.BoolOutput { return v.SupportNamespaceEndpoint }).(pulumi.BoolOutput)
}

// VPC access address.
func (o RocketmqClusterOutput) VpcEndPoint() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqCluster) pulumi.StringOutput { return v.VpcEndPoint }).(pulumi.StringOutput)
}

// Vpc list.
func (o RocketmqClusterOutput) Vpcs() RocketmqClusterVpcArrayOutput {
	return o.ApplyT(func(v *RocketmqCluster) RocketmqClusterVpcArrayOutput { return v.Vpcs }).(RocketmqClusterVpcArrayOutput)
}

type RocketmqClusterArrayOutput struct{ *pulumi.OutputState }

func (RocketmqClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RocketmqCluster)(nil)).Elem()
}

func (o RocketmqClusterArrayOutput) ToRocketmqClusterArrayOutput() RocketmqClusterArrayOutput {
	return o
}

func (o RocketmqClusterArrayOutput) ToRocketmqClusterArrayOutputWithContext(ctx context.Context) RocketmqClusterArrayOutput {
	return o
}

func (o RocketmqClusterArrayOutput) Index(i pulumi.IntInput) RocketmqClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RocketmqCluster {
		return vs[0].([]*RocketmqCluster)[vs[1].(int)]
	}).(RocketmqClusterOutput)
}

type RocketmqClusterMapOutput struct{ *pulumi.OutputState }

func (RocketmqClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RocketmqCluster)(nil)).Elem()
}

func (o RocketmqClusterMapOutput) ToRocketmqClusterMapOutput() RocketmqClusterMapOutput {
	return o
}

func (o RocketmqClusterMapOutput) ToRocketmqClusterMapOutputWithContext(ctx context.Context) RocketmqClusterMapOutput {
	return o
}

func (o RocketmqClusterMapOutput) MapIndex(k pulumi.StringInput) RocketmqClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RocketmqCluster {
		return vs[0].(map[string]*RocketmqCluster)[vs[1].(string)]
	}).(RocketmqClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RocketmqClusterInput)(nil)).Elem(), &RocketmqCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*RocketmqClusterArrayInput)(nil)).Elem(), RocketmqClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RocketmqClusterMapInput)(nil)).Elem(), RocketmqClusterMap{})
	pulumi.RegisterOutputType(RocketmqClusterOutput{})
	pulumi.RegisterOutputType(RocketmqClusterArrayOutput{})
	pulumi.RegisterOutputType(RocketmqClusterMapOutput{})
}
