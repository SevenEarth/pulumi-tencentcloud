// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticsearch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of elasticsearch instance plugin list
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Elasticsearch"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Elasticsearch"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Elasticsearch.GetInstancePluginList(ctx, &elasticsearch.GetInstancePluginListArgs{
// 			InstanceId: "es-xxxxxx",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetInstancePluginList(ctx *pulumi.Context, args *GetInstancePluginListArgs, opts ...pulumi.InvokeOption) (*GetInstancePluginListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetInstancePluginListResult
	err := ctx.Invoke("tencentcloud:Elasticsearch/getInstancePluginList:getInstancePluginList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstancePluginList.
type GetInstancePluginListArgs struct {
	// Instance id.
	InstanceId string `pulumi:"instanceId"`
	// order field. Valid values: `pluginName`.
	OrderBy *string `pulumi:"orderBy"`
	// Order type. Valid values:
	// - asc: Ascending asc
	// - desc: Descending Desc.
	OrderByType *string `pulumi:"orderByType"`
	// Plugin type. Valid values: `0`: System plugin.
	PluginType *int `pulumi:"pluginType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getInstancePluginList.
type GetInstancePluginListResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	InstanceId  string  `pulumi:"instanceId"`
	OrderBy     *string `pulumi:"orderBy"`
	OrderByType *string `pulumi:"orderByType"`
	// Plugin information list.
	PluginLists []GetInstancePluginListPluginList `pulumi:"pluginLists"`
	// Plugin type. Valid values: `0`: System plugin.
	PluginType       *int    `pulumi:"pluginType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetInstancePluginListOutput(ctx *pulumi.Context, args GetInstancePluginListOutputArgs, opts ...pulumi.InvokeOption) GetInstancePluginListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstancePluginListResult, error) {
			args := v.(GetInstancePluginListArgs)
			r, err := GetInstancePluginList(ctx, &args, opts...)
			var s GetInstancePluginListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstancePluginListResultOutput)
}

// A collection of arguments for invoking getInstancePluginList.
type GetInstancePluginListOutputArgs struct {
	// Instance id.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// order field. Valid values: `pluginName`.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// Order type. Valid values:
	// - asc: Ascending asc
	// - desc: Descending Desc.
	OrderByType pulumi.StringPtrInput `pulumi:"orderByType"`
	// Plugin type. Valid values: `0`: System plugin.
	PluginType pulumi.IntPtrInput `pulumi:"pluginType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetInstancePluginListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancePluginListArgs)(nil)).Elem()
}

// A collection of values returned by getInstancePluginList.
type GetInstancePluginListResultOutput struct{ *pulumi.OutputState }

func (GetInstancePluginListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancePluginListResult)(nil)).Elem()
}

func (o GetInstancePluginListResultOutput) ToGetInstancePluginListResultOutput() GetInstancePluginListResultOutput {
	return o
}

func (o GetInstancePluginListResultOutput) ToGetInstancePluginListResultOutputWithContext(ctx context.Context) GetInstancePluginListResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstancePluginListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancePluginListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstancePluginListResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancePluginListResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetInstancePluginListResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancePluginListResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o GetInstancePluginListResultOutput) OrderByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancePluginListResult) *string { return v.OrderByType }).(pulumi.StringPtrOutput)
}

// Plugin information list.
func (o GetInstancePluginListResultOutput) PluginLists() GetInstancePluginListPluginListArrayOutput {
	return o.ApplyT(func(v GetInstancePluginListResult) []GetInstancePluginListPluginList { return v.PluginLists }).(GetInstancePluginListPluginListArrayOutput)
}

// Plugin type. Valid values: `0`: System plugin.
func (o GetInstancePluginListResultOutput) PluginType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstancePluginListResult) *int { return v.PluginType }).(pulumi.IntPtrOutput)
}

func (o GetInstancePluginListResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancePluginListResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancePluginListResultOutput{})
}