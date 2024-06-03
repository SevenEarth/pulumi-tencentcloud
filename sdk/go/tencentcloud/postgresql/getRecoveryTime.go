// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of postgresql recoveryTime
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Postgresql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Postgresql.GetRecoveryTime(ctx, &postgresql.GetRecoveryTimeArgs{
//				DbInstanceId: local.Pgsql_id,
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
func GetRecoveryTime(ctx *pulumi.Context, args *GetRecoveryTimeArgs, opts ...pulumi.InvokeOption) (*GetRecoveryTimeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRecoveryTimeResult
	err := ctx.Invoke("tencentcloud:Postgresql/getRecoveryTime:getRecoveryTime", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRecoveryTime.
type GetRecoveryTimeArgs struct {
	// Instance ID.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getRecoveryTime.
type GetRecoveryTimeResult struct {
	DbInstanceId string `pulumi:"dbInstanceId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The earliest restoration time (UTC+8).
	RecoveryBeginTime string `pulumi:"recoveryBeginTime"`
	// The latest restoration time (UTC+8).
	RecoveryEndTime  string  `pulumi:"recoveryEndTime"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetRecoveryTimeOutput(ctx *pulumi.Context, args GetRecoveryTimeOutputArgs, opts ...pulumi.InvokeOption) GetRecoveryTimeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRecoveryTimeResult, error) {
			args := v.(GetRecoveryTimeArgs)
			r, err := GetRecoveryTime(ctx, &args, opts...)
			var s GetRecoveryTimeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRecoveryTimeResultOutput)
}

// A collection of arguments for invoking getRecoveryTime.
type GetRecoveryTimeOutputArgs struct {
	// Instance ID.
	DbInstanceId pulumi.StringInput `pulumi:"dbInstanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetRecoveryTimeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecoveryTimeArgs)(nil)).Elem()
}

// A collection of values returned by getRecoveryTime.
type GetRecoveryTimeResultOutput struct{ *pulumi.OutputState }

func (GetRecoveryTimeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecoveryTimeResult)(nil)).Elem()
}

func (o GetRecoveryTimeResultOutput) ToGetRecoveryTimeResultOutput() GetRecoveryTimeResultOutput {
	return o
}

func (o GetRecoveryTimeResultOutput) ToGetRecoveryTimeResultOutputWithContext(ctx context.Context) GetRecoveryTimeResultOutput {
	return o
}

func (o GetRecoveryTimeResultOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecoveryTimeResult) string { return v.DbInstanceId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRecoveryTimeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecoveryTimeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The earliest restoration time (UTC+8).
func (o GetRecoveryTimeResultOutput) RecoveryBeginTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecoveryTimeResult) string { return v.RecoveryBeginTime }).(pulumi.StringOutput)
}

// The latest restoration time (UTC+8).
func (o GetRecoveryTimeResultOutput) RecoveryEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecoveryTimeResult) string { return v.RecoveryEndTime }).(pulumi.StringOutput)
}

func (o GetRecoveryTimeResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecoveryTimeResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRecoveryTimeResultOutput{})
}
