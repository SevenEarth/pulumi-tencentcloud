// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of mariadb dcnDetail
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mariadb.GetDcnDetail(ctx, &mariadb.GetDcnDetailArgs{
//				InstanceId: "tdsql-9vqvls95",
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
func GetDcnDetail(ctx *pulumi.Context, args *GetDcnDetailArgs, opts ...pulumi.InvokeOption) (*GetDcnDetailResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDcnDetailResult
	err := ctx.Invoke("tencentcloud:Mariadb/getDcnDetail:getDcnDetail", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDcnDetail.
type GetDcnDetailArgs struct {
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDcnDetail.
type GetDcnDetailResult struct {
	// DCN synchronization details.
	DcnDetails []GetDcnDetailDcnDetail `pulumi:"dcnDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Instance ID.
	InstanceId       string  `pulumi:"instanceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetDcnDetailOutput(ctx *pulumi.Context, args GetDcnDetailOutputArgs, opts ...pulumi.InvokeOption) GetDcnDetailResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDcnDetailResult, error) {
			args := v.(GetDcnDetailArgs)
			r, err := GetDcnDetail(ctx, &args, opts...)
			var s GetDcnDetailResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDcnDetailResultOutput)
}

// A collection of arguments for invoking getDcnDetail.
type GetDcnDetailOutputArgs struct {
	// Instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDcnDetailOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDcnDetailArgs)(nil)).Elem()
}

// A collection of values returned by getDcnDetail.
type GetDcnDetailResultOutput struct{ *pulumi.OutputState }

func (GetDcnDetailResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDcnDetailResult)(nil)).Elem()
}

func (o GetDcnDetailResultOutput) ToGetDcnDetailResultOutput() GetDcnDetailResultOutput {
	return o
}

func (o GetDcnDetailResultOutput) ToGetDcnDetailResultOutputWithContext(ctx context.Context) GetDcnDetailResultOutput {
	return o
}

// DCN synchronization details.
func (o GetDcnDetailResultOutput) DcnDetails() GetDcnDetailDcnDetailArrayOutput {
	return o.ApplyT(func(v GetDcnDetailResult) []GetDcnDetailDcnDetail { return v.DcnDetails }).(GetDcnDetailDcnDetailArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDcnDetailResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDcnDetailResult) string { return v.Id }).(pulumi.StringOutput)
}

// Instance ID.
func (o GetDcnDetailResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDcnDetailResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetDcnDetailResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDcnDetailResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDcnDetailResultOutput{})
}
