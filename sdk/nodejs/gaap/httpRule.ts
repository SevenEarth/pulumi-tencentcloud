// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a forward rule of layer7 listener.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const fooProxy = new tencentcloud.gaap.Proxy("fooProxy", {
 *     bandwidth: 10,
 *     concurrent: 2,
 *     accessRegion: "SouthChina",
 *     realserverRegion: "NorthChina",
 * });
 * const fooLayer7Listener = new tencentcloud.gaap.Layer7Listener("fooLayer7Listener", {
 *     protocol: "HTTP",
 *     port: 80,
 *     proxyId: fooProxy.id,
 * });
 * const fooRealserver = new tencentcloud.gaap.Realserver("fooRealserver", {ip: "1.1.1.1"});
 * const bar = new tencentcloud.gaap.Realserver("bar", {ip: "8.8.8.8"});
 * const fooHttpDomain = new tencentcloud.gaap.HttpDomain("fooHttpDomain", {
 *     listenerId: fooLayer7Listener.id,
 *     domain: "www.qq.com",
 * });
 * const fooHttpRule = new tencentcloud.gaap.HttpRule("fooHttpRule", {
 *     listenerId: fooLayer7Listener.id,
 *     domain: fooHttpDomain.domain,
 *     path: "/",
 *     realserverType: "IP",
 *     healthCheck: true,
 *     healthCheckPath: "/",
 *     healthCheckMethod: "GET",
 *     healthCheckStatusCodes: [200],
 *     realservers: [
 *         {
 *             id: fooRealserver.id,
 *             ip: fooRealserver.ip,
 *             port: 80,
 *         },
 *         {
 *             id: bar.id,
 *             ip: bar.ip,
 *             port: 80,
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * GAAP http rule can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Gaap/httpRule:HttpRule tencentcloud_gaap_http_rule.foo rule-3bsuu01r
 * ```
 */
export class HttpRule extends pulumi.CustomResource {
    /**
     * Get an existing HttpRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HttpRuleState, opts?: pulumi.CustomResourceOptions): HttpRule {
        return new HttpRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Gaap/httpRule:HttpRule';

    /**
     * Returns true if the given object is an instance of HttpRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HttpRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HttpRule.__pulumiType;
    }

    /**
     * Timeout of the health check response, default value is 2s.
     */
    public readonly connectTimeout!: pulumi.Output<number | undefined>;
    /**
     * Forward domain of the forward rule.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * The default value of requested host which is forwarded to the realserver by the listener is `default`.
     */
    public readonly forwardHost!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether health check is enable.
     */
    public readonly healthCheck!: pulumi.Output<boolean>;
    /**
     * Method of the health check. Valid value: `GET` and `HEAD`.
     */
    public readonly healthCheckMethod!: pulumi.Output<string | undefined>;
    /**
     * Path of health check. Maximum length is 80.
     */
    public readonly healthCheckPath!: pulumi.Output<string | undefined>;
    /**
     * Return code of confirmed normal. Valid value: `100`, `200`, `300`, `400` and `500`.
     */
    public readonly healthCheckStatusCodes!: pulumi.Output<number[]>;
    /**
     * Interval of the health check, default value is 5s.
     */
    public readonly interval!: pulumi.Output<number | undefined>;
    /**
     * ID of the layer7 listener.
     */
    public readonly listenerId!: pulumi.Output<string>;
    /**
     * Path of the forward rule. Maximum length is 80.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * Type of the realserver. Valid value: `IP` and `DOMAIN`.
     */
    public readonly realserverType!: pulumi.Output<string>;
    /**
     * An information list of GAAP realserver.
     */
    public readonly realservers!: pulumi.Output<outputs.Gaap.HttpRuleRealserver[] | undefined>;
    /**
     * Scheduling policy of the forward rule, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
     */
    public readonly scheduler!: pulumi.Output<string | undefined>;
    /**
     * ServerNameIndication (SNI) is required when the SNI switch is turned on.
     */
    public readonly sni!: pulumi.Output<string>;
    /**
     * ServerNameIndication (SNI) switch. ON means on and OFF means off.
     */
    public readonly sniSwitch!: pulumi.Output<string>;

    /**
     * Create a HttpRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HttpRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HttpRuleArgs | HttpRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HttpRuleState | undefined;
            resourceInputs["connectTimeout"] = state ? state.connectTimeout : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["forwardHost"] = state ? state.forwardHost : undefined;
            resourceInputs["healthCheck"] = state ? state.healthCheck : undefined;
            resourceInputs["healthCheckMethod"] = state ? state.healthCheckMethod : undefined;
            resourceInputs["healthCheckPath"] = state ? state.healthCheckPath : undefined;
            resourceInputs["healthCheckStatusCodes"] = state ? state.healthCheckStatusCodes : undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["realserverType"] = state ? state.realserverType : undefined;
            resourceInputs["realservers"] = state ? state.realservers : undefined;
            resourceInputs["scheduler"] = state ? state.scheduler : undefined;
            resourceInputs["sni"] = state ? state.sni : undefined;
            resourceInputs["sniSwitch"] = state ? state.sniSwitch : undefined;
        } else {
            const args = argsOrState as HttpRuleArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.healthCheck === undefined) && !opts.urn) {
                throw new Error("Missing required property 'healthCheck'");
            }
            if ((!args || args.listenerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerId'");
            }
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            if ((!args || args.realserverType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realserverType'");
            }
            resourceInputs["connectTimeout"] = args ? args.connectTimeout : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["forwardHost"] = args ? args.forwardHost : undefined;
            resourceInputs["healthCheck"] = args ? args.healthCheck : undefined;
            resourceInputs["healthCheckMethod"] = args ? args.healthCheckMethod : undefined;
            resourceInputs["healthCheckPath"] = args ? args.healthCheckPath : undefined;
            resourceInputs["healthCheckStatusCodes"] = args ? args.healthCheckStatusCodes : undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["listenerId"] = args ? args.listenerId : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["realserverType"] = args ? args.realserverType : undefined;
            resourceInputs["realservers"] = args ? args.realservers : undefined;
            resourceInputs["scheduler"] = args ? args.scheduler : undefined;
            resourceInputs["sni"] = args ? args.sni : undefined;
            resourceInputs["sniSwitch"] = args ? args.sniSwitch : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HttpRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HttpRule resources.
 */
export interface HttpRuleState {
    /**
     * Timeout of the health check response, default value is 2s.
     */
    connectTimeout?: pulumi.Input<number>;
    /**
     * Forward domain of the forward rule.
     */
    domain?: pulumi.Input<string>;
    /**
     * The default value of requested host which is forwarded to the realserver by the listener is `default`.
     */
    forwardHost?: pulumi.Input<string>;
    /**
     * Indicates whether health check is enable.
     */
    healthCheck?: pulumi.Input<boolean>;
    /**
     * Method of the health check. Valid value: `GET` and `HEAD`.
     */
    healthCheckMethod?: pulumi.Input<string>;
    /**
     * Path of health check. Maximum length is 80.
     */
    healthCheckPath?: pulumi.Input<string>;
    /**
     * Return code of confirmed normal. Valid value: `100`, `200`, `300`, `400` and `500`.
     */
    healthCheckStatusCodes?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Interval of the health check, default value is 5s.
     */
    interval?: pulumi.Input<number>;
    /**
     * ID of the layer7 listener.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * Path of the forward rule. Maximum length is 80.
     */
    path?: pulumi.Input<string>;
    /**
     * Type of the realserver. Valid value: `IP` and `DOMAIN`.
     */
    realserverType?: pulumi.Input<string>;
    /**
     * An information list of GAAP realserver.
     */
    realservers?: pulumi.Input<pulumi.Input<inputs.Gaap.HttpRuleRealserver>[]>;
    /**
     * Scheduling policy of the forward rule, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
     */
    scheduler?: pulumi.Input<string>;
    /**
     * ServerNameIndication (SNI) is required when the SNI switch is turned on.
     */
    sni?: pulumi.Input<string>;
    /**
     * ServerNameIndication (SNI) switch. ON means on and OFF means off.
     */
    sniSwitch?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HttpRule resource.
 */
export interface HttpRuleArgs {
    /**
     * Timeout of the health check response, default value is 2s.
     */
    connectTimeout?: pulumi.Input<number>;
    /**
     * Forward domain of the forward rule.
     */
    domain: pulumi.Input<string>;
    /**
     * The default value of requested host which is forwarded to the realserver by the listener is `default`.
     */
    forwardHost?: pulumi.Input<string>;
    /**
     * Indicates whether health check is enable.
     */
    healthCheck: pulumi.Input<boolean>;
    /**
     * Method of the health check. Valid value: `GET` and `HEAD`.
     */
    healthCheckMethod?: pulumi.Input<string>;
    /**
     * Path of health check. Maximum length is 80.
     */
    healthCheckPath?: pulumi.Input<string>;
    /**
     * Return code of confirmed normal. Valid value: `100`, `200`, `300`, `400` and `500`.
     */
    healthCheckStatusCodes?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Interval of the health check, default value is 5s.
     */
    interval?: pulumi.Input<number>;
    /**
     * ID of the layer7 listener.
     */
    listenerId: pulumi.Input<string>;
    /**
     * Path of the forward rule. Maximum length is 80.
     */
    path: pulumi.Input<string>;
    /**
     * Type of the realserver. Valid value: `IP` and `DOMAIN`.
     */
    realserverType: pulumi.Input<string>;
    /**
     * An information list of GAAP realserver.
     */
    realservers?: pulumi.Input<pulumi.Input<inputs.Gaap.HttpRuleRealserver>[]>;
    /**
     * Scheduling policy of the forward rule, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
     */
    scheduler?: pulumi.Input<string>;
    /**
     * ServerNameIndication (SNI) is required when the SNI switch is turned on.
     */
    sni?: pulumi.Input<string>;
    /**
     * ServerNameIndication (SNI) switch. ON means on and OFF means off.
     */
    sniSwitch?: pulumi.Input<string>;
}
