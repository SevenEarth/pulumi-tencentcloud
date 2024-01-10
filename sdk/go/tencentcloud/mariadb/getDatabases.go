// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of mariadb databases
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mariadb.GetDatabases(ctx, &mariadb.GetDatabasesArgs{
//				InstanceId: "tdsql-e9tklsgz",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDatabases(ctx *pulumi.Context, args *GetDatabasesArgs, opts ...pulumi.InvokeOption) (*GetDatabasesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDatabasesResult
	err := ctx.Invoke("tencentcloud:Mariadb/getDatabases:getDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabases.
type GetDatabasesArgs struct {
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDatabases.
type GetDatabasesResult struct {
	// The database list of this instance.
	Databases []GetDatabasesDatabase `pulumi:"databases"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	InstanceId       string  `pulumi:"instanceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetDatabasesOutput(ctx *pulumi.Context, args GetDatabasesOutputArgs, opts ...pulumi.InvokeOption) GetDatabasesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabasesResult, error) {
			args := v.(GetDatabasesArgs)
			r, err := GetDatabases(ctx, &args, opts...)
			var s GetDatabasesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatabasesResultOutput)
}

// A collection of arguments for invoking getDatabases.
type GetDatabasesOutputArgs struct {
	// instance id.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDatabasesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesArgs)(nil)).Elem()
}

// A collection of values returned by getDatabases.
type GetDatabasesResultOutput struct{ *pulumi.OutputState }

func (GetDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesResult)(nil)).Elem()
}

func (o GetDatabasesResultOutput) ToGetDatabasesResultOutput() GetDatabasesResultOutput {
	return o
}

func (o GetDatabasesResultOutput) ToGetDatabasesResultOutputWithContext(ctx context.Context) GetDatabasesResultOutput {
	return o
}

// The database list of this instance.
func (o GetDatabasesResultOutput) Databases() GetDatabasesDatabaseArrayOutput {
	return o.ApplyT(func(v GetDatabasesResult) []GetDatabasesDatabase { return v.Databases }).(GetDatabasesDatabaseArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDatabasesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDatabasesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetDatabasesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabasesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabasesResultOutput{})
}
