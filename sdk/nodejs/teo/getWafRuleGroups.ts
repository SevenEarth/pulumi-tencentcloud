// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of teo wafRuleGroups
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const wafRuleGroups = pulumi.output(tencentcloud.Teo.getWafRuleGroups({
 *     entity: "",
 *     zoneId: "",
 * }));
 * ```
 */
export function getWafRuleGroups(args: GetWafRuleGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetWafRuleGroupsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Teo/getWafRuleGroups:getWafRuleGroups", {
        "entity": args.entity,
        "resultOutputFile": args.resultOutputFile,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getWafRuleGroups.
 */
export interface GetWafRuleGroupsArgs {
    /**
     * Subdomain or application name.
     */
    entity: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Site ID.
     */
    zoneId: string;
}

/**
 * A collection of values returned by getWafRuleGroups.
 */
export interface GetWafRuleGroupsResult {
    readonly entity: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * List of WAF rule groups.
     */
    readonly wafRuleGroups: outputs.Teo.GetWafRuleGroupsWafRuleGroup[];
    readonly zoneId: string;
}

export function getWafRuleGroupsOutput(args: GetWafRuleGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWafRuleGroupsResult> {
    return pulumi.output(args).apply(a => getWafRuleGroups(a, opts))
}

/**
 * A collection of arguments for invoking getWafRuleGroups.
 */
export interface GetWafRuleGroupsOutputArgs {
    /**
     * Subdomain or application name.
     */
    entity: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Site ID.
     */
    zoneId: pulumi.Input<string>;
}