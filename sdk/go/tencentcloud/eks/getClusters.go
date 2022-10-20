// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query elastic kubernetes cluster resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Eks"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Eks"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Eks.GetClusters(ctx, &eks.GetClustersArgs{
//				ClusterId: pulumi.StringRef("cls-xxxxxxxx"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetClusters(ctx *pulumi.Context, args *GetClustersArgs, opts ...pulumi.InvokeOption) (*GetClustersResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetClustersResult
	err := ctx.Invoke("tencentcloud:Eks/getClusters:getClusters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusters.
type GetClustersArgs struct {
	// ID of the cluster. Conflict with cluster_name, can not be set at the same time.
	ClusterId *string `pulumi:"clusterId"`
	// Name of the cluster. Conflict with cluster_id, can not be set at the same time.
	ClusterName *string `pulumi:"clusterName"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getClusters.
type GetClustersResult struct {
	// ID of the cluster.
	ClusterId *string `pulumi:"clusterId"`
	// Name of the cluster.
	ClusterName *string `pulumi:"clusterName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// EKS cluster list.
	Lists            []GetClustersList `pulumi:"lists"`
	ResultOutputFile *string           `pulumi:"resultOutputFile"`
}

func GetClustersOutput(ctx *pulumi.Context, args GetClustersOutputArgs, opts ...pulumi.InvokeOption) GetClustersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClustersResult, error) {
			args := v.(GetClustersArgs)
			r, err := GetClusters(ctx, &args, opts...)
			var s GetClustersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetClustersResultOutput)
}

// A collection of arguments for invoking getClusters.
type GetClustersOutputArgs struct {
	// ID of the cluster. Conflict with cluster_name, can not be set at the same time.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// Name of the cluster. Conflict with cluster_id, can not be set at the same time.
	ClusterName pulumi.StringPtrInput `pulumi:"clusterName"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetClustersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClustersArgs)(nil)).Elem()
}

// A collection of values returned by getClusters.
type GetClustersResultOutput struct{ *pulumi.OutputState }

func (GetClustersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClustersResult)(nil)).Elem()
}

func (o GetClustersResultOutput) ToGetClustersResultOutput() GetClustersResultOutput {
	return o
}

func (o GetClustersResultOutput) ToGetClustersResultOutputWithContext(ctx context.Context) GetClustersResultOutput {
	return o
}

// ID of the cluster.
func (o GetClustersResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

// Name of the cluster.
func (o GetClustersResultOutput) ClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.ClusterName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClustersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersResult) string { return v.Id }).(pulumi.StringOutput)
}

// EKS cluster list.
func (o GetClustersResultOutput) Lists() GetClustersListArrayOutput {
	return o.ApplyT(func(v GetClustersResult) []GetClustersList { return v.Lists }).(GetClustersListArrayOutput)
}

func (o GetClustersResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClustersResultOutput{})
}
