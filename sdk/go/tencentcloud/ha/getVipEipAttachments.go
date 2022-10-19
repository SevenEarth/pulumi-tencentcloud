// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of HA VIP EIP attachments
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Ha"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ha"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Ha.GetVipEipAttachments(ctx, &ha.GetVipEipAttachmentsArgs{
//				AddressIp: pulumi.StringRef("1.1.1.1"),
//				HavipId:   "havip-kjqwe4ba",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetVipEipAttachments(ctx *pulumi.Context, args *GetVipEipAttachmentsArgs, opts ...pulumi.InvokeOption) (*GetVipEipAttachmentsResult, error) {
	var rv GetVipEipAttachmentsResult
	err := ctx.Invoke("tencentcloud:Ha/getVipEipAttachments:getVipEipAttachments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVipEipAttachments.
type GetVipEipAttachmentsArgs struct {
	// Public IP address of EIP to be queried.
	AddressIp *string `pulumi:"addressIp"`
	// ID of the attached HA VIP to be queried.
	HavipId string `pulumi:"havipId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getVipEipAttachments.
type GetVipEipAttachmentsResult struct {
	// Public IP address of EIP.
	AddressIp *string `pulumi:"addressIp"`
	// A list of HA VIP EIP attachments. Each element contains the following attributes:
	HaVipEipAttachmentLists []GetVipEipAttachmentsHaVipEipAttachmentList `pulumi:"haVipEipAttachmentLists"`
	// ID of the attached HA VIP.
	HavipId string `pulumi:"havipId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetVipEipAttachmentsOutput(ctx *pulumi.Context, args GetVipEipAttachmentsOutputArgs, opts ...pulumi.InvokeOption) GetVipEipAttachmentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVipEipAttachmentsResult, error) {
			args := v.(GetVipEipAttachmentsArgs)
			r, err := GetVipEipAttachments(ctx, &args, opts...)
			var s GetVipEipAttachmentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVipEipAttachmentsResultOutput)
}

// A collection of arguments for invoking getVipEipAttachments.
type GetVipEipAttachmentsOutputArgs struct {
	// Public IP address of EIP to be queried.
	AddressIp pulumi.StringPtrInput `pulumi:"addressIp"`
	// ID of the attached HA VIP to be queried.
	HavipId pulumi.StringInput `pulumi:"havipId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetVipEipAttachmentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVipEipAttachmentsArgs)(nil)).Elem()
}

// A collection of values returned by getVipEipAttachments.
type GetVipEipAttachmentsResultOutput struct{ *pulumi.OutputState }

func (GetVipEipAttachmentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVipEipAttachmentsResult)(nil)).Elem()
}

func (o GetVipEipAttachmentsResultOutput) ToGetVipEipAttachmentsResultOutput() GetVipEipAttachmentsResultOutput {
	return o
}

func (o GetVipEipAttachmentsResultOutput) ToGetVipEipAttachmentsResultOutputWithContext(ctx context.Context) GetVipEipAttachmentsResultOutput {
	return o
}

// Public IP address of EIP.
func (o GetVipEipAttachmentsResultOutput) AddressIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVipEipAttachmentsResult) *string { return v.AddressIp }).(pulumi.StringPtrOutput)
}

// A list of HA VIP EIP attachments. Each element contains the following attributes:
func (o GetVipEipAttachmentsResultOutput) HaVipEipAttachmentLists() GetVipEipAttachmentsHaVipEipAttachmentListArrayOutput {
	return o.ApplyT(func(v GetVipEipAttachmentsResult) []GetVipEipAttachmentsHaVipEipAttachmentList {
		return v.HaVipEipAttachmentLists
	}).(GetVipEipAttachmentsHaVipEipAttachmentListArrayOutput)
}

// ID of the attached HA VIP.
func (o GetVipEipAttachmentsResultOutput) HavipId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVipEipAttachmentsResult) string { return v.HavipId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVipEipAttachmentsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVipEipAttachmentsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVipEipAttachmentsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVipEipAttachmentsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVipEipAttachmentsResultOutput{})
}
