// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of scf layerVersions
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Scf"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Scf"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Scf.GetLayerVersions(ctx, &scf.GetLayerVersionsArgs{
// 			LayerName: "tf-test",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetLayerVersions(ctx *pulumi.Context, args *GetLayerVersionsArgs, opts ...pulumi.InvokeOption) (*GetLayerVersionsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetLayerVersionsResult
	err := ctx.Invoke("tencentcloud:Scf/getLayerVersions:getLayerVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLayerVersions.
type GetLayerVersionsArgs struct {
	// Compatible runtimes.
	CompatibleRuntimes []string `pulumi:"compatibleRuntimes"`
	// Layer name.
	LayerName string `pulumi:"layerName"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getLayerVersions.
type GetLayerVersionsResult struct {
	CompatibleRuntimes []string `pulumi:"compatibleRuntimes"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Layer name.
	LayerName string `pulumi:"layerName"`
	// Layer version list.
	LayerVersions    []GetLayerVersionsLayerVersion `pulumi:"layerVersions"`
	ResultOutputFile *string                        `pulumi:"resultOutputFile"`
}

func GetLayerVersionsOutput(ctx *pulumi.Context, args GetLayerVersionsOutputArgs, opts ...pulumi.InvokeOption) GetLayerVersionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLayerVersionsResult, error) {
			args := v.(GetLayerVersionsArgs)
			r, err := GetLayerVersions(ctx, &args, opts...)
			var s GetLayerVersionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLayerVersionsResultOutput)
}

// A collection of arguments for invoking getLayerVersions.
type GetLayerVersionsOutputArgs struct {
	// Compatible runtimes.
	CompatibleRuntimes pulumi.StringArrayInput `pulumi:"compatibleRuntimes"`
	// Layer name.
	LayerName pulumi.StringInput `pulumi:"layerName"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetLayerVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLayerVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getLayerVersions.
type GetLayerVersionsResultOutput struct{ *pulumi.OutputState }

func (GetLayerVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLayerVersionsResult)(nil)).Elem()
}

func (o GetLayerVersionsResultOutput) ToGetLayerVersionsResultOutput() GetLayerVersionsResultOutput {
	return o
}

func (o GetLayerVersionsResultOutput) ToGetLayerVersionsResultOutputWithContext(ctx context.Context) GetLayerVersionsResultOutput {
	return o
}

func (o GetLayerVersionsResultOutput) CompatibleRuntimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLayerVersionsResult) []string { return v.CompatibleRuntimes }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLayerVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLayerVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Layer name.
func (o GetLayerVersionsResultOutput) LayerName() pulumi.StringOutput {
	return o.ApplyT(func(v GetLayerVersionsResult) string { return v.LayerName }).(pulumi.StringOutput)
}

// Layer version list.
func (o GetLayerVersionsResultOutput) LayerVersions() GetLayerVersionsLayerVersionArrayOutput {
	return o.ApplyT(func(v GetLayerVersionsResult) []GetLayerVersionsLayerVersion { return v.LayerVersions }).(GetLayerVersionsLayerVersionArrayOutput)
}

func (o GetLayerVersionsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLayerVersionsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLayerVersionsResultOutput{})
}
