// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of cynosdb describeInstanceErrorLogs
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cynosdb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cynosdb.GetDescribeInstanceErrorLogs(ctx, &cynosdb.GetDescribeInstanceErrorLogsArgs{
//				EndTime:    pulumi.StringRef("2023-06-19 15:04:05"),
//				InstanceId: "cynosdbmysql-ins-afqx1hy0",
//				KeyWords: []string{
//					"Aborted",
//				},
//				LogLevels: []string{
//					"note",
//					"warning",
//				},
//				OrderBy:     pulumi.StringRef("Timestamp"),
//				OrderByType: pulumi.StringRef("DESC"),
//				StartTime:   pulumi.StringRef("2023-06-01 15:04:05"),
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
func GetDescribeInstanceErrorLogs(ctx *pulumi.Context, args *GetDescribeInstanceErrorLogsArgs, opts ...pulumi.InvokeOption) (*GetDescribeInstanceErrorLogsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDescribeInstanceErrorLogsResult
	err := ctx.Invoke("tencentcloud:Cynosdb/getDescribeInstanceErrorLogs:getDescribeInstanceErrorLogs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDescribeInstanceErrorLogs.
type GetDescribeInstanceErrorLogsArgs struct {
	// End time.
	EndTime *string `pulumi:"endTime"`
	// Instance Id.
	InstanceId string `pulumi:"instanceId"`
	// Keywords, supports fuzzy search.
	KeyWords []string `pulumi:"keyWords"`
	// Log levels, including error, warning, and note, support simultaneous search of multiple levels.
	LogLevels []string `pulumi:"logLevels"`
	// Sort fields with Timestamp enumeration values.
	OrderBy *string `pulumi:"orderBy"`
	// Sort type, with ASC and DESC enumeration values.
	OrderByType *string `pulumi:"orderByType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// start time.
	StartTime *string `pulumi:"startTime"`
}

// A collection of values returned by getDescribeInstanceErrorLogs.
type GetDescribeInstanceErrorLogsResult struct {
	EndTime *string `pulumi:"endTime"`
	// Error log list note: This field may return null, indicating that a valid value cannot be obtained.
	ErrorLogs []GetDescribeInstanceErrorLogsErrorLog `pulumi:"errorLogs"`
	// The provider-assigned unique ID for this managed resource.
	Id               string   `pulumi:"id"`
	InstanceId       string   `pulumi:"instanceId"`
	KeyWords         []string `pulumi:"keyWords"`
	LogLevels        []string `pulumi:"logLevels"`
	OrderBy          *string  `pulumi:"orderBy"`
	OrderByType      *string  `pulumi:"orderByType"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
	StartTime        *string  `pulumi:"startTime"`
}

func GetDescribeInstanceErrorLogsOutput(ctx *pulumi.Context, args GetDescribeInstanceErrorLogsOutputArgs, opts ...pulumi.InvokeOption) GetDescribeInstanceErrorLogsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDescribeInstanceErrorLogsResult, error) {
			args := v.(GetDescribeInstanceErrorLogsArgs)
			r, err := GetDescribeInstanceErrorLogs(ctx, &args, opts...)
			var s GetDescribeInstanceErrorLogsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDescribeInstanceErrorLogsResultOutput)
}

// A collection of arguments for invoking getDescribeInstanceErrorLogs.
type GetDescribeInstanceErrorLogsOutputArgs struct {
	// End time.
	EndTime pulumi.StringPtrInput `pulumi:"endTime"`
	// Instance Id.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Keywords, supports fuzzy search.
	KeyWords pulumi.StringArrayInput `pulumi:"keyWords"`
	// Log levels, including error, warning, and note, support simultaneous search of multiple levels.
	LogLevels pulumi.StringArrayInput `pulumi:"logLevels"`
	// Sort fields with Timestamp enumeration values.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// Sort type, with ASC and DESC enumeration values.
	OrderByType pulumi.StringPtrInput `pulumi:"orderByType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// start time.
	StartTime pulumi.StringPtrInput `pulumi:"startTime"`
}

func (GetDescribeInstanceErrorLogsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeInstanceErrorLogsArgs)(nil)).Elem()
}

// A collection of values returned by getDescribeInstanceErrorLogs.
type GetDescribeInstanceErrorLogsResultOutput struct{ *pulumi.OutputState }

func (GetDescribeInstanceErrorLogsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeInstanceErrorLogsResult)(nil)).Elem()
}

func (o GetDescribeInstanceErrorLogsResultOutput) ToGetDescribeInstanceErrorLogsResultOutput() GetDescribeInstanceErrorLogsResultOutput {
	return o
}

func (o GetDescribeInstanceErrorLogsResultOutput) ToGetDescribeInstanceErrorLogsResultOutputWithContext(ctx context.Context) GetDescribeInstanceErrorLogsResultOutput {
	return o
}

func (o GetDescribeInstanceErrorLogsResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeInstanceErrorLogsResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// Error log list note: This field may return null, indicating that a valid value cannot be obtained.
func (o GetDescribeInstanceErrorLogsResultOutput) ErrorLogs() GetDescribeInstanceErrorLogsErrorLogArrayOutput {
	return o.ApplyT(func(v GetDescribeInstanceErrorLogsResult) []GetDescribeInstanceErrorLogsErrorLog { return v.ErrorLogs }).(GetDescribeInstanceErrorLogsErrorLogArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDescribeInstanceErrorLogsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeInstanceErrorLogsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDescribeInstanceErrorLogsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeInstanceErrorLogsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetDescribeInstanceErrorLogsResultOutput) KeyWords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDescribeInstanceErrorLogsResult) []string { return v.KeyWords }).(pulumi.StringArrayOutput)
}

func (o GetDescribeInstanceErrorLogsResultOutput) LogLevels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDescribeInstanceErrorLogsResult) []string { return v.LogLevels }).(pulumi.StringArrayOutput)
}

func (o GetDescribeInstanceErrorLogsResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeInstanceErrorLogsResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o GetDescribeInstanceErrorLogsResultOutput) OrderByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeInstanceErrorLogsResult) *string { return v.OrderByType }).(pulumi.StringPtrOutput)
}

func (o GetDescribeInstanceErrorLogsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeInstanceErrorLogsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetDescribeInstanceErrorLogsResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeInstanceErrorLogsResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDescribeInstanceErrorLogsResultOutput{})
}
