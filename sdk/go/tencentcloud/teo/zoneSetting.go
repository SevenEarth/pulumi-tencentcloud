// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package teo

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a teo zoneSetting
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Teo"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Teo"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Teo.NewZoneSetting(ctx, "zoneSetting", &Teo.ZoneSettingArgs{
//				Cache: &teo.ZoneSettingCacheArgs{
//					FollowOrigin: &teo.ZoneSettingCacheFollowOriginArgs{
//						Switch: pulumi.String("on"),
//					},
//					NoCache: &teo.ZoneSettingCacheNoCacheArgs{
//						Switch: pulumi.String("off"),
//					},
//				},
//				CacheKey: &teo.ZoneSettingCacheKeyArgs{
//					FullUrlCache: pulumi.String("on"),
//					IgnoreCase:   pulumi.String("off"),
//					QueryString: &teo.ZoneSettingCacheKeyQueryStringArgs{
//						Action: pulumi.String("includeCustom"),
//						Switch: pulumi.String("off"),
//						Value:  []interface{}{},
//					},
//				},
//				CachePrefresh: &teo.ZoneSettingCachePrefreshArgs{
//					Percent: pulumi.Int(90),
//					Switch:  pulumi.String("off"),
//				},
//				ClientIpHeader: &teo.ZoneSettingClientIpHeaderArgs{
//					Switch: pulumi.String("off"),
//				},
//				Compression: &teo.ZoneSettingCompressionArgs{
//					Algorithms: pulumi.StringArray{
//						pulumi.String("brotli"),
//						pulumi.String("gzip"),
//					},
//					Switch: pulumi.String("on"),
//				},
//				ForceRedirect: &teo.ZoneSettingForceRedirectArgs{
//					RedirectStatusCode: pulumi.Int(302),
//					Switch:             pulumi.String("off"),
//				},
//				Https: &teo.ZoneSettingHttpsArgs{
//					Hsts: &teo.ZoneSettingHttpsHstsArgs{
//						IncludeSubDomains: pulumi.String("off"),
//						MaxAge:            pulumi.Int(0),
//						Preload:           pulumi.String("off"),
//						Switch:            pulumi.String("off"),
//					},
//					Http2:        pulumi.String("on"),
//					OcspStapling: pulumi.String("off"),
//					TlsVersions: pulumi.StringArray{
//						pulumi.String("TLSv1"),
//						pulumi.String("TLSv1.1"),
//						pulumi.String("TLSv1.2"),
//						pulumi.String("TLSv1.3"),
//					},
//				},
//				Ipv6: &teo.ZoneSettingIpv6Args{
//					Switch: pulumi.String("off"),
//				},
//				MaxAge: &teo.ZoneSettingMaxAgeArgs{
//					FollowOrigin: pulumi.String("on"),
//					MaxAgeTime:   pulumi.Int(0),
//				},
//				OfflineCache: &teo.ZoneSettingOfflineCacheArgs{
//					Switch: pulumi.String("on"),
//				},
//				Origin: &teo.ZoneSettingOriginArgs{
//					BackupOrigins:      pulumi.StringArray{},
//					OriginPullProtocol: pulumi.String("follow"),
//					Origins:            pulumi.StringArray{},
//				},
//				PostMaxSize: &teo.ZoneSettingPostMaxSizeArgs{
//					MaxSize: pulumi.Int(524288000),
//					Switch:  pulumi.String("on"),
//				},
//				Quic: &teo.ZoneSettingQuicArgs{
//					Switch: pulumi.String("off"),
//				},
//				SmartRouting: &teo.ZoneSettingSmartRoutingArgs{
//					Switch: pulumi.String("off"),
//				},
//				UpstreamHttp2: &teo.ZoneSettingUpstreamHttp2Args{
//					Switch: pulumi.String("off"),
//				},
//				WebSocket: &teo.ZoneSettingWebSocketArgs{
//					Switch:  pulumi.String("off"),
//					Timeout: pulumi.Int(30),
//				},
//				ZoneId: pulumi.String("zone-297z8rf93cfw"),
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
// teo zone_setting can be imported using the zone_id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Teo/zoneSetting:ZoneSetting zone_setting zone-297z8rf93cfw#
//
// ```
type ZoneSetting struct {
	pulumi.CustomResourceState

	// Acceleration area of the zone. Valid values: `mainland`, `overseas`.
	Area pulumi.StringOutput `pulumi:"area"`
	// Cache expiration time configuration.
	Cache ZoneSettingCacheOutput `pulumi:"cache"`
	// Node cache key configuration.
	CacheKey ZoneSettingCacheKeyOutput `pulumi:"cacheKey"`
	// Cache pre-refresh configuration.
	CachePrefresh ZoneSettingCachePrefreshOutput `pulumi:"cachePrefresh"`
	// Origin-pull client IP header configuration.
	ClientIpHeader ZoneSettingClientIpHeaderOutput `pulumi:"clientIpHeader"`
	// Smart compression configuration.
	Compression ZoneSettingCompressionOutput `pulumi:"compression"`
	// Force HTTPS redirect configuration.
	ForceRedirect ZoneSettingForceRedirectOutput `pulumi:"forceRedirect"`
	// HTTPS acceleration configuration.
	Https ZoneSettingHttpsOutput `pulumi:"https"`
	// IPv6 access configuration.
	Ipv6 ZoneSettingIpv6Output `pulumi:"ipv6"`
	// Browser cache configuration.
	MaxAge ZoneSettingMaxAgeOutput `pulumi:"maxAge"`
	// Offline cache configuration.
	OfflineCache ZoneSettingOfflineCacheOutput `pulumi:"offlineCache"`
	// Origin server configuration.
	Origin ZoneSettingOriginOutput `pulumi:"origin"`
	// Maximum size of files transferred over POST request.
	PostMaxSize ZoneSettingPostMaxSizeOutput `pulumi:"postMaxSize"`
	// QUIC access configuration.
	Quic ZoneSettingQuicOutput `pulumi:"quic"`
	// Smart acceleration configuration.
	SmartRouting ZoneSettingSmartRoutingOutput `pulumi:"smartRouting"`
	// HTTP2 origin-pull configuration.
	UpstreamHttp2 ZoneSettingUpstreamHttp2Output `pulumi:"upstreamHttp2"`
	// WebSocket configuration.
	WebSocket ZoneSettingWebSocketOutput `pulumi:"webSocket"`
	// Site ID.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewZoneSetting registers a new resource with the given unique name, arguments, and options.
func NewZoneSetting(ctx *pulumi.Context,
	name string, args *ZoneSettingArgs, opts ...pulumi.ResourceOption) (*ZoneSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ZoneSetting
	err := ctx.RegisterResource("tencentcloud:Teo/zoneSetting:ZoneSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZoneSetting gets an existing ZoneSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZoneSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneSettingState, opts ...pulumi.ResourceOption) (*ZoneSetting, error) {
	var resource ZoneSetting
	err := ctx.ReadResource("tencentcloud:Teo/zoneSetting:ZoneSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZoneSetting resources.
type zoneSettingState struct {
	// Acceleration area of the zone. Valid values: `mainland`, `overseas`.
	Area *string `pulumi:"area"`
	// Cache expiration time configuration.
	Cache *ZoneSettingCache `pulumi:"cache"`
	// Node cache key configuration.
	CacheKey *ZoneSettingCacheKey `pulumi:"cacheKey"`
	// Cache pre-refresh configuration.
	CachePrefresh *ZoneSettingCachePrefresh `pulumi:"cachePrefresh"`
	// Origin-pull client IP header configuration.
	ClientIpHeader *ZoneSettingClientIpHeader `pulumi:"clientIpHeader"`
	// Smart compression configuration.
	Compression *ZoneSettingCompression `pulumi:"compression"`
	// Force HTTPS redirect configuration.
	ForceRedirect *ZoneSettingForceRedirect `pulumi:"forceRedirect"`
	// HTTPS acceleration configuration.
	Https *ZoneSettingHttps `pulumi:"https"`
	// IPv6 access configuration.
	Ipv6 *ZoneSettingIpv6 `pulumi:"ipv6"`
	// Browser cache configuration.
	MaxAge *ZoneSettingMaxAge `pulumi:"maxAge"`
	// Offline cache configuration.
	OfflineCache *ZoneSettingOfflineCache `pulumi:"offlineCache"`
	// Origin server configuration.
	Origin *ZoneSettingOrigin `pulumi:"origin"`
	// Maximum size of files transferred over POST request.
	PostMaxSize *ZoneSettingPostMaxSize `pulumi:"postMaxSize"`
	// QUIC access configuration.
	Quic *ZoneSettingQuic `pulumi:"quic"`
	// Smart acceleration configuration.
	SmartRouting *ZoneSettingSmartRouting `pulumi:"smartRouting"`
	// HTTP2 origin-pull configuration.
	UpstreamHttp2 *ZoneSettingUpstreamHttp2 `pulumi:"upstreamHttp2"`
	// WebSocket configuration.
	WebSocket *ZoneSettingWebSocket `pulumi:"webSocket"`
	// Site ID.
	ZoneId *string `pulumi:"zoneId"`
}

type ZoneSettingState struct {
	// Acceleration area of the zone. Valid values: `mainland`, `overseas`.
	Area pulumi.StringPtrInput
	// Cache expiration time configuration.
	Cache ZoneSettingCachePtrInput
	// Node cache key configuration.
	CacheKey ZoneSettingCacheKeyPtrInput
	// Cache pre-refresh configuration.
	CachePrefresh ZoneSettingCachePrefreshPtrInput
	// Origin-pull client IP header configuration.
	ClientIpHeader ZoneSettingClientIpHeaderPtrInput
	// Smart compression configuration.
	Compression ZoneSettingCompressionPtrInput
	// Force HTTPS redirect configuration.
	ForceRedirect ZoneSettingForceRedirectPtrInput
	// HTTPS acceleration configuration.
	Https ZoneSettingHttpsPtrInput
	// IPv6 access configuration.
	Ipv6 ZoneSettingIpv6PtrInput
	// Browser cache configuration.
	MaxAge ZoneSettingMaxAgePtrInput
	// Offline cache configuration.
	OfflineCache ZoneSettingOfflineCachePtrInput
	// Origin server configuration.
	Origin ZoneSettingOriginPtrInput
	// Maximum size of files transferred over POST request.
	PostMaxSize ZoneSettingPostMaxSizePtrInput
	// QUIC access configuration.
	Quic ZoneSettingQuicPtrInput
	// Smart acceleration configuration.
	SmartRouting ZoneSettingSmartRoutingPtrInput
	// HTTP2 origin-pull configuration.
	UpstreamHttp2 ZoneSettingUpstreamHttp2PtrInput
	// WebSocket configuration.
	WebSocket ZoneSettingWebSocketPtrInput
	// Site ID.
	ZoneId pulumi.StringPtrInput
}

func (ZoneSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneSettingState)(nil)).Elem()
}

type zoneSettingArgs struct {
	// Cache expiration time configuration.
	Cache *ZoneSettingCache `pulumi:"cache"`
	// Node cache key configuration.
	CacheKey *ZoneSettingCacheKey `pulumi:"cacheKey"`
	// Cache pre-refresh configuration.
	CachePrefresh *ZoneSettingCachePrefresh `pulumi:"cachePrefresh"`
	// Origin-pull client IP header configuration.
	ClientIpHeader *ZoneSettingClientIpHeader `pulumi:"clientIpHeader"`
	// Smart compression configuration.
	Compression *ZoneSettingCompression `pulumi:"compression"`
	// Force HTTPS redirect configuration.
	ForceRedirect *ZoneSettingForceRedirect `pulumi:"forceRedirect"`
	// HTTPS acceleration configuration.
	Https *ZoneSettingHttps `pulumi:"https"`
	// IPv6 access configuration.
	Ipv6 *ZoneSettingIpv6 `pulumi:"ipv6"`
	// Browser cache configuration.
	MaxAge *ZoneSettingMaxAge `pulumi:"maxAge"`
	// Offline cache configuration.
	OfflineCache *ZoneSettingOfflineCache `pulumi:"offlineCache"`
	// Origin server configuration.
	Origin *ZoneSettingOrigin `pulumi:"origin"`
	// Maximum size of files transferred over POST request.
	PostMaxSize *ZoneSettingPostMaxSize `pulumi:"postMaxSize"`
	// QUIC access configuration.
	Quic *ZoneSettingQuic `pulumi:"quic"`
	// Smart acceleration configuration.
	SmartRouting *ZoneSettingSmartRouting `pulumi:"smartRouting"`
	// HTTP2 origin-pull configuration.
	UpstreamHttp2 *ZoneSettingUpstreamHttp2 `pulumi:"upstreamHttp2"`
	// WebSocket configuration.
	WebSocket *ZoneSettingWebSocket `pulumi:"webSocket"`
	// Site ID.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a ZoneSetting resource.
type ZoneSettingArgs struct {
	// Cache expiration time configuration.
	Cache ZoneSettingCachePtrInput
	// Node cache key configuration.
	CacheKey ZoneSettingCacheKeyPtrInput
	// Cache pre-refresh configuration.
	CachePrefresh ZoneSettingCachePrefreshPtrInput
	// Origin-pull client IP header configuration.
	ClientIpHeader ZoneSettingClientIpHeaderPtrInput
	// Smart compression configuration.
	Compression ZoneSettingCompressionPtrInput
	// Force HTTPS redirect configuration.
	ForceRedirect ZoneSettingForceRedirectPtrInput
	// HTTPS acceleration configuration.
	Https ZoneSettingHttpsPtrInput
	// IPv6 access configuration.
	Ipv6 ZoneSettingIpv6PtrInput
	// Browser cache configuration.
	MaxAge ZoneSettingMaxAgePtrInput
	// Offline cache configuration.
	OfflineCache ZoneSettingOfflineCachePtrInput
	// Origin server configuration.
	Origin ZoneSettingOriginPtrInput
	// Maximum size of files transferred over POST request.
	PostMaxSize ZoneSettingPostMaxSizePtrInput
	// QUIC access configuration.
	Quic ZoneSettingQuicPtrInput
	// Smart acceleration configuration.
	SmartRouting ZoneSettingSmartRoutingPtrInput
	// HTTP2 origin-pull configuration.
	UpstreamHttp2 ZoneSettingUpstreamHttp2PtrInput
	// WebSocket configuration.
	WebSocket ZoneSettingWebSocketPtrInput
	// Site ID.
	ZoneId pulumi.StringInput
}

func (ZoneSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneSettingArgs)(nil)).Elem()
}

type ZoneSettingInput interface {
	pulumi.Input

	ToZoneSettingOutput() ZoneSettingOutput
	ToZoneSettingOutputWithContext(ctx context.Context) ZoneSettingOutput
}

func (*ZoneSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneSetting)(nil)).Elem()
}

func (i *ZoneSetting) ToZoneSettingOutput() ZoneSettingOutput {
	return i.ToZoneSettingOutputWithContext(context.Background())
}

func (i *ZoneSetting) ToZoneSettingOutputWithContext(ctx context.Context) ZoneSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneSettingOutput)
}

// ZoneSettingArrayInput is an input type that accepts ZoneSettingArray and ZoneSettingArrayOutput values.
// You can construct a concrete instance of `ZoneSettingArrayInput` via:
//
//	ZoneSettingArray{ ZoneSettingArgs{...} }
type ZoneSettingArrayInput interface {
	pulumi.Input

	ToZoneSettingArrayOutput() ZoneSettingArrayOutput
	ToZoneSettingArrayOutputWithContext(context.Context) ZoneSettingArrayOutput
}

type ZoneSettingArray []ZoneSettingInput

func (ZoneSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZoneSetting)(nil)).Elem()
}

func (i ZoneSettingArray) ToZoneSettingArrayOutput() ZoneSettingArrayOutput {
	return i.ToZoneSettingArrayOutputWithContext(context.Background())
}

func (i ZoneSettingArray) ToZoneSettingArrayOutputWithContext(ctx context.Context) ZoneSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneSettingArrayOutput)
}

// ZoneSettingMapInput is an input type that accepts ZoneSettingMap and ZoneSettingMapOutput values.
// You can construct a concrete instance of `ZoneSettingMapInput` via:
//
//	ZoneSettingMap{ "key": ZoneSettingArgs{...} }
type ZoneSettingMapInput interface {
	pulumi.Input

	ToZoneSettingMapOutput() ZoneSettingMapOutput
	ToZoneSettingMapOutputWithContext(context.Context) ZoneSettingMapOutput
}

type ZoneSettingMap map[string]ZoneSettingInput

func (ZoneSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZoneSetting)(nil)).Elem()
}

func (i ZoneSettingMap) ToZoneSettingMapOutput() ZoneSettingMapOutput {
	return i.ToZoneSettingMapOutputWithContext(context.Background())
}

func (i ZoneSettingMap) ToZoneSettingMapOutputWithContext(ctx context.Context) ZoneSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneSettingMapOutput)
}

type ZoneSettingOutput struct{ *pulumi.OutputState }

func (ZoneSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneSetting)(nil)).Elem()
}

func (o ZoneSettingOutput) ToZoneSettingOutput() ZoneSettingOutput {
	return o
}

func (o ZoneSettingOutput) ToZoneSettingOutputWithContext(ctx context.Context) ZoneSettingOutput {
	return o
}

// Acceleration area of the zone. Valid values: `mainland`, `overseas`.
func (o ZoneSettingOutput) Area() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneSetting) pulumi.StringOutput { return v.Area }).(pulumi.StringOutput)
}

// Cache expiration time configuration.
func (o ZoneSettingOutput) Cache() ZoneSettingCacheOutput {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingCacheOutput { return v.Cache }).(ZoneSettingCacheOutput)
}

// Node cache key configuration.
func (o ZoneSettingOutput) CacheKey() ZoneSettingCacheKeyOutput {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingCacheKeyOutput { return v.CacheKey }).(ZoneSettingCacheKeyOutput)
}

// Cache pre-refresh configuration.
func (o ZoneSettingOutput) CachePrefresh() ZoneSettingCachePrefreshOutput {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingCachePrefreshOutput { return v.CachePrefresh }).(ZoneSettingCachePrefreshOutput)
}

// Origin-pull client IP header configuration.
func (o ZoneSettingOutput) ClientIpHeader() ZoneSettingClientIpHeaderOutput {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingClientIpHeaderOutput { return v.ClientIpHeader }).(ZoneSettingClientIpHeaderOutput)
}

// Smart compression configuration.
func (o ZoneSettingOutput) Compression() ZoneSettingCompressionOutput {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingCompressionOutput { return v.Compression }).(ZoneSettingCompressionOutput)
}

// Force HTTPS redirect configuration.
func (o ZoneSettingOutput) ForceRedirect() ZoneSettingForceRedirectOutput {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingForceRedirectOutput { return v.ForceRedirect }).(ZoneSettingForceRedirectOutput)
}

// HTTPS acceleration configuration.
func (o ZoneSettingOutput) Https() ZoneSettingHttpsOutput {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingHttpsOutput { return v.Https }).(ZoneSettingHttpsOutput)
}

// IPv6 access configuration.
func (o ZoneSettingOutput) Ipv6() ZoneSettingIpv6Output {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingIpv6Output { return v.Ipv6 }).(ZoneSettingIpv6Output)
}

// Browser cache configuration.
func (o ZoneSettingOutput) MaxAge() ZoneSettingMaxAgeOutput {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingMaxAgeOutput { return v.MaxAge }).(ZoneSettingMaxAgeOutput)
}

// Offline cache configuration.
func (o ZoneSettingOutput) OfflineCache() ZoneSettingOfflineCacheOutput {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingOfflineCacheOutput { return v.OfflineCache }).(ZoneSettingOfflineCacheOutput)
}

// Origin server configuration.
func (o ZoneSettingOutput) Origin() ZoneSettingOriginOutput {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingOriginOutput { return v.Origin }).(ZoneSettingOriginOutput)
}

// Maximum size of files transferred over POST request.
func (o ZoneSettingOutput) PostMaxSize() ZoneSettingPostMaxSizeOutput {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingPostMaxSizeOutput { return v.PostMaxSize }).(ZoneSettingPostMaxSizeOutput)
}

// QUIC access configuration.
func (o ZoneSettingOutput) Quic() ZoneSettingQuicOutput {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingQuicOutput { return v.Quic }).(ZoneSettingQuicOutput)
}

// Smart acceleration configuration.
func (o ZoneSettingOutput) SmartRouting() ZoneSettingSmartRoutingOutput {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingSmartRoutingOutput { return v.SmartRouting }).(ZoneSettingSmartRoutingOutput)
}

// HTTP2 origin-pull configuration.
func (o ZoneSettingOutput) UpstreamHttp2() ZoneSettingUpstreamHttp2Output {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingUpstreamHttp2Output { return v.UpstreamHttp2 }).(ZoneSettingUpstreamHttp2Output)
}

// WebSocket configuration.
func (o ZoneSettingOutput) WebSocket() ZoneSettingWebSocketOutput {
	return o.ApplyT(func(v *ZoneSetting) ZoneSettingWebSocketOutput { return v.WebSocket }).(ZoneSettingWebSocketOutput)
}

// Site ID.
func (o ZoneSettingOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneSetting) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type ZoneSettingArrayOutput struct{ *pulumi.OutputState }

func (ZoneSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZoneSetting)(nil)).Elem()
}

func (o ZoneSettingArrayOutput) ToZoneSettingArrayOutput() ZoneSettingArrayOutput {
	return o
}

func (o ZoneSettingArrayOutput) ToZoneSettingArrayOutputWithContext(ctx context.Context) ZoneSettingArrayOutput {
	return o
}

func (o ZoneSettingArrayOutput) Index(i pulumi.IntInput) ZoneSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ZoneSetting {
		return vs[0].([]*ZoneSetting)[vs[1].(int)]
	}).(ZoneSettingOutput)
}

type ZoneSettingMapOutput struct{ *pulumi.OutputState }

func (ZoneSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZoneSetting)(nil)).Elem()
}

func (o ZoneSettingMapOutput) ToZoneSettingMapOutput() ZoneSettingMapOutput {
	return o
}

func (o ZoneSettingMapOutput) ToZoneSettingMapOutputWithContext(ctx context.Context) ZoneSettingMapOutput {
	return o
}

func (o ZoneSettingMapOutput) MapIndex(k pulumi.StringInput) ZoneSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ZoneSetting {
		return vs[0].(map[string]*ZoneSetting)[vs[1].(string)]
	}).(ZoneSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneSettingInput)(nil)).Elem(), &ZoneSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneSettingArrayInput)(nil)).Elem(), ZoneSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneSettingMapInput)(nil)).Elem(), ZoneSettingMap{})
	pulumi.RegisterOutputType(ZoneSettingOutput{})
	pulumi.RegisterOutputType(ZoneSettingArrayOutput{})
	pulumi.RegisterOutputType(ZoneSettingMapOutput{})
}
