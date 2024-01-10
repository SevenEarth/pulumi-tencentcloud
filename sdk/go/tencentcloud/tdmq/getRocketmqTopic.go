// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tdmq

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tdmqRocketmq topic
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tdmq"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tdmq"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleRocketmqCluster, err := Tdmq.NewRocketmqCluster(ctx, "exampleRocketmqCluster", &Tdmq.RocketmqClusterArgs{
//				ClusterName: pulumi.String("tf_example"),
//				Remark:      pulumi.String("remark."),
//			})
//			if err != nil {
//				return err
//			}
//			exampleRocketmqNamespace, err := Tdmq.NewRocketmqNamespace(ctx, "exampleRocketmqNamespace", &Tdmq.RocketmqNamespaceArgs{
//				ClusterId:     exampleRocketmqCluster.ClusterId,
//				NamespaceName: pulumi.String("tf_example"),
//				Remark:        pulumi.String("remark."),
//			})
//			if err != nil {
//				return err
//			}
//			_ = Tdmq.GetRocketmqTopicOutput(ctx, tdmq.GetRocketmqTopicOutputArgs{
//				ClusterId:   exampleRocketmqCluster.ClusterId,
//				NamespaceId: exampleRocketmqNamespace.NamespaceName,
//				FilterName:  exampleTdmq / rocketmqTopicRocketmqTopic.TopicName,
//			}, nil)
//			_, err = Tdmq.NewRocketmqTopic(ctx, "exampleTdmq/rocketmqTopicRocketmqTopic", &Tdmq.RocketmqTopicArgs{
//				TopicName:     pulumi.String("tf_example"),
//				NamespaceName: exampleRocketmqNamespace.NamespaceName,
//				ClusterId:     exampleRocketmqCluster.ClusterId,
//				Type:          pulumi.String("Normal"),
//				Remark:        pulumi.String("remark."),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRocketmqTopic(ctx *pulumi.Context, args *LookupRocketmqTopicArgs, opts ...pulumi.InvokeOption) (*LookupRocketmqTopicResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRocketmqTopicResult
	err := ctx.Invoke("tencentcloud:Tdmq/getRocketmqTopic:getRocketmqTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRocketmqTopic.
type LookupRocketmqTopicArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Search by topic name. Fuzzy query is supported.
	FilterName *string `pulumi:"filterName"`
	// Filter by topic type. Valid values: `Normal`, `GlobalOrder`, `PartitionedOrder`, `Transaction`.
	FilterTypes []string `pulumi:"filterTypes"`
	// Namespace.
	NamespaceId string `pulumi:"namespaceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getRocketmqTopic.
type LookupRocketmqTopicResult struct {
	ClusterId   string   `pulumi:"clusterId"`
	FilterName  *string  `pulumi:"filterName"`
	FilterTypes []string `pulumi:"filterTypes"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	NamespaceId      string  `pulumi:"namespaceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// List of topic information.
	Topics []GetRocketmqTopicTopic `pulumi:"topics"`
}

func LookupRocketmqTopicOutput(ctx *pulumi.Context, args LookupRocketmqTopicOutputArgs, opts ...pulumi.InvokeOption) LookupRocketmqTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRocketmqTopicResult, error) {
			args := v.(LookupRocketmqTopicArgs)
			r, err := LookupRocketmqTopic(ctx, &args, opts...)
			var s LookupRocketmqTopicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRocketmqTopicResultOutput)
}

// A collection of arguments for invoking getRocketmqTopic.
type LookupRocketmqTopicOutputArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// Search by topic name. Fuzzy query is supported.
	FilterName pulumi.StringPtrInput `pulumi:"filterName"`
	// Filter by topic type. Valid values: `Normal`, `GlobalOrder`, `PartitionedOrder`, `Transaction`.
	FilterTypes pulumi.StringArrayInput `pulumi:"filterTypes"`
	// Namespace.
	NamespaceId pulumi.StringInput `pulumi:"namespaceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (LookupRocketmqTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRocketmqTopicArgs)(nil)).Elem()
}

// A collection of values returned by getRocketmqTopic.
type LookupRocketmqTopicResultOutput struct{ *pulumi.OutputState }

func (LookupRocketmqTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRocketmqTopicResult)(nil)).Elem()
}

func (o LookupRocketmqTopicResultOutput) ToLookupRocketmqTopicResultOutput() LookupRocketmqTopicResultOutput {
	return o
}

func (o LookupRocketmqTopicResultOutput) ToLookupRocketmqTopicResultOutputWithContext(ctx context.Context) LookupRocketmqTopicResultOutput {
	return o
}

func (o LookupRocketmqTopicResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRocketmqTopicResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o LookupRocketmqTopicResultOutput) FilterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRocketmqTopicResult) *string { return v.FilterName }).(pulumi.StringPtrOutput)
}

func (o LookupRocketmqTopicResultOutput) FilterTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRocketmqTopicResult) []string { return v.FilterTypes }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRocketmqTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRocketmqTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRocketmqTopicResultOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRocketmqTopicResult) string { return v.NamespaceId }).(pulumi.StringOutput)
}

func (o LookupRocketmqTopicResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRocketmqTopicResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// List of topic information.
func (o LookupRocketmqTopicResultOutput) Topics() GetRocketmqTopicTopicArrayOutput {
	return o.ApplyT(func(v LookupRocketmqTopicResult) []GetRocketmqTopicTopic { return v.Topics }).(GetRocketmqTopicTopicArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRocketmqTopicResultOutput{})
}
