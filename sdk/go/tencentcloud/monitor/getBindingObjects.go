// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query policy group binding objects.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		name, err := Monitor.GetPolicyGroups(ctx, &monitor.GetPolicyGroupsArgs{
// 			Name: pulumi.StringRef("test"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Monitor.GetBindingObjects(ctx, &monitor.GetBindingObjectsArgs{
// 			GroupId: name.Lists[0].GroupId,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetBindingObjects(ctx *pulumi.Context, args *GetBindingObjectsArgs, opts ...pulumi.InvokeOption) (*GetBindingObjectsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetBindingObjectsResult
	err := ctx.Invoke("tencentcloud:Monitor/getBindingObjects:getBindingObjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBindingObjects.
type GetBindingObjectsArgs struct {
	// Policy group ID for query.
	GroupId int `pulumi:"groupId"`
	// Used to store results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getBindingObjects.
type GetBindingObjectsResult struct {
	GroupId int `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list objects. Each element contains the following attributes:
	Lists            []GetBindingObjectsList `pulumi:"lists"`
	ResultOutputFile *string                 `pulumi:"resultOutputFile"`
}

func GetBindingObjectsOutput(ctx *pulumi.Context, args GetBindingObjectsOutputArgs, opts ...pulumi.InvokeOption) GetBindingObjectsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBindingObjectsResult, error) {
			args := v.(GetBindingObjectsArgs)
			r, err := GetBindingObjects(ctx, &args, opts...)
			var s GetBindingObjectsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBindingObjectsResultOutput)
}

// A collection of arguments for invoking getBindingObjects.
type GetBindingObjectsOutputArgs struct {
	// Policy group ID for query.
	GroupId pulumi.IntInput `pulumi:"groupId"`
	// Used to store results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetBindingObjectsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBindingObjectsArgs)(nil)).Elem()
}

// A collection of values returned by getBindingObjects.
type GetBindingObjectsResultOutput struct{ *pulumi.OutputState }

func (GetBindingObjectsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBindingObjectsResult)(nil)).Elem()
}

func (o GetBindingObjectsResultOutput) ToGetBindingObjectsResultOutput() GetBindingObjectsResultOutput {
	return o
}

func (o GetBindingObjectsResultOutput) ToGetBindingObjectsResultOutputWithContext(ctx context.Context) GetBindingObjectsResultOutput {
	return o
}

func (o GetBindingObjectsResultOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v GetBindingObjectsResult) int { return v.GroupId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBindingObjectsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBindingObjectsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list objects. Each element contains the following attributes:
func (o GetBindingObjectsResultOutput) Lists() GetBindingObjectsListArrayOutput {
	return o.ApplyT(func(v GetBindingObjectsResult) []GetBindingObjectsList { return v.Lists }).(GetBindingObjectsListArrayOutput)
}

func (o GetBindingObjectsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBindingObjectsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBindingObjectsResultOutput{})
}
