// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticsearch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query elasticsearch instances.
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
// 		_, err := Elasticsearch.GetInstances(ctx, &elasticsearch.GetInstancesArgs{
// 			InstanceId: pulumi.StringRef("es-17634f05"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetInstancesResult
	err := ctx.Invoke("tencentcloud:Elasticsearch/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// ID of the instance to be queried.
	InstanceId *string `pulumi:"instanceId"`
	// Name of the instance to be queried.
	InstanceName *string `pulumi:"instanceName"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Tag of the instance to be queried.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the instance.
	InstanceId *string `pulumi:"instanceId"`
	// An information list of elasticsearch instance. Each element contains the following attributes:
	InstanceLists []GetInstancesInstanceList `pulumi:"instanceLists"`
	// Name of the instance.
	InstanceName     *string `pulumi:"instanceName"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// A mapping of tags to assign to the instance.
	Tags map[string]interface{} `pulumi:"tags"`
}

func GetInstancesOutput(ctx *pulumi.Context, args GetInstancesOutputArgs, opts ...pulumi.InvokeOption) GetInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstancesResult, error) {
			args := v.(GetInstancesArgs)
			r, err := GetInstances(ctx, &args, opts...)
			var s GetInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstancesResultOutput)
}

// A collection of arguments for invoking getInstances.
type GetInstancesOutputArgs struct {
	// ID of the instance to be queried.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// Name of the instance to be queried.
	InstanceName pulumi.StringPtrInput `pulumi:"instanceName"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Tag of the instance to be queried.
	Tags pulumi.MapInput `pulumi:"tags"`
}

func (GetInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getInstances.
type GetInstancesResultOutput struct{ *pulumi.OutputState }

func (GetInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesResult)(nil)).Elem()
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutput() GetInstancesResultOutput {
	return o
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutputWithContext(ctx context.Context) GetInstancesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the instance.
func (o GetInstancesResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// An information list of elasticsearch instance. Each element contains the following attributes:
func (o GetInstancesResultOutput) InstanceLists() GetInstancesInstanceListArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesInstanceList { return v.InstanceLists }).(GetInstancesInstanceListArrayOutput)
}

// Name of the instance.
func (o GetInstancesResultOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.InstanceName }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// A mapping of tags to assign to the instance.
func (o GetInstancesResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetInstancesResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}
