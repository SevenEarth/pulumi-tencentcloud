// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a CAM SAML provider.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cam"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cam.NewSamlProvider(ctx, "samlProviderBasic", &Cam.SamlProviderArgs{
// 			Description: pulumi.String("test"),
// 			MetaData:    pulumi.String("PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48bWQ6RW50aXR5RGVzY3JpcHRvciBlbnRpdHlJRD0iaHR0cDovL3d3dy5va3RhLmNvbS9leGsxa3F4bWNqUW1HQURNeTM1NyIgeG1sbnM6bWQ9InVybjpvYXNpczpuYW1lczp0YzpTQU1MOjIuMDptZXRhZGF0YSI+PG1kOklEUFNTT0Rlc2NyaXB0b3IgV2FudEF1dGhuUmVxdWVzdHNTaWduZWQ9ImZhbHNlIiBwcm90b2NvbFN1cHBvcnRFbnVtZXJhdGlvbj0idXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6Mi4wOnByb3RvY29sIj48bWQ6S2V5RGVzY3JpcHRvciB1c2U9InNpZ25pbmciPjxkczpLZXlJbmZvIHhtbG5zOmRzPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwLzA5L3htbGRzaWcjIj48ZHM6WDUwOURhdGE+PGRzOlg1MDlDZXJ0aWZpY2F0ZT5NSUlEb0RDQ0FvaWdBd0lCQWdJR0FXM0lTcExvTUEwR0NTcUdTSWIzRFFFQkN3VUFNSUdRTVFzd0NRWURWUVFHRXdKVlV6RVRNQkVHDQpBMVVFQ0F3S1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ3d05VMkZ1SUVaeVlXNWphWE5qYnpFTk1Bc0dBMVVFQ2d3RVQydDBZVEVVDQpNQklHQTFVRUN3d0xVMU5QVUhKdmRtbGtaWEl4RVRBUEJnTlZCQU1NQ0dsa2VIVmxkblJoTVJ3d0dnWUpLb1pJaHZjTkFRa0JGZzFwDQpibVp2UUc5cmRHRXVZMjl0TUI0WERURTVNVEF4TkRBek1qSXhNMW9YRFRJNU1UQXhOREF6TWpNeE0xb3dnWkF4Q3pBSkJnTlZCQVlUDQpBbFZUTVJNd0VRWURWUVFJREFwRFlXeHBabTl5Ym1saE1SWXdGQVlEVlFRSERBMVRZVzRnUm5KaGJtTnBjMk52TVEwd0N3WURWUVFLDQpEQVJQYTNSaE1SUXdFZ1lEVlFRTERBdFRVMDlRY205MmFXUmxjakVSTUE4R0ExVUVBd3dJYVdSNGRXVjJkR0V4SERBYUJna3Foa2lHDQo5dzBCQ1FFV0RXbHVabTlBYjJ0MFlTNWpiMjB3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLQW9JQkFRQ2g4b3dqDQpZK2dQSUM3blQvNTduLzdmeXJzcDlHMXdxa2UxdXhjMHVrTndnQXozOVNpelY3QVhLMWRReTFLaThXWjJJMzFEczJkT0FNQ1FKR2pWDQpUWWNNbnA3KzhqUzNLdmxNUkRJamk5cmxuUi9vcnBvMll1RHVWby9jVzdidlRIS2h2REo1QWZRaWxzYlNPTXdUOWM2TVlYZGhBNVBwDQpzelFsK1UrdHJmcXUrdUorSER4SVQxdlhWaVI5YlY2SUFRSzZpbWZoc2wxWmVSUytjbVFVNEpjQWlYT0xtTnFVVWM2UkpxUzhrMW1mDQpBLzhmb2VyMGc3SG4xZDVXclpCc2gyUlR2Vzh1ZVdadHQ3dmh4QTlGdE5kSVlEcXJ0eElmMlZXcXhrSHM3WFZDSm5wTnJITVovT1BRDQpGY21YSGVxNlJJMlB3Q1RlOW8zZHZpM0hqeXBaOEl4dkFnTUJBQUV3RFFZSktvWklodmNOQVFFTEJRQURnZ0VCQUFHaHk1bG9nbGtTDQoyVHg2YS90MnF5VEx0YVV5cEwrNGhySGJoMVAweVVMc0NrSnFsM2wrWG9VZDZCY2FJaFNSVGFPQk95ODViL0UzelJ4K3JzQXJwTjVVDQp5ZThuUEM4a05PYW5vTk9wWnZvYmhpTzFlMFIvYmxEcnRBL0o5UlBwMWtmdlhmS2NSTTU3TlRCWXppTURlbnFQUTRFOWN1U2lGdFFxDQpJYmpIbThaM1B1YXgwRitldkZ3U1pJMDNCWXNISGw1d1EraEJBS3hTdTJINEZRdU93Zmpnb2EveEN6Z1NKYjJ2UXdEc1MxMk9mSkNiDQpSRm1ZL1VYZXQramFhdEVORktLZStZSUJpU0J2WG1adTN0MHN5NDZTNzlPVzBacXJ0NUh2bElsT2lpTFpaN1FZamxjM1kxeG1LZ1luDQpXM2M2WGZkdmhGWHo0ZDdkbWYvTUdpNGY0enM9PC9kczpYNTA5Q2VydGlmaWNhdGU+PC9kczpYNTA5RGF0YT48L2RzOktleUluZm8+PC9tZDpLZXlEZXNjcmlwdG9yPjxtZDpOYW1lSURGb3JtYXQ+dXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6MS4xOm5hbWVpZC1mb3JtYXQ6dW5zcGVjaWZpZWQ8L21kOk5hbWVJREZvcm1hdD48bWQ6TmFtZUlERm9ybWF0PnVybjpvYXNpczpuYW1lczp0YzpTQU1MOjEuMTpuYW1laWQtZm9ybWF0OmVtYWlsQWRkcmVzczwvbWQ6TmFtZUlERm9ybWF0PjxtZDpTaW5nbGVTaWduT25TZXJ2aWNlIEJpbmRpbmc9InVybjpvYXNpczpuYW1lczp0YzpTQU1MOjIuMDpiaW5kaW5nczpIVFRQLVBPU1QiIExvY2F0aW9uPSJodHRwczovL2lkeHVldnRhLm9rdGEuY29tL2FwcC9pZHh1ZW9yZzYzNzM1OF90ZXN0XzEvZXhrMWtxeG1jalFtR0FETXkzNTcvc3NvL3NhbWwiLz48bWQ6U2luZ2xlU2lnbk9uU2VydmljZSBCaW5kaW5nPSJ1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6YmluZGluZ3M6SFRUUC1SZWRpcmVjdCIgTG9jYXRpb249Imh0dHBzOi8vaWR4dWV2dGEub2t0YS5jb20vYXBwL2lkeHVlb3JnNjM3MzU4X3Rlc3RfMS9leGsxa3F4bWNqUW1HQURNeTM1Ny9zc28vc2FtbCIvPjwvbWQ6SURQU1NPRGVzY3JpcHRvcj48L21kOkVudGl0eURlc2NyaXB0b3I+"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// CAM SAML provider can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Cam/samlProvider:SamlProvider foo cam-SAML-provider-test
// ```
type SamlProvider struct {
	pulumi.CustomResourceState

	// The create time of the CAM SAML provider.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The description of the CAM SAML provider.
	Description pulumi.StringOutput `pulumi:"description"`
	// The meta data document of the CAM SAML provider.
	MetaData pulumi.StringOutput `pulumi:"metaData"`
	// Name of CAM SAML provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARN of the CAM SAML provider.
	ProviderArn pulumi.StringOutput `pulumi:"providerArn"`
	// The last update time of the CAM SAML provider.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSamlProvider registers a new resource with the given unique name, arguments, and options.
func NewSamlProvider(ctx *pulumi.Context,
	name string, args *SamlProviderArgs, opts ...pulumi.ResourceOption) (*SamlProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.MetaData == nil {
		return nil, errors.New("invalid value for required argument 'MetaData'")
	}
	var resource SamlProvider
	err := ctx.RegisterResource("tencentcloud:Cam/samlProvider:SamlProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSamlProvider gets an existing SamlProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSamlProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SamlProviderState, opts ...pulumi.ResourceOption) (*SamlProvider, error) {
	var resource SamlProvider
	err := ctx.ReadResource("tencentcloud:Cam/samlProvider:SamlProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SamlProvider resources.
type samlProviderState struct {
	// The create time of the CAM SAML provider.
	CreateTime *string `pulumi:"createTime"`
	// The description of the CAM SAML provider.
	Description *string `pulumi:"description"`
	// The meta data document of the CAM SAML provider.
	MetaData *string `pulumi:"metaData"`
	// Name of CAM SAML provider.
	Name *string `pulumi:"name"`
	// The ARN of the CAM SAML provider.
	ProviderArn *string `pulumi:"providerArn"`
	// The last update time of the CAM SAML provider.
	UpdateTime *string `pulumi:"updateTime"`
}

type SamlProviderState struct {
	// The create time of the CAM SAML provider.
	CreateTime pulumi.StringPtrInput
	// The description of the CAM SAML provider.
	Description pulumi.StringPtrInput
	// The meta data document of the CAM SAML provider.
	MetaData pulumi.StringPtrInput
	// Name of CAM SAML provider.
	Name pulumi.StringPtrInput
	// The ARN of the CAM SAML provider.
	ProviderArn pulumi.StringPtrInput
	// The last update time of the CAM SAML provider.
	UpdateTime pulumi.StringPtrInput
}

func (SamlProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*samlProviderState)(nil)).Elem()
}

type samlProviderArgs struct {
	// The description of the CAM SAML provider.
	Description string `pulumi:"description"`
	// The meta data document of the CAM SAML provider.
	MetaData string `pulumi:"metaData"`
	// Name of CAM SAML provider.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a SamlProvider resource.
type SamlProviderArgs struct {
	// The description of the CAM SAML provider.
	Description pulumi.StringInput
	// The meta data document of the CAM SAML provider.
	MetaData pulumi.StringInput
	// Name of CAM SAML provider.
	Name pulumi.StringPtrInput
}

func (SamlProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*samlProviderArgs)(nil)).Elem()
}

type SamlProviderInput interface {
	pulumi.Input

	ToSamlProviderOutput() SamlProviderOutput
	ToSamlProviderOutputWithContext(ctx context.Context) SamlProviderOutput
}

func (*SamlProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**SamlProvider)(nil)).Elem()
}

func (i *SamlProvider) ToSamlProviderOutput() SamlProviderOutput {
	return i.ToSamlProviderOutputWithContext(context.Background())
}

func (i *SamlProvider) ToSamlProviderOutputWithContext(ctx context.Context) SamlProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlProviderOutput)
}

// SamlProviderArrayInput is an input type that accepts SamlProviderArray and SamlProviderArrayOutput values.
// You can construct a concrete instance of `SamlProviderArrayInput` via:
//
//          SamlProviderArray{ SamlProviderArgs{...} }
type SamlProviderArrayInput interface {
	pulumi.Input

	ToSamlProviderArrayOutput() SamlProviderArrayOutput
	ToSamlProviderArrayOutputWithContext(context.Context) SamlProviderArrayOutput
}

type SamlProviderArray []SamlProviderInput

func (SamlProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SamlProvider)(nil)).Elem()
}

func (i SamlProviderArray) ToSamlProviderArrayOutput() SamlProviderArrayOutput {
	return i.ToSamlProviderArrayOutputWithContext(context.Background())
}

func (i SamlProviderArray) ToSamlProviderArrayOutputWithContext(ctx context.Context) SamlProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlProviderArrayOutput)
}

// SamlProviderMapInput is an input type that accepts SamlProviderMap and SamlProviderMapOutput values.
// You can construct a concrete instance of `SamlProviderMapInput` via:
//
//          SamlProviderMap{ "key": SamlProviderArgs{...} }
type SamlProviderMapInput interface {
	pulumi.Input

	ToSamlProviderMapOutput() SamlProviderMapOutput
	ToSamlProviderMapOutputWithContext(context.Context) SamlProviderMapOutput
}

type SamlProviderMap map[string]SamlProviderInput

func (SamlProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SamlProvider)(nil)).Elem()
}

func (i SamlProviderMap) ToSamlProviderMapOutput() SamlProviderMapOutput {
	return i.ToSamlProviderMapOutputWithContext(context.Background())
}

func (i SamlProviderMap) ToSamlProviderMapOutputWithContext(ctx context.Context) SamlProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlProviderMapOutput)
}

type SamlProviderOutput struct{ *pulumi.OutputState }

func (SamlProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SamlProvider)(nil)).Elem()
}

func (o SamlProviderOutput) ToSamlProviderOutput() SamlProviderOutput {
	return o
}

func (o SamlProviderOutput) ToSamlProviderOutputWithContext(ctx context.Context) SamlProviderOutput {
	return o
}

// The create time of the CAM SAML provider.
func (o SamlProviderOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The description of the CAM SAML provider.
func (o SamlProviderOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The meta data document of the CAM SAML provider.
func (o SamlProviderOutput) MetaData() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.MetaData }).(pulumi.StringOutput)
}

// Name of CAM SAML provider.
func (o SamlProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ARN of the CAM SAML provider.
func (o SamlProviderOutput) ProviderArn() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.ProviderArn }).(pulumi.StringOutput)
}

// The last update time of the CAM SAML provider.
func (o SamlProviderOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type SamlProviderArrayOutput struct{ *pulumi.OutputState }

func (SamlProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SamlProvider)(nil)).Elem()
}

func (o SamlProviderArrayOutput) ToSamlProviderArrayOutput() SamlProviderArrayOutput {
	return o
}

func (o SamlProviderArrayOutput) ToSamlProviderArrayOutputWithContext(ctx context.Context) SamlProviderArrayOutput {
	return o
}

func (o SamlProviderArrayOutput) Index(i pulumi.IntInput) SamlProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SamlProvider {
		return vs[0].([]*SamlProvider)[vs[1].(int)]
	}).(SamlProviderOutput)
}

type SamlProviderMapOutput struct{ *pulumi.OutputState }

func (SamlProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SamlProvider)(nil)).Elem()
}

func (o SamlProviderMapOutput) ToSamlProviderMapOutput() SamlProviderMapOutput {
	return o
}

func (o SamlProviderMapOutput) ToSamlProviderMapOutputWithContext(ctx context.Context) SamlProviderMapOutput {
	return o
}

func (o SamlProviderMapOutput) MapIndex(k pulumi.StringInput) SamlProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SamlProvider {
		return vs[0].(map[string]*SamlProvider)[vs[1].(string)]
	}).(SamlProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SamlProviderInput)(nil)).Elem(), &SamlProvider{})
	pulumi.RegisterInputType(reflect.TypeOf((*SamlProviderArrayInput)(nil)).Elem(), SamlProviderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SamlProviderMapInput)(nil)).Elem(), SamlProviderMap{})
	pulumi.RegisterOutputType(SamlProviderOutput{})
	pulumi.RegisterOutputType(SamlProviderArrayOutput{})
	pulumi.RegisterOutputType(SamlProviderMapOutput{})
}
