// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of postgresql backupDownloadUrls
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const logBackups = tencentcloud.Postgresql.getLogBackups({
 *     minFinishTime: "%s",
 *     maxFinishTime: "%s",
 *     filters: [{
 *         name: "db-instance-id",
 *         values: [local.pgsql_id],
 *     }],
 *     orderBy: "StartTime",
 *     orderByType: "desc",
 * });
 * const backupDownloadUrls = logBackups.then(logBackups => tencentcloud.Postgresql.getBackupDownloadUrls({
 *     dbInstanceId: local.pgsql_id,
 *     backupType: "LogBackup",
 *     backupId: logBackups.logBackupSets?.[0]?.id,
 *     urlExpireTime: 12,
 *     backupDownloadRestriction: {
 *         restrictionType: "NONE",
 *         vpcRestrictionEffect: "ALLOW",
 *         vpcIdSets: [local.vpc_id],
 *         ipRestrictionEffect: "ALLOW",
 *         ipSets: ["0.0.0.0"],
 *     },
 * }));
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBackupDownloadUrls(args: GetBackupDownloadUrlsArgs, opts?: pulumi.InvokeOptions): Promise<GetBackupDownloadUrlsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Postgresql/getBackupDownloadUrls:getBackupDownloadUrls", {
        "backupDownloadRestriction": args.backupDownloadRestriction,
        "backupId": args.backupId,
        "backupType": args.backupType,
        "dbInstanceId": args.dbInstanceId,
        "resultOutputFile": args.resultOutputFile,
        "urlExpireTime": args.urlExpireTime,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackupDownloadUrls.
 */
export interface GetBackupDownloadUrlsArgs {
    /**
     * Backup download restriction.
     */
    backupDownloadRestriction?: inputs.Postgresql.GetBackupDownloadUrlsBackupDownloadRestriction;
    /**
     * Unique backup ID.
     */
    backupId: string;
    /**
     * Backup type. Valid values: `LogBackup`, `BaseBackup`.
     */
    backupType: string;
    /**
     * Instance ID.
     */
    dbInstanceId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Validity period of a URL, which is 12 hours by default.
     */
    urlExpireTime?: number;
}

/**
 * A collection of values returned by getBackupDownloadUrls.
 */
export interface GetBackupDownloadUrlsResult {
    readonly backupDownloadRestriction?: outputs.Postgresql.GetBackupDownloadUrlsBackupDownloadRestriction;
    /**
     * Backup download URL.
     */
    readonly backupDownloadUrl: string;
    readonly backupId: string;
    readonly backupType: string;
    readonly dbInstanceId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly urlExpireTime?: number;
}
/**
 * Use this data source to query detailed information of postgresql backupDownloadUrls
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const logBackups = tencentcloud.Postgresql.getLogBackups({
 *     minFinishTime: "%s",
 *     maxFinishTime: "%s",
 *     filters: [{
 *         name: "db-instance-id",
 *         values: [local.pgsql_id],
 *     }],
 *     orderBy: "StartTime",
 *     orderByType: "desc",
 * });
 * const backupDownloadUrls = logBackups.then(logBackups => tencentcloud.Postgresql.getBackupDownloadUrls({
 *     dbInstanceId: local.pgsql_id,
 *     backupType: "LogBackup",
 *     backupId: logBackups.logBackupSets?.[0]?.id,
 *     urlExpireTime: 12,
 *     backupDownloadRestriction: {
 *         restrictionType: "NONE",
 *         vpcRestrictionEffect: "ALLOW",
 *         vpcIdSets: [local.vpc_id],
 *         ipRestrictionEffect: "ALLOW",
 *         ipSets: ["0.0.0.0"],
 *     },
 * }));
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBackupDownloadUrlsOutput(args: GetBackupDownloadUrlsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackupDownloadUrlsResult> {
    return pulumi.output(args).apply((a: any) => getBackupDownloadUrls(a, opts))
}

/**
 * A collection of arguments for invoking getBackupDownloadUrls.
 */
export interface GetBackupDownloadUrlsOutputArgs {
    /**
     * Backup download restriction.
     */
    backupDownloadRestriction?: pulumi.Input<inputs.Postgresql.GetBackupDownloadUrlsBackupDownloadRestrictionArgs>;
    /**
     * Unique backup ID.
     */
    backupId: pulumi.Input<string>;
    /**
     * Backup type. Valid values: `LogBackup`, `BaseBackup`.
     */
    backupType: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    dbInstanceId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Validity period of a URL, which is 12 hours by default.
     */
    urlExpireTime?: pulumi.Input<number>;
}
