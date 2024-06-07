// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query API gateway throttling services.
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/ApiGateway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			service, err := ApiGateway.NewService(ctx, "service", &ApiGateway.ServiceArgs{
//				ServiceName: pulumi.String("niceservice"),
//				Protocol:    pulumi.String("http&https"),
//				ServiceDesc: pulumi.String("your nice service"),
//				NetTypes: pulumi.StringArray{
//					pulumi.String("INNER"),
//					pulumi.String("OUTER"),
//				},
//				IpVersion:    pulumi.String("IPv4"),
//				ReleaseLimit: pulumi.Int(100),
//				PreLimit:     pulumi.Int(100),
//				TestLimit:    pulumi.Int(100),
//			})
//			if err != nil {
//				return err
//			}
//			_ = ApiGateway.GetThrottlingServicesOutput(ctx, apigateway.GetThrottlingServicesOutputArgs{
//				ServiceId: service.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetThrottlingServices(ctx *pulumi.Context, args *GetThrottlingServicesArgs, opts ...pulumi.InvokeOption) (*GetThrottlingServicesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetThrottlingServicesResult
	err := ctx.Invoke("tencentcloud:ApiGateway/getThrottlingServices:getThrottlingServices", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getThrottlingServices.
type GetThrottlingServicesArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Service ID for query.
	ServiceId *string `pulumi:"serviceId"`
}

// A collection of values returned by getThrottlingServices.
type GetThrottlingServicesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Throttling policy.
	Lists            []GetThrottlingServicesList `pulumi:"lists"`
	ResultOutputFile *string                     `pulumi:"resultOutputFile"`
	// Service ID for query.
	ServiceId *string `pulumi:"serviceId"`
}

func GetThrottlingServicesOutput(ctx *pulumi.Context, args GetThrottlingServicesOutputArgs, opts ...pulumi.InvokeOption) GetThrottlingServicesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetThrottlingServicesResult, error) {
			args := v.(GetThrottlingServicesArgs)
			r, err := GetThrottlingServices(ctx, &args, opts...)
			var s GetThrottlingServicesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetThrottlingServicesResultOutput)
}

// A collection of arguments for invoking getThrottlingServices.
type GetThrottlingServicesOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Service ID for query.
	ServiceId pulumi.StringPtrInput `pulumi:"serviceId"`
}

func (GetThrottlingServicesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetThrottlingServicesArgs)(nil)).Elem()
}

// A collection of values returned by getThrottlingServices.
type GetThrottlingServicesResultOutput struct{ *pulumi.OutputState }

func (GetThrottlingServicesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetThrottlingServicesResult)(nil)).Elem()
}

func (o GetThrottlingServicesResultOutput) ToGetThrottlingServicesResultOutput() GetThrottlingServicesResultOutput {
	return o
}

func (o GetThrottlingServicesResultOutput) ToGetThrottlingServicesResultOutputWithContext(ctx context.Context) GetThrottlingServicesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetThrottlingServicesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetThrottlingServicesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of Throttling policy.
func (o GetThrottlingServicesResultOutput) Lists() GetThrottlingServicesListArrayOutput {
	return o.ApplyT(func(v GetThrottlingServicesResult) []GetThrottlingServicesList { return v.Lists }).(GetThrottlingServicesListArrayOutput)
}

func (o GetThrottlingServicesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetThrottlingServicesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Service ID for query.
func (o GetThrottlingServicesResultOutput) ServiceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetThrottlingServicesResult) *string { return v.ServiceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetThrottlingServicesResultOutput{})
}
