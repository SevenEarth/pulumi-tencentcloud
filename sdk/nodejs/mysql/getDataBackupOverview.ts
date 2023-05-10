// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of mysql dataBackupOverview
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const dataBackupOverview = pulumi.output(tencentcloud.Mysql.getDataBackupOverview({
 *     product: "mysql",
 * }));
 * ```
 */
export function getDataBackupOverview(args: GetDataBackupOverviewArgs, opts?: pulumi.InvokeOptions): Promise<GetDataBackupOverviewResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Mysql/getDataBackupOverview:getDataBackupOverview", {
        "product": args.product,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDataBackupOverview.
 */
export interface GetDataBackupOverviewArgs {
    /**
     * The type of cloud database product to be queried, currently only supports `mysql`.
     */
    product: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDataBackupOverview.
 */
export interface GetDataBackupOverviewResult {
    /**
     * The total number of automatic backups in the current region.
     */
    readonly autoBackupCount: number;
    /**
     * The total automatic backup capacity of the current region.
     */
    readonly autoBackupVolume: number;
    /**
     * The total number of archive backups in the current region.
     */
    readonly dataBackupArchiveCount: number;
    /**
     * The total capacity of the current regional archive backup.
     */
    readonly dataBackupArchiveVolume: number;
    /**
     * The total number of data backups in the current region.
     */
    readonly dataBackupCount: number;
    /**
     * The total number of standard storage backups in the current region.
     */
    readonly dataBackupStandbyCount: number;
    /**
     * The total backup capacity of the current regional standard storage.
     */
    readonly dataBackupStandbyVolume: number;
    /**
     * Total data backup capacity of the current region (including automatic backup and manual backup, in bytes).
     */
    readonly dataBackupVolume: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The total number of manual backups in the current region.
     */
    readonly manualBackupCount: number;
    /**
     * The total manual backup capacity of the current region.
     */
    readonly manualBackupVolume: number;
    readonly product: string;
    /**
     * The total number of remote backups.
     */
    readonly remoteBackupCount: number;
    /**
     * The total capacity of remote backup.
     */
    readonly remoteBackupVolume: number;
    readonly resultOutputFile?: string;
}

export function getDataBackupOverviewOutput(args: GetDataBackupOverviewOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataBackupOverviewResult> {
    return pulumi.output(args).apply(a => getDataBackupOverview(a, opts))
}

/**
 * A collection of arguments for invoking getDataBackupOverview.
 */
export interface GetDataBackupOverviewOutputArgs {
    /**
     * The type of cloud database product to be queried, currently only supports `mysql`.
     */
    product: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
