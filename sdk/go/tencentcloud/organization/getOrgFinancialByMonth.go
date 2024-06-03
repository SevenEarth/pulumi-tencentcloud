// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of organization orgFinancialByMonth
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Organization"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// _, err := Organization.GetOrgFinancialByMonth(ctx, &organization.GetOrgFinancialByMonthArgs{
// EndMonth: pulumi.StringRef("2023-05"),
// MemberUins: interface{}{
// 100026517717,
// },
// }, nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
// <!--End PulumiCodeChooser -->
func GetOrgFinancialByMonth(ctx *pulumi.Context, args *GetOrgFinancialByMonthArgs, opts ...pulumi.InvokeOption) (*GetOrgFinancialByMonthResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOrgFinancialByMonthResult
	err := ctx.Invoke("tencentcloud:Organization/getOrgFinancialByMonth:getOrgFinancialByMonth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrgFinancialByMonth.
type GetOrgFinancialByMonthArgs struct {
	// Query for the end month. Format:yyyy-mm, for example:2021-01.
	EndMonth *string `pulumi:"endMonth"`
	// Member uin list. Up to 100.
	MemberUins []int `pulumi:"memberUins"`
	// Product code list. Up to 100.
	ProductCodes []string `pulumi:"productCodes"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getOrgFinancialByMonth.
type GetOrgFinancialByMonthResult struct {
	EndMonth *string `pulumi:"endMonth"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Organization financial info by month.
	Items            []GetOrgFinancialByMonthItem `pulumi:"items"`
	MemberUins       []int                        `pulumi:"memberUins"`
	ProductCodes     []string                     `pulumi:"productCodes"`
	ResultOutputFile *string                      `pulumi:"resultOutputFile"`
}

func GetOrgFinancialByMonthOutput(ctx *pulumi.Context, args GetOrgFinancialByMonthOutputArgs, opts ...pulumi.InvokeOption) GetOrgFinancialByMonthResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOrgFinancialByMonthResult, error) {
			args := v.(GetOrgFinancialByMonthArgs)
			r, err := GetOrgFinancialByMonth(ctx, &args, opts...)
			var s GetOrgFinancialByMonthResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOrgFinancialByMonthResultOutput)
}

// A collection of arguments for invoking getOrgFinancialByMonth.
type GetOrgFinancialByMonthOutputArgs struct {
	// Query for the end month. Format:yyyy-mm, for example:2021-01.
	EndMonth pulumi.StringPtrInput `pulumi:"endMonth"`
	// Member uin list. Up to 100.
	MemberUins pulumi.IntArrayInput `pulumi:"memberUins"`
	// Product code list. Up to 100.
	ProductCodes pulumi.StringArrayInput `pulumi:"productCodes"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetOrgFinancialByMonthOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrgFinancialByMonthArgs)(nil)).Elem()
}

// A collection of values returned by getOrgFinancialByMonth.
type GetOrgFinancialByMonthResultOutput struct{ *pulumi.OutputState }

func (GetOrgFinancialByMonthResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrgFinancialByMonthResult)(nil)).Elem()
}

func (o GetOrgFinancialByMonthResultOutput) ToGetOrgFinancialByMonthResultOutput() GetOrgFinancialByMonthResultOutput {
	return o
}

func (o GetOrgFinancialByMonthResultOutput) ToGetOrgFinancialByMonthResultOutputWithContext(ctx context.Context) GetOrgFinancialByMonthResultOutput {
	return o
}

func (o GetOrgFinancialByMonthResultOutput) EndMonth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOrgFinancialByMonthResult) *string { return v.EndMonth }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrgFinancialByMonthResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgFinancialByMonthResult) string { return v.Id }).(pulumi.StringOutput)
}

// Organization financial info by month.
func (o GetOrgFinancialByMonthResultOutput) Items() GetOrgFinancialByMonthItemArrayOutput {
	return o.ApplyT(func(v GetOrgFinancialByMonthResult) []GetOrgFinancialByMonthItem { return v.Items }).(GetOrgFinancialByMonthItemArrayOutput)
}

func (o GetOrgFinancialByMonthResultOutput) MemberUins() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetOrgFinancialByMonthResult) []int { return v.MemberUins }).(pulumi.IntArrayOutput)
}

func (o GetOrgFinancialByMonthResultOutput) ProductCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOrgFinancialByMonthResult) []string { return v.ProductCodes }).(pulumi.StringArrayOutput)
}

func (o GetOrgFinancialByMonthResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOrgFinancialByMonthResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrgFinancialByMonthResultOutput{})
}
