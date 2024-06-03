// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oceanus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of oceanus treeJobs
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Oceanus"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Oceanus.GetTreeJobs(ctx, &oceanus.GetTreeJobsArgs{
//				WorkSpaceId: pulumi.StringRef("space-2idq8wbr"),
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
func GetTreeJobs(ctx *pulumi.Context, args *GetTreeJobsArgs, opts ...pulumi.InvokeOption) (*GetTreeJobsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTreeJobsResult
	err := ctx.Invoke("tencentcloud:Oceanus/getTreeJobs:getTreeJobs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTreeJobs.
type GetTreeJobsArgs struct {
	// Filter rules.
	Filters []GetTreeJobsFilter `pulumi:"filters"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Workspace SerialId.
	WorkSpaceId *string `pulumi:"workSpaceId"`
}

// A collection of values returned by getTreeJobs.
type GetTreeJobsResult struct {
	Filters []GetTreeJobsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Tree structure information.
	TreeInfos   []GetTreeJobsTreeInfo `pulumi:"treeInfos"`
	WorkSpaceId *string               `pulumi:"workSpaceId"`
}

func GetTreeJobsOutput(ctx *pulumi.Context, args GetTreeJobsOutputArgs, opts ...pulumi.InvokeOption) GetTreeJobsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTreeJobsResult, error) {
			args := v.(GetTreeJobsArgs)
			r, err := GetTreeJobs(ctx, &args, opts...)
			var s GetTreeJobsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTreeJobsResultOutput)
}

// A collection of arguments for invoking getTreeJobs.
type GetTreeJobsOutputArgs struct {
	// Filter rules.
	Filters GetTreeJobsFilterArrayInput `pulumi:"filters"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Workspace SerialId.
	WorkSpaceId pulumi.StringPtrInput `pulumi:"workSpaceId"`
}

func (GetTreeJobsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTreeJobsArgs)(nil)).Elem()
}

// A collection of values returned by getTreeJobs.
type GetTreeJobsResultOutput struct{ *pulumi.OutputState }

func (GetTreeJobsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTreeJobsResult)(nil)).Elem()
}

func (o GetTreeJobsResultOutput) ToGetTreeJobsResultOutput() GetTreeJobsResultOutput {
	return o
}

func (o GetTreeJobsResultOutput) ToGetTreeJobsResultOutputWithContext(ctx context.Context) GetTreeJobsResultOutput {
	return o
}

func (o GetTreeJobsResultOutput) Filters() GetTreeJobsFilterArrayOutput {
	return o.ApplyT(func(v GetTreeJobsResult) []GetTreeJobsFilter { return v.Filters }).(GetTreeJobsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTreeJobsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTreeJobsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTreeJobsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTreeJobsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Tree structure information.
func (o GetTreeJobsResultOutput) TreeInfos() GetTreeJobsTreeInfoArrayOutput {
	return o.ApplyT(func(v GetTreeJobsResult) []GetTreeJobsTreeInfo { return v.TreeInfos }).(GetTreeJobsTreeInfoArrayOutput)
}

func (o GetTreeJobsResultOutput) WorkSpaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTreeJobsResult) *string { return v.WorkSpaceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTreeJobsResultOutput{})
}
