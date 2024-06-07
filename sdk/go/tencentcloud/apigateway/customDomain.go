// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this resource to create custom domain of API gateway.
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
//			_, err := ApiGateway.NewCustomDomain(ctx, "foo", &ApiGateway.CustomDomainArgs{
//				DefaultDomain:    pulumi.String("service-ohxqslqe-1259649581.gz.apigw.tencentcs.com"),
//				IsDefaultMapping: pulumi.Bool(false),
//				NetType:          pulumi.String("OUTER"),
//				PathMappings: pulumi.StringArray{
//					pulumi.String("/good#test"),
//					pulumi.String("/root#release"),
//				},
//				Protocol:  pulumi.String("http"),
//				ServiceId: pulumi.String("service-ohxqslqe"),
//				SubDomain: pulumi.String("tic-test.dnsv1.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
type CustomDomain struct {
	pulumi.CustomResourceState

	// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&https`.
	CertificateId pulumi.StringOutput `pulumi:"certificateId"`
	// Default domain name.
	DefaultDomain pulumi.StringOutput `pulumi:"defaultDomain"`
	// Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `pathMappings` attribute is required.
	IsDefaultMapping pulumi.BoolPtrOutput `pulumi:"isDefaultMapping"`
	// Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
	IsForcedHttps pulumi.BoolPtrOutput `pulumi:"isForcedHttps"`
	// Network type. Valid values: `OUTER`, `INNER`.
	NetType pulumi.StringOutput `pulumi:"netType"`
	// Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
	PathMappings pulumi.StringArrayOutput `pulumi:"pathMappings"`
	// Protocol supported by service. Valid values: `http`, `https`, `http&https`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Unique service ID.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// Domain name resolution status. `1` means normal analysis, `0` means parsing failed.
	Status pulumi.IntOutput `pulumi:"status"`
	// Custom domain name to be bound.
	SubDomain pulumi.StringOutput `pulumi:"subDomain"`
}

// NewCustomDomain registers a new resource with the given unique name, arguments, and options.
func NewCustomDomain(ctx *pulumi.Context,
	name string, args *CustomDomainArgs, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultDomain == nil {
		return nil, errors.New("invalid value for required argument 'DefaultDomain'")
	}
	if args.NetType == nil {
		return nil, errors.New("invalid value for required argument 'NetType'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	if args.SubDomain == nil {
		return nil, errors.New("invalid value for required argument 'SubDomain'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CustomDomain
	err := ctx.RegisterResource("tencentcloud:ApiGateway/customDomain:CustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomDomain gets an existing CustomDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomDomainState, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	var resource CustomDomain
	err := ctx.ReadResource("tencentcloud:ApiGateway/customDomain:CustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomDomain resources.
type customDomainState struct {
	// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&https`.
	CertificateId *string `pulumi:"certificateId"`
	// Default domain name.
	DefaultDomain *string `pulumi:"defaultDomain"`
	// Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `pathMappings` attribute is required.
	IsDefaultMapping *bool `pulumi:"isDefaultMapping"`
	// Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
	IsForcedHttps *bool `pulumi:"isForcedHttps"`
	// Network type. Valid values: `OUTER`, `INNER`.
	NetType *string `pulumi:"netType"`
	// Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
	PathMappings []string `pulumi:"pathMappings"`
	// Protocol supported by service. Valid values: `http`, `https`, `http&https`.
	Protocol *string `pulumi:"protocol"`
	// Unique service ID.
	ServiceId *string `pulumi:"serviceId"`
	// Domain name resolution status. `1` means normal analysis, `0` means parsing failed.
	Status *int `pulumi:"status"`
	// Custom domain name to be bound.
	SubDomain *string `pulumi:"subDomain"`
}

type CustomDomainState struct {
	// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&https`.
	CertificateId pulumi.StringPtrInput
	// Default domain name.
	DefaultDomain pulumi.StringPtrInput
	// Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `pathMappings` attribute is required.
	IsDefaultMapping pulumi.BoolPtrInput
	// Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
	IsForcedHttps pulumi.BoolPtrInput
	// Network type. Valid values: `OUTER`, `INNER`.
	NetType pulumi.StringPtrInput
	// Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
	PathMappings pulumi.StringArrayInput
	// Protocol supported by service. Valid values: `http`, `https`, `http&https`.
	Protocol pulumi.StringPtrInput
	// Unique service ID.
	ServiceId pulumi.StringPtrInput
	// Domain name resolution status. `1` means normal analysis, `0` means parsing failed.
	Status pulumi.IntPtrInput
	// Custom domain name to be bound.
	SubDomain pulumi.StringPtrInput
}

func (CustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainState)(nil)).Elem()
}

type customDomainArgs struct {
	// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&https`.
	CertificateId *string `pulumi:"certificateId"`
	// Default domain name.
	DefaultDomain string `pulumi:"defaultDomain"`
	// Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `pathMappings` attribute is required.
	IsDefaultMapping *bool `pulumi:"isDefaultMapping"`
	// Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
	IsForcedHttps *bool `pulumi:"isForcedHttps"`
	// Network type. Valid values: `OUTER`, `INNER`.
	NetType string `pulumi:"netType"`
	// Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
	PathMappings []string `pulumi:"pathMappings"`
	// Protocol supported by service. Valid values: `http`, `https`, `http&https`.
	Protocol string `pulumi:"protocol"`
	// Unique service ID.
	ServiceId string `pulumi:"serviceId"`
	// Custom domain name to be bound.
	SubDomain string `pulumi:"subDomain"`
}

// The set of arguments for constructing a CustomDomain resource.
type CustomDomainArgs struct {
	// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&https`.
	CertificateId pulumi.StringPtrInput
	// Default domain name.
	DefaultDomain pulumi.StringInput
	// Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `pathMappings` attribute is required.
	IsDefaultMapping pulumi.BoolPtrInput
	// Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
	IsForcedHttps pulumi.BoolPtrInput
	// Network type. Valid values: `OUTER`, `INNER`.
	NetType pulumi.StringInput
	// Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
	PathMappings pulumi.StringArrayInput
	// Protocol supported by service. Valid values: `http`, `https`, `http&https`.
	Protocol pulumi.StringInput
	// Unique service ID.
	ServiceId pulumi.StringInput
	// Custom domain name to be bound.
	SubDomain pulumi.StringInput
}

func (CustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainArgs)(nil)).Elem()
}

type CustomDomainInput interface {
	pulumi.Input

	ToCustomDomainOutput() CustomDomainOutput
	ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput
}

func (*CustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (i *CustomDomain) ToCustomDomainOutput() CustomDomainOutput {
	return i.ToCustomDomainOutputWithContext(context.Background())
}

func (i *CustomDomain) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput)
}

// CustomDomainArrayInput is an input type that accepts CustomDomainArray and CustomDomainArrayOutput values.
// You can construct a concrete instance of `CustomDomainArrayInput` via:
//
//	CustomDomainArray{ CustomDomainArgs{...} }
type CustomDomainArrayInput interface {
	pulumi.Input

	ToCustomDomainArrayOutput() CustomDomainArrayOutput
	ToCustomDomainArrayOutputWithContext(context.Context) CustomDomainArrayOutput
}

type CustomDomainArray []CustomDomainInput

func (CustomDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomDomain)(nil)).Elem()
}

func (i CustomDomainArray) ToCustomDomainArrayOutput() CustomDomainArrayOutput {
	return i.ToCustomDomainArrayOutputWithContext(context.Background())
}

func (i CustomDomainArray) ToCustomDomainArrayOutputWithContext(ctx context.Context) CustomDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainArrayOutput)
}

// CustomDomainMapInput is an input type that accepts CustomDomainMap and CustomDomainMapOutput values.
// You can construct a concrete instance of `CustomDomainMapInput` via:
//
//	CustomDomainMap{ "key": CustomDomainArgs{...} }
type CustomDomainMapInput interface {
	pulumi.Input

	ToCustomDomainMapOutput() CustomDomainMapOutput
	ToCustomDomainMapOutputWithContext(context.Context) CustomDomainMapOutput
}

type CustomDomainMap map[string]CustomDomainInput

func (CustomDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomDomain)(nil)).Elem()
}

func (i CustomDomainMap) ToCustomDomainMapOutput() CustomDomainMapOutput {
	return i.ToCustomDomainMapOutputWithContext(context.Background())
}

func (i CustomDomainMap) ToCustomDomainMapOutputWithContext(ctx context.Context) CustomDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainMapOutput)
}

type CustomDomainOutput struct{ *pulumi.OutputState }

func (CustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (o CustomDomainOutput) ToCustomDomainOutput() CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return o
}

// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&https`.
func (o CustomDomainOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.CertificateId }).(pulumi.StringOutput)
}

// Default domain name.
func (o CustomDomainOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.DefaultDomain }).(pulumi.StringOutput)
}

// Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `pathMappings` attribute is required.
func (o CustomDomainOutput) IsDefaultMapping() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.BoolPtrOutput { return v.IsDefaultMapping }).(pulumi.BoolPtrOutput)
}

// Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
func (o CustomDomainOutput) IsForcedHttps() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.BoolPtrOutput { return v.IsForcedHttps }).(pulumi.BoolPtrOutput)
}

// Network type. Valid values: `OUTER`, `INNER`.
func (o CustomDomainOutput) NetType() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.NetType }).(pulumi.StringOutput)
}

// Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
func (o CustomDomainOutput) PathMappings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringArrayOutput { return v.PathMappings }).(pulumi.StringArrayOutput)
}

// Protocol supported by service. Valid values: `http`, `https`, `http&https`.
func (o CustomDomainOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Unique service ID.
func (o CustomDomainOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// Domain name resolution status. `1` means normal analysis, `0` means parsing failed.
func (o CustomDomainOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// Custom domain name to be bound.
func (o CustomDomainOutput) SubDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.SubDomain }).(pulumi.StringOutput)
}

type CustomDomainArrayOutput struct{ *pulumi.OutputState }

func (CustomDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomDomain)(nil)).Elem()
}

func (o CustomDomainArrayOutput) ToCustomDomainArrayOutput() CustomDomainArrayOutput {
	return o
}

func (o CustomDomainArrayOutput) ToCustomDomainArrayOutputWithContext(ctx context.Context) CustomDomainArrayOutput {
	return o
}

func (o CustomDomainArrayOutput) Index(i pulumi.IntInput) CustomDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomDomain {
		return vs[0].([]*CustomDomain)[vs[1].(int)]
	}).(CustomDomainOutput)
}

type CustomDomainMapOutput struct{ *pulumi.OutputState }

func (CustomDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomDomain)(nil)).Elem()
}

func (o CustomDomainMapOutput) ToCustomDomainMapOutput() CustomDomainMapOutput {
	return o
}

func (o CustomDomainMapOutput) ToCustomDomainMapOutputWithContext(ctx context.Context) CustomDomainMapOutput {
	return o
}

func (o CustomDomainMapOutput) MapIndex(k pulumi.StringInput) CustomDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomDomain {
		return vs[0].(map[string]*CustomDomain)[vs[1].(string)]
	}).(CustomDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDomainInput)(nil)).Elem(), &CustomDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDomainArrayInput)(nil)).Elem(), CustomDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDomainMapInput)(nil)).Elem(), CustomDomainMap{})
	pulumi.RegisterOutputType(CustomDomainOutput{})
	pulumi.RegisterOutputType(CustomDomainArrayOutput{})
	pulumi.RegisterOutputType(CustomDomainMapOutput{})
}
