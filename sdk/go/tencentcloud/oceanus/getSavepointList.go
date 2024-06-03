// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oceanus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of oceanus savepointList
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
//			_, err := Oceanus.GetSavepointList(ctx, &oceanus.GetSavepointListArgs{
//				JobId:       "cql-314rw6w0",
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
func GetSavepointList(ctx *pulumi.Context, args *GetSavepointListArgs, opts ...pulumi.InvokeOption) (*GetSavepointListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSavepointListResult
	err := ctx.Invoke("tencentcloud:Oceanus/getSavepointList:getSavepointList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSavepointList.
type GetSavepointListArgs struct {
	// Job SerialId.
	JobId string `pulumi:"jobId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Workspace SerialId.
	WorkSpaceId *string `pulumi:"workSpaceId"`
}

// A collection of values returned by getSavepointList.
type GetSavepointListResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	JobId            string  `pulumi:"jobId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Snapshot listNote: This field may return null, indicating that no valid value was found.
	Savepoints  []GetSavepointListSavepoint `pulumi:"savepoints"`
	WorkSpaceId *string                     `pulumi:"workSpaceId"`
}

func GetSavepointListOutput(ctx *pulumi.Context, args GetSavepointListOutputArgs, opts ...pulumi.InvokeOption) GetSavepointListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSavepointListResult, error) {
			args := v.(GetSavepointListArgs)
			r, err := GetSavepointList(ctx, &args, opts...)
			var s GetSavepointListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSavepointListResultOutput)
}

// A collection of arguments for invoking getSavepointList.
type GetSavepointListOutputArgs struct {
	// Job SerialId.
	JobId pulumi.StringInput `pulumi:"jobId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Workspace SerialId.
	WorkSpaceId pulumi.StringPtrInput `pulumi:"workSpaceId"`
}

func (GetSavepointListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSavepointListArgs)(nil)).Elem()
}

// A collection of values returned by getSavepointList.
type GetSavepointListResultOutput struct{ *pulumi.OutputState }

func (GetSavepointListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSavepointListResult)(nil)).Elem()
}

func (o GetSavepointListResultOutput) ToGetSavepointListResultOutput() GetSavepointListResultOutput {
	return o
}

func (o GetSavepointListResultOutput) ToGetSavepointListResultOutputWithContext(ctx context.Context) GetSavepointListResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSavepointListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSavepointListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSavepointListResultOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSavepointListResult) string { return v.JobId }).(pulumi.StringOutput)
}

func (o GetSavepointListResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSavepointListResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Snapshot listNote: This field may return null, indicating that no valid value was found.
func (o GetSavepointListResultOutput) Savepoints() GetSavepointListSavepointArrayOutput {
	return o.ApplyT(func(v GetSavepointListResult) []GetSavepointListSavepoint { return v.Savepoints }).(GetSavepointListSavepointArrayOutput)
}

func (o GetSavepointListResultOutput) WorkSpaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSavepointListResult) *string { return v.WorkSpaceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSavepointListResultOutput{})
}
