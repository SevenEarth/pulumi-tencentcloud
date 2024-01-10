// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of apigateway plugin
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/ApiGateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/ApiGateway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleService, err := ApiGateway.NewService(ctx, "exampleService", &ApiGateway.ServiceArgs{
//				ServiceName: pulumi.String("tf_example"),
//				Protocol:    pulumi.String("http&https"),
//				ServiceDesc: pulumi.String("desc."),
//				NetTypes: pulumi.StringArray{
//					pulumi.String("INNER"),
//					pulumi.String("OUTER"),
//				},
//				IpVersion: pulumi.String("IPv4"),
//				Tags: pulumi.AnyMap{
//					"testKey": pulumi.Any("testValue"),
//				},
//				ReleaseLimit: pulumi.Int(500),
//				PreLimit:     pulumi.Int(500),
//				TestLimit:    pulumi.Int(500),
//			})
//			if err != nil {
//				return err
//			}
//			exampleApi, err := ApiGateway.NewApi(ctx, "exampleApi", &ApiGateway.ApiArgs{
//				ServiceId:           exampleService.ID(),
//				ApiName:             pulumi.String("hello"),
//				ApiDesc:             pulumi.String("my hello api"),
//				AuthType:            pulumi.String("NONE"),
//				Protocol:            pulumi.String("HTTP"),
//				EnableCors:          pulumi.Bool(true),
//				RequestConfigPath:   pulumi.String("/user/info"),
//				RequestConfigMethod: pulumi.String("GET"),
//				RequestParameters: apigateway.ApiRequestParameterArray{
//					&apigateway.ApiRequestParameterArgs{
//						Name:         pulumi.String("name"),
//						Position:     pulumi.String("QUERY"),
//						Type:         pulumi.String("string"),
//						Desc:         pulumi.String("who are you?"),
//						DefaultValue: pulumi.String("tom"),
//						Required:     pulumi.Bool(true),
//					},
//				},
//				ServiceConfigType:      pulumi.String("HTTP"),
//				ServiceConfigTimeout:   pulumi.Int(15),
//				ServiceConfigUrl:       pulumi.String("http://www.qq.com"),
//				ServiceConfigPath:      pulumi.String("/user"),
//				ServiceConfigMethod:    pulumi.String("GET"),
//				ResponseType:           pulumi.String("HTML"),
//				ResponseSuccessExample: pulumi.String("success"),
//				ResponseFailExample:    pulumi.String("fail"),
//				ResponseErrorCodes: apigateway.ApiResponseErrorCodeArray{
//					&apigateway.ApiResponseErrorCodeArgs{
//						Code:          pulumi.Int(500),
//						Msg:           pulumi.String("system error"),
//						Desc:          pulumi.String("system error code"),
//						ConvertedCode: pulumi.Int(5000),
//						NeedConvert:   pulumi.Bool(true),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleServiceRelease, err := ApiGateway.NewServiceRelease(ctx, "exampleServiceRelease", &ApiGateway.ServiceReleaseArgs{
//				ServiceId:       exampleApi.ServiceId,
//				EnvironmentName: pulumi.String("release"),
//				ReleaseDesc:     pulumi.String("desc."),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"type":   "white_list",
//				"blocks": "1.1.1.1",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			examplePlugin, err := ApiGateway.NewPlugin(ctx, "examplePlugin", &ApiGateway.PluginArgs{
//				PluginName:  pulumi.String("tf-example"),
//				PluginType:  pulumi.String("IPControl"),
//				PluginData:  pulumi.String(json0),
//				Description: pulumi.String("desc."),
//			})
//			if err != nil {
//				return err
//			}
//			_ = ApiGateway.GetPluginsOutput(ctx, apigateway.GetPluginsOutputArgs{
//				ServiceId:       exampleServiceRelease.ServiceId,
//				PluginId:        examplePlugin.ID(),
//				EnvironmentName: pulumi.String("release"),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupPlugins(ctx *pulumi.Context, args *LookupPluginsArgs, opts ...pulumi.InvokeOption) (*LookupPluginsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupPluginsResult
	err := ctx.Invoke("tencentcloud:ApiGateway/getPlugins:getPlugins", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPlugins.
type LookupPluginsArgs struct {
	// Environmental information.
	EnvironmentName string `pulumi:"environmentName"`
	// The plugin ID to query.
	PluginId string `pulumi:"pluginId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// The service ID to query.
	ServiceId string `pulumi:"serviceId"`
}

// A collection of values returned by getPlugins.
type LookupPluginsResult struct {
	EnvironmentName string `pulumi:"environmentName"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	PluginId         string  `pulumi:"pluginId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// List of plugin related APIs.
	Results   []GetPluginsResult `pulumi:"results"`
	ServiceId string             `pulumi:"serviceId"`
}

func LookupPluginsOutput(ctx *pulumi.Context, args LookupPluginsOutputArgs, opts ...pulumi.InvokeOption) LookupPluginsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPluginsResult, error) {
			args := v.(LookupPluginsArgs)
			r, err := LookupPlugins(ctx, &args, opts...)
			var s LookupPluginsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPluginsResultOutput)
}

// A collection of arguments for invoking getPlugins.
type LookupPluginsOutputArgs struct {
	// Environmental information.
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
	// The plugin ID to query.
	PluginId pulumi.StringInput `pulumi:"pluginId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// The service ID to query.
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (LookupPluginsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPluginsArgs)(nil)).Elem()
}

// A collection of values returned by getPlugins.
type LookupPluginsResultOutput struct{ *pulumi.OutputState }

func (LookupPluginsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPluginsResult)(nil)).Elem()
}

func (o LookupPluginsResultOutput) ToLookupPluginsResultOutput() LookupPluginsResultOutput {
	return o
}

func (o LookupPluginsResultOutput) ToLookupPluginsResultOutputWithContext(ctx context.Context) LookupPluginsResultOutput {
	return o
}

func (o LookupPluginsResultOutput) EnvironmentName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPluginsResult) string { return v.EnvironmentName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPluginsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPluginsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPluginsResultOutput) PluginId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPluginsResult) string { return v.PluginId }).(pulumi.StringOutput)
}

func (o LookupPluginsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPluginsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// List of plugin related APIs.
func (o LookupPluginsResultOutput) Results() GetPluginsResultArrayOutput {
	return o.ApplyT(func(v LookupPluginsResult) []GetPluginsResult { return v.Results }).(GetPluginsResultArrayOutput)
}

func (o LookupPluginsResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPluginsResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPluginsResultOutput{})
}
