// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of CAM roles
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
//			_, err := Cam.GetRoles(ctx, &cam.GetRolesArgs{
//				RoleId: pulumi.StringRef(tencentcloud_cam_role.Foo.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Cam.GetRoles(ctx, &cam.GetRolesArgs{
//				Name: pulumi.StringRef("cam-role-test"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetRoles(ctx *pulumi.Context, args *GetRolesArgs, opts ...pulumi.InvokeOption) (*GetRolesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRolesResult
	err := ctx.Invoke("tencentcloud:Cam/getRoles:getRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRoles.
type GetRolesArgs struct {
	// The description of the CAM role to be queried.
	Description *string `pulumi:"description"`
	// Name of the CAM policy to be queried.
	Name *string `pulumi:"name"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of the CAM role to be queried.
	RoleId *string `pulumi:"roleId"`
}

// A collection of values returned by getRoles.
type GetRolesResult struct {
	// Description of CAM role.
	Description *string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of CAM role.
	Name             *string `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Id of CAM role.
	RoleId *string `pulumi:"roleId"`
	// A list of CAM roles. Each element contains the following attributes:
	RoleLists []GetRolesRoleList `pulumi:"roleLists"`
}

func GetRolesOutput(ctx *pulumi.Context, args GetRolesOutputArgs, opts ...pulumi.InvokeOption) GetRolesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRolesResult, error) {
			args := v.(GetRolesArgs)
			r, err := GetRoles(ctx, &args, opts...)
			var s GetRolesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRolesResultOutput)
}

// A collection of arguments for invoking getRoles.
type GetRolesOutputArgs struct {
	// The description of the CAM role to be queried.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Name of the CAM policy to be queried.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// ID of the CAM role to be queried.
	RoleId pulumi.StringPtrInput `pulumi:"roleId"`
}

func (GetRolesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRolesArgs)(nil)).Elem()
}

// A collection of values returned by getRoles.
type GetRolesResultOutput struct{ *pulumi.OutputState }

func (GetRolesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRolesResult)(nil)).Elem()
}

func (o GetRolesResultOutput) ToGetRolesResultOutput() GetRolesResultOutput {
	return o
}

func (o GetRolesResultOutput) ToGetRolesResultOutputWithContext(ctx context.Context) GetRolesResultOutput {
	return o
}

// Description of CAM role.
func (o GetRolesResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRolesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRolesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of CAM role.
func (o GetRolesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetRolesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Id of CAM role.
func (o GetRolesResultOutput) RoleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.RoleId }).(pulumi.StringPtrOutput)
}

// A list of CAM roles. Each element contains the following attributes:
func (o GetRolesResultOutput) RoleLists() GetRolesRoleListArrayOutput {
	return o.ApplyT(func(v GetRolesResult) []GetRolesRoleList { return v.RoleLists }).(GetRolesRoleListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRolesResultOutput{})
}
