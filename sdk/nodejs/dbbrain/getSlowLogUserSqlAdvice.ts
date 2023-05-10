// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dbbrain slowLogUserSqlAdvice
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const test = pulumi.output(tencentcloud.Dbbrain.getSlowLogUserSqlAdvice({
 *     instanceId: "%s",
 *     product: "mysql",
 *     sqlText: "%s",
 * }));
 * ```
 */
export function getSlowLogUserSqlAdvice(args: GetSlowLogUserSqlAdviceArgs, opts?: pulumi.InvokeOptions): Promise<GetSlowLogUserSqlAdviceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Dbbrain/getSlowLogUserSqlAdvice:getSlowLogUserSqlAdvice", {
        "instanceId": args.instanceId,
        "product": args.product,
        "resultOutputFile": args.resultOutputFile,
        "schema": args.schema,
        "sqlText": args.sqlText,
    }, opts);
}

/**
 * A collection of arguments for invoking getSlowLogUserSqlAdvice.
 */
export interface GetSlowLogUserSqlAdviceArgs {
    /**
     * instance id.
     */
    instanceId: string;
    /**
     * Service product type, supported values: `mysql` - cloud database MySQL; `cynosdb` - cloud database TDSQL-C for MySQL; `dbbrain-mysql` - self-built MySQL, the default is `mysql`.
     */
    product?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * library name.
     */
    schema?: string;
    /**
     * SQL statements.
     */
    sqlText: string;
}

/**
 * A collection of values returned by getSlowLogUserSqlAdvice.
 */
export interface GetSlowLogUserSqlAdviceResult {
    /**
     * SQL optimization suggestion, which can be parsed into a JSON array, and the output is empty when no optimization is required.
     */
    readonly advices: string;
    /**
     * SQL optimization suggestion remarks, which can be parsed into a String array, and the output is empty when optimization is not required.
     */
    readonly comments: string;
    /**
     * The cost saving details after SQL optimization can be parsed as JSON, and the output is empty when no optimization is required.
     */
    readonly cost: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly product?: string;
    /**
     * Unique request ID, returned for every request. The RequestId of the request needs to be provided when locating the problem.
     */
    readonly requestId: string;
    readonly resultOutputFile?: string;
    readonly schema: string;
    /**
     * The SQL execution plan can be parsed into JSON, and the output is empty when no optimization is required.
     */
    readonly sqlPlan: string;
    readonly sqlText: string;
    /**
     * The DDL information of related tables can be parsed into a JSON array.
     */
    readonly tables: string;
}

export function getSlowLogUserSqlAdviceOutput(args: GetSlowLogUserSqlAdviceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSlowLogUserSqlAdviceResult> {
    return pulumi.output(args).apply(a => getSlowLogUserSqlAdvice(a, opts))
}

/**
 * A collection of arguments for invoking getSlowLogUserSqlAdvice.
 */
export interface GetSlowLogUserSqlAdviceOutputArgs {
    /**
     * instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Service product type, supported values: `mysql` - cloud database MySQL; `cynosdb` - cloud database TDSQL-C for MySQL; `dbbrain-mysql` - self-built MySQL, the default is `mysql`.
     */
    product?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * library name.
     */
    schema?: pulumi.Input<string>;
    /**
     * SQL statements.
     */
    sqlText: pulumi.Input<string>;
}
