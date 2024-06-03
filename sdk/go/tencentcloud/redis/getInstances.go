// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query the detail information of redis instance.
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Redis"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Redis.GetInstances(ctx, &redis.GetInstancesArgs{
//				Limit:            pulumi.IntRef(20),
//				ProjectId:        pulumi.IntRef(0),
//				ResultOutputFile: pulumi.StringRef("/tmp/redis_instances"),
//				SearchKey:        pulumi.StringRef("myredis"),
//				Zone:             pulumi.StringRef("ap-hongkong-1"),
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
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInstancesResult
	err := ctx.Invoke("tencentcloud:Redis/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// The number limitation of results for a query.
	Limit *int `pulumi:"limit"`
	// ID of the project to which redis instance belongs.
	ProjectId *int `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Key words used to match the results, and the key words can be: instance ID, instance name and IP address.
	SearchKey *string `pulumi:"searchKey"`
	// Tags of redis instance.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of an available zone.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of redis instance. Each element contains the following attributes:
	InstanceLists []GetInstancesInstanceList `pulumi:"instanceLists"`
	Limit         *int                       `pulumi:"limit"`
	// ID of the project to which a redis instance belongs.
	ProjectId        *int    `pulumi:"projectId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	SearchKey        *string `pulumi:"searchKey"`
	// Tags of an instance.
	Tags map[string]interface{} `pulumi:"tags"`
	// Available zone to which a redis instance belongs.
	Zone *string `pulumi:"zone"`
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
	// The number limitation of results for a query.
	Limit pulumi.IntPtrInput `pulumi:"limit"`
	// ID of the project to which redis instance belongs.
	ProjectId pulumi.IntPtrInput `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Key words used to match the results, and the key words can be: instance ID, instance name and IP address.
	SearchKey pulumi.StringPtrInput `pulumi:"searchKey"`
	// Tags of redis instance.
	Tags pulumi.MapInput `pulumi:"tags"`
	// ID of an available zone.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
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

// The provider-assigned unique ID for this managed resource.
func (o GetInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of redis instance. Each element contains the following attributes:
func (o GetInstancesResultOutput) InstanceLists() GetInstancesInstanceListArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesInstanceList { return v.InstanceLists }).(GetInstancesInstanceListArrayOutput)
}

func (o GetInstancesResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

// ID of the project to which a redis instance belongs.
func (o GetInstancesResultOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *int { return v.ProjectId }).(pulumi.IntPtrOutput)
}

func (o GetInstancesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) SearchKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.SearchKey }).(pulumi.StringPtrOutput)
}

// Tags of an instance.
func (o GetInstancesResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetInstancesResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

// Available zone to which a redis instance belongs.
func (o GetInstancesResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}
