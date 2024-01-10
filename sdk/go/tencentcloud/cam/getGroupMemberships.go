// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of CAM group memberships
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
//			_, err := Cam.GetGroupMemberships(ctx, &cam.GetGroupMembershipsArgs{
//				GroupId: pulumi.StringRef(tencentcloud_cam_group.Foo.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetGroupMemberships(ctx *pulumi.Context, args *GetGroupMembershipsArgs, opts ...pulumi.InvokeOption) (*GetGroupMembershipsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetGroupMembershipsResult
	err := ctx.Invoke("tencentcloud:Cam/getGroupMemberships:getGroupMemberships", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroupMemberships.
type GetGroupMembershipsArgs struct {
	// ID of CAM group to be queried.
	GroupId *string `pulumi:"groupId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getGroupMemberships.
type GetGroupMembershipsResult struct {
	// ID of CAM group.
	GroupId *string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of CAM group membership. Each element contains the following attributes:
	MembershipLists  []GetGroupMembershipsMembershipList `pulumi:"membershipLists"`
	ResultOutputFile *string                             `pulumi:"resultOutputFile"`
}

func GetGroupMembershipsOutput(ctx *pulumi.Context, args GetGroupMembershipsOutputArgs, opts ...pulumi.InvokeOption) GetGroupMembershipsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupMembershipsResult, error) {
			args := v.(GetGroupMembershipsArgs)
			r, err := GetGroupMemberships(ctx, &args, opts...)
			var s GetGroupMembershipsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGroupMembershipsResultOutput)
}

// A collection of arguments for invoking getGroupMemberships.
type GetGroupMembershipsOutputArgs struct {
	// ID of CAM group to be queried.
	GroupId pulumi.StringPtrInput `pulumi:"groupId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetGroupMembershipsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupMembershipsArgs)(nil)).Elem()
}

// A collection of values returned by getGroupMemberships.
type GetGroupMembershipsResultOutput struct{ *pulumi.OutputState }

func (GetGroupMembershipsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupMembershipsResult)(nil)).Elem()
}

func (o GetGroupMembershipsResultOutput) ToGetGroupMembershipsResultOutput() GetGroupMembershipsResultOutput {
	return o
}

func (o GetGroupMembershipsResultOutput) ToGetGroupMembershipsResultOutputWithContext(ctx context.Context) GetGroupMembershipsResultOutput {
	return o
}

// ID of CAM group.
func (o GetGroupMembershipsResultOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupMembershipsResult) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGroupMembershipsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupMembershipsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of CAM group membership. Each element contains the following attributes:
func (o GetGroupMembershipsResultOutput) MembershipLists() GetGroupMembershipsMembershipListArrayOutput {
	return o.ApplyT(func(v GetGroupMembershipsResult) []GetGroupMembershipsMembershipList { return v.MembershipLists }).(GetGroupMembershipsMembershipListArrayOutput)
}

func (o GetGroupMembershipsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupMembershipsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupMembershipsResultOutput{})
}
