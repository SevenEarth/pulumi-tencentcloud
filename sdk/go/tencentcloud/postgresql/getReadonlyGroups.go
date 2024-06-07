// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of postgresql readOnlyGroups
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
//			group, err := Postgresql.NewReadonlyGroup(ctx, "group", &Postgresql.ReadonlyGroupArgs{
//				MasterDbInstanceId:       pulumi.String("postgres-gzg9jb2n"),
//				ProjectId:                pulumi.Int(0),
//				VpcId:                    pulumi.String("vpc-86v957zb"),
//				SubnetId:                 pulumi.String("subnet-enm92y0m"),
//				ReplayLagEliminate:       pulumi.Int(1),
//				ReplayLatencyEliminate:   pulumi.Int(1),
//				MaxReplayLag:             pulumi.Int(100),
//				MaxReplayLatency:         pulumi.Int(512),
//				MinDelayEliminateReserve: pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			_ = Postgresql.GetReadonlyGroupsOutput(ctx, postgresql.GetReadonlyGroupsOutputArgs{
//				Filters: postgresql.GetReadonlyGroupsFilterArray{
//					&postgresql.GetReadonlyGroupsFilterArgs{
//						Name: pulumi.String("db-master-instance-id"),
//						Values: pulumi.StringArray{
//							group.MasterDbInstanceId,
//						},
//					},
//				},
//				OrderBy:     pulumi.String("CreateTime"),
//				OrderByType: pulumi.String("asc"),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetReadonlyGroups(ctx *pulumi.Context, args *GetReadonlyGroupsArgs, opts ...pulumi.InvokeOption) (*GetReadonlyGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetReadonlyGroupsResult
	err := ctx.Invoke("tencentcloud:Postgresql/getReadonlyGroups:getReadonlyGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReadonlyGroups.
type GetReadonlyGroupsArgs struct {
	// Filter condition. The primary ID must be specified in the format of db-master-instance-id to filter results, or else null will be returned.
	Filters []GetReadonlyGroupsFilter `pulumi:"filters"`
	// Sorting criterion. Valid values:ROGroupId, CreateTime, Name.
	OrderBy *string `pulumi:"orderBy"`
	// Sorting order. Valid values:desc, asc.
	OrderByType *string `pulumi:"orderByType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getReadonlyGroups.
type GetReadonlyGroupsResult struct {
	Filters []GetReadonlyGroupsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	OrderBy     *string `pulumi:"orderBy"`
	OrderByType *string `pulumi:"orderByType"`
	// list of read-only groups.
	ReadOnlyGroupLists []GetReadonlyGroupsReadOnlyGroupList `pulumi:"readOnlyGroupLists"`
	ResultOutputFile   *string                              `pulumi:"resultOutputFile"`
}

func GetReadonlyGroupsOutput(ctx *pulumi.Context, args GetReadonlyGroupsOutputArgs, opts ...pulumi.InvokeOption) GetReadonlyGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetReadonlyGroupsResult, error) {
			args := v.(GetReadonlyGroupsArgs)
			r, err := GetReadonlyGroups(ctx, &args, opts...)
			var s GetReadonlyGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetReadonlyGroupsResultOutput)
}

// A collection of arguments for invoking getReadonlyGroups.
type GetReadonlyGroupsOutputArgs struct {
	// Filter condition. The primary ID must be specified in the format of db-master-instance-id to filter results, or else null will be returned.
	Filters GetReadonlyGroupsFilterArrayInput `pulumi:"filters"`
	// Sorting criterion. Valid values:ROGroupId, CreateTime, Name.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// Sorting order. Valid values:desc, asc.
	OrderByType pulumi.StringPtrInput `pulumi:"orderByType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetReadonlyGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReadonlyGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getReadonlyGroups.
type GetReadonlyGroupsResultOutput struct{ *pulumi.OutputState }

func (GetReadonlyGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReadonlyGroupsResult)(nil)).Elem()
}

func (o GetReadonlyGroupsResultOutput) ToGetReadonlyGroupsResultOutput() GetReadonlyGroupsResultOutput {
	return o
}

func (o GetReadonlyGroupsResultOutput) ToGetReadonlyGroupsResultOutputWithContext(ctx context.Context) GetReadonlyGroupsResultOutput {
	return o
}

func (o GetReadonlyGroupsResultOutput) Filters() GetReadonlyGroupsFilterArrayOutput {
	return o.ApplyT(func(v GetReadonlyGroupsResult) []GetReadonlyGroupsFilter { return v.Filters }).(GetReadonlyGroupsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetReadonlyGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetReadonlyGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetReadonlyGroupsResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetReadonlyGroupsResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o GetReadonlyGroupsResultOutput) OrderByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetReadonlyGroupsResult) *string { return v.OrderByType }).(pulumi.StringPtrOutput)
}

// list of read-only groups.
func (o GetReadonlyGroupsResultOutput) ReadOnlyGroupLists() GetReadonlyGroupsReadOnlyGroupListArrayOutput {
	return o.ApplyT(func(v GetReadonlyGroupsResult) []GetReadonlyGroupsReadOnlyGroupList { return v.ReadOnlyGroupLists }).(GetReadonlyGroupsReadOnlyGroupListArrayOutput)
}

func (o GetReadonlyGroupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetReadonlyGroupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetReadonlyGroupsResultOutput{})
}
