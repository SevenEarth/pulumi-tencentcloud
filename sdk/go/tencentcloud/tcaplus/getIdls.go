// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcaplus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query  IDL information of the TcaplusDB table.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tcaplus"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tcaplus"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Tcaplus.GetIdls(ctx, &tcaplus.GetIdlsArgs{
// 			ClusterId: "19162256624",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetIdls(ctx *pulumi.Context, args *GetIdlsArgs, opts ...pulumi.InvokeOption) (*GetIdlsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetIdlsResult
	err := ctx.Invoke("tencentcloud:Tcaplus/getIdls:getIdls", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIdls.
type GetIdlsArgs struct {
	// ID of the TcaplusDB cluster to be query.
	ClusterId string `pulumi:"clusterId"`
	// File for saving results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getIdls.
type GetIdlsResult struct {
	ClusterId string `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of TcaplusDB table IDL. Each element contains the following attributes.
	Lists            []GetIdlsList `pulumi:"lists"`
	ResultOutputFile *string       `pulumi:"resultOutputFile"`
}

func GetIdlsOutput(ctx *pulumi.Context, args GetIdlsOutputArgs, opts ...pulumi.InvokeOption) GetIdlsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIdlsResult, error) {
			args := v.(GetIdlsArgs)
			r, err := GetIdls(ctx, &args, opts...)
			var s GetIdlsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIdlsResultOutput)
}

// A collection of arguments for invoking getIdls.
type GetIdlsOutputArgs struct {
	// ID of the TcaplusDB cluster to be query.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// File for saving results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetIdlsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIdlsArgs)(nil)).Elem()
}

// A collection of values returned by getIdls.
type GetIdlsResultOutput struct{ *pulumi.OutputState }

func (GetIdlsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIdlsResult)(nil)).Elem()
}

func (o GetIdlsResultOutput) ToGetIdlsResultOutput() GetIdlsResultOutput {
	return o
}

func (o GetIdlsResultOutput) ToGetIdlsResultOutputWithContext(ctx context.Context) GetIdlsResultOutput {
	return o
}

func (o GetIdlsResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdlsResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIdlsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdlsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of TcaplusDB table IDL. Each element contains the following attributes.
func (o GetIdlsResultOutput) Lists() GetIdlsListArrayOutput {
	return o.ApplyT(func(v GetIdlsResult) []GetIdlsList { return v.Lists }).(GetIdlsListArrayOutput)
}

func (o GetIdlsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIdlsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIdlsResultOutput{})
}
