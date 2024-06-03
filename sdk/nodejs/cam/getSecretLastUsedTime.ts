// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cam secretLastUsedTime
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const secretLastUsedTime = tencentcloud.Cam.getSecretLastUsedTime({
 *     secretIdLists: ["xxxx"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSecretLastUsedTime(args: GetSecretLastUsedTimeArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretLastUsedTimeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cam/getSecretLastUsedTime:getSecretLastUsedTime", {
        "resultOutputFile": args.resultOutputFile,
        "secretIdLists": args.secretIdLists,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecretLastUsedTime.
 */
export interface GetSecretLastUsedTimeArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Query the key ID list. Supports up to 10.
     */
    secretIdLists: string[];
}

/**
 * A collection of values returned by getSecretLastUsedTime.
 */
export interface GetSecretLastUsedTimeResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * Last used time list.
     */
    readonly secretIdLastUsedRows: outputs.Cam.GetSecretLastUsedTimeSecretIdLastUsedRow[];
    readonly secretIdLists: string[];
}
/**
 * Use this data source to query detailed information of cam secretLastUsedTime
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const secretLastUsedTime = tencentcloud.Cam.getSecretLastUsedTime({
 *     secretIdLists: ["xxxx"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSecretLastUsedTimeOutput(args: GetSecretLastUsedTimeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecretLastUsedTimeResult> {
    return pulumi.output(args).apply((a: any) => getSecretLastUsedTime(a, opts))
}

/**
 * A collection of arguments for invoking getSecretLastUsedTime.
 */
export interface GetSecretLastUsedTimeOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Query the key ID list. Supports up to 10.
     */
    secretIdLists: pulumi.Input<pulumi.Input<string>[]>;
}
