// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query gaap layer7 listeners.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
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
 * const listenerId = tencentcloud.Gaap.getLayer7ListenersOutput({
 *     protocol: "HTTP",
 *     proxyId: fooProxy.id,
 *     listenerId: fooLayer7Listener.id,
 * });
 * ```
 */
export function getLayer7Listeners(args: GetLayer7ListenersArgs, opts?: pulumi.InvokeOptions): Promise<GetLayer7ListenersResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Gaap/getLayer7Listeners:getLayer7Listeners", {
        "listenerId": args.listenerId,
        "listenerName": args.listenerName,
        "port": args.port,
        "protocol": args.protocol,
        "proxyId": args.proxyId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getLayer7Listeners.
 */
export interface GetLayer7ListenersArgs {
    /**
     * ID of the layer7 listener to be queried.
     */
    listenerId?: string;
    /**
     * Name of the layer7 listener to be queried.
     */
    listenerName?: string;
    /**
     * Port of the layer7 listener to be queried.
     */
    port?: number;
    /**
     * Protocol of the layer7 listener to be queried. Valid values: `HTTP` and `HTTPS`.
     */
    protocol: string;
    /**
     * ID of the GAAP proxy to be queried.
     */
    proxyId?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getLayer7Listeners.
 */
export interface GetLayer7ListenersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly listenerId?: string;
    readonly listenerName?: string;
    /**
     * An information list of layer7 listeners. Each element contains the following attributes:
     */
    readonly listeners: outputs.Gaap.GetLayer7ListenersListener[];
    /**
     * Port of the layer7 listener.
     */
    readonly port?: number;
    /**
     * Protocol of the layer7 listener.
     */
    readonly protocol: string;
    readonly proxyId?: string;
    readonly resultOutputFile?: string;
}

export function getLayer7ListenersOutput(args: GetLayer7ListenersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLayer7ListenersResult> {
    return pulumi.output(args).apply(a => getLayer7Listeners(a, opts))
}

/**
 * A collection of arguments for invoking getLayer7Listeners.
 */
export interface GetLayer7ListenersOutputArgs {
    /**
     * ID of the layer7 listener to be queried.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * Name of the layer7 listener to be queried.
     */
    listenerName?: pulumi.Input<string>;
    /**
     * Port of the layer7 listener to be queried.
     */
    port?: pulumi.Input<number>;
    /**
     * Protocol of the layer7 listener to be queried. Valid values: `HTTP` and `HTTPS`.
     */
    protocol: pulumi.Input<string>;
    /**
     * ID of the GAAP proxy to be queried.
     */
    proxyId?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
