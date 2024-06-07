// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rum

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

func GetLogList(ctx *pulumi.Context, args *GetLogListArgs, opts ...pulumi.InvokeOption) (*GetLogListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLogListResult
	err := ctx.Invoke("tencentcloud:Rum/getLogList:getLogList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogList.
type GetLogListArgs struct {
	EndTime          string  `pulumi:"endTime"`
	OrderBy          string  `pulumi:"orderBy"`
	ProjectId        int     `pulumi:"projectId"`
	Query            string  `pulumi:"query"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	StartTime        string  `pulumi:"startTime"`
}

// A collection of values returned by getLogList.
type GetLogListResult struct {
	EndTime string `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	OrderBy          string  `pulumi:"orderBy"`
	ProjectId        int     `pulumi:"projectId"`
	Query            string  `pulumi:"query"`
	Result           string  `pulumi:"result"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	StartTime        string  `pulumi:"startTime"`
}

func GetLogListOutput(ctx *pulumi.Context, args GetLogListOutputArgs, opts ...pulumi.InvokeOption) GetLogListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLogListResult, error) {
			args := v.(GetLogListArgs)
			r, err := GetLogList(ctx, &args, opts...)
			var s GetLogListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLogListResultOutput)
}

// A collection of arguments for invoking getLogList.
type GetLogListOutputArgs struct {
	EndTime          pulumi.StringInput    `pulumi:"endTime"`
	OrderBy          pulumi.StringInput    `pulumi:"orderBy"`
	ProjectId        pulumi.IntInput       `pulumi:"projectId"`
	Query            pulumi.StringInput    `pulumi:"query"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	StartTime        pulumi.StringInput    `pulumi:"startTime"`
}

func (GetLogListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogListArgs)(nil)).Elem()
}

// A collection of values returned by getLogList.
type GetLogListResultOutput struct{ *pulumi.OutputState }

func (GetLogListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogListResult)(nil)).Elem()
}

func (o GetLogListResultOutput) ToGetLogListResultOutput() GetLogListResultOutput {
	return o
}

func (o GetLogListResultOutput) ToGetLogListResultOutputWithContext(ctx context.Context) GetLogListResultOutput {
	return o
}

func (o GetLogListResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogListResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLogListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLogListResultOutput) OrderBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogListResult) string { return v.OrderBy }).(pulumi.StringOutput)
}

func (o GetLogListResultOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v GetLogListResult) int { return v.ProjectId }).(pulumi.IntOutput)
}

func (o GetLogListResultOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogListResult) string { return v.Query }).(pulumi.StringOutput)
}

func (o GetLogListResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogListResult) string { return v.Result }).(pulumi.StringOutput)
}

func (o GetLogListResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogListResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetLogListResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogListResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLogListResultOutput{})
}
