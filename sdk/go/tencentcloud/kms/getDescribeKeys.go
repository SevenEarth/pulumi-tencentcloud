// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of kms keyLists
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Kms"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Kms.GetDescribeKeys(ctx, &kms.GetDescribeKeysArgs{
//				KeyIds: []string{
//					"9ffacc8b-6461-11ee-a54e-525400dd8a7d",
//					"bffae4ed-6465-11ee-90b2-5254000ef00e",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetDescribeKeys(ctx *pulumi.Context, args *GetDescribeKeysArgs, opts ...pulumi.InvokeOption) (*GetDescribeKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDescribeKeysResult
	err := ctx.Invoke("tencentcloud:Kms/getDescribeKeys:getDescribeKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDescribeKeys.
type GetDescribeKeysArgs struct {
	// Query the ID list of CMK, batch query supports up to 100 KeyIds at a time.
	KeyIds []string `pulumi:"keyIds"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDescribeKeys.
type GetDescribeKeysResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id     string   `pulumi:"id"`
	KeyIds []string `pulumi:"keyIds"`
	// A list of KMS keys.
	KeyLists         []GetDescribeKeysKeyList `pulumi:"keyLists"`
	ResultOutputFile *string                  `pulumi:"resultOutputFile"`
}

func GetDescribeKeysOutput(ctx *pulumi.Context, args GetDescribeKeysOutputArgs, opts ...pulumi.InvokeOption) GetDescribeKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDescribeKeysResult, error) {
			args := v.(GetDescribeKeysArgs)
			r, err := GetDescribeKeys(ctx, &args, opts...)
			var s GetDescribeKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDescribeKeysResultOutput)
}

// A collection of arguments for invoking getDescribeKeys.
type GetDescribeKeysOutputArgs struct {
	// Query the ID list of CMK, batch query supports up to 100 KeyIds at a time.
	KeyIds pulumi.StringArrayInput `pulumi:"keyIds"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDescribeKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeKeysArgs)(nil)).Elem()
}

// A collection of values returned by getDescribeKeys.
type GetDescribeKeysResultOutput struct{ *pulumi.OutputState }

func (GetDescribeKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeKeysResult)(nil)).Elem()
}

func (o GetDescribeKeysResultOutput) ToGetDescribeKeysResultOutput() GetDescribeKeysResultOutput {
	return o
}

func (o GetDescribeKeysResultOutput) ToGetDescribeKeysResultOutputWithContext(ctx context.Context) GetDescribeKeysResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetDescribeKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDescribeKeysResultOutput) KeyIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDescribeKeysResult) []string { return v.KeyIds }).(pulumi.StringArrayOutput)
}

// A list of KMS keys.
func (o GetDescribeKeysResultOutput) KeyLists() GetDescribeKeysKeyListArrayOutput {
	return o.ApplyT(func(v GetDescribeKeysResult) []GetDescribeKeysKeyList { return v.KeyLists }).(GetDescribeKeysKeyListArrayOutput)
}

func (o GetDescribeKeysResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeKeysResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDescribeKeysResultOutput{})
}
