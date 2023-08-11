// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of sqlserver datasourceBackupCommand
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = pulumi.output(tencentcloud.Sqlserver.getBackupCommands({
 *     backupFileType: "FULL",
 *     dataBaseName: "keep-publish-instance",
 *     isRecovery: "NO",
 * }));
 * ```
 */
export function getBackupCommands(args: GetBackupCommandsArgs, opts?: pulumi.InvokeOptions): Promise<GetBackupCommandsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Sqlserver/getBackupCommands:getBackupCommands", {
        "backupFileType": args.backupFileType,
        "dataBaseName": args.dataBaseName,
        "isRecovery": args.isRecovery,
        "localPath": args.localPath,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackupCommands.
 */
export interface GetBackupCommandsArgs {
    /**
     * Backup file type. Full: full backup. FULL_LOG: full backup which needs log increments. FULL_DIFF: full backup which needs differential increments. LOG: log backup. DIFF: differential backup.
     */
    backupFileType: string;
    /**
     * Database name.
     */
    dataBaseName: string;
    /**
     * Whether restoration is required. No: not required. Yes: required.
     */
    isRecovery: string;
    /**
     * Storage path of backup files. If this parameter is left empty, the default storage path will be D:.
     */
    localPath?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getBackupCommands.
 */
export interface GetBackupCommandsResult {
    readonly backupFileType: string;
    readonly dataBaseName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isRecovery: string;
    /**
     * Command list.
     */
    readonly lists: outputs.Sqlserver.GetBackupCommandsList[];
    readonly localPath?: string;
    readonly resultOutputFile?: string;
}

export function getBackupCommandsOutput(args: GetBackupCommandsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackupCommandsResult> {
    return pulumi.output(args).apply(a => getBackupCommands(a, opts))
}

/**
 * A collection of arguments for invoking getBackupCommands.
 */
export interface GetBackupCommandsOutputArgs {
    /**
     * Backup file type. Full: full backup. FULL_LOG: full backup which needs log increments. FULL_DIFF: full backup which needs differential increments. LOG: log backup. DIFF: differential backup.
     */
    backupFileType: pulumi.Input<string>;
    /**
     * Database name.
     */
    dataBaseName: pulumi.Input<string>;
    /**
     * Whether restoration is required. No: not required. Yes: required.
     */
    isRecovery: pulumi.Input<string>;
    /**
     * Storage path of backup files. If this parameter is left empty, the default storage path will be D:.
     */
    localPath?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
