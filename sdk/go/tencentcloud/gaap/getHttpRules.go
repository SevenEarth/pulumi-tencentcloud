// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query forward rule of layer7 listeners.
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Gaap"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooProxy, err := Gaap.NewProxy(ctx, "fooProxy", &Gaap.ProxyArgs{
//				Bandwidth:        pulumi.Int(10),
//				Concurrent:       pulumi.Int(2),
//				AccessRegion:     pulumi.String("SouthChina"),
//				RealserverRegion: pulumi.String("NorthChina"),
//			})
//			if err != nil {
//				return err
//			}
//			fooLayer7Listener, err := Gaap.NewLayer7Listener(ctx, "fooLayer7Listener", &Gaap.Layer7ListenerArgs{
//				Protocol: pulumi.String("HTTP"),
//				Port:     pulumi.Int(80),
//				ProxyId:  fooProxy.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			fooRealserver, err := Gaap.NewRealserver(ctx, "fooRealserver", &Gaap.RealserverArgs{
//				Ip: pulumi.String("1.1.1.1"),
//			})
//			if err != nil {
//				return err
//			}
//			fooHttpRule, err := Gaap.NewHttpRule(ctx, "fooHttpRule", &Gaap.HttpRuleArgs{
//				ListenerId:     fooLayer7Listener.ID(),
//				Domain:         pulumi.String("www.qq.com"),
//				Path:           pulumi.String("/"),
//				RealserverType: pulumi.String("IP"),
//				HealthCheck:    pulumi.Bool(true),
//				Realservers: gaap.HttpRuleRealserverArray{
//					&gaap.HttpRuleRealserverArgs{
//						Id:   fooRealserver.ID(),
//						Ip:   fooRealserver.Ip,
//						Port: pulumi.Int(80),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = Gaap.GetHttpRulesOutput(ctx, gaap.GetHttpRulesOutputArgs{
//				ListenerId: fooLayer7Listener.ID(),
//				Domain:     fooHttpRule.Domain,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetHttpRules(ctx *pulumi.Context, args *GetHttpRulesArgs, opts ...pulumi.InvokeOption) (*GetHttpRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetHttpRulesResult
	err := ctx.Invoke("tencentcloud:Gaap/getHttpRules:getHttpRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHttpRules.
type GetHttpRulesArgs struct {
	// Forward domain of the layer7 listener to be queried.
	Domain *string `pulumi:"domain"`
	// Requested host which is forwarded to the realserver by the listener to be queried.
	ForwardHost *string `pulumi:"forwardHost"`
	// ID of the layer7 listener to be queried.
	ListenerId string `pulumi:"listenerId"`
	// Path of the forward rule to be queried.
	Path *string `pulumi:"path"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getHttpRules.
type GetHttpRulesResult struct {
	// Domain of the GAAP realserver.
	Domain *string `pulumi:"domain"`
	// Requested host which is forwarded to the realserver by the listener.
	ForwardHost *string `pulumi:"forwardHost"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the layer7 listener.
	ListenerId string `pulumi:"listenerId"`
	// Path of the forward rule.
	Path             *string `pulumi:"path"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// An information list of forward rule of the layer7 listeners. Each element contains the following attributes:
	Rules []GetHttpRulesRule `pulumi:"rules"`
}

func GetHttpRulesOutput(ctx *pulumi.Context, args GetHttpRulesOutputArgs, opts ...pulumi.InvokeOption) GetHttpRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetHttpRulesResult, error) {
			args := v.(GetHttpRulesArgs)
			r, err := GetHttpRules(ctx, &args, opts...)
			var s GetHttpRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetHttpRulesResultOutput)
}

// A collection of arguments for invoking getHttpRules.
type GetHttpRulesOutputArgs struct {
	// Forward domain of the layer7 listener to be queried.
	Domain pulumi.StringPtrInput `pulumi:"domain"`
	// Requested host which is forwarded to the realserver by the listener to be queried.
	ForwardHost pulumi.StringPtrInput `pulumi:"forwardHost"`
	// ID of the layer7 listener to be queried.
	ListenerId pulumi.StringInput `pulumi:"listenerId"`
	// Path of the forward rule to be queried.
	Path pulumi.StringPtrInput `pulumi:"path"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetHttpRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHttpRulesArgs)(nil)).Elem()
}

// A collection of values returned by getHttpRules.
type GetHttpRulesResultOutput struct{ *pulumi.OutputState }

func (GetHttpRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHttpRulesResult)(nil)).Elem()
}

func (o GetHttpRulesResultOutput) ToGetHttpRulesResultOutput() GetHttpRulesResultOutput {
	return o
}

func (o GetHttpRulesResultOutput) ToGetHttpRulesResultOutputWithContext(ctx context.Context) GetHttpRulesResultOutput {
	return o
}

// Domain of the GAAP realserver.
func (o GetHttpRulesResultOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHttpRulesResult) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

// Requested host which is forwarded to the realserver by the listener.
func (o GetHttpRulesResultOutput) ForwardHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHttpRulesResult) *string { return v.ForwardHost }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetHttpRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetHttpRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the layer7 listener.
func (o GetHttpRulesResultOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetHttpRulesResult) string { return v.ListenerId }).(pulumi.StringOutput)
}

// Path of the forward rule.
func (o GetHttpRulesResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHttpRulesResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o GetHttpRulesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHttpRulesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// An information list of forward rule of the layer7 listeners. Each element contains the following attributes:
func (o GetHttpRulesResultOutput) Rules() GetHttpRulesRuleArrayOutput {
	return o.ApplyT(func(v GetHttpRulesResult) []GetHttpRulesRule { return v.Rules }).(GetHttpRulesRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetHttpRulesResultOutput{})
}
