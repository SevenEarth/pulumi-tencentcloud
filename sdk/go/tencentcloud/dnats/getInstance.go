// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dnats

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of DNATs.
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dnats"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dnats.GetInstance(ctx, &dnats.GetInstanceArgs{
//				ElasticIp: pulumi.StringRef("123.207.115.136"),
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
func GetInstance(ctx *pulumi.Context, args *GetInstanceArgs, opts ...pulumi.InvokeOption) (*GetInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInstanceResult
	err := ctx.Invoke("tencentcloud:Dnats/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type GetInstanceArgs struct {
	// Description of the NAT forward.
	Description *string `pulumi:"description"`
	// Network address of the EIP.
	ElasticIp *string `pulumi:"elasticIp"`
	// Port of the EIP.
	ElasticPort *string `pulumi:"elasticPort"`
	// ID of the NAT gateway.
	NatId *string `pulumi:"natId"`
	// Network address of the backend service.
	PrivateIp *string `pulumi:"privateIp"`
	// Port of intranet.
	PrivatePort *string `pulumi:"privatePort"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of the VPC.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getInstance.
type GetInstanceResult struct {
	Description *string `pulumi:"description"`
	// Information list of the DNATs.
	DnatLists []GetInstanceDnatList `pulumi:"dnatLists"`
	// Network address of the EIP.
	ElasticIp *string `pulumi:"elasticIp"`
	// Port of the EIP.
	ElasticPort *string `pulumi:"elasticPort"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the NAT.
	NatId *string `pulumi:"natId"`
	// Network address of the backend service.
	PrivateIp *string `pulumi:"privateIp"`
	// Port of intranet.
	PrivatePort      *string `pulumi:"privatePort"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of the VPC.
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
	// Description of the NAT forward.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Network address of the EIP.
	ElasticIp pulumi.StringPtrInput `pulumi:"elasticIp"`
	// Port of the EIP.
	ElasticPort pulumi.StringPtrInput `pulumi:"elasticPort"`
	// ID of the NAT gateway.
	NatId pulumi.StringPtrInput `pulumi:"natId"`
	// Network address of the backend service.
	PrivateIp pulumi.StringPtrInput `pulumi:"privateIp"`
	// Port of intranet.
	PrivatePort pulumi.StringPtrInput `pulumi:"privatePort"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// ID of the VPC.
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

func (o GetInstanceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Information list of the DNATs.
func (o GetInstanceResultOutput) DnatLists() GetInstanceDnatListArrayOutput {
	return o.ApplyT(func(v GetInstanceResult) []GetInstanceDnatList { return v.DnatLists }).(GetInstanceDnatListArrayOutput)
}

// Network address of the EIP.
func (o GetInstanceResultOutput) ElasticIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.ElasticIp }).(pulumi.StringPtrOutput)
}

// Port of the EIP.
func (o GetInstanceResultOutput) ElasticPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.ElasticPort }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the NAT.
func (o GetInstanceResultOutput) NatId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.NatId }).(pulumi.StringPtrOutput)
}

// Network address of the backend service.
func (o GetInstanceResultOutput) PrivateIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.PrivateIp }).(pulumi.StringPtrOutput)
}

// Port of intranet.
func (o GetInstanceResultOutput) PrivatePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.PrivatePort }).(pulumi.StringPtrOutput)
}

func (o GetInstanceResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// ID of the VPC.
func (o GetInstanceResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceResultOutput{})
}
