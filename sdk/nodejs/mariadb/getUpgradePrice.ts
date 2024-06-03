// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of mariadb upgradePrice
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const upgradePrice = tencentcloud.Mariadb.getUpgradePrice({
 *     instanceId: "tdsql-9vqvls95",
 *     memory: 4,
 *     nodeCount: 2,
 *     storage: 40,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getUpgradePrice(args: GetUpgradePriceArgs, opts?: pulumi.InvokeOptions): Promise<GetUpgradePriceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Mariadb/getUpgradePrice:getUpgradePrice", {
        "amountUnit": args.amountUnit,
        "instanceId": args.instanceId,
        "memory": args.memory,
        "nodeCount": args.nodeCount,
        "resultOutputFile": args.resultOutputFile,
        "storage": args.storage,
    }, opts);
}

/**
 * A collection of arguments for invoking getUpgradePrice.
 */
export interface GetUpgradePriceArgs {
    /**
     * Price unit. Valid values: `* pent` (cent), `* microPent` (microcent).
     */
    amountUnit?: string;
    /**
     * Instance ID.
     */
    instanceId: string;
    /**
     * Memory size in GB, which can be obtained by querying the instance specification through the `DescribeDBInstanceSpecs` API.
     */
    memory: number;
    /**
     * New instance nodes, zero means not change.
     */
    nodeCount?: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Storage capacity in GB. The maximum and minimum storage space can be obtained by querying instance specification through the `DescribeDBInstanceSpecs` API.
     */
    storage: number;
}

/**
 * A collection of values returned by getUpgradePrice.
 */
export interface GetUpgradePriceResult {
    readonly amountUnit?: string;
    /**
     * Price calculation formula.
     */
    readonly formula: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly memory: number;
    readonly nodeCount?: number;
    /**
     * Original price * Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description. * Currency: CNY (Chinese site), USD (international site).
     */
    readonly originalPrice: number;
    /**
     * The actual price may be different from the original price due to discounts. * Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description. * Currency: CNY (Chinese site), USD (international site).
     */
    readonly price: number;
    readonly resultOutputFile?: string;
    readonly storage: number;
}
/**
 * Use this data source to query detailed information of mariadb upgradePrice
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const upgradePrice = tencentcloud.Mariadb.getUpgradePrice({
 *     instanceId: "tdsql-9vqvls95",
 *     memory: 4,
 *     nodeCount: 2,
 *     storage: 40,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getUpgradePriceOutput(args: GetUpgradePriceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUpgradePriceResult> {
    return pulumi.output(args).apply((a: any) => getUpgradePrice(a, opts))
}

/**
 * A collection of arguments for invoking getUpgradePrice.
 */
export interface GetUpgradePriceOutputArgs {
    /**
     * Price unit. Valid values: `* pent` (cent), `* microPent` (microcent).
     */
    amountUnit?: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Memory size in GB, which can be obtained by querying the instance specification through the `DescribeDBInstanceSpecs` API.
     */
    memory: pulumi.Input<number>;
    /**
     * New instance nodes, zero means not change.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Storage capacity in GB. The maximum and minimum storage space can be obtained by querying instance specification through the `DescribeDBInstanceSpecs` API.
     */
    storage: pulumi.Input<number>;
}
