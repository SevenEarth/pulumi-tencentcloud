// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of monitor alarmMetric
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Monitor.GetAlarmMetric(ctx, &monitor.GetAlarmMetricArgs{
//				Module:      "monitor",
//				MonitorType: "Monitoring",
//				Namespace:   "cvm_device",
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
func GetAlarmMetric(ctx *pulumi.Context, args *GetAlarmMetricArgs, opts ...pulumi.InvokeOption) (*GetAlarmMetricResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAlarmMetricResult
	err := ctx.Invoke("tencentcloud:Monitor/getAlarmMetric:getAlarmMetric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlarmMetric.
type GetAlarmMetricArgs struct {
	// Fixed value, as `monitor`.
	Module string `pulumi:"module"`
	// Monitoring Type Filter MT_QCE=Cloud Product Monitoring.
	MonitorType string `pulumi:"monitorType"`
	// Alarm policy type, obtained from DescribeAllNamespaces, such as cvm_device.
	Namespace string `pulumi:"namespace"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getAlarmMetric.
type GetAlarmMetricResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Alarm indicator list.
	Metrics     []GetAlarmMetricMetric `pulumi:"metrics"`
	Module      string                 `pulumi:"module"`
	MonitorType string                 `pulumi:"monitorType"`
	// Alarm strategy type.
	Namespace        string  `pulumi:"namespace"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetAlarmMetricOutput(ctx *pulumi.Context, args GetAlarmMetricOutputArgs, opts ...pulumi.InvokeOption) GetAlarmMetricResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAlarmMetricResult, error) {
			args := v.(GetAlarmMetricArgs)
			r, err := GetAlarmMetric(ctx, &args, opts...)
			var s GetAlarmMetricResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAlarmMetricResultOutput)
}

// A collection of arguments for invoking getAlarmMetric.
type GetAlarmMetricOutputArgs struct {
	// Fixed value, as `monitor`.
	Module pulumi.StringInput `pulumi:"module"`
	// Monitoring Type Filter MT_QCE=Cloud Product Monitoring.
	MonitorType pulumi.StringInput `pulumi:"monitorType"`
	// Alarm policy type, obtained from DescribeAllNamespaces, such as cvm_device.
	Namespace pulumi.StringInput `pulumi:"namespace"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetAlarmMetricOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlarmMetricArgs)(nil)).Elem()
}

// A collection of values returned by getAlarmMetric.
type GetAlarmMetricResultOutput struct{ *pulumi.OutputState }

func (GetAlarmMetricResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlarmMetricResult)(nil)).Elem()
}

func (o GetAlarmMetricResultOutput) ToGetAlarmMetricResultOutput() GetAlarmMetricResultOutput {
	return o
}

func (o GetAlarmMetricResultOutput) ToGetAlarmMetricResultOutputWithContext(ctx context.Context) GetAlarmMetricResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetAlarmMetricResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlarmMetricResult) string { return v.Id }).(pulumi.StringOutput)
}

// Alarm indicator list.
func (o GetAlarmMetricResultOutput) Metrics() GetAlarmMetricMetricArrayOutput {
	return o.ApplyT(func(v GetAlarmMetricResult) []GetAlarmMetricMetric { return v.Metrics }).(GetAlarmMetricMetricArrayOutput)
}

func (o GetAlarmMetricResultOutput) Module() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlarmMetricResult) string { return v.Module }).(pulumi.StringOutput)
}

func (o GetAlarmMetricResultOutput) MonitorType() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlarmMetricResult) string { return v.MonitorType }).(pulumi.StringOutput)
}

// Alarm strategy type.
func (o GetAlarmMetricResultOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlarmMetricResult) string { return v.Namespace }).(pulumi.StringOutput)
}

func (o GetAlarmMetricResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmMetricResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAlarmMetricResultOutput{})
}
