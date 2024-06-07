// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lighthouse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of lighthouse zone
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Lighthouse"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Lighthouse.GetZone(ctx, &lighthouse.GetZoneArgs{
//				Order:      pulumi.StringRef("ASC"),
//				OrderField: pulumi.StringRef("ZONE"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetZone(ctx *pulumi.Context, args *GetZoneArgs, opts ...pulumi.InvokeOption) (*GetZoneResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetZoneResult
	err := ctx.Invoke("tencentcloud:Lighthouse/getZone:getZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZone.
type GetZoneArgs struct {
	// Specifies how availability zones are listed. Valid values:
	// - ASC: Ascending sort.
	// - DESC: Descending sort.
	//   The default value is ASC.
	Order *string `pulumi:"order"`
	// Sorting field. Valid values:
	// - ZONE: Sort by the availability zone.
	// - INSTANCE_DISPLAY_LABEL: Sort by visibility labels (HIDDEN, NORMAL and SELECTED). Default: [HIDDEN, NORMAL, SELECTED].
	//   Sort by availability zone by default.
	OrderField *string `pulumi:"orderField"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getZone.
type GetZoneResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	Order            *string `pulumi:"order"`
	OrderField       *string `pulumi:"orderField"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// List of zone info.
	ZoneInfoSets []GetZoneZoneInfoSet `pulumi:"zoneInfoSets"`
}

func GetZoneOutput(ctx *pulumi.Context, args GetZoneOutputArgs, opts ...pulumi.InvokeOption) GetZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetZoneResult, error) {
			args := v.(GetZoneArgs)
			r, err := GetZone(ctx, &args, opts...)
			var s GetZoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetZoneResultOutput)
}

// A collection of arguments for invoking getZone.
type GetZoneOutputArgs struct {
	// Specifies how availability zones are listed. Valid values:
	// - ASC: Ascending sort.
	// - DESC: Descending sort.
	//   The default value is ASC.
	Order pulumi.StringPtrInput `pulumi:"order"`
	// Sorting field. Valid values:
	// - ZONE: Sort by the availability zone.
	// - INSTANCE_DISPLAY_LABEL: Sort by visibility labels (HIDDEN, NORMAL and SELECTED). Default: [HIDDEN, NORMAL, SELECTED].
	//   Sort by availability zone by default.
	OrderField pulumi.StringPtrInput `pulumi:"orderField"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZoneArgs)(nil)).Elem()
}

// A collection of values returned by getZone.
type GetZoneResultOutput struct{ *pulumi.OutputState }

func (GetZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZoneResult)(nil)).Elem()
}

func (o GetZoneResultOutput) ToGetZoneResultOutput() GetZoneResultOutput {
	return o
}

func (o GetZoneResultOutput) ToGetZoneResultOutputWithContext(ctx context.Context) GetZoneResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetZoneResultOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZoneResult) *string { return v.Order }).(pulumi.StringPtrOutput)
}

func (o GetZoneResultOutput) OrderField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZoneResult) *string { return v.OrderField }).(pulumi.StringPtrOutput)
}

func (o GetZoneResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZoneResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// List of zone info.
func (o GetZoneResultOutput) ZoneInfoSets() GetZoneZoneInfoSetArrayOutput {
	return o.ApplyT(func(v GetZoneResult) []GetZoneZoneInfoSet { return v.ZoneInfoSets }).(GetZoneZoneInfoSetArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetZoneResultOutput{})
}
