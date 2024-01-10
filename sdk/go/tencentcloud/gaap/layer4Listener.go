// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a layer4 listener of GAAP.
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
//			fooRealserver, err := Gaap.NewRealserver(ctx, "fooRealserver", &Gaap.RealserverArgs{
//				Ip: pulumi.String("1.1.1.1"),
//			})
//			if err != nil {
//				return err
//			}
//			bar, err := Gaap.NewRealserver(ctx, "bar", &Gaap.RealserverArgs{
//				Ip: pulumi.String("119.29.29.29"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Gaap.NewLayer4Listener(ctx, "fooLayer4Listener", &Gaap.Layer4ListenerArgs{
//				Protocol:       pulumi.String("TCP"),
//				Port:           pulumi.Int(80),
//				RealserverType: pulumi.String("IP"),
//				ProxyId:        fooProxy.ID(),
//				HealthCheck:    pulumi.Bool(true),
//				RealserverBindSets: gaap.Layer4ListenerRealserverBindSetArray{
//					&gaap.Layer4ListenerRealserverBindSetArgs{
//						Id:   fooRealserver.ID(),
//						Ip:   fooRealserver.Ip,
//						Port: pulumi.Int(80),
//					},
//					&gaap.Layer4ListenerRealserverBindSetArgs{
//						Id:   bar.ID(),
//						Ip:   bar.Ip,
//						Port: pulumi.Int(80),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// GAAP layer4 listener can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Gaap/layer4Listener:Layer4Listener tencentcloud_gaap_layer4_listener.foo listener-11112222
//
// ```
type Layer4Listener struct {
	pulumi.CustomResourceState

	// The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports listeners of `TCP` protocol.
	ClientIpMethod pulumi.IntPtrOutput `pulumi:"clientIpMethod"`
	// Timeout of the health check response, should less than interval, default value is 2s. NOTES: Only supports listeners of `TCP` protocol and require less than `interval`.
	ConnectTimeout pulumi.IntPtrOutput `pulumi:"connectTimeout"`
	// Creation time of the layer4 listener.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Indicates whether health check is enable, default value is `false`. NOTES: Only supports listeners of `TCP` protocol.
	HealthCheck pulumi.BoolPtrOutput `pulumi:"healthCheck"`
	// Interval of the health check, default value is 5s. NOTES: Only supports listeners of `TCP` protocol.
	Interval pulumi.IntPtrOutput `pulumi:"interval"`
	// Name of the layer4 listener, the maximum length is 30.
	Name pulumi.StringOutput `pulumi:"name"`
	// Port of the layer4 listener.
	Port pulumi.IntOutput `pulumi:"port"`
	// Protocol of the layer4 listener. Valid value: `TCP` and `UDP`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// ID of the GAAP proxy.
	ProxyId pulumi.StringOutput `pulumi:"proxyId"`
	// An information list of GAAP realserver.
	RealserverBindSets Layer4ListenerRealserverBindSetArrayOutput `pulumi:"realserverBindSets"`
	// Type of the realserver. Valid value: `IP` and `DOMAIN`. NOTES: when the `protocol` is specified as `TCP` and the `scheduler` is specified as `wrr`, the item can only be set to `IP`.
	RealserverType pulumi.StringOutput `pulumi:"realserverType"`
	// Scheduling policy of the layer4 listener, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler pulumi.StringPtrOutput `pulumi:"scheduler"`
	// Status of the layer4 listener.
	Status pulumi.IntOutput `pulumi:"status"`
}

// NewLayer4Listener registers a new resource with the given unique name, arguments, and options.
func NewLayer4Listener(ctx *pulumi.Context,
	name string, args *Layer4ListenerArgs, opts ...pulumi.ResourceOption) (*Layer4Listener, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ProxyId == nil {
		return nil, errors.New("invalid value for required argument 'ProxyId'")
	}
	if args.RealserverType == nil {
		return nil, errors.New("invalid value for required argument 'RealserverType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Layer4Listener
	err := ctx.RegisterResource("tencentcloud:Gaap/layer4Listener:Layer4Listener", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLayer4Listener gets an existing Layer4Listener resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLayer4Listener(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Layer4ListenerState, opts ...pulumi.ResourceOption) (*Layer4Listener, error) {
	var resource Layer4Listener
	err := ctx.ReadResource("tencentcloud:Gaap/layer4Listener:Layer4Listener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Layer4Listener resources.
type layer4ListenerState struct {
	// The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports listeners of `TCP` protocol.
	ClientIpMethod *int `pulumi:"clientIpMethod"`
	// Timeout of the health check response, should less than interval, default value is 2s. NOTES: Only supports listeners of `TCP` protocol and require less than `interval`.
	ConnectTimeout *int `pulumi:"connectTimeout"`
	// Creation time of the layer4 listener.
	CreateTime *string `pulumi:"createTime"`
	// Indicates whether health check is enable, default value is `false`. NOTES: Only supports listeners of `TCP` protocol.
	HealthCheck *bool `pulumi:"healthCheck"`
	// Interval of the health check, default value is 5s. NOTES: Only supports listeners of `TCP` protocol.
	Interval *int `pulumi:"interval"`
	// Name of the layer4 listener, the maximum length is 30.
	Name *string `pulumi:"name"`
	// Port of the layer4 listener.
	Port *int `pulumi:"port"`
	// Protocol of the layer4 listener. Valid value: `TCP` and `UDP`.
	Protocol *string `pulumi:"protocol"`
	// ID of the GAAP proxy.
	ProxyId *string `pulumi:"proxyId"`
	// An information list of GAAP realserver.
	RealserverBindSets []Layer4ListenerRealserverBindSet `pulumi:"realserverBindSets"`
	// Type of the realserver. Valid value: `IP` and `DOMAIN`. NOTES: when the `protocol` is specified as `TCP` and the `scheduler` is specified as `wrr`, the item can only be set to `IP`.
	RealserverType *string `pulumi:"realserverType"`
	// Scheduling policy of the layer4 listener, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler *string `pulumi:"scheduler"`
	// Status of the layer4 listener.
	Status *int `pulumi:"status"`
}

type Layer4ListenerState struct {
	// The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports listeners of `TCP` protocol.
	ClientIpMethod pulumi.IntPtrInput
	// Timeout of the health check response, should less than interval, default value is 2s. NOTES: Only supports listeners of `TCP` protocol and require less than `interval`.
	ConnectTimeout pulumi.IntPtrInput
	// Creation time of the layer4 listener.
	CreateTime pulumi.StringPtrInput
	// Indicates whether health check is enable, default value is `false`. NOTES: Only supports listeners of `TCP` protocol.
	HealthCheck pulumi.BoolPtrInput
	// Interval of the health check, default value is 5s. NOTES: Only supports listeners of `TCP` protocol.
	Interval pulumi.IntPtrInput
	// Name of the layer4 listener, the maximum length is 30.
	Name pulumi.StringPtrInput
	// Port of the layer4 listener.
	Port pulumi.IntPtrInput
	// Protocol of the layer4 listener. Valid value: `TCP` and `UDP`.
	Protocol pulumi.StringPtrInput
	// ID of the GAAP proxy.
	ProxyId pulumi.StringPtrInput
	// An information list of GAAP realserver.
	RealserverBindSets Layer4ListenerRealserverBindSetArrayInput
	// Type of the realserver. Valid value: `IP` and `DOMAIN`. NOTES: when the `protocol` is specified as `TCP` and the `scheduler` is specified as `wrr`, the item can only be set to `IP`.
	RealserverType pulumi.StringPtrInput
	// Scheduling policy of the layer4 listener, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler pulumi.StringPtrInput
	// Status of the layer4 listener.
	Status pulumi.IntPtrInput
}

func (Layer4ListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*layer4ListenerState)(nil)).Elem()
}

type layer4ListenerArgs struct {
	// The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports listeners of `TCP` protocol.
	ClientIpMethod *int `pulumi:"clientIpMethod"`
	// Timeout of the health check response, should less than interval, default value is 2s. NOTES: Only supports listeners of `TCP` protocol and require less than `interval`.
	ConnectTimeout *int `pulumi:"connectTimeout"`
	// Indicates whether health check is enable, default value is `false`. NOTES: Only supports listeners of `TCP` protocol.
	HealthCheck *bool `pulumi:"healthCheck"`
	// Interval of the health check, default value is 5s. NOTES: Only supports listeners of `TCP` protocol.
	Interval *int `pulumi:"interval"`
	// Name of the layer4 listener, the maximum length is 30.
	Name *string `pulumi:"name"`
	// Port of the layer4 listener.
	Port int `pulumi:"port"`
	// Protocol of the layer4 listener. Valid value: `TCP` and `UDP`.
	Protocol string `pulumi:"protocol"`
	// ID of the GAAP proxy.
	ProxyId string `pulumi:"proxyId"`
	// An information list of GAAP realserver.
	RealserverBindSets []Layer4ListenerRealserverBindSet `pulumi:"realserverBindSets"`
	// Type of the realserver. Valid value: `IP` and `DOMAIN`. NOTES: when the `protocol` is specified as `TCP` and the `scheduler` is specified as `wrr`, the item can only be set to `IP`.
	RealserverType string `pulumi:"realserverType"`
	// Scheduling policy of the layer4 listener, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler *string `pulumi:"scheduler"`
}

// The set of arguments for constructing a Layer4Listener resource.
type Layer4ListenerArgs struct {
	// The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports listeners of `TCP` protocol.
	ClientIpMethod pulumi.IntPtrInput
	// Timeout of the health check response, should less than interval, default value is 2s. NOTES: Only supports listeners of `TCP` protocol and require less than `interval`.
	ConnectTimeout pulumi.IntPtrInput
	// Indicates whether health check is enable, default value is `false`. NOTES: Only supports listeners of `TCP` protocol.
	HealthCheck pulumi.BoolPtrInput
	// Interval of the health check, default value is 5s. NOTES: Only supports listeners of `TCP` protocol.
	Interval pulumi.IntPtrInput
	// Name of the layer4 listener, the maximum length is 30.
	Name pulumi.StringPtrInput
	// Port of the layer4 listener.
	Port pulumi.IntInput
	// Protocol of the layer4 listener. Valid value: `TCP` and `UDP`.
	Protocol pulumi.StringInput
	// ID of the GAAP proxy.
	ProxyId pulumi.StringInput
	// An information list of GAAP realserver.
	RealserverBindSets Layer4ListenerRealserverBindSetArrayInput
	// Type of the realserver. Valid value: `IP` and `DOMAIN`. NOTES: when the `protocol` is specified as `TCP` and the `scheduler` is specified as `wrr`, the item can only be set to `IP`.
	RealserverType pulumi.StringInput
	// Scheduling policy of the layer4 listener, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler pulumi.StringPtrInput
}

func (Layer4ListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*layer4ListenerArgs)(nil)).Elem()
}

type Layer4ListenerInput interface {
	pulumi.Input

	ToLayer4ListenerOutput() Layer4ListenerOutput
	ToLayer4ListenerOutputWithContext(ctx context.Context) Layer4ListenerOutput
}

func (*Layer4Listener) ElementType() reflect.Type {
	return reflect.TypeOf((**Layer4Listener)(nil)).Elem()
}

func (i *Layer4Listener) ToLayer4ListenerOutput() Layer4ListenerOutput {
	return i.ToLayer4ListenerOutputWithContext(context.Background())
}

func (i *Layer4Listener) ToLayer4ListenerOutputWithContext(ctx context.Context) Layer4ListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Layer4ListenerOutput)
}

// Layer4ListenerArrayInput is an input type that accepts Layer4ListenerArray and Layer4ListenerArrayOutput values.
// You can construct a concrete instance of `Layer4ListenerArrayInput` via:
//
//	Layer4ListenerArray{ Layer4ListenerArgs{...} }
type Layer4ListenerArrayInput interface {
	pulumi.Input

	ToLayer4ListenerArrayOutput() Layer4ListenerArrayOutput
	ToLayer4ListenerArrayOutputWithContext(context.Context) Layer4ListenerArrayOutput
}

type Layer4ListenerArray []Layer4ListenerInput

func (Layer4ListenerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Layer4Listener)(nil)).Elem()
}

func (i Layer4ListenerArray) ToLayer4ListenerArrayOutput() Layer4ListenerArrayOutput {
	return i.ToLayer4ListenerArrayOutputWithContext(context.Background())
}

func (i Layer4ListenerArray) ToLayer4ListenerArrayOutputWithContext(ctx context.Context) Layer4ListenerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Layer4ListenerArrayOutput)
}

// Layer4ListenerMapInput is an input type that accepts Layer4ListenerMap and Layer4ListenerMapOutput values.
// You can construct a concrete instance of `Layer4ListenerMapInput` via:
//
//	Layer4ListenerMap{ "key": Layer4ListenerArgs{...} }
type Layer4ListenerMapInput interface {
	pulumi.Input

	ToLayer4ListenerMapOutput() Layer4ListenerMapOutput
	ToLayer4ListenerMapOutputWithContext(context.Context) Layer4ListenerMapOutput
}

type Layer4ListenerMap map[string]Layer4ListenerInput

func (Layer4ListenerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Layer4Listener)(nil)).Elem()
}

func (i Layer4ListenerMap) ToLayer4ListenerMapOutput() Layer4ListenerMapOutput {
	return i.ToLayer4ListenerMapOutputWithContext(context.Background())
}

func (i Layer4ListenerMap) ToLayer4ListenerMapOutputWithContext(ctx context.Context) Layer4ListenerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Layer4ListenerMapOutput)
}

type Layer4ListenerOutput struct{ *pulumi.OutputState }

func (Layer4ListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Layer4Listener)(nil)).Elem()
}

func (o Layer4ListenerOutput) ToLayer4ListenerOutput() Layer4ListenerOutput {
	return o
}

func (o Layer4ListenerOutput) ToLayer4ListenerOutputWithContext(ctx context.Context) Layer4ListenerOutput {
	return o
}

// The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports listeners of `TCP` protocol.
func (o Layer4ListenerOutput) ClientIpMethod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.IntPtrOutput { return v.ClientIpMethod }).(pulumi.IntPtrOutput)
}

// Timeout of the health check response, should less than interval, default value is 2s. NOTES: Only supports listeners of `TCP` protocol and require less than `interval`.
func (o Layer4ListenerOutput) ConnectTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.IntPtrOutput { return v.ConnectTimeout }).(pulumi.IntPtrOutput)
}

// Creation time of the layer4 listener.
func (o Layer4ListenerOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Indicates whether health check is enable, default value is `false`. NOTES: Only supports listeners of `TCP` protocol.
func (o Layer4ListenerOutput) HealthCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.BoolPtrOutput { return v.HealthCheck }).(pulumi.BoolPtrOutput)
}

// Interval of the health check, default value is 5s. NOTES: Only supports listeners of `TCP` protocol.
func (o Layer4ListenerOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

// Name of the layer4 listener, the maximum length is 30.
func (o Layer4ListenerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Port of the layer4 listener.
func (o Layer4ListenerOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Protocol of the layer4 listener. Valid value: `TCP` and `UDP`.
func (o Layer4ListenerOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// ID of the GAAP proxy.
func (o Layer4ListenerOutput) ProxyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringOutput { return v.ProxyId }).(pulumi.StringOutput)
}

// An information list of GAAP realserver.
func (o Layer4ListenerOutput) RealserverBindSets() Layer4ListenerRealserverBindSetArrayOutput {
	return o.ApplyT(func(v *Layer4Listener) Layer4ListenerRealserverBindSetArrayOutput { return v.RealserverBindSets }).(Layer4ListenerRealserverBindSetArrayOutput)
}

// Type of the realserver. Valid value: `IP` and `DOMAIN`. NOTES: when the `protocol` is specified as `TCP` and the `scheduler` is specified as `wrr`, the item can only be set to `IP`.
func (o Layer4ListenerOutput) RealserverType() pulumi.StringOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringOutput { return v.RealserverType }).(pulumi.StringOutput)
}

// Scheduling policy of the layer4 listener, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
func (o Layer4ListenerOutput) Scheduler() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringPtrOutput { return v.Scheduler }).(pulumi.StringPtrOutput)
}

// Status of the layer4 listener.
func (o Layer4ListenerOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

type Layer4ListenerArrayOutput struct{ *pulumi.OutputState }

func (Layer4ListenerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Layer4Listener)(nil)).Elem()
}

func (o Layer4ListenerArrayOutput) ToLayer4ListenerArrayOutput() Layer4ListenerArrayOutput {
	return o
}

func (o Layer4ListenerArrayOutput) ToLayer4ListenerArrayOutputWithContext(ctx context.Context) Layer4ListenerArrayOutput {
	return o
}

func (o Layer4ListenerArrayOutput) Index(i pulumi.IntInput) Layer4ListenerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Layer4Listener {
		return vs[0].([]*Layer4Listener)[vs[1].(int)]
	}).(Layer4ListenerOutput)
}

type Layer4ListenerMapOutput struct{ *pulumi.OutputState }

func (Layer4ListenerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Layer4Listener)(nil)).Elem()
}

func (o Layer4ListenerMapOutput) ToLayer4ListenerMapOutput() Layer4ListenerMapOutput {
	return o
}

func (o Layer4ListenerMapOutput) ToLayer4ListenerMapOutputWithContext(ctx context.Context) Layer4ListenerMapOutput {
	return o
}

func (o Layer4ListenerMapOutput) MapIndex(k pulumi.StringInput) Layer4ListenerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Layer4Listener {
		return vs[0].(map[string]*Layer4Listener)[vs[1].(string)]
	}).(Layer4ListenerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Layer4ListenerInput)(nil)).Elem(), &Layer4Listener{})
	pulumi.RegisterInputType(reflect.TypeOf((*Layer4ListenerArrayInput)(nil)).Elem(), Layer4ListenerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Layer4ListenerMapInput)(nil)).Elem(), Layer4ListenerMap{})
	pulumi.RegisterOutputType(Layer4ListenerOutput{})
	pulumi.RegisterOutputType(Layer4ListenerArrayOutput{})
	pulumi.RegisterOutputType(Layer4ListenerMapOutput{})
}
