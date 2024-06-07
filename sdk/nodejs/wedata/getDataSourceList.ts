// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of wedata dataSourceList
 *
 * ## Example Usage
 *
 * ### Query All
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Wedata.getDataSourceList({});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Query By filter
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Wedata.getDataSourceList({
 *     filters: [{
 *         name: "Name",
 *         values: ["tf_example"],
 *     }],
 *     orderFields: [{
 *         direction: "DESC",
 *         name: "create_time",
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDataSourceList(args?: GetDataSourceListArgs, opts?: pulumi.InvokeOptions): Promise<GetDataSourceListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Wedata/getDataSourceList:getDataSourceList", {
        "filters": args.filters,
        "orderFields": args.orderFields,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDataSourceList.
 */
export interface GetDataSourceListArgs {
    /**
     * Filters.
     */
    filters?: inputs.Wedata.GetDataSourceListFilter[];
    /**
     * OrderFields.
     */
    orderFields?: inputs.Wedata.GetDataSourceListOrderField[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDataSourceList.
 */
export interface GetDataSourceListResult {
    readonly filters?: outputs.Wedata.GetDataSourceListFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly orderFields?: outputs.Wedata.GetDataSourceListOrderField[];
    readonly resultOutputFile?: string;
    /**
     * Data rows.
     */
    readonly rows: outputs.Wedata.GetDataSourceListRow[];
}
/**
 * Use this data source to query detailed information of wedata dataSourceList
 *
 * ## Example Usage
 *
 * ### Query All
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Wedata.getDataSourceList({});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Query By filter
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Wedata.getDataSourceList({
 *     filters: [{
 *         name: "Name",
 *         values: ["tf_example"],
 *     }],
 *     orderFields: [{
 *         direction: "DESC",
 *         name: "create_time",
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDataSourceListOutput(args?: GetDataSourceListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataSourceListResult> {
    return pulumi.output(args).apply((a: any) => getDataSourceList(a, opts))
}

/**
 * A collection of arguments for invoking getDataSourceList.
 */
export interface GetDataSourceListOutputArgs {
    /**
     * Filters.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.Wedata.GetDataSourceListFilterArgs>[]>;
    /**
     * OrderFields.
     */
    orderFields?: pulumi.Input<pulumi.Input<inputs.Wedata.GetDataSourceListOrderFieldArgs>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
