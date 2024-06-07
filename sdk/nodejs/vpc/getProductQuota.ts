// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vpc productQuota
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const productQuota = tencentcloud.Vpc.getProductQuota({
 *     product: "vpc",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getProductQuota(args: GetProductQuotaArgs, opts?: pulumi.InvokeOptions): Promise<GetProductQuotaResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Vpc/getProductQuota:getProductQuota", {
        "product": args.product,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getProductQuota.
 */
export interface GetProductQuotaArgs {
    /**
     * The name of the network product to be queried. The products that can be queried are:vpc, ccn, vpn, dc, dfw, clb, eip.
     */
    product: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getProductQuota.
 */
export interface GetProductQuotaResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly product: string;
    /**
     * ProductQuota Array.
     */
    readonly productQuotaSets: outputs.Vpc.GetProductQuotaProductQuotaSet[];
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of vpc productQuota
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const productQuota = tencentcloud.Vpc.getProductQuota({
 *     product: "vpc",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getProductQuotaOutput(args: GetProductQuotaOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProductQuotaResult> {
    return pulumi.output(args).apply((a: any) => getProductQuota(a, opts))
}

/**
 * A collection of arguments for invoking getProductQuota.
 */
export interface GetProductQuotaOutputArgs {
    /**
     * The name of the network product to be queried. The products that can be queried are:vpc, ccn, vpn, dc, dfw, clb, eip.
     */
    product: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
