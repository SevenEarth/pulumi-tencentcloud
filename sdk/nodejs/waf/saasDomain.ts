// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a waf saasDomain
 *
 * ## Example Usage
 *
 * ### If upstreamType is 0
 *
 * Create a basic waf saas domain
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.waf.SaasDomain("example", {
 *     domain: "tf.example.com",
 *     instanceId: "waf_2kxtlbky01b3wceb",
 *     ports: [{
 *         port: "80",
 *         protocol: "http",
 *         upstreamPort: "80",
 *         upstreamProtocol: "http",
 *     }],
 *     srcLists: ["1.1.1.1"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Create a load balancing strategy is weighted polling saas domain
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.waf.SaasDomain("example", {
 *     domain: "tf.example.com",
 *     instanceId: "waf_2kxtlbky01b3wceb",
 *     loadBalance: "2",
 *     ports: [{
 *         port: "80",
 *         protocol: "http",
 *         upstreamPort: "80",
 *         upstreamProtocol: "http",
 *     }],
 *     srcLists: [
 *         "1.1.1.1",
 *         "2.2.2.2",
 *     ],
 *     weights: [
 *         30,
 *         50,
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### If upstreamType is 1
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.waf.SaasDomain("example", {
 *     domain: "tf.example.com",
 *     instanceId: "waf_2kxtlbky01b3wceb",
 *     ports: [{
 *         port: "80",
 *         protocol: "http",
 *         upstreamPort: "80",
 *         upstreamProtocol: "http",
 *     }],
 *     upstreamDomain: "test.com",
 *     upstreamType: 1,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Create a waf saas domain with set Http&Https
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.waf.SaasDomain("example", {
 *     certType: 2,
 *     domain: "tf.example.com",
 *     httpsRewrite: 1,
 *     instanceId: "waf_2kxtlbky01b3wceb",
 *     ipHeaders: [
 *         "headers_1",
 *         "headers_2",
 *         "headers_3",
 *     ],
 *     isCdn: 3,
 *     loadBalance: "2",
 *     ports: [
 *         {
 *             port: "80",
 *             protocol: "http",
 *             upstreamPort: "80",
 *             upstreamProtocol: "http",
 *         },
 *         {
 *             port: "443",
 *             protocol: "https",
 *             upstreamPort: "443",
 *             upstreamProtocol: "https",
 *         },
 *     ],
 *     srcLists: [
 *         "1.1.1.1",
 *         "2.2.2.2",
 *     ],
 *     sslId: "3a6B5y8v",
 *     upstreamScheme: "https",
 *     weights: [
 *         50,
 *         60,
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Create a complete waf saas domain
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.waf.SaasDomain("example", {
 *     activeCheck: 1,
 *     apiSafeStatus: 1,
 *     botStatus: 1,
 *     certType: 2,
 *     cipherTemplate: 1,
 *     domain: "tf.example.com",
 *     httpsRewrite: 1,
 *     instanceId: "waf_2kxtlbky01b3wceb",
 *     ipHeaders: [
 *         "headers_1",
 *         "headers_2",
 *         "headers_3",
 *     ],
 *     isCdn: 3,
 *     isHttp2: 1,
 *     isKeepAlive: "1",
 *     loadBalance: "2",
 *     ports: [
 *         {
 *             port: "80",
 *             protocol: "http",
 *             upstreamPort: "80",
 *             upstreamProtocol: "http",
 *         },
 *         {
 *             port: "443",
 *             protocol: "https",
 *             upstreamPort: "443",
 *             upstreamProtocol: "https",
 *         },
 *     ],
 *     proxyReadTimeout: 500,
 *     proxySendTimeout: 500,
 *     sniHost: "3.3.3.3",
 *     sniType: 3,
 *     srcLists: [
 *         "1.1.1.1",
 *         "2.2.2.2",
 *     ],
 *     sslId: "3a6B5y8v",
 *     tlsVersion: 3,
 *     upstreamScheme: "https",
 *     weights: [
 *         50,
 *         60,
 *     ],
 *     xffReset: 1,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * waf saas_domain can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Waf/saasDomain:SaasDomain example waf_2kxtlbky01b3wceb#tf.example.com#9647c91da0aa5f5aaa49d0ca40e2af24
 * ```
 */
export class SaasDomain extends pulumi.CustomResource {
    /**
     * Get an existing SaasDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SaasDomainState, opts?: pulumi.CustomResourceOptions): SaasDomain {
        return new SaasDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Waf/saasDomain:SaasDomain';

    /**
     * Returns true if the given object is an instance of SaasDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SaasDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SaasDomain.__pulumiType;
    }

    /**
     * Whether to enable active health detection, 0 represents disable and 1 represents enable.
     */
    public readonly activeCheck!: pulumi.Output<number | undefined>;
    /**
     * Whether to enable api safe, 1 enable, 0 disable.
     */
    public readonly apiSafeStatus!: pulumi.Output<number | undefined>;
    /**
     * Whether to enable bot, 1 enable, 0 disable.
     */
    public readonly botStatus!: pulumi.Output<number | undefined>;
    /**
     * Certificate content, When CertType=1, this parameter needs to be filled.
     */
    public readonly cert!: pulumi.Output<string | undefined>;
    /**
     * Certificate type, 0 represents no certificate, CertType=1 represents self owned certificate, and 2 represents managed certificate.
     */
    public readonly certType!: pulumi.Output<number | undefined>;
    /**
     * Encryption Suite Template, 0:default  1:Universal template 2:Security template 3:Custom template.
     */
    public readonly cipherTemplate!: pulumi.Output<number | undefined>;
    /**
     * Encryption Suite Information.
     */
    public readonly ciphers!: pulumi.Output<number[]>;
    /**
     * Whether to enable access logs, 1 enable, 0 disable.
     */
    public readonly clsStatus!: pulumi.Output<number | undefined>;
    /**
     * Domain names that require defense.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Domain id.
     */
    public /*out*/ readonly domainId!: pulumi.Output<string>;
    /**
     * Whether redirect to https, 1 will redirect and 0 will not.
     */
    public readonly httpsRewrite!: pulumi.Output<number | undefined>;
    /**
     * Upstream port for https, When listen ports has https port and UpstreamScheme is HTTP, the current field needs to be filled.
     */
    public readonly httpsUpstreamPort!: pulumi.Output<string | undefined>;
    /**
     * Unique ID of Instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * When is_cdn=3, this parameter needs to be filled in to indicate a custom header.
     */
    public readonly ipHeaders!: pulumi.Output<string[] | undefined>;
    /**
     * Whether a proxy has been enabled before WAF, 0 no deployment, 1 deployment and use first IP in X-Forwarded-For as client IP, 2 deployment and use remoteAddr as client IP, 3 deployment and use values of custom headers as client IP.
     */
    public readonly isCdn!: pulumi.Output<number | undefined>;
    /**
     * Whether enable HTTP2, Enabling HTTP2 requires HTTPS support, 1 means enabled, 0 does not.
     */
    public readonly isHttp2!: pulumi.Output<number | undefined>;
    /**
     * Whether to enable keep-alive, 0 disable, 1 enable.
     */
    public readonly isKeepAlive!: pulumi.Output<string | undefined>;
    /**
     * Is WebSocket support enabled. 1 means enabled, 0 does not.
     */
    public readonly isWebsocket!: pulumi.Output<number | undefined>;
    /**
     * Load balancing strategy, where 0 represents polling and 1 represents IP hash and 2 weighted round robin.
     */
    public readonly loadBalance!: pulumi.Output<string | undefined>;
    /**
     * This field needs to be set for multiple ports in the upstream server.
     */
    public readonly ports!: pulumi.Output<outputs.Waf.SaasDomainPort[]>;
    /**
     * Certificate key, When CertType=1, this parameter needs to be filled.
     */
    public readonly privateKey!: pulumi.Output<string | undefined>;
    /**
     * 300s.
     */
    public readonly proxyReadTimeout!: pulumi.Output<number | undefined>;
    /**
     * 300s.
     */
    public readonly proxySendTimeout!: pulumi.Output<number | undefined>;
    /**
     * When SniType=3, this parameter needs to be filled in to represent a custom host.
     */
    public readonly sniHost!: pulumi.Output<string | undefined>;
    /**
     * Sni type fo upstream, 0:disable SNI; 1:enable SNI and SNI equal original request host; 2:and SNI equal upstream host 3:enable SNI and equal customize host.
     */
    public readonly sniType!: pulumi.Output<number | undefined>;
    /**
     * Upstream IP List, When UpstreamType=0, this parameter needs to be filled.
     */
    public readonly srcLists!: pulumi.Output<string[] | undefined>;
    /**
     * Certificate ID, When CertType=2, this parameter needs to be filled.
     */
    public readonly sslId!: pulumi.Output<string | undefined>;
    /**
     * Binding status between waf and LB, 0:not bind, 1:binding.
     */
    public readonly status!: pulumi.Output<number | undefined>;
    /**
     * Version of TLS Protocol.
     */
    public readonly tlsVersion!: pulumi.Output<number | undefined>;
    /**
     * Upstream domain, When UpstreamType=1, this parameter needs to be filled.
     */
    public readonly upstreamDomain!: pulumi.Output<string | undefined>;
    /**
     * Upstream scheme for https, http or https.
     */
    public readonly upstreamScheme!: pulumi.Output<string | undefined>;
    /**
     * Upstream type, 0 represents IP, 1 represents domain name.
     */
    public readonly upstreamType!: pulumi.Output<number | undefined>;
    /**
     * Weight of each upstream.
     */
    public readonly weights!: pulumi.Output<number[] | undefined>;
    /**
     * 0:disable xff reset; 1:enable xff reset.
     */
    public readonly xffReset!: pulumi.Output<number | undefined>;

    /**
     * Create a SaasDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SaasDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SaasDomainArgs | SaasDomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SaasDomainState | undefined;
            resourceInputs["activeCheck"] = state ? state.activeCheck : undefined;
            resourceInputs["apiSafeStatus"] = state ? state.apiSafeStatus : undefined;
            resourceInputs["botStatus"] = state ? state.botStatus : undefined;
            resourceInputs["cert"] = state ? state.cert : undefined;
            resourceInputs["certType"] = state ? state.certType : undefined;
            resourceInputs["cipherTemplate"] = state ? state.cipherTemplate : undefined;
            resourceInputs["ciphers"] = state ? state.ciphers : undefined;
            resourceInputs["clsStatus"] = state ? state.clsStatus : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["domainId"] = state ? state.domainId : undefined;
            resourceInputs["httpsRewrite"] = state ? state.httpsRewrite : undefined;
            resourceInputs["httpsUpstreamPort"] = state ? state.httpsUpstreamPort : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["ipHeaders"] = state ? state.ipHeaders : undefined;
            resourceInputs["isCdn"] = state ? state.isCdn : undefined;
            resourceInputs["isHttp2"] = state ? state.isHttp2 : undefined;
            resourceInputs["isKeepAlive"] = state ? state.isKeepAlive : undefined;
            resourceInputs["isWebsocket"] = state ? state.isWebsocket : undefined;
            resourceInputs["loadBalance"] = state ? state.loadBalance : undefined;
            resourceInputs["ports"] = state ? state.ports : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["proxyReadTimeout"] = state ? state.proxyReadTimeout : undefined;
            resourceInputs["proxySendTimeout"] = state ? state.proxySendTimeout : undefined;
            resourceInputs["sniHost"] = state ? state.sniHost : undefined;
            resourceInputs["sniType"] = state ? state.sniType : undefined;
            resourceInputs["srcLists"] = state ? state.srcLists : undefined;
            resourceInputs["sslId"] = state ? state.sslId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tlsVersion"] = state ? state.tlsVersion : undefined;
            resourceInputs["upstreamDomain"] = state ? state.upstreamDomain : undefined;
            resourceInputs["upstreamScheme"] = state ? state.upstreamScheme : undefined;
            resourceInputs["upstreamType"] = state ? state.upstreamType : undefined;
            resourceInputs["weights"] = state ? state.weights : undefined;
            resourceInputs["xffReset"] = state ? state.xffReset : undefined;
        } else {
            const args = argsOrState as SaasDomainArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.ports === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ports'");
            }
            resourceInputs["activeCheck"] = args ? args.activeCheck : undefined;
            resourceInputs["apiSafeStatus"] = args ? args.apiSafeStatus : undefined;
            resourceInputs["botStatus"] = args ? args.botStatus : undefined;
            resourceInputs["cert"] = args ? args.cert : undefined;
            resourceInputs["certType"] = args ? args.certType : undefined;
            resourceInputs["cipherTemplate"] = args ? args.cipherTemplate : undefined;
            resourceInputs["ciphers"] = args ? args.ciphers : undefined;
            resourceInputs["clsStatus"] = args ? args.clsStatus : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["httpsRewrite"] = args ? args.httpsRewrite : undefined;
            resourceInputs["httpsUpstreamPort"] = args ? args.httpsUpstreamPort : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["ipHeaders"] = args ? args.ipHeaders : undefined;
            resourceInputs["isCdn"] = args ? args.isCdn : undefined;
            resourceInputs["isHttp2"] = args ? args.isHttp2 : undefined;
            resourceInputs["isKeepAlive"] = args ? args.isKeepAlive : undefined;
            resourceInputs["isWebsocket"] = args ? args.isWebsocket : undefined;
            resourceInputs["loadBalance"] = args ? args.loadBalance : undefined;
            resourceInputs["ports"] = args ? args.ports : undefined;
            resourceInputs["privateKey"] = args ? args.privateKey : undefined;
            resourceInputs["proxyReadTimeout"] = args ? args.proxyReadTimeout : undefined;
            resourceInputs["proxySendTimeout"] = args ? args.proxySendTimeout : undefined;
            resourceInputs["sniHost"] = args ? args.sniHost : undefined;
            resourceInputs["sniType"] = args ? args.sniType : undefined;
            resourceInputs["srcLists"] = args ? args.srcLists : undefined;
            resourceInputs["sslId"] = args ? args.sslId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tlsVersion"] = args ? args.tlsVersion : undefined;
            resourceInputs["upstreamDomain"] = args ? args.upstreamDomain : undefined;
            resourceInputs["upstreamScheme"] = args ? args.upstreamScheme : undefined;
            resourceInputs["upstreamType"] = args ? args.upstreamType : undefined;
            resourceInputs["weights"] = args ? args.weights : undefined;
            resourceInputs["xffReset"] = args ? args.xffReset : undefined;
            resourceInputs["domainId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SaasDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SaasDomain resources.
 */
export interface SaasDomainState {
    /**
     * Whether to enable active health detection, 0 represents disable and 1 represents enable.
     */
    activeCheck?: pulumi.Input<number>;
    /**
     * Whether to enable api safe, 1 enable, 0 disable.
     */
    apiSafeStatus?: pulumi.Input<number>;
    /**
     * Whether to enable bot, 1 enable, 0 disable.
     */
    botStatus?: pulumi.Input<number>;
    /**
     * Certificate content, When CertType=1, this parameter needs to be filled.
     */
    cert?: pulumi.Input<string>;
    /**
     * Certificate type, 0 represents no certificate, CertType=1 represents self owned certificate, and 2 represents managed certificate.
     */
    certType?: pulumi.Input<number>;
    /**
     * Encryption Suite Template, 0:default  1:Universal template 2:Security template 3:Custom template.
     */
    cipherTemplate?: pulumi.Input<number>;
    /**
     * Encryption Suite Information.
     */
    ciphers?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Whether to enable access logs, 1 enable, 0 disable.
     */
    clsStatus?: pulumi.Input<number>;
    /**
     * Domain names that require defense.
     */
    domain?: pulumi.Input<string>;
    /**
     * Domain id.
     */
    domainId?: pulumi.Input<string>;
    /**
     * Whether redirect to https, 1 will redirect and 0 will not.
     */
    httpsRewrite?: pulumi.Input<number>;
    /**
     * Upstream port for https, When listen ports has https port and UpstreamScheme is HTTP, the current field needs to be filled.
     */
    httpsUpstreamPort?: pulumi.Input<string>;
    /**
     * Unique ID of Instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * When is_cdn=3, this parameter needs to be filled in to indicate a custom header.
     */
    ipHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether a proxy has been enabled before WAF, 0 no deployment, 1 deployment and use first IP in X-Forwarded-For as client IP, 2 deployment and use remoteAddr as client IP, 3 deployment and use values of custom headers as client IP.
     */
    isCdn?: pulumi.Input<number>;
    /**
     * Whether enable HTTP2, Enabling HTTP2 requires HTTPS support, 1 means enabled, 0 does not.
     */
    isHttp2?: pulumi.Input<number>;
    /**
     * Whether to enable keep-alive, 0 disable, 1 enable.
     */
    isKeepAlive?: pulumi.Input<string>;
    /**
     * Is WebSocket support enabled. 1 means enabled, 0 does not.
     */
    isWebsocket?: pulumi.Input<number>;
    /**
     * Load balancing strategy, where 0 represents polling and 1 represents IP hash and 2 weighted round robin.
     */
    loadBalance?: pulumi.Input<string>;
    /**
     * This field needs to be set for multiple ports in the upstream server.
     */
    ports?: pulumi.Input<pulumi.Input<inputs.Waf.SaasDomainPort>[]>;
    /**
     * Certificate key, When CertType=1, this parameter needs to be filled.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * 300s.
     */
    proxyReadTimeout?: pulumi.Input<number>;
    /**
     * 300s.
     */
    proxySendTimeout?: pulumi.Input<number>;
    /**
     * When SniType=3, this parameter needs to be filled in to represent a custom host.
     */
    sniHost?: pulumi.Input<string>;
    /**
     * Sni type fo upstream, 0:disable SNI; 1:enable SNI and SNI equal original request host; 2:and SNI equal upstream host 3:enable SNI and equal customize host.
     */
    sniType?: pulumi.Input<number>;
    /**
     * Upstream IP List, When UpstreamType=0, this parameter needs to be filled.
     */
    srcLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Certificate ID, When CertType=2, this parameter needs to be filled.
     */
    sslId?: pulumi.Input<string>;
    /**
     * Binding status between waf and LB, 0:not bind, 1:binding.
     */
    status?: pulumi.Input<number>;
    /**
     * Version of TLS Protocol.
     */
    tlsVersion?: pulumi.Input<number>;
    /**
     * Upstream domain, When UpstreamType=1, this parameter needs to be filled.
     */
    upstreamDomain?: pulumi.Input<string>;
    /**
     * Upstream scheme for https, http or https.
     */
    upstreamScheme?: pulumi.Input<string>;
    /**
     * Upstream type, 0 represents IP, 1 represents domain name.
     */
    upstreamType?: pulumi.Input<number>;
    /**
     * Weight of each upstream.
     */
    weights?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * 0:disable xff reset; 1:enable xff reset.
     */
    xffReset?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SaasDomain resource.
 */
export interface SaasDomainArgs {
    /**
     * Whether to enable active health detection, 0 represents disable and 1 represents enable.
     */
    activeCheck?: pulumi.Input<number>;
    /**
     * Whether to enable api safe, 1 enable, 0 disable.
     */
    apiSafeStatus?: pulumi.Input<number>;
    /**
     * Whether to enable bot, 1 enable, 0 disable.
     */
    botStatus?: pulumi.Input<number>;
    /**
     * Certificate content, When CertType=1, this parameter needs to be filled.
     */
    cert?: pulumi.Input<string>;
    /**
     * Certificate type, 0 represents no certificate, CertType=1 represents self owned certificate, and 2 represents managed certificate.
     */
    certType?: pulumi.Input<number>;
    /**
     * Encryption Suite Template, 0:default  1:Universal template 2:Security template 3:Custom template.
     */
    cipherTemplate?: pulumi.Input<number>;
    /**
     * Encryption Suite Information.
     */
    ciphers?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Whether to enable access logs, 1 enable, 0 disable.
     */
    clsStatus?: pulumi.Input<number>;
    /**
     * Domain names that require defense.
     */
    domain: pulumi.Input<string>;
    /**
     * Whether redirect to https, 1 will redirect and 0 will not.
     */
    httpsRewrite?: pulumi.Input<number>;
    /**
     * Upstream port for https, When listen ports has https port and UpstreamScheme is HTTP, the current field needs to be filled.
     */
    httpsUpstreamPort?: pulumi.Input<string>;
    /**
     * Unique ID of Instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * When is_cdn=3, this parameter needs to be filled in to indicate a custom header.
     */
    ipHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether a proxy has been enabled before WAF, 0 no deployment, 1 deployment and use first IP in X-Forwarded-For as client IP, 2 deployment and use remoteAddr as client IP, 3 deployment and use values of custom headers as client IP.
     */
    isCdn?: pulumi.Input<number>;
    /**
     * Whether enable HTTP2, Enabling HTTP2 requires HTTPS support, 1 means enabled, 0 does not.
     */
    isHttp2?: pulumi.Input<number>;
    /**
     * Whether to enable keep-alive, 0 disable, 1 enable.
     */
    isKeepAlive?: pulumi.Input<string>;
    /**
     * Is WebSocket support enabled. 1 means enabled, 0 does not.
     */
    isWebsocket?: pulumi.Input<number>;
    /**
     * Load balancing strategy, where 0 represents polling and 1 represents IP hash and 2 weighted round robin.
     */
    loadBalance?: pulumi.Input<string>;
    /**
     * This field needs to be set for multiple ports in the upstream server.
     */
    ports: pulumi.Input<pulumi.Input<inputs.Waf.SaasDomainPort>[]>;
    /**
     * Certificate key, When CertType=1, this parameter needs to be filled.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * 300s.
     */
    proxyReadTimeout?: pulumi.Input<number>;
    /**
     * 300s.
     */
    proxySendTimeout?: pulumi.Input<number>;
    /**
     * When SniType=3, this parameter needs to be filled in to represent a custom host.
     */
    sniHost?: pulumi.Input<string>;
    /**
     * Sni type fo upstream, 0:disable SNI; 1:enable SNI and SNI equal original request host; 2:and SNI equal upstream host 3:enable SNI and equal customize host.
     */
    sniType?: pulumi.Input<number>;
    /**
     * Upstream IP List, When UpstreamType=0, this parameter needs to be filled.
     */
    srcLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Certificate ID, When CertType=2, this parameter needs to be filled.
     */
    sslId?: pulumi.Input<string>;
    /**
     * Binding status between waf and LB, 0:not bind, 1:binding.
     */
    status?: pulumi.Input<number>;
    /**
     * Version of TLS Protocol.
     */
    tlsVersion?: pulumi.Input<number>;
    /**
     * Upstream domain, When UpstreamType=1, this parameter needs to be filled.
     */
    upstreamDomain?: pulumi.Input<string>;
    /**
     * Upstream scheme for https, http or https.
     */
    upstreamScheme?: pulumi.Input<string>;
    /**
     * Upstream type, 0 represents IP, 1 represents domain name.
     */
    upstreamType?: pulumi.Input<number>;
    /**
     * Weight of each upstream.
     */
    weights?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * 0:disable xff reset; 1:enable xff reset.
     */
    xffReset?: pulumi.Input<number>;
}
