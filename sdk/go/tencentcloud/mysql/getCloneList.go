// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of mysql cloneList
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mysql.GetCloneList(ctx, &mysql.GetCloneListArgs{
//				InstanceId: "cdb-fitq5t9h",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCloneList(ctx *pulumi.Context, args *GetCloneListArgs, opts ...pulumi.InvokeOption) (*GetCloneListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetCloneListResult
	err := ctx.Invoke("tencentcloud:Mysql/getCloneList:getCloneList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloneList.
type GetCloneListArgs struct {
	// Query the list of cloning tasks for the specified source instance.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getCloneList.
type GetCloneListResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// Clone task list.
	Items            []GetCloneListItem `pulumi:"items"`
	ResultOutputFile *string            `pulumi:"resultOutputFile"`
}

func GetCloneListOutput(ctx *pulumi.Context, args GetCloneListOutputArgs, opts ...pulumi.InvokeOption) GetCloneListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloneListResult, error) {
			args := v.(GetCloneListArgs)
			r, err := GetCloneList(ctx, &args, opts...)
			var s GetCloneListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloneListResultOutput)
}

// A collection of arguments for invoking getCloneList.
type GetCloneListOutputArgs struct {
	// Query the list of cloning tasks for the specified source instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetCloneListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloneListArgs)(nil)).Elem()
}

// A collection of values returned by getCloneList.
type GetCloneListResultOutput struct{ *pulumi.OutputState }

func (GetCloneListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloneListResult)(nil)).Elem()
}

func (o GetCloneListResultOutput) ToGetCloneListResultOutput() GetCloneListResultOutput {
	return o
}

func (o GetCloneListResultOutput) ToGetCloneListResultOutputWithContext(ctx context.Context) GetCloneListResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetCloneListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloneListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCloneListResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloneListResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Clone task list.
func (o GetCloneListResultOutput) Items() GetCloneListItemArrayOutput {
	return o.ApplyT(func(v GetCloneListResult) []GetCloneListItem { return v.Items }).(GetCloneListItemArrayOutput)
}

func (o GetCloneListResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCloneListResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloneListResultOutput{})
}
