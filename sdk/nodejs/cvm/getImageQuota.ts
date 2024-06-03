// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cvm imageQuota
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const imageQuota = tencentcloud.Cvm.getImageQuota({});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getImageQuota(args?: GetImageQuotaArgs, opts?: pulumi.InvokeOptions): Promise<GetImageQuotaResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cvm/getImageQuota:getImageQuota", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getImageQuota.
 */
export interface GetImageQuotaArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getImageQuota.
 */
export interface GetImageQuotaResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The image quota of an account.
     */
    readonly imageNumQuota: number;
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of cvm imageQuota
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const imageQuota = tencentcloud.Cvm.getImageQuota({});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getImageQuotaOutput(args?: GetImageQuotaOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImageQuotaResult> {
    return pulumi.output(args).apply((a: any) => getImageQuota(a, opts))
}

/**
 * A collection of arguments for invoking getImageQuota.
 */
export interface GetImageQuotaOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
