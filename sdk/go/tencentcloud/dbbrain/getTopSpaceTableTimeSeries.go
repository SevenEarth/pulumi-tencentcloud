// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbbrain

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of dbbrain topSpaceTableTimeSeries
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dbbrain.GetTopSpaceTableTimeSeries(ctx, &dbbrain.GetTopSpaceTableTimeSeriesArgs{
//				EndDate:    pulumi.StringRef("%s"),
//				InstanceId: "%s",
//				Product:    pulumi.StringRef("mysql"),
//				SortBy:     pulumi.StringRef("DataLength"),
//				StartDate:  pulumi.StringRef("%s"),
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
func GetTopSpaceTableTimeSeries(ctx *pulumi.Context, args *GetTopSpaceTableTimeSeriesArgs, opts ...pulumi.InvokeOption) (*GetTopSpaceTableTimeSeriesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTopSpaceTableTimeSeriesResult
	err := ctx.Invoke("tencentcloud:Dbbrain/getTopSpaceTableTimeSeries:getTopSpaceTableTimeSeries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTopSpaceTableTimeSeries.
type GetTopSpaceTableTimeSeriesArgs struct {
	// The deadline, such as 2021-01-01, the earliest is the 29th day before the current day, and the default is the current day.
	EndDate *string `pulumi:"endDate"`
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// The number of Top tables returned, the maximum value is 100, and the default is 20.
	Limit *int `pulumi:"limit"`
	// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
	Product *string `pulumi:"product"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// The sorting field used to filter the Top table. The optional fields include DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, and PhysicalFileSize. The default is PhysicalFileSize.
	SortBy *string `pulumi:"sortBy"`
	// The start date, such as 2021-01-01, the earliest is the 29th day before the current day, and the default is the 6th day before the deadline.
	StartDate *string `pulumi:"startDate"`
}

// A collection of values returned by getTopSpaceTableTimeSeries.
type GetTopSpaceTableTimeSeriesResult struct {
	EndDate *string `pulumi:"endDate"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	InstanceId       string  `pulumi:"instanceId"`
	Limit            *int    `pulumi:"limit"`
	Product          *string `pulumi:"product"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	SortBy           *string `pulumi:"sortBy"`
	StartDate        *string `pulumi:"startDate"`
	// The time-series data list of the returned Top tablespace statistics.
	TopSpaceTableTimeSeries []GetTopSpaceTableTimeSeriesTopSpaceTableTimeSeries `pulumi:"topSpaceTableTimeSeries"`
}

func GetTopSpaceTableTimeSeriesOutput(ctx *pulumi.Context, args GetTopSpaceTableTimeSeriesOutputArgs, opts ...pulumi.InvokeOption) GetTopSpaceTableTimeSeriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTopSpaceTableTimeSeriesResult, error) {
			args := v.(GetTopSpaceTableTimeSeriesArgs)
			r, err := GetTopSpaceTableTimeSeries(ctx, &args, opts...)
			var s GetTopSpaceTableTimeSeriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTopSpaceTableTimeSeriesResultOutput)
}

// A collection of arguments for invoking getTopSpaceTableTimeSeries.
type GetTopSpaceTableTimeSeriesOutputArgs struct {
	// The deadline, such as 2021-01-01, the earliest is the 29th day before the current day, and the default is the current day.
	EndDate pulumi.StringPtrInput `pulumi:"endDate"`
	// instance id.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// The number of Top tables returned, the maximum value is 100, and the default is 20.
	Limit pulumi.IntPtrInput `pulumi:"limit"`
	// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
	Product pulumi.StringPtrInput `pulumi:"product"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// The sorting field used to filter the Top table. The optional fields include DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, and PhysicalFileSize. The default is PhysicalFileSize.
	SortBy pulumi.StringPtrInput `pulumi:"sortBy"`
	// The start date, such as 2021-01-01, the earliest is the 29th day before the current day, and the default is the 6th day before the deadline.
	StartDate pulumi.StringPtrInput `pulumi:"startDate"`
}

func (GetTopSpaceTableTimeSeriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTopSpaceTableTimeSeriesArgs)(nil)).Elem()
}

// A collection of values returned by getTopSpaceTableTimeSeries.
type GetTopSpaceTableTimeSeriesResultOutput struct{ *pulumi.OutputState }

func (GetTopSpaceTableTimeSeriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTopSpaceTableTimeSeriesResult)(nil)).Elem()
}

func (o GetTopSpaceTableTimeSeriesResultOutput) ToGetTopSpaceTableTimeSeriesResultOutput() GetTopSpaceTableTimeSeriesResultOutput {
	return o
}

func (o GetTopSpaceTableTimeSeriesResultOutput) ToGetTopSpaceTableTimeSeriesResultOutputWithContext(ctx context.Context) GetTopSpaceTableTimeSeriesResultOutput {
	return o
}

func (o GetTopSpaceTableTimeSeriesResultOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTopSpaceTableTimeSeriesResult) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTopSpaceTableTimeSeriesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTopSpaceTableTimeSeriesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTopSpaceTableTimeSeriesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTopSpaceTableTimeSeriesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetTopSpaceTableTimeSeriesResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetTopSpaceTableTimeSeriesResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

func (o GetTopSpaceTableTimeSeriesResultOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTopSpaceTableTimeSeriesResult) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o GetTopSpaceTableTimeSeriesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTopSpaceTableTimeSeriesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetTopSpaceTableTimeSeriesResultOutput) SortBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTopSpaceTableTimeSeriesResult) *string { return v.SortBy }).(pulumi.StringPtrOutput)
}

func (o GetTopSpaceTableTimeSeriesResultOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTopSpaceTableTimeSeriesResult) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

// The time-series data list of the returned Top tablespace statistics.
func (o GetTopSpaceTableTimeSeriesResultOutput) TopSpaceTableTimeSeries() GetTopSpaceTableTimeSeriesTopSpaceTableTimeSeriesArrayOutput {
	return o.ApplyT(func(v GetTopSpaceTableTimeSeriesResult) []GetTopSpaceTableTimeSeriesTopSpaceTableTimeSeries {
		return v.TopSpaceTableTimeSeries
	}).(GetTopSpaceTableTimeSeriesTopSpaceTableTimeSeriesArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTopSpaceTableTimeSeriesResultOutput{})
}
