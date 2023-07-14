// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of cynosdb projectSecurityGroups
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cynosdb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cynosdb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cynosdb.GetProjectSecurityGroups(ctx, &cynosdb.GetProjectSecurityGroupsArgs{
// 			ProjectId: pulumi.IntRef(1250480),
// 			SearchKey: pulumi.StringRef("自定义模版"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetProjectSecurityGroups(ctx *pulumi.Context, args *GetProjectSecurityGroupsArgs, opts ...pulumi.InvokeOption) (*GetProjectSecurityGroupsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetProjectSecurityGroupsResult
	err := ctx.Invoke("tencentcloud:Cynosdb/getProjectSecurityGroups:getProjectSecurityGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectSecurityGroups.
type GetProjectSecurityGroupsArgs struct {
	// Project ID.
	ProjectId *int `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Search Keywords.
	SearchKey *string `pulumi:"searchKey"`
}

// A collection of values returned by getProjectSecurityGroups.
type GetProjectSecurityGroupsResult struct {
	// Security Group Details.
	Groups []GetProjectSecurityGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Project ID.
	ProjectId        *int    `pulumi:"projectId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	SearchKey        *string `pulumi:"searchKey"`
}

func GetProjectSecurityGroupsOutput(ctx *pulumi.Context, args GetProjectSecurityGroupsOutputArgs, opts ...pulumi.InvokeOption) GetProjectSecurityGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectSecurityGroupsResult, error) {
			args := v.(GetProjectSecurityGroupsArgs)
			r, err := GetProjectSecurityGroups(ctx, &args, opts...)
			var s GetProjectSecurityGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectSecurityGroupsResultOutput)
}

// A collection of arguments for invoking getProjectSecurityGroups.
type GetProjectSecurityGroupsOutputArgs struct {
	// Project ID.
	ProjectId pulumi.IntPtrInput `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Search Keywords.
	SearchKey pulumi.StringPtrInput `pulumi:"searchKey"`
}

func (GetProjectSecurityGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectSecurityGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getProjectSecurityGroups.
type GetProjectSecurityGroupsResultOutput struct{ *pulumi.OutputState }

func (GetProjectSecurityGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectSecurityGroupsResult)(nil)).Elem()
}

func (o GetProjectSecurityGroupsResultOutput) ToGetProjectSecurityGroupsResultOutput() GetProjectSecurityGroupsResultOutput {
	return o
}

func (o GetProjectSecurityGroupsResultOutput) ToGetProjectSecurityGroupsResultOutputWithContext(ctx context.Context) GetProjectSecurityGroupsResultOutput {
	return o
}

// Security Group Details.
func (o GetProjectSecurityGroupsResultOutput) Groups() GetProjectSecurityGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetProjectSecurityGroupsResult) []GetProjectSecurityGroupsGroup { return v.Groups }).(GetProjectSecurityGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectSecurityGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectSecurityGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Project ID.
func (o GetProjectSecurityGroupsResultOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetProjectSecurityGroupsResult) *int { return v.ProjectId }).(pulumi.IntPtrOutput)
}

func (o GetProjectSecurityGroupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectSecurityGroupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetProjectSecurityGroupsResultOutput) SearchKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectSecurityGroupsResult) *string { return v.SearchKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectSecurityGroupsResultOutput{})
}
