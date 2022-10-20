// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query gaap layer7 listeners.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Gaap"
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
//			_ = Gaap.GetLayer7ListenersOutput(ctx, gaap.GetLayer7ListenersOutputArgs{
//				Protocol:   pulumi.String("HTTP"),
//				ProxyId:    fooProxy.ID(),
//				ListenerId: fooLayer7Listener.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetLayer7Listeners(ctx *pulumi.Context, args *GetLayer7ListenersArgs, opts ...pulumi.InvokeOption) (*GetLayer7ListenersResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetLayer7ListenersResult
	err := ctx.Invoke("tencentcloud:Gaap/getLayer7Listeners:getLayer7Listeners", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLayer7Listeners.
type GetLayer7ListenersArgs struct {
	// ID of the layer7 listener to be queried.
	ListenerId *string `pulumi:"listenerId"`
	// Name of the layer7 listener to be queried.
	ListenerName *string `pulumi:"listenerName"`
	// Port of the layer7 listener to be queried.
	Port *int `pulumi:"port"`
	// Protocol of the layer7 listener to be queried. Valid values: `HTTP` and `HTTPS`.
	Protocol string `pulumi:"protocol"`
	// ID of the GAAP proxy to be queried.
	ProxyId *string `pulumi:"proxyId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getLayer7Listeners.
type GetLayer7ListenersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id           string  `pulumi:"id"`
	ListenerId   *string `pulumi:"listenerId"`
	ListenerName *string `pulumi:"listenerName"`
	// An information list of layer7 listeners. Each element contains the following attributes:
	Listeners []GetLayer7ListenersListener `pulumi:"listeners"`
	// Port of the layer7 listener.
	Port *int `pulumi:"port"`
	// Protocol of the layer7 listener.
	Protocol         string  `pulumi:"protocol"`
	ProxyId          *string `pulumi:"proxyId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetLayer7ListenersOutput(ctx *pulumi.Context, args GetLayer7ListenersOutputArgs, opts ...pulumi.InvokeOption) GetLayer7ListenersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLayer7ListenersResult, error) {
			args := v.(GetLayer7ListenersArgs)
			r, err := GetLayer7Listeners(ctx, &args, opts...)
			var s GetLayer7ListenersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLayer7ListenersResultOutput)
}

// A collection of arguments for invoking getLayer7Listeners.
type GetLayer7ListenersOutputArgs struct {
	// ID of the layer7 listener to be queried.
	ListenerId pulumi.StringPtrInput `pulumi:"listenerId"`
	// Name of the layer7 listener to be queried.
	ListenerName pulumi.StringPtrInput `pulumi:"listenerName"`
	// Port of the layer7 listener to be queried.
	Port pulumi.IntPtrInput `pulumi:"port"`
	// Protocol of the layer7 listener to be queried. Valid values: `HTTP` and `HTTPS`.
	Protocol pulumi.StringInput `pulumi:"protocol"`
	// ID of the GAAP proxy to be queried.
	ProxyId pulumi.StringPtrInput `pulumi:"proxyId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetLayer7ListenersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLayer7ListenersArgs)(nil)).Elem()
}

// A collection of values returned by getLayer7Listeners.
type GetLayer7ListenersResultOutput struct{ *pulumi.OutputState }

func (GetLayer7ListenersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLayer7ListenersResult)(nil)).Elem()
}

func (o GetLayer7ListenersResultOutput) ToGetLayer7ListenersResultOutput() GetLayer7ListenersResultOutput {
	return o
}

func (o GetLayer7ListenersResultOutput) ToGetLayer7ListenersResultOutputWithContext(ctx context.Context) GetLayer7ListenersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetLayer7ListenersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLayer7ListenersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLayer7ListenersResultOutput) ListenerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLayer7ListenersResult) *string { return v.ListenerId }).(pulumi.StringPtrOutput)
}

func (o GetLayer7ListenersResultOutput) ListenerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLayer7ListenersResult) *string { return v.ListenerName }).(pulumi.StringPtrOutput)
}

// An information list of layer7 listeners. Each element contains the following attributes:
func (o GetLayer7ListenersResultOutput) Listeners() GetLayer7ListenersListenerArrayOutput {
	return o.ApplyT(func(v GetLayer7ListenersResult) []GetLayer7ListenersListener { return v.Listeners }).(GetLayer7ListenersListenerArrayOutput)
}

// Port of the layer7 listener.
func (o GetLayer7ListenersResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetLayer7ListenersResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

// Protocol of the layer7 listener.
func (o GetLayer7ListenersResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v GetLayer7ListenersResult) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o GetLayer7ListenersResultOutput) ProxyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLayer7ListenersResult) *string { return v.ProxyId }).(pulumi.StringPtrOutput)
}

func (o GetLayer7ListenersResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLayer7ListenersResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLayer7ListenersResultOutput{})
}
