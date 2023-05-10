// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbbrain

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of dbbrain slowLogTimeSeriesStats
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dbbrain.GetSlowLogTimeSeriesStats(ctx, &dbbrain.GetSlowLogTimeSeriesStatsArgs{
// 			EndTime:    fmt.Sprintf("%v%v", "%", "s"),
// 			InstanceId: fmt.Sprintf("%v%v", "%", "s"),
// 			Product:    pulumi.StringRef("mysql"),
// 			StartTime:  fmt.Sprintf("%v%v", "%", "s"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetSlowLogTimeSeriesStats(ctx *pulumi.Context, args *GetSlowLogTimeSeriesStatsArgs, opts ...pulumi.InvokeOption) (*GetSlowLogTimeSeriesStatsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSlowLogTimeSeriesStatsResult
	err := ctx.Invoke("tencentcloud:Dbbrain/getSlowLogTimeSeriesStats:getSlowLogTimeSeriesStats", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSlowLogTimeSeriesStats.
type GetSlowLogTimeSeriesStatsArgs struct {
	// End time, such as `2019-09-10 12:13:14`, the interval between the end time and the start time can be up to 7 days.
	EndTime string `pulumi:"endTime"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database CynosDB for MySQL, the default is `mysql`.
	Product *string `pulumi:"product"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Start time, such as `2019-09-10 12:13:14`.
	StartTime string `pulumi:"startTime"`
}

// A collection of values returned by getSlowLogTimeSeriesStats.
type GetSlowLogTimeSeriesStatsResult struct {
	EndTime string `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// The unit time interval between bars, in seconds.
	Period           int     `pulumi:"period"`
	Product          *string `pulumi:"product"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Instan1ce cpu utilization monitoring data within a unit time interval.
	SeriesDatas []GetSlowLogTimeSeriesStatsSeriesData `pulumi:"seriesDatas"`
	StartTime   string                                `pulumi:"startTime"`
	// Statistics on the number of slow logs in a unit time interval.
	TimeSeries []GetSlowLogTimeSeriesStatsTimeSeries `pulumi:"timeSeries"`
}

func GetSlowLogTimeSeriesStatsOutput(ctx *pulumi.Context, args GetSlowLogTimeSeriesStatsOutputArgs, opts ...pulumi.InvokeOption) GetSlowLogTimeSeriesStatsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSlowLogTimeSeriesStatsResult, error) {
			args := v.(GetSlowLogTimeSeriesStatsArgs)
			r, err := GetSlowLogTimeSeriesStats(ctx, &args, opts...)
			var s GetSlowLogTimeSeriesStatsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSlowLogTimeSeriesStatsResultOutput)
}

// A collection of arguments for invoking getSlowLogTimeSeriesStats.
type GetSlowLogTimeSeriesStatsOutputArgs struct {
	// End time, such as `2019-09-10 12:13:14`, the interval between the end time and the start time can be up to 7 days.
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// Instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database CynosDB for MySQL, the default is `mysql`.
	Product pulumi.StringPtrInput `pulumi:"product"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Start time, such as `2019-09-10 12:13:14`.
	StartTime pulumi.StringInput `pulumi:"startTime"`
}

func (GetSlowLogTimeSeriesStatsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlowLogTimeSeriesStatsArgs)(nil)).Elem()
}

// A collection of values returned by getSlowLogTimeSeriesStats.
type GetSlowLogTimeSeriesStatsResultOutput struct{ *pulumi.OutputState }

func (GetSlowLogTimeSeriesStatsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlowLogTimeSeriesStatsResult)(nil)).Elem()
}

func (o GetSlowLogTimeSeriesStatsResultOutput) ToGetSlowLogTimeSeriesStatsResultOutput() GetSlowLogTimeSeriesStatsResultOutput {
	return o
}

func (o GetSlowLogTimeSeriesStatsResultOutput) ToGetSlowLogTimeSeriesStatsResultOutputWithContext(ctx context.Context) GetSlowLogTimeSeriesStatsResultOutput {
	return o
}

func (o GetSlowLogTimeSeriesStatsResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogTimeSeriesStatsResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSlowLogTimeSeriesStatsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogTimeSeriesStatsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSlowLogTimeSeriesStatsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogTimeSeriesStatsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The unit time interval between bars, in seconds.
func (o GetSlowLogTimeSeriesStatsResultOutput) Period() pulumi.IntOutput {
	return o.ApplyT(func(v GetSlowLogTimeSeriesStatsResult) int { return v.Period }).(pulumi.IntOutput)
}

func (o GetSlowLogTimeSeriesStatsResultOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogTimeSeriesStatsResult) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogTimeSeriesStatsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogTimeSeriesStatsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Instan1ce cpu utilization monitoring data within a unit time interval.
func (o GetSlowLogTimeSeriesStatsResultOutput) SeriesDatas() GetSlowLogTimeSeriesStatsSeriesDataArrayOutput {
	return o.ApplyT(func(v GetSlowLogTimeSeriesStatsResult) []GetSlowLogTimeSeriesStatsSeriesData { return v.SeriesDatas }).(GetSlowLogTimeSeriesStatsSeriesDataArrayOutput)
}

func (o GetSlowLogTimeSeriesStatsResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogTimeSeriesStatsResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// Statistics on the number of slow logs in a unit time interval.
func (o GetSlowLogTimeSeriesStatsResultOutput) TimeSeries() GetSlowLogTimeSeriesStatsTimeSeriesArrayOutput {
	return o.ApplyT(func(v GetSlowLogTimeSeriesStatsResult) []GetSlowLogTimeSeriesStatsTimeSeries { return v.TimeSeries }).(GetSlowLogTimeSeriesStatsTimeSeriesArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSlowLogTimeSeriesStatsResultOutput{})
}