// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a CLB listener rule.
 *
 * > **NOTE:** This resource only be applied to the HTTP or HTTPS listeners.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.Clb.ListenerRule("foo", {
 *     certificateCaId: "VfqO4zkB",
 *     certificateId: "VjANRdz8",
 *     certificateSslMode: "MUTUAL",
 *     clbId: "lb-k2zjp9lv",
 *     domain: "foo.net",
 *     healthCheckHealthNum: 3,
 *     healthCheckHttpCode: 2,
 *     healthCheckHttpDomain: "Default Domain",
 *     healthCheckHttpMethod: "GET",
 *     healthCheckHttpPath: "Default Path",
 *     healthCheckIntervalTime: 5,
 *     healthCheckSwitch: true,
 *     healthCheckUnhealthNum: 3,
 *     listenerId: "lbl-hh141sn9",
 *     scheduler: "WRR",
 *     sessionExpireTime: 30,
 *     url: "/bar",
 * });
 * ```
 *
 * ## Import
 *
 * CLB listener rule can be imported using the id (version >= 1.47.0), e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Clb/listenerRule:ListenerRule foo lb-7a0t6zqb#lbl-hh141sn9#loc-agg236ys
 * ```
 */
export class ListenerRule extends pulumi.CustomResource {
    /**
     * Get an existing ListenerRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerRuleState, opts?: pulumi.CustomResourceOptions): ListenerRule {
        return new ListenerRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Clb/listenerRule:ListenerRule';

    /**
     * Returns true if the given object is an instance of ListenerRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ListenerRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ListenerRule.__pulumiType;
    }

    /**
     * ID of the client certificate. NOTES: Only supports listeners of HTTPS protocol.
     */
    public readonly certificateCaId!: pulumi.Output<string | undefined>;
    /**
     * ID of the server certificate. NOTES: Only supports listeners of HTTPS protocol.
     */
    public readonly certificateId!: pulumi.Output<string | undefined>;
    /**
     * Type of certificate. Valid values: `UNIDIRECTIONAL`, `MUTUAL`. NOTES: Only supports listeners of HTTPS protocol.
     */
    public readonly certificateSslMode!: pulumi.Output<string | undefined>;
    /**
     * ID of CLB instance.
     */
    public readonly clbId!: pulumi.Output<string>;
    /**
     * Domain name of the listener rule.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Forwarding protocol between the CLB instance and real server. Valid values: `HTTP`, `HTTPS`, `TRPC`.
     */
    public readonly forwardType!: pulumi.Output<string>;
    /**
     * Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10]. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
     */
    public readonly healthCheckHealthNum!: pulumi.Output<number>;
    /**
     * HTTP Status Code. The default is 31. Valid value ranges: [1~31]. `1 means the return value '1xx' is health. `2` means the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.
     */
    public readonly healthCheckHttpCode!: pulumi.Output<number>;
    /**
     * Domain name of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
     */
    public readonly healthCheckHttpDomain!: pulumi.Output<string>;
    /**
     * Methods of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol. The default is `HEAD`, the available value are `HEAD` and `GET`.
     */
    public readonly healthCheckHttpMethod!: pulumi.Output<string>;
    /**
     * Path of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
     */
    public readonly healthCheckHttpPath!: pulumi.Output<string>;
    /**
     * Interval time of health check. Valid value ranges: (2~300) sec. and the default is `5` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
     */
    public readonly healthCheckIntervalTime!: pulumi.Output<number>;
    /**
     * Indicates whether health check is enabled.
     */
    public readonly healthCheckSwitch!: pulumi.Output<boolean>;
    /**
     * Time out of health check. The value range is 2-60.
     */
    public readonly healthCheckTimeOut!: pulumi.Output<number>;
    /**
     * Type of health check. Valid value is `CUSTOM`, `TCP`, `HTTP`.
     */
    public readonly healthCheckType!: pulumi.Output<string>;
    /**
     * Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
     */
    public readonly healthCheckUnhealthNum!: pulumi.Output<number>;
    /**
     * Indicate to apply HTTP2.0 protocol or not.
     */
    public readonly http2Switch!: pulumi.Output<boolean>;
    /**
     * ID of CLB listener.
     */
    public readonly listenerId!: pulumi.Output<string>;
    /**
     * ID of this CLB listener rule.
     */
    public /*out*/ readonly ruleId!: pulumi.Output<string>;
    /**
     * Scheduling method of the CLB listener rules. Valid values: `WRR`, `IP HASH`, `LEAST_CONN`. The default is `WRR`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
     */
    public readonly scheduler!: pulumi.Output<string | undefined>;
    /**
     * Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as `WRR`, and not available when listener protocol is `TCP_SSL`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
     */
    public readonly sessionExpireTime!: pulumi.Output<number | undefined>;
    /**
     * Backend target type. Valid values: `NODE`, `TARGETGROUP`. `NODE` means to bind ordinary nodes, `TARGETGROUP` means to bind target group.
     */
    public readonly targetType!: pulumi.Output<string | undefined>;
    /**
     * Url of the listener rule.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ListenerRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListenerRuleArgs | ListenerRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ListenerRuleState | undefined;
            resourceInputs["certificateCaId"] = state ? state.certificateCaId : undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["certificateSslMode"] = state ? state.certificateSslMode : undefined;
            resourceInputs["clbId"] = state ? state.clbId : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["forwardType"] = state ? state.forwardType : undefined;
            resourceInputs["healthCheckHealthNum"] = state ? state.healthCheckHealthNum : undefined;
            resourceInputs["healthCheckHttpCode"] = state ? state.healthCheckHttpCode : undefined;
            resourceInputs["healthCheckHttpDomain"] = state ? state.healthCheckHttpDomain : undefined;
            resourceInputs["healthCheckHttpMethod"] = state ? state.healthCheckHttpMethod : undefined;
            resourceInputs["healthCheckHttpPath"] = state ? state.healthCheckHttpPath : undefined;
            resourceInputs["healthCheckIntervalTime"] = state ? state.healthCheckIntervalTime : undefined;
            resourceInputs["healthCheckSwitch"] = state ? state.healthCheckSwitch : undefined;
            resourceInputs["healthCheckTimeOut"] = state ? state.healthCheckTimeOut : undefined;
            resourceInputs["healthCheckType"] = state ? state.healthCheckType : undefined;
            resourceInputs["healthCheckUnhealthNum"] = state ? state.healthCheckUnhealthNum : undefined;
            resourceInputs["http2Switch"] = state ? state.http2Switch : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["scheduler"] = state ? state.scheduler : undefined;
            resourceInputs["sessionExpireTime"] = state ? state.sessionExpireTime : undefined;
            resourceInputs["targetType"] = state ? state.targetType : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ListenerRuleArgs | undefined;
            if ((!args || args.clbId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clbId'");
            }
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.listenerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerId'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["certificateCaId"] = args ? args.certificateCaId : undefined;
            resourceInputs["certificateId"] = args ? args.certificateId : undefined;
            resourceInputs["certificateSslMode"] = args ? args.certificateSslMode : undefined;
            resourceInputs["clbId"] = args ? args.clbId : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["forwardType"] = args ? args.forwardType : undefined;
            resourceInputs["healthCheckHealthNum"] = args ? args.healthCheckHealthNum : undefined;
            resourceInputs["healthCheckHttpCode"] = args ? args.healthCheckHttpCode : undefined;
            resourceInputs["healthCheckHttpDomain"] = args ? args.healthCheckHttpDomain : undefined;
            resourceInputs["healthCheckHttpMethod"] = args ? args.healthCheckHttpMethod : undefined;
            resourceInputs["healthCheckHttpPath"] = args ? args.healthCheckHttpPath : undefined;
            resourceInputs["healthCheckIntervalTime"] = args ? args.healthCheckIntervalTime : undefined;
            resourceInputs["healthCheckSwitch"] = args ? args.healthCheckSwitch : undefined;
            resourceInputs["healthCheckTimeOut"] = args ? args.healthCheckTimeOut : undefined;
            resourceInputs["healthCheckType"] = args ? args.healthCheckType : undefined;
            resourceInputs["healthCheckUnhealthNum"] = args ? args.healthCheckUnhealthNum : undefined;
            resourceInputs["http2Switch"] = args ? args.http2Switch : undefined;
            resourceInputs["listenerId"] = args ? args.listenerId : undefined;
            resourceInputs["scheduler"] = args ? args.scheduler : undefined;
            resourceInputs["sessionExpireTime"] = args ? args.sessionExpireTime : undefined;
            resourceInputs["targetType"] = args ? args.targetType : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["ruleId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ListenerRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ListenerRule resources.
 */
export interface ListenerRuleState {
    /**
     * ID of the client certificate. NOTES: Only supports listeners of HTTPS protocol.
     */
    certificateCaId?: pulumi.Input<string>;
    /**
     * ID of the server certificate. NOTES: Only supports listeners of HTTPS protocol.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * Type of certificate. Valid values: `UNIDIRECTIONAL`, `MUTUAL`. NOTES: Only supports listeners of HTTPS protocol.
     */
    certificateSslMode?: pulumi.Input<string>;
    /**
     * ID of CLB instance.
     */
    clbId?: pulumi.Input<string>;
    /**
     * Domain name of the listener rule.
     */
    domain?: pulumi.Input<string>;
    /**
     * Forwarding protocol between the CLB instance and real server. Valid values: `HTTP`, `HTTPS`, `TRPC`.
     */
    forwardType?: pulumi.Input<string>;
    /**
     * Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10]. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
     */
    healthCheckHealthNum?: pulumi.Input<number>;
    /**
     * HTTP Status Code. The default is 31. Valid value ranges: [1~31]. `1 means the return value '1xx' is health. `2` means the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.
     */
    healthCheckHttpCode?: pulumi.Input<number>;
    /**
     * Domain name of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
     */
    healthCheckHttpDomain?: pulumi.Input<string>;
    /**
     * Methods of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol. The default is `HEAD`, the available value are `HEAD` and `GET`.
     */
    healthCheckHttpMethod?: pulumi.Input<string>;
    /**
     * Path of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
     */
    healthCheckHttpPath?: pulumi.Input<string>;
    /**
     * Interval time of health check. Valid value ranges: (2~300) sec. and the default is `5` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
     */
    healthCheckIntervalTime?: pulumi.Input<number>;
    /**
     * Indicates whether health check is enabled.
     */
    healthCheckSwitch?: pulumi.Input<boolean>;
    /**
     * Time out of health check. The value range is 2-60.
     */
    healthCheckTimeOut?: pulumi.Input<number>;
    /**
     * Type of health check. Valid value is `CUSTOM`, `TCP`, `HTTP`.
     */
    healthCheckType?: pulumi.Input<string>;
    /**
     * Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
     */
    healthCheckUnhealthNum?: pulumi.Input<number>;
    /**
     * Indicate to apply HTTP2.0 protocol or not.
     */
    http2Switch?: pulumi.Input<boolean>;
    /**
     * ID of CLB listener.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * ID of this CLB listener rule.
     */
    ruleId?: pulumi.Input<string>;
    /**
     * Scheduling method of the CLB listener rules. Valid values: `WRR`, `IP HASH`, `LEAST_CONN`. The default is `WRR`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
     */
    scheduler?: pulumi.Input<string>;
    /**
     * Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as `WRR`, and not available when listener protocol is `TCP_SSL`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
     */
    sessionExpireTime?: pulumi.Input<number>;
    /**
     * Backend target type. Valid values: `NODE`, `TARGETGROUP`. `NODE` means to bind ordinary nodes, `TARGETGROUP` means to bind target group.
     */
    targetType?: pulumi.Input<string>;
    /**
     * Url of the listener rule.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ListenerRule resource.
 */
export interface ListenerRuleArgs {
    /**
     * ID of the client certificate. NOTES: Only supports listeners of HTTPS protocol.
     */
    certificateCaId?: pulumi.Input<string>;
    /**
     * ID of the server certificate. NOTES: Only supports listeners of HTTPS protocol.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * Type of certificate. Valid values: `UNIDIRECTIONAL`, `MUTUAL`. NOTES: Only supports listeners of HTTPS protocol.
     */
    certificateSslMode?: pulumi.Input<string>;
    /**
     * ID of CLB instance.
     */
    clbId: pulumi.Input<string>;
    /**
     * Domain name of the listener rule.
     */
    domain: pulumi.Input<string>;
    /**
     * Forwarding protocol between the CLB instance and real server. Valid values: `HTTP`, `HTTPS`, `TRPC`.
     */
    forwardType?: pulumi.Input<string>;
    /**
     * Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10]. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
     */
    healthCheckHealthNum?: pulumi.Input<number>;
    /**
     * HTTP Status Code. The default is 31. Valid value ranges: [1~31]. `1 means the return value '1xx' is health. `2` means the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.
     */
    healthCheckHttpCode?: pulumi.Input<number>;
    /**
     * Domain name of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
     */
    healthCheckHttpDomain?: pulumi.Input<string>;
    /**
     * Methods of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol. The default is `HEAD`, the available value are `HEAD` and `GET`.
     */
    healthCheckHttpMethod?: pulumi.Input<string>;
    /**
     * Path of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
     */
    healthCheckHttpPath?: pulumi.Input<string>;
    /**
     * Interval time of health check. Valid value ranges: (2~300) sec. and the default is `5` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
     */
    healthCheckIntervalTime?: pulumi.Input<number>;
    /**
     * Indicates whether health check is enabled.
     */
    healthCheckSwitch?: pulumi.Input<boolean>;
    /**
     * Time out of health check. The value range is 2-60.
     */
    healthCheckTimeOut?: pulumi.Input<number>;
    /**
     * Type of health check. Valid value is `CUSTOM`, `TCP`, `HTTP`.
     */
    healthCheckType?: pulumi.Input<string>;
    /**
     * Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
     */
    healthCheckUnhealthNum?: pulumi.Input<number>;
    /**
     * Indicate to apply HTTP2.0 protocol or not.
     */
    http2Switch?: pulumi.Input<boolean>;
    /**
     * ID of CLB listener.
     */
    listenerId: pulumi.Input<string>;
    /**
     * Scheduling method of the CLB listener rules. Valid values: `WRR`, `IP HASH`, `LEAST_CONN`. The default is `WRR`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
     */
    scheduler?: pulumi.Input<string>;
    /**
     * Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as `WRR`, and not available when listener protocol is `TCP_SSL`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
     */
    sessionExpireTime?: pulumi.Input<number>;
    /**
     * Backend target type. Valid values: `NODE`, `TARGETGROUP`. `NODE` means to bind ordinary nodes, `TARGETGROUP` means to bind target group.
     */
    targetType?: pulumi.Input<string>;
    /**
     * Url of the listener rule.
     */
    url: pulumi.Input<string>;
}
