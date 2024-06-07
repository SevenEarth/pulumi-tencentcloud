// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of ssm products
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const products = tencentcloud.Ssm.getProducts({});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getProducts(args?: GetProductsArgs, opts?: pulumi.InvokeOptions): Promise<GetProductsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Ssm/getProducts:getProducts", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getProducts.
 */
export interface GetProductsArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getProducts.
 */
export interface GetProductsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of supported services.
     */
    readonly products: string[];
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of ssm products
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const products = tencentcloud.Ssm.getProducts({});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getProductsOutput(args?: GetProductsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProductsResult> {
    return pulumi.output(args).apply((a: any) => getProducts(a, opts))
}

/**
 * A collection of arguments for invoking getProducts.
 */
export interface GetProductsOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
