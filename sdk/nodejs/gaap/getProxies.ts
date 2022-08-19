// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query gaap proxies.
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
 * const fooProxies = tencentcloud.Gaap.getProxiesOutput({
 *     ids: [fooProxy.id],
 * });
 * ```
 */
export function getProxies(args?: GetProxiesArgs, opts?: pulumi.InvokeOptions): Promise<GetProxiesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Gaap/getProxies:getProxies", {
        "accessRegion": args.accessRegion,
        "ids": args.ids,
        "projectId": args.projectId,
        "realserverRegion": args.realserverRegion,
        "resultOutputFile": args.resultOutputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getProxies.
 */
export interface GetProxiesArgs {
    /**
     * Access region of the GAAP proxy to be queried. Conflict with `ids`.
     */
    accessRegion?: string;
    /**
     * ID of the GAAP proxy to be queried. Conflict with `projectId`, `accessRegion` and `realserverRegion`.
     */
    ids?: string[];
    /**
     * Project ID of the GAAP proxy to be queried. Conflict with `ids`.
     */
    projectId?: number;
    /**
     * Region of the GAAP realserver to be queried. Conflict with `ids`.
     */
    realserverRegion?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Tags of the GAAP proxy to be queried. Support up to 5, display the information as long as it matches one.
     */
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getProxies.
 */
export interface GetProxiesResult {
    /**
     * Access region of the GAAP proxy.
     */
    readonly accessRegion?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * ID of the project within the GAAP proxy, '0' means is default project.
     */
    readonly projectId?: number;
    /**
     * An information list of GAAP proxy. Each element contains the following attributes:
     */
    readonly proxies: outputs.Gaap.GetProxiesProxy[];
    /**
     * Region of the GAAP realserver.
     */
    readonly realserverRegion?: string;
    readonly resultOutputFile?: string;
    /**
     * Tags of the GAAP proxy.
     */
    readonly tags?: {[key: string]: any};
}

export function getProxiesOutput(args?: GetProxiesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProxiesResult> {
    return pulumi.output(args).apply(a => getProxies(a, opts))
}

/**
 * A collection of arguments for invoking getProxies.
 */
export interface GetProxiesOutputArgs {
    /**
     * Access region of the GAAP proxy to be queried. Conflict with `ids`.
     */
    accessRegion?: pulumi.Input<string>;
    /**
     * ID of the GAAP proxy to be queried. Conflict with `projectId`, `accessRegion` and `realserverRegion`.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project ID of the GAAP proxy to be queried. Conflict with `ids`.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Region of the GAAP realserver to be queried. Conflict with `ids`.
     */
    realserverRegion?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Tags of the GAAP proxy to be queried. Support up to 5, display the information as long as it matches one.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
