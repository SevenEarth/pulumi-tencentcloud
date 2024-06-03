// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of vpc cvmInstances
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
//			_, err := Vpc.GetCvmInstances(ctx, &vpc.GetCvmInstancesArgs{
//				Filters: []vpc.GetCvmInstancesFilter{
//					{
//						Name: "vpc-id",
//						Values: []string{
//							"vpc-lh4nqig9",
//						},
//					},
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
func GetCvmInstances(ctx *pulumi.Context, args *GetCvmInstancesArgs, opts ...pulumi.InvokeOption) (*GetCvmInstancesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCvmInstancesResult
	err := ctx.Invoke("tencentcloud:Vpc/getCvmInstances:getCvmInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCvmInstances.
type GetCvmInstancesArgs struct {
	// Filter condition. `RouteTableIds` and `Filters` cannot be specified at the same time. vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`;instance-type - String - (Filter condition) CVM instance ID;instance-name - String - (Filter condition) CVM name.
	Filters []GetCvmInstancesFilter `pulumi:"filters"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getCvmInstances.
type GetCvmInstancesResult struct {
	Filters []GetCvmInstancesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of CVM instances.
	InstanceSets     []GetCvmInstancesInstanceSet `pulumi:"instanceSets"`
	ResultOutputFile *string                      `pulumi:"resultOutputFile"`
}

func GetCvmInstancesOutput(ctx *pulumi.Context, args GetCvmInstancesOutputArgs, opts ...pulumi.InvokeOption) GetCvmInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCvmInstancesResult, error) {
			args := v.(GetCvmInstancesArgs)
			r, err := GetCvmInstances(ctx, &args, opts...)
			var s GetCvmInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCvmInstancesResultOutput)
}

// A collection of arguments for invoking getCvmInstances.
type GetCvmInstancesOutputArgs struct {
	// Filter condition. `RouteTableIds` and `Filters` cannot be specified at the same time. vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`;instance-type - String - (Filter condition) CVM instance ID;instance-name - String - (Filter condition) CVM name.
	Filters GetCvmInstancesFilterArrayInput `pulumi:"filters"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetCvmInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCvmInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getCvmInstances.
type GetCvmInstancesResultOutput struct{ *pulumi.OutputState }

func (GetCvmInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCvmInstancesResult)(nil)).Elem()
}

func (o GetCvmInstancesResultOutput) ToGetCvmInstancesResultOutput() GetCvmInstancesResultOutput {
	return o
}

func (o GetCvmInstancesResultOutput) ToGetCvmInstancesResultOutputWithContext(ctx context.Context) GetCvmInstancesResultOutput {
	return o
}

func (o GetCvmInstancesResultOutput) Filters() GetCvmInstancesFilterArrayOutput {
	return o.ApplyT(func(v GetCvmInstancesResult) []GetCvmInstancesFilter { return v.Filters }).(GetCvmInstancesFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCvmInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCvmInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of CVM instances.
func (o GetCvmInstancesResultOutput) InstanceSets() GetCvmInstancesInstanceSetArrayOutput {
	return o.ApplyT(func(v GetCvmInstancesResult) []GetCvmInstancesInstanceSet { return v.InstanceSets }).(GetCvmInstancesInstanceSetArrayOutput)
}

func (o GetCvmInstancesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCvmInstancesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCvmInstancesResultOutput{})
}
