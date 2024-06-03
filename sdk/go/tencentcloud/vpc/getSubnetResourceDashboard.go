// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of vpc subnetResourceDashboard
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vpc.GetSubnetResourceDashboard(ctx, &vpc.GetSubnetResourceDashboardArgs{
//				SubnetIds: []string{
//					"subnet-i9tpf6hq",
//				},
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
func GetSubnetResourceDashboard(ctx *pulumi.Context, args *GetSubnetResourceDashboardArgs, opts ...pulumi.InvokeOption) (*GetSubnetResourceDashboardResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSubnetResourceDashboardResult
	err := ctx.Invoke("tencentcloud:Vpc/getSubnetResourceDashboard:getSubnetResourceDashboard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnetResourceDashboard.
type GetSubnetResourceDashboardArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Subnet instance ID, such as `subnet-f1xjkw1b`.
	SubnetIds []string `pulumi:"subnetIds"`
}

// A collection of values returned by getSubnetResourceDashboard.
type GetSubnetResourceDashboardResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Information of resources returned.
	ResourceStatisticsSets []GetSubnetResourceDashboardResourceStatisticsSet `pulumi:"resourceStatisticsSets"`
	ResultOutputFile       *string                                           `pulumi:"resultOutputFile"`
	SubnetIds              []string                                          `pulumi:"subnetIds"`
}

func GetSubnetResourceDashboardOutput(ctx *pulumi.Context, args GetSubnetResourceDashboardOutputArgs, opts ...pulumi.InvokeOption) GetSubnetResourceDashboardResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSubnetResourceDashboardResult, error) {
			args := v.(GetSubnetResourceDashboardArgs)
			r, err := GetSubnetResourceDashboard(ctx, &args, opts...)
			var s GetSubnetResourceDashboardResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSubnetResourceDashboardResultOutput)
}

// A collection of arguments for invoking getSubnetResourceDashboard.
type GetSubnetResourceDashboardOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Subnet instance ID, such as `subnet-f1xjkw1b`.
	SubnetIds pulumi.StringArrayInput `pulumi:"subnetIds"`
}

func (GetSubnetResourceDashboardOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetResourceDashboardArgs)(nil)).Elem()
}

// A collection of values returned by getSubnetResourceDashboard.
type GetSubnetResourceDashboardResultOutput struct{ *pulumi.OutputState }

func (GetSubnetResourceDashboardResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetResourceDashboardResult)(nil)).Elem()
}

func (o GetSubnetResourceDashboardResultOutput) ToGetSubnetResourceDashboardResultOutput() GetSubnetResourceDashboardResultOutput {
	return o
}

func (o GetSubnetResourceDashboardResultOutput) ToGetSubnetResourceDashboardResultOutputWithContext(ctx context.Context) GetSubnetResourceDashboardResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSubnetResourceDashboardResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetResourceDashboardResult) string { return v.Id }).(pulumi.StringOutput)
}

// Information of resources returned.
func (o GetSubnetResourceDashboardResultOutput) ResourceStatisticsSets() GetSubnetResourceDashboardResourceStatisticsSetArrayOutput {
	return o.ApplyT(func(v GetSubnetResourceDashboardResult) []GetSubnetResourceDashboardResourceStatisticsSet {
		return v.ResourceStatisticsSets
	}).(GetSubnetResourceDashboardResourceStatisticsSetArrayOutput)
}

func (o GetSubnetResourceDashboardResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetResourceDashboardResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetSubnetResourceDashboardResultOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSubnetResourceDashboardResult) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSubnetResourceDashboardResultOutput{})
}
