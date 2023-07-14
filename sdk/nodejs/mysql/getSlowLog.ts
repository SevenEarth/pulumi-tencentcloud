// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of mysql slowLog
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const slowLog = pulumi.output(tencentcloud.Mysql.getSlowLog({
 *     instanceId: "cdb-fitq5t9h",
 * }));
 * ```
 */
export function getSlowLog(args: GetSlowLogArgs, opts?: pulumi.InvokeOptions): Promise<GetSlowLogResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Mysql/getSlowLog:getSlowLog", {
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getSlowLog.
 */
export interface GetSlowLogArgs {
    /**
     * Instance ID, in the format: cdb-c1nl9rpv. Same instance ID as displayed in the ApsaraDB for Console page.
     */
    instanceId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getSlowLog.
 */
export interface GetSlowLogResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    /**
     * Details of slow query logs that meet the query conditions.
     */
    readonly items: outputs.Mysql.GetSlowLogItem[];
    readonly resultOutputFile?: string;
}

export function getSlowLogOutput(args: GetSlowLogOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSlowLogResult> {
    return pulumi.output(args).apply(a => getSlowLog(a, opts))
}

/**
 * A collection of arguments for invoking getSlowLog.
 */
export interface GetSlowLogOutputArgs {
    /**
     * Instance ID, in the format: cdb-c1nl9rpv. Same instance ID as displayed in the ApsaraDB for Console page.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
