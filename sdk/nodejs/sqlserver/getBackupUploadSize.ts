// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of sqlserver datasourceBackupUploadSize
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Sqlserver.getBackupUploadSize({
 *     backupMigrationId: "mssql-backup-migration-9tj0sxnz",
 *     instanceId: "mssql-4gmc5805",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBackupUploadSize(args: GetBackupUploadSizeArgs, opts?: pulumi.InvokeOptions): Promise<GetBackupUploadSizeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Sqlserver/getBackupUploadSize:getBackupUploadSize", {
        "backupMigrationId": args.backupMigrationId,
        "incrementalMigrationId": args.incrementalMigrationId,
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackupUploadSize.
 */
export interface GetBackupUploadSizeArgs {
    /**
     * Backup import task ID, which is returned through the API CreateBackupMigration.
     */
    backupMigrationId: string;
    /**
     * Incremental import task ID.
     */
    incrementalMigrationId?: string;
    /**
     * ID of imported target instance.
     */
    instanceId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getBackupUploadSize.
 */
export interface GetBackupUploadSizeResult {
    readonly backupMigrationId: string;
    /**
     * Information of uploaded backups.
     */
    readonly cosUploadBackupFileSets: outputs.Sqlserver.GetBackupUploadSizeCosUploadBackupFileSet[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly incrementalMigrationId?: string;
    readonly instanceId: string;
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of sqlserver datasourceBackupUploadSize
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Sqlserver.getBackupUploadSize({
 *     backupMigrationId: "mssql-backup-migration-9tj0sxnz",
 *     instanceId: "mssql-4gmc5805",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBackupUploadSizeOutput(args: GetBackupUploadSizeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackupUploadSizeResult> {
    return pulumi.output(args).apply((a: any) => getBackupUploadSize(a, opts))
}

/**
 * A collection of arguments for invoking getBackupUploadSize.
 */
export interface GetBackupUploadSizeOutputArgs {
    /**
     * Backup import task ID, which is returned through the API CreateBackupMigration.
     */
    backupMigrationId: pulumi.Input<string>;
    /**
     * Incremental import task ID.
     */
    incrementalMigrationId?: pulumi.Input<string>;
    /**
     * ID of imported target instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
