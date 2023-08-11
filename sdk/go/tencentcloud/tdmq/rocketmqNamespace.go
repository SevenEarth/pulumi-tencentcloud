// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tdmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tdmqRocketmq namespace
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tdmq"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleRocketmqCluster, err := Tdmq.NewRocketmqCluster(ctx, "exampleRocketmqCluster", &Tdmq.RocketmqClusterArgs{
// 			ClusterName: pulumi.String("tf_example"),
// 			Remark:      pulumi.String("remark."),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Tdmq.NewRocketmqNamespace(ctx, "exampleRocketmqNamespace", &Tdmq.RocketmqNamespaceArgs{
// 			ClusterId:     exampleRocketmqCluster.ClusterId,
// 			NamespaceName: pulumi.String("tf_example_namespace"),
// 			Remark:        pulumi.String("remark."),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// tdmqRocketmq namespace can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Tdmq/rocketmqNamespace:RocketmqNamespace namespace namespace_id
// ```
type RocketmqNamespace struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Namespace name, which can contain 3-64 letters, digits, hyphens, and underscores.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// Public network access point address.
	PublicEndpoint pulumi.StringOutput `pulumi:"publicEndpoint"`
	// Remarks (up to 128 characters).
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored. Retention time of persisted messages in milliseconds.
	//
	// Deprecated: It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored.
	RetentionTime pulumi.IntPtrOutput `pulumi:"retentionTime"`
	// It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored. Retention time of unconsumed messages in milliseconds. Value range: 60 seconds-15 days.
	//
	// Deprecated: It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored.
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
	// VPC access point address.
	VpcEndpoint pulumi.StringOutput `pulumi:"vpcEndpoint"`
}

// NewRocketmqNamespace registers a new resource with the given unique name, arguments, and options.
func NewRocketmqNamespace(ctx *pulumi.Context,
	name string, args *RocketmqNamespaceArgs, opts ...pulumi.ResourceOption) (*RocketmqNamespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RocketmqNamespace
	err := ctx.RegisterResource("tencentcloud:Tdmq/rocketmqNamespace:RocketmqNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRocketmqNamespace gets an existing RocketmqNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRocketmqNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RocketmqNamespaceState, opts ...pulumi.ResourceOption) (*RocketmqNamespace, error) {
	var resource RocketmqNamespace
	err := ctx.ReadResource("tencentcloud:Tdmq/rocketmqNamespace:RocketmqNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RocketmqNamespace resources.
type rocketmqNamespaceState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Namespace name, which can contain 3-64 letters, digits, hyphens, and underscores.
	NamespaceName *string `pulumi:"namespaceName"`
	// Public network access point address.
	PublicEndpoint *string `pulumi:"publicEndpoint"`
	// Remarks (up to 128 characters).
	Remark *string `pulumi:"remark"`
	// It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored. Retention time of persisted messages in milliseconds.
	//
	// Deprecated: It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored.
	RetentionTime *int `pulumi:"retentionTime"`
	// It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored. Retention time of unconsumed messages in milliseconds. Value range: 60 seconds-15 days.
	//
	// Deprecated: It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored.
	Ttl *int `pulumi:"ttl"`
	// VPC access point address.
	VpcEndpoint *string `pulumi:"vpcEndpoint"`
}

type RocketmqNamespaceState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Namespace name, which can contain 3-64 letters, digits, hyphens, and underscores.
	NamespaceName pulumi.StringPtrInput
	// Public network access point address.
	PublicEndpoint pulumi.StringPtrInput
	// Remarks (up to 128 characters).
	Remark pulumi.StringPtrInput
	// It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored. Retention time of persisted messages in milliseconds.
	//
	// Deprecated: It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored.
	RetentionTime pulumi.IntPtrInput
	// It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored. Retention time of unconsumed messages in milliseconds. Value range: 60 seconds-15 days.
	//
	// Deprecated: It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored.
	Ttl pulumi.IntPtrInput
	// VPC access point address.
	VpcEndpoint pulumi.StringPtrInput
}

func (RocketmqNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*rocketmqNamespaceState)(nil)).Elem()
}

type rocketmqNamespaceArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Namespace name, which can contain 3-64 letters, digits, hyphens, and underscores.
	NamespaceName string `pulumi:"namespaceName"`
	// Remarks (up to 128 characters).
	Remark *string `pulumi:"remark"`
	// It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored. Retention time of persisted messages in milliseconds.
	//
	// Deprecated: It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored.
	RetentionTime *int `pulumi:"retentionTime"`
	// It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored. Retention time of unconsumed messages in milliseconds. Value range: 60 seconds-15 days.
	//
	// Deprecated: It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored.
	Ttl *int `pulumi:"ttl"`
}

// The set of arguments for constructing a RocketmqNamespace resource.
type RocketmqNamespaceArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Namespace name, which can contain 3-64 letters, digits, hyphens, and underscores.
	NamespaceName pulumi.StringInput
	// Remarks (up to 128 characters).
	Remark pulumi.StringPtrInput
	// It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored. Retention time of persisted messages in milliseconds.
	//
	// Deprecated: It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored.
	RetentionTime pulumi.IntPtrInput
	// It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored. Retention time of unconsumed messages in milliseconds. Value range: 60 seconds-15 days.
	//
	// Deprecated: It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored.
	Ttl pulumi.IntPtrInput
}

func (RocketmqNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rocketmqNamespaceArgs)(nil)).Elem()
}

type RocketmqNamespaceInput interface {
	pulumi.Input

	ToRocketmqNamespaceOutput() RocketmqNamespaceOutput
	ToRocketmqNamespaceOutputWithContext(ctx context.Context) RocketmqNamespaceOutput
}

func (*RocketmqNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((**RocketmqNamespace)(nil)).Elem()
}

func (i *RocketmqNamespace) ToRocketmqNamespaceOutput() RocketmqNamespaceOutput {
	return i.ToRocketmqNamespaceOutputWithContext(context.Background())
}

func (i *RocketmqNamespace) ToRocketmqNamespaceOutputWithContext(ctx context.Context) RocketmqNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RocketmqNamespaceOutput)
}

// RocketmqNamespaceArrayInput is an input type that accepts RocketmqNamespaceArray and RocketmqNamespaceArrayOutput values.
// You can construct a concrete instance of `RocketmqNamespaceArrayInput` via:
//
//          RocketmqNamespaceArray{ RocketmqNamespaceArgs{...} }
type RocketmqNamespaceArrayInput interface {
	pulumi.Input

	ToRocketmqNamespaceArrayOutput() RocketmqNamespaceArrayOutput
	ToRocketmqNamespaceArrayOutputWithContext(context.Context) RocketmqNamespaceArrayOutput
}

type RocketmqNamespaceArray []RocketmqNamespaceInput

func (RocketmqNamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RocketmqNamespace)(nil)).Elem()
}

func (i RocketmqNamespaceArray) ToRocketmqNamespaceArrayOutput() RocketmqNamespaceArrayOutput {
	return i.ToRocketmqNamespaceArrayOutputWithContext(context.Background())
}

func (i RocketmqNamespaceArray) ToRocketmqNamespaceArrayOutputWithContext(ctx context.Context) RocketmqNamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RocketmqNamespaceArrayOutput)
}

// RocketmqNamespaceMapInput is an input type that accepts RocketmqNamespaceMap and RocketmqNamespaceMapOutput values.
// You can construct a concrete instance of `RocketmqNamespaceMapInput` via:
//
//          RocketmqNamespaceMap{ "key": RocketmqNamespaceArgs{...} }
type RocketmqNamespaceMapInput interface {
	pulumi.Input

	ToRocketmqNamespaceMapOutput() RocketmqNamespaceMapOutput
	ToRocketmqNamespaceMapOutputWithContext(context.Context) RocketmqNamespaceMapOutput
}

type RocketmqNamespaceMap map[string]RocketmqNamespaceInput

func (RocketmqNamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RocketmqNamespace)(nil)).Elem()
}

func (i RocketmqNamespaceMap) ToRocketmqNamespaceMapOutput() RocketmqNamespaceMapOutput {
	return i.ToRocketmqNamespaceMapOutputWithContext(context.Background())
}

func (i RocketmqNamespaceMap) ToRocketmqNamespaceMapOutputWithContext(ctx context.Context) RocketmqNamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RocketmqNamespaceMapOutput)
}

type RocketmqNamespaceOutput struct{ *pulumi.OutputState }

func (RocketmqNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RocketmqNamespace)(nil)).Elem()
}

func (o RocketmqNamespaceOutput) ToRocketmqNamespaceOutput() RocketmqNamespaceOutput {
	return o
}

func (o RocketmqNamespaceOutput) ToRocketmqNamespaceOutputWithContext(ctx context.Context) RocketmqNamespaceOutput {
	return o
}

// Cluster ID.
func (o RocketmqNamespaceOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqNamespace) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Namespace name, which can contain 3-64 letters, digits, hyphens, and underscores.
func (o RocketmqNamespaceOutput) NamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqNamespace) pulumi.StringOutput { return v.NamespaceName }).(pulumi.StringOutput)
}

// Public network access point address.
func (o RocketmqNamespaceOutput) PublicEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqNamespace) pulumi.StringOutput { return v.PublicEndpoint }).(pulumi.StringOutput)
}

// Remarks (up to 128 characters).
func (o RocketmqNamespaceOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RocketmqNamespace) pulumi.StringPtrOutput { return v.Remark }).(pulumi.StringPtrOutput)
}

// It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored. Retention time of persisted messages in milliseconds.
//
// Deprecated: It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored.
func (o RocketmqNamespaceOutput) RetentionTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RocketmqNamespace) pulumi.IntPtrOutput { return v.RetentionTime }).(pulumi.IntPtrOutput)
}

// It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored. Retention time of unconsumed messages in milliseconds. Value range: 60 seconds-15 days.
//
// Deprecated: It has been deprecated from version 1.81.20. Due to the adjustment of RocketMQ, the creation or modification of this parameter will be ignored.
func (o RocketmqNamespaceOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RocketmqNamespace) pulumi.IntPtrOutput { return v.Ttl }).(pulumi.IntPtrOutput)
}

// VPC access point address.
func (o RocketmqNamespaceOutput) VpcEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqNamespace) pulumi.StringOutput { return v.VpcEndpoint }).(pulumi.StringOutput)
}

type RocketmqNamespaceArrayOutput struct{ *pulumi.OutputState }

func (RocketmqNamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RocketmqNamespace)(nil)).Elem()
}

func (o RocketmqNamespaceArrayOutput) ToRocketmqNamespaceArrayOutput() RocketmqNamespaceArrayOutput {
	return o
}

func (o RocketmqNamespaceArrayOutput) ToRocketmqNamespaceArrayOutputWithContext(ctx context.Context) RocketmqNamespaceArrayOutput {
	return o
}

func (o RocketmqNamespaceArrayOutput) Index(i pulumi.IntInput) RocketmqNamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RocketmqNamespace {
		return vs[0].([]*RocketmqNamespace)[vs[1].(int)]
	}).(RocketmqNamespaceOutput)
}

type RocketmqNamespaceMapOutput struct{ *pulumi.OutputState }

func (RocketmqNamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RocketmqNamespace)(nil)).Elem()
}

func (o RocketmqNamespaceMapOutput) ToRocketmqNamespaceMapOutput() RocketmqNamespaceMapOutput {
	return o
}

func (o RocketmqNamespaceMapOutput) ToRocketmqNamespaceMapOutputWithContext(ctx context.Context) RocketmqNamespaceMapOutput {
	return o
}

func (o RocketmqNamespaceMapOutput) MapIndex(k pulumi.StringInput) RocketmqNamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RocketmqNamespace {
		return vs[0].(map[string]*RocketmqNamespace)[vs[1].(string)]
	}).(RocketmqNamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RocketmqNamespaceInput)(nil)).Elem(), &RocketmqNamespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*RocketmqNamespaceArrayInput)(nil)).Elem(), RocketmqNamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RocketmqNamespaceMapInput)(nil)).Elem(), RocketmqNamespaceMap{})
	pulumi.RegisterOutputType(RocketmqNamespaceOutput{})
	pulumi.RegisterOutputType(RocketmqNamespaceArrayOutput{})
	pulumi.RegisterOutputType(RocketmqNamespaceMapOutput{})
}
