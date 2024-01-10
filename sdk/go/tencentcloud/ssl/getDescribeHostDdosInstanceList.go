// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of ssl describeHostDdosInstanceList
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
//			_, err := Ssl.GetDescribeHostDdosInstanceList(ctx, &ssl.GetDescribeHostDdosInstanceListArgs{
//				CertificateId: "8u8DII0l",
//				ResourceType:  "ddos",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDescribeHostDdosInstanceList(ctx *pulumi.Context, args *GetDescribeHostDdosInstanceListArgs, opts ...pulumi.InvokeOption) (*GetDescribeHostDdosInstanceListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDescribeHostDdosInstanceListResult
	err := ctx.Invoke("tencentcloud:Ssl/getDescribeHostDdosInstanceList:getDescribeHostDdosInstanceList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDescribeHostDdosInstanceList.
type GetDescribeHostDdosInstanceListArgs struct {
	// Certificate ID to be deployed.
	CertificateId string `pulumi:"certificateId"`
	// List of filtering parameters; Filterkey: domainmatch.
	Filters []GetDescribeHostDdosInstanceListFilter `pulumi:"filters"`
	// Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
	IsCache *int `pulumi:"isCache"`
	// Deployed certificate ID.
	OldCertificateId *string `pulumi:"oldCertificateId"`
	// Deploy resource type.
	ResourceType string `pulumi:"resourceType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDescribeHostDdosInstanceList.
type GetDescribeHostDdosInstanceListResult struct {
	CertificateId string                                  `pulumi:"certificateId"`
	Filters       []GetDescribeHostDdosInstanceListFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// DDOS example listNote: This field may return NULL, indicating that the valid value cannot be obtained.
	InstanceLists    []GetDescribeHostDdosInstanceListInstanceList `pulumi:"instanceLists"`
	IsCache          *int                                          `pulumi:"isCache"`
	OldCertificateId *string                                       `pulumi:"oldCertificateId"`
	ResourceType     string                                        `pulumi:"resourceType"`
	ResultOutputFile *string                                       `pulumi:"resultOutputFile"`
}

func GetDescribeHostDdosInstanceListOutput(ctx *pulumi.Context, args GetDescribeHostDdosInstanceListOutputArgs, opts ...pulumi.InvokeOption) GetDescribeHostDdosInstanceListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDescribeHostDdosInstanceListResult, error) {
			args := v.(GetDescribeHostDdosInstanceListArgs)
			r, err := GetDescribeHostDdosInstanceList(ctx, &args, opts...)
			var s GetDescribeHostDdosInstanceListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDescribeHostDdosInstanceListResultOutput)
}

// A collection of arguments for invoking getDescribeHostDdosInstanceList.
type GetDescribeHostDdosInstanceListOutputArgs struct {
	// Certificate ID to be deployed.
	CertificateId pulumi.StringInput `pulumi:"certificateId"`
	// List of filtering parameters; Filterkey: domainmatch.
	Filters GetDescribeHostDdosInstanceListFilterArrayInput `pulumi:"filters"`
	// Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
	IsCache pulumi.IntPtrInput `pulumi:"isCache"`
	// Deployed certificate ID.
	OldCertificateId pulumi.StringPtrInput `pulumi:"oldCertificateId"`
	// Deploy resource type.
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDescribeHostDdosInstanceListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeHostDdosInstanceListArgs)(nil)).Elem()
}

// A collection of values returned by getDescribeHostDdosInstanceList.
type GetDescribeHostDdosInstanceListResultOutput struct{ *pulumi.OutputState }

func (GetDescribeHostDdosInstanceListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeHostDdosInstanceListResult)(nil)).Elem()
}

func (o GetDescribeHostDdosInstanceListResultOutput) ToGetDescribeHostDdosInstanceListResultOutput() GetDescribeHostDdosInstanceListResultOutput {
	return o
}

func (o GetDescribeHostDdosInstanceListResultOutput) ToGetDescribeHostDdosInstanceListResultOutputWithContext(ctx context.Context) GetDescribeHostDdosInstanceListResultOutput {
	return o
}

func (o GetDescribeHostDdosInstanceListResultOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeHostDdosInstanceListResult) string { return v.CertificateId }).(pulumi.StringOutput)
}

func (o GetDescribeHostDdosInstanceListResultOutput) Filters() GetDescribeHostDdosInstanceListFilterArrayOutput {
	return o.ApplyT(func(v GetDescribeHostDdosInstanceListResult) []GetDescribeHostDdosInstanceListFilter {
		return v.Filters
	}).(GetDescribeHostDdosInstanceListFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDescribeHostDdosInstanceListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeHostDdosInstanceListResult) string { return v.Id }).(pulumi.StringOutput)
}

// DDOS example listNote: This field may return NULL, indicating that the valid value cannot be obtained.
func (o GetDescribeHostDdosInstanceListResultOutput) InstanceLists() GetDescribeHostDdosInstanceListInstanceListArrayOutput {
	return o.ApplyT(func(v GetDescribeHostDdosInstanceListResult) []GetDescribeHostDdosInstanceListInstanceList {
		return v.InstanceLists
	}).(GetDescribeHostDdosInstanceListInstanceListArrayOutput)
}

func (o GetDescribeHostDdosInstanceListResultOutput) IsCache() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDescribeHostDdosInstanceListResult) *int { return v.IsCache }).(pulumi.IntPtrOutput)
}

func (o GetDescribeHostDdosInstanceListResultOutput) OldCertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeHostDdosInstanceListResult) *string { return v.OldCertificateId }).(pulumi.StringPtrOutput)
}

func (o GetDescribeHostDdosInstanceListResultOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeHostDdosInstanceListResult) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o GetDescribeHostDdosInstanceListResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeHostDdosInstanceListResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDescribeHostDdosInstanceListResultOutput{})
}
