// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbbrain

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of dbbrain diagEvent
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
// 		diagHistory, err := Dbbrain.GetDiagHistory(ctx, &dbbrain.GetDiagHistoryArgs{
// 			InstanceId: fmt.Sprintf("%v%v", "%", "s"),
// 			StartTime:  fmt.Sprintf("%v%v", "%", "s"),
// 			EndTime:    fmt.Sprintf("%v%v", "%", "s"),
// 			Product:    pulumi.StringRef("mysql"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Dbbrain.GetDiagEvent(ctx, &dbbrain.GetDiagEventArgs{
// 			InstanceId: fmt.Sprintf("%v%v", "%", "s"),
// 			EventId:    pulumi.IntRef(diagHistory.Events[0].EventId),
// 			Product:    pulumi.StringRef("mysql"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDiagEvent(ctx *pulumi.Context, args *GetDiagEventArgs, opts ...pulumi.InvokeOption) (*GetDiagEventResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDiagEventResult
	err := ctx.Invoke("tencentcloud:Dbbrain/getDiagEvent:getDiagEvent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDiagEvent.
type GetDiagEventArgs struct {
	// Event ID. Obtain it through `Get Instance Diagnosis History DescribeDBDiagHistory`.
	EventId *int `pulumi:"eventId"`
	// isntance id.
	InstanceId string `pulumi:"instanceId"`
	// Service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database CynosDB for MySQL, the default is `mysql`.
	Product *string `pulumi:"product"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDiagEvent.
type GetDiagEventResult struct {
	// diagnostic item.
	DiagItem string `pulumi:"diagItem"`
	// Diagnostic type.
	DiagType string `pulumi:"diagType"`
	// End Time.
	EndTime string `pulumi:"endTime"`
	EventId int    `pulumi:"eventId"`
	// Diagnostic event details, output is empty if there is no additional explanatory information.
	Explanation string `pulumi:"explanation"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// reserved text. Note: This field may return null, indicating that no valid value can be obtained.
	Metric string `pulumi:"metric"`
	// Diagnostic summary.
	Outline string `pulumi:"outline"`
	// Diagnosed problem.
	Problem          string  `pulumi:"problem"`
	Product          *string `pulumi:"product"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// severity. The severity is divided into 5 levels, according to the degree of impact from high to low: 1: Fatal, 2: Serious, 3: Warning, 4: Prompt, 5: Healthy.
	Severity int `pulumi:"severity"`
	// Starting time.
	StartTime string `pulumi:"startTime"`
	// A diagnostic suggestion, or empty if there is no suggestion.
	Suggestions string `pulumi:"suggestions"`
}

func GetDiagEventOutput(ctx *pulumi.Context, args GetDiagEventOutputArgs, opts ...pulumi.InvokeOption) GetDiagEventResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDiagEventResult, error) {
			args := v.(GetDiagEventArgs)
			r, err := GetDiagEvent(ctx, &args, opts...)
			var s GetDiagEventResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDiagEventResultOutput)
}

// A collection of arguments for invoking getDiagEvent.
type GetDiagEventOutputArgs struct {
	// Event ID. Obtain it through `Get Instance Diagnosis History DescribeDBDiagHistory`.
	EventId pulumi.IntPtrInput `pulumi:"eventId"`
	// isntance id.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database CynosDB for MySQL, the default is `mysql`.
	Product pulumi.StringPtrInput `pulumi:"product"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDiagEventOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDiagEventArgs)(nil)).Elem()
}

// A collection of values returned by getDiagEvent.
type GetDiagEventResultOutput struct{ *pulumi.OutputState }

func (GetDiagEventResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDiagEventResult)(nil)).Elem()
}

func (o GetDiagEventResultOutput) ToGetDiagEventResultOutput() GetDiagEventResultOutput {
	return o
}

func (o GetDiagEventResultOutput) ToGetDiagEventResultOutputWithContext(ctx context.Context) GetDiagEventResultOutput {
	return o
}

// diagnostic item.
func (o GetDiagEventResultOutput) DiagItem() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiagEventResult) string { return v.DiagItem }).(pulumi.StringOutput)
}

// Diagnostic type.
func (o GetDiagEventResultOutput) DiagType() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiagEventResult) string { return v.DiagType }).(pulumi.StringOutput)
}

// End Time.
func (o GetDiagEventResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiagEventResult) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o GetDiagEventResultOutput) EventId() pulumi.IntOutput {
	return o.ApplyT(func(v GetDiagEventResult) int { return v.EventId }).(pulumi.IntOutput)
}

// Diagnostic event details, output is empty if there is no additional explanatory information.
func (o GetDiagEventResultOutput) Explanation() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiagEventResult) string { return v.Explanation }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDiagEventResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiagEventResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDiagEventResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiagEventResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// reserved text. Note: This field may return null, indicating that no valid value can be obtained.
func (o GetDiagEventResultOutput) Metric() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiagEventResult) string { return v.Metric }).(pulumi.StringOutput)
}

// Diagnostic summary.
func (o GetDiagEventResultOutput) Outline() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiagEventResult) string { return v.Outline }).(pulumi.StringOutput)
}

// Diagnosed problem.
func (o GetDiagEventResultOutput) Problem() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiagEventResult) string { return v.Problem }).(pulumi.StringOutput)
}

func (o GetDiagEventResultOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDiagEventResult) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o GetDiagEventResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDiagEventResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// severity. The severity is divided into 5 levels, according to the degree of impact from high to low: 1: Fatal, 2: Serious, 3: Warning, 4: Prompt, 5: Healthy.
func (o GetDiagEventResultOutput) Severity() pulumi.IntOutput {
	return o.ApplyT(func(v GetDiagEventResult) int { return v.Severity }).(pulumi.IntOutput)
}

// Starting time.
func (o GetDiagEventResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiagEventResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// A diagnostic suggestion, or empty if there is no suggestion.
func (o GetDiagEventResultOutput) Suggestions() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiagEventResult) string { return v.Suggestions }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDiagEventResultOutput{})
}
