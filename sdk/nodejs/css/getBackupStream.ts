// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of css backupStream
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const backupStream = tencentcloud.Css.getBackupStream({
 *     streamName: "live",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBackupStream(args?: GetBackupStreamArgs, opts?: pulumi.InvokeOptions): Promise<GetBackupStreamResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Css/getBackupStream:getBackupStream", {
        "resultOutputFile": args.resultOutputFile,
        "streamName": args.streamName,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackupStream.
 */
export interface GetBackupStreamArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Stream id.
     */
    streamName?: string;
}

/**
 * A collection of values returned by getBackupStream.
 */
export interface GetBackupStreamResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * Backup stream group info.
     */
    readonly streamInfoLists: outputs.Css.GetBackupStreamStreamInfoList[];
    /**
     * Stream name.
     */
    readonly streamName?: string;
}
/**
 * Use this data source to query detailed information of css backupStream
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const backupStream = tencentcloud.Css.getBackupStream({
 *     streamName: "live",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBackupStreamOutput(args?: GetBackupStreamOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackupStreamResult> {
    return pulumi.output(args).apply((a: any) => getBackupStream(a, opts))
}

/**
 * A collection of arguments for invoking getBackupStream.
 */
export interface GetBackupStreamOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Stream id.
     */
    streamName?: pulumi.Input<string>;
}
