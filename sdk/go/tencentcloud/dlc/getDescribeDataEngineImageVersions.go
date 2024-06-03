// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dlc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of dlc describeDataEngineImageVersions
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dlc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dlc.GetDescribeDataEngineImageVersions(ctx, &dlc.GetDescribeDataEngineImageVersionsArgs{
//				EngineType: "SparkBatch",
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
func GetDescribeDataEngineImageVersions(ctx *pulumi.Context, args *GetDescribeDataEngineImageVersionsArgs, opts ...pulumi.InvokeOption) (*GetDescribeDataEngineImageVersionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDescribeDataEngineImageVersionsResult
	err := ctx.Invoke("tencentcloud:Dlc/getDescribeDataEngineImageVersions:getDescribeDataEngineImageVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDescribeDataEngineImageVersions.
type GetDescribeDataEngineImageVersionsArgs struct {
	// Engine type only support: SparkSQL/PrestoSQL/SparkBatch.
	EngineType string `pulumi:"engineType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDescribeDataEngineImageVersions.
type GetDescribeDataEngineImageVersionsResult struct {
	// Engine type only support: SparkSQL/PrestoSQL/SparkBatch.
	EngineType string `pulumi:"engineType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Cluster large version image information list.
	ImageParentVersions []GetDescribeDataEngineImageVersionsImageParentVersion `pulumi:"imageParentVersions"`
	ResultOutputFile    *string                                                `pulumi:"resultOutputFile"`
}

func GetDescribeDataEngineImageVersionsOutput(ctx *pulumi.Context, args GetDescribeDataEngineImageVersionsOutputArgs, opts ...pulumi.InvokeOption) GetDescribeDataEngineImageVersionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDescribeDataEngineImageVersionsResult, error) {
			args := v.(GetDescribeDataEngineImageVersionsArgs)
			r, err := GetDescribeDataEngineImageVersions(ctx, &args, opts...)
			var s GetDescribeDataEngineImageVersionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDescribeDataEngineImageVersionsResultOutput)
}

// A collection of arguments for invoking getDescribeDataEngineImageVersions.
type GetDescribeDataEngineImageVersionsOutputArgs struct {
	// Engine type only support: SparkSQL/PrestoSQL/SparkBatch.
	EngineType pulumi.StringInput `pulumi:"engineType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDescribeDataEngineImageVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeDataEngineImageVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getDescribeDataEngineImageVersions.
type GetDescribeDataEngineImageVersionsResultOutput struct{ *pulumi.OutputState }

func (GetDescribeDataEngineImageVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeDataEngineImageVersionsResult)(nil)).Elem()
}

func (o GetDescribeDataEngineImageVersionsResultOutput) ToGetDescribeDataEngineImageVersionsResultOutput() GetDescribeDataEngineImageVersionsResultOutput {
	return o
}

func (o GetDescribeDataEngineImageVersionsResultOutput) ToGetDescribeDataEngineImageVersionsResultOutputWithContext(ctx context.Context) GetDescribeDataEngineImageVersionsResultOutput {
	return o
}

// Engine type only support: SparkSQL/PrestoSQL/SparkBatch.
func (o GetDescribeDataEngineImageVersionsResultOutput) EngineType() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeDataEngineImageVersionsResult) string { return v.EngineType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDescribeDataEngineImageVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeDataEngineImageVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Cluster large version image information list.
func (o GetDescribeDataEngineImageVersionsResultOutput) ImageParentVersions() GetDescribeDataEngineImageVersionsImageParentVersionArrayOutput {
	return o.ApplyT(func(v GetDescribeDataEngineImageVersionsResult) []GetDescribeDataEngineImageVersionsImageParentVersion {
		return v.ImageParentVersions
	}).(GetDescribeDataEngineImageVersionsImageParentVersionArrayOutput)
}

func (o GetDescribeDataEngineImageVersionsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeDataEngineImageVersionsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDescribeDataEngineImageVersionsResultOutput{})
}
