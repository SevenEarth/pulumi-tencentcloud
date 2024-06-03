// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of rum whitelist
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const whitelist = tencentcloud.Rum.getWhitelist({
 *     instanceId: "rum-pasZKEI3RLgakj",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getWhitelist(args: GetWhitelistArgs, opts?: pulumi.InvokeOptions): Promise<GetWhitelistResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Rum/getWhitelist:getWhitelist", {
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getWhitelist.
 */
export interface GetWhitelistArgs {
    /**
     * Instance ID, such as taw-123.
     */
    instanceId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getWhitelist.
 */
export interface GetWhitelistResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly resultOutputFile?: string;
    /**
     * While list.
     */
    readonly whitelistSets: outputs.Rum.GetWhitelistWhitelistSet[];
}
/**
 * Use this data source to query detailed information of rum whitelist
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const whitelist = tencentcloud.Rum.getWhitelist({
 *     instanceId: "rum-pasZKEI3RLgakj",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getWhitelistOutput(args: GetWhitelistOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWhitelistResult> {
    return pulumi.output(args).apply((a: any) => getWhitelist(a, opts))
}

/**
 * A collection of arguments for invoking getWhitelist.
 */
export interface GetWhitelistOutputArgs {
    /**
     * Instance ID, such as taw-123.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
