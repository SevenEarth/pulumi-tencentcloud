// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

func GetDescHaLog(ctx *pulumi.Context, args *GetDescHaLogArgs, opts ...pulumi.InvokeOption) (*GetDescHaLogResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDescHaLogResult
	err := ctx.Invoke("tencentcloud:Sqlserver/getDescHaLog:getDescHaLog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDescHaLog.
type GetDescHaLogArgs struct {
	EndTime          string  `pulumi:"endTime"`
	InstanceId       string  `pulumi:"instanceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	StartTime        string  `pulumi:"startTime"`
	SwitchType       *int    `pulumi:"switchType"`
}

// A collection of values returned by getDescHaLog.
type GetDescHaLogResult struct {
	EndTime string `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                  `pulumi:"id"`
	InstanceId       string                  `pulumi:"instanceId"`
	ResultOutputFile *string                 `pulumi:"resultOutputFile"`
	StartTime        string                  `pulumi:"startTime"`
	SwitchLogs       []GetDescHaLogSwitchLog `pulumi:"switchLogs"`
	SwitchType       *int                    `pulumi:"switchType"`
}

func GetDescHaLogOutput(ctx *pulumi.Context, args GetDescHaLogOutputArgs, opts ...pulumi.InvokeOption) GetDescHaLogResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDescHaLogResult, error) {
			args := v.(GetDescHaLogArgs)
			r, err := GetDescHaLog(ctx, &args, opts...)
			var s GetDescHaLogResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDescHaLogResultOutput)
}

// A collection of arguments for invoking getDescHaLog.
type GetDescHaLogOutputArgs struct {
	EndTime          pulumi.StringInput    `pulumi:"endTime"`
	InstanceId       pulumi.StringInput    `pulumi:"instanceId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	StartTime        pulumi.StringInput    `pulumi:"startTime"`
	SwitchType       pulumi.IntPtrInput    `pulumi:"switchType"`
}

func (GetDescHaLogOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescHaLogArgs)(nil)).Elem()
}

// A collection of values returned by getDescHaLog.
type GetDescHaLogResultOutput struct{ *pulumi.OutputState }

func (GetDescHaLogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescHaLogResult)(nil)).Elem()
}

func (o GetDescHaLogResultOutput) ToGetDescHaLogResultOutput() GetDescHaLogResultOutput {
	return o
}

func (o GetDescHaLogResultOutput) ToGetDescHaLogResultOutputWithContext(ctx context.Context) GetDescHaLogResultOutput {
	return o
}

func (o GetDescHaLogResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescHaLogResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDescHaLogResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescHaLogResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDescHaLogResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescHaLogResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetDescHaLogResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescHaLogResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetDescHaLogResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescHaLogResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o GetDescHaLogResultOutput) SwitchLogs() GetDescHaLogSwitchLogArrayOutput {
	return o.ApplyT(func(v GetDescHaLogResult) []GetDescHaLogSwitchLog { return v.SwitchLogs }).(GetDescHaLogSwitchLogArrayOutput)
}

func (o GetDescHaLogResultOutput) SwitchType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDescHaLogResult) *int { return v.SwitchType }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDescHaLogResultOutput{})
}
