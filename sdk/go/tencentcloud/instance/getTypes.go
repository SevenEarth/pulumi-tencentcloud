// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package instance

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query instances type.
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Instance"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Instance.GetTypes(ctx, &instance.GetTypesArgs{
//				AvailabilityZone: pulumi.StringRef("ap-guangzhou-2"),
//				CpuCoreCount:     pulumi.IntRef(2),
//				MemorySize:       pulumi.IntRef(4),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Instance.GetTypes(ctx, &instance.GetTypesArgs{
//				CpuCoreCount:   pulumi.IntRef(1),
//				ExcludeSoldOut: pulumi.BoolRef(true),
//				Filters: []instance.GetTypesFilter{
//					{
//						Name: "instance-charge-type",
//						Values: []string{
//							"POSTPAID_BY_HOUR",
//						},
//					},
//					{
//						Name: "zone",
//						Values: []string{
//							"ap-shanghai-2",
//						},
//					},
//				},
//				MemorySize: pulumi.IntRef(1),
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
func GetTypes(ctx *pulumi.Context, args *GetTypesArgs, opts ...pulumi.InvokeOption) (*GetTypesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTypesResult
	err := ctx.Invoke("tencentcloud:Instance/getTypes:getTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTypes.
type GetTypesArgs struct {
	// The available zone that the CVM instance locates at. This field is conflict with `filter`.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The number of CPU cores of the instance.
	CpuCoreCount *int `pulumi:"cpuCoreCount"`
	// Indicate to filter instances types that is sold out or not, default is false.
	ExcludeSoldOut *bool `pulumi:"excludeSoldOut"`
	// One or more name/value pairs to filter. This field is conflict with `availabilityZone`.
	Filters []GetTypesFilter `pulumi:"filters"`
	// The number of GPU cores of the instance.
	GpuCoreCount *int `pulumi:"gpuCoreCount"`
	// Instance memory capacity, unit in GB.
	MemorySize *int `pulumi:"memorySize"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getTypes.
type GetTypesResult struct {
	// The available zone that the CVM instance locates at.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The number of CPU cores of the instance.
	CpuCoreCount   *int             `pulumi:"cpuCoreCount"`
	ExcludeSoldOut *bool            `pulumi:"excludeSoldOut"`
	Filters        []GetTypesFilter `pulumi:"filters"`
	// The number of GPU cores of the instance.
	GpuCoreCount *int `pulumi:"gpuCoreCount"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An information list of cvm instance. Each element contains the following attributes:
	InstanceTypes []GetTypesInstanceType `pulumi:"instanceTypes"`
	// Instance memory capacity, unit in GB.
	MemorySize       *int    `pulumi:"memorySize"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetTypesOutput(ctx *pulumi.Context, args GetTypesOutputArgs, opts ...pulumi.InvokeOption) GetTypesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTypesResult, error) {
			args := v.(GetTypesArgs)
			r, err := GetTypes(ctx, &args, opts...)
			var s GetTypesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTypesResultOutput)
}

// A collection of arguments for invoking getTypes.
type GetTypesOutputArgs struct {
	// The available zone that the CVM instance locates at. This field is conflict with `filter`.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	// The number of CPU cores of the instance.
	CpuCoreCount pulumi.IntPtrInput `pulumi:"cpuCoreCount"`
	// Indicate to filter instances types that is sold out or not, default is false.
	ExcludeSoldOut pulumi.BoolPtrInput `pulumi:"excludeSoldOut"`
	// One or more name/value pairs to filter. This field is conflict with `availabilityZone`.
	Filters GetTypesFilterArrayInput `pulumi:"filters"`
	// The number of GPU cores of the instance.
	GpuCoreCount pulumi.IntPtrInput `pulumi:"gpuCoreCount"`
	// Instance memory capacity, unit in GB.
	MemorySize pulumi.IntPtrInput `pulumi:"memorySize"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetTypesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTypesArgs)(nil)).Elem()
}

// A collection of values returned by getTypes.
type GetTypesResultOutput struct{ *pulumi.OutputState }

func (GetTypesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTypesResult)(nil)).Elem()
}

func (o GetTypesResultOutput) ToGetTypesResultOutput() GetTypesResultOutput {
	return o
}

func (o GetTypesResultOutput) ToGetTypesResultOutputWithContext(ctx context.Context) GetTypesResultOutput {
	return o
}

// The available zone that the CVM instance locates at.
func (o GetTypesResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTypesResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// The number of CPU cores of the instance.
func (o GetTypesResultOutput) CpuCoreCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetTypesResult) *int { return v.CpuCoreCount }).(pulumi.IntPtrOutput)
}

func (o GetTypesResultOutput) ExcludeSoldOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetTypesResult) *bool { return v.ExcludeSoldOut }).(pulumi.BoolPtrOutput)
}

func (o GetTypesResultOutput) Filters() GetTypesFilterArrayOutput {
	return o.ApplyT(func(v GetTypesResult) []GetTypesFilter { return v.Filters }).(GetTypesFilterArrayOutput)
}

// The number of GPU cores of the instance.
func (o GetTypesResultOutput) GpuCoreCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetTypesResult) *int { return v.GpuCoreCount }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTypesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTypesResult) string { return v.Id }).(pulumi.StringOutput)
}

// An information list of cvm instance. Each element contains the following attributes:
func (o GetTypesResultOutput) InstanceTypes() GetTypesInstanceTypeArrayOutput {
	return o.ApplyT(func(v GetTypesResult) []GetTypesInstanceType { return v.InstanceTypes }).(GetTypesInstanceTypeArrayOutput)
}

// Instance memory capacity, unit in GB.
func (o GetTypesResultOutput) MemorySize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetTypesResult) *int { return v.MemorySize }).(pulumi.IntPtrOutput)
}

func (o GetTypesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTypesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTypesResultOutput{})
}
