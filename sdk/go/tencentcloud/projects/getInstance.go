// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tag project
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Projects"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Projects"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Projects.GetInstance(ctx, &projects.GetInstanceArgs{
// 			AllList: 1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetInstance(ctx *pulumi.Context, args *GetInstanceArgs, opts ...pulumi.InvokeOption) (*GetInstanceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetInstanceResult
	err := ctx.Invoke("tencentcloud:Projects/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type GetInstanceArgs struct {
	// 1 means to list all project, 0 means to list visible project.
	AllList int `pulumi:"allList"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getInstance.
type GetInstanceResult struct {
	AllList int `pulumi:"allList"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of projects.
	Projects         []GetInstanceProject `pulumi:"projects"`
	ResultOutputFile *string              `pulumi:"resultOutputFile"`
}

func GetInstanceOutput(ctx *pulumi.Context, args GetInstanceOutputArgs, opts ...pulumi.InvokeOption) GetInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceResult, error) {
			args := v.(GetInstanceArgs)
			r, err := GetInstance(ctx, &args, opts...)
			var s GetInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceResultOutput)
}

// A collection of arguments for invoking getInstance.
type GetInstanceOutputArgs struct {
	// 1 means to list all project, 0 means to list visible project.
	AllList pulumi.IntInput `pulumi:"allList"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getInstance.
type GetInstanceResultOutput struct{ *pulumi.OutputState }

func (GetInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceResult)(nil)).Elem()
}

func (o GetInstanceResultOutput) ToGetInstanceResultOutput() GetInstanceResultOutput {
	return o
}

func (o GetInstanceResultOutput) ToGetInstanceResultOutputWithContext(ctx context.Context) GetInstanceResultOutput {
	return o
}

func (o GetInstanceResultOutput) AllList() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstanceResult) int { return v.AllList }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of projects.
func (o GetInstanceResultOutput) Projects() GetInstanceProjectArrayOutput {
	return o.ApplyT(func(v GetInstanceResult) []GetInstanceProject { return v.Projects }).(GetInstanceProjectArrayOutput)
}

func (o GetInstanceResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceResultOutput{})
}