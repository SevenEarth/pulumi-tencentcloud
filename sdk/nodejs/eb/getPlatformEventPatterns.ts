// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of eb platformEventPatterns
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const platformEventPatterns = pulumi.output(tencentcloud.Eb.getPlatformEventPatterns({
 *     productType: "",
 * }));
 * ```
 */
export function getPlatformEventPatterns(args: GetPlatformEventPatternsArgs, opts?: pulumi.InvokeOptions): Promise<GetPlatformEventPatternsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Eb/getPlatformEventPatterns:getPlatformEventPatterns", {
        "productType": args.productType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getPlatformEventPatterns.
 */
export interface GetPlatformEventPatternsArgs {
    /**
     * Platform product type.
     */
    productType: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getPlatformEventPatterns.
 */
export interface GetPlatformEventPatternsResult {
    /**
     * Platform product event matching rules.
     */
    readonly eventPatterns: outputs.Eb.GetPlatformEventPatternsEventPattern[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly productType: string;
    readonly resultOutputFile?: string;
}

export function getPlatformEventPatternsOutput(args: GetPlatformEventPatternsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPlatformEventPatternsResult> {
    return pulumi.output(args).apply(a => getPlatformEventPatterns(a, opts))
}

/**
 * A collection of arguments for invoking getPlatformEventPatterns.
 */
export interface GetPlatformEventPatternsOutputArgs {
    /**
     * Platform product type.
     */
    productType: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
