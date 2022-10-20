// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of CAM role policy attachments
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cam"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cam.GetRolePolicyAttachments(ctx, &cam.GetRolePolicyAttachmentsArgs{
//				RoleId: tencentcloud_cam_role.Foo.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Cam.GetRolePolicyAttachments(ctx, &cam.GetRolePolicyAttachmentsArgs{
//				RoleId:   tencentcloud_cam_role.Foo.Id,
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
func GetRolePolicyAttachments(ctx *pulumi.Context, args *GetRolePolicyAttachmentsArgs, opts ...pulumi.InvokeOption) (*GetRolePolicyAttachmentsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRolePolicyAttachmentsResult
	err := ctx.Invoke("tencentcloud:Cam/getRolePolicyAttachments:getRolePolicyAttachments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRolePolicyAttachments.
type GetRolePolicyAttachmentsArgs struct {
	// Mode of Creation of the CAM user policy attachment. `1` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
	CreateMode *int `pulumi:"createMode"`
	// ID of CAM policy to be queried.
	PolicyId *string `pulumi:"policyId"`
	// Type of the policy strategy. Valid values are 'User', 'QCS'. 'User' means customer strategy and 'QCS' means preset strategy.
	PolicyType *string `pulumi:"policyType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of the attached CAM role to be queried.
	RoleId string `pulumi:"roleId"`
}

// A collection of values returned by getRolePolicyAttachments.
type GetRolePolicyAttachmentsResult struct {
	// Mode of Creation of the CAM role policy attachment. `1` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
	CreateMode *int `pulumi:"createMode"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of CAM role.
	PolicyId *string `pulumi:"policyId"`
	// Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
	PolicyType       *string `pulumi:"policyType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of CAM role.
	RoleId string `pulumi:"roleId"`
	// A list of CAM role policy attachments. Each element contains the following attributes:
	RolePolicyAttachmentLists []GetRolePolicyAttachmentsRolePolicyAttachmentList `pulumi:"rolePolicyAttachmentLists"`
}

func GetRolePolicyAttachmentsOutput(ctx *pulumi.Context, args GetRolePolicyAttachmentsOutputArgs, opts ...pulumi.InvokeOption) GetRolePolicyAttachmentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRolePolicyAttachmentsResult, error) {
			args := v.(GetRolePolicyAttachmentsArgs)
			r, err := GetRolePolicyAttachments(ctx, &args, opts...)
			var s GetRolePolicyAttachmentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRolePolicyAttachmentsResultOutput)
}

// A collection of arguments for invoking getRolePolicyAttachments.
type GetRolePolicyAttachmentsOutputArgs struct {
	// Mode of Creation of the CAM user policy attachment. `1` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
	CreateMode pulumi.IntPtrInput `pulumi:"createMode"`
	// ID of CAM policy to be queried.
	PolicyId pulumi.StringPtrInput `pulumi:"policyId"`
	// Type of the policy strategy. Valid values are 'User', 'QCS'. 'User' means customer strategy and 'QCS' means preset strategy.
	PolicyType pulumi.StringPtrInput `pulumi:"policyType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// ID of the attached CAM role to be queried.
	RoleId pulumi.StringInput `pulumi:"roleId"`
}

func (GetRolePolicyAttachmentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRolePolicyAttachmentsArgs)(nil)).Elem()
}

// A collection of values returned by getRolePolicyAttachments.
type GetRolePolicyAttachmentsResultOutput struct{ *pulumi.OutputState }

func (GetRolePolicyAttachmentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRolePolicyAttachmentsResult)(nil)).Elem()
}

func (o GetRolePolicyAttachmentsResultOutput) ToGetRolePolicyAttachmentsResultOutput() GetRolePolicyAttachmentsResultOutput {
	return o
}

func (o GetRolePolicyAttachmentsResultOutput) ToGetRolePolicyAttachmentsResultOutputWithContext(ctx context.Context) GetRolePolicyAttachmentsResultOutput {
	return o
}

// Mode of Creation of the CAM role policy attachment. `1` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
func (o GetRolePolicyAttachmentsResultOutput) CreateMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetRolePolicyAttachmentsResult) *int { return v.CreateMode }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRolePolicyAttachmentsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRolePolicyAttachmentsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of CAM role.
func (o GetRolePolicyAttachmentsResultOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolePolicyAttachmentsResult) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

// Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
func (o GetRolePolicyAttachmentsResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolePolicyAttachmentsResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

func (o GetRolePolicyAttachmentsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolePolicyAttachmentsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// ID of CAM role.
func (o GetRolePolicyAttachmentsResultOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRolePolicyAttachmentsResult) string { return v.RoleId }).(pulumi.StringOutput)
}

// A list of CAM role policy attachments. Each element contains the following attributes:
func (o GetRolePolicyAttachmentsResultOutput) RolePolicyAttachmentLists() GetRolePolicyAttachmentsRolePolicyAttachmentListArrayOutput {
	return o.ApplyT(func(v GetRolePolicyAttachmentsResult) []GetRolePolicyAttachmentsRolePolicyAttachmentList {
		return v.RolePolicyAttachmentLists
	}).(GetRolePolicyAttachmentsRolePolicyAttachmentListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRolePolicyAttachmentsResultOutput{})
}
