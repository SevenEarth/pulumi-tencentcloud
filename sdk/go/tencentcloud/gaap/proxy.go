// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a GAAP proxy.
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
//			_, err := Gaap.NewProxy(ctx, "foo", &Gaap.ProxyArgs{
//				AccessRegion:     pulumi.String("SouthChina"),
//				Bandwidth:        pulumi.Int(10),
//				Concurrent:       pulumi.Int(2),
//				RealserverRegion: pulumi.String("NorthChina"),
//				Tags: pulumi.AnyMap{
//					"test": pulumi.Any("test"),
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
// GAAP proxy can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Gaap/proxy:Proxy tencentcloud_gaap_proxy.foo link-11112222
//
// ```
type Proxy struct {
	pulumi.CustomResourceState

	// Access region of the GAAP proxy. Valid value: `Hongkong`, `SoutheastAsia`, `Korea`, `Europe`, `NorthAmerica`, `Canada`, `WestIndia`, `Thailand`, `Virginia`, `Japan`, `Taipei`, `SL_AZURE_NorthUAE`, `SL_AZURE_EastAUS`, `SL_AZURE_NorthCentralUSA`, `SL_AZURE_SouthIndia`, `SL_AZURE_SouthBrazil`, `SL_AZURE_NorthZAF`, `SL_AZURE_SoutheastAsia`, `SL_AZURE_CentralFrance`, `SL_AZURE_SouthEngland`, `SL_AZURE_EastUS`, `SL_AZURE_WestUS`, `SL_AZURE_SouthCentralUSA`, `Jakarta`, `Beijing`, `Shanghai`, `Guangzhou`, `Chengdu`, `SL_AZURE_NorwayEast`, `Chongqing`, `Nanjing`, `SaoPaulo`, `SL_AZURE_JapanEast`, `Changsha`, `Xian`, `Wuhan`, `Fuzhou`, `Shenyang`, `Zhengzhou`, `Jinan`, `Hangzhou`, `Shijiazhuang`, `Hefei`.
	AccessRegion pulumi.StringOutput `pulumi:"accessRegion"`
	// Maximum bandwidth of the GAAP proxy, unit is Mbps. Valid value: `10`, `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000` and `10000`. To set `2000`, `5000` or `10000`, you need to apply for a whitelist from Tencent Cloud.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// Maximum concurrency of the GAAP proxy, unit is 10k. Valid value: `2`, `5`, `10`, `20`, `30`, `40`, `50`, `60`, `70`, `80`, `90`, `100`, `150`, `200`, `250` and `300`. To set `150`, `200`, `250` or `300`, you need to apply for a whitelist from Tencent Cloud.
	Concurrent pulumi.IntOutput `pulumi:"concurrent"`
	// Creation time of the GAAP proxy.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Access domain of the GAAP proxy.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Indicates whether GAAP proxy is enabled, default value is `true`.
	Enable pulumi.BoolPtrOutput `pulumi:"enable"`
	// Forwarding IP of the GAAP proxy.
	ForwardIp pulumi.StringOutput `pulumi:"forwardIp"`
	// Access IP of the GAAP proxy.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Name of the GAAP proxy, the maximum length is 30.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network type. `normal`: regular BGP, `cn2`: boutique BGP, `triple`: triple play.
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
	// ID of the project within the GAAP proxy, `0` means is default project.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// Region of the GAAP realserver. Valid value: `Hongkong`, `SoutheastAsia`, `Korea`, `Europe`, `NorthAmerica`, `Canada`, `WestIndia`, `Thailand`, `Virginia`, `Japan`, `Taipei`, `SL_AZURE_NorthUAE`, `SL_AZURE_EastAUS`, `SL_AZURE_NorthCentralUSA`, `SL_AZURE_SouthIndia`, `SL_AZURE_SouthBrazil`, `SL_AZURE_NorthZAF`, `SL_AZURE_SoutheastAsia`, `SL_AZURE_CentralFrance`, `SL_AZURE_SouthEngland`, `SL_AZURE_EastUS`, `SL_AZURE_WestUS`, `SL_AZURE_SouthCentralUSA`, `Jakarta`, `Beijing`, `Shanghai`, `Guangzhou`, `Chengdu`, `SL_AZURE_NorwayEast`, `Chongqing`, `Nanjing`, `SaoPaulo`, `SL_AZURE_JapanEast`.
	RealserverRegion pulumi.StringOutput `pulumi:"realserverRegion"`
	// Indicates whether GAAP proxy can scalable.
	Scalable pulumi.BoolOutput `pulumi:"scalable"`
	// Status of the GAAP proxy.
	Status pulumi.StringOutput `pulumi:"status"`
	// Supported protocols of the GAAP proxy.
	SupportProtocols pulumi.StringArrayOutput `pulumi:"supportProtocols"`
	// Tags of the GAAP proxy. Tags that do not exist are not created automatically.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewProxy registers a new resource with the given unique name, arguments, and options.
func NewProxy(ctx *pulumi.Context,
	name string, args *ProxyArgs, opts ...pulumi.ResourceOption) (*Proxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessRegion == nil {
		return nil, errors.New("invalid value for required argument 'AccessRegion'")
	}
	if args.Bandwidth == nil {
		return nil, errors.New("invalid value for required argument 'Bandwidth'")
	}
	if args.Concurrent == nil {
		return nil, errors.New("invalid value for required argument 'Concurrent'")
	}
	if args.RealserverRegion == nil {
		return nil, errors.New("invalid value for required argument 'RealserverRegion'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Proxy
	err := ctx.RegisterResource("tencentcloud:Gaap/proxy:Proxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProxy gets an existing Proxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProxyState, opts ...pulumi.ResourceOption) (*Proxy, error) {
	var resource Proxy
	err := ctx.ReadResource("tencentcloud:Gaap/proxy:Proxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Proxy resources.
type proxyState struct {
	// Access region of the GAAP proxy. Valid value: `Hongkong`, `SoutheastAsia`, `Korea`, `Europe`, `NorthAmerica`, `Canada`, `WestIndia`, `Thailand`, `Virginia`, `Japan`, `Taipei`, `SL_AZURE_NorthUAE`, `SL_AZURE_EastAUS`, `SL_AZURE_NorthCentralUSA`, `SL_AZURE_SouthIndia`, `SL_AZURE_SouthBrazil`, `SL_AZURE_NorthZAF`, `SL_AZURE_SoutheastAsia`, `SL_AZURE_CentralFrance`, `SL_AZURE_SouthEngland`, `SL_AZURE_EastUS`, `SL_AZURE_WestUS`, `SL_AZURE_SouthCentralUSA`, `Jakarta`, `Beijing`, `Shanghai`, `Guangzhou`, `Chengdu`, `SL_AZURE_NorwayEast`, `Chongqing`, `Nanjing`, `SaoPaulo`, `SL_AZURE_JapanEast`, `Changsha`, `Xian`, `Wuhan`, `Fuzhou`, `Shenyang`, `Zhengzhou`, `Jinan`, `Hangzhou`, `Shijiazhuang`, `Hefei`.
	AccessRegion *string `pulumi:"accessRegion"`
	// Maximum bandwidth of the GAAP proxy, unit is Mbps. Valid value: `10`, `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000` and `10000`. To set `2000`, `5000` or `10000`, you need to apply for a whitelist from Tencent Cloud.
	Bandwidth *int `pulumi:"bandwidth"`
	// Maximum concurrency of the GAAP proxy, unit is 10k. Valid value: `2`, `5`, `10`, `20`, `30`, `40`, `50`, `60`, `70`, `80`, `90`, `100`, `150`, `200`, `250` and `300`. To set `150`, `200`, `250` or `300`, you need to apply for a whitelist from Tencent Cloud.
	Concurrent *int `pulumi:"concurrent"`
	// Creation time of the GAAP proxy.
	CreateTime *string `pulumi:"createTime"`
	// Access domain of the GAAP proxy.
	Domain *string `pulumi:"domain"`
	// Indicates whether GAAP proxy is enabled, default value is `true`.
	Enable *bool `pulumi:"enable"`
	// Forwarding IP of the GAAP proxy.
	ForwardIp *string `pulumi:"forwardIp"`
	// Access IP of the GAAP proxy.
	Ip *string `pulumi:"ip"`
	// Name of the GAAP proxy, the maximum length is 30.
	Name *string `pulumi:"name"`
	// Network type. `normal`: regular BGP, `cn2`: boutique BGP, `triple`: triple play.
	NetworkType *string `pulumi:"networkType"`
	// ID of the project within the GAAP proxy, `0` means is default project.
	ProjectId *int `pulumi:"projectId"`
	// Region of the GAAP realserver. Valid value: `Hongkong`, `SoutheastAsia`, `Korea`, `Europe`, `NorthAmerica`, `Canada`, `WestIndia`, `Thailand`, `Virginia`, `Japan`, `Taipei`, `SL_AZURE_NorthUAE`, `SL_AZURE_EastAUS`, `SL_AZURE_NorthCentralUSA`, `SL_AZURE_SouthIndia`, `SL_AZURE_SouthBrazil`, `SL_AZURE_NorthZAF`, `SL_AZURE_SoutheastAsia`, `SL_AZURE_CentralFrance`, `SL_AZURE_SouthEngland`, `SL_AZURE_EastUS`, `SL_AZURE_WestUS`, `SL_AZURE_SouthCentralUSA`, `Jakarta`, `Beijing`, `Shanghai`, `Guangzhou`, `Chengdu`, `SL_AZURE_NorwayEast`, `Chongqing`, `Nanjing`, `SaoPaulo`, `SL_AZURE_JapanEast`.
	RealserverRegion *string `pulumi:"realserverRegion"`
	// Indicates whether GAAP proxy can scalable.
	Scalable *bool `pulumi:"scalable"`
	// Status of the GAAP proxy.
	Status *string `pulumi:"status"`
	// Supported protocols of the GAAP proxy.
	SupportProtocols []string `pulumi:"supportProtocols"`
	// Tags of the GAAP proxy. Tags that do not exist are not created automatically.
	Tags map[string]interface{} `pulumi:"tags"`
}

type ProxyState struct {
	// Access region of the GAAP proxy. Valid value: `Hongkong`, `SoutheastAsia`, `Korea`, `Europe`, `NorthAmerica`, `Canada`, `WestIndia`, `Thailand`, `Virginia`, `Japan`, `Taipei`, `SL_AZURE_NorthUAE`, `SL_AZURE_EastAUS`, `SL_AZURE_NorthCentralUSA`, `SL_AZURE_SouthIndia`, `SL_AZURE_SouthBrazil`, `SL_AZURE_NorthZAF`, `SL_AZURE_SoutheastAsia`, `SL_AZURE_CentralFrance`, `SL_AZURE_SouthEngland`, `SL_AZURE_EastUS`, `SL_AZURE_WestUS`, `SL_AZURE_SouthCentralUSA`, `Jakarta`, `Beijing`, `Shanghai`, `Guangzhou`, `Chengdu`, `SL_AZURE_NorwayEast`, `Chongqing`, `Nanjing`, `SaoPaulo`, `SL_AZURE_JapanEast`, `Changsha`, `Xian`, `Wuhan`, `Fuzhou`, `Shenyang`, `Zhengzhou`, `Jinan`, `Hangzhou`, `Shijiazhuang`, `Hefei`.
	AccessRegion pulumi.StringPtrInput
	// Maximum bandwidth of the GAAP proxy, unit is Mbps. Valid value: `10`, `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000` and `10000`. To set `2000`, `5000` or `10000`, you need to apply for a whitelist from Tencent Cloud.
	Bandwidth pulumi.IntPtrInput
	// Maximum concurrency of the GAAP proxy, unit is 10k. Valid value: `2`, `5`, `10`, `20`, `30`, `40`, `50`, `60`, `70`, `80`, `90`, `100`, `150`, `200`, `250` and `300`. To set `150`, `200`, `250` or `300`, you need to apply for a whitelist from Tencent Cloud.
	Concurrent pulumi.IntPtrInput
	// Creation time of the GAAP proxy.
	CreateTime pulumi.StringPtrInput
	// Access domain of the GAAP proxy.
	Domain pulumi.StringPtrInput
	// Indicates whether GAAP proxy is enabled, default value is `true`.
	Enable pulumi.BoolPtrInput
	// Forwarding IP of the GAAP proxy.
	ForwardIp pulumi.StringPtrInput
	// Access IP of the GAAP proxy.
	Ip pulumi.StringPtrInput
	// Name of the GAAP proxy, the maximum length is 30.
	Name pulumi.StringPtrInput
	// Network type. `normal`: regular BGP, `cn2`: boutique BGP, `triple`: triple play.
	NetworkType pulumi.StringPtrInput
	// ID of the project within the GAAP proxy, `0` means is default project.
	ProjectId pulumi.IntPtrInput
	// Region of the GAAP realserver. Valid value: `Hongkong`, `SoutheastAsia`, `Korea`, `Europe`, `NorthAmerica`, `Canada`, `WestIndia`, `Thailand`, `Virginia`, `Japan`, `Taipei`, `SL_AZURE_NorthUAE`, `SL_AZURE_EastAUS`, `SL_AZURE_NorthCentralUSA`, `SL_AZURE_SouthIndia`, `SL_AZURE_SouthBrazil`, `SL_AZURE_NorthZAF`, `SL_AZURE_SoutheastAsia`, `SL_AZURE_CentralFrance`, `SL_AZURE_SouthEngland`, `SL_AZURE_EastUS`, `SL_AZURE_WestUS`, `SL_AZURE_SouthCentralUSA`, `Jakarta`, `Beijing`, `Shanghai`, `Guangzhou`, `Chengdu`, `SL_AZURE_NorwayEast`, `Chongqing`, `Nanjing`, `SaoPaulo`, `SL_AZURE_JapanEast`.
	RealserverRegion pulumi.StringPtrInput
	// Indicates whether GAAP proxy can scalable.
	Scalable pulumi.BoolPtrInput
	// Status of the GAAP proxy.
	Status pulumi.StringPtrInput
	// Supported protocols of the GAAP proxy.
	SupportProtocols pulumi.StringArrayInput
	// Tags of the GAAP proxy. Tags that do not exist are not created automatically.
	Tags pulumi.MapInput
}

func (ProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyState)(nil)).Elem()
}

type proxyArgs struct {
	// Access region of the GAAP proxy. Valid value: `Hongkong`, `SoutheastAsia`, `Korea`, `Europe`, `NorthAmerica`, `Canada`, `WestIndia`, `Thailand`, `Virginia`, `Japan`, `Taipei`, `SL_AZURE_NorthUAE`, `SL_AZURE_EastAUS`, `SL_AZURE_NorthCentralUSA`, `SL_AZURE_SouthIndia`, `SL_AZURE_SouthBrazil`, `SL_AZURE_NorthZAF`, `SL_AZURE_SoutheastAsia`, `SL_AZURE_CentralFrance`, `SL_AZURE_SouthEngland`, `SL_AZURE_EastUS`, `SL_AZURE_WestUS`, `SL_AZURE_SouthCentralUSA`, `Jakarta`, `Beijing`, `Shanghai`, `Guangzhou`, `Chengdu`, `SL_AZURE_NorwayEast`, `Chongqing`, `Nanjing`, `SaoPaulo`, `SL_AZURE_JapanEast`, `Changsha`, `Xian`, `Wuhan`, `Fuzhou`, `Shenyang`, `Zhengzhou`, `Jinan`, `Hangzhou`, `Shijiazhuang`, `Hefei`.
	AccessRegion string `pulumi:"accessRegion"`
	// Maximum bandwidth of the GAAP proxy, unit is Mbps. Valid value: `10`, `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000` and `10000`. To set `2000`, `5000` or `10000`, you need to apply for a whitelist from Tencent Cloud.
	Bandwidth int `pulumi:"bandwidth"`
	// Maximum concurrency of the GAAP proxy, unit is 10k. Valid value: `2`, `5`, `10`, `20`, `30`, `40`, `50`, `60`, `70`, `80`, `90`, `100`, `150`, `200`, `250` and `300`. To set `150`, `200`, `250` or `300`, you need to apply for a whitelist from Tencent Cloud.
	Concurrent int `pulumi:"concurrent"`
	// Indicates whether GAAP proxy is enabled, default value is `true`.
	Enable *bool `pulumi:"enable"`
	// Name of the GAAP proxy, the maximum length is 30.
	Name *string `pulumi:"name"`
	// Network type. `normal`: regular BGP, `cn2`: boutique BGP, `triple`: triple play.
	NetworkType *string `pulumi:"networkType"`
	// ID of the project within the GAAP proxy, `0` means is default project.
	ProjectId *int `pulumi:"projectId"`
	// Region of the GAAP realserver. Valid value: `Hongkong`, `SoutheastAsia`, `Korea`, `Europe`, `NorthAmerica`, `Canada`, `WestIndia`, `Thailand`, `Virginia`, `Japan`, `Taipei`, `SL_AZURE_NorthUAE`, `SL_AZURE_EastAUS`, `SL_AZURE_NorthCentralUSA`, `SL_AZURE_SouthIndia`, `SL_AZURE_SouthBrazil`, `SL_AZURE_NorthZAF`, `SL_AZURE_SoutheastAsia`, `SL_AZURE_CentralFrance`, `SL_AZURE_SouthEngland`, `SL_AZURE_EastUS`, `SL_AZURE_WestUS`, `SL_AZURE_SouthCentralUSA`, `Jakarta`, `Beijing`, `Shanghai`, `Guangzhou`, `Chengdu`, `SL_AZURE_NorwayEast`, `Chongqing`, `Nanjing`, `SaoPaulo`, `SL_AZURE_JapanEast`.
	RealserverRegion string `pulumi:"realserverRegion"`
	// Tags of the GAAP proxy. Tags that do not exist are not created automatically.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Proxy resource.
type ProxyArgs struct {
	// Access region of the GAAP proxy. Valid value: `Hongkong`, `SoutheastAsia`, `Korea`, `Europe`, `NorthAmerica`, `Canada`, `WestIndia`, `Thailand`, `Virginia`, `Japan`, `Taipei`, `SL_AZURE_NorthUAE`, `SL_AZURE_EastAUS`, `SL_AZURE_NorthCentralUSA`, `SL_AZURE_SouthIndia`, `SL_AZURE_SouthBrazil`, `SL_AZURE_NorthZAF`, `SL_AZURE_SoutheastAsia`, `SL_AZURE_CentralFrance`, `SL_AZURE_SouthEngland`, `SL_AZURE_EastUS`, `SL_AZURE_WestUS`, `SL_AZURE_SouthCentralUSA`, `Jakarta`, `Beijing`, `Shanghai`, `Guangzhou`, `Chengdu`, `SL_AZURE_NorwayEast`, `Chongqing`, `Nanjing`, `SaoPaulo`, `SL_AZURE_JapanEast`, `Changsha`, `Xian`, `Wuhan`, `Fuzhou`, `Shenyang`, `Zhengzhou`, `Jinan`, `Hangzhou`, `Shijiazhuang`, `Hefei`.
	AccessRegion pulumi.StringInput
	// Maximum bandwidth of the GAAP proxy, unit is Mbps. Valid value: `10`, `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000` and `10000`. To set `2000`, `5000` or `10000`, you need to apply for a whitelist from Tencent Cloud.
	Bandwidth pulumi.IntInput
	// Maximum concurrency of the GAAP proxy, unit is 10k. Valid value: `2`, `5`, `10`, `20`, `30`, `40`, `50`, `60`, `70`, `80`, `90`, `100`, `150`, `200`, `250` and `300`. To set `150`, `200`, `250` or `300`, you need to apply for a whitelist from Tencent Cloud.
	Concurrent pulumi.IntInput
	// Indicates whether GAAP proxy is enabled, default value is `true`.
	Enable pulumi.BoolPtrInput
	// Name of the GAAP proxy, the maximum length is 30.
	Name pulumi.StringPtrInput
	// Network type. `normal`: regular BGP, `cn2`: boutique BGP, `triple`: triple play.
	NetworkType pulumi.StringPtrInput
	// ID of the project within the GAAP proxy, `0` means is default project.
	ProjectId pulumi.IntPtrInput
	// Region of the GAAP realserver. Valid value: `Hongkong`, `SoutheastAsia`, `Korea`, `Europe`, `NorthAmerica`, `Canada`, `WestIndia`, `Thailand`, `Virginia`, `Japan`, `Taipei`, `SL_AZURE_NorthUAE`, `SL_AZURE_EastAUS`, `SL_AZURE_NorthCentralUSA`, `SL_AZURE_SouthIndia`, `SL_AZURE_SouthBrazil`, `SL_AZURE_NorthZAF`, `SL_AZURE_SoutheastAsia`, `SL_AZURE_CentralFrance`, `SL_AZURE_SouthEngland`, `SL_AZURE_EastUS`, `SL_AZURE_WestUS`, `SL_AZURE_SouthCentralUSA`, `Jakarta`, `Beijing`, `Shanghai`, `Guangzhou`, `Chengdu`, `SL_AZURE_NorwayEast`, `Chongqing`, `Nanjing`, `SaoPaulo`, `SL_AZURE_JapanEast`.
	RealserverRegion pulumi.StringInput
	// Tags of the GAAP proxy. Tags that do not exist are not created automatically.
	Tags pulumi.MapInput
}

func (ProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyArgs)(nil)).Elem()
}

type ProxyInput interface {
	pulumi.Input

	ToProxyOutput() ProxyOutput
	ToProxyOutputWithContext(ctx context.Context) ProxyOutput
}

func (*Proxy) ElementType() reflect.Type {
	return reflect.TypeOf((**Proxy)(nil)).Elem()
}

func (i *Proxy) ToProxyOutput() ProxyOutput {
	return i.ToProxyOutputWithContext(context.Background())
}

func (i *Proxy) ToProxyOutputWithContext(ctx context.Context) ProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyOutput)
}

// ProxyArrayInput is an input type that accepts ProxyArray and ProxyArrayOutput values.
// You can construct a concrete instance of `ProxyArrayInput` via:
//
//	ProxyArray{ ProxyArgs{...} }
type ProxyArrayInput interface {
	pulumi.Input

	ToProxyArrayOutput() ProxyArrayOutput
	ToProxyArrayOutputWithContext(context.Context) ProxyArrayOutput
}

type ProxyArray []ProxyInput

func (ProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Proxy)(nil)).Elem()
}

func (i ProxyArray) ToProxyArrayOutput() ProxyArrayOutput {
	return i.ToProxyArrayOutputWithContext(context.Background())
}

func (i ProxyArray) ToProxyArrayOutputWithContext(ctx context.Context) ProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyArrayOutput)
}

// ProxyMapInput is an input type that accepts ProxyMap and ProxyMapOutput values.
// You can construct a concrete instance of `ProxyMapInput` via:
//
//	ProxyMap{ "key": ProxyArgs{...} }
type ProxyMapInput interface {
	pulumi.Input

	ToProxyMapOutput() ProxyMapOutput
	ToProxyMapOutputWithContext(context.Context) ProxyMapOutput
}

type ProxyMap map[string]ProxyInput

func (ProxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Proxy)(nil)).Elem()
}

func (i ProxyMap) ToProxyMapOutput() ProxyMapOutput {
	return i.ToProxyMapOutputWithContext(context.Background())
}

func (i ProxyMap) ToProxyMapOutputWithContext(ctx context.Context) ProxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyMapOutput)
}

type ProxyOutput struct{ *pulumi.OutputState }

func (ProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Proxy)(nil)).Elem()
}

func (o ProxyOutput) ToProxyOutput() ProxyOutput {
	return o
}

func (o ProxyOutput) ToProxyOutputWithContext(ctx context.Context) ProxyOutput {
	return o
}

// Access region of the GAAP proxy. Valid value: `Hongkong`, `SoutheastAsia`, `Korea`, `Europe`, `NorthAmerica`, `Canada`, `WestIndia`, `Thailand`, `Virginia`, `Japan`, `Taipei`, `SL_AZURE_NorthUAE`, `SL_AZURE_EastAUS`, `SL_AZURE_NorthCentralUSA`, `SL_AZURE_SouthIndia`, `SL_AZURE_SouthBrazil`, `SL_AZURE_NorthZAF`, `SL_AZURE_SoutheastAsia`, `SL_AZURE_CentralFrance`, `SL_AZURE_SouthEngland`, `SL_AZURE_EastUS`, `SL_AZURE_WestUS`, `SL_AZURE_SouthCentralUSA`, `Jakarta`, `Beijing`, `Shanghai`, `Guangzhou`, `Chengdu`, `SL_AZURE_NorwayEast`, `Chongqing`, `Nanjing`, `SaoPaulo`, `SL_AZURE_JapanEast`, `Changsha`, `Xian`, `Wuhan`, `Fuzhou`, `Shenyang`, `Zhengzhou`, `Jinan`, `Hangzhou`, `Shijiazhuang`, `Hefei`.
func (o ProxyOutput) AccessRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringOutput { return v.AccessRegion }).(pulumi.StringOutput)
}

// Maximum bandwidth of the GAAP proxy, unit is Mbps. Valid value: `10`, `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000` and `10000`. To set `2000`, `5000` or `10000`, you need to apply for a whitelist from Tencent Cloud.
func (o ProxyOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *Proxy) pulumi.IntOutput { return v.Bandwidth }).(pulumi.IntOutput)
}

// Maximum concurrency of the GAAP proxy, unit is 10k. Valid value: `2`, `5`, `10`, `20`, `30`, `40`, `50`, `60`, `70`, `80`, `90`, `100`, `150`, `200`, `250` and `300`. To set `150`, `200`, `250` or `300`, you need to apply for a whitelist from Tencent Cloud.
func (o ProxyOutput) Concurrent() pulumi.IntOutput {
	return o.ApplyT(func(v *Proxy) pulumi.IntOutput { return v.Concurrent }).(pulumi.IntOutput)
}

// Creation time of the GAAP proxy.
func (o ProxyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Access domain of the GAAP proxy.
func (o ProxyOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Indicates whether GAAP proxy is enabled, default value is `true`.
func (o ProxyOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Proxy) pulumi.BoolPtrOutput { return v.Enable }).(pulumi.BoolPtrOutput)
}

// Forwarding IP of the GAAP proxy.
func (o ProxyOutput) ForwardIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringOutput { return v.ForwardIp }).(pulumi.StringOutput)
}

// Access IP of the GAAP proxy.
func (o ProxyOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// Name of the GAAP proxy, the maximum length is 30.
func (o ProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network type. `normal`: regular BGP, `cn2`: boutique BGP, `triple`: triple play.
func (o ProxyOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringOutput { return v.NetworkType }).(pulumi.StringOutput)
}

// ID of the project within the GAAP proxy, `0` means is default project.
func (o ProxyOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Proxy) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// Region of the GAAP realserver. Valid value: `Hongkong`, `SoutheastAsia`, `Korea`, `Europe`, `NorthAmerica`, `Canada`, `WestIndia`, `Thailand`, `Virginia`, `Japan`, `Taipei`, `SL_AZURE_NorthUAE`, `SL_AZURE_EastAUS`, `SL_AZURE_NorthCentralUSA`, `SL_AZURE_SouthIndia`, `SL_AZURE_SouthBrazil`, `SL_AZURE_NorthZAF`, `SL_AZURE_SoutheastAsia`, `SL_AZURE_CentralFrance`, `SL_AZURE_SouthEngland`, `SL_AZURE_EastUS`, `SL_AZURE_WestUS`, `SL_AZURE_SouthCentralUSA`, `Jakarta`, `Beijing`, `Shanghai`, `Guangzhou`, `Chengdu`, `SL_AZURE_NorwayEast`, `Chongqing`, `Nanjing`, `SaoPaulo`, `SL_AZURE_JapanEast`.
func (o ProxyOutput) RealserverRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringOutput { return v.RealserverRegion }).(pulumi.StringOutput)
}

// Indicates whether GAAP proxy can scalable.
func (o ProxyOutput) Scalable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Proxy) pulumi.BoolOutput { return v.Scalable }).(pulumi.BoolOutput)
}

// Status of the GAAP proxy.
func (o ProxyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Supported protocols of the GAAP proxy.
func (o ProxyOutput) SupportProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringArrayOutput { return v.SupportProtocols }).(pulumi.StringArrayOutput)
}

// Tags of the GAAP proxy. Tags that do not exist are not created automatically.
func (o ProxyOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Proxy) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type ProxyArrayOutput struct{ *pulumi.OutputState }

func (ProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Proxy)(nil)).Elem()
}

func (o ProxyArrayOutput) ToProxyArrayOutput() ProxyArrayOutput {
	return o
}

func (o ProxyArrayOutput) ToProxyArrayOutputWithContext(ctx context.Context) ProxyArrayOutput {
	return o
}

func (o ProxyArrayOutput) Index(i pulumi.IntInput) ProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Proxy {
		return vs[0].([]*Proxy)[vs[1].(int)]
	}).(ProxyOutput)
}

type ProxyMapOutput struct{ *pulumi.OutputState }

func (ProxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Proxy)(nil)).Elem()
}

func (o ProxyMapOutput) ToProxyMapOutput() ProxyMapOutput {
	return o
}

func (o ProxyMapOutput) ToProxyMapOutputWithContext(ctx context.Context) ProxyMapOutput {
	return o
}

func (o ProxyMapOutput) MapIndex(k pulumi.StringInput) ProxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Proxy {
		return vs[0].(map[string]*Proxy)[vs[1].(string)]
	}).(ProxyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyInput)(nil)).Elem(), &Proxy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyArrayInput)(nil)).Elem(), ProxyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyMapInput)(nil)).Elem(), ProxyMap{})
	pulumi.RegisterOutputType(ProxyOutput{})
	pulumi.RegisterOutputType(ProxyArrayOutput{})
	pulumi.RegisterOutputType(ProxyMapOutput{})
}
