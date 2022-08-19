// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query purchasable specification configuration for each availability zone in this specific region.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const mysqlserver = pulumi.output(tencentcloud.Sqlserver.getZoneConfig());
 * ```
 */
export function getZoneConfig(args?: GetZoneConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetZoneConfigResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Sqlserver/getZoneConfig:getZoneConfig", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getZoneConfig.
 */
export interface GetZoneConfigArgs {
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
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * A list of availability zones. Each element contains the following attributes:
     */
    readonly zoneLists: outputs.Sqlserver.GetZoneConfigZoneList[];
}

export function getZoneConfigOutput(args?: GetZoneConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZoneConfigResult> {
    return pulumi.output(args).apply(a => getZoneConfig(a, opts))
}

/**
 * A collection of arguments for invoking getZoneConfig.
 */
export interface GetZoneConfigOutputArgs {
    /**
     * Used to store results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
