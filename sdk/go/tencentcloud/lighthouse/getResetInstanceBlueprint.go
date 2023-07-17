// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lighthouse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of lighthouse resetInstanceBlueprint
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Lighthouse"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Lighthouse"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Lighthouse.GetResetInstanceBlueprint(ctx, &lighthouse.GetResetInstanceBlueprintArgs{
// 			InstanceId: "lhins-123456",
// 			Limit:      pulumi.IntRef(20),
// 			Offset:     pulumi.IntRef(0),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetResetInstanceBlueprint(ctx *pulumi.Context, args *GetResetInstanceBlueprintArgs, opts ...pulumi.InvokeOption) (*GetResetInstanceBlueprintResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetResetInstanceBlueprintResult
	err := ctx.Invoke("tencentcloud:Lighthouse/getResetInstanceBlueprint:getResetInstanceBlueprint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResetInstanceBlueprint.
type GetResetInstanceBlueprintArgs struct {
	// Filter listblueprint-idFilter by image ID.Type: StringRequired: noblueprint-typeFilter by image type.Valid values: APP_OS: application image; PURE_OS: system image; PRIVATE: custom imageType: StringRequired: noplatform-typeFilter by image platform type.Valid values: LINUX_UNIX: Linux or Unix; WINDOWS: WindowsType: StringRequired: noblueprint-nameFilter by image name.Type: StringRequired: noblueprint-stateFilter by image status.Type: StringRequired: noEach request can contain up to 10 Filters and 5 Filter.Values. BlueprintIds and Filters cannot be specified at the same time.
	Filters []GetResetInstanceBlueprintFilter `pulumi:"filters"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Number of returned results. Default value is 20. Maximum value is 100.
	Limit *int `pulumi:"limit"`
	// Offset. Default value is 0.
	Offset *int `pulumi:"offset"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getResetInstanceBlueprint.
type GetResetInstanceBlueprintResult struct {
	Filters []GetResetInstanceBlueprintFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	Limit      *int   `pulumi:"limit"`
	Offset     *int   `pulumi:"offset"`
	// List of scene info.
	ResetInstanceBlueprintSets []GetResetInstanceBlueprintResetInstanceBlueprintSet `pulumi:"resetInstanceBlueprintSets"`
	ResultOutputFile           *string                                              `pulumi:"resultOutputFile"`
}

func GetResetInstanceBlueprintOutput(ctx *pulumi.Context, args GetResetInstanceBlueprintOutputArgs, opts ...pulumi.InvokeOption) GetResetInstanceBlueprintResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetResetInstanceBlueprintResult, error) {
			args := v.(GetResetInstanceBlueprintArgs)
			r, err := GetResetInstanceBlueprint(ctx, &args, opts...)
			var s GetResetInstanceBlueprintResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetResetInstanceBlueprintResultOutput)
}

// A collection of arguments for invoking getResetInstanceBlueprint.
type GetResetInstanceBlueprintOutputArgs struct {
	// Filter listblueprint-idFilter by image ID.Type: StringRequired: noblueprint-typeFilter by image type.Valid values: APP_OS: application image; PURE_OS: system image; PRIVATE: custom imageType: StringRequired: noplatform-typeFilter by image platform type.Valid values: LINUX_UNIX: Linux or Unix; WINDOWS: WindowsType: StringRequired: noblueprint-nameFilter by image name.Type: StringRequired: noblueprint-stateFilter by image status.Type: StringRequired: noEach request can contain up to 10 Filters and 5 Filter.Values. BlueprintIds and Filters cannot be specified at the same time.
	Filters GetResetInstanceBlueprintFilterArrayInput `pulumi:"filters"`
	// Instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Number of returned results. Default value is 20. Maximum value is 100.
	Limit pulumi.IntPtrInput `pulumi:"limit"`
	// Offset. Default value is 0.
	Offset pulumi.IntPtrInput `pulumi:"offset"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetResetInstanceBlueprintOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResetInstanceBlueprintArgs)(nil)).Elem()
}

// A collection of values returned by getResetInstanceBlueprint.
type GetResetInstanceBlueprintResultOutput struct{ *pulumi.OutputState }

func (GetResetInstanceBlueprintResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResetInstanceBlueprintResult)(nil)).Elem()
}

func (o GetResetInstanceBlueprintResultOutput) ToGetResetInstanceBlueprintResultOutput() GetResetInstanceBlueprintResultOutput {
	return o
}

func (o GetResetInstanceBlueprintResultOutput) ToGetResetInstanceBlueprintResultOutputWithContext(ctx context.Context) GetResetInstanceBlueprintResultOutput {
	return o
}

func (o GetResetInstanceBlueprintResultOutput) Filters() GetResetInstanceBlueprintFilterArrayOutput {
	return o.ApplyT(func(v GetResetInstanceBlueprintResult) []GetResetInstanceBlueprintFilter { return v.Filters }).(GetResetInstanceBlueprintFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetResetInstanceBlueprintResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetResetInstanceBlueprintResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetResetInstanceBlueprintResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetResetInstanceBlueprintResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetResetInstanceBlueprintResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetResetInstanceBlueprintResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

func (o GetResetInstanceBlueprintResultOutput) Offset() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetResetInstanceBlueprintResult) *int { return v.Offset }).(pulumi.IntPtrOutput)
}

// List of scene info.
func (o GetResetInstanceBlueprintResultOutput) ResetInstanceBlueprintSets() GetResetInstanceBlueprintResetInstanceBlueprintSetArrayOutput {
	return o.ApplyT(func(v GetResetInstanceBlueprintResult) []GetResetInstanceBlueprintResetInstanceBlueprintSet {
		return v.ResetInstanceBlueprintSets
	}).(GetResetInstanceBlueprintResetInstanceBlueprintSetArrayOutput)
}

func (o GetResetInstanceBlueprintResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetResetInstanceBlueprintResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetResetInstanceBlueprintResultOutput{})
}
