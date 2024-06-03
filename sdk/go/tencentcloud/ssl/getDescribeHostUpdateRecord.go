// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of ssl describeHostUpdateRecord
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ssl"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Ssl.GetDescribeHostUpdateRecord(ctx, &ssl.GetDescribeHostUpdateRecordArgs{
//				OldCertificateId: pulumi.StringRef("8u8DII0l"),
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
func GetDescribeHostUpdateRecord(ctx *pulumi.Context, args *GetDescribeHostUpdateRecordArgs, opts ...pulumi.InvokeOption) (*GetDescribeHostUpdateRecordResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDescribeHostUpdateRecordResult
	err := ctx.Invoke("tencentcloud:Ssl/getDescribeHostUpdateRecord:getDescribeHostUpdateRecord", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDescribeHostUpdateRecord.
type GetDescribeHostUpdateRecordArgs struct {
	// New certificate ID.
	CertificateId *string `pulumi:"certificateId"`
	// Original certificate ID.
	OldCertificateId *string `pulumi:"oldCertificateId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDescribeHostUpdateRecord.
type GetDescribeHostUpdateRecordResult struct {
	CertificateId *string `pulumi:"certificateId"`
	// Certificate deployment record listNote: This field may return NULL, indicating that the valid value cannot be obtained.
	DeployRecordLists []GetDescribeHostUpdateRecordDeployRecordList `pulumi:"deployRecordLists"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	OldCertificateId *string `pulumi:"oldCertificateId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetDescribeHostUpdateRecordOutput(ctx *pulumi.Context, args GetDescribeHostUpdateRecordOutputArgs, opts ...pulumi.InvokeOption) GetDescribeHostUpdateRecordResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDescribeHostUpdateRecordResult, error) {
			args := v.(GetDescribeHostUpdateRecordArgs)
			r, err := GetDescribeHostUpdateRecord(ctx, &args, opts...)
			var s GetDescribeHostUpdateRecordResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDescribeHostUpdateRecordResultOutput)
}

// A collection of arguments for invoking getDescribeHostUpdateRecord.
type GetDescribeHostUpdateRecordOutputArgs struct {
	// New certificate ID.
	CertificateId pulumi.StringPtrInput `pulumi:"certificateId"`
	// Original certificate ID.
	OldCertificateId pulumi.StringPtrInput `pulumi:"oldCertificateId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDescribeHostUpdateRecordOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeHostUpdateRecordArgs)(nil)).Elem()
}

// A collection of values returned by getDescribeHostUpdateRecord.
type GetDescribeHostUpdateRecordResultOutput struct{ *pulumi.OutputState }

func (GetDescribeHostUpdateRecordResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeHostUpdateRecordResult)(nil)).Elem()
}

func (o GetDescribeHostUpdateRecordResultOutput) ToGetDescribeHostUpdateRecordResultOutput() GetDescribeHostUpdateRecordResultOutput {
	return o
}

func (o GetDescribeHostUpdateRecordResultOutput) ToGetDescribeHostUpdateRecordResultOutputWithContext(ctx context.Context) GetDescribeHostUpdateRecordResultOutput {
	return o
}

func (o GetDescribeHostUpdateRecordResultOutput) CertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeHostUpdateRecordResult) *string { return v.CertificateId }).(pulumi.StringPtrOutput)
}

// Certificate deployment record listNote: This field may return NULL, indicating that the valid value cannot be obtained.
func (o GetDescribeHostUpdateRecordResultOutput) DeployRecordLists() GetDescribeHostUpdateRecordDeployRecordListArrayOutput {
	return o.ApplyT(func(v GetDescribeHostUpdateRecordResult) []GetDescribeHostUpdateRecordDeployRecordList {
		return v.DeployRecordLists
	}).(GetDescribeHostUpdateRecordDeployRecordListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDescribeHostUpdateRecordResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeHostUpdateRecordResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDescribeHostUpdateRecordResultOutput) OldCertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeHostUpdateRecordResult) *string { return v.OldCertificateId }).(pulumi.StringPtrOutput)
}

func (o GetDescribeHostUpdateRecordResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeHostUpdateRecordResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDescribeHostUpdateRecordResultOutput{})
}
