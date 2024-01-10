// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of organization orgFinancialByProduct
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const orgFinancialByProduct = pulumi.output(tencentcloud.Organization.getOrgFinancialByProduct({
 *     endMonth: "2023-09",
 *     month: "2023-05",
 *     productCodes: ["p_eip"],
 * }));
 * ```
 */
export function getOrgFinancialByProduct(args: GetOrgFinancialByProductArgs, opts?: pulumi.InvokeOptions): Promise<GetOrgFinancialByProductResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Organization/getOrgFinancialByProduct:getOrgFinancialByProduct", {
        "endMonth": args.endMonth,
        "memberUins": args.memberUins,
        "month": args.month,
        "productCodes": args.productCodes,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrgFinancialByProduct.
 */
export interface GetOrgFinancialByProductArgs {
    /**
     * Query for the end month. Format:yyyy-mm, for example:2021-01.The default value is the `Month`.
     */
    endMonth?: string;
    /**
     * Member uin list. Up to 100.
     */
    memberUins?: number[];
    /**
     * Query for the start month. Format:yyyy-mm, for example:2021-01.
     */
    month: string;
    /**
     * Product code list. Up to 100.
     */
    productCodes?: string[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getOrgFinancialByProduct.
 */
export interface GetOrgFinancialByProductResult {
    readonly endMonth?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Organization financial info by products.
     */
    readonly items: outputs.Organization.GetOrgFinancialByProductItem[];
    readonly memberUins?: number[];
    readonly month: string;
    readonly productCodes?: string[];
    readonly resultOutputFile?: string;
    /**
     * Total cost of the product.
     */
    readonly totalCost: number;
}

export function getOrgFinancialByProductOutput(args: GetOrgFinancialByProductOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrgFinancialByProductResult> {
    return pulumi.output(args).apply(a => getOrgFinancialByProduct(a, opts))
}

/**
 * A collection of arguments for invoking getOrgFinancialByProduct.
 */
export interface GetOrgFinancialByProductOutputArgs {
    /**
     * Query for the end month. Format:yyyy-mm, for example:2021-01.The default value is the `Month`.
     */
    endMonth?: pulumi.Input<string>;
    /**
     * Member uin list. Up to 100.
     */
    memberUins?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Query for the start month. Format:yyyy-mm, for example:2021-01.
     */
    month: pulumi.Input<string>;
    /**
     * Product code list. Up to 100.
     */
    productCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
