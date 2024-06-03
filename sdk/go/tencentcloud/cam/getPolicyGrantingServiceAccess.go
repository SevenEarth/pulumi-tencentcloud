// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of cam policyGrantingServiceAccess
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
//			_, err := Cam.GetPolicyGrantingServiceAccess(ctx, &cam.GetPolicyGrantingServiceAccessArgs{
//				RoleId:      pulumi.IntRef(4611686018436804608),
//				ServiceType: pulumi.StringRef("cam"),
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
func GetPolicyGrantingServiceAccess(ctx *pulumi.Context, args *GetPolicyGrantingServiceAccessArgs, opts ...pulumi.InvokeOption) (*GetPolicyGrantingServiceAccessResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPolicyGrantingServiceAccessResult
	err := ctx.Invoke("tencentcloud:Cam/getPolicyGrantingServiceAccess:getPolicyGrantingServiceAccess", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicyGrantingServiceAccess.
type GetPolicyGrantingServiceAccessArgs struct {
	// Group Id, one of the three (TargetUin, RoleId, GroupId) must be passed.
	GroupId *int `pulumi:"groupId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Role Id, one of the three (TargetUin, RoleId, GroupId) must be passed.
	RoleId *int `pulumi:"roleId"`
	// Service type, this field needs to be passed when viewing the details of the service authorization interface.
	ServiceType *string `pulumi:"serviceType"`
	// Sub-account uin, one of the three (TargetUin, RoleId, GroupId) must be passed.
	TargetUin *int `pulumi:"targetUin"`
}

// A collection of values returned by getPolicyGrantingServiceAccess.
type GetPolicyGrantingServiceAccessResult struct {
	GroupId *int `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List.
	Lists            []GetPolicyGrantingServiceAccessList `pulumi:"lists"`
	ResultOutputFile *string                              `pulumi:"resultOutputFile"`
	RoleId           *int                                 `pulumi:"roleId"`
	// Service type.
	ServiceType *string `pulumi:"serviceType"`
	TargetUin   *int    `pulumi:"targetUin"`
}

func GetPolicyGrantingServiceAccessOutput(ctx *pulumi.Context, args GetPolicyGrantingServiceAccessOutputArgs, opts ...pulumi.InvokeOption) GetPolicyGrantingServiceAccessResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPolicyGrantingServiceAccessResult, error) {
			args := v.(GetPolicyGrantingServiceAccessArgs)
			r, err := GetPolicyGrantingServiceAccess(ctx, &args, opts...)
			var s GetPolicyGrantingServiceAccessResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPolicyGrantingServiceAccessResultOutput)
}

// A collection of arguments for invoking getPolicyGrantingServiceAccess.
type GetPolicyGrantingServiceAccessOutputArgs struct {
	// Group Id, one of the three (TargetUin, RoleId, GroupId) must be passed.
	GroupId pulumi.IntPtrInput `pulumi:"groupId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Role Id, one of the three (TargetUin, RoleId, GroupId) must be passed.
	RoleId pulumi.IntPtrInput `pulumi:"roleId"`
	// Service type, this field needs to be passed when viewing the details of the service authorization interface.
	ServiceType pulumi.StringPtrInput `pulumi:"serviceType"`
	// Sub-account uin, one of the three (TargetUin, RoleId, GroupId) must be passed.
	TargetUin pulumi.IntPtrInput `pulumi:"targetUin"`
}

func (GetPolicyGrantingServiceAccessOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyGrantingServiceAccessArgs)(nil)).Elem()
}

// A collection of values returned by getPolicyGrantingServiceAccess.
type GetPolicyGrantingServiceAccessResultOutput struct{ *pulumi.OutputState }

func (GetPolicyGrantingServiceAccessResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyGrantingServiceAccessResult)(nil)).Elem()
}

func (o GetPolicyGrantingServiceAccessResultOutput) ToGetPolicyGrantingServiceAccessResultOutput() GetPolicyGrantingServiceAccessResultOutput {
	return o
}

func (o GetPolicyGrantingServiceAccessResultOutput) ToGetPolicyGrantingServiceAccessResultOutputWithContext(ctx context.Context) GetPolicyGrantingServiceAccessResultOutput {
	return o
}

func (o GetPolicyGrantingServiceAccessResultOutput) GroupId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetPolicyGrantingServiceAccessResult) *int { return v.GroupId }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPolicyGrantingServiceAccessResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyGrantingServiceAccessResult) string { return v.Id }).(pulumi.StringOutput)
}

// List.
func (o GetPolicyGrantingServiceAccessResultOutput) Lists() GetPolicyGrantingServiceAccessListArrayOutput {
	return o.ApplyT(func(v GetPolicyGrantingServiceAccessResult) []GetPolicyGrantingServiceAccessList { return v.Lists }).(GetPolicyGrantingServiceAccessListArrayOutput)
}

func (o GetPolicyGrantingServiceAccessResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyGrantingServiceAccessResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetPolicyGrantingServiceAccessResultOutput) RoleId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetPolicyGrantingServiceAccessResult) *int { return v.RoleId }).(pulumi.IntPtrOutput)
}

// Service type.
func (o GetPolicyGrantingServiceAccessResultOutput) ServiceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyGrantingServiceAccessResult) *string { return v.ServiceType }).(pulumi.StringPtrOutput)
}

func (o GetPolicyGrantingServiceAccessResultOutput) TargetUin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetPolicyGrantingServiceAccessResult) *int { return v.TargetUin }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPolicyGrantingServiceAccessResultOutput{})
}
