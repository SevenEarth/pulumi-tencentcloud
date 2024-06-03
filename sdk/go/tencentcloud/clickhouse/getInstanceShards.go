// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clickhouse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of clickhouse instanceShards
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Clickhouse"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Clickhouse.GetInstanceShards(ctx, &clickhouse.GetInstanceShardsArgs{
//				InstanceId: "cdwch-datuhk3z",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetInstanceShards(ctx *pulumi.Context, args *GetInstanceShardsArgs, opts ...pulumi.InvokeOption) (*GetInstanceShardsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInstanceShardsResult
	err := ctx.Invoke("tencentcloud:Clickhouse/getInstanceShards:getInstanceShards", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceShards.
type GetInstanceShardsArgs struct {
	// Cluster instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getInstanceShards.
type GetInstanceShardsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// Instance shard information.
	InstanceShardsList string  `pulumi:"instanceShardsList"`
	ResultOutputFile   *string `pulumi:"resultOutputFile"`
}

func GetInstanceShardsOutput(ctx *pulumi.Context, args GetInstanceShardsOutputArgs, opts ...pulumi.InvokeOption) GetInstanceShardsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceShardsResult, error) {
			args := v.(GetInstanceShardsArgs)
			r, err := GetInstanceShards(ctx, &args, opts...)
			var s GetInstanceShardsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceShardsResultOutput)
}

// A collection of arguments for invoking getInstanceShards.
type GetInstanceShardsOutputArgs struct {
	// Cluster instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetInstanceShardsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceShardsArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceShards.
type GetInstanceShardsResultOutput struct{ *pulumi.OutputState }

func (GetInstanceShardsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceShardsResult)(nil)).Elem()
}

func (o GetInstanceShardsResultOutput) ToGetInstanceShardsResultOutput() GetInstanceShardsResultOutput {
	return o
}

func (o GetInstanceShardsResultOutput) ToGetInstanceShardsResultOutputWithContext(ctx context.Context) GetInstanceShardsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceShardsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceShardsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstanceShardsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceShardsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Instance shard information.
func (o GetInstanceShardsResultOutput) InstanceShardsList() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceShardsResult) string { return v.InstanceShardsList }).(pulumi.StringOutput)
}

func (o GetInstanceShardsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceShardsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceShardsResultOutput{})
}
