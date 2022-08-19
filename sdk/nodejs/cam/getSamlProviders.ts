// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of CAM SAML providers
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = pulumi.output(tencentcloud.Cam.getSamlProviders({
 *     name: "cam-test-provider",
 * }));
 * ```
 */
export function getSamlProviders(args?: GetSamlProvidersArgs, opts?: pulumi.InvokeOptions): Promise<GetSamlProvidersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Cam/getSamlProviders:getSamlProviders", {
        "description": args.description,
        "name": args.name,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getSamlProviders.
 */
export interface GetSamlProvidersArgs {
    /**
     * The description of the CAM SAML provider.
     */
    description?: string;
    /**
     * Name of the CAM SAML provider to be queried.
     */
    name?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getSamlProviders.
 */
export interface GetSamlProvidersResult {
    /**
     * Description of CAM SAML provider.
     */
    readonly description?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of CAM SAML provider.
     */
    readonly name?: string;
    /**
     * A list of CAM SAML providers. Each element contains the following attributes:
     */
    readonly providerLists: outputs.Cam.GetSamlProvidersProviderList[];
    readonly resultOutputFile?: string;
}

export function getSamlProvidersOutput(args?: GetSamlProvidersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSamlProvidersResult> {
    return pulumi.output(args).apply(a => getSamlProviders(a, opts))
}

/**
 * A collection of arguments for invoking getSamlProviders.
 */
export interface GetSamlProvidersOutputArgs {
    /**
     * The description of the CAM SAML provider.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the CAM SAML provider to be queried.
     */
    name?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
