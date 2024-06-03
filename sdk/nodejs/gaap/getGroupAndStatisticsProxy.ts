// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of gaap and statistics proxy
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const groupAndStatisticsProxy = tencentcloud.Gaap.getGroupAndStatisticsProxy({
 *     projectId: 0,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getGroupAndStatisticsProxy(args: GetGroupAndStatisticsProxyArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupAndStatisticsProxyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Gaap/getGroupAndStatisticsProxy:getGroupAndStatisticsProxy", {
        "projectId": args.projectId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupAndStatisticsProxy.
 */
export interface GetGroupAndStatisticsProxyArgs {
    /**
     * Project Id.
     */
    projectId: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getGroupAndStatisticsProxy.
 */
export interface GetGroupAndStatisticsProxyResult {
    /**
     * Channel group information that can be counted.
     */
    readonly groupSets: outputs.Gaap.GetGroupAndStatisticsProxyGroupSet[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly projectId: number;
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of gaap and statistics proxy
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const groupAndStatisticsProxy = tencentcloud.Gaap.getGroupAndStatisticsProxy({
 *     projectId: 0,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getGroupAndStatisticsProxyOutput(args: GetGroupAndStatisticsProxyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupAndStatisticsProxyResult> {
    return pulumi.output(args).apply((a: any) => getGroupAndStatisticsProxy(a, opts))
}

/**
 * A collection of arguments for invoking getGroupAndStatisticsProxy.
 */
export interface GetGroupAndStatisticsProxyOutputArgs {
    /**
     * Project Id.
     */
    projectId: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
