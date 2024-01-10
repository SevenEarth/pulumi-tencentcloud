// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of waf attackTotalCount
 *
 * ## Example Usage
 * ### Obtain the specified domain name attack log
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = pulumi.output(tencentcloud.Waf.getAttackTotalCount({
 *     domain: "domain.com",
 *     endTime: "2023-09-07 00:00:00",
 *     queryString: "method:GET",
 *     startTime: "2023-09-01 00:00:00",
 * }));
 * ```
 * ### Obtain all domain name attack log
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = pulumi.output(tencentcloud.Waf.getAttackTotalCount({
 *     domain: "all",
 *     endTime: "2023-09-07 00:00:00",
 *     queryString: "method:GET",
 *     startTime: "2023-09-01 00:00:00",
 * }));
 * ```
 */
export function getAttackTotalCount(args: GetAttackTotalCountArgs, opts?: pulumi.InvokeOptions): Promise<GetAttackTotalCountResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Waf/getAttackTotalCount:getAttackTotalCount", {
        "domain": args.domain,
        "endTime": args.endTime,
        "queryString": args.queryString,
        "resultOutputFile": args.resultOutputFile,
        "startTime": args.startTime,
    }, opts);
}

/**
 * A collection of arguments for invoking getAttackTotalCount.
 */
export interface GetAttackTotalCountArgs {
    /**
     * Query domain name, all domain use all.
     */
    domain: string;
    /**
     * End time.
     */
    endTime: string;
    /**
     * Query conditions.
     */
    queryString?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Begin time.
     */
    startTime: string;
}

/**
 * A collection of values returned by getAttackTotalCount.
 */
export interface GetAttackTotalCountResult {
    readonly domain: string;
    readonly endTime: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly queryString?: string;
    readonly resultOutputFile?: string;
    readonly startTime: string;
    /**
     * Total number of attacks.
     */
    readonly totalCount: number;
}

export function getAttackTotalCountOutput(args: GetAttackTotalCountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAttackTotalCountResult> {
    return pulumi.output(args).apply(a => getAttackTotalCount(a, opts))
}

/**
 * A collection of arguments for invoking getAttackTotalCount.
 */
export interface GetAttackTotalCountOutputArgs {
    /**
     * Query domain name, all domain use all.
     */
    domain: pulumi.Input<string>;
    /**
     * End time.
     */
    endTime: pulumi.Input<string>;
    /**
     * Query conditions.
     */
    queryString?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Begin time.
     */
    startTime: pulumi.Input<string>;
}
