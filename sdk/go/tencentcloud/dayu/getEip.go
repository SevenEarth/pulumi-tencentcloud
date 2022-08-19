// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query dayu eip rules
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dayu"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dayu"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dayu.GetEip(ctx, &dayu.GetEipArgs{
// 			ResourceId: "bgpip-000004xg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupEip(ctx *pulumi.Context, args *LookupEipArgs, opts ...pulumi.InvokeOption) (*LookupEipResult, error) {
	var rv LookupEipResult
	err := ctx.Invoke("tencentcloud:Dayu/getEip:getEip", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEip.
type LookupEipArgs struct {
	// The binding state of the instance, value range [BINDING, BIND, UNBINDING, UNBIND], default is [BINDING, BIND, UNBINDING, UNBIND].
	BindStatuses []string `pulumi:"bindStatuses"`
	// The number of pages, default is `10`.
	Limit *int `pulumi:"limit"`
	// The page start offset, default is `0`.
	Offset *int `pulumi:"offset"`
	// Id of the resource.
	ResourceId string `pulumi:"resourceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getEip.
type LookupEipResult struct {
	BindStatuses []string `pulumi:"bindStatuses"`
	// The provider-assigned unique ID for this managed resource.
	Id    string `pulumi:"id"`
	Limit *int   `pulumi:"limit"`
	// A list of layer 4 rules. Each element contains the following attributes:
	Lists            []GetEipList `pulumi:"lists"`
	Offset           *int         `pulumi:"offset"`
	ResourceId       string       `pulumi:"resourceId"`
	ResultOutputFile *string      `pulumi:"resultOutputFile"`
}

func LookupEipOutput(ctx *pulumi.Context, args LookupEipOutputArgs, opts ...pulumi.InvokeOption) LookupEipResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEipResult, error) {
			args := v.(LookupEipArgs)
			r, err := LookupEip(ctx, &args, opts...)
			var s LookupEipResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEipResultOutput)
}

// A collection of arguments for invoking getEip.
type LookupEipOutputArgs struct {
	// The binding state of the instance, value range [BINDING, BIND, UNBINDING, UNBIND], default is [BINDING, BIND, UNBINDING, UNBIND].
	BindStatuses pulumi.StringArrayInput `pulumi:"bindStatuses"`
	// The number of pages, default is `10`.
	Limit pulumi.IntPtrInput `pulumi:"limit"`
	// The page start offset, default is `0`.
	Offset pulumi.IntPtrInput `pulumi:"offset"`
	// Id of the resource.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (LookupEipOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEipArgs)(nil)).Elem()
}

// A collection of values returned by getEip.
type LookupEipResultOutput struct{ *pulumi.OutputState }

func (LookupEipResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEipResult)(nil)).Elem()
}

func (o LookupEipResultOutput) ToLookupEipResultOutput() LookupEipResultOutput {
	return o
}

func (o LookupEipResultOutput) ToLookupEipResultOutputWithContext(ctx context.Context) LookupEipResultOutput {
	return o
}

func (o LookupEipResultOutput) BindStatuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEipResult) []string { return v.BindStatuses }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupEipResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEipResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEipResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEipResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

// A list of layer 4 rules. Each element contains the following attributes:
func (o LookupEipResultOutput) Lists() GetEipListArrayOutput {
	return o.ApplyT(func(v LookupEipResult) []GetEipList { return v.Lists }).(GetEipListArrayOutput)
}

func (o LookupEipResultOutput) Offset() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEipResult) *int { return v.Offset }).(pulumi.IntPtrOutput)
}

func (o LookupEipResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEipResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o LookupEipResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEipResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEipResultOutput{})
}
