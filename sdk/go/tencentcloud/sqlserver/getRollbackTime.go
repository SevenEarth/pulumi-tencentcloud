// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of sqlserver rollbackTime
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Sqlserver.GetRollbackTime(ctx, &sqlserver.GetRollbackTimeArgs{
//				Dbs: []string{
//					"keep_pubsub_db",
//				},
//				InstanceId: "mssql-qelbzgwf",
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
func GetRollbackTime(ctx *pulumi.Context, args *GetRollbackTimeArgs, opts ...pulumi.InvokeOption) (*GetRollbackTimeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRollbackTimeResult
	err := ctx.Invoke("tencentcloud:Sqlserver/getRollbackTime:getRollbackTime", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRollbackTime.
type GetRollbackTimeArgs struct {
	// List of databases to be queried.
	Dbs []string `pulumi:"dbs"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getRollbackTime.
type GetRollbackTimeResult struct {
	Dbs []string `pulumi:"dbs"`
	// Information of time range available for database rollback.
	Details []GetRollbackTimeDetail `pulumi:"details"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	InstanceId       string  `pulumi:"instanceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetRollbackTimeOutput(ctx *pulumi.Context, args GetRollbackTimeOutputArgs, opts ...pulumi.InvokeOption) GetRollbackTimeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRollbackTimeResult, error) {
			args := v.(GetRollbackTimeArgs)
			r, err := GetRollbackTime(ctx, &args, opts...)
			var s GetRollbackTimeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRollbackTimeResultOutput)
}

// A collection of arguments for invoking getRollbackTime.
type GetRollbackTimeOutputArgs struct {
	// List of databases to be queried.
	Dbs pulumi.StringArrayInput `pulumi:"dbs"`
	// Instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetRollbackTimeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRollbackTimeArgs)(nil)).Elem()
}

// A collection of values returned by getRollbackTime.
type GetRollbackTimeResultOutput struct{ *pulumi.OutputState }

func (GetRollbackTimeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRollbackTimeResult)(nil)).Elem()
}

func (o GetRollbackTimeResultOutput) ToGetRollbackTimeResultOutput() GetRollbackTimeResultOutput {
	return o
}

func (o GetRollbackTimeResultOutput) ToGetRollbackTimeResultOutputWithContext(ctx context.Context) GetRollbackTimeResultOutput {
	return o
}

func (o GetRollbackTimeResultOutput) Dbs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRollbackTimeResult) []string { return v.Dbs }).(pulumi.StringArrayOutput)
}

// Information of time range available for database rollback.
func (o GetRollbackTimeResultOutput) Details() GetRollbackTimeDetailArrayOutput {
	return o.ApplyT(func(v GetRollbackTimeResult) []GetRollbackTimeDetail { return v.Details }).(GetRollbackTimeDetailArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRollbackTimeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRollbackTimeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRollbackTimeResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRollbackTimeResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetRollbackTimeResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRollbackTimeResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRollbackTimeResultOutput{})
}
