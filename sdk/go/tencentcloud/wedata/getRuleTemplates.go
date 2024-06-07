// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wedata

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of wedata rule templates
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Wedata"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// _, err := Wedata.GetRuleTemplates(ctx, &wedata.GetRuleTemplatesArgs{
// ProjectId: pulumi.StringRef("1840731346428280832"),
// SourceEngineTypes: interface{}{
// 2,
// 4,
// 16,
// },
// SourceObjectType: pulumi.IntRef(2),
// Type: pulumi.IntRef(2),
// }, nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
// <!--End PulumiCodeChooser -->
func GetRuleTemplates(ctx *pulumi.Context, args *GetRuleTemplatesArgs, opts ...pulumi.InvokeOption) (*GetRuleTemplatesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRuleTemplatesResult
	err := ctx.Invoke("tencentcloud:Wedata/getRuleTemplates:getRuleTemplates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRuleTemplates.
type GetRuleTemplatesArgs struct {
	// Project ID.
	ProjectId *string `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Applicable type of source data.
	SourceEngineTypes []int `pulumi:"sourceEngineTypes"`
	// Source data object type. `1`: Constant, `2`: Offline table level, `3`: Offline field level.
	SourceObjectType *int `pulumi:"sourceObjectType"`
	// Template type. `1` means System template, `2` means Custom template.
	Type *int `pulumi:"type"`
}

// A collection of values returned by getRuleTemplates.
type GetRuleTemplatesResult struct {
	// rule template list.
	Datas []GetRuleTemplatesData `pulumi:"datas"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ProjectId        *string `pulumi:"projectId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Applicable type of source data.
	SourceEngineTypes []int `pulumi:"sourceEngineTypes"`
	// Source object type. `1`: Constant, `2`: Offline table level, `3`: Offline field level.
	SourceObjectType *int `pulumi:"sourceObjectType"`
	// Template type. `1` means System template, `2` means Custom template.
	Type *int `pulumi:"type"`
}

func GetRuleTemplatesOutput(ctx *pulumi.Context, args GetRuleTemplatesOutputArgs, opts ...pulumi.InvokeOption) GetRuleTemplatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRuleTemplatesResult, error) {
			args := v.(GetRuleTemplatesArgs)
			r, err := GetRuleTemplates(ctx, &args, opts...)
			var s GetRuleTemplatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRuleTemplatesResultOutput)
}

// A collection of arguments for invoking getRuleTemplates.
type GetRuleTemplatesOutputArgs struct {
	// Project ID.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Applicable type of source data.
	SourceEngineTypes pulumi.IntArrayInput `pulumi:"sourceEngineTypes"`
	// Source data object type. `1`: Constant, `2`: Offline table level, `3`: Offline field level.
	SourceObjectType pulumi.IntPtrInput `pulumi:"sourceObjectType"`
	// Template type. `1` means System template, `2` means Custom template.
	Type pulumi.IntPtrInput `pulumi:"type"`
}

func (GetRuleTemplatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRuleTemplatesArgs)(nil)).Elem()
}

// A collection of values returned by getRuleTemplates.
type GetRuleTemplatesResultOutput struct{ *pulumi.OutputState }

func (GetRuleTemplatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRuleTemplatesResult)(nil)).Elem()
}

func (o GetRuleTemplatesResultOutput) ToGetRuleTemplatesResultOutput() GetRuleTemplatesResultOutput {
	return o
}

func (o GetRuleTemplatesResultOutput) ToGetRuleTemplatesResultOutputWithContext(ctx context.Context) GetRuleTemplatesResultOutput {
	return o
}

// rule template list.
func (o GetRuleTemplatesResultOutput) Datas() GetRuleTemplatesDataArrayOutput {
	return o.ApplyT(func(v GetRuleTemplatesResult) []GetRuleTemplatesData { return v.Datas }).(GetRuleTemplatesDataArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRuleTemplatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRuleTemplatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRuleTemplatesResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRuleTemplatesResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o GetRuleTemplatesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRuleTemplatesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Applicable type of source data.
func (o GetRuleTemplatesResultOutput) SourceEngineTypes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetRuleTemplatesResult) []int { return v.SourceEngineTypes }).(pulumi.IntArrayOutput)
}

// Source object type. `1`: Constant, `2`: Offline table level, `3`: Offline field level.
func (o GetRuleTemplatesResultOutput) SourceObjectType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetRuleTemplatesResult) *int { return v.SourceObjectType }).(pulumi.IntPtrOutput)
}

// Template type. `1` means System template, `2` means Custom template.
func (o GetRuleTemplatesResultOutput) Type() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetRuleTemplatesResult) *int { return v.Type }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRuleTemplatesResultOutput{})
}
