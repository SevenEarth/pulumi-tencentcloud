// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of redis backup
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const backup = tencentcloud.Redis.getBackup({
 *     beginTime: "2023-04-07 03:57:30",
 *     endTime: "2023-04-07 03:57:56",
 *     instanceId: "crs-c1nl9rpv",
 *     instanceName: "Keep-terraform",
 *     statuses: [2],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBackup(args?: GetBackupArgs, opts?: pulumi.InvokeOptions): Promise<GetBackupResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Redis/getBackup:getBackup", {
        "beginTime": args.beginTime,
        "endTime": args.endTime,
        "instanceId": args.instanceId,
        "instanceName": args.instanceName,
        "resultOutputFile": args.resultOutputFile,
        "statuses": args.statuses,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackup.
 */
export interface GetBackupArgs {
    /**
     * start time, such as 2017-02-08 19:09:26.Query the list of backups that the instance started backing up during the [beginTime, endTime] time period.
     */
    beginTime?: string;
    /**
     * End time, such as 2017-02-08 19:09:26.Query the list of backups that the instance started backing up during the [beginTime, endTime] time period.
     */
    endTime?: string;
    /**
     * The ID of instance.
     */
    instanceId?: string;
    /**
     * Instance name, which supports fuzzy search based on instance name.
     */
    instanceName?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Status of the backup task:1: Backup is in the process.2: The backup is normal.3: Backup to RDB file processing.4: RDB conversion completed.-1: The backup has expired.-2: Backup deleted.
     */
    statuses?: number[];
}

/**
 * A collection of values returned by getBackup.
 */
export interface GetBackupResult {
    /**
     * An array of backups for the instance.
     */
    readonly backupSets: outputs.Redis.GetBackupBackupSet[];
    readonly beginTime?: string;
    /**
     * Backup end time.
     */
    readonly endTime?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ID of instance.
     */
    readonly instanceId?: string;
    /**
     * The name of instance.
     */
    readonly instanceName?: string;
    readonly resultOutputFile?: string;
    /**
     * Backup status.1: The backup is locked by another process.2: The backup is normal and not locked by any process.-1: The backup has expired.3: The backup is being exported.4: The backup export is successful.
     */
    readonly statuses?: number[];
}
/**
 * Use this data source to query detailed information of redis backup
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const backup = tencentcloud.Redis.getBackup({
 *     beginTime: "2023-04-07 03:57:30",
 *     endTime: "2023-04-07 03:57:56",
 *     instanceId: "crs-c1nl9rpv",
 *     instanceName: "Keep-terraform",
 *     statuses: [2],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBackupOutput(args?: GetBackupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackupResult> {
    return pulumi.output(args).apply((a: any) => getBackup(a, opts))
}

/**
 * A collection of arguments for invoking getBackup.
 */
export interface GetBackupOutputArgs {
    /**
     * start time, such as 2017-02-08 19:09:26.Query the list of backups that the instance started backing up during the [beginTime, endTime] time period.
     */
    beginTime?: pulumi.Input<string>;
    /**
     * End time, such as 2017-02-08 19:09:26.Query the list of backups that the instance started backing up during the [beginTime, endTime] time period.
     */
    endTime?: pulumi.Input<string>;
    /**
     * The ID of instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Instance name, which supports fuzzy search based on instance name.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Status of the backup task:1: Backup is in the process.2: The backup is normal.3: Backup to RDB file processing.4: RDB conversion completed.-1: The backup has expired.-2: Backup deleted.
     */
    statuses?: pulumi.Input<pulumi.Input<number>[]>;
}
