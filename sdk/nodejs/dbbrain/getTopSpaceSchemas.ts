// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dbbrain topSpaceSchemas
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const topSpaceSchemas = tencentcloud.Dbbrain.getTopSpaceSchemas({
 *     instanceId: "%s",
 *     product: "mysql",
 *     sortBy: "DataLength",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getTopSpaceSchemas(args: GetTopSpaceSchemasArgs, opts?: pulumi.InvokeOptions): Promise<GetTopSpaceSchemasResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dbbrain/getTopSpaceSchemas:getTopSpaceSchemas", {
        "instanceId": args.instanceId,
        "limit": args.limit,
        "product": args.product,
        "resultOutputFile": args.resultOutputFile,
        "sortBy": args.sortBy,
    }, opts);
}

/**
 * A collection of arguments for invoking getTopSpaceSchemas.
 */
export interface GetTopSpaceSchemasArgs {
    /**
     * instance id.
     */
    instanceId: string;
    /**
     * The number of Top libraries to return, the maximum value is 100, and the default is 20.
     */
    limit?: number;
    /**
     * Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
     */
    product?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * The sorting field used to filter the Top library. The optional fields include DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, and PhysicalFileSize (only supported by ApsaraDB for MySQL instances). The default for ApsaraDB for MySQL instances is PhysicalFileSize, and the default for other product instances is TotalLength.
     */
    sortBy?: string;
}

/**
 * A collection of values returned by getTopSpaceSchemas.
 */
export interface GetTopSpaceSchemasResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly limit?: number;
    readonly product?: string;
    readonly resultOutputFile?: string;
    readonly sortBy?: string;
    /**
     * Timestamp (in seconds) when library space data is collected.
     */
    readonly timestamp: number;
    /**
     * The returned list of top library space statistics.
     */
    readonly topSpaceSchemas: outputs.Dbbrain.GetTopSpaceSchemasTopSpaceSchema[];
}
/**
 * Use this data source to query detailed information of dbbrain topSpaceSchemas
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const topSpaceSchemas = tencentcloud.Dbbrain.getTopSpaceSchemas({
 *     instanceId: "%s",
 *     product: "mysql",
 *     sortBy: "DataLength",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getTopSpaceSchemasOutput(args: GetTopSpaceSchemasOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTopSpaceSchemasResult> {
    return pulumi.output(args).apply((a: any) => getTopSpaceSchemas(a, opts))
}

/**
 * A collection of arguments for invoking getTopSpaceSchemas.
 */
export interface GetTopSpaceSchemasOutputArgs {
    /**
     * instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The number of Top libraries to return, the maximum value is 100, and the default is 20.
     */
    limit?: pulumi.Input<number>;
    /**
     * Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
     */
    product?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * The sorting field used to filter the Top library. The optional fields include DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, and PhysicalFileSize (only supported by ApsaraDB for MySQL instances). The default for ApsaraDB for MySQL instances is PhysicalFileSize, and the default for other product instances is TotalLength.
     */
    sortBy?: pulumi.Input<string>;
}
