// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ckafka

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of ckafka topicProduceConnection
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Ckafka"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ckafka"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Ckafka.GetTopicProduceConnection(ctx, &ckafka.GetTopicProduceConnectionArgs{
//				InstanceId: "ckafka-xxxxxx",
//				TopicName:  "topic-xxxxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTopicProduceConnection(ctx *pulumi.Context, args *LookupTopicProduceConnectionArgs, opts ...pulumi.InvokeOption) (*LookupTopicProduceConnectionResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupTopicProduceConnectionResult
	err := ctx.Invoke("tencentcloud:Ckafka/getTopicProduceConnection:getTopicProduceConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTopicProduceConnection.
type LookupTopicProduceConnectionArgs struct {
	// InstanceId.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// TopicName.
	TopicName string `pulumi:"topicName"`
}

// A collection of values returned by getTopicProduceConnection.
type LookupTopicProduceConnectionResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	InstanceId       string  `pulumi:"instanceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// link information return result set.
	Results   []GetTopicProduceConnectionResult `pulumi:"results"`
	TopicName string                            `pulumi:"topicName"`
}

func LookupTopicProduceConnectionOutput(ctx *pulumi.Context, args LookupTopicProduceConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupTopicProduceConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTopicProduceConnectionResult, error) {
			args := v.(LookupTopicProduceConnectionArgs)
			r, err := LookupTopicProduceConnection(ctx, &args, opts...)
			var s LookupTopicProduceConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTopicProduceConnectionResultOutput)
}

// A collection of arguments for invoking getTopicProduceConnection.
type LookupTopicProduceConnectionOutputArgs struct {
	// InstanceId.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// TopicName.
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

func (LookupTopicProduceConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicProduceConnectionArgs)(nil)).Elem()
}

// A collection of values returned by getTopicProduceConnection.
type LookupTopicProduceConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupTopicProduceConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicProduceConnectionResult)(nil)).Elem()
}

func (o LookupTopicProduceConnectionResultOutput) ToLookupTopicProduceConnectionResultOutput() LookupTopicProduceConnectionResultOutput {
	return o
}

func (o LookupTopicProduceConnectionResultOutput) ToLookupTopicProduceConnectionResultOutputWithContext(ctx context.Context) LookupTopicProduceConnectionResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTopicProduceConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicProduceConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTopicProduceConnectionResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicProduceConnectionResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o LookupTopicProduceConnectionResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicProduceConnectionResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// link information return result set.
func (o LookupTopicProduceConnectionResultOutput) Results() GetTopicProduceConnectionResultArrayOutput {
	return o.ApplyT(func(v LookupTopicProduceConnectionResult) []GetTopicProduceConnectionResult { return v.Results }).(GetTopicProduceConnectionResultArrayOutput)
}

func (o LookupTopicProduceConnectionResultOutput) TopicName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicProduceConnectionResult) string { return v.TopicName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTopicProduceConnectionResultOutput{})
}
