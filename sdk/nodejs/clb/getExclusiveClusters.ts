// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of clb exclusiveClusters
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const exclusiveClusters = pulumi.output(tencentcloud.Clb.getExclusiveClusters({
 *     filters: [{
 *         name: "zone",
 *         values: ["ap-guangzhou-1"],
 *     }],
 * }));
 * ```
 */
export function getExclusiveClusters(args?: GetExclusiveClustersArgs, opts?: pulumi.InvokeOptions): Promise<GetExclusiveClustersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Clb/getExclusiveClusters:getExclusiveClusters", {
        "filters": args.filters,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getExclusiveClusters.
 */
export interface GetExclusiveClustersArgs {
    /**
     * Filter to query the list of AZ resources as detailed below: cluster-type - String - Required: No - (Filter condition) Filter by cluster type, such as TGW. cluster-id - String - Required: No - (Filter condition) Filter by cluster ID, such as tgw-xxxxxxxx. cluster-name - String - Required: No - (Filter condition) Filter by cluster name, such as test-xxxxxx. cluster-tag - String - Required: No - (Filter condition) Filter by cluster tag, such as TAG-xxxxx. vip - String - Required: No - (Filter condition) Filter by vip in the cluster, such as x.x.x.x. network - String - Required: No - (Filter condition) Filter by cluster network type, such as Public or Private. zone - String - Required: No - (Filter condition) Filter by cluster zone, such as ap-guangzhou-1. isp - String - Required: No - (Filter condition) Filter by TGW cluster isp type, such as BGP. loadblancer-id - String - Required: No - (Filter condition) Filter by loadblancer-id in the cluste, such as lb-xxxxxxxx.
     */
    filters?: inputs.Clb.GetExclusiveClustersFilter[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getExclusiveClusters.
 */
export interface GetExclusiveClustersResult {
    /**
     * cluster list.
     */
    readonly clusterSets: outputs.Clb.GetExclusiveClustersClusterSet[];
    readonly filters?: outputs.Clb.GetExclusiveClustersFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
}

export function getExclusiveClustersOutput(args?: GetExclusiveClustersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetExclusiveClustersResult> {
    return pulumi.output(args).apply(a => getExclusiveClusters(a, opts))
}

/**
 * A collection of arguments for invoking getExclusiveClusters.
 */
export interface GetExclusiveClustersOutputArgs {
    /**
     * Filter to query the list of AZ resources as detailed below: cluster-type - String - Required: No - (Filter condition) Filter by cluster type, such as TGW. cluster-id - String - Required: No - (Filter condition) Filter by cluster ID, such as tgw-xxxxxxxx. cluster-name - String - Required: No - (Filter condition) Filter by cluster name, such as test-xxxxxx. cluster-tag - String - Required: No - (Filter condition) Filter by cluster tag, such as TAG-xxxxx. vip - String - Required: No - (Filter condition) Filter by vip in the cluster, such as x.x.x.x. network - String - Required: No - (Filter condition) Filter by cluster network type, such as Public or Private. zone - String - Required: No - (Filter condition) Filter by cluster zone, such as ap-guangzhou-1. isp - String - Required: No - (Filter condition) Filter by TGW cluster isp type, such as BGP. loadblancer-id - String - Required: No - (Filter condition) Filter by loadblancer-id in the cluste, such as lb-xxxxxxxx.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.Clb.GetExclusiveClustersFilterArgs>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
