// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oceanus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of oceanus systemResource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Oceanus"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Oceanus"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Oceanus.GetSystemResource(ctx, &oceanus.GetSystemResourceArgs{
//				ClusterId: pulumi.StringRef("cluster-n8yaia0p"),
//				Filters: []oceanus.GetSystemResourceFilter{
//					oceanus.GetSystemResourceFilter{
//						Name: "Name",
//						Values: []string{
//							"tf_example",
//						},
//					},
//				},
//				FlinkVersion: pulumi.StringRef("Flink-1.11"),
//				ResourceIds: []string{
//					"resource-abd503yt",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSystemResource(ctx *pulumi.Context, args *GetSystemResourceArgs, opts ...pulumi.InvokeOption) (*GetSystemResourceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemResourceResult
	err := ctx.Invoke("tencentcloud:Oceanus/getSystemResource:getSystemResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemResource.
type GetSystemResourceArgs struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Query the resource configuration list. If not specified, return all job configuration lists under ResourceIds.N.
	Filters []GetSystemResourceFilter `pulumi:"filters"`
	// Query built-in connectors for the corresponding Flink version.
	FlinkVersion *string `pulumi:"flinkVersion"`
	// Array of resource IDs to be queried.
	ResourceIds []string `pulumi:"resourceIds"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getSystemResource.
type GetSystemResourceResult struct {
	ClusterId    *string                   `pulumi:"clusterId"`
	Filters      []GetSystemResourceFilter `pulumi:"filters"`
	FlinkVersion *string                   `pulumi:"flinkVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id          string   `pulumi:"id"`
	ResourceIds []string `pulumi:"resourceIds"`
	// Collection of resource details.
	ResourceSets     []GetSystemResourceResourceSet `pulumi:"resourceSets"`
	ResultOutputFile *string                        `pulumi:"resultOutputFile"`
}

func GetSystemResourceOutput(ctx *pulumi.Context, args GetSystemResourceOutputArgs, opts ...pulumi.InvokeOption) GetSystemResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemResourceResult, error) {
			args := v.(GetSystemResourceArgs)
			r, err := GetSystemResource(ctx, &args, opts...)
			var s GetSystemResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemResourceResultOutput)
}

// A collection of arguments for invoking getSystemResource.
type GetSystemResourceOutputArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// Query the resource configuration list. If not specified, return all job configuration lists under ResourceIds.N.
	Filters GetSystemResourceFilterArrayInput `pulumi:"filters"`
	// Query built-in connectors for the corresponding Flink version.
	FlinkVersion pulumi.StringPtrInput `pulumi:"flinkVersion"`
	// Array of resource IDs to be queried.
	ResourceIds pulumi.StringArrayInput `pulumi:"resourceIds"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetSystemResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemResourceArgs)(nil)).Elem()
}

// A collection of values returned by getSystemResource.
type GetSystemResourceResultOutput struct{ *pulumi.OutputState }

func (GetSystemResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemResourceResult)(nil)).Elem()
}

func (o GetSystemResourceResultOutput) ToGetSystemResourceResultOutput() GetSystemResourceResultOutput {
	return o
}

func (o GetSystemResourceResultOutput) ToGetSystemResourceResultOutputWithContext(ctx context.Context) GetSystemResourceResultOutput {
	return o
}

func (o GetSystemResourceResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemResourceResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o GetSystemResourceResultOutput) Filters() GetSystemResourceFilterArrayOutput {
	return o.ApplyT(func(v GetSystemResourceResult) []GetSystemResourceFilter { return v.Filters }).(GetSystemResourceFilterArrayOutput)
}

func (o GetSystemResourceResultOutput) FlinkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemResourceResult) *string { return v.FlinkVersion }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemResourceResultOutput) ResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemResourceResult) []string { return v.ResourceIds }).(pulumi.StringArrayOutput)
}

// Collection of resource details.
func (o GetSystemResourceResultOutput) ResourceSets() GetSystemResourceResourceSetArrayOutput {
	return o.ApplyT(func(v GetSystemResourceResult) []GetSystemResourceResourceSet { return v.ResourceSets }).(GetSystemResourceResourceSetArrayOutput)
}

func (o GetSystemResourceResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemResourceResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemResourceResultOutput{})
}
