// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of lighthouse resetInstanceBlueprint
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const resetInstanceBlueprint = tencentcloud.Lighthouse.getResetInstanceBlueprint({
 *     instanceId: "lhins-123456",
 *     limit: 20,
 *     offset: 0,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getResetInstanceBlueprint(args: GetResetInstanceBlueprintArgs, opts?: pulumi.InvokeOptions): Promise<GetResetInstanceBlueprintResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Lighthouse/getResetInstanceBlueprint:getResetInstanceBlueprint", {
        "filters": args.filters,
        "instanceId": args.instanceId,
        "limit": args.limit,
        "offset": args.offset,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getResetInstanceBlueprint.
 */
export interface GetResetInstanceBlueprintArgs {
    /**
     * Filter listblueprint-idFilter by image ID.Type: StringRequired: noblueprint-typeFilter by image type.Valid values: APP_OS: application image; PURE_OS: system image; PRIVATE: custom imageType: StringRequired: noplatform-typeFilter by image platform type.Valid values: LINUX_UNIX: Linux or Unix; WINDOWS: WindowsType: StringRequired: noblueprint-nameFilter by image name.Type: StringRequired: noblueprint-stateFilter by image status.Type: StringRequired: noEach request can contain up to 10 Filters and 5 Filter.Values. BlueprintIds and Filters cannot be specified at the same time.
     */
    filters?: inputs.Lighthouse.GetResetInstanceBlueprintFilter[];
    /**
     * Instance ID.
     */
    instanceId: string;
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
}

/**
 * A collection of values returned by getResetInstanceBlueprint.
 */
export interface GetResetInstanceBlueprintResult {
    readonly filters?: outputs.Lighthouse.GetResetInstanceBlueprintFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly limit?: number;
    readonly offset?: number;
    /**
     * List of scene info.
     */
    readonly resetInstanceBlueprintSets: outputs.Lighthouse.GetResetInstanceBlueprintResetInstanceBlueprintSet[];
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of lighthouse resetInstanceBlueprint
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const resetInstanceBlueprint = tencentcloud.Lighthouse.getResetInstanceBlueprint({
 *     instanceId: "lhins-123456",
 *     limit: 20,
 *     offset: 0,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getResetInstanceBlueprintOutput(args: GetResetInstanceBlueprintOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResetInstanceBlueprintResult> {
    return pulumi.output(args).apply((a: any) => getResetInstanceBlueprint(a, opts))
}

/**
 * A collection of arguments for invoking getResetInstanceBlueprint.
 */
export interface GetResetInstanceBlueprintOutputArgs {
    /**
     * Filter listblueprint-idFilter by image ID.Type: StringRequired: noblueprint-typeFilter by image type.Valid values: APP_OS: application image; PURE_OS: system image; PRIVATE: custom imageType: StringRequired: noplatform-typeFilter by image platform type.Valid values: LINUX_UNIX: Linux or Unix; WINDOWS: WindowsType: StringRequired: noblueprint-nameFilter by image name.Type: StringRequired: noblueprint-stateFilter by image status.Type: StringRequired: noEach request can contain up to 10 Filters and 5 Filter.Values. BlueprintIds and Filters cannot be specified at the same time.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.Lighthouse.GetResetInstanceBlueprintFilterArgs>[]>;
    /**
     * Instance ID.
     */
    instanceId: pulumi.Input<string>;
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
}
