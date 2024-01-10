// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of organization orgFinancialByMember
func GetOrgFinancialByMember(ctx *pulumi.Context, args *GetOrgFinancialByMemberArgs, opts ...pulumi.InvokeOption) (*GetOrgFinancialByMemberResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetOrgFinancialByMemberResult
	err := ctx.Invoke("tencentcloud:Organization/getOrgFinancialByMember:getOrgFinancialByMember", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrgFinancialByMember.
type GetOrgFinancialByMemberArgs struct {
	// Query for the end month. Format:yyyy-mm, for example:2021-01.The default value is the `Month`.
	EndMonth *string `pulumi:"endMonth"`
	// Member uin list. Up to 100.
	MemberUins []int `pulumi:"memberUins"`
	// Query for the start month. Format:yyyy-mm, for example:2021-01.
	Month string `pulumi:"month"`
	// Product code list. Up to 100.
	ProductCodes []string `pulumi:"productCodes"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getOrgFinancialByMember.
type GetOrgFinancialByMemberResult struct {
	EndMonth *string `pulumi:"endMonth"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Member financial detail.
	Items            []GetOrgFinancialByMemberItem `pulumi:"items"`
	MemberUins       []int                         `pulumi:"memberUins"`
	Month            string                        `pulumi:"month"`
	ProductCodes     []string                      `pulumi:"productCodes"`
	ResultOutputFile *string                       `pulumi:"resultOutputFile"`
	// Total cost of the member.
	TotalCost float64 `pulumi:"totalCost"`
}

func GetOrgFinancialByMemberOutput(ctx *pulumi.Context, args GetOrgFinancialByMemberOutputArgs, opts ...pulumi.InvokeOption) GetOrgFinancialByMemberResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOrgFinancialByMemberResult, error) {
			args := v.(GetOrgFinancialByMemberArgs)
			r, err := GetOrgFinancialByMember(ctx, &args, opts...)
			var s GetOrgFinancialByMemberResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOrgFinancialByMemberResultOutput)
}

// A collection of arguments for invoking getOrgFinancialByMember.
type GetOrgFinancialByMemberOutputArgs struct {
	// Query for the end month. Format:yyyy-mm, for example:2021-01.The default value is the `Month`.
	EndMonth pulumi.StringPtrInput `pulumi:"endMonth"`
	// Member uin list. Up to 100.
	MemberUins pulumi.IntArrayInput `pulumi:"memberUins"`
	// Query for the start month. Format:yyyy-mm, for example:2021-01.
	Month pulumi.StringInput `pulumi:"month"`
	// Product code list. Up to 100.
	ProductCodes pulumi.StringArrayInput `pulumi:"productCodes"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetOrgFinancialByMemberOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrgFinancialByMemberArgs)(nil)).Elem()
}

// A collection of values returned by getOrgFinancialByMember.
type GetOrgFinancialByMemberResultOutput struct{ *pulumi.OutputState }

func (GetOrgFinancialByMemberResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrgFinancialByMemberResult)(nil)).Elem()
}

func (o GetOrgFinancialByMemberResultOutput) ToGetOrgFinancialByMemberResultOutput() GetOrgFinancialByMemberResultOutput {
	return o
}

func (o GetOrgFinancialByMemberResultOutput) ToGetOrgFinancialByMemberResultOutputWithContext(ctx context.Context) GetOrgFinancialByMemberResultOutput {
	return o
}

func (o GetOrgFinancialByMemberResultOutput) EndMonth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOrgFinancialByMemberResult) *string { return v.EndMonth }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrgFinancialByMemberResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgFinancialByMemberResult) string { return v.Id }).(pulumi.StringOutput)
}

// Member financial detail.
func (o GetOrgFinancialByMemberResultOutput) Items() GetOrgFinancialByMemberItemArrayOutput {
	return o.ApplyT(func(v GetOrgFinancialByMemberResult) []GetOrgFinancialByMemberItem { return v.Items }).(GetOrgFinancialByMemberItemArrayOutput)
}

func (o GetOrgFinancialByMemberResultOutput) MemberUins() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetOrgFinancialByMemberResult) []int { return v.MemberUins }).(pulumi.IntArrayOutput)
}

func (o GetOrgFinancialByMemberResultOutput) Month() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgFinancialByMemberResult) string { return v.Month }).(pulumi.StringOutput)
}

func (o GetOrgFinancialByMemberResultOutput) ProductCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOrgFinancialByMemberResult) []string { return v.ProductCodes }).(pulumi.StringArrayOutput)
}

func (o GetOrgFinancialByMemberResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOrgFinancialByMemberResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Total cost of the member.
func (o GetOrgFinancialByMemberResultOutput) TotalCost() pulumi.Float64Output {
	return o.ApplyT(func(v GetOrgFinancialByMemberResult) float64 { return v.TotalCost }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(GetOrgFinancialByMemberResultOutput{})
}
