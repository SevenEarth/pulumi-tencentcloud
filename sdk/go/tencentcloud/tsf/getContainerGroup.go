// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tsf containerGroup
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tsf.GetContainerGroup(ctx, &tsf.GetContainerGroupArgs{
//				ApplicationId: pulumi.StringRef("application-a24x29xv"),
//				ClusterId:     pulumi.StringRef("cluster-vwgj5e6y"),
//				NamespaceId:   pulumi.StringRef("namespace-aemrg36v"),
//				OrderBy:       pulumi.StringRef("createTime"),
//				OrderType:     pulumi.IntRef(0),
//				SearchWord:    pulumi.StringRef("keep"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupContainerGroup(ctx *pulumi.Context, args *LookupContainerGroupArgs, opts ...pulumi.InvokeOption) (*LookupContainerGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupContainerGroupResult
	err := ctx.Invoke("tencentcloud:Tsf/getContainerGroup:getContainerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerGroup.
type LookupContainerGroupArgs struct {
	// ApplicationId, required.
	ApplicationId *string `pulumi:"applicationId"`
	// Cluster Id.
	ClusterId *string `pulumi:"clusterId"`
	// Namespace Id.
	NamespaceId *string `pulumi:"namespaceId"`
	// The sorting field. By default, it is the createTime field. Supports id, name, createTime.
	OrderBy *string `pulumi:"orderBy"`
	// The sorting order. By default, it is 1, indicating descending order. 0 indicates ascending order, and 1 indicates descending order.
	OrderType *int `pulumi:"orderType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// search word, support group name.
	SearchWord *string `pulumi:"searchWord"`
}

// A collection of values returned by getContainerGroup.
type LookupContainerGroupResult struct {
	ApplicationId *string `pulumi:"applicationId"`
	// Cluster Id.Note: This field may return null, indicating that no valid value was found.
	ClusterId *string `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Namespace Id.Note: This field may return null, indicating that no valid value was found.
	NamespaceId      *string `pulumi:"namespaceId"`
	OrderBy          *string `pulumi:"orderBy"`
	OrderType        *int    `pulumi:"orderType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// result list.
	Results    []GetContainerGroupResult `pulumi:"results"`
	SearchWord *string                   `pulumi:"searchWord"`
}

func LookupContainerGroupOutput(ctx *pulumi.Context, args LookupContainerGroupOutputArgs, opts ...pulumi.InvokeOption) LookupContainerGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerGroupResult, error) {
			args := v.(LookupContainerGroupArgs)
			r, err := LookupContainerGroup(ctx, &args, opts...)
			var s LookupContainerGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerGroupResultOutput)
}

// A collection of arguments for invoking getContainerGroup.
type LookupContainerGroupOutputArgs struct {
	// ApplicationId, required.
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
	// Cluster Id.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// Namespace Id.
	NamespaceId pulumi.StringPtrInput `pulumi:"namespaceId"`
	// The sorting field. By default, it is the createTime field. Supports id, name, createTime.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// The sorting order. By default, it is 1, indicating descending order. 0 indicates ascending order, and 1 indicates descending order.
	OrderType pulumi.IntPtrInput `pulumi:"orderType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// search word, support group name.
	SearchWord pulumi.StringPtrInput `pulumi:"searchWord"`
}

func (LookupContainerGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerGroupArgs)(nil)).Elem()
}

// A collection of values returned by getContainerGroup.
type LookupContainerGroupResultOutput struct{ *pulumi.OutputState }

func (LookupContainerGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerGroupResult)(nil)).Elem()
}

func (o LookupContainerGroupResultOutput) ToLookupContainerGroupResultOutput() LookupContainerGroupResultOutput {
	return o
}

func (o LookupContainerGroupResultOutput) ToLookupContainerGroupResultOutputWithContext(ctx context.Context) LookupContainerGroupResultOutput {
	return o
}

func (o LookupContainerGroupResultOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

// Cluster Id.Note: This field may return null, indicating that no valid value was found.
func (o LookupContainerGroupResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupContainerGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Namespace Id.Note: This field may return null, indicating that no valid value was found.
func (o LookupContainerGroupResultOutput) NamespaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.NamespaceId }).(pulumi.StringPtrOutput)
}

func (o LookupContainerGroupResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o LookupContainerGroupResultOutput) OrderType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *int { return v.OrderType }).(pulumi.IntPtrOutput)
}

func (o LookupContainerGroupResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// result list.
func (o LookupContainerGroupResultOutput) Results() GetContainerGroupResultArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []GetContainerGroupResult { return v.Results }).(GetContainerGroupResultArrayOutput)
}

func (o LookupContainerGroupResultOutput) SearchWord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.SearchWord }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerGroupResultOutput{})
}
