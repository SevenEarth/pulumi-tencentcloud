// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of gaap proxies status
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const proxiesStatus = pulumi.output(tencentcloud.Gaap.getProxiesStatus({
 *     proxyIds: ["link-xxxxxx"],
 * }));
 * ```
 */
export function getProxiesStatus(args?: GetProxiesStatusArgs, opts?: pulumi.InvokeOptions): Promise<GetProxiesStatusResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Gaap/getProxiesStatus:getProxiesStatus", {
        "proxyIds": args.proxyIds,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getProxiesStatus.
 */
export interface GetProxiesStatusArgs {
    /**
     * List of Proxy IDs.
     */
    proxyIds?: string[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getProxiesStatus.
 */
export interface GetProxiesStatusResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Proxy status list.
     */
    readonly instanceStatusSets: outputs.Gaap.GetProxiesStatusInstanceStatusSet[];
    readonly proxyIds?: string[];
    readonly resultOutputFile?: string;
}

export function getProxiesStatusOutput(args?: GetProxiesStatusOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProxiesStatusResult> {
    return pulumi.output(args).apply(a => getProxiesStatus(a, opts))
}

/**
 * A collection of arguments for invoking getProxiesStatus.
 */
export interface GetProxiesStatusOutputArgs {
    /**
     * List of Proxy IDs.
     */
    proxyIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
