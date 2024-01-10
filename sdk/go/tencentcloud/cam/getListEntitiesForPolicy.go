// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of cam listEntitiesForPolicy
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cam"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cam.GetListEntitiesForPolicy(ctx, &cam.GetListEntitiesForPolicyArgs{
//				EntityFilter: pulumi.StringRef("All"),
//				PolicyId:     1,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetListEntitiesForPolicy(ctx *pulumi.Context, args *GetListEntitiesForPolicyArgs, opts ...pulumi.InvokeOption) (*GetListEntitiesForPolicyResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetListEntitiesForPolicyResult
	err := ctx.Invoke("tencentcloud:Cam/getListEntitiesForPolicy:getListEntitiesForPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getListEntitiesForPolicy.
type GetListEntitiesForPolicyArgs struct {
	// Can take values of &amp;amp;#39;All&amp;amp;#39;, &amp;amp;#39;User&amp;amp;#39;, &amp;amp;#39;Group&amp;amp;#39;, and &amp;amp;#39;Role&amp;amp;#39;. &amp;amp;#39;All&amp;amp;#39; represents obtaining all entity types, &amp;amp;#39;User&amp;amp;#39; represents only obtaining sub accounts, &amp;amp;#39;Group&amp;amp;#39; represents only obtaining user groups, and &amp;amp;#39;Role&amp;amp;#39; represents only obtaining roles. The default value is&amp;amp;#39; All &amp;amp;#39;.
	EntityFilter *string `pulumi:"entityFilter"`
	// Policy Id.
	PolicyId int `pulumi:"policyId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Per page size, default value is 20.
	Rp *int `pulumi:"rp"`
}

// A collection of values returned by getListEntitiesForPolicy.
type GetListEntitiesForPolicyResult struct {
	EntityFilter *string `pulumi:"entityFilter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Entity ListNote: This field may return null, indicating that a valid value cannot be obtained.
	Lists            []GetListEntitiesForPolicyList `pulumi:"lists"`
	PolicyId         int                            `pulumi:"policyId"`
	ResultOutputFile *string                        `pulumi:"resultOutputFile"`
	Rp               *int                           `pulumi:"rp"`
}

func GetListEntitiesForPolicyOutput(ctx *pulumi.Context, args GetListEntitiesForPolicyOutputArgs, opts ...pulumi.InvokeOption) GetListEntitiesForPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetListEntitiesForPolicyResult, error) {
			args := v.(GetListEntitiesForPolicyArgs)
			r, err := GetListEntitiesForPolicy(ctx, &args, opts...)
			var s GetListEntitiesForPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetListEntitiesForPolicyResultOutput)
}

// A collection of arguments for invoking getListEntitiesForPolicy.
type GetListEntitiesForPolicyOutputArgs struct {
	// Can take values of &amp;amp;#39;All&amp;amp;#39;, &amp;amp;#39;User&amp;amp;#39;, &amp;amp;#39;Group&amp;amp;#39;, and &amp;amp;#39;Role&amp;amp;#39;. &amp;amp;#39;All&amp;amp;#39; represents obtaining all entity types, &amp;amp;#39;User&amp;amp;#39; represents only obtaining sub accounts, &amp;amp;#39;Group&amp;amp;#39; represents only obtaining user groups, and &amp;amp;#39;Role&amp;amp;#39; represents only obtaining roles. The default value is&amp;amp;#39; All &amp;amp;#39;.
	EntityFilter pulumi.StringPtrInput `pulumi:"entityFilter"`
	// Policy Id.
	PolicyId pulumi.IntInput `pulumi:"policyId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Per page size, default value is 20.
	Rp pulumi.IntPtrInput `pulumi:"rp"`
}

func (GetListEntitiesForPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetListEntitiesForPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getListEntitiesForPolicy.
type GetListEntitiesForPolicyResultOutput struct{ *pulumi.OutputState }

func (GetListEntitiesForPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetListEntitiesForPolicyResult)(nil)).Elem()
}

func (o GetListEntitiesForPolicyResultOutput) ToGetListEntitiesForPolicyResultOutput() GetListEntitiesForPolicyResultOutput {
	return o
}

func (o GetListEntitiesForPolicyResultOutput) ToGetListEntitiesForPolicyResultOutputWithContext(ctx context.Context) GetListEntitiesForPolicyResultOutput {
	return o
}

func (o GetListEntitiesForPolicyResultOutput) EntityFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListEntitiesForPolicyResult) *string { return v.EntityFilter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetListEntitiesForPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetListEntitiesForPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Entity ListNote: This field may return null, indicating that a valid value cannot be obtained.
func (o GetListEntitiesForPolicyResultOutput) Lists() GetListEntitiesForPolicyListArrayOutput {
	return o.ApplyT(func(v GetListEntitiesForPolicyResult) []GetListEntitiesForPolicyList { return v.Lists }).(GetListEntitiesForPolicyListArrayOutput)
}

func (o GetListEntitiesForPolicyResultOutput) PolicyId() pulumi.IntOutput {
	return o.ApplyT(func(v GetListEntitiesForPolicyResult) int { return v.PolicyId }).(pulumi.IntOutput)
}

func (o GetListEntitiesForPolicyResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListEntitiesForPolicyResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetListEntitiesForPolicyResultOutput) Rp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetListEntitiesForPolicyResult) *int { return v.Rp }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetListEntitiesForPolicyResultOutput{})
}
