// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tdmq

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tdmqRocketmq namespace
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
// 		exampleRocketmqCluster, err := Tdmq.NewRocketmqCluster(ctx, "exampleRocketmqCluster", &Tdmq.RocketmqClusterArgs{
// 			ClusterName: pulumi.String("tf_example"),
// 			Remark:      pulumi.String("remark."),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_ = Tdmq.GetRocketmqNamespaceOutput(ctx, tdmq.GetRocketmqNamespaceOutputArgs{
// 			ClusterId:   exampleRocketmqCluster.ClusterId,
// 			NameKeyword: exampleTdmq / rocketmqNamespaceRocketmqNamespace.NamespaceName,
// 		}, nil)
// 		_, err = Tdmq.NewRocketmqNamespace(ctx, "exampleTdmq/rocketmqNamespaceRocketmqNamespace", &Tdmq.RocketmqNamespaceArgs{
// 			ClusterId:     exampleRocketmqCluster.ClusterId,
// 			NamespaceName: pulumi.String("tf_example"),
// 			Remark:        pulumi.String("remark."),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupRocketmqNamespace(ctx *pulumi.Context, args *LookupRocketmqNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupRocketmqNamespaceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRocketmqNamespaceResult
	err := ctx.Invoke("tencentcloud:Tdmq/getRocketmqNamespace:getRocketmqNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRocketmqNamespace.
type LookupRocketmqNamespaceArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Search by name.
	NameKeyword *string `pulumi:"nameKeyword"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getRocketmqNamespace.
type LookupRocketmqNamespaceResult struct {
	ClusterId string `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	NameKeyword *string `pulumi:"nameKeyword"`
	// List of namespaces.
	Namespaces       []GetRocketmqNamespaceNamespace `pulumi:"namespaces"`
	ResultOutputFile *string                         `pulumi:"resultOutputFile"`
}

func LookupRocketmqNamespaceOutput(ctx *pulumi.Context, args LookupRocketmqNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupRocketmqNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRocketmqNamespaceResult, error) {
			args := v.(LookupRocketmqNamespaceArgs)
			r, err := LookupRocketmqNamespace(ctx, &args, opts...)
			var s LookupRocketmqNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRocketmqNamespaceResultOutput)
}

// A collection of arguments for invoking getRocketmqNamespace.
type LookupRocketmqNamespaceOutputArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// Search by name.
	NameKeyword pulumi.StringPtrInput `pulumi:"nameKeyword"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (LookupRocketmqNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRocketmqNamespaceArgs)(nil)).Elem()
}

// A collection of values returned by getRocketmqNamespace.
type LookupRocketmqNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupRocketmqNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRocketmqNamespaceResult)(nil)).Elem()
}

func (o LookupRocketmqNamespaceResultOutput) ToLookupRocketmqNamespaceResultOutput() LookupRocketmqNamespaceResultOutput {
	return o
}

func (o LookupRocketmqNamespaceResultOutput) ToLookupRocketmqNamespaceResultOutputWithContext(ctx context.Context) LookupRocketmqNamespaceResultOutput {
	return o
}

func (o LookupRocketmqNamespaceResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRocketmqNamespaceResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRocketmqNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRocketmqNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRocketmqNamespaceResultOutput) NameKeyword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRocketmqNamespaceResult) *string { return v.NameKeyword }).(pulumi.StringPtrOutput)
}

// List of namespaces.
func (o LookupRocketmqNamespaceResultOutput) Namespaces() GetRocketmqNamespaceNamespaceArrayOutput {
	return o.ApplyT(func(v LookupRocketmqNamespaceResult) []GetRocketmqNamespaceNamespace { return v.Namespaces }).(GetRocketmqNamespaceNamespaceArrayOutput)
}

func (o LookupRocketmqNamespaceResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRocketmqNamespaceResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRocketmqNamespaceResultOutput{})
}
