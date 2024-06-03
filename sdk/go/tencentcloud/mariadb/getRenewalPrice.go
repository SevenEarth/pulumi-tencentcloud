// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of mariadb renewalPrice
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mariadb.GetRenewalPrice(ctx, &mariadb.GetRenewalPriceArgs{
//				InstanceId: "tdsql-9vqvls95",
//				Period:     pulumi.IntRef(2),
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
func GetRenewalPrice(ctx *pulumi.Context, args *GetRenewalPriceArgs, opts ...pulumi.InvokeOption) (*GetRenewalPriceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRenewalPriceResult
	err := ctx.Invoke("tencentcloud:Mariadb/getRenewalPrice:getRenewalPrice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRenewalPrice.
type GetRenewalPriceArgs struct {
	// Price unit. Valid values: `* pent` (cent), `* microPent` (microcent).
	AmountUnit *string `pulumi:"amountUnit"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Renewal duration, default: 1 month.
	Period *int `pulumi:"period"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getRenewalPrice.
type GetRenewalPriceResult struct {
	AmountUnit *string `pulumi:"amountUnit"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// Original price * Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description. * Currency: CNY (Chinese site), USD (international site).
	OriginalPrice int  `pulumi:"originalPrice"`
	Period        *int `pulumi:"period"`
	// The actual price may be different from the original price due to discounts. * Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description. * Currency: CNY (Chinese site), USD (international site).
	Price            int     `pulumi:"price"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetRenewalPriceOutput(ctx *pulumi.Context, args GetRenewalPriceOutputArgs, opts ...pulumi.InvokeOption) GetRenewalPriceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRenewalPriceResult, error) {
			args := v.(GetRenewalPriceArgs)
			r, err := GetRenewalPrice(ctx, &args, opts...)
			var s GetRenewalPriceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRenewalPriceResultOutput)
}

// A collection of arguments for invoking getRenewalPrice.
type GetRenewalPriceOutputArgs struct {
	// Price unit. Valid values: `* pent` (cent), `* microPent` (microcent).
	AmountUnit pulumi.StringPtrInput `pulumi:"amountUnit"`
	// Instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Renewal duration, default: 1 month.
	Period pulumi.IntPtrInput `pulumi:"period"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetRenewalPriceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRenewalPriceArgs)(nil)).Elem()
}

// A collection of values returned by getRenewalPrice.
type GetRenewalPriceResultOutput struct{ *pulumi.OutputState }

func (GetRenewalPriceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRenewalPriceResult)(nil)).Elem()
}

func (o GetRenewalPriceResultOutput) ToGetRenewalPriceResultOutput() GetRenewalPriceResultOutput {
	return o
}

func (o GetRenewalPriceResultOutput) ToGetRenewalPriceResultOutputWithContext(ctx context.Context) GetRenewalPriceResultOutput {
	return o
}

func (o GetRenewalPriceResultOutput) AmountUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRenewalPriceResult) *string { return v.AmountUnit }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRenewalPriceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRenewalPriceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRenewalPriceResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRenewalPriceResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Original price * Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description. * Currency: CNY (Chinese site), USD (international site).
func (o GetRenewalPriceResultOutput) OriginalPrice() pulumi.IntOutput {
	return o.ApplyT(func(v GetRenewalPriceResult) int { return v.OriginalPrice }).(pulumi.IntOutput)
}

func (o GetRenewalPriceResultOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetRenewalPriceResult) *int { return v.Period }).(pulumi.IntPtrOutput)
}

// The actual price may be different from the original price due to discounts. * Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description. * Currency: CNY (Chinese site), USD (international site).
func (o GetRenewalPriceResultOutput) Price() pulumi.IntOutput {
	return o.ApplyT(func(v GetRenewalPriceResult) int { return v.Price }).(pulumi.IntOutput)
}

func (o GetRenewalPriceResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRenewalPriceResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRenewalPriceResultOutput{})
}
