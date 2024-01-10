// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of dcdb securityGroups
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dcdb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dcdb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dcdb.GetSecurityGroups(ctx, &dcdb.GetSecurityGroupsArgs{
//				InstanceId: "your_instance_id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSecurityGroups(ctx *pulumi.Context, args *GetSecurityGroupsArgs, opts ...pulumi.InvokeOption) (*GetSecurityGroupsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSecurityGroupsResult
	err := ctx.Invoke("tencentcloud:Dcdb/getSecurityGroups:getSecurityGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecurityGroups.
type GetSecurityGroupsArgs struct {
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getSecurityGroups.
type GetSecurityGroupsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// security group list.
	Lists            []GetSecurityGroupsList `pulumi:"lists"`
	ResultOutputFile *string                 `pulumi:"resultOutputFile"`
}

func GetSecurityGroupsOutput(ctx *pulumi.Context, args GetSecurityGroupsOutputArgs, opts ...pulumi.InvokeOption) GetSecurityGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecurityGroupsResult, error) {
			args := v.(GetSecurityGroupsArgs)
			r, err := GetSecurityGroups(ctx, &args, opts...)
			var s GetSecurityGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSecurityGroupsResultOutput)
}

// A collection of arguments for invoking getSecurityGroups.
type GetSecurityGroupsOutputArgs struct {
	// instance id.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetSecurityGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getSecurityGroups.
type GetSecurityGroupsResultOutput struct{ *pulumi.OutputState }

func (GetSecurityGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityGroupsResult)(nil)).Elem()
}

func (o GetSecurityGroupsResultOutput) ToGetSecurityGroupsResultOutput() GetSecurityGroupsResultOutput {
	return o
}

func (o GetSecurityGroupsResultOutput) ToGetSecurityGroupsResultOutputWithContext(ctx context.Context) GetSecurityGroupsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSecurityGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSecurityGroupsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// security group list.
func (o GetSecurityGroupsResultOutput) Lists() GetSecurityGroupsListArrayOutput {
	return o.ApplyT(func(v GetSecurityGroupsResult) []GetSecurityGroupsList { return v.Lists }).(GetSecurityGroupsListArrayOutput)
}

func (o GetSecurityGroupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecurityGroupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecurityGroupsResultOutput{})
}
