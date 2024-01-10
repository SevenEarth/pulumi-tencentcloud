// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package antiddos

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of antiddos overview ddos trend
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Antiddos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Antiddos"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Antiddos.GetOverviewDdosTrend(ctx, &antiddos.GetOverviewDdosTrendArgs{
//				Business:   pulumi.StringRef("bgpip"),
//				EndTime:    "2023-11-21 14:16:23",
//				MetricName: "bps",
//				Period:     300,
//				StartTime:  "2023-11-20 14:16:23",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOverviewDdosTrend(ctx *pulumi.Context, args *GetOverviewDdosTrendArgs, opts ...pulumi.InvokeOption) (*GetOverviewDdosTrendResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetOverviewDdosTrendResult
	err := ctx.Invoke("tencentcloud:Antiddos/getOverviewDdosTrend:getOverviewDdosTrend", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOverviewDdosTrend.
type GetOverviewDdosTrendArgs struct {
	// Dayu sub product code (bgpip represents advanced defense IP; net represents professional version of advanced defense IP).
	Business *string `pulumi:"business"`
	// EndTime.
	EndTime string `pulumi:"endTime"`
	// instance IpList.
	IpLists []string `pulumi:"ipLists"`
	// Indicator, value [bps (attack traffic bandwidth, pps (attack packet rate)].
	MetricName string `pulumi:"metricName"`
	// Statistical granularity, values [300 (5 minutes), 3600 (hours), 86400 (days)].
	Period int `pulumi:"period"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// StartTime.
	StartTime string `pulumi:"startTime"`
}

// A collection of values returned by getOverviewDdosTrend.
type GetOverviewDdosTrendResult struct {
	Business *string `pulumi:"business"`
	// Array, attack traffic bandwidth in Mbps, packet rate in pps.
	Datas   []int  `pulumi:"datas"`
	EndTime string `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id               string   `pulumi:"id"`
	IpLists          []string `pulumi:"ipLists"`
	MetricName       string   `pulumi:"metricName"`
	Period           int      `pulumi:"period"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
	StartTime        string   `pulumi:"startTime"`
}

func GetOverviewDdosTrendOutput(ctx *pulumi.Context, args GetOverviewDdosTrendOutputArgs, opts ...pulumi.InvokeOption) GetOverviewDdosTrendResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOverviewDdosTrendResult, error) {
			args := v.(GetOverviewDdosTrendArgs)
			r, err := GetOverviewDdosTrend(ctx, &args, opts...)
			var s GetOverviewDdosTrendResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOverviewDdosTrendResultOutput)
}

// A collection of arguments for invoking getOverviewDdosTrend.
type GetOverviewDdosTrendOutputArgs struct {
	// Dayu sub product code (bgpip represents advanced defense IP; net represents professional version of advanced defense IP).
	Business pulumi.StringPtrInput `pulumi:"business"`
	// EndTime.
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// instance IpList.
	IpLists pulumi.StringArrayInput `pulumi:"ipLists"`
	// Indicator, value [bps (attack traffic bandwidth, pps (attack packet rate)].
	MetricName pulumi.StringInput `pulumi:"metricName"`
	// Statistical granularity, values [300 (5 minutes), 3600 (hours), 86400 (days)].
	Period pulumi.IntInput `pulumi:"period"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// StartTime.
	StartTime pulumi.StringInput `pulumi:"startTime"`
}

func (GetOverviewDdosTrendOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOverviewDdosTrendArgs)(nil)).Elem()
}

// A collection of values returned by getOverviewDdosTrend.
type GetOverviewDdosTrendResultOutput struct{ *pulumi.OutputState }

func (GetOverviewDdosTrendResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOverviewDdosTrendResult)(nil)).Elem()
}

func (o GetOverviewDdosTrendResultOutput) ToGetOverviewDdosTrendResultOutput() GetOverviewDdosTrendResultOutput {
	return o
}

func (o GetOverviewDdosTrendResultOutput) ToGetOverviewDdosTrendResultOutputWithContext(ctx context.Context) GetOverviewDdosTrendResultOutput {
	return o
}

func (o GetOverviewDdosTrendResultOutput) Business() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOverviewDdosTrendResult) *string { return v.Business }).(pulumi.StringPtrOutput)
}

// Array, attack traffic bandwidth in Mbps, packet rate in pps.
func (o GetOverviewDdosTrendResultOutput) Datas() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetOverviewDdosTrendResult) []int { return v.Datas }).(pulumi.IntArrayOutput)
}

func (o GetOverviewDdosTrendResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetOverviewDdosTrendResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOverviewDdosTrendResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOverviewDdosTrendResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOverviewDdosTrendResultOutput) IpLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOverviewDdosTrendResult) []string { return v.IpLists }).(pulumi.StringArrayOutput)
}

func (o GetOverviewDdosTrendResultOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v GetOverviewDdosTrendResult) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o GetOverviewDdosTrendResultOutput) Period() pulumi.IntOutput {
	return o.ApplyT(func(v GetOverviewDdosTrendResult) int { return v.Period }).(pulumi.IntOutput)
}

func (o GetOverviewDdosTrendResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOverviewDdosTrendResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetOverviewDdosTrendResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetOverviewDdosTrendResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOverviewDdosTrendResultOutput{})
}
