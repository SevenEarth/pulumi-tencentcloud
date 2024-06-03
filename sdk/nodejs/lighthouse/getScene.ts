// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of lighthouse scene with region
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const scene = tencentcloud.Lighthouse.getScene({
 *     limit: 20,
 *     offset: 0,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getScene(args?: GetSceneArgs, opts?: pulumi.InvokeOptions): Promise<GetSceneResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Lighthouse/getScene:getScene", {
        "limit": args.limit,
        "offset": args.offset,
        "resultOutputFile": args.resultOutputFile,
        "sceneIds": args.sceneIds,
    }, opts);
}

/**
 * A collection of arguments for invoking getScene.
 */
export interface GetSceneArgs {
    /**
     * Number of returned results. Default value is 20. Maximum value is 100.
     */
    limit?: number;
    /**
     * Offset. Default value is 0.
     */
    offset?: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * List of scene IDs.
     */
    sceneIds?: string[];
}

/**
 * A collection of values returned by getScene.
 */
export interface GetSceneResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly limit?: number;
    readonly offset?: number;
    readonly resultOutputFile?: string;
    readonly sceneIds?: string[];
    /**
     * List of scene info.
     */
    readonly sceneSets: outputs.Lighthouse.GetSceneSceneSet[];
}
/**
 * Use this data source to query detailed information of lighthouse scene with region
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const scene = tencentcloud.Lighthouse.getScene({
 *     limit: 20,
 *     offset: 0,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSceneOutput(args?: GetSceneOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSceneResult> {
    return pulumi.output(args).apply((a: any) => getScene(a, opts))
}

/**
 * A collection of arguments for invoking getScene.
 */
export interface GetSceneOutputArgs {
    /**
     * Number of returned results. Default value is 20. Maximum value is 100.
     */
    limit?: pulumi.Input<number>;
    /**
     * Offset. Default value is 0.
     */
    offset?: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * List of scene IDs.
     */
    sceneIds?: pulumi.Input<pulumi.Input<string>[]>;
}
