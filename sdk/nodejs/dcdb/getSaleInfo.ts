// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dcdb saleInfo
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const saleInfo = tencentcloud.Dcdb.getSaleInfo({});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSaleInfo(args?: GetSaleInfoArgs, opts?: pulumi.InvokeOptions): Promise<GetSaleInfoResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dcdb/getSaleInfo:getSaleInfo", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getSaleInfo.
 */
export interface GetSaleInfoArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getSaleInfo.
 */
export interface GetSaleInfoResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * list of sale region info.
     */
    readonly regionLists: outputs.Dcdb.GetSaleInfoRegionList[];
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of dcdb saleInfo
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const saleInfo = tencentcloud.Dcdb.getSaleInfo({});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSaleInfoOutput(args?: GetSaleInfoOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSaleInfoResult> {
    return pulumi.output(args).apply((a: any) => getSaleInfo(a, opts))
}

/**
 * A collection of arguments for invoking getSaleInfo.
 */
export interface GetSaleInfoOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
