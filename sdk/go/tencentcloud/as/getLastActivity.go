// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package as

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of as lastActivity
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/As"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/As"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := As.GetLastActivity(ctx, &as.GetLastActivityArgs{
// 			AutoScalingGroupIds: []string{
// 				"asc-lo0b94oy",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetLastActivity(ctx *pulumi.Context, args *GetLastActivityArgs, opts ...pulumi.InvokeOption) (*GetLastActivityResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetLastActivityResult
	err := ctx.Invoke("tencentcloud:As/getLastActivity:getLastActivity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLastActivity.
type GetLastActivityArgs struct {
	// ID list of an auto scaling group.
	AutoScalingGroupIds []string `pulumi:"autoScalingGroupIds"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getLastActivity.
type GetLastActivityResult struct {
	// Information set of eligible scaling activities. Scaling groups without scaling activities are not returned. For example, if there are 50 auto scaling group IDs but only 45 records are returned, it indicates that 5 of the auto scaling groups do not have scaling activities.
	ActivitySets        []GetLastActivityActivitySet `pulumi:"activitySets"`
	AutoScalingGroupIds []string                     `pulumi:"autoScalingGroupIds"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetLastActivityOutput(ctx *pulumi.Context, args GetLastActivityOutputArgs, opts ...pulumi.InvokeOption) GetLastActivityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLastActivityResult, error) {
			args := v.(GetLastActivityArgs)
			r, err := GetLastActivity(ctx, &args, opts...)
			var s GetLastActivityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLastActivityResultOutput)
}

// A collection of arguments for invoking getLastActivity.
type GetLastActivityOutputArgs struct {
	// ID list of an auto scaling group.
	AutoScalingGroupIds pulumi.StringArrayInput `pulumi:"autoScalingGroupIds"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetLastActivityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLastActivityArgs)(nil)).Elem()
}

// A collection of values returned by getLastActivity.
type GetLastActivityResultOutput struct{ *pulumi.OutputState }

func (GetLastActivityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLastActivityResult)(nil)).Elem()
}

func (o GetLastActivityResultOutput) ToGetLastActivityResultOutput() GetLastActivityResultOutput {
	return o
}

func (o GetLastActivityResultOutput) ToGetLastActivityResultOutputWithContext(ctx context.Context) GetLastActivityResultOutput {
	return o
}

// Information set of eligible scaling activities. Scaling groups without scaling activities are not returned. For example, if there are 50 auto scaling group IDs but only 45 records are returned, it indicates that 5 of the auto scaling groups do not have scaling activities.
func (o GetLastActivityResultOutput) ActivitySets() GetLastActivityActivitySetArrayOutput {
	return o.ApplyT(func(v GetLastActivityResult) []GetLastActivityActivitySet { return v.ActivitySets }).(GetLastActivityActivitySetArrayOutput)
}

func (o GetLastActivityResultOutput) AutoScalingGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLastActivityResult) []string { return v.AutoScalingGroupIds }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLastActivityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLastActivityResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLastActivityResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLastActivityResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLastActivityResultOutput{})
}
