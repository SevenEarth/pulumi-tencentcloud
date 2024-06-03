// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tcr replicationInstanceSyncStatus
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const syncStatus = tencentcloud.Tcr.getReplicationInstanceSyncStatus({
 *     registryId: local.src_registry_id,
 *     replicationRegistryId: local.dst_registry_id,
 *     replicationRegionId: local.dst_region_id,
 *     showReplicationLog: false,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getReplicationInstanceSyncStatus(args: GetReplicationInstanceSyncStatusArgs, opts?: pulumi.InvokeOptions): Promise<GetReplicationInstanceSyncStatusResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Tcr/getReplicationInstanceSyncStatus:getReplicationInstanceSyncStatus", {
        "registryId": args.registryId,
        "replicationRegionId": args.replicationRegionId,
        "replicationRegistryId": args.replicationRegistryId,
        "resultOutputFile": args.resultOutputFile,
        "showReplicationLog": args.showReplicationLog,
    }, opts);
}

/**
 * A collection of arguments for invoking getReplicationInstanceSyncStatus.
 */
export interface GetReplicationInstanceSyncStatusArgs {
    /**
     * master registry id.
     */
    registryId: string;
    /**
     * synchronization instance region id.
     */
    replicationRegionId?: number;
    /**
     * synchronization instance id.
     */
    replicationRegistryId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * whether to display the synchronization log.
     */
    showReplicationLog?: boolean;
}

/**
 * A collection of values returned by getReplicationInstanceSyncStatus.
 */
export interface GetReplicationInstanceSyncStatusResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly registryId: string;
    /**
     * sync log. Note: This field may return null, indicating that no valid value can be obtained.
     */
    readonly replicationLogs: outputs.Tcr.GetReplicationInstanceSyncStatusReplicationLog[];
    readonly replicationRegionId?: number;
    readonly replicationRegistryId: string;
    /**
     * sync status.
     */
    readonly replicationStatus: string;
    /**
     * sync complete time.
     */
    readonly replicationTime: string;
    readonly resultOutputFile?: string;
    readonly showReplicationLog?: boolean;
}
/**
 * Use this data source to query detailed information of tcr replicationInstanceSyncStatus
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const syncStatus = tencentcloud.Tcr.getReplicationInstanceSyncStatus({
 *     registryId: local.src_registry_id,
 *     replicationRegistryId: local.dst_registry_id,
 *     replicationRegionId: local.dst_region_id,
 *     showReplicationLog: false,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getReplicationInstanceSyncStatusOutput(args: GetReplicationInstanceSyncStatusOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReplicationInstanceSyncStatusResult> {
    return pulumi.output(args).apply((a: any) => getReplicationInstanceSyncStatus(a, opts))
}

/**
 * A collection of arguments for invoking getReplicationInstanceSyncStatus.
 */
export interface GetReplicationInstanceSyncStatusOutputArgs {
    /**
     * master registry id.
     */
    registryId: pulumi.Input<string>;
    /**
     * synchronization instance region id.
     */
    replicationRegionId?: pulumi.Input<number>;
    /**
     * synchronization instance id.
     */
    replicationRegistryId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * whether to display the synchronization log.
     */
    showReplicationLog?: pulumi.Input<boolean>;
}
