// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create a dayu CC policy
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dayu"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dayu"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dayu.NewCcPolicyV2(ctx, "demo", &Dayu.CcPolicyV2Args{
//				Business: pulumi.String("bgpip"),
//				CcBlackWhiteIps: dayu.CcPolicyV2CcBlackWhiteIpArray{
//					&dayu.CcPolicyV2CcBlackWhiteIpArgs{
//						BlackWhiteIp: pulumi.String("1.2.3.4"),
//						Domain:       pulumi.String("12.com"),
//						Protocol:     pulumi.String("http"),
//						Type:         pulumi.String("black"),
//					},
//				},
//				CcGeoIpPolicys: dayu.CcPolicyV2CcGeoIpPolicyArray{
//					&dayu.CcPolicyV2CcGeoIpPolicyArgs{
//						Action:     pulumi.String("drop"),
//						Domain:     pulumi.String("12.com"),
//						Protocol:   pulumi.String("http"),
//						RegionType: pulumi.String("china"),
//					},
//				},
//				CcPrecisionPolicys: dayu.CcPolicyV2CcPrecisionPolicyArray{
//					&dayu.CcPolicyV2CcPrecisionPolicyArgs{
//						Domain:       pulumi.String("1.com"),
//						Ip:           pulumi.String("162.62.163.34"),
//						PolicyAction: pulumi.String("drop"),
//						Policys: dayu.CcPolicyV2CcPrecisionPolicyPolicyArray{
//							&dayu.CcPolicyV2CcPrecisionPolicyPolicyArgs{
//								FieldName:     pulumi.String("cgi"),
//								FieldType:     pulumi.String("value"),
//								Value:         pulumi.String("12123.com"),
//								ValueOperator: pulumi.String("equal"),
//							},
//						},
//						Protocol: pulumi.String("http"),
//					},
//				},
//				CcPrecisionReqLimits: dayu.CcPolicyV2CcPrecisionReqLimitArray{
//					&dayu.CcPolicyV2CcPrecisionReqLimitArgs{
//						Domain: pulumi.String("11.com"),
//						Level:  pulumi.String("loose"),
//						Policys: dayu.CcPolicyV2CcPrecisionReqLimitPolicyArray{
//							&dayu.CcPolicyV2CcPrecisionReqLimitPolicyArgs{
//								Action:          pulumi.String("alg"),
//								ExecuteDuration: pulumi.Int(2),
//								Mode:            pulumi.String("equal"),
//								Period:          pulumi.Int(5),
//								RequestNum:      pulumi.Int(12),
//								Uri:             pulumi.String("15.com"),
//							},
//						},
//						Protocol: pulumi.String("http"),
//					},
//				},
//				ResourceId: pulumi.String("bgpip-000004xf"),
//				Thresholds: dayu.CcPolicyV2ThresholdArray{
//					&dayu.CcPolicyV2ThresholdArgs{
//						Domain:    pulumi.String("12.com"),
//						Threshold: pulumi.Int(0),
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
type CcPolicyV2 struct {
	pulumi.CustomResourceState

	// Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared packet; net indicates anti-anti-ip pro version.
	Business pulumi.StringOutput `pulumi:"business"`
	// Blacklist and whitelist.
	CcBlackWhiteIps CcPolicyV2CcBlackWhiteIpArrayOutput `pulumi:"ccBlackWhiteIps"`
	// Details of the CC region blocking policy list.
	CcGeoIpPolicys CcPolicyV2CcGeoIpPolicyArrayOutput `pulumi:"ccGeoIpPolicys"`
	// CC Precision Protection List.
	CcPrecisionPolicys CcPolicyV2CcPrecisionPolicyArrayOutput `pulumi:"ccPrecisionPolicys"`
	// CC frequency throttling policy.
	CcPrecisionReqLimits CcPolicyV2CcPrecisionReqLimitArrayOutput `pulumi:"ccPrecisionReqLimits"`
	// The ID of the resource instance.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// List of protection threshold configurations.
	Thresholds CcPolicyV2ThresholdArrayOutput `pulumi:"thresholds"`
}

// NewCcPolicyV2 registers a new resource with the given unique name, arguments, and options.
func NewCcPolicyV2(ctx *pulumi.Context,
	name string, args *CcPolicyV2Args, opts ...pulumi.ResourceOption) (*CcPolicyV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Business == nil {
		return nil, errors.New("invalid value for required argument 'Business'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CcPolicyV2
	err := ctx.RegisterResource("tencentcloud:Dayu/ccPolicyV2:CcPolicyV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCcPolicyV2 gets an existing CcPolicyV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCcPolicyV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CcPolicyV2State, opts ...pulumi.ResourceOption) (*CcPolicyV2, error) {
	var resource CcPolicyV2
	err := ctx.ReadResource("tencentcloud:Dayu/ccPolicyV2:CcPolicyV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CcPolicyV2 resources.
type ccPolicyV2State struct {
	// Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared packet; net indicates anti-anti-ip pro version.
	Business *string `pulumi:"business"`
	// Blacklist and whitelist.
	CcBlackWhiteIps []CcPolicyV2CcBlackWhiteIp `pulumi:"ccBlackWhiteIps"`
	// Details of the CC region blocking policy list.
	CcGeoIpPolicys []CcPolicyV2CcGeoIpPolicy `pulumi:"ccGeoIpPolicys"`
	// CC Precision Protection List.
	CcPrecisionPolicys []CcPolicyV2CcPrecisionPolicy `pulumi:"ccPrecisionPolicys"`
	// CC frequency throttling policy.
	CcPrecisionReqLimits []CcPolicyV2CcPrecisionReqLimit `pulumi:"ccPrecisionReqLimits"`
	// The ID of the resource instance.
	ResourceId *string `pulumi:"resourceId"`
	// List of protection threshold configurations.
	Thresholds []CcPolicyV2Threshold `pulumi:"thresholds"`
}

type CcPolicyV2State struct {
	// Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared packet; net indicates anti-anti-ip pro version.
	Business pulumi.StringPtrInput
	// Blacklist and whitelist.
	CcBlackWhiteIps CcPolicyV2CcBlackWhiteIpArrayInput
	// Details of the CC region blocking policy list.
	CcGeoIpPolicys CcPolicyV2CcGeoIpPolicyArrayInput
	// CC Precision Protection List.
	CcPrecisionPolicys CcPolicyV2CcPrecisionPolicyArrayInput
	// CC frequency throttling policy.
	CcPrecisionReqLimits CcPolicyV2CcPrecisionReqLimitArrayInput
	// The ID of the resource instance.
	ResourceId pulumi.StringPtrInput
	// List of protection threshold configurations.
	Thresholds CcPolicyV2ThresholdArrayInput
}

func (CcPolicyV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*ccPolicyV2State)(nil)).Elem()
}

type ccPolicyV2Args struct {
	// Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared packet; net indicates anti-anti-ip pro version.
	Business string `pulumi:"business"`
	// Blacklist and whitelist.
	CcBlackWhiteIps []CcPolicyV2CcBlackWhiteIp `pulumi:"ccBlackWhiteIps"`
	// Details of the CC region blocking policy list.
	CcGeoIpPolicys []CcPolicyV2CcGeoIpPolicy `pulumi:"ccGeoIpPolicys"`
	// CC Precision Protection List.
	CcPrecisionPolicys []CcPolicyV2CcPrecisionPolicy `pulumi:"ccPrecisionPolicys"`
	// CC frequency throttling policy.
	CcPrecisionReqLimits []CcPolicyV2CcPrecisionReqLimit `pulumi:"ccPrecisionReqLimits"`
	// The ID of the resource instance.
	ResourceId string `pulumi:"resourceId"`
	// List of protection threshold configurations.
	Thresholds []CcPolicyV2Threshold `pulumi:"thresholds"`
}

// The set of arguments for constructing a CcPolicyV2 resource.
type CcPolicyV2Args struct {
	// Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared packet; net indicates anti-anti-ip pro version.
	Business pulumi.StringInput
	// Blacklist and whitelist.
	CcBlackWhiteIps CcPolicyV2CcBlackWhiteIpArrayInput
	// Details of the CC region blocking policy list.
	CcGeoIpPolicys CcPolicyV2CcGeoIpPolicyArrayInput
	// CC Precision Protection List.
	CcPrecisionPolicys CcPolicyV2CcPrecisionPolicyArrayInput
	// CC frequency throttling policy.
	CcPrecisionReqLimits CcPolicyV2CcPrecisionReqLimitArrayInput
	// The ID of the resource instance.
	ResourceId pulumi.StringInput
	// List of protection threshold configurations.
	Thresholds CcPolicyV2ThresholdArrayInput
}

func (CcPolicyV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*ccPolicyV2Args)(nil)).Elem()
}

type CcPolicyV2Input interface {
	pulumi.Input

	ToCcPolicyV2Output() CcPolicyV2Output
	ToCcPolicyV2OutputWithContext(ctx context.Context) CcPolicyV2Output
}

func (*CcPolicyV2) ElementType() reflect.Type {
	return reflect.TypeOf((**CcPolicyV2)(nil)).Elem()
}

func (i *CcPolicyV2) ToCcPolicyV2Output() CcPolicyV2Output {
	return i.ToCcPolicyV2OutputWithContext(context.Background())
}

func (i *CcPolicyV2) ToCcPolicyV2OutputWithContext(ctx context.Context) CcPolicyV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(CcPolicyV2Output)
}

// CcPolicyV2ArrayInput is an input type that accepts CcPolicyV2Array and CcPolicyV2ArrayOutput values.
// You can construct a concrete instance of `CcPolicyV2ArrayInput` via:
//
//	CcPolicyV2Array{ CcPolicyV2Args{...} }
type CcPolicyV2ArrayInput interface {
	pulumi.Input

	ToCcPolicyV2ArrayOutput() CcPolicyV2ArrayOutput
	ToCcPolicyV2ArrayOutputWithContext(context.Context) CcPolicyV2ArrayOutput
}

type CcPolicyV2Array []CcPolicyV2Input

func (CcPolicyV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CcPolicyV2)(nil)).Elem()
}

func (i CcPolicyV2Array) ToCcPolicyV2ArrayOutput() CcPolicyV2ArrayOutput {
	return i.ToCcPolicyV2ArrayOutputWithContext(context.Background())
}

func (i CcPolicyV2Array) ToCcPolicyV2ArrayOutputWithContext(ctx context.Context) CcPolicyV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CcPolicyV2ArrayOutput)
}

// CcPolicyV2MapInput is an input type that accepts CcPolicyV2Map and CcPolicyV2MapOutput values.
// You can construct a concrete instance of `CcPolicyV2MapInput` via:
//
//	CcPolicyV2Map{ "key": CcPolicyV2Args{...} }
type CcPolicyV2MapInput interface {
	pulumi.Input

	ToCcPolicyV2MapOutput() CcPolicyV2MapOutput
	ToCcPolicyV2MapOutputWithContext(context.Context) CcPolicyV2MapOutput
}

type CcPolicyV2Map map[string]CcPolicyV2Input

func (CcPolicyV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CcPolicyV2)(nil)).Elem()
}

func (i CcPolicyV2Map) ToCcPolicyV2MapOutput() CcPolicyV2MapOutput {
	return i.ToCcPolicyV2MapOutputWithContext(context.Background())
}

func (i CcPolicyV2Map) ToCcPolicyV2MapOutputWithContext(ctx context.Context) CcPolicyV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CcPolicyV2MapOutput)
}

type CcPolicyV2Output struct{ *pulumi.OutputState }

func (CcPolicyV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**CcPolicyV2)(nil)).Elem()
}

func (o CcPolicyV2Output) ToCcPolicyV2Output() CcPolicyV2Output {
	return o
}

func (o CcPolicyV2Output) ToCcPolicyV2OutputWithContext(ctx context.Context) CcPolicyV2Output {
	return o
}

// Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared packet; net indicates anti-anti-ip pro version.
func (o CcPolicyV2Output) Business() pulumi.StringOutput {
	return o.ApplyT(func(v *CcPolicyV2) pulumi.StringOutput { return v.Business }).(pulumi.StringOutput)
}

// Blacklist and whitelist.
func (o CcPolicyV2Output) CcBlackWhiteIps() CcPolicyV2CcBlackWhiteIpArrayOutput {
	return o.ApplyT(func(v *CcPolicyV2) CcPolicyV2CcBlackWhiteIpArrayOutput { return v.CcBlackWhiteIps }).(CcPolicyV2CcBlackWhiteIpArrayOutput)
}

// Details of the CC region blocking policy list.
func (o CcPolicyV2Output) CcGeoIpPolicys() CcPolicyV2CcGeoIpPolicyArrayOutput {
	return o.ApplyT(func(v *CcPolicyV2) CcPolicyV2CcGeoIpPolicyArrayOutput { return v.CcGeoIpPolicys }).(CcPolicyV2CcGeoIpPolicyArrayOutput)
}

// CC Precision Protection List.
func (o CcPolicyV2Output) CcPrecisionPolicys() CcPolicyV2CcPrecisionPolicyArrayOutput {
	return o.ApplyT(func(v *CcPolicyV2) CcPolicyV2CcPrecisionPolicyArrayOutput { return v.CcPrecisionPolicys }).(CcPolicyV2CcPrecisionPolicyArrayOutput)
}

// CC frequency throttling policy.
func (o CcPolicyV2Output) CcPrecisionReqLimits() CcPolicyV2CcPrecisionReqLimitArrayOutput {
	return o.ApplyT(func(v *CcPolicyV2) CcPolicyV2CcPrecisionReqLimitArrayOutput { return v.CcPrecisionReqLimits }).(CcPolicyV2CcPrecisionReqLimitArrayOutput)
}

// The ID of the resource instance.
func (o CcPolicyV2Output) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CcPolicyV2) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// List of protection threshold configurations.
func (o CcPolicyV2Output) Thresholds() CcPolicyV2ThresholdArrayOutput {
	return o.ApplyT(func(v *CcPolicyV2) CcPolicyV2ThresholdArrayOutput { return v.Thresholds }).(CcPolicyV2ThresholdArrayOutput)
}

type CcPolicyV2ArrayOutput struct{ *pulumi.OutputState }

func (CcPolicyV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CcPolicyV2)(nil)).Elem()
}

func (o CcPolicyV2ArrayOutput) ToCcPolicyV2ArrayOutput() CcPolicyV2ArrayOutput {
	return o
}

func (o CcPolicyV2ArrayOutput) ToCcPolicyV2ArrayOutputWithContext(ctx context.Context) CcPolicyV2ArrayOutput {
	return o
}

func (o CcPolicyV2ArrayOutput) Index(i pulumi.IntInput) CcPolicyV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CcPolicyV2 {
		return vs[0].([]*CcPolicyV2)[vs[1].(int)]
	}).(CcPolicyV2Output)
}

type CcPolicyV2MapOutput struct{ *pulumi.OutputState }

func (CcPolicyV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CcPolicyV2)(nil)).Elem()
}

func (o CcPolicyV2MapOutput) ToCcPolicyV2MapOutput() CcPolicyV2MapOutput {
	return o
}

func (o CcPolicyV2MapOutput) ToCcPolicyV2MapOutputWithContext(ctx context.Context) CcPolicyV2MapOutput {
	return o
}

func (o CcPolicyV2MapOutput) MapIndex(k pulumi.StringInput) CcPolicyV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CcPolicyV2 {
		return vs[0].(map[string]*CcPolicyV2)[vs[1].(string)]
	}).(CcPolicyV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CcPolicyV2Input)(nil)).Elem(), &CcPolicyV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*CcPolicyV2ArrayInput)(nil)).Elem(), CcPolicyV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*CcPolicyV2MapInput)(nil)).Elem(), CcPolicyV2Map{})
	pulumi.RegisterOutputType(CcPolicyV2Output{})
	pulumi.RegisterOutputType(CcPolicyV2ArrayOutput{})
	pulumi.RegisterOutputType(CcPolicyV2MapOutput{})
}
