// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query the list of backup databases.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mysql.GetBackupList(ctx, &mysql.GetBackupListArgs{
//				MaxNumber:        pulumi.IntRef(10),
//				MysqlId:          "terraform-test-local-database",
//				ResultOutputFile: pulumi.StringRef("mytestpath"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetBackupList(ctx *pulumi.Context, args *GetBackupListArgs, opts ...pulumi.InvokeOption) (*GetBackupListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetBackupListResult
	err := ctx.Invoke("tencentcloud:Mysql/getBackupList:getBackupList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackupList.
type GetBackupListArgs struct {
	// The latest files to list, rang from 1 to 10000. And the default value is `10`.
	MaxNumber *int `pulumi:"maxNumber"`
	// Instance ID, such as `cdb-c1nl9rpv`. It is identical to the instance ID displayed in the database console page.
	MysqlId string `pulumi:"mysqlId"`
	// Used to store results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getBackupList.
type GetBackupListResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of MySQL backup. Each element contains the following attributes:
	Lists            []GetBackupListList `pulumi:"lists"`
	MaxNumber        *int                `pulumi:"maxNumber"`
	MysqlId          string              `pulumi:"mysqlId"`
	ResultOutputFile *string             `pulumi:"resultOutputFile"`
}

func GetBackupListOutput(ctx *pulumi.Context, args GetBackupListOutputArgs, opts ...pulumi.InvokeOption) GetBackupListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBackupListResult, error) {
			args := v.(GetBackupListArgs)
			r, err := GetBackupList(ctx, &args, opts...)
			var s GetBackupListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBackupListResultOutput)
}

// A collection of arguments for invoking getBackupList.
type GetBackupListOutputArgs struct {
	// The latest files to list, rang from 1 to 10000. And the default value is `10`.
	MaxNumber pulumi.IntPtrInput `pulumi:"maxNumber"`
	// Instance ID, such as `cdb-c1nl9rpv`. It is identical to the instance ID displayed in the database console page.
	MysqlId pulumi.StringInput `pulumi:"mysqlId"`
	// Used to store results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetBackupListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupListArgs)(nil)).Elem()
}

// A collection of values returned by getBackupList.
type GetBackupListResultOutput struct{ *pulumi.OutputState }

func (GetBackupListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupListResult)(nil)).Elem()
}

func (o GetBackupListResultOutput) ToGetBackupListResultOutput() GetBackupListResultOutput {
	return o
}

func (o GetBackupListResultOutput) ToGetBackupListResultOutputWithContext(ctx context.Context) GetBackupListResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetBackupListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of MySQL backup. Each element contains the following attributes:
func (o GetBackupListResultOutput) Lists() GetBackupListListArrayOutput {
	return o.ApplyT(func(v GetBackupListResult) []GetBackupListList { return v.Lists }).(GetBackupListListArrayOutput)
}

func (o GetBackupListResultOutput) MaxNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetBackupListResult) *int { return v.MaxNumber }).(pulumi.IntPtrOutput)
}

func (o GetBackupListResultOutput) MysqlId() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupListResult) string { return v.MysqlId }).(pulumi.StringOutput)
}

func (o GetBackupListResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupListResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBackupListResultOutput{})
}
