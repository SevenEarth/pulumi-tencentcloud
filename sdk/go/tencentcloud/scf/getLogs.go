// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query SCF function logs.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Scf"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Scf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooFunction, err := Scf.NewFunction(ctx, "fooFunction", &Scf.FunctionArgs{
//				Handler:         pulumi.String("main.do_it"),
//				Runtime:         pulumi.String("Python3.6"),
//				CosBucketName:   pulumi.String("scf-code-1234567890"),
//				CosObjectName:   pulumi.String("code.zip"),
//				CosBucketRegion: pulumi.String("ap-guangzhou"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = Scf.GetLogsOutput(ctx, scf.GetLogsOutputArgs{
//				FunctionName: fooFunction.Name,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetLogs(ctx *pulumi.Context, args *GetLogsArgs, opts ...pulumi.InvokeOption) (*GetLogsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetLogsResult
	err := ctx.Invoke("tencentcloud:Scf/getLogs:getLogs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogs.
type GetLogsArgs struct {
	// The end time of the query, the format is `2017-05-16 20:00:00`, which can only be within one day from `startTime`.
	EndTime *string `pulumi:"endTime"`
	// Name of the SCF function to be queried.
	FunctionName string `pulumi:"functionName"`
	// Corresponding requestId when executing function.
	InvokeRequestId *string `pulumi:"invokeRequestId"`
	// Number of logs, the default is `10000`, offset+limit cannot be greater than 10000.
	Limit *int `pulumi:"limit"`
	// Namespace of the SCF function to be queried.
	Namespace *string `pulumi:"namespace"`
	// Log offset, default is `0`, offset+limit cannot be greater than 10000.
	Offset *int `pulumi:"offset"`
	// Order to sort the log, optional values `desc` and `asc`, default `desc`.
	Order *string `pulumi:"order"`
	// Sort the logs according to the following fields: `functionName`, `duration`, `memUsage`, `startTime`, default `startTime`.
	OrderBy *string `pulumi:"orderBy"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Use to filter log, optional value: `not0` only returns the error log. `is0` only returns the correct log. `TimeLimitExceeded` returns the log of the function call timeout. `ResourceLimitExceeded` returns the function call generation resource overrun log. `UserCodeException` returns logs of the user code error that occurred in the function call. Not passing the parameter means returning all logs.
	RetCode *string `pulumi:"retCode"`
	// The start time of the query, the format is `2017-05-16 20:00:00`, which can only be within one day from `endTime`.
	StartTime *string `pulumi:"startTime"`
}

// A collection of values returned by getLogs.
type GetLogsResult struct {
	EndTime *string `pulumi:"endTime"`
	// Name of the SCF function.
	FunctionName string `pulumi:"functionName"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	InvokeRequestId *string `pulumi:"invokeRequestId"`
	Limit           *int    `pulumi:"limit"`
	// An information list of logs. Each element contains the following attributes:
	Logs             []GetLogsLog `pulumi:"logs"`
	Namespace        *string      `pulumi:"namespace"`
	Offset           *int         `pulumi:"offset"`
	Order            *string      `pulumi:"order"`
	OrderBy          *string      `pulumi:"orderBy"`
	ResultOutputFile *string      `pulumi:"resultOutputFile"`
	// Execution result of function, `0` means the execution is successful, other values indicate failure.
	RetCode *string `pulumi:"retCode"`
	// Point in time at which the function begins execution.
	StartTime *string `pulumi:"startTime"`
}

func GetLogsOutput(ctx *pulumi.Context, args GetLogsOutputArgs, opts ...pulumi.InvokeOption) GetLogsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLogsResult, error) {
			args := v.(GetLogsArgs)
			r, err := GetLogs(ctx, &args, opts...)
			var s GetLogsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLogsResultOutput)
}

// A collection of arguments for invoking getLogs.
type GetLogsOutputArgs struct {
	// The end time of the query, the format is `2017-05-16 20:00:00`, which can only be within one day from `startTime`.
	EndTime pulumi.StringPtrInput `pulumi:"endTime"`
	// Name of the SCF function to be queried.
	FunctionName pulumi.StringInput `pulumi:"functionName"`
	// Corresponding requestId when executing function.
	InvokeRequestId pulumi.StringPtrInput `pulumi:"invokeRequestId"`
	// Number of logs, the default is `10000`, offset+limit cannot be greater than 10000.
	Limit pulumi.IntPtrInput `pulumi:"limit"`
	// Namespace of the SCF function to be queried.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// Log offset, default is `0`, offset+limit cannot be greater than 10000.
	Offset pulumi.IntPtrInput `pulumi:"offset"`
	// Order to sort the log, optional values `desc` and `asc`, default `desc`.
	Order pulumi.StringPtrInput `pulumi:"order"`
	// Sort the logs according to the following fields: `functionName`, `duration`, `memUsage`, `startTime`, default `startTime`.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Use to filter log, optional value: `not0` only returns the error log. `is0` only returns the correct log. `TimeLimitExceeded` returns the log of the function call timeout. `ResourceLimitExceeded` returns the function call generation resource overrun log. `UserCodeException` returns logs of the user code error that occurred in the function call. Not passing the parameter means returning all logs.
	RetCode pulumi.StringPtrInput `pulumi:"retCode"`
	// The start time of the query, the format is `2017-05-16 20:00:00`, which can only be within one day from `endTime`.
	StartTime pulumi.StringPtrInput `pulumi:"startTime"`
}

func (GetLogsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogsArgs)(nil)).Elem()
}

// A collection of values returned by getLogs.
type GetLogsResultOutput struct{ *pulumi.OutputState }

func (GetLogsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogsResult)(nil)).Elem()
}

func (o GetLogsResultOutput) ToGetLogsResultOutput() GetLogsResultOutput {
	return o
}

func (o GetLogsResultOutput) ToGetLogsResultOutputWithContext(ctx context.Context) GetLogsResultOutput {
	return o
}

func (o GetLogsResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogsResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// Name of the SCF function.
func (o GetLogsResultOutput) FunctionName() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsResult) string { return v.FunctionName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLogsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLogsResultOutput) InvokeRequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogsResult) *string { return v.InvokeRequestId }).(pulumi.StringPtrOutput)
}

func (o GetLogsResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetLogsResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

// An information list of logs. Each element contains the following attributes:
func (o GetLogsResultOutput) Logs() GetLogsLogArrayOutput {
	return o.ApplyT(func(v GetLogsResult) []GetLogsLog { return v.Logs }).(GetLogsLogArrayOutput)
}

func (o GetLogsResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogsResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetLogsResultOutput) Offset() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetLogsResult) *int { return v.Offset }).(pulumi.IntPtrOutput)
}

func (o GetLogsResultOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogsResult) *string { return v.Order }).(pulumi.StringPtrOutput)
}

func (o GetLogsResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogsResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o GetLogsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Execution result of function, `0` means the execution is successful, other values indicate failure.
func (o GetLogsResultOutput) RetCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogsResult) *string { return v.RetCode }).(pulumi.StringPtrOutput)
}

// Point in time at which the function begins execution.
func (o GetLogsResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogsResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLogsResultOutput{})
}
