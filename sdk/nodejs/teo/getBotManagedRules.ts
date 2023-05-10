// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of teo botManagedRules
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const botManagedRules = pulumi.output(tencentcloud.Teo.getBotManagedRules({
 *     entity: "",
 *     zoneId: "",
 * }));
 * ```
 */
export function getBotManagedRules(args: GetBotManagedRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetBotManagedRulesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Teo/getBotManagedRules:getBotManagedRules", {
        "entity": args.entity,
        "resultOutputFile": args.resultOutputFile,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getBotManagedRules.
 */
export interface GetBotManagedRulesArgs {
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
 * A collection of values returned by getBotManagedRules.
 */
export interface GetBotManagedRulesResult {
    readonly entity: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * Managed rules list.
     */
    readonly rules: outputs.Teo.GetBotManagedRulesRule[];
    readonly zoneId: string;
}

export function getBotManagedRulesOutput(args: GetBotManagedRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBotManagedRulesResult> {
    return pulumi.output(args).apply(a => getBotManagedRules(a, opts))
}

/**
 * A collection of arguments for invoking getBotManagedRules.
 */
export interface GetBotManagedRulesOutputArgs {
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
