// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dbbrain sqlTemplates
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const sqlTemplates = pulumi.output(tencentcloud.Dbbrain.getSqlTemplates({
 *     instanceId: "",
 *     product: "",
 *     schema: "",
 *     sqlText: "",
 * }));
 * ```
 */
export function getSqlTemplates(args: GetSqlTemplatesArgs, opts?: pulumi.InvokeOptions): Promise<GetSqlTemplatesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Dbbrain/getSqlTemplates:getSqlTemplates", {
        "instanceId": args.instanceId,
        "product": args.product,
        "resultOutputFile": args.resultOutputFile,
        "schema": args.schema,
        "sqlText": args.sqlText,
    }, opts);
}

/**
 * A collection of arguments for invoking getSqlTemplates.
 */
export interface GetSqlTemplatesArgs {
    /**
     * instance id.
     */
    instanceId: string;
    /**
     * Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
     */
    product?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * database name.
     */
    schema: string;
    /**
     * SQL statements.
     */
    sqlText: string;
}

/**
 * A collection of values returned by getSqlTemplates.
 */
export interface GetSqlTemplatesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly product?: string;
    readonly resultOutputFile?: string;
    readonly schema: string;
    /**
     * SQL template ID.
     */
    readonly sqlId: number;
    /**
     * SQL template content.
     */
    readonly sqlTemplate: string;
    readonly sqlText: string;
    /**
     * sql type.
     */
    readonly sqlType: string;
}

export function getSqlTemplatesOutput(args: GetSqlTemplatesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSqlTemplatesResult> {
    return pulumi.output(args).apply(a => getSqlTemplates(a, opts))
}

/**
 * A collection of arguments for invoking getSqlTemplates.
 */
export interface GetSqlTemplatesOutputArgs {
    /**
     * instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
     */
    product?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * database name.
     */
    schema: pulumi.Input<string>;
    /**
     * SQL statements.
     */
    sqlText: pulumi.Input<string>;
}
