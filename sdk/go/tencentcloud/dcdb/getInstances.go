// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of dcdb instances
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInstancesResult
	err := ctx.Invoke("tencentcloud:Dcdb/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// cluster excluster type.
	ExclusterType *int `pulumi:"exclusterType"`
	// instance ids.
	InstanceIds []string `pulumi:"instanceIds"`
	// search according to the cluster excluter type.
	IsFilterExcluster *bool `pulumi:"isFilterExcluster"`
	// search according to the vpc.
	IsFilterVpc *bool `pulumi:"isFilterVpc"`
	// project ids.
	ProjectIds []int `pulumi:"projectIds"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// search key, support fuzzy query.
	SearchKey *string `pulumi:"searchKey"`
	// search name, support instancename, vip, all.
	SearchName *string `pulumi:"searchName"`
	// subnet id, valid when IsFilterVpc is true.
	SubnetId *string `pulumi:"subnetId"`
	// vpc id, valid when IsFilterVpc is true.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	ExclusterType *int `pulumi:"exclusterType"`
	// The provider-assigned unique ID for this managed resource.
	Id                string   `pulumi:"id"`
	InstanceIds       []string `pulumi:"instanceIds"`
	IsFilterExcluster *bool    `pulumi:"isFilterExcluster"`
	IsFilterVpc       *bool    `pulumi:"isFilterVpc"`
	// instance list.
	Lists            []GetInstancesList `pulumi:"lists"`
	ProjectIds       []int              `pulumi:"projectIds"`
	ResultOutputFile *string            `pulumi:"resultOutputFile"`
	SearchKey        *string            `pulumi:"searchKey"`
	SearchName       *string            `pulumi:"searchName"`
	// subnet id.
	SubnetId *string `pulumi:"subnetId"`
	// vpc id.
	VpcId string `pulumi:"vpcId"`
}

func GetInstancesOutput(ctx *pulumi.Context, args GetInstancesOutputArgs, opts ...pulumi.InvokeOption) GetInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstancesResult, error) {
			args := v.(GetInstancesArgs)
			r, err := GetInstances(ctx, &args, opts...)
			var s GetInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstancesResultOutput)
}

// A collection of arguments for invoking getInstances.
type GetInstancesOutputArgs struct {
	// cluster excluster type.
	ExclusterType pulumi.IntPtrInput `pulumi:"exclusterType"`
	// instance ids.
	InstanceIds pulumi.StringArrayInput `pulumi:"instanceIds"`
	// search according to the cluster excluter type.
	IsFilterExcluster pulumi.BoolPtrInput `pulumi:"isFilterExcluster"`
	// search according to the vpc.
	IsFilterVpc pulumi.BoolPtrInput `pulumi:"isFilterVpc"`
	// project ids.
	ProjectIds pulumi.IntArrayInput `pulumi:"projectIds"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// search key, support fuzzy query.
	SearchKey pulumi.StringPtrInput `pulumi:"searchKey"`
	// search name, support instancename, vip, all.
	SearchName pulumi.StringPtrInput `pulumi:"searchName"`
	// subnet id, valid when IsFilterVpc is true.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// vpc id, valid when IsFilterVpc is true.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getInstances.
type GetInstancesResultOutput struct{ *pulumi.OutputState }

func (GetInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesResult)(nil)).Elem()
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutput() GetInstancesResultOutput {
	return o
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutputWithContext(ctx context.Context) GetInstancesResultOutput {
	return o
}

func (o GetInstancesResultOutput) ExclusterType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *int { return v.ExclusterType }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstancesResultOutput) InstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.InstanceIds }).(pulumi.StringArrayOutput)
}

func (o GetInstancesResultOutput) IsFilterExcluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *bool { return v.IsFilterExcluster }).(pulumi.BoolPtrOutput)
}

func (o GetInstancesResultOutput) IsFilterVpc() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *bool { return v.IsFilterVpc }).(pulumi.BoolPtrOutput)
}

// instance list.
func (o GetInstancesResultOutput) Lists() GetInstancesListArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesList { return v.Lists }).(GetInstancesListArrayOutput)
}

func (o GetInstancesResultOutput) ProjectIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []int { return v.ProjectIds }).(pulumi.IntArrayOutput)
}

func (o GetInstancesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) SearchKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.SearchKey }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) SearchName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.SearchName }).(pulumi.StringPtrOutput)
}

// subnet id.
func (o GetInstancesResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// vpc id.
func (o GetInstancesResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}
