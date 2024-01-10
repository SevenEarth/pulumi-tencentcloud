// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of monitor alarmConditionsTemplate
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Monitor.GetAlarmConditionsTemplate(ctx, &monitor.GetAlarmConditionsTemplateArgs{
//				GroupId:          pulumi.StringRef("7803070"),
//				GroupName:        pulumi.StringRef("keep-template"),
//				Module:           "monitor",
//				PolicyCountOrder: pulumi.StringRef("asc=ascending"),
//				UpdateTimeOrder:  pulumi.StringRef("desc=descending"),
//				ViewName:         pulumi.StringRef("cvm_device"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetAlarmConditionsTemplate(ctx *pulumi.Context, args *GetAlarmConditionsTemplateArgs, opts ...pulumi.InvokeOption) (*GetAlarmConditionsTemplateResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetAlarmConditionsTemplateResult
	err := ctx.Invoke("tencentcloud:Monitor/getAlarmConditionsTemplate:getAlarmConditionsTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlarmConditionsTemplate.
type GetAlarmConditionsTemplateArgs struct {
	// Filter queries based on trigger condition template ID.
	GroupId *string `pulumi:"groupId"`
	// Filter queries based on trigger condition template names.
	GroupName *string `pulumi:"groupName"`
	// Fixed value, as&amp;amp;#39; monitor &amp;amp;#39;.
	Module string `pulumi:"module"`
	// Specify the sorting method based on the number of binding policies, asc=ascending, desc=descending.
	PolicyCountOrder *string `pulumi:"policyCountOrder"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Specify the sorting method by update time, asc=ascending, desc=descending.
	UpdateTimeOrder *string `pulumi:"updateTimeOrder"`
	// View name, composed of DescribeAllNamespacesObtain. For cloud product monitoring, retrieve the QceNamespacesNew. N.ID parameter from the interface, such as cvm_ Device.
	ViewName *string `pulumi:"viewName"`
}

// A collection of values returned by getAlarmConditionsTemplate.
type GetAlarmConditionsTemplateResult struct {
	// Alarm Policy Group ID.
	GroupId *string `pulumi:"groupId"`
	// Alarm Policy Group Name.
	GroupName *string `pulumi:"groupName"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	Module           string  `pulumi:"module"`
	PolicyCountOrder *string `pulumi:"policyCountOrder"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Template List.
	TemplateGroupLists []GetAlarmConditionsTemplateTemplateGroupList `pulumi:"templateGroupLists"`
	UpdateTimeOrder    *string                                       `pulumi:"updateTimeOrder"`
	// View.
	ViewName *string `pulumi:"viewName"`
}

func GetAlarmConditionsTemplateOutput(ctx *pulumi.Context, args GetAlarmConditionsTemplateOutputArgs, opts ...pulumi.InvokeOption) GetAlarmConditionsTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAlarmConditionsTemplateResult, error) {
			args := v.(GetAlarmConditionsTemplateArgs)
			r, err := GetAlarmConditionsTemplate(ctx, &args, opts...)
			var s GetAlarmConditionsTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAlarmConditionsTemplateResultOutput)
}

// A collection of arguments for invoking getAlarmConditionsTemplate.
type GetAlarmConditionsTemplateOutputArgs struct {
	// Filter queries based on trigger condition template ID.
	GroupId pulumi.StringPtrInput `pulumi:"groupId"`
	// Filter queries based on trigger condition template names.
	GroupName pulumi.StringPtrInput `pulumi:"groupName"`
	// Fixed value, as&amp;amp;#39; monitor &amp;amp;#39;.
	Module pulumi.StringInput `pulumi:"module"`
	// Specify the sorting method based on the number of binding policies, asc=ascending, desc=descending.
	PolicyCountOrder pulumi.StringPtrInput `pulumi:"policyCountOrder"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Specify the sorting method by update time, asc=ascending, desc=descending.
	UpdateTimeOrder pulumi.StringPtrInput `pulumi:"updateTimeOrder"`
	// View name, composed of DescribeAllNamespacesObtain. For cloud product monitoring, retrieve the QceNamespacesNew. N.ID parameter from the interface, such as cvm_ Device.
	ViewName pulumi.StringPtrInput `pulumi:"viewName"`
}

func (GetAlarmConditionsTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlarmConditionsTemplateArgs)(nil)).Elem()
}

// A collection of values returned by getAlarmConditionsTemplate.
type GetAlarmConditionsTemplateResultOutput struct{ *pulumi.OutputState }

func (GetAlarmConditionsTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlarmConditionsTemplateResult)(nil)).Elem()
}

func (o GetAlarmConditionsTemplateResultOutput) ToGetAlarmConditionsTemplateResultOutput() GetAlarmConditionsTemplateResultOutput {
	return o
}

func (o GetAlarmConditionsTemplateResultOutput) ToGetAlarmConditionsTemplateResultOutputWithContext(ctx context.Context) GetAlarmConditionsTemplateResultOutput {
	return o
}

// Alarm Policy Group ID.
func (o GetAlarmConditionsTemplateResultOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmConditionsTemplateResult) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

// Alarm Policy Group Name.
func (o GetAlarmConditionsTemplateResultOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmConditionsTemplateResult) *string { return v.GroupName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAlarmConditionsTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlarmConditionsTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAlarmConditionsTemplateResultOutput) Module() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlarmConditionsTemplateResult) string { return v.Module }).(pulumi.StringOutput)
}

func (o GetAlarmConditionsTemplateResultOutput) PolicyCountOrder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmConditionsTemplateResult) *string { return v.PolicyCountOrder }).(pulumi.StringPtrOutput)
}

func (o GetAlarmConditionsTemplateResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmConditionsTemplateResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Template List.
func (o GetAlarmConditionsTemplateResultOutput) TemplateGroupLists() GetAlarmConditionsTemplateTemplateGroupListArrayOutput {
	return o.ApplyT(func(v GetAlarmConditionsTemplateResult) []GetAlarmConditionsTemplateTemplateGroupList {
		return v.TemplateGroupLists
	}).(GetAlarmConditionsTemplateTemplateGroupListArrayOutput)
}

func (o GetAlarmConditionsTemplateResultOutput) UpdateTimeOrder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmConditionsTemplateResult) *string { return v.UpdateTimeOrder }).(pulumi.StringPtrOutput)
}

// View.
func (o GetAlarmConditionsTemplateResultOutput) ViewName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmConditionsTemplateResult) *string { return v.ViewName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAlarmConditionsTemplateResultOutput{})
}
