// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of kubernetes availableClusterVersions
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const queryById = tencentcloud.Kubernetes.getAvailableClusterVersions({
 *     clusterId: "xxx",
 * });
 * export const versionsId = queryById.then(queryById => queryById.versions);
 * const queryByIds = tencentcloud.Kubernetes.getAvailableClusterVersions({
 *     clusterIds: ["xxx"],
 * });
 * export const versionsIds = queryByIds.then(queryByIds => queryByIds.clusters);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getAvailableClusterVersions(args?: GetAvailableClusterVersionsArgs, opts?: pulumi.InvokeOptions): Promise<GetAvailableClusterVersionsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Kubernetes/getAvailableClusterVersions:getAvailableClusterVersions", {
        "clusterId": args.clusterId,
        "clusterIds": args.clusterIds,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getAvailableClusterVersions.
 */
export interface GetAvailableClusterVersionsArgs {
    /**
     * Cluster Id.
     */
    clusterId?: string;
    /**
     * list of cluster IDs.
     */
    clusterIds?: string[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getAvailableClusterVersions.
 */
export interface GetAvailableClusterVersionsResult {
    /**
     * Cluster ID.
     */
    readonly clusterId?: string;
    readonly clusterIds?: string[];
    /**
     * cluster information. Note: This field may return null, indicating that no valid value can be obtained.
     */
    readonly clusters: outputs.Kubernetes.GetAvailableClusterVersionsCluster[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * Upgradable cluster version number. Note: This field may return null, indicating that no valid value can be obtained.
     */
    readonly versions: string[];
}
/**
 * Use this data source to query detailed information of kubernetes availableClusterVersions
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const queryById = tencentcloud.Kubernetes.getAvailableClusterVersions({
 *     clusterId: "xxx",
 * });
 * export const versionsId = queryById.then(queryById => queryById.versions);
 * const queryByIds = tencentcloud.Kubernetes.getAvailableClusterVersions({
 *     clusterIds: ["xxx"],
 * });
 * export const versionsIds = queryByIds.then(queryByIds => queryByIds.clusters);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getAvailableClusterVersionsOutput(args?: GetAvailableClusterVersionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAvailableClusterVersionsResult> {
    return pulumi.output(args).apply((a: any) => getAvailableClusterVersions(a, opts))
}

/**
 * A collection of arguments for invoking getAvailableClusterVersions.
 */
export interface GetAvailableClusterVersionsOutputArgs {
    /**
     * Cluster Id.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * list of cluster IDs.
     */
    clusterIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
