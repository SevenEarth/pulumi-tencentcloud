// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of bi project
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Bi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Bi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Bi.GetProject(ctx, &bi.GetProjectArgs{
//				AllPage:          pulumi.BoolRef(true),
//				Keyword:          pulumi.StringRef("abc"),
//				ModuleCollection: pulumi.StringRef("sys_common_user"),
//				PageNo:           pulumi.IntRef(1),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("tencentcloud:Bi/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type LookupProjectArgs struct {
	// Whether to display all, if true, ignore paging.
	AllPage *bool `pulumi:"allPage"`
	// Retrieve fuzzy fields.
	Keyword *string `pulumi:"keyword"`
	// Role information, can be ignored.
	ModuleCollection *string `pulumi:"moduleCollection"`
	// Page number.
	PageNo *int `pulumi:"pageNo"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	AllPage *bool `pulumi:"allPage"`
	// Additional information(Note: This field may return null, indicating that no valid value can be obtained).
	Extra string `pulumi:"extra"`
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	Keyword *string `pulumi:"keyword"`
	// Array(Note: This field may return null, indicating that no valid value can be obtained).
	Lists            []GetProjectList `pulumi:"lists"`
	ModuleCollection *string          `pulumi:"moduleCollection"`
	// Interface information(Note: This field may return null, indicating that no valid value can be obtained).
	Msg              string  `pulumi:"msg"`
	PageNo           *int    `pulumi:"pageNo"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type LookupProjectOutputArgs struct {
	// Whether to display all, if true, ignore paging.
	AllPage pulumi.BoolPtrInput `pulumi:"allPage"`
	// Retrieve fuzzy fields.
	Keyword pulumi.StringPtrInput `pulumi:"keyword"`
	// Role information, can be ignored.
	ModuleCollection pulumi.StringPtrInput `pulumi:"moduleCollection"`
	// Page number.
	PageNo pulumi.IntPtrInput `pulumi:"pageNo"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// A collection of values returned by getProject.
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) AllPage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *bool { return v.AllPage }).(pulumi.BoolPtrOutput)
}

// Additional information(Note: This field may return null, indicating that no valid value can be obtained).
func (o LookupProjectResultOutput) Extra() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Extra }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Keyword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Keyword }).(pulumi.StringPtrOutput)
}

// Array(Note: This field may return null, indicating that no valid value can be obtained).
func (o LookupProjectResultOutput) Lists() GetProjectListArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []GetProjectList { return v.Lists }).(GetProjectListArrayOutput)
}

func (o LookupProjectResultOutput) ModuleCollection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ModuleCollection }).(pulumi.StringPtrOutput)
}

// Interface information(Note: This field may return null, indicating that no valid value can be obtained).
func (o LookupProjectResultOutput) Msg() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Msg }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) PageNo() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *int { return v.PageNo }).(pulumi.IntPtrOutput)
}

func (o LookupProjectResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
