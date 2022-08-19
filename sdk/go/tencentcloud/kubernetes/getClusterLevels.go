// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a datasource to query TKE cluster levels.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Kubernetes"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Kubernetes"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foo, err := Kubernetes.GetClusterLevels(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("level5", foo.Lists[0].Alias)
// 		return nil
// 	})
// }
// ```
func GetClusterLevels(ctx *pulumi.Context, args *GetClusterLevelsArgs, opts ...pulumi.InvokeOption) (*GetClusterLevelsResult, error) {
	var rv GetClusterLevelsResult
	err := ctx.Invoke("tencentcloud:Kubernetes/getClusterLevels:getClusterLevels", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusterLevels.
type GetClusterLevelsArgs struct {
	// Specify cluster Id, if set will only query current cluster's available levels.
	ClusterId *string `pulumi:"clusterId"`
	// Used for save result.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getClusterLevels.
type GetClusterLevelsResult struct {
	ClusterId *string `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of level information.
	Lists            []GetClusterLevelsList `pulumi:"lists"`
	ResultOutputFile *string                `pulumi:"resultOutputFile"`
}

func GetClusterLevelsOutput(ctx *pulumi.Context, args GetClusterLevelsOutputArgs, opts ...pulumi.InvokeOption) GetClusterLevelsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClusterLevelsResult, error) {
			args := v.(GetClusterLevelsArgs)
			r, err := GetClusterLevels(ctx, &args, opts...)
			var s GetClusterLevelsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetClusterLevelsResultOutput)
}

// A collection of arguments for invoking getClusterLevels.
type GetClusterLevelsOutputArgs struct {
	// Specify cluster Id, if set will only query current cluster's available levels.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// Used for save result.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetClusterLevelsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterLevelsArgs)(nil)).Elem()
}

// A collection of values returned by getClusterLevels.
type GetClusterLevelsResultOutput struct{ *pulumi.OutputState }

func (GetClusterLevelsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterLevelsResult)(nil)).Elem()
}

func (o GetClusterLevelsResultOutput) ToGetClusterLevelsResultOutput() GetClusterLevelsResultOutput {
	return o
}

func (o GetClusterLevelsResultOutput) ToGetClusterLevelsResultOutputWithContext(ctx context.Context) GetClusterLevelsResultOutput {
	return o
}

func (o GetClusterLevelsResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClusterLevelsResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClusterLevelsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterLevelsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of level information.
func (o GetClusterLevelsResultOutput) Lists() GetClusterLevelsListArrayOutput {
	return o.ApplyT(func(v GetClusterLevelsResult) []GetClusterLevelsList { return v.Lists }).(GetClusterLevelsListArrayOutput)
}

func (o GetClusterLevelsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClusterLevelsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClusterLevelsResultOutput{})
}
