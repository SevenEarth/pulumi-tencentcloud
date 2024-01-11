// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dlc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of dlc checkDataEngineImageCanBeUpgrade
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dlc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dlc"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dlc.GetCheckDataEngineImageCanBeUpgrade(ctx, &dlc.GetCheckDataEngineImageCanBeUpgradeArgs{
// 			DataEngineId: "DataEngine-cgkvbas6",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetCheckDataEngineImageCanBeUpgrade(ctx *pulumi.Context, args *GetCheckDataEngineImageCanBeUpgradeArgs, opts ...pulumi.InvokeOption) (*GetCheckDataEngineImageCanBeUpgradeResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetCheckDataEngineImageCanBeUpgradeResult
	err := ctx.Invoke("tencentcloud:Dlc/getCheckDataEngineImageCanBeUpgrade:getCheckDataEngineImageCanBeUpgrade", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCheckDataEngineImageCanBeUpgrade.
type GetCheckDataEngineImageCanBeUpgradeArgs struct {
	// Engine unique id.
	DataEngineId string `pulumi:"dataEngineId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getCheckDataEngineImageCanBeUpgrade.
type GetCheckDataEngineImageCanBeUpgradeResult struct {
	// The latest image version id that can be upgraded.
	ChildImageVersionId string `pulumi:"childImageVersionId"`
	DataEngineId        string `pulumi:"dataEngineId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Is it possible to upgrade.
	IsUpgrade        bool    `pulumi:"isUpgrade"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetCheckDataEngineImageCanBeUpgradeOutput(ctx *pulumi.Context, args GetCheckDataEngineImageCanBeUpgradeOutputArgs, opts ...pulumi.InvokeOption) GetCheckDataEngineImageCanBeUpgradeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCheckDataEngineImageCanBeUpgradeResult, error) {
			args := v.(GetCheckDataEngineImageCanBeUpgradeArgs)
			r, err := GetCheckDataEngineImageCanBeUpgrade(ctx, &args, opts...)
			var s GetCheckDataEngineImageCanBeUpgradeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCheckDataEngineImageCanBeUpgradeResultOutput)
}

// A collection of arguments for invoking getCheckDataEngineImageCanBeUpgrade.
type GetCheckDataEngineImageCanBeUpgradeOutputArgs struct {
	// Engine unique id.
	DataEngineId pulumi.StringInput `pulumi:"dataEngineId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetCheckDataEngineImageCanBeUpgradeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCheckDataEngineImageCanBeUpgradeArgs)(nil)).Elem()
}

// A collection of values returned by getCheckDataEngineImageCanBeUpgrade.
type GetCheckDataEngineImageCanBeUpgradeResultOutput struct{ *pulumi.OutputState }

func (GetCheckDataEngineImageCanBeUpgradeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCheckDataEngineImageCanBeUpgradeResult)(nil)).Elem()
}

func (o GetCheckDataEngineImageCanBeUpgradeResultOutput) ToGetCheckDataEngineImageCanBeUpgradeResultOutput() GetCheckDataEngineImageCanBeUpgradeResultOutput {
	return o
}

func (o GetCheckDataEngineImageCanBeUpgradeResultOutput) ToGetCheckDataEngineImageCanBeUpgradeResultOutputWithContext(ctx context.Context) GetCheckDataEngineImageCanBeUpgradeResultOutput {
	return o
}

// The latest image version id that can be upgraded.
func (o GetCheckDataEngineImageCanBeUpgradeResultOutput) ChildImageVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCheckDataEngineImageCanBeUpgradeResult) string { return v.ChildImageVersionId }).(pulumi.StringOutput)
}

func (o GetCheckDataEngineImageCanBeUpgradeResultOutput) DataEngineId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCheckDataEngineImageCanBeUpgradeResult) string { return v.DataEngineId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCheckDataEngineImageCanBeUpgradeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCheckDataEngineImageCanBeUpgradeResult) string { return v.Id }).(pulumi.StringOutput)
}

// Is it possible to upgrade.
func (o GetCheckDataEngineImageCanBeUpgradeResultOutput) IsUpgrade() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCheckDataEngineImageCanBeUpgradeResult) bool { return v.IsUpgrade }).(pulumi.BoolOutput)
}

func (o GetCheckDataEngineImageCanBeUpgradeResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCheckDataEngineImageCanBeUpgradeResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCheckDataEngineImageCanBeUpgradeResultOutput{})
}