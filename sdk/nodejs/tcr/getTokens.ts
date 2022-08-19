// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of TCR tokens.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const name = pulumi.output(tencentcloud.Tcr.getTokens({
 *     instanceId: "cls-satg5125",
 * }));
 * ```
 */
export function getTokens(args: GetTokensArgs, opts?: pulumi.InvokeOptions): Promise<GetTokensResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Tcr/getTokens:getTokens", {
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
        "tokenId": args.tokenId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTokens.
 */
export interface GetTokensArgs {
    /**
     * ID of the instance that the token belongs to.
     */
    instanceId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * ID of the TCR token to query.
     */
    tokenId?: string;
}

/**
 * A collection of values returned by getTokens.
 */
export interface GetTokensResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly resultOutputFile?: string;
    /**
     * Id of TCR token.
     */
    readonly tokenId?: string;
    /**
     * Information list of the dedicated TCR tokens.
     */
    readonly tokenLists: outputs.Tcr.GetTokensTokenList[];
}

export function getTokensOutput(args: GetTokensOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTokensResult> {
    return pulumi.output(args).apply(a => getTokens(a, opts))
}

/**
 * A collection of arguments for invoking getTokens.
 */
export interface GetTokensOutputArgs {
    /**
     * ID of the instance that the token belongs to.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * ID of the TCR token to query.
     */
    tokenId?: pulumi.Input<string>;
}
