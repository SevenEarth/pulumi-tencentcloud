// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query the available mongodb specifications for different zone.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const mongodb = tencentcloud.Mongodb.getZoneConfig({
 *     availableZone: "ap-guangzhou-2",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getZoneConfig(args?: GetZoneConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetZoneConfigResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Mongodb/getZoneConfig:getZoneConfig", {
        "availableZone": args.availableZone,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getZoneConfig.
 */
export interface GetZoneConfigArgs {
    /**
     * The available zone of the Mongodb.
     */
    availableZone?: string;
    /**
     * Used to store results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getZoneConfig.
 */
export interface GetZoneConfigResult {
    /**
     * The available zone of the Mongodb.
     */
    readonly availableZone?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of zone config. Each element contains the following attributes:
     */
    readonly lists: outputs.Mongodb.GetZoneConfigList[];
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query the available mongodb specifications for different zone.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const mongodb = tencentcloud.Mongodb.getZoneConfig({
 *     availableZone: "ap-guangzhou-2",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getZoneConfigOutput(args?: GetZoneConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZoneConfigResult> {
    return pulumi.output(args).apply((a: any) => getZoneConfig(a, opts))
}

/**
 * A collection of arguments for invoking getZoneConfig.
 */
export interface GetZoneConfigOutputArgs {
    /**
     * The available zone of the Mongodb.
     */
    availableZone?: pulumi.Input<string>;
    /**
     * Used to store results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
