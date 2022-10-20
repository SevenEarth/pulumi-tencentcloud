// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query DB resources for the specific SQL Server instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Sqlserver.GetDbs(ctx, &sqlserver.GetDbsArgs{
//				InstanceId: "mssql-3cdq7kx5",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDbs(ctx *pulumi.Context, args *GetDbsArgs, opts ...pulumi.InvokeOption) (*GetDbsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDbsResult
	err := ctx.Invoke("tencentcloud:Sqlserver/getDbs:getDbs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDbs.
type GetDbsArgs struct {
	// SQL Server instance ID which DB belongs to.
	InstanceId string `pulumi:"instanceId"`
	// Used to store results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDbs.
type GetDbsResult struct {
	// A list of dbs belong to the specific instance. Each element contains the following attributes:
	DbLists []GetDbsDbList `pulumi:"dbLists"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// SQL Server instance ID which DB belongs to.
	InstanceId       string  `pulumi:"instanceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetDbsOutput(ctx *pulumi.Context, args GetDbsOutputArgs, opts ...pulumi.InvokeOption) GetDbsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDbsResult, error) {
			args := v.(GetDbsArgs)
			r, err := GetDbs(ctx, &args, opts...)
			var s GetDbsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDbsResultOutput)
}

// A collection of arguments for invoking getDbs.
type GetDbsOutputArgs struct {
	// SQL Server instance ID which DB belongs to.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to store results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDbsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDbsArgs)(nil)).Elem()
}

// A collection of values returned by getDbs.
type GetDbsResultOutput struct{ *pulumi.OutputState }

func (GetDbsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDbsResult)(nil)).Elem()
}

func (o GetDbsResultOutput) ToGetDbsResultOutput() GetDbsResultOutput {
	return o
}

func (o GetDbsResultOutput) ToGetDbsResultOutputWithContext(ctx context.Context) GetDbsResultOutput {
	return o
}

// A list of dbs belong to the specific instance. Each element contains the following attributes:
func (o GetDbsResultOutput) DbLists() GetDbsDbListArrayOutput {
	return o.ApplyT(func(v GetDbsResult) []GetDbsDbList { return v.DbLists }).(GetDbsDbListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDbsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDbsResult) string { return v.Id }).(pulumi.StringOutput)
}

// SQL Server instance ID which DB belongs to.
func (o GetDbsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDbsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetDbsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDbsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDbsResultOutput{})
}
