// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a forward domain of layer7 listener.
//
// ## Example Usage
//
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
//			_, err = Gaap.NewHttpDomain(ctx, "fooHttpDomain", &Gaap.HttpDomainArgs{
//				ListenerId: fooLayer7Listener.ID(),
//				Domain:     pulumi.String("www.qq.com"),
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
// GAAP http domain can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Gaap/httpDomain:HttpDomain tencentcloud_gaap_http_domain.foo listener-11112222+HTTP+www.qq.com
//
// ```
type HttpDomain struct {
	pulumi.CustomResourceState

	// Indicates whether basic authentication is enable, default value is `false`.
	BasicAuth pulumi.BoolPtrOutput `pulumi:"basicAuth"`
	// ID of the basic authentication.
	BasicAuthId pulumi.StringOutput `pulumi:"basicAuthId"`
	// ID of the server certificate, default value is `default`.
	CertificateId pulumi.StringPtrOutput `pulumi:"certificateId"`
	// It has been deprecated from version 1.26.0. Set `clientCertificateIds` instead. ID of the client certificate, default value is `default`.
	//
	// Deprecated: It has been deprecated from version 1.26.0. Set `client_certificate_ids` instead.
	ClientCertificateId pulumi.StringOutput `pulumi:"clientCertificateId"`
	// ID list of the poly client certificate.
	ClientCertificateIds pulumi.StringArrayOutput `pulumi:"clientCertificateIds"`
	// Forward domain of the layer7 listener.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Indicates whether SSL certificate authentication is enable, default value is `false`.
	GaapAuth pulumi.BoolPtrOutput `pulumi:"gaapAuth"`
	// ID of the SSL certificate.
	GaapAuthId pulumi.StringOutput `pulumi:"gaapAuthId"`
	// ID of the layer7 listener.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// Indicates whether realserver authentication is enable, default value is `false`.
	RealserverAuth pulumi.BoolPtrOutput `pulumi:"realserverAuth"`
	// CA certificate domain of the realserver. It has been deprecated.
	RealserverCertificateDomain pulumi.StringOutput `pulumi:"realserverCertificateDomain"`
	// It has been deprecated from version 1.28.0. Set `realserverCertificateIds` instead. CA certificate ID of the realserver.
	//
	// Deprecated: It has been deprecated from version 1.28.0. Set `realserver_certificate_ids` instead.
	RealserverCertificateId pulumi.StringOutput `pulumi:"realserverCertificateId"`
	// CA certificate ID list of the realserver.
	RealserverCertificateIds pulumi.StringArrayOutput `pulumi:"realserverCertificateIds"`
}

// NewHttpDomain registers a new resource with the given unique name, arguments, and options.
func NewHttpDomain(ctx *pulumi.Context,
	name string, args *HttpDomainArgs, opts ...pulumi.ResourceOption) (*HttpDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.ListenerId == nil {
		return nil, errors.New("invalid value for required argument 'ListenerId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource HttpDomain
	err := ctx.RegisterResource("tencentcloud:Gaap/httpDomain:HttpDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHttpDomain gets an existing HttpDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHttpDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HttpDomainState, opts ...pulumi.ResourceOption) (*HttpDomain, error) {
	var resource HttpDomain
	err := ctx.ReadResource("tencentcloud:Gaap/httpDomain:HttpDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HttpDomain resources.
type httpDomainState struct {
	// Indicates whether basic authentication is enable, default value is `false`.
	BasicAuth *bool `pulumi:"basicAuth"`
	// ID of the basic authentication.
	BasicAuthId *string `pulumi:"basicAuthId"`
	// ID of the server certificate, default value is `default`.
	CertificateId *string `pulumi:"certificateId"`
	// It has been deprecated from version 1.26.0. Set `clientCertificateIds` instead. ID of the client certificate, default value is `default`.
	//
	// Deprecated: It has been deprecated from version 1.26.0. Set `client_certificate_ids` instead.
	ClientCertificateId *string `pulumi:"clientCertificateId"`
	// ID list of the poly client certificate.
	ClientCertificateIds []string `pulumi:"clientCertificateIds"`
	// Forward domain of the layer7 listener.
	Domain *string `pulumi:"domain"`
	// Indicates whether SSL certificate authentication is enable, default value is `false`.
	GaapAuth *bool `pulumi:"gaapAuth"`
	// ID of the SSL certificate.
	GaapAuthId *string `pulumi:"gaapAuthId"`
	// ID of the layer7 listener.
	ListenerId *string `pulumi:"listenerId"`
	// Indicates whether realserver authentication is enable, default value is `false`.
	RealserverAuth *bool `pulumi:"realserverAuth"`
	// CA certificate domain of the realserver. It has been deprecated.
	RealserverCertificateDomain *string `pulumi:"realserverCertificateDomain"`
	// It has been deprecated from version 1.28.0. Set `realserverCertificateIds` instead. CA certificate ID of the realserver.
	//
	// Deprecated: It has been deprecated from version 1.28.0. Set `realserver_certificate_ids` instead.
	RealserverCertificateId *string `pulumi:"realserverCertificateId"`
	// CA certificate ID list of the realserver.
	RealserverCertificateIds []string `pulumi:"realserverCertificateIds"`
}

type HttpDomainState struct {
	// Indicates whether basic authentication is enable, default value is `false`.
	BasicAuth pulumi.BoolPtrInput
	// ID of the basic authentication.
	BasicAuthId pulumi.StringPtrInput
	// ID of the server certificate, default value is `default`.
	CertificateId pulumi.StringPtrInput
	// It has been deprecated from version 1.26.0. Set `clientCertificateIds` instead. ID of the client certificate, default value is `default`.
	//
	// Deprecated: It has been deprecated from version 1.26.0. Set `client_certificate_ids` instead.
	ClientCertificateId pulumi.StringPtrInput
	// ID list of the poly client certificate.
	ClientCertificateIds pulumi.StringArrayInput
	// Forward domain of the layer7 listener.
	Domain pulumi.StringPtrInput
	// Indicates whether SSL certificate authentication is enable, default value is `false`.
	GaapAuth pulumi.BoolPtrInput
	// ID of the SSL certificate.
	GaapAuthId pulumi.StringPtrInput
	// ID of the layer7 listener.
	ListenerId pulumi.StringPtrInput
	// Indicates whether realserver authentication is enable, default value is `false`.
	RealserverAuth pulumi.BoolPtrInput
	// CA certificate domain of the realserver. It has been deprecated.
	RealserverCertificateDomain pulumi.StringPtrInput
	// It has been deprecated from version 1.28.0. Set `realserverCertificateIds` instead. CA certificate ID of the realserver.
	//
	// Deprecated: It has been deprecated from version 1.28.0. Set `realserver_certificate_ids` instead.
	RealserverCertificateId pulumi.StringPtrInput
	// CA certificate ID list of the realserver.
	RealserverCertificateIds pulumi.StringArrayInput
}

func (HttpDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*httpDomainState)(nil)).Elem()
}

type httpDomainArgs struct {
	// Indicates whether basic authentication is enable, default value is `false`.
	BasicAuth *bool `pulumi:"basicAuth"`
	// ID of the basic authentication.
	BasicAuthId *string `pulumi:"basicAuthId"`
	// ID of the server certificate, default value is `default`.
	CertificateId *string `pulumi:"certificateId"`
	// It has been deprecated from version 1.26.0. Set `clientCertificateIds` instead. ID of the client certificate, default value is `default`.
	//
	// Deprecated: It has been deprecated from version 1.26.0. Set `client_certificate_ids` instead.
	ClientCertificateId *string `pulumi:"clientCertificateId"`
	// ID list of the poly client certificate.
	ClientCertificateIds []string `pulumi:"clientCertificateIds"`
	// Forward domain of the layer7 listener.
	Domain string `pulumi:"domain"`
	// Indicates whether SSL certificate authentication is enable, default value is `false`.
	GaapAuth *bool `pulumi:"gaapAuth"`
	// ID of the SSL certificate.
	GaapAuthId *string `pulumi:"gaapAuthId"`
	// ID of the layer7 listener.
	ListenerId string `pulumi:"listenerId"`
	// Indicates whether realserver authentication is enable, default value is `false`.
	RealserverAuth *bool `pulumi:"realserverAuth"`
	// CA certificate domain of the realserver. It has been deprecated.
	RealserverCertificateDomain *string `pulumi:"realserverCertificateDomain"`
	// It has been deprecated from version 1.28.0. Set `realserverCertificateIds` instead. CA certificate ID of the realserver.
	//
	// Deprecated: It has been deprecated from version 1.28.0. Set `realserver_certificate_ids` instead.
	RealserverCertificateId *string `pulumi:"realserverCertificateId"`
	// CA certificate ID list of the realserver.
	RealserverCertificateIds []string `pulumi:"realserverCertificateIds"`
}

// The set of arguments for constructing a HttpDomain resource.
type HttpDomainArgs struct {
	// Indicates whether basic authentication is enable, default value is `false`.
	BasicAuth pulumi.BoolPtrInput
	// ID of the basic authentication.
	BasicAuthId pulumi.StringPtrInput
	// ID of the server certificate, default value is `default`.
	CertificateId pulumi.StringPtrInput
	// It has been deprecated from version 1.26.0. Set `clientCertificateIds` instead. ID of the client certificate, default value is `default`.
	//
	// Deprecated: It has been deprecated from version 1.26.0. Set `client_certificate_ids` instead.
	ClientCertificateId pulumi.StringPtrInput
	// ID list of the poly client certificate.
	ClientCertificateIds pulumi.StringArrayInput
	// Forward domain of the layer7 listener.
	Domain pulumi.StringInput
	// Indicates whether SSL certificate authentication is enable, default value is `false`.
	GaapAuth pulumi.BoolPtrInput
	// ID of the SSL certificate.
	GaapAuthId pulumi.StringPtrInput
	// ID of the layer7 listener.
	ListenerId pulumi.StringInput
	// Indicates whether realserver authentication is enable, default value is `false`.
	RealserverAuth pulumi.BoolPtrInput
	// CA certificate domain of the realserver. It has been deprecated.
	RealserverCertificateDomain pulumi.StringPtrInput
	// It has been deprecated from version 1.28.0. Set `realserverCertificateIds` instead. CA certificate ID of the realserver.
	//
	// Deprecated: It has been deprecated from version 1.28.0. Set `realserver_certificate_ids` instead.
	RealserverCertificateId pulumi.StringPtrInput
	// CA certificate ID list of the realserver.
	RealserverCertificateIds pulumi.StringArrayInput
}

func (HttpDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*httpDomainArgs)(nil)).Elem()
}

type HttpDomainInput interface {
	pulumi.Input

	ToHttpDomainOutput() HttpDomainOutput
	ToHttpDomainOutputWithContext(ctx context.Context) HttpDomainOutput
}

func (*HttpDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpDomain)(nil)).Elem()
}

func (i *HttpDomain) ToHttpDomainOutput() HttpDomainOutput {
	return i.ToHttpDomainOutputWithContext(context.Background())
}

func (i *HttpDomain) ToHttpDomainOutputWithContext(ctx context.Context) HttpDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpDomainOutput)
}

// HttpDomainArrayInput is an input type that accepts HttpDomainArray and HttpDomainArrayOutput values.
// You can construct a concrete instance of `HttpDomainArrayInput` via:
//
//	HttpDomainArray{ HttpDomainArgs{...} }
type HttpDomainArrayInput interface {
	pulumi.Input

	ToHttpDomainArrayOutput() HttpDomainArrayOutput
	ToHttpDomainArrayOutputWithContext(context.Context) HttpDomainArrayOutput
}

type HttpDomainArray []HttpDomainInput

func (HttpDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HttpDomain)(nil)).Elem()
}

func (i HttpDomainArray) ToHttpDomainArrayOutput() HttpDomainArrayOutput {
	return i.ToHttpDomainArrayOutputWithContext(context.Background())
}

func (i HttpDomainArray) ToHttpDomainArrayOutputWithContext(ctx context.Context) HttpDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpDomainArrayOutput)
}

// HttpDomainMapInput is an input type that accepts HttpDomainMap and HttpDomainMapOutput values.
// You can construct a concrete instance of `HttpDomainMapInput` via:
//
//	HttpDomainMap{ "key": HttpDomainArgs{...} }
type HttpDomainMapInput interface {
	pulumi.Input

	ToHttpDomainMapOutput() HttpDomainMapOutput
	ToHttpDomainMapOutputWithContext(context.Context) HttpDomainMapOutput
}

type HttpDomainMap map[string]HttpDomainInput

func (HttpDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HttpDomain)(nil)).Elem()
}

func (i HttpDomainMap) ToHttpDomainMapOutput() HttpDomainMapOutput {
	return i.ToHttpDomainMapOutputWithContext(context.Background())
}

func (i HttpDomainMap) ToHttpDomainMapOutputWithContext(ctx context.Context) HttpDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpDomainMapOutput)
}

type HttpDomainOutput struct{ *pulumi.OutputState }

func (HttpDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpDomain)(nil)).Elem()
}

func (o HttpDomainOutput) ToHttpDomainOutput() HttpDomainOutput {
	return o
}

func (o HttpDomainOutput) ToHttpDomainOutputWithContext(ctx context.Context) HttpDomainOutput {
	return o
}

// Indicates whether basic authentication is enable, default value is `false`.
func (o HttpDomainOutput) BasicAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HttpDomain) pulumi.BoolPtrOutput { return v.BasicAuth }).(pulumi.BoolPtrOutput)
}

// ID of the basic authentication.
func (o HttpDomainOutput) BasicAuthId() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpDomain) pulumi.StringOutput { return v.BasicAuthId }).(pulumi.StringOutput)
}

// ID of the server certificate, default value is `default`.
func (o HttpDomainOutput) CertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpDomain) pulumi.StringPtrOutput { return v.CertificateId }).(pulumi.StringPtrOutput)
}

// It has been deprecated from version 1.26.0. Set `clientCertificateIds` instead. ID of the client certificate, default value is `default`.
//
// Deprecated: It has been deprecated from version 1.26.0. Set `client_certificate_ids` instead.
func (o HttpDomainOutput) ClientCertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpDomain) pulumi.StringOutput { return v.ClientCertificateId }).(pulumi.StringOutput)
}

// ID list of the poly client certificate.
func (o HttpDomainOutput) ClientCertificateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HttpDomain) pulumi.StringArrayOutput { return v.ClientCertificateIds }).(pulumi.StringArrayOutput)
}

// Forward domain of the layer7 listener.
func (o HttpDomainOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpDomain) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Indicates whether SSL certificate authentication is enable, default value is `false`.
func (o HttpDomainOutput) GaapAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HttpDomain) pulumi.BoolPtrOutput { return v.GaapAuth }).(pulumi.BoolPtrOutput)
}

// ID of the SSL certificate.
func (o HttpDomainOutput) GaapAuthId() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpDomain) pulumi.StringOutput { return v.GaapAuthId }).(pulumi.StringOutput)
}

// ID of the layer7 listener.
func (o HttpDomainOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpDomain) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

// Indicates whether realserver authentication is enable, default value is `false`.
func (o HttpDomainOutput) RealserverAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HttpDomain) pulumi.BoolPtrOutput { return v.RealserverAuth }).(pulumi.BoolPtrOutput)
}

// CA certificate domain of the realserver. It has been deprecated.
func (o HttpDomainOutput) RealserverCertificateDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpDomain) pulumi.StringOutput { return v.RealserverCertificateDomain }).(pulumi.StringOutput)
}

// It has been deprecated from version 1.28.0. Set `realserverCertificateIds` instead. CA certificate ID of the realserver.
//
// Deprecated: It has been deprecated from version 1.28.0. Set `realserver_certificate_ids` instead.
func (o HttpDomainOutput) RealserverCertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpDomain) pulumi.StringOutput { return v.RealserverCertificateId }).(pulumi.StringOutput)
}

// CA certificate ID list of the realserver.
func (o HttpDomainOutput) RealserverCertificateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HttpDomain) pulumi.StringArrayOutput { return v.RealserverCertificateIds }).(pulumi.StringArrayOutput)
}

type HttpDomainArrayOutput struct{ *pulumi.OutputState }

func (HttpDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HttpDomain)(nil)).Elem()
}

func (o HttpDomainArrayOutput) ToHttpDomainArrayOutput() HttpDomainArrayOutput {
	return o
}

func (o HttpDomainArrayOutput) ToHttpDomainArrayOutputWithContext(ctx context.Context) HttpDomainArrayOutput {
	return o
}

func (o HttpDomainArrayOutput) Index(i pulumi.IntInput) HttpDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HttpDomain {
		return vs[0].([]*HttpDomain)[vs[1].(int)]
	}).(HttpDomainOutput)
}

type HttpDomainMapOutput struct{ *pulumi.OutputState }

func (HttpDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HttpDomain)(nil)).Elem()
}

func (o HttpDomainMapOutput) ToHttpDomainMapOutput() HttpDomainMapOutput {
	return o
}

func (o HttpDomainMapOutput) ToHttpDomainMapOutputWithContext(ctx context.Context) HttpDomainMapOutput {
	return o
}

func (o HttpDomainMapOutput) MapIndex(k pulumi.StringInput) HttpDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HttpDomain {
		return vs[0].(map[string]*HttpDomain)[vs[1].(string)]
	}).(HttpDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HttpDomainInput)(nil)).Elem(), &HttpDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpDomainArrayInput)(nil)).Elem(), HttpDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpDomainMapInput)(nil)).Elem(), HttpDomainMap{})
	pulumi.RegisterOutputType(HttpDomainOutput{})
	pulumi.RegisterOutputType(HttpDomainArrayOutput{})
	pulumi.RegisterOutputType(HttpDomainMapOutput{})
}
