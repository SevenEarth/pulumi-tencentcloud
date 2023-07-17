// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dcdb price
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const price = tencentcloud.Dcdb.getPrice({
 *     instanceCount: 1,
 *     zone: _var.default_az,
 *     period: 1,
 *     shardNodeCount: 2,
 *     shardMemory: 2,
 *     shardStorage: 10,
 *     shardCount: 2,
 *     paymode: "postpaid",
 *     amountUnit: "pent",
 * });
 * ```
 */
export function getPrice(args: GetPriceArgs, opts?: pulumi.InvokeOptions): Promise<GetPriceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Dcdb/getPrice:getPrice", {
        "amountUnit": args.amountUnit,
        "instanceCount": args.instanceCount,
        "paymode": args.paymode,
        "period": args.period,
        "resultOutputFile": args.resultOutputFile,
        "shardCount": args.shardCount,
        "shardMemory": args.shardMemory,
        "shardNodeCount": args.shardNodeCount,
        "shardStorage": args.shardStorage,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrice.
 */
export interface GetPriceArgs {
    /**
     * Price unit. Valid values: `pent` (cent), `microPent` (microcent).
     */
    amountUnit?: string;
    /**
     * The count of instances wants to buy.
     */
    instanceCount: number;
    /**
     * Billing type. Valid values: `postpaid` (pay-as-you-go), `prepaid` (monthly subscription).
     */
    paymode?: string;
    /**
     * Purchase period in months.
     */
    period: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Number of instance shards.
     */
    shardCount: number;
    /**
     * Shard memory size in GB.
     */
    shardMemory: number;
    /**
     * Number of instance shard nodes.
     */
    shardNodeCount: number;
    /**
     * Shard storage capacity in GB.
     */
    shardStorage: number;
    /**
     * AZ ID of the purchased instance.
     */
    zone: string;
}

/**
 * A collection of values returned by getPrice.
 */
export interface GetPriceResult {
    readonly amountUnit?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceCount: number;
    /**
     * Original price. Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description. Currency: CNY (Chinese site), USD (international site).
     */
    readonly originalPrice: number;
    readonly paymode?: string;
    readonly period: number;
    /**
     * The actual price may be different from the original price due to discounts. Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description. Currency: CNY (Chinese site), USD (international site).
     */
    readonly price: number;
    readonly resultOutputFile?: string;
    readonly shardCount: number;
    readonly shardMemory: number;
    readonly shardNodeCount: number;
    readonly shardStorage: number;
    readonly zone: string;
}

export function getPriceOutput(args: GetPriceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPriceResult> {
    return pulumi.output(args).apply(a => getPrice(a, opts))
}

/**
 * A collection of arguments for invoking getPrice.
 */
export interface GetPriceOutputArgs {
    /**
     * Price unit. Valid values: `pent` (cent), `microPent` (microcent).
     */
    amountUnit?: pulumi.Input<string>;
    /**
     * The count of instances wants to buy.
     */
    instanceCount: pulumi.Input<number>;
    /**
     * Billing type. Valid values: `postpaid` (pay-as-you-go), `prepaid` (monthly subscription).
     */
    paymode?: pulumi.Input<string>;
    /**
     * Purchase period in months.
     */
    period: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Number of instance shards.
     */
    shardCount: pulumi.Input<number>;
    /**
     * Shard memory size in GB.
     */
    shardMemory: pulumi.Input<number>;
    /**
     * Number of instance shard nodes.
     */
    shardNodeCount: pulumi.Input<number>;
    /**
     * Shard storage capacity in GB.
     */
    shardStorage: pulumi.Input<number>;
    /**
     * AZ ID of the purchased instance.
     */
    zone: pulumi.Input<string>;
}
