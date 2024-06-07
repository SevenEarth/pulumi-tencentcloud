// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of mysql errorLog
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mysql.GetErrorLog(ctx, &mysql.GetErrorLogArgs{
//				EndTime:    1686043908,
//				InstType:   pulumi.StringRef("slave"),
//				InstanceId: "cdb-fitq5t9h",
//				KeyWords: []string{
//					"Shutting",
//				},
//				StartTime: 1683538307,
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
func GetErrorLog(ctx *pulumi.Context, args *GetErrorLogArgs, opts ...pulumi.InvokeOption) (*GetErrorLogResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetErrorLogResult
	err := ctx.Invoke("tencentcloud:Mysql/getErrorLog:getErrorLog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getErrorLog.
type GetErrorLogArgs struct {
	// End timestamp. For example 1585142640.
	EndTime int `pulumi:"endTime"`
	// Only valid when the instance is the master instance or disaster recovery instance, the optional value: slave, which means to pull the log of the slave machine.
	InstType *string `pulumi:"instType"`
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// A list of keywords to match, up to 15 keywords are supported.
	KeyWords []string `pulumi:"keyWords"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Start timestamp. For example 1585142640.
	StartTime int `pulumi:"startTime"`
}

// A collection of values returned by getErrorLog.
type GetErrorLogResult struct {
	EndTime int `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstType   *string `pulumi:"instType"`
	InstanceId string  `pulumi:"instanceId"`
	// The records returned.
	Items            []GetErrorLogItem `pulumi:"items"`
	KeyWords         []string          `pulumi:"keyWords"`
	ResultOutputFile *string           `pulumi:"resultOutputFile"`
	StartTime        int               `pulumi:"startTime"`
}

func GetErrorLogOutput(ctx *pulumi.Context, args GetErrorLogOutputArgs, opts ...pulumi.InvokeOption) GetErrorLogResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetErrorLogResult, error) {
			args := v.(GetErrorLogArgs)
			r, err := GetErrorLog(ctx, &args, opts...)
			var s GetErrorLogResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetErrorLogResultOutput)
}

// A collection of arguments for invoking getErrorLog.
type GetErrorLogOutputArgs struct {
	// End timestamp. For example 1585142640.
	EndTime pulumi.IntInput `pulumi:"endTime"`
	// Only valid when the instance is the master instance or disaster recovery instance, the optional value: slave, which means to pull the log of the slave machine.
	InstType pulumi.StringPtrInput `pulumi:"instType"`
	// instance id.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// A list of keywords to match, up to 15 keywords are supported.
	KeyWords pulumi.StringArrayInput `pulumi:"keyWords"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Start timestamp. For example 1585142640.
	StartTime pulumi.IntInput `pulumi:"startTime"`
}

func (GetErrorLogOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetErrorLogArgs)(nil)).Elem()
}

// A collection of values returned by getErrorLog.
type GetErrorLogResultOutput struct{ *pulumi.OutputState }

func (GetErrorLogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetErrorLogResult)(nil)).Elem()
}

func (o GetErrorLogResultOutput) ToGetErrorLogResultOutput() GetErrorLogResultOutput {
	return o
}

func (o GetErrorLogResultOutput) ToGetErrorLogResultOutputWithContext(ctx context.Context) GetErrorLogResultOutput {
	return o
}

func (o GetErrorLogResultOutput) EndTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetErrorLogResult) int { return v.EndTime }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetErrorLogResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetErrorLogResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetErrorLogResultOutput) InstType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetErrorLogResult) *string { return v.InstType }).(pulumi.StringPtrOutput)
}

func (o GetErrorLogResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetErrorLogResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The records returned.
func (o GetErrorLogResultOutput) Items() GetErrorLogItemArrayOutput {
	return o.ApplyT(func(v GetErrorLogResult) []GetErrorLogItem { return v.Items }).(GetErrorLogItemArrayOutput)
}

func (o GetErrorLogResultOutput) KeyWords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetErrorLogResult) []string { return v.KeyWords }).(pulumi.StringArrayOutput)
}

func (o GetErrorLogResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetErrorLogResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetErrorLogResultOutput) StartTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetErrorLogResult) int { return v.StartTime }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetErrorLogResultOutput{})
}
