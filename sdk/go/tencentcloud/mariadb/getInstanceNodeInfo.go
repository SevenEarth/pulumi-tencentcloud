// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetInstanceNodeInfo(ctx *pulumi.Context, args *GetInstanceNodeInfoArgs, opts ...pulumi.InvokeOption) (*GetInstanceNodeInfoResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetInstanceNodeInfoResult
	err := ctx.Invoke("tencentcloud:Mariadb/getInstanceNodeInfo:getInstanceNodeInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceNodeInfo.
type GetInstanceNodeInfoArgs struct {
	InstanceId       string  `pulumi:"instanceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getInstanceNodeInfo.
type GetInstanceNodeInfoResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                         `pulumi:"id"`
	InstanceId       string                         `pulumi:"instanceId"`
	NodesInfos       []GetInstanceNodeInfoNodesInfo `pulumi:"nodesInfos"`
	ResultOutputFile *string                        `pulumi:"resultOutputFile"`
}

func GetInstanceNodeInfoOutput(ctx *pulumi.Context, args GetInstanceNodeInfoOutputArgs, opts ...pulumi.InvokeOption) GetInstanceNodeInfoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceNodeInfoResult, error) {
			args := v.(GetInstanceNodeInfoArgs)
			r, err := GetInstanceNodeInfo(ctx, &args, opts...)
			var s GetInstanceNodeInfoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceNodeInfoResultOutput)
}

// A collection of arguments for invoking getInstanceNodeInfo.
type GetInstanceNodeInfoOutputArgs struct {
	InstanceId       pulumi.StringInput    `pulumi:"instanceId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetInstanceNodeInfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceNodeInfoArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceNodeInfo.
type GetInstanceNodeInfoResultOutput struct{ *pulumi.OutputState }

func (GetInstanceNodeInfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceNodeInfoResult)(nil)).Elem()
}

func (o GetInstanceNodeInfoResultOutput) ToGetInstanceNodeInfoResultOutput() GetInstanceNodeInfoResultOutput {
	return o
}

func (o GetInstanceNodeInfoResultOutput) ToGetInstanceNodeInfoResultOutputWithContext(ctx context.Context) GetInstanceNodeInfoResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceNodeInfoResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceNodeInfoResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstanceNodeInfoResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceNodeInfoResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetInstanceNodeInfoResultOutput) NodesInfos() GetInstanceNodeInfoNodesInfoArrayOutput {
	return o.ApplyT(func(v GetInstanceNodeInfoResult) []GetInstanceNodeInfoNodesInfo { return v.NodesInfos }).(GetInstanceNodeInfoNodesInfoArrayOutput)
}

func (o GetInstanceNodeInfoResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceNodeInfoResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceNodeInfoResultOutput{})
}
