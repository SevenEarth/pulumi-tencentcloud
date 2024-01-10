// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of ssl describeHostVodInstanceList
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Ssl"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ssl"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Ssl.GetDescribeHostVodInstanceList(ctx, &ssl.GetDescribeHostVodInstanceListArgs{
//				CertificateId: "8u8DII0l",
//				ResourceType:  "vod",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDescribeHostVodInstanceList(ctx *pulumi.Context, args *GetDescribeHostVodInstanceListArgs, opts ...pulumi.InvokeOption) (*GetDescribeHostVodInstanceListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDescribeHostVodInstanceListResult
	err := ctx.Invoke("tencentcloud:Ssl/getDescribeHostVodInstanceList:getDescribeHostVodInstanceList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDescribeHostVodInstanceList.
type GetDescribeHostVodInstanceListArgs struct {
	// Certificate ID to be deployed.
	CertificateId string `pulumi:"certificateId"`
	// List of filter parameters.
	Filters []GetDescribeHostVodInstanceListFilter `pulumi:"filters"`
	// Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
	IsCache *int `pulumi:"isCache"`
	// Deployed certificate ID.
	OldCertificateId *string `pulumi:"oldCertificateId"`
	// Deploy resource type VOD.
	ResourceType string `pulumi:"resourceType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDescribeHostVodInstanceList.
type GetDescribeHostVodInstanceListResult struct {
	CertificateId string                                 `pulumi:"certificateId"`
	Filters       []GetDescribeHostVodInstanceListFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// VOD example listNote: This field may return NULL, indicating that the valid value cannot be obtained.
	InstanceLists    []GetDescribeHostVodInstanceListInstanceList `pulumi:"instanceLists"`
	IsCache          *int                                         `pulumi:"isCache"`
	OldCertificateId *string                                      `pulumi:"oldCertificateId"`
	ResourceType     string                                       `pulumi:"resourceType"`
	ResultOutputFile *string                                      `pulumi:"resultOutputFile"`
}

func GetDescribeHostVodInstanceListOutput(ctx *pulumi.Context, args GetDescribeHostVodInstanceListOutputArgs, opts ...pulumi.InvokeOption) GetDescribeHostVodInstanceListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDescribeHostVodInstanceListResult, error) {
			args := v.(GetDescribeHostVodInstanceListArgs)
			r, err := GetDescribeHostVodInstanceList(ctx, &args, opts...)
			var s GetDescribeHostVodInstanceListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDescribeHostVodInstanceListResultOutput)
}

// A collection of arguments for invoking getDescribeHostVodInstanceList.
type GetDescribeHostVodInstanceListOutputArgs struct {
	// Certificate ID to be deployed.
	CertificateId pulumi.StringInput `pulumi:"certificateId"`
	// List of filter parameters.
	Filters GetDescribeHostVodInstanceListFilterArrayInput `pulumi:"filters"`
	// Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
	IsCache pulumi.IntPtrInput `pulumi:"isCache"`
	// Deployed certificate ID.
	OldCertificateId pulumi.StringPtrInput `pulumi:"oldCertificateId"`
	// Deploy resource type VOD.
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDescribeHostVodInstanceListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeHostVodInstanceListArgs)(nil)).Elem()
}

// A collection of values returned by getDescribeHostVodInstanceList.
type GetDescribeHostVodInstanceListResultOutput struct{ *pulumi.OutputState }

func (GetDescribeHostVodInstanceListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeHostVodInstanceListResult)(nil)).Elem()
}

func (o GetDescribeHostVodInstanceListResultOutput) ToGetDescribeHostVodInstanceListResultOutput() GetDescribeHostVodInstanceListResultOutput {
	return o
}

func (o GetDescribeHostVodInstanceListResultOutput) ToGetDescribeHostVodInstanceListResultOutputWithContext(ctx context.Context) GetDescribeHostVodInstanceListResultOutput {
	return o
}

func (o GetDescribeHostVodInstanceListResultOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeHostVodInstanceListResult) string { return v.CertificateId }).(pulumi.StringOutput)
}

func (o GetDescribeHostVodInstanceListResultOutput) Filters() GetDescribeHostVodInstanceListFilterArrayOutput {
	return o.ApplyT(func(v GetDescribeHostVodInstanceListResult) []GetDescribeHostVodInstanceListFilter { return v.Filters }).(GetDescribeHostVodInstanceListFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDescribeHostVodInstanceListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeHostVodInstanceListResult) string { return v.Id }).(pulumi.StringOutput)
}

// VOD example listNote: This field may return NULL, indicating that the valid value cannot be obtained.
func (o GetDescribeHostVodInstanceListResultOutput) InstanceLists() GetDescribeHostVodInstanceListInstanceListArrayOutput {
	return o.ApplyT(func(v GetDescribeHostVodInstanceListResult) []GetDescribeHostVodInstanceListInstanceList {
		return v.InstanceLists
	}).(GetDescribeHostVodInstanceListInstanceListArrayOutput)
}

func (o GetDescribeHostVodInstanceListResultOutput) IsCache() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDescribeHostVodInstanceListResult) *int { return v.IsCache }).(pulumi.IntPtrOutput)
}

func (o GetDescribeHostVodInstanceListResultOutput) OldCertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeHostVodInstanceListResult) *string { return v.OldCertificateId }).(pulumi.StringPtrOutput)
}

func (o GetDescribeHostVodInstanceListResultOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeHostVodInstanceListResult) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o GetDescribeHostVodInstanceListResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeHostVodInstanceListResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDescribeHostVodInstanceListResultOutput{})
}
