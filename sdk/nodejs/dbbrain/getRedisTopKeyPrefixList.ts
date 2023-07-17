// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dbbrain redisTopKeyPrefixList
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const redisTopKeyPrefixList = tencentcloud.Dbbrain.getRedisTopKeyPrefixList({
 *     instanceId: local.redis_id,
 *     date: `%s`,
 *     product: "redis",
 * });
 * ```
 */
export function getRedisTopKeyPrefixList(args: GetRedisTopKeyPrefixListArgs, opts?: pulumi.InvokeOptions): Promise<GetRedisTopKeyPrefixListResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Dbbrain/getRedisTopKeyPrefixList:getRedisTopKeyPrefixList", {
        "date": args.date,
        "instanceId": args.instanceId,
        "product": args.product,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getRedisTopKeyPrefixList.
 */
export interface GetRedisTopKeyPrefixListArgs {
    /**
     * Query date, such as 2021-05-27, the earliest date can be the previous 30 days.
     */
    date: string;
    /**
     * instance id.
     */
    instanceId: string;
    /**
     * Service product type, supported values include `redis` - cloud database Redis.
     */
    product: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getRedisTopKeyPrefixList.
 */
export interface GetRedisTopKeyPrefixListResult {
    readonly date: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    /**
     * list of top key prefixes.
     */
    readonly items: outputs.Dbbrain.GetRedisTopKeyPrefixListItem[];
    readonly product: string;
    readonly resultOutputFile?: string;
}

export function getRedisTopKeyPrefixListOutput(args: GetRedisTopKeyPrefixListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRedisTopKeyPrefixListResult> {
    return pulumi.output(args).apply(a => getRedisTopKeyPrefixList(a, opts))
}

/**
 * A collection of arguments for invoking getRedisTopKeyPrefixList.
 */
export interface GetRedisTopKeyPrefixListOutputArgs {
    /**
     * Query date, such as 2021-05-27, the earliest date can be the previous 30 days.
     */
    date: pulumi.Input<string>;
    /**
     * instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Service product type, supported values include `redis` - cloud database Redis.
     */
    product: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
