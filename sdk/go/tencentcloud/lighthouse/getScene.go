// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lighthouse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of lighthouse scene with region
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Lighthouse"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Lighthouse.GetScene(ctx, &lighthouse.GetSceneArgs{
//				Limit:  pulumi.IntRef(20),
//				Offset: pulumi.IntRef(0),
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
func GetScene(ctx *pulumi.Context, args *GetSceneArgs, opts ...pulumi.InvokeOption) (*GetSceneResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSceneResult
	err := ctx.Invoke("tencentcloud:Lighthouse/getScene:getScene", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getScene.
type GetSceneArgs struct {
	// Number of returned results. Default value is 20. Maximum value is 100.
	Limit *int `pulumi:"limit"`
	// Offset. Default value is 0.
	Offset *int `pulumi:"offset"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// List of scene IDs.
	SceneIds []string `pulumi:"sceneIds"`
}

// A collection of values returned by getScene.
type GetSceneResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string   `pulumi:"id"`
	Limit            *int     `pulumi:"limit"`
	Offset           *int     `pulumi:"offset"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
	SceneIds         []string `pulumi:"sceneIds"`
	// List of scene info.
	SceneSets []GetSceneSceneSet `pulumi:"sceneSets"`
}

func GetSceneOutput(ctx *pulumi.Context, args GetSceneOutputArgs, opts ...pulumi.InvokeOption) GetSceneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSceneResult, error) {
			args := v.(GetSceneArgs)
			r, err := GetScene(ctx, &args, opts...)
			var s GetSceneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSceneResultOutput)
}

// A collection of arguments for invoking getScene.
type GetSceneOutputArgs struct {
	// Number of returned results. Default value is 20. Maximum value is 100.
	Limit pulumi.IntPtrInput `pulumi:"limit"`
	// Offset. Default value is 0.
	Offset pulumi.IntPtrInput `pulumi:"offset"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// List of scene IDs.
	SceneIds pulumi.StringArrayInput `pulumi:"sceneIds"`
}

func (GetSceneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSceneArgs)(nil)).Elem()
}

// A collection of values returned by getScene.
type GetSceneResultOutput struct{ *pulumi.OutputState }

func (GetSceneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSceneResult)(nil)).Elem()
}

func (o GetSceneResultOutput) ToGetSceneResultOutput() GetSceneResultOutput {
	return o
}

func (o GetSceneResultOutput) ToGetSceneResultOutputWithContext(ctx context.Context) GetSceneResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSceneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSceneResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSceneResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetSceneResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

func (o GetSceneResultOutput) Offset() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetSceneResult) *int { return v.Offset }).(pulumi.IntPtrOutput)
}

func (o GetSceneResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSceneResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetSceneResultOutput) SceneIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSceneResult) []string { return v.SceneIds }).(pulumi.StringArrayOutput)
}

// List of scene info.
func (o GetSceneResultOutput) SceneSets() GetSceneSceneSetArrayOutput {
	return o.ApplyT(func(v GetSceneResult) []GetSceneSceneSet { return v.SceneSets }).(GetSceneSceneSetArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSceneResultOutput{})
}
