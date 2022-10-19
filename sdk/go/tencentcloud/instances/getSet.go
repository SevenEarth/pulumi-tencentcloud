// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package instances

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query cvm instances in parallel.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Instances"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Instances"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Instances.GetSet(ctx, &instances.GetSetArgs{
//				VpcId: pulumi.StringRef("vpc-4owdpnwr"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSet(ctx *pulumi.Context, args *GetSetArgs, opts ...pulumi.InvokeOption) (*GetSetResult, error) {
	var rv GetSetResult
	err := ctx.Invoke("tencentcloud:Instances/getSet:getSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSet.
type GetSetArgs struct {
	// The available zone that the CVM instance locates at.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// ID of the instances to be queried.
	InstanceId *string `pulumi:"instanceId"`
	// Name of the instances to be queried.
	InstanceName *string `pulumi:"instanceName"`
	// The project CVM belongs to.
	ProjectId *int `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of a vpc subnetwork.
	SubnetId *string `pulumi:"subnetId"`
	// Tags of the instance.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the vpc to be queried.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getSet.
type GetSetResult struct {
	// The available zone that the CVM instance locates at.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the instances.
	InstanceId *string `pulumi:"instanceId"`
	// An information list of cvm instance. Each element contains the following attributes:
	InstanceLists []GetSetInstanceList `pulumi:"instanceLists"`
	// Name of the instances.
	InstanceName *string `pulumi:"instanceName"`
	// The project CVM belongs to.
	ProjectId        *int    `pulumi:"projectId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of a vpc subnetwork.
	SubnetId *string `pulumi:"subnetId"`
	// Tags of the instance.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the vpc.
	VpcId *string `pulumi:"vpcId"`
}

func GetSetOutput(ctx *pulumi.Context, args GetSetOutputArgs, opts ...pulumi.InvokeOption) GetSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSetResult, error) {
			args := v.(GetSetArgs)
			r, err := GetSet(ctx, &args, opts...)
			var s GetSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSetResultOutput)
}

// A collection of arguments for invoking getSet.
type GetSetOutputArgs struct {
	// The available zone that the CVM instance locates at.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	// ID of the instances to be queried.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// Name of the instances to be queried.
	InstanceName pulumi.StringPtrInput `pulumi:"instanceName"`
	// The project CVM belongs to.
	ProjectId pulumi.IntPtrInput `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// ID of a vpc subnetwork.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// Tags of the instance.
	Tags pulumi.MapInput `pulumi:"tags"`
	// ID of the vpc to be queried.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSetArgs)(nil)).Elem()
}

// A collection of values returned by getSet.
type GetSetResultOutput struct{ *pulumi.OutputState }

func (GetSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSetResult)(nil)).Elem()
}

func (o GetSetResultOutput) ToGetSetResultOutput() GetSetResultOutput {
	return o
}

func (o GetSetResultOutput) ToGetSetResultOutputWithContext(ctx context.Context) GetSetResultOutput {
	return o
}

// The available zone that the CVM instance locates at.
func (o GetSetResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSetResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the instances.
func (o GetSetResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSetResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// An information list of cvm instance. Each element contains the following attributes:
func (o GetSetResultOutput) InstanceLists() GetSetInstanceListArrayOutput {
	return o.ApplyT(func(v GetSetResult) []GetSetInstanceList { return v.InstanceLists }).(GetSetInstanceListArrayOutput)
}

// Name of the instances.
func (o GetSetResultOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSetResult) *string { return v.InstanceName }).(pulumi.StringPtrOutput)
}

// The project CVM belongs to.
func (o GetSetResultOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetSetResult) *int { return v.ProjectId }).(pulumi.IntPtrOutput)
}

func (o GetSetResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSetResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// ID of a vpc subnetwork.
func (o GetSetResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSetResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Tags of the instance.
func (o GetSetResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetSetResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

// ID of the vpc.
func (o GetSetResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSetResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSetResultOutput{})
}
