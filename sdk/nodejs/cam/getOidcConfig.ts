// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cam oidcConfig
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const oidcConfig = tencentcloud.Cam.getOidcConfig({
 *     name: "cls-kzilgv5m",
 * });
 * export const identityKey = oidcConfig.then(oidcConfig => oidcConfig.identityKey);
 * export const identityUrl = oidcConfig.then(oidcConfig => oidcConfig.identityUrl);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getOidcConfig(args: GetOidcConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetOidcConfigResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cam/getOidcConfig:getOidcConfig", {
        "name": args.name,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getOidcConfig.
 */
export interface GetOidcConfigArgs {
    /**
     * Name.
     */
    name: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getOidcConfig.
 */
export interface GetOidcConfigResult {
    /**
     * Client ID.
     */
    readonly clientIds: string[];
    /**
     * Description.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Public key for signature.
     */
    readonly identityKey: string;
    /**
     * IdP URL.
     */
    readonly identityUrl: string;
    readonly name: string;
    /**
     * IdP type. 11: Role IdP.
     */
    readonly providerType: number;
    readonly resultOutputFile?: string;
    /**
     * Status. 0: Not set; 2: Disabled; 11: Enabled.
     */
    readonly status: number;
}
/**
 * Use this data source to query detailed information of cam oidcConfig
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const oidcConfig = tencentcloud.Cam.getOidcConfig({
 *     name: "cls-kzilgv5m",
 * });
 * export const identityKey = oidcConfig.then(oidcConfig => oidcConfig.identityKey);
 * export const identityUrl = oidcConfig.then(oidcConfig => oidcConfig.identityUrl);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getOidcConfigOutput(args: GetOidcConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOidcConfigResult> {
    return pulumi.output(args).apply((a: any) => getOidcConfig(a, opts))
}

/**
 * A collection of arguments for invoking getOidcConfig.
 */
export interface GetOidcConfigOutputArgs {
    /**
     * Name.
     */
    name: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
