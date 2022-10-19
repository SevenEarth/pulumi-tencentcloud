// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of CLB
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Clb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Clb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Clb.GetInstances(ctx, &clb.GetInstancesArgs{
//				ClbId:            pulumi.StringRef("lb-k2zjp9lv"),
//				ClbName:          pulumi.StringRef("myclb"),
//				NetworkType:      pulumi.StringRef("OPEN"),
//				ProjectId:        pulumi.IntRef(0),
//				ResultOutputFile: pulumi.StringRef("mytestpath"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("tencentcloud:Clb/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// ID of the CLB to be queried.
	ClbId *string `pulumi:"clbId"`
	// Name of the CLB to be queried.
	ClbName *string `pulumi:"clbName"`
	// Master available zone id.
	MasterZone *string `pulumi:"masterZone"`
	// Type of CLB instance, and available values include `OPEN` and `INTERNAL`.
	NetworkType *string `pulumi:"networkType"`
	// Project ID of the CLB.
	ProjectId *int `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// ID of CLB.
	ClbId *string `pulumi:"clbId"`
	// A list of cloud load balancers. Each element contains the following attributes:
	ClbLists []GetInstancesClbList `pulumi:"clbLists"`
	// Name of CLB.
	ClbName *string `pulumi:"clbName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	MasterZone *string `pulumi:"masterZone"`
	// Types of CLB.
	NetworkType *string `pulumi:"networkType"`
	// ID of the project.
	ProjectId        *int    `pulumi:"projectId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
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
	// ID of the CLB to be queried.
	ClbId pulumi.StringPtrInput `pulumi:"clbId"`
	// Name of the CLB to be queried.
	ClbName pulumi.StringPtrInput `pulumi:"clbName"`
	// Master available zone id.
	MasterZone pulumi.StringPtrInput `pulumi:"masterZone"`
	// Type of CLB instance, and available values include `OPEN` and `INTERNAL`.
	NetworkType pulumi.StringPtrInput `pulumi:"networkType"`
	// Project ID of the CLB.
	ProjectId pulumi.IntPtrInput `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
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

// ID of CLB.
func (o GetInstancesResultOutput) ClbId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.ClbId }).(pulumi.StringPtrOutput)
}

// A list of cloud load balancers. Each element contains the following attributes:
func (o GetInstancesResultOutput) ClbLists() GetInstancesClbListArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesClbList { return v.ClbLists }).(GetInstancesClbListArrayOutput)
}

// Name of CLB.
func (o GetInstancesResultOutput) ClbName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.ClbName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstancesResultOutput) MasterZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.MasterZone }).(pulumi.StringPtrOutput)
}

// Types of CLB.
func (o GetInstancesResultOutput) NetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.NetworkType }).(pulumi.StringPtrOutput)
}

// ID of the project.
func (o GetInstancesResultOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *int { return v.ProjectId }).(pulumi.IntPtrOutput)
}

func (o GetInstancesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}
