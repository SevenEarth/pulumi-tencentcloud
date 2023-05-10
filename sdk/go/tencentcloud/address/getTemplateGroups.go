// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package address

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of address template groups.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Address"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Address"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Address.GetTemplateGroups(ctx, &address.GetTemplateGroupsArgs{
// 			Name: pulumi.StringRef("test"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetTemplateGroups(ctx *pulumi.Context, args *GetTemplateGroupsArgs, opts ...pulumi.InvokeOption) (*GetTemplateGroupsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetTemplateGroupsResult
	err := ctx.Invoke("tencentcloud:Address/getTemplateGroups:getTemplateGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTemplateGroups.
type GetTemplateGroupsArgs struct {
	// Id of the address template group to query.
	Id *string `pulumi:"id"`
	// Name of the address template group to query.
	Name *string `pulumi:"name"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getTemplateGroups.
type GetTemplateGroupsResult struct {
	// Information list of the dedicated address template groups.
	GroupLists []GetTemplateGroupsGroupList `pulumi:"groupLists"`
	// Id of the address template group.
	Id *string `pulumi:"id"`
	// Name of address template group.
	Name             *string `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetTemplateGroupsOutput(ctx *pulumi.Context, args GetTemplateGroupsOutputArgs, opts ...pulumi.InvokeOption) GetTemplateGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTemplateGroupsResult, error) {
			args := v.(GetTemplateGroupsArgs)
			r, err := GetTemplateGroups(ctx, &args, opts...)
			var s GetTemplateGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTemplateGroupsResultOutput)
}

// A collection of arguments for invoking getTemplateGroups.
type GetTemplateGroupsOutputArgs struct {
	// Id of the address template group to query.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of the address template group to query.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetTemplateGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplateGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getTemplateGroups.
type GetTemplateGroupsResultOutput struct{ *pulumi.OutputState }

func (GetTemplateGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplateGroupsResult)(nil)).Elem()
}

func (o GetTemplateGroupsResultOutput) ToGetTemplateGroupsResultOutput() GetTemplateGroupsResultOutput {
	return o
}

func (o GetTemplateGroupsResultOutput) ToGetTemplateGroupsResultOutputWithContext(ctx context.Context) GetTemplateGroupsResultOutput {
	return o
}

// Information list of the dedicated address template groups.
func (o GetTemplateGroupsResultOutput) GroupLists() GetTemplateGroupsGroupListArrayOutput {
	return o.ApplyT(func(v GetTemplateGroupsResult) []GetTemplateGroupsGroupList { return v.GroupLists }).(GetTemplateGroupsGroupListArrayOutput)
}

// Id of the address template group.
func (o GetTemplateGroupsResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTemplateGroupsResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Name of address template group.
func (o GetTemplateGroupsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTemplateGroupsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetTemplateGroupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTemplateGroupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTemplateGroupsResultOutput{})
}
