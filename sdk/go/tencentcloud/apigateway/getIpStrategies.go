// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query API gateway IP strategy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/ApiGateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/ApiGateway"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		service, err := ApiGateway.NewService(ctx, "service", &ApiGateway.ServiceArgs{
// 			ServiceName: pulumi.String("ck"),
// 			Protocol:    pulumi.String("http&https"),
// 			ServiceDesc: pulumi.String("your nice service"),
// 			NetTypes: pulumi.StringArray{
// 				pulumi.String("INNER"),
// 				pulumi.String("OUTER"),
// 			},
// 			IpVersion: pulumi.String("IPv4"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		test, err := ApiGateway.NewIpStrategy(ctx, "test", &ApiGateway.IpStrategyArgs{
// 			ServiceId:    service.ID(),
// 			StrategyName: pulumi.String("tf_test"),
// 			StrategyType: pulumi.String("BLACK"),
// 			StrategyData: pulumi.String("9.9.9.9"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_ = ApiGateway.GetIpStrategiesOutput(ctx, apigateway.GetIpStrategiesOutputArgs{
// 			ServiceId: test.ServiceId,
// 		}, nil)
// 		_ = ApiGateway.GetIpStrategiesOutput(ctx, apigateway.GetIpStrategiesOutputArgs{
// 			ServiceId:    test.ServiceId,
// 			StrategyName: test.StrategyName,
// 		}, nil)
// 		return nil
// 	})
// }
// ```
func GetIpStrategies(ctx *pulumi.Context, args *GetIpStrategiesArgs, opts ...pulumi.InvokeOption) (*GetIpStrategiesResult, error) {
	var rv GetIpStrategiesResult
	err := ctx.Invoke("tencentcloud:ApiGateway/getIpStrategies:getIpStrategies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpStrategies.
type GetIpStrategiesArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// The service ID to be queried.
	ServiceId string `pulumi:"serviceId"`
	// Name of IP policy.
	StrategyName *string `pulumi:"strategyName"`
}

// A collection of values returned by getIpStrategies.
type GetIpStrategiesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of strategy.
	Lists            []GetIpStrategiesList `pulumi:"lists"`
	ResultOutputFile *string               `pulumi:"resultOutputFile"`
	// The service ID.
	ServiceId string `pulumi:"serviceId"`
	// Name of the strategy.
	StrategyName *string `pulumi:"strategyName"`
}

func GetIpStrategiesOutput(ctx *pulumi.Context, args GetIpStrategiesOutputArgs, opts ...pulumi.InvokeOption) GetIpStrategiesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIpStrategiesResult, error) {
			args := v.(GetIpStrategiesArgs)
			r, err := GetIpStrategies(ctx, &args, opts...)
			var s GetIpStrategiesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIpStrategiesResultOutput)
}

// A collection of arguments for invoking getIpStrategies.
type GetIpStrategiesOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// The service ID to be queried.
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
	// Name of IP policy.
	StrategyName pulumi.StringPtrInput `pulumi:"strategyName"`
}

func (GetIpStrategiesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpStrategiesArgs)(nil)).Elem()
}

// A collection of values returned by getIpStrategies.
type GetIpStrategiesResultOutput struct{ *pulumi.OutputState }

func (GetIpStrategiesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpStrategiesResult)(nil)).Elem()
}

func (o GetIpStrategiesResultOutput) ToGetIpStrategiesResultOutput() GetIpStrategiesResultOutput {
	return o
}

func (o GetIpStrategiesResultOutput) ToGetIpStrategiesResultOutputWithContext(ctx context.Context) GetIpStrategiesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpStrategiesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpStrategiesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of strategy.
func (o GetIpStrategiesResultOutput) Lists() GetIpStrategiesListArrayOutput {
	return o.ApplyT(func(v GetIpStrategiesResult) []GetIpStrategiesList { return v.Lists }).(GetIpStrategiesListArrayOutput)
}

func (o GetIpStrategiesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpStrategiesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// The service ID.
func (o GetIpStrategiesResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpStrategiesResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

// Name of the strategy.
func (o GetIpStrategiesResultOutput) StrategyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpStrategiesResult) *string { return v.StrategyName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpStrategiesResultOutput{})
}
