// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package enis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query query ENIs.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Enis"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Enis"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Enis.GetInstance(ctx, &enis.GetInstanceArgs{
//				Name: pulumi.StringRef("test eni"),
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
	var rv GetInstanceResult
	err := ctx.Invoke("tencentcloud:Enis/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type GetInstanceArgs struct {
	// Description of the ENI. Conflict with `ids`.
	Description *string `pulumi:"description"`
	// ID of the ENIs to be queried. Conflict with `vpcId`,`subnetId`,`instanceId`,`securityGroup`,`name`,`ipv4` and `tags`.
	Ids []string `pulumi:"ids"`
	// ID of the instance which bind the ENI. Conflict with `ids`.
	InstanceId *string `pulumi:"instanceId"`
	// Intranet IP of the ENI. Conflict with `ids`.
	Ipv4 *string `pulumi:"ipv4"`
	// Name of the ENI to be queried. Conflict with `ids`.
	Name *string `pulumi:"name"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// A set of security group IDs which bind the ENI. Conflict with `ids`.
	SecurityGroup *string `pulumi:"securityGroup"`
	// ID of the subnet within this vpc to be queried. Conflict with `ids`.
	SubnetId *string `pulumi:"subnetId"`
	// Tags of the ENI. Conflict with `ids`.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the vpc to be queried. Conflict with `ids`.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getInstance.
type GetInstanceResult struct {
	// Description of the IP.
	Description *string `pulumi:"description"`
	// An information list of ENIs. Each element contains the following attributes:
	Enis []GetInstanceEni `pulumi:"enis"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// ID of the instance which bind the ENI.
	InstanceId *string `pulumi:"instanceId"`
	Ipv4       *string `pulumi:"ipv4"`
	// Name of the ENI.
	Name             *string `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	SecurityGroup    *string `pulumi:"securityGroup"`
	// ID of the subnet within this vpc.
	SubnetId *string `pulumi:"subnetId"`
	// Tags of the ENI.
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
	// Description of the ENI. Conflict with `ids`.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// ID of the ENIs to be queried. Conflict with `vpcId`,`subnetId`,`instanceId`,`securityGroup`,`name`,`ipv4` and `tags`.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// ID of the instance which bind the ENI. Conflict with `ids`.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// Intranet IP of the ENI. Conflict with `ids`.
	Ipv4 pulumi.StringPtrInput `pulumi:"ipv4"`
	// Name of the ENI to be queried. Conflict with `ids`.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// A set of security group IDs which bind the ENI. Conflict with `ids`.
	SecurityGroup pulumi.StringPtrInput `pulumi:"securityGroup"`
	// ID of the subnet within this vpc to be queried. Conflict with `ids`.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// Tags of the ENI. Conflict with `ids`.
	Tags pulumi.MapInput `pulumi:"tags"`
	// ID of the vpc to be queried. Conflict with `ids`.
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

// Description of the IP.
func (o GetInstanceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// An information list of ENIs. Each element contains the following attributes:
func (o GetInstanceResultOutput) Enis() GetInstanceEniArrayOutput {
	return o.ApplyT(func(v GetInstanceResult) []GetInstanceEni { return v.Enis }).(GetInstanceEniArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstanceResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// ID of the instance which bind the ENI.
func (o GetInstanceResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o GetInstanceResultOutput) Ipv4() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.Ipv4 }).(pulumi.StringPtrOutput)
}

// Name of the ENI.
func (o GetInstanceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetInstanceResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetInstanceResultOutput) SecurityGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.SecurityGroup }).(pulumi.StringPtrOutput)
}

// ID of the subnet within this vpc.
func (o GetInstanceResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Tags of the ENI.
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
