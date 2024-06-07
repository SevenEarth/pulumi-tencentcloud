// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tsf deliveryConfigs
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const deliveryConfigs = tencentcloud.Tsf.getDeliveryConfigs({
 *     searchWord: "test",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDeliveryConfigs(args?: GetDeliveryConfigsArgs, opts?: pulumi.InvokeOptions): Promise<GetDeliveryConfigsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Tsf/getDeliveryConfigs:getDeliveryConfigs", {
        "resultOutputFile": args.resultOutputFile,
        "searchWord": args.searchWord,
    }, opts);
}

/**
 * A collection of arguments for invoking getDeliveryConfigs.
 */
export interface GetDeliveryConfigsArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * search word.
     */
    searchWord?: string;
}

/**
 * A collection of values returned by getDeliveryConfigs.
 */
export interface GetDeliveryConfigsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * deploy group information about the deployment group associated with a delivery item.Note: This field may return null, which means that no valid value was obtained.
     */
    readonly results: outputs.Tsf.GetDeliveryConfigsResult[];
    readonly searchWord?: string;
}
/**
 * Use this data source to query detailed information of tsf deliveryConfigs
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const deliveryConfigs = tencentcloud.Tsf.getDeliveryConfigs({
 *     searchWord: "test",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDeliveryConfigsOutput(args?: GetDeliveryConfigsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeliveryConfigsResult> {
    return pulumi.output(args).apply((a: any) => getDeliveryConfigs(a, opts))
}

/**
 * A collection of arguments for invoking getDeliveryConfigs.
 */
export interface GetDeliveryConfigsOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * search word.
     */
    searchWord?: pulumi.Input<string>;
}
