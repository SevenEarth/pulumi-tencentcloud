// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of CAM user policy attachments
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cam"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cam.GetUserPolicyAttachments(ctx, &cam.GetUserPolicyAttachmentsArgs{
//				UserId: pulumi.StringRef(tencentcloud_cam_user.Foo.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Cam.GetUserPolicyAttachments(ctx, &cam.GetUserPolicyAttachmentsArgs{
//				UserId:   pulumi.StringRef(tencentcloud_cam_user.Foo.Id),
//				PolicyId: pulumi.StringRef(tencentcloud_cam_policy.Foo.Id),
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
func GetUserPolicyAttachments(ctx *pulumi.Context, args *GetUserPolicyAttachmentsArgs, opts ...pulumi.InvokeOption) (*GetUserPolicyAttachmentsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUserPolicyAttachmentsResult
	err := ctx.Invoke("tencentcloud:Cam/getUserPolicyAttachments:getUserPolicyAttachments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserPolicyAttachments.
type GetUserPolicyAttachmentsArgs struct {
	// Mode of Creation of the CAM user policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
	CreateMode *int `pulumi:"createMode"`
	// ID of CAM policy to be queried.
	PolicyId *string `pulumi:"policyId"`
	// Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
	PolicyType *string `pulumi:"policyType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// It has been deprecated from version 1.59.6. Use `userName` instead. ID of the attached CAM user to be queried.
	//
	// Deprecated: It has been deprecated from version 1.59.6. Use `userName` instead.
	UserId *string `pulumi:"userId"`
	// Name of the attached CAM user as unique key to be queried.
	UserName *string `pulumi:"userName"`
}

// A collection of values returned by getUserPolicyAttachments.
type GetUserPolicyAttachmentsResult struct {
	// Mode of Creation of the CAM user policy attachment. `1` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
	CreateMode *int `pulumi:"createMode"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of CAM user.
	PolicyId *string `pulumi:"policyId"`
	// Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
	PolicyType       *string `pulumi:"policyType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// (**Deprecated**) It has been deprecated from version 1.59.6. Use `userName` instead. ID of CAM user.
	//
	// Deprecated: It has been deprecated from version 1.59.6. Use `userName` instead.
	UserId *string `pulumi:"userId"`
	// Name of CAM user as unique key.
	UserName *string `pulumi:"userName"`
	// A list of CAM user policy attachments. Each element contains the following attributes:
	UserPolicyAttachmentLists []GetUserPolicyAttachmentsUserPolicyAttachmentList `pulumi:"userPolicyAttachmentLists"`
}

func GetUserPolicyAttachmentsOutput(ctx *pulumi.Context, args GetUserPolicyAttachmentsOutputArgs, opts ...pulumi.InvokeOption) GetUserPolicyAttachmentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserPolicyAttachmentsResult, error) {
			args := v.(GetUserPolicyAttachmentsArgs)
			r, err := GetUserPolicyAttachments(ctx, &args, opts...)
			var s GetUserPolicyAttachmentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUserPolicyAttachmentsResultOutput)
}

// A collection of arguments for invoking getUserPolicyAttachments.
type GetUserPolicyAttachmentsOutputArgs struct {
	// Mode of Creation of the CAM user policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
	CreateMode pulumi.IntPtrInput `pulumi:"createMode"`
	// ID of CAM policy to be queried.
	PolicyId pulumi.StringPtrInput `pulumi:"policyId"`
	// Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
	PolicyType pulumi.StringPtrInput `pulumi:"policyType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// It has been deprecated from version 1.59.6. Use `userName` instead. ID of the attached CAM user to be queried.
	//
	// Deprecated: It has been deprecated from version 1.59.6. Use `userName` instead.
	UserId pulumi.StringPtrInput `pulumi:"userId"`
	// Name of the attached CAM user as unique key to be queried.
	UserName pulumi.StringPtrInput `pulumi:"userName"`
}

func (GetUserPolicyAttachmentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserPolicyAttachmentsArgs)(nil)).Elem()
}

// A collection of values returned by getUserPolicyAttachments.
type GetUserPolicyAttachmentsResultOutput struct{ *pulumi.OutputState }

func (GetUserPolicyAttachmentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserPolicyAttachmentsResult)(nil)).Elem()
}

func (o GetUserPolicyAttachmentsResultOutput) ToGetUserPolicyAttachmentsResultOutput() GetUserPolicyAttachmentsResultOutput {
	return o
}

func (o GetUserPolicyAttachmentsResultOutput) ToGetUserPolicyAttachmentsResultOutputWithContext(ctx context.Context) GetUserPolicyAttachmentsResultOutput {
	return o
}

// Mode of Creation of the CAM user policy attachment. `1` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
func (o GetUserPolicyAttachmentsResultOutput) CreateMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetUserPolicyAttachmentsResult) *int { return v.CreateMode }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetUserPolicyAttachmentsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserPolicyAttachmentsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of CAM user.
func (o GetUserPolicyAttachmentsResultOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserPolicyAttachmentsResult) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

// Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
func (o GetUserPolicyAttachmentsResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserPolicyAttachmentsResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

func (o GetUserPolicyAttachmentsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserPolicyAttachmentsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// (**Deprecated**) It has been deprecated from version 1.59.6. Use `userName` instead. ID of CAM user.
//
// Deprecated: It has been deprecated from version 1.59.6. Use `userName` instead.
func (o GetUserPolicyAttachmentsResultOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserPolicyAttachmentsResult) *string { return v.UserId }).(pulumi.StringPtrOutput)
}

// Name of CAM user as unique key.
func (o GetUserPolicyAttachmentsResultOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserPolicyAttachmentsResult) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

// A list of CAM user policy attachments. Each element contains the following attributes:
func (o GetUserPolicyAttachmentsResultOutput) UserPolicyAttachmentLists() GetUserPolicyAttachmentsUserPolicyAttachmentListArrayOutput {
	return o.ApplyT(func(v GetUserPolicyAttachmentsResult) []GetUserPolicyAttachmentsUserPolicyAttachmentList {
		return v.UserPolicyAttachmentLists
	}).(GetUserPolicyAttachmentsUserPolicyAttachmentListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserPolicyAttachmentsResultOutput{})
}
