// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of dcdb saleInfo
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dcdb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dcdb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dcdb.GetSaleInfo(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetSaleInfo(ctx *pulumi.Context, args *GetSaleInfoArgs, opts ...pulumi.InvokeOption) (*GetSaleInfoResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSaleInfoResult
	err := ctx.Invoke("tencentcloud:Dcdb/getSaleInfo:getSaleInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSaleInfo.
type GetSaleInfoArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getSaleInfo.
type GetSaleInfoResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// list of sale region info.
	RegionLists      []GetSaleInfoRegionList `pulumi:"regionLists"`
	ResultOutputFile *string                 `pulumi:"resultOutputFile"`
}

func GetSaleInfoOutput(ctx *pulumi.Context, args GetSaleInfoOutputArgs, opts ...pulumi.InvokeOption) GetSaleInfoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSaleInfoResult, error) {
			args := v.(GetSaleInfoArgs)
			r, err := GetSaleInfo(ctx, &args, opts...)
			var s GetSaleInfoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSaleInfoResultOutput)
}

// A collection of arguments for invoking getSaleInfo.
type GetSaleInfoOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetSaleInfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSaleInfoArgs)(nil)).Elem()
}

// A collection of values returned by getSaleInfo.
type GetSaleInfoResultOutput struct{ *pulumi.OutputState }

func (GetSaleInfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSaleInfoResult)(nil)).Elem()
}

func (o GetSaleInfoResultOutput) ToGetSaleInfoResultOutput() GetSaleInfoResultOutput {
	return o
}

func (o GetSaleInfoResultOutput) ToGetSaleInfoResultOutputWithContext(ctx context.Context) GetSaleInfoResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSaleInfoResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSaleInfoResult) string { return v.Id }).(pulumi.StringOutput)
}

// list of sale region info.
func (o GetSaleInfoResultOutput) RegionLists() GetSaleInfoRegionListArrayOutput {
	return o.ApplyT(func(v GetSaleInfoResult) []GetSaleInfoRegionList { return v.RegionLists }).(GetSaleInfoRegionListArrayOutput)
}

func (o GetSaleInfoResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSaleInfoResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSaleInfoResultOutput{})
}
