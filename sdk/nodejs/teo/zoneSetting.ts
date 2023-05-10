// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a teo zoneSetting
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const zoneSetting = new tencentcloud.Teo.ZoneSetting("zone_setting", {
 *     cache: {
 *         followOrigin: {
 *             switch: "on",
 *         },
 *         noCache: {
 *             switch: "off",
 *         },
 *     },
 *     cacheKey: {
 *         fullUrlCache: "on",
 *         ignoreCase: "off",
 *         queryString: {
 *             action: "includeCustom",
 *             switch: "off",
 *             values: [],
 *         },
 *     },
 *     cachePrefresh: {
 *         percent: 90,
 *         switch: "off",
 *     },
 *     clientIpHeader: {
 *         switch: "off",
 *     },
 *     compression: {
 *         algorithms: [
 *             "brotli",
 *             "gzip",
 *         ],
 *         switch: "on",
 *     },
 *     forceRedirect: {
 *         redirectStatusCode: 302,
 *         switch: "off",
 *     },
 *     https: {
 *         hsts: {
 *             includeSubDomains: "off",
 *             maxAge: 0,
 *             preload: "off",
 *             switch: "off",
 *         },
 *         http2: "on",
 *         ocspStapling: "off",
 *         tlsVersions: [
 *             "TLSv1",
 *             "TLSv1.1",
 *             "TLSv1.2",
 *             "TLSv1.3",
 *         ],
 *     },
 *     ipv6: {
 *         switch: "off",
 *     },
 *     maxAge: {
 *         followOrigin: "on",
 *         maxAgeTime: 0,
 *     },
 *     offlineCache: {
 *         switch: "on",
 *     },
 *     origin: {
 *         backupOrigins: [],
 *         originPullProtocol: "follow",
 *         origins: [],
 *     },
 *     postMaxSize: {
 *         maxSize: 524288000,
 *         switch: "on",
 *     },
 *     quic: {
 *         switch: "off",
 *     },
 *     smartRouting: {
 *         switch: "off",
 *     },
 *     upstreamHttp2: {
 *         switch: "off",
 *     },
 *     webSocket: {
 *         switch: "off",
 *         timeout: 30,
 *     },
 *     zoneId: "zone-297z8rf93cfw",
 * });
 * ```
 *
 * ## Import
 *
 * teo zone_setting can be imported using the zone_id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Teo/zoneSetting:ZoneSetting zone_setting zone-297z8rf93cfw#
 * ```
 */
export class ZoneSetting extends pulumi.CustomResource {
    /**
     * Get an existing ZoneSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneSettingState, opts?: pulumi.CustomResourceOptions): ZoneSetting {
        return new ZoneSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Teo/zoneSetting:ZoneSetting';

    /**
     * Returns true if the given object is an instance of ZoneSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ZoneSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ZoneSetting.__pulumiType;
    }

    /**
     * Acceleration area of the zone. Valid values: `mainland`, `overseas`.
     */
    public /*out*/ readonly area!: pulumi.Output<string>;
    /**
     * Cache expiration time configuration.
     */
    public readonly cache!: pulumi.Output<outputs.Teo.ZoneSettingCache>;
    /**
     * Node cache key configuration.
     */
    public readonly cacheKey!: pulumi.Output<outputs.Teo.ZoneSettingCacheKey>;
    /**
     * Cache pre-refresh configuration.
     */
    public readonly cachePrefresh!: pulumi.Output<outputs.Teo.ZoneSettingCachePrefresh>;
    /**
     * Origin-pull client IP header configuration.
     */
    public readonly clientIpHeader!: pulumi.Output<outputs.Teo.ZoneSettingClientIpHeader>;
    /**
     * Smart compression configuration.
     */
    public readonly compression!: pulumi.Output<outputs.Teo.ZoneSettingCompression>;
    /**
     * Force HTTPS redirect configuration.
     */
    public readonly forceRedirect!: pulumi.Output<outputs.Teo.ZoneSettingForceRedirect>;
    /**
     * HTTPS acceleration configuration.
     */
    public readonly https!: pulumi.Output<outputs.Teo.ZoneSettingHttps>;
    /**
     * IPv6 access configuration.
     */
    public readonly ipv6!: pulumi.Output<outputs.Teo.ZoneSettingIpv6>;
    /**
     * Browser cache configuration.
     */
    public readonly maxAge!: pulumi.Output<outputs.Teo.ZoneSettingMaxAge>;
    /**
     * Offline cache configuration.
     */
    public readonly offlineCache!: pulumi.Output<outputs.Teo.ZoneSettingOfflineCache>;
    /**
     * Origin server configuration.
     */
    public readonly origin!: pulumi.Output<outputs.Teo.ZoneSettingOrigin>;
    /**
     * Maximum size of files transferred over POST request.
     */
    public readonly postMaxSize!: pulumi.Output<outputs.Teo.ZoneSettingPostMaxSize>;
    /**
     * QUIC access configuration.
     */
    public readonly quic!: pulumi.Output<outputs.Teo.ZoneSettingQuic>;
    /**
     * Smart acceleration configuration.
     */
    public readonly smartRouting!: pulumi.Output<outputs.Teo.ZoneSettingSmartRouting>;
    /**
     * HTTP2 origin-pull configuration.
     */
    public readonly upstreamHttp2!: pulumi.Output<outputs.Teo.ZoneSettingUpstreamHttp2>;
    /**
     * WebSocket configuration.
     */
    public readonly webSocket!: pulumi.Output<outputs.Teo.ZoneSettingWebSocket>;
    /**
     * Site ID.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a ZoneSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZoneSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZoneSettingArgs | ZoneSettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ZoneSettingState | undefined;
            resourceInputs["area"] = state ? state.area : undefined;
            resourceInputs["cache"] = state ? state.cache : undefined;
            resourceInputs["cacheKey"] = state ? state.cacheKey : undefined;
            resourceInputs["cachePrefresh"] = state ? state.cachePrefresh : undefined;
            resourceInputs["clientIpHeader"] = state ? state.clientIpHeader : undefined;
            resourceInputs["compression"] = state ? state.compression : undefined;
            resourceInputs["forceRedirect"] = state ? state.forceRedirect : undefined;
            resourceInputs["https"] = state ? state.https : undefined;
            resourceInputs["ipv6"] = state ? state.ipv6 : undefined;
            resourceInputs["maxAge"] = state ? state.maxAge : undefined;
            resourceInputs["offlineCache"] = state ? state.offlineCache : undefined;
            resourceInputs["origin"] = state ? state.origin : undefined;
            resourceInputs["postMaxSize"] = state ? state.postMaxSize : undefined;
            resourceInputs["quic"] = state ? state.quic : undefined;
            resourceInputs["smartRouting"] = state ? state.smartRouting : undefined;
            resourceInputs["upstreamHttp2"] = state ? state.upstreamHttp2 : undefined;
            resourceInputs["webSocket"] = state ? state.webSocket : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ZoneSettingArgs | undefined;
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["cache"] = args ? args.cache : undefined;
            resourceInputs["cacheKey"] = args ? args.cacheKey : undefined;
            resourceInputs["cachePrefresh"] = args ? args.cachePrefresh : undefined;
            resourceInputs["clientIpHeader"] = args ? args.clientIpHeader : undefined;
            resourceInputs["compression"] = args ? args.compression : undefined;
            resourceInputs["forceRedirect"] = args ? args.forceRedirect : undefined;
            resourceInputs["https"] = args ? args.https : undefined;
            resourceInputs["ipv6"] = args ? args.ipv6 : undefined;
            resourceInputs["maxAge"] = args ? args.maxAge : undefined;
            resourceInputs["offlineCache"] = args ? args.offlineCache : undefined;
            resourceInputs["origin"] = args ? args.origin : undefined;
            resourceInputs["postMaxSize"] = args ? args.postMaxSize : undefined;
            resourceInputs["quic"] = args ? args.quic : undefined;
            resourceInputs["smartRouting"] = args ? args.smartRouting : undefined;
            resourceInputs["upstreamHttp2"] = args ? args.upstreamHttp2 : undefined;
            resourceInputs["webSocket"] = args ? args.webSocket : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["area"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ZoneSetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ZoneSetting resources.
 */
export interface ZoneSettingState {
    /**
     * Acceleration area of the zone. Valid values: `mainland`, `overseas`.
     */
    area?: pulumi.Input<string>;
    /**
     * Cache expiration time configuration.
     */
    cache?: pulumi.Input<inputs.Teo.ZoneSettingCache>;
    /**
     * Node cache key configuration.
     */
    cacheKey?: pulumi.Input<inputs.Teo.ZoneSettingCacheKey>;
    /**
     * Cache pre-refresh configuration.
     */
    cachePrefresh?: pulumi.Input<inputs.Teo.ZoneSettingCachePrefresh>;
    /**
     * Origin-pull client IP header configuration.
     */
    clientIpHeader?: pulumi.Input<inputs.Teo.ZoneSettingClientIpHeader>;
    /**
     * Smart compression configuration.
     */
    compression?: pulumi.Input<inputs.Teo.ZoneSettingCompression>;
    /**
     * Force HTTPS redirect configuration.
     */
    forceRedirect?: pulumi.Input<inputs.Teo.ZoneSettingForceRedirect>;
    /**
     * HTTPS acceleration configuration.
     */
    https?: pulumi.Input<inputs.Teo.ZoneSettingHttps>;
    /**
     * IPv6 access configuration.
     */
    ipv6?: pulumi.Input<inputs.Teo.ZoneSettingIpv6>;
    /**
     * Browser cache configuration.
     */
    maxAge?: pulumi.Input<inputs.Teo.ZoneSettingMaxAge>;
    /**
     * Offline cache configuration.
     */
    offlineCache?: pulumi.Input<inputs.Teo.ZoneSettingOfflineCache>;
    /**
     * Origin server configuration.
     */
    origin?: pulumi.Input<inputs.Teo.ZoneSettingOrigin>;
    /**
     * Maximum size of files transferred over POST request.
     */
    postMaxSize?: pulumi.Input<inputs.Teo.ZoneSettingPostMaxSize>;
    /**
     * QUIC access configuration.
     */
    quic?: pulumi.Input<inputs.Teo.ZoneSettingQuic>;
    /**
     * Smart acceleration configuration.
     */
    smartRouting?: pulumi.Input<inputs.Teo.ZoneSettingSmartRouting>;
    /**
     * HTTP2 origin-pull configuration.
     */
    upstreamHttp2?: pulumi.Input<inputs.Teo.ZoneSettingUpstreamHttp2>;
    /**
     * WebSocket configuration.
     */
    webSocket?: pulumi.Input<inputs.Teo.ZoneSettingWebSocket>;
    /**
     * Site ID.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ZoneSetting resource.
 */
export interface ZoneSettingArgs {
    /**
     * Cache expiration time configuration.
     */
    cache?: pulumi.Input<inputs.Teo.ZoneSettingCache>;
    /**
     * Node cache key configuration.
     */
    cacheKey?: pulumi.Input<inputs.Teo.ZoneSettingCacheKey>;
    /**
     * Cache pre-refresh configuration.
     */
    cachePrefresh?: pulumi.Input<inputs.Teo.ZoneSettingCachePrefresh>;
    /**
     * Origin-pull client IP header configuration.
     */
    clientIpHeader?: pulumi.Input<inputs.Teo.ZoneSettingClientIpHeader>;
    /**
     * Smart compression configuration.
     */
    compression?: pulumi.Input<inputs.Teo.ZoneSettingCompression>;
    /**
     * Force HTTPS redirect configuration.
     */
    forceRedirect?: pulumi.Input<inputs.Teo.ZoneSettingForceRedirect>;
    /**
     * HTTPS acceleration configuration.
     */
    https?: pulumi.Input<inputs.Teo.ZoneSettingHttps>;
    /**
     * IPv6 access configuration.
     */
    ipv6?: pulumi.Input<inputs.Teo.ZoneSettingIpv6>;
    /**
     * Browser cache configuration.
     */
    maxAge?: pulumi.Input<inputs.Teo.ZoneSettingMaxAge>;
    /**
     * Offline cache configuration.
     */
    offlineCache?: pulumi.Input<inputs.Teo.ZoneSettingOfflineCache>;
    /**
     * Origin server configuration.
     */
    origin?: pulumi.Input<inputs.Teo.ZoneSettingOrigin>;
    /**
     * Maximum size of files transferred over POST request.
     */
    postMaxSize?: pulumi.Input<inputs.Teo.ZoneSettingPostMaxSize>;
    /**
     * QUIC access configuration.
     */
    quic?: pulumi.Input<inputs.Teo.ZoneSettingQuic>;
    /**
     * Smart acceleration configuration.
     */
    smartRouting?: pulumi.Input<inputs.Teo.ZoneSettingSmartRouting>;
    /**
     * HTTP2 origin-pull configuration.
     */
    upstreamHttp2?: pulumi.Input<inputs.Teo.ZoneSettingUpstreamHttp2>;
    /**
     * WebSocket configuration.
     */
    webSocket?: pulumi.Input<inputs.Teo.ZoneSettingWebSocket>;
    /**
     * Site ID.
     */
    zoneId: pulumi.Input<string>;
}
