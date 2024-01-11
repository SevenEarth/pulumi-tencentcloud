// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of ssl describeHostClbInstanceList
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Ssl"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ssl"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Ssl.GetDescribeHostClbInstanceList(ctx, &ssl.GetDescribeHostClbInstanceListArgs{
// 			CertificateId: "8u8DII0l",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDescribeHostClbInstanceList(ctx *pulumi.Context, args *GetDescribeHostClbInstanceListArgs, opts ...pulumi.InvokeOption) (*GetDescribeHostClbInstanceListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDescribeHostClbInstanceListResult
	err := ctx.Invoke("tencentcloud:Ssl/getDescribeHostClbInstanceList:getDescribeHostClbInstanceList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDescribeHostClbInstanceList.
type GetDescribeHostClbInstanceListArgs struct {
	// Whether to cache asynchronous.
	AsyncCache *int `pulumi:"asyncCache"`
	// Certificate ID to be deployed.
	CertificateId string `pulumi:"certificateId"`
	// List of filtering parameters; Filterkey: domainmatch.
	Filters []GetDescribeHostClbInstanceListFilter `pulumi:"filters"`
	// Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
	IsCache *int `pulumi:"isCache"`
	// Original certificate ID.
	OldCertificateId *string `pulumi:"oldCertificateId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDescribeHostClbInstanceList.
type GetDescribeHostClbInstanceListResult struct {
	AsyncCache *int `pulumi:"asyncCache"`
	// Current cache read timeNote: This field may return NULL, indicating that the valid value cannot be obtained.
	AsyncCacheTime string `pulumi:"asyncCacheTime"`
	// Asynchronous refresh current execution numberNote: This field may return NULL, indicating that the valid value cannot be obtained.
	AsyncOffset int `pulumi:"asyncOffset"`
	// The total number of asynchronous refreshNote: This field may return NULL, indicating that the valid value cannot be obtained.
	AsyncTotalNum int                                    `pulumi:"asyncTotalNum"`
	CertificateId string                                 `pulumi:"certificateId"`
	Filters       []GetDescribeHostClbInstanceListFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// CLB instance listener listNote: This field may return NULL, indicating that the valid value cannot be obtained.
	InstanceLists    []GetDescribeHostClbInstanceListInstanceList `pulumi:"instanceLists"`
	IsCache          *int                                         `pulumi:"isCache"`
	OldCertificateId *string                                      `pulumi:"oldCertificateId"`
	ResultOutputFile *string                                      `pulumi:"resultOutputFile"`
}

func GetDescribeHostClbInstanceListOutput(ctx *pulumi.Context, args GetDescribeHostClbInstanceListOutputArgs, opts ...pulumi.InvokeOption) GetDescribeHostClbInstanceListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDescribeHostClbInstanceListResult, error) {
			args := v.(GetDescribeHostClbInstanceListArgs)
			r, err := GetDescribeHostClbInstanceList(ctx, &args, opts...)
			var s GetDescribeHostClbInstanceListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDescribeHostClbInstanceListResultOutput)
}

// A collection of arguments for invoking getDescribeHostClbInstanceList.
type GetDescribeHostClbInstanceListOutputArgs struct {
	// Whether to cache asynchronous.
	AsyncCache pulumi.IntPtrInput `pulumi:"asyncCache"`
	// Certificate ID to be deployed.
	CertificateId pulumi.StringInput `pulumi:"certificateId"`
	// List of filtering parameters; Filterkey: domainmatch.
	Filters GetDescribeHostClbInstanceListFilterArrayInput `pulumi:"filters"`
	// Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
	IsCache pulumi.IntPtrInput `pulumi:"isCache"`
	// Original certificate ID.
	OldCertificateId pulumi.StringPtrInput `pulumi:"oldCertificateId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDescribeHostClbInstanceListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeHostClbInstanceListArgs)(nil)).Elem()
}

// A collection of values returned by getDescribeHostClbInstanceList.
type GetDescribeHostClbInstanceListResultOutput struct{ *pulumi.OutputState }

func (GetDescribeHostClbInstanceListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeHostClbInstanceListResult)(nil)).Elem()
}

func (o GetDescribeHostClbInstanceListResultOutput) ToGetDescribeHostClbInstanceListResultOutput() GetDescribeHostClbInstanceListResultOutput {
	return o
}

func (o GetDescribeHostClbInstanceListResultOutput) ToGetDescribeHostClbInstanceListResultOutputWithContext(ctx context.Context) GetDescribeHostClbInstanceListResultOutput {
	return o
}

func (o GetDescribeHostClbInstanceListResultOutput) AsyncCache() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDescribeHostClbInstanceListResult) *int { return v.AsyncCache }).(pulumi.IntPtrOutput)
}

// Current cache read timeNote: This field may return NULL, indicating that the valid value cannot be obtained.
func (o GetDescribeHostClbInstanceListResultOutput) AsyncCacheTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeHostClbInstanceListResult) string { return v.AsyncCacheTime }).(pulumi.StringOutput)
}

// Asynchronous refresh current execution numberNote: This field may return NULL, indicating that the valid value cannot be obtained.
func (o GetDescribeHostClbInstanceListResultOutput) AsyncOffset() pulumi.IntOutput {
	return o.ApplyT(func(v GetDescribeHostClbInstanceListResult) int { return v.AsyncOffset }).(pulumi.IntOutput)
}

// The total number of asynchronous refreshNote: This field may return NULL, indicating that the valid value cannot be obtained.
func (o GetDescribeHostClbInstanceListResultOutput) AsyncTotalNum() pulumi.IntOutput {
	return o.ApplyT(func(v GetDescribeHostClbInstanceListResult) int { return v.AsyncTotalNum }).(pulumi.IntOutput)
}

func (o GetDescribeHostClbInstanceListResultOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeHostClbInstanceListResult) string { return v.CertificateId }).(pulumi.StringOutput)
}

func (o GetDescribeHostClbInstanceListResultOutput) Filters() GetDescribeHostClbInstanceListFilterArrayOutput {
	return o.ApplyT(func(v GetDescribeHostClbInstanceListResult) []GetDescribeHostClbInstanceListFilter { return v.Filters }).(GetDescribeHostClbInstanceListFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDescribeHostClbInstanceListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeHostClbInstanceListResult) string { return v.Id }).(pulumi.StringOutput)
}

// CLB instance listener listNote: This field may return NULL, indicating that the valid value cannot be obtained.
func (o GetDescribeHostClbInstanceListResultOutput) InstanceLists() GetDescribeHostClbInstanceListInstanceListArrayOutput {
	return o.ApplyT(func(v GetDescribeHostClbInstanceListResult) []GetDescribeHostClbInstanceListInstanceList {
		return v.InstanceLists
	}).(GetDescribeHostClbInstanceListInstanceListArrayOutput)
}

func (o GetDescribeHostClbInstanceListResultOutput) IsCache() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDescribeHostClbInstanceListResult) *int { return v.IsCache }).(pulumi.IntPtrOutput)
}

func (o GetDescribeHostClbInstanceListResultOutput) OldCertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeHostClbInstanceListResult) *string { return v.OldCertificateId }).(pulumi.StringPtrOutput)
}

func (o GetDescribeHostClbInstanceListResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeHostClbInstanceListResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDescribeHostClbInstanceListResultOutput{})
}