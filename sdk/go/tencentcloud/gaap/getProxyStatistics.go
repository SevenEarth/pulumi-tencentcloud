// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of gaap proxy statistics
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Gaap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Gaap"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Gaap.GetProxyStatistics(ctx, &gaap.GetProxyStatisticsArgs{
//				EndTime:     "2023-10-09 23:59:59",
//				Granularity: 300,
//				MetricNames: []string{
//					"InBandwidth",
//					"OutBandwidth",
//					"InFlow",
//					"OutFlow",
//					"InPackets",
//					"OutPackets",
//					"Concurrent",
//					"HttpQPS",
//					"HttpsQPS",
//					"Latency",
//					"PacketLoss",
//				},
//				ProxyId:   "link-8lpyo88p",
//				StartTime: "2023-10-09 00:00:00",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetProxyStatistics(ctx *pulumi.Context, args *GetProxyStatisticsArgs, opts ...pulumi.InvokeOption) (*GetProxyStatisticsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetProxyStatisticsResult
	err := ctx.Invoke("tencentcloud:Gaap/getProxyStatistics:getProxyStatistics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProxyStatistics.
type GetProxyStatisticsArgs struct {
	// End Time(2019-03-25 12:00:00).
	EndTime string `pulumi:"endTime"`
	// Monitoring granularity, currently supporting 60 300 3600 86400, in seconds.When the time range does not exceed 3 days, support a minimum granularity of 60 seconds;When the time range does not exceed 7 days, support a minimum granularity of 300 seconds;When the time range does not exceed 30 days, the minimum granularity supported is 3600 seconds.
	Granularity int `pulumi:"granularity"`
	// Operator (valid when the proxy is a three network proxy), supports CMCC, CUCC, CTCC, and merges data from the three operators if null values are passed or not passed.
	Isp *string `pulumi:"isp"`
	// Metric Names. Valid values: InBandwidth,OutBandwidth, Concurrent, InPackets, OutPackets, PacketLoss, Latency, HttpQPS, HttpsQPS.
	MetricNames []string `pulumi:"metricNames"`
	// Proxy Id.
	ProxyId string `pulumi:"proxyId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Start Time(2019-03-25 12:00:00).
	StartTime string `pulumi:"startTime"`
}

// A collection of values returned by getProxyStatistics.
type GetProxyStatisticsResult struct {
	EndTime     string `pulumi:"endTime"`
	Granularity int    `pulumi:"granularity"`
	// The provider-assigned unique ID for this managed resource.
	Id               string   `pulumi:"id"`
	Isp              *string  `pulumi:"isp"`
	MetricNames      []string `pulumi:"metricNames"`
	ProxyId          string   `pulumi:"proxyId"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
	StartTime        string   `pulumi:"startTime"`
	// proxy Statistics.
	StatisticsDatas []GetProxyStatisticsStatisticsData `pulumi:"statisticsDatas"`
}

func GetProxyStatisticsOutput(ctx *pulumi.Context, args GetProxyStatisticsOutputArgs, opts ...pulumi.InvokeOption) GetProxyStatisticsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProxyStatisticsResult, error) {
			args := v.(GetProxyStatisticsArgs)
			r, err := GetProxyStatistics(ctx, &args, opts...)
			var s GetProxyStatisticsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProxyStatisticsResultOutput)
}

// A collection of arguments for invoking getProxyStatistics.
type GetProxyStatisticsOutputArgs struct {
	// End Time(2019-03-25 12:00:00).
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// Monitoring granularity, currently supporting 60 300 3600 86400, in seconds.When the time range does not exceed 3 days, support a minimum granularity of 60 seconds;When the time range does not exceed 7 days, support a minimum granularity of 300 seconds;When the time range does not exceed 30 days, the minimum granularity supported is 3600 seconds.
	Granularity pulumi.IntInput `pulumi:"granularity"`
	// Operator (valid when the proxy is a three network proxy), supports CMCC, CUCC, CTCC, and merges data from the three operators if null values are passed or not passed.
	Isp pulumi.StringPtrInput `pulumi:"isp"`
	// Metric Names. Valid values: InBandwidth,OutBandwidth, Concurrent, InPackets, OutPackets, PacketLoss, Latency, HttpQPS, HttpsQPS.
	MetricNames pulumi.StringArrayInput `pulumi:"metricNames"`
	// Proxy Id.
	ProxyId pulumi.StringInput `pulumi:"proxyId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Start Time(2019-03-25 12:00:00).
	StartTime pulumi.StringInput `pulumi:"startTime"`
}

func (GetProxyStatisticsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxyStatisticsArgs)(nil)).Elem()
}

// A collection of values returned by getProxyStatistics.
type GetProxyStatisticsResultOutput struct{ *pulumi.OutputState }

func (GetProxyStatisticsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxyStatisticsResult)(nil)).Elem()
}

func (o GetProxyStatisticsResultOutput) ToGetProxyStatisticsResultOutput() GetProxyStatisticsResultOutput {
	return o
}

func (o GetProxyStatisticsResultOutput) ToGetProxyStatisticsResultOutputWithContext(ctx context.Context) GetProxyStatisticsResultOutput {
	return o
}

func (o GetProxyStatisticsResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxyStatisticsResult) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o GetProxyStatisticsResultOutput) Granularity() pulumi.IntOutput {
	return o.ApplyT(func(v GetProxyStatisticsResult) int { return v.Granularity }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProxyStatisticsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxyStatisticsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetProxyStatisticsResultOutput) Isp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProxyStatisticsResult) *string { return v.Isp }).(pulumi.StringPtrOutput)
}

func (o GetProxyStatisticsResultOutput) MetricNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProxyStatisticsResult) []string { return v.MetricNames }).(pulumi.StringArrayOutput)
}

func (o GetProxyStatisticsResultOutput) ProxyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxyStatisticsResult) string { return v.ProxyId }).(pulumi.StringOutput)
}

func (o GetProxyStatisticsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProxyStatisticsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetProxyStatisticsResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxyStatisticsResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// proxy Statistics.
func (o GetProxyStatisticsResultOutput) StatisticsDatas() GetProxyStatisticsStatisticsDataArrayOutput {
	return o.ApplyT(func(v GetProxyStatisticsResult) []GetProxyStatisticsStatisticsData { return v.StatisticsDatas }).(GetProxyStatisticsStatisticsDataArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProxyStatisticsResultOutput{})
}
