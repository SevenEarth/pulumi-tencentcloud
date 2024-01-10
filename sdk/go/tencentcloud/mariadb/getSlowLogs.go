// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of mariadb slowLogs
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mariadb.GetSlowLogs(ctx, &mariadb.GetSlowLogsArgs{
//				InstanceId:  "tdsql-9vqvls95",
//				OrderBy:     pulumi.StringRef("query_time_sum"),
//				OrderByType: pulumi.StringRef("desc"),
//				Slave:       pulumi.IntRef(0),
//				StartTime:   "2023-06-01 14:55:20",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSlowLogs(ctx *pulumi.Context, args *GetSlowLogsArgs, opts ...pulumi.InvokeOption) (*GetSlowLogsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSlowLogsResult
	err := ctx.Invoke("tencentcloud:Mariadb/getSlowLogs:getSlowLogs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSlowLogs.
type GetSlowLogsArgs struct {
	// Specific name of the database to be queried.
	Db *string `pulumi:"db"`
	// Query end time in the format of 2016-08-22 14:55:20.
	EndTime *string `pulumi:"endTime"`
	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId string `pulumi:"instanceId"`
	// Sorting metric. Valid values: query_time_sum, query_count.
	OrderBy *string `pulumi:"orderBy"`
	// Sorting order. Valid values: desc, asc.
	OrderByType *string `pulumi:"orderByType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Query slow queries from either the primary or the replica. Valid values: 0 (primary), 1 (replica).
	Slave *int `pulumi:"slave"`
	// Query start time in the format of 2016-07-23 14:55:20.
	StartTime string `pulumi:"startTime"`
}

// A collection of values returned by getSlowLogs.
type GetSlowLogsResult struct {
	// Slow query log data.
	Datas []GetSlowLogsData `pulumi:"datas"`
	// Database name.
	Db      *string `pulumi:"db"`
	EndTime *string `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// Total statement lock time.
	LockTimeSum float64 `pulumi:"lockTimeSum"`
	OrderBy     *string `pulumi:"orderBy"`
	OrderByType *string `pulumi:"orderByType"`
	// Total number of statement queries.
	QueryCount int `pulumi:"queryCount"`
	// Total statement query time.
	QueryTimeSum     float64 `pulumi:"queryTimeSum"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	Slave            *int    `pulumi:"slave"`
	StartTime        string  `pulumi:"startTime"`
}

func GetSlowLogsOutput(ctx *pulumi.Context, args GetSlowLogsOutputArgs, opts ...pulumi.InvokeOption) GetSlowLogsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSlowLogsResult, error) {
			args := v.(GetSlowLogsArgs)
			r, err := GetSlowLogs(ctx, &args, opts...)
			var s GetSlowLogsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSlowLogsResultOutput)
}

// A collection of arguments for invoking getSlowLogs.
type GetSlowLogsOutputArgs struct {
	// Specific name of the database to be queried.
	Db pulumi.StringPtrInput `pulumi:"db"`
	// Query end time in the format of 2016-08-22 14:55:20.
	EndTime pulumi.StringPtrInput `pulumi:"endTime"`
	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Sorting metric. Valid values: query_time_sum, query_count.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// Sorting order. Valid values: desc, asc.
	OrderByType pulumi.StringPtrInput `pulumi:"orderByType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Query slow queries from either the primary or the replica. Valid values: 0 (primary), 1 (replica).
	Slave pulumi.IntPtrInput `pulumi:"slave"`
	// Query start time in the format of 2016-07-23 14:55:20.
	StartTime pulumi.StringInput `pulumi:"startTime"`
}

func (GetSlowLogsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlowLogsArgs)(nil)).Elem()
}

// A collection of values returned by getSlowLogs.
type GetSlowLogsResultOutput struct{ *pulumi.OutputState }

func (GetSlowLogsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlowLogsResult)(nil)).Elem()
}

func (o GetSlowLogsResultOutput) ToGetSlowLogsResultOutput() GetSlowLogsResultOutput {
	return o
}

func (o GetSlowLogsResultOutput) ToGetSlowLogsResultOutputWithContext(ctx context.Context) GetSlowLogsResultOutput {
	return o
}

// Slow query log data.
func (o GetSlowLogsResultOutput) Datas() GetSlowLogsDataArrayOutput {
	return o.ApplyT(func(v GetSlowLogsResult) []GetSlowLogsData { return v.Datas }).(GetSlowLogsDataArrayOutput)
}

// Database name.
func (o GetSlowLogsResultOutput) Db() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogsResult) *string { return v.Db }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogsResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogsResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSlowLogsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSlowLogsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Total statement lock time.
func (o GetSlowLogsResultOutput) LockTimeSum() pulumi.Float64Output {
	return o.ApplyT(func(v GetSlowLogsResult) float64 { return v.LockTimeSum }).(pulumi.Float64Output)
}

func (o GetSlowLogsResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogsResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogsResultOutput) OrderByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogsResult) *string { return v.OrderByType }).(pulumi.StringPtrOutput)
}

// Total number of statement queries.
func (o GetSlowLogsResultOutput) QueryCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetSlowLogsResult) int { return v.QueryCount }).(pulumi.IntOutput)
}

// Total statement query time.
func (o GetSlowLogsResultOutput) QueryTimeSum() pulumi.Float64Output {
	return o.ApplyT(func(v GetSlowLogsResult) float64 { return v.QueryTimeSum }).(pulumi.Float64Output)
}

func (o GetSlowLogsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogsResultOutput) Slave() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetSlowLogsResult) *int { return v.Slave }).(pulumi.IntPtrOutput)
}

func (o GetSlowLogsResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogsResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSlowLogsResultOutput{})
}
