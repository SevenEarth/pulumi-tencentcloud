// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ckafka

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of ckafka datahubGroupOffsets
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
//			_, err := Ckafka.GetDatahubGroupOffsets(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDatahubGroupOffsets(ctx *pulumi.Context, args *GetDatahubGroupOffsetsArgs, opts ...pulumi.InvokeOption) (*GetDatahubGroupOffsetsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDatahubGroupOffsetsResult
	err := ctx.Invoke("tencentcloud:Ckafka/getDatahubGroupOffsets:getDatahubGroupOffsets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatahubGroupOffsets.
type GetDatahubGroupOffsetsArgs struct {
	// Kafka consumer group.
	Group string `pulumi:"group"`
	// topic name that the task subscribe.
	Name string `pulumi:"name"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// fuzzy match topicName.
	SearchWord *string `pulumi:"searchWord"`
}

// A collection of values returned by getDatahubGroupOffsets.
type GetDatahubGroupOffsetsResult struct {
	Group string `pulumi:"group"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	Name             string  `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	SearchWord       *string `pulumi:"searchWord"`
	// The topic array, where each element is a json object.
	TopicLists []GetDatahubGroupOffsetsTopicList `pulumi:"topicLists"`
}

func GetDatahubGroupOffsetsOutput(ctx *pulumi.Context, args GetDatahubGroupOffsetsOutputArgs, opts ...pulumi.InvokeOption) GetDatahubGroupOffsetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatahubGroupOffsetsResult, error) {
			args := v.(GetDatahubGroupOffsetsArgs)
			r, err := GetDatahubGroupOffsets(ctx, &args, opts...)
			var s GetDatahubGroupOffsetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatahubGroupOffsetsResultOutput)
}

// A collection of arguments for invoking getDatahubGroupOffsets.
type GetDatahubGroupOffsetsOutputArgs struct {
	// Kafka consumer group.
	Group pulumi.StringInput `pulumi:"group"`
	// topic name that the task subscribe.
	Name pulumi.StringInput `pulumi:"name"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// fuzzy match topicName.
	SearchWord pulumi.StringPtrInput `pulumi:"searchWord"`
}

func (GetDatahubGroupOffsetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatahubGroupOffsetsArgs)(nil)).Elem()
}

// A collection of values returned by getDatahubGroupOffsets.
type GetDatahubGroupOffsetsResultOutput struct{ *pulumi.OutputState }

func (GetDatahubGroupOffsetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatahubGroupOffsetsResult)(nil)).Elem()
}

func (o GetDatahubGroupOffsetsResultOutput) ToGetDatahubGroupOffsetsResultOutput() GetDatahubGroupOffsetsResultOutput {
	return o
}

func (o GetDatahubGroupOffsetsResultOutput) ToGetDatahubGroupOffsetsResultOutputWithContext(ctx context.Context) GetDatahubGroupOffsetsResultOutput {
	return o
}

func (o GetDatahubGroupOffsetsResultOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatahubGroupOffsetsResult) string { return v.Group }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDatahubGroupOffsetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatahubGroupOffsetsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDatahubGroupOffsetsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatahubGroupOffsetsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetDatahubGroupOffsetsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatahubGroupOffsetsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetDatahubGroupOffsetsResultOutput) SearchWord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatahubGroupOffsetsResult) *string { return v.SearchWord }).(pulumi.StringPtrOutput)
}

// The topic array, where each element is a json object.
func (o GetDatahubGroupOffsetsResultOutput) TopicLists() GetDatahubGroupOffsetsTopicListArrayOutput {
	return o.ApplyT(func(v GetDatahubGroupOffsetsResult) []GetDatahubGroupOffsetsTopicList { return v.TopicLists }).(GetDatahubGroupOffsetsTopicListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatahubGroupOffsetsResultOutput{})
}
