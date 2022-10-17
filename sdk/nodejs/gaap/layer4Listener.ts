// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a layer4 listener of GAAP.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const fooProxy = new tencentcloud.gaap.Proxy("fooProxy", {
 *     bandwidth: 10,
 *     concurrent: 2,
 *     accessRegion: "SouthChina",
 *     realserverRegion: "NorthChina",
 * });
 * const fooRealserver = new tencentcloud.gaap.Realserver("fooRealserver", {ip: "1.1.1.1"});
 * const bar = new tencentcloud.gaap.Realserver("bar", {ip: "119.29.29.29"});
 * const fooLayer4Listener = new tencentcloud.gaap.Layer4Listener("fooLayer4Listener", {
 *     protocol: "TCP",
 *     port: 80,
 *     realserverType: "IP",
 *     proxyId: fooProxy.id,
 *     healthCheck: true,
 *     realserverBindSets: [
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
 * GAAP layer4 listener can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Gaap/layer4Listener:Layer4Listener tencentcloud_gaap_layer4_listener.foo listener-11112222
 * ```
 */
export class Layer4Listener extends pulumi.CustomResource {
    /**
     * Get an existing Layer4Listener resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Layer4ListenerState, opts?: pulumi.CustomResourceOptions): Layer4Listener {
        return new Layer4Listener(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Gaap/layer4Listener:Layer4Listener';

    /**
     * Returns true if the given object is an instance of Layer4Listener.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Layer4Listener {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Layer4Listener.__pulumiType;
    }

    /**
     * The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports listeners of `TCP` protocol.
     */
    public readonly clientIpMethod!: pulumi.Output<number | undefined>;
    /**
     * Timeout of the health check response, should less than interval, default value is 2s. NOTES: Only supports listeners of `TCP` protocol and require less than `interval`.
     */
    public readonly connectTimeout!: pulumi.Output<number | undefined>;
    /**
     * Creation time of the layer4 listener.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Indicates whether health check is enable, default value is `false`. NOTES: Only supports listeners of `TCP` protocol.
     */
    public readonly healthCheck!: pulumi.Output<boolean | undefined>;
    /**
     * Interval of the health check, default value is 5s. NOTES: Only supports listeners of `TCP` protocol.
     */
    public readonly interval!: pulumi.Output<number | undefined>;
    /**
     * Name of the layer4 listener, the maximum length is 30.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Port of the layer4 listener.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Protocol of the layer4 listener. Valid value: `TCP` and `UDP`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * ID of the GAAP proxy.
     */
    public readonly proxyId!: pulumi.Output<string>;
    /**
     * An information list of GAAP realserver.
     */
    public readonly realserverBindSets!: pulumi.Output<outputs.Gaap.Layer4ListenerRealserverBindSet[] | undefined>;
    /**
     * Type of the realserver. Valid value: `IP` and `DOMAIN`. NOTES: when the `protocol` is specified as `TCP` and the `scheduler` is specified as `wrr`, the item can only be set to `IP`.
     */
    public readonly realserverType!: pulumi.Output<string>;
    /**
     * Scheduling policy of the layer4 listener, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
     */
    public readonly scheduler!: pulumi.Output<string | undefined>;
    /**
     * Status of the layer4 listener.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;

    /**
     * Create a Layer4Listener resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Layer4ListenerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Layer4ListenerArgs | Layer4ListenerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Layer4ListenerState | undefined;
            resourceInputs["clientIpMethod"] = state ? state.clientIpMethod : undefined;
            resourceInputs["connectTimeout"] = state ? state.connectTimeout : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["healthCheck"] = state ? state.healthCheck : undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["proxyId"] = state ? state.proxyId : undefined;
            resourceInputs["realserverBindSets"] = state ? state.realserverBindSets : undefined;
            resourceInputs["realserverType"] = state ? state.realserverType : undefined;
            resourceInputs["scheduler"] = state ? state.scheduler : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as Layer4ListenerArgs | undefined;
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.proxyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'proxyId'");
            }
            if ((!args || args.realserverType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realserverType'");
            }
            resourceInputs["clientIpMethod"] = args ? args.clientIpMethod : undefined;
            resourceInputs["connectTimeout"] = args ? args.connectTimeout : undefined;
            resourceInputs["healthCheck"] = args ? args.healthCheck : undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["proxyId"] = args ? args.proxyId : undefined;
            resourceInputs["realserverBindSets"] = args ? args.realserverBindSets : undefined;
            resourceInputs["realserverType"] = args ? args.realserverType : undefined;
            resourceInputs["scheduler"] = args ? args.scheduler : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Layer4Listener.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Layer4Listener resources.
 */
export interface Layer4ListenerState {
    /**
     * The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports listeners of `TCP` protocol.
     */
    clientIpMethod?: pulumi.Input<number>;
    /**
     * Timeout of the health check response, should less than interval, default value is 2s. NOTES: Only supports listeners of `TCP` protocol and require less than `interval`.
     */
    connectTimeout?: pulumi.Input<number>;
    /**
     * Creation time of the layer4 listener.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Indicates whether health check is enable, default value is `false`. NOTES: Only supports listeners of `TCP` protocol.
     */
    healthCheck?: pulumi.Input<boolean>;
    /**
     * Interval of the health check, default value is 5s. NOTES: Only supports listeners of `TCP` protocol.
     */
    interval?: pulumi.Input<number>;
    /**
     * Name of the layer4 listener, the maximum length is 30.
     */
    name?: pulumi.Input<string>;
    /**
     * Port of the layer4 listener.
     */
    port?: pulumi.Input<number>;
    /**
     * Protocol of the layer4 listener. Valid value: `TCP` and `UDP`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * ID of the GAAP proxy.
     */
    proxyId?: pulumi.Input<string>;
    /**
     * An information list of GAAP realserver.
     */
    realserverBindSets?: pulumi.Input<pulumi.Input<inputs.Gaap.Layer4ListenerRealserverBindSet>[]>;
    /**
     * Type of the realserver. Valid value: `IP` and `DOMAIN`. NOTES: when the `protocol` is specified as `TCP` and the `scheduler` is specified as `wrr`, the item can only be set to `IP`.
     */
    realserverType?: pulumi.Input<string>;
    /**
     * Scheduling policy of the layer4 listener, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
     */
    scheduler?: pulumi.Input<string>;
    /**
     * Status of the layer4 listener.
     */
    status?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Layer4Listener resource.
 */
export interface Layer4ListenerArgs {
    /**
     * The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports listeners of `TCP` protocol.
     */
    clientIpMethod?: pulumi.Input<number>;
    /**
     * Timeout of the health check response, should less than interval, default value is 2s. NOTES: Only supports listeners of `TCP` protocol and require less than `interval`.
     */
    connectTimeout?: pulumi.Input<number>;
    /**
     * Indicates whether health check is enable, default value is `false`. NOTES: Only supports listeners of `TCP` protocol.
     */
    healthCheck?: pulumi.Input<boolean>;
    /**
     * Interval of the health check, default value is 5s. NOTES: Only supports listeners of `TCP` protocol.
     */
    interval?: pulumi.Input<number>;
    /**
     * Name of the layer4 listener, the maximum length is 30.
     */
    name?: pulumi.Input<string>;
    /**
     * Port of the layer4 listener.
     */
    port: pulumi.Input<number>;
    /**
     * Protocol of the layer4 listener. Valid value: `TCP` and `UDP`.
     */
    protocol: pulumi.Input<string>;
    /**
     * ID of the GAAP proxy.
     */
    proxyId: pulumi.Input<string>;
    /**
     * An information list of GAAP realserver.
     */
    realserverBindSets?: pulumi.Input<pulumi.Input<inputs.Gaap.Layer4ListenerRealserverBindSet>[]>;
    /**
     * Type of the realserver. Valid value: `IP` and `DOMAIN`. NOTES: when the `protocol` is specified as `TCP` and the `scheduler` is specified as `wrr`, the item can only be set to `IP`.
     */
    realserverType: pulumi.Input<string>;
    /**
     * Scheduling policy of the layer4 listener, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
     */
    scheduler?: pulumi.Input<string>;
}
