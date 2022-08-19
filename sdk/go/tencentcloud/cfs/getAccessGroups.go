// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query the detail information of CFS access group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cfs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cfs"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cfs.GetAccessGroups(ctx, &cfs.GetAccessGroupsArgs{
// 			AccessGroupId: pulumi.StringRef("pgroup-7nx89k7l"),
// 			Name:          pulumi.StringRef("test"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetAccessGroups(ctx *pulumi.Context, args *GetAccessGroupsArgs, opts ...pulumi.InvokeOption) (*GetAccessGroupsResult, error) {
	var rv GetAccessGroupsResult
	err := ctx.Invoke("tencentcloud:Cfs/getAccessGroups:getAccessGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessGroups.
type GetAccessGroupsArgs struct {
	// A specified access group ID used to query.
	AccessGroupId *string `pulumi:"accessGroupId"`
	// A access group Name used to query.
	Name *string `pulumi:"name"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getAccessGroups.
type GetAccessGroupsResult struct {
	// ID of the access group.
	AccessGroupId *string `pulumi:"accessGroupId"`
	// An information list of CFS access group. Each element contains the following attributes:
	AccessGroupLists []GetAccessGroupsAccessGroupList `pulumi:"accessGroupLists"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the access group.
	Name             *string `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetAccessGroupsOutput(ctx *pulumi.Context, args GetAccessGroupsOutputArgs, opts ...pulumi.InvokeOption) GetAccessGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccessGroupsResult, error) {
			args := v.(GetAccessGroupsArgs)
			r, err := GetAccessGroups(ctx, &args, opts...)
			var s GetAccessGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAccessGroupsResultOutput)
}

// A collection of arguments for invoking getAccessGroups.
type GetAccessGroupsOutputArgs struct {
	// A specified access group ID used to query.
	AccessGroupId pulumi.StringPtrInput `pulumi:"accessGroupId"`
	// A access group Name used to query.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetAccessGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getAccessGroups.
type GetAccessGroupsResultOutput struct{ *pulumi.OutputState }

func (GetAccessGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessGroupsResult)(nil)).Elem()
}

func (o GetAccessGroupsResultOutput) ToGetAccessGroupsResultOutput() GetAccessGroupsResultOutput {
	return o
}

func (o GetAccessGroupsResultOutput) ToGetAccessGroupsResultOutputWithContext(ctx context.Context) GetAccessGroupsResultOutput {
	return o
}

// ID of the access group.
func (o GetAccessGroupsResultOutput) AccessGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessGroupsResult) *string { return v.AccessGroupId }).(pulumi.StringPtrOutput)
}

// An information list of CFS access group. Each element contains the following attributes:
func (o GetAccessGroupsResultOutput) AccessGroupLists() GetAccessGroupsAccessGroupListArrayOutput {
	return o.ApplyT(func(v GetAccessGroupsResult) []GetAccessGroupsAccessGroupList { return v.AccessGroupLists }).(GetAccessGroupsAccessGroupListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccessGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the access group.
func (o GetAccessGroupsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessGroupsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetAccessGroupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessGroupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccessGroupsResultOutput{})
}
