// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tdmq

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tdmqRocketmq cluster
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tdmq"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tdmq"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Tdmq.GetRocketmqCluster(ctx, &tdmq.GetRocketmqClusterArgs{
// 			NameKeyword: pulumi.StringRef("test_rocketmq"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupRocketmqCluster(ctx *pulumi.Context, args *LookupRocketmqClusterArgs, opts ...pulumi.InvokeOption) (*LookupRocketmqClusterResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRocketmqClusterResult
	err := ctx.Invoke("tencentcloud:Tdmq/getRocketmqCluster:getRocketmqCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRocketmqCluster.
type LookupRocketmqClusterArgs struct {
	// Filter by cluster ID.
	ClusterIdLists []string `pulumi:"clusterIdLists"`
	// Search by cluster ID.
	IdKeyword *string `pulumi:"idKeyword"`
	// Search by cluster name.
	NameKeyword *string `pulumi:"nameKeyword"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getRocketmqCluster.
type LookupRocketmqClusterResult struct {
	ClusterIdLists []string `pulumi:"clusterIdLists"`
	// Cluster information.
	ClusterLists []GetRocketmqClusterClusterList `pulumi:"clusterLists"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	IdKeyword        *string `pulumi:"idKeyword"`
	NameKeyword      *string `pulumi:"nameKeyword"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func LookupRocketmqClusterOutput(ctx *pulumi.Context, args LookupRocketmqClusterOutputArgs, opts ...pulumi.InvokeOption) LookupRocketmqClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRocketmqClusterResult, error) {
			args := v.(LookupRocketmqClusterArgs)
			r, err := LookupRocketmqCluster(ctx, &args, opts...)
			var s LookupRocketmqClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRocketmqClusterResultOutput)
}

// A collection of arguments for invoking getRocketmqCluster.
type LookupRocketmqClusterOutputArgs struct {
	// Filter by cluster ID.
	ClusterIdLists pulumi.StringArrayInput `pulumi:"clusterIdLists"`
	// Search by cluster ID.
	IdKeyword pulumi.StringPtrInput `pulumi:"idKeyword"`
	// Search by cluster name.
	NameKeyword pulumi.StringPtrInput `pulumi:"nameKeyword"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (LookupRocketmqClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRocketmqClusterArgs)(nil)).Elem()
}

// A collection of values returned by getRocketmqCluster.
type LookupRocketmqClusterResultOutput struct{ *pulumi.OutputState }

func (LookupRocketmqClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRocketmqClusterResult)(nil)).Elem()
}

func (o LookupRocketmqClusterResultOutput) ToLookupRocketmqClusterResultOutput() LookupRocketmqClusterResultOutput {
	return o
}

func (o LookupRocketmqClusterResultOutput) ToLookupRocketmqClusterResultOutputWithContext(ctx context.Context) LookupRocketmqClusterResultOutput {
	return o
}

func (o LookupRocketmqClusterResultOutput) ClusterIdLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRocketmqClusterResult) []string { return v.ClusterIdLists }).(pulumi.StringArrayOutput)
}

// Cluster information.
func (o LookupRocketmqClusterResultOutput) ClusterLists() GetRocketmqClusterClusterListArrayOutput {
	return o.ApplyT(func(v LookupRocketmqClusterResult) []GetRocketmqClusterClusterList { return v.ClusterLists }).(GetRocketmqClusterClusterListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRocketmqClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRocketmqClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRocketmqClusterResultOutput) IdKeyword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRocketmqClusterResult) *string { return v.IdKeyword }).(pulumi.StringPtrOutput)
}

func (o LookupRocketmqClusterResultOutput) NameKeyword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRocketmqClusterResult) *string { return v.NameKeyword }).(pulumi.StringPtrOutput)
}

func (o LookupRocketmqClusterResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRocketmqClusterResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRocketmqClusterResultOutput{})
}
