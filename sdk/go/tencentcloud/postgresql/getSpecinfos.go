// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the available product configs of the postgresql instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Postgresql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Postgresql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Postgresql.GetSpecinfos(ctx, &postgresql.GetSpecinfosArgs{
//				AvailabilityZone: "ap-shanghai-2",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSpecinfos(ctx *pulumi.Context, args *GetSpecinfosArgs, opts ...pulumi.InvokeOption) (*GetSpecinfosResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSpecinfosResult
	err := ctx.Invoke("tencentcloud:Postgresql/getSpecinfos:getSpecinfos", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSpecinfos.
type GetSpecinfosArgs struct {
	// The zone of the postgresql instance to query.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getSpecinfos.
type GetSpecinfosResult struct {
	AvailabilityZone string `pulumi:"availabilityZone"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of zones will be exported and its every element contains the following attributes:
	Lists            []GetSpecinfosList `pulumi:"lists"`
	ResultOutputFile *string            `pulumi:"resultOutputFile"`
}

func GetSpecinfosOutput(ctx *pulumi.Context, args GetSpecinfosOutputArgs, opts ...pulumi.InvokeOption) GetSpecinfosResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSpecinfosResult, error) {
			args := v.(GetSpecinfosArgs)
			r, err := GetSpecinfos(ctx, &args, opts...)
			var s GetSpecinfosResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSpecinfosResultOutput)
}

// A collection of arguments for invoking getSpecinfos.
type GetSpecinfosOutputArgs struct {
	// The zone of the postgresql instance to query.
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetSpecinfosOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSpecinfosArgs)(nil)).Elem()
}

// A collection of values returned by getSpecinfos.
type GetSpecinfosResultOutput struct{ *pulumi.OutputState }

func (GetSpecinfosResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSpecinfosResult)(nil)).Elem()
}

func (o GetSpecinfosResultOutput) ToGetSpecinfosResultOutput() GetSpecinfosResultOutput {
	return o
}

func (o GetSpecinfosResultOutput) ToGetSpecinfosResultOutputWithContext(ctx context.Context) GetSpecinfosResultOutput {
	return o
}

func (o GetSpecinfosResultOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v GetSpecinfosResult) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSpecinfosResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSpecinfosResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of zones will be exported and its every element contains the following attributes:
func (o GetSpecinfosResultOutput) Lists() GetSpecinfosListArrayOutput {
	return o.ApplyT(func(v GetSpecinfosResult) []GetSpecinfosList { return v.Lists }).(GetSpecinfosListArrayOutput)
}

func (o GetSpecinfosResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSpecinfosResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSpecinfosResultOutput{})
}
