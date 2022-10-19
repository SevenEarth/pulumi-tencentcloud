// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of CAM users
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
//			_, err := Cam.GetUsers(ctx, &cam.GetUsersArgs{
//				Name: pulumi.StringRef("cam-user-test"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Cam.GetUsers(ctx, &cam.GetUsersArgs{
//				Email: pulumi.StringRef("hello@test.com"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Cam.GetUsers(ctx, &cam.GetUsersArgs{
//				PhoneNum: pulumi.StringRef("12345678910"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetUsers(ctx *pulumi.Context, args *GetUsersArgs, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	var rv GetUsersResult
	err := ctx.Invoke("tencentcloud:Cam/getUsers:getUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsers.
type GetUsersArgs struct {
	// Indicate whether the user can login in.
	ConsoleLogin *bool `pulumi:"consoleLogin"`
	// Country code of the CAM user to be queried.
	CountryCode *string `pulumi:"countryCode"`
	// Email of the CAM user to be queried.
	Email *string `pulumi:"email"`
	// Name of CAM user to be queried.
	Name *string `pulumi:"name"`
	// Phone num of the CAM user to be queried.
	PhoneNum *string `pulumi:"phoneNum"`
	// Remark of the CAM user to be queried.
	Remark *string `pulumi:"remark"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Uid of the CAM user to be queried.
	Uid *int `pulumi:"uid"`
	// Uin of the CAM user to be queried.
	Uin *int `pulumi:"uin"`
}

// A collection of values returned by getUsers.
type GetUsersResult struct {
	ConsoleLogin *bool `pulumi:"consoleLogin"`
	// Country code of the CAM user.
	CountryCode *string `pulumi:"countryCode"`
	// Email of the CAM user.
	Email *string `pulumi:"email"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of CAM user.
	Name *string `pulumi:"name"`
	// Phone num of the CAM user.
	PhoneNum *string `pulumi:"phoneNum"`
	// Remark of the CAM user.
	Remark           *string `pulumi:"remark"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Uid of the CAM user.
	Uid *int `pulumi:"uid"`
	// Uin of the CAM user.
	Uin *int `pulumi:"uin"`
	// A list of CAM users. Each element contains the following attributes:
	UserLists []GetUsersUserList `pulumi:"userLists"`
}

func GetUsersOutput(ctx *pulumi.Context, args GetUsersOutputArgs, opts ...pulumi.InvokeOption) GetUsersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUsersResult, error) {
			args := v.(GetUsersArgs)
			r, err := GetUsers(ctx, &args, opts...)
			var s GetUsersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUsersResultOutput)
}

// A collection of arguments for invoking getUsers.
type GetUsersOutputArgs struct {
	// Indicate whether the user can login in.
	ConsoleLogin pulumi.BoolPtrInput `pulumi:"consoleLogin"`
	// Country code of the CAM user to be queried.
	CountryCode pulumi.StringPtrInput `pulumi:"countryCode"`
	// Email of the CAM user to be queried.
	Email pulumi.StringPtrInput `pulumi:"email"`
	// Name of CAM user to be queried.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Phone num of the CAM user to be queried.
	PhoneNum pulumi.StringPtrInput `pulumi:"phoneNum"`
	// Remark of the CAM user to be queried.
	Remark pulumi.StringPtrInput `pulumi:"remark"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Uid of the CAM user to be queried.
	Uid pulumi.IntPtrInput `pulumi:"uid"`
	// Uin of the CAM user to be queried.
	Uin pulumi.IntPtrInput `pulumi:"uin"`
}

func (GetUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersArgs)(nil)).Elem()
}

// A collection of values returned by getUsers.
type GetUsersResultOutput struct{ *pulumi.OutputState }

func (GetUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersResult)(nil)).Elem()
}

func (o GetUsersResultOutput) ToGetUsersResultOutput() GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) ToGetUsersResultOutputWithContext(ctx context.Context) GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) ConsoleLogin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *bool { return v.ConsoleLogin }).(pulumi.BoolPtrOutput)
}

// Country code of the CAM user.
func (o GetUsersResultOutput) CountryCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.CountryCode }).(pulumi.StringPtrOutput)
}

// Email of the CAM user.
func (o GetUsersResultOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Email }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of CAM user.
func (o GetUsersResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Phone num of the CAM user.
func (o GetUsersResultOutput) PhoneNum() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.PhoneNum }).(pulumi.StringPtrOutput)
}

// Remark of the CAM user.
func (o GetUsersResultOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Remark }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Uid of the CAM user.
func (o GetUsersResultOutput) Uid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *int { return v.Uid }).(pulumi.IntPtrOutput)
}

// Uin of the CAM user.
func (o GetUsersResultOutput) Uin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *int { return v.Uin }).(pulumi.IntPtrOutput)
}

// A list of CAM users. Each element contains the following attributes:
func (o GetUsersResultOutput) UserLists() GetUsersUserListArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []GetUsersUserList { return v.UserLists }).(GetUsersUserListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUsersResultOutput{})
}
