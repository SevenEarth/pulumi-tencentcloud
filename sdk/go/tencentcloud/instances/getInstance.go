// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package instances

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query cvm instances.
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
//			_, err := Instances.GetInstance(ctx, &instances.GetInstanceArgs{
//				InstanceId: pulumi.StringRef("ins-da412f5a"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetInstance(ctx *pulumi.Context, args *GetInstanceArgs, opts ...pulumi.InvokeOption) (*GetInstanceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetInstanceResult
	err := ctx.Invoke("tencentcloud:Instances/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type GetInstanceArgs struct {
	// The available zone that the CVM instance locates at.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// ID of the instances to be queried.
	InstanceId *string `pulumi:"instanceId"`
	// Name of the instances to be queried.
	InstanceName *string `pulumi:"instanceName"`
	// Instance set ids, max length is 100, conflict with other field.
	InstanceSetIds []string `pulumi:"instanceSetIds"`
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

// A collection of values returned by getInstance.
type GetInstanceResult struct {
	// The available zone that the CVM instance locates at.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the instances.
	InstanceId *string `pulumi:"instanceId"`
	// An information list of cvm instance. Each element contains the following attributes:
	InstanceLists []GetInstanceInstanceList `pulumi:"instanceLists"`
	// Name of the instances.
	InstanceName   *string  `pulumi:"instanceName"`
	InstanceSetIds []string `pulumi:"instanceSetIds"`
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

func GetInstanceOutput(ctx *pulumi.Context, args GetInstanceOutputArgs, opts ...pulumi.InvokeOption) GetInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceResult, error) {
			args := v.(GetInstanceArgs)
			r, err := GetInstance(ctx, &args, opts...)
			var s GetInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceResultOutput)
}

// A collection of arguments for invoking getInstance.
type GetInstanceOutputArgs struct {
	// The available zone that the CVM instance locates at.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	// ID of the instances to be queried.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// Name of the instances to be queried.
	InstanceName pulumi.StringPtrInput `pulumi:"instanceName"`
	// Instance set ids, max length is 100, conflict with other field.
	InstanceSetIds pulumi.StringArrayInput `pulumi:"instanceSetIds"`
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

func (GetInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getInstance.
type GetInstanceResultOutput struct{ *pulumi.OutputState }

func (GetInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceResult)(nil)).Elem()
}

func (o GetInstanceResultOutput) ToGetInstanceResultOutput() GetInstanceResultOutput {
	return o
}

func (o GetInstanceResultOutput) ToGetInstanceResultOutputWithContext(ctx context.Context) GetInstanceResultOutput {
	return o
}

// The available zone that the CVM instance locates at.
func (o GetInstanceResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the instances.
func (o GetInstanceResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// An information list of cvm instance. Each element contains the following attributes:
func (o GetInstanceResultOutput) InstanceLists() GetInstanceInstanceListArrayOutput {
	return o.ApplyT(func(v GetInstanceResult) []GetInstanceInstanceList { return v.InstanceLists }).(GetInstanceInstanceListArrayOutput)
}

// Name of the instances.
func (o GetInstanceResultOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.InstanceName }).(pulumi.StringPtrOutput)
}

func (o GetInstanceResultOutput) InstanceSetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceResult) []string { return v.InstanceSetIds }).(pulumi.StringArrayOutput)
}

// The project CVM belongs to.
func (o GetInstanceResultOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *int { return v.ProjectId }).(pulumi.IntPtrOutput)
}

func (o GetInstanceResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// ID of a vpc subnetwork.
func (o GetInstanceResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Tags of the instance.
func (o GetInstanceResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetInstanceResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

// ID of the vpc.
func (o GetInstanceResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceResultOutput{})
}
