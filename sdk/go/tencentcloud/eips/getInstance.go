// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eips

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query eip instances.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Eips"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Eips"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Eips.GetInstance(ctx, &eips.GetInstanceArgs{
// 			EipId: pulumi.StringRef("eip-ry9h95hg"),
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
	err := ctx.Invoke("tencentcloud:Eips/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type GetInstanceArgs struct {
	// ID of the EIP to be queried.
	EipId *string `pulumi:"eipId"`
	// Name of the EIP to be queried.
	EipName *string `pulumi:"eipName"`
	// The elastic ip address.
	PublicIp *string `pulumi:"publicIp"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// The tags of EIP.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getInstance.
type GetInstanceResult struct {
	// ID of the EIP.
	EipId *string `pulumi:"eipId"`
	// An information list of EIP. Each element contains the following attributes:
	EipLists []GetInstanceEipList `pulumi:"eipLists"`
	// Name of the EIP.
	EipName *string `pulumi:"eipName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The elastic ip address.
	PublicIp         *string `pulumi:"publicIp"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Tags of the EIP.
	Tags map[string]interface{} `pulumi:"tags"`
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
	// ID of the EIP to be queried.
	EipId pulumi.StringPtrInput `pulumi:"eipId"`
	// Name of the EIP to be queried.
	EipName pulumi.StringPtrInput `pulumi:"eipName"`
	// The elastic ip address.
	PublicIp pulumi.StringPtrInput `pulumi:"publicIp"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// The tags of EIP.
	Tags pulumi.MapInput `pulumi:"tags"`
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

// ID of the EIP.
func (o GetInstanceResultOutput) EipId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.EipId }).(pulumi.StringPtrOutput)
}

// An information list of EIP. Each element contains the following attributes:
func (o GetInstanceResultOutput) EipLists() GetInstanceEipListArrayOutput {
	return o.ApplyT(func(v GetInstanceResult) []GetInstanceEipList { return v.EipLists }).(GetInstanceEipListArrayOutput)
}

// Name of the EIP.
func (o GetInstanceResultOutput) EipName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.EipName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The elastic ip address.
func (o GetInstanceResultOutput) PublicIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.PublicIp }).(pulumi.StringPtrOutput)
}

func (o GetInstanceResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Tags of the EIP.
func (o GetInstanceResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetInstanceResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceResultOutput{})
}
