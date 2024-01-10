// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of mariadb instanceSpecs
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mariadb.GetInstanceSpecs(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetInstanceSpecs(ctx *pulumi.Context, args *GetInstanceSpecsArgs, opts ...pulumi.InvokeOption) (*GetInstanceSpecsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetInstanceSpecsResult
	err := ctx.Invoke("tencentcloud:Mariadb/getInstanceSpecs:getInstanceSpecs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceSpecs.
type GetInstanceSpecsArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getInstanceSpecs.
type GetInstanceSpecsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// list of instance specifications.
	Specs []GetInstanceSpecsSpec `pulumi:"specs"`
}

func GetInstanceSpecsOutput(ctx *pulumi.Context, args GetInstanceSpecsOutputArgs, opts ...pulumi.InvokeOption) GetInstanceSpecsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceSpecsResult, error) {
			args := v.(GetInstanceSpecsArgs)
			r, err := GetInstanceSpecs(ctx, &args, opts...)
			var s GetInstanceSpecsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceSpecsResultOutput)
}

// A collection of arguments for invoking getInstanceSpecs.
type GetInstanceSpecsOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetInstanceSpecsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceSpecsArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceSpecs.
type GetInstanceSpecsResultOutput struct{ *pulumi.OutputState }

func (GetInstanceSpecsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceSpecsResult)(nil)).Elem()
}

func (o GetInstanceSpecsResultOutput) ToGetInstanceSpecsResultOutput() GetInstanceSpecsResultOutput {
	return o
}

func (o GetInstanceSpecsResultOutput) ToGetInstanceSpecsResultOutputWithContext(ctx context.Context) GetInstanceSpecsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceSpecsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceSpecsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstanceSpecsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceSpecsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// list of instance specifications.
func (o GetInstanceSpecsResultOutput) Specs() GetInstanceSpecsSpecArrayOutput {
	return o.ApplyT(func(v GetInstanceSpecsResult) []GetInstanceSpecsSpec { return v.Specs }).(GetInstanceSpecsSpecArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceSpecsResultOutput{})
}
