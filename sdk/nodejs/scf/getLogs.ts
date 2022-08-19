// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query SCF function logs.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const fooFunction = new tencentcloud.scf.Function("fooFunction", {
 *     handler: "main.do_it",
 *     runtime: "Python3.6",
 *     cosBucketName: "scf-code-1234567890",
 *     cosObjectName: "code.zip",
 *     cosBucketRegion: "ap-guangzhou",
 * });
 * const fooLogs = tencentcloud.Scf.getLogsOutput({
 *     functionName: fooFunction.name,
 * });
 * ```
 */
export function getLogs(args: GetLogsArgs, opts?: pulumi.InvokeOptions): Promise<GetLogsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Scf/getLogs:getLogs", {
        "endTime": args.endTime,
        "functionName": args.functionName,
        "invokeRequestId": args.invokeRequestId,
        "limit": args.limit,
        "namespace": args.namespace,
        "offset": args.offset,
        "order": args.order,
        "orderBy": args.orderBy,
        "resultOutputFile": args.resultOutputFile,
        "retCode": args.retCode,
        "startTime": args.startTime,
    }, opts);
}

/**
 * A collection of arguments for invoking getLogs.
 */
export interface GetLogsArgs {
    /**
     * The end time of the query, the format is `2017-05-16 20:00:00`, which can only be within one day from `startTime`.
     */
    endTime?: string;
    /**
     * Name of the SCF function to be queried.
     */
    functionName: string;
    /**
     * Corresponding requestId when executing function.
     */
    invokeRequestId?: string;
    /**
     * Number of logs, the default is `10000`, offset+limit cannot be greater than 10000.
     */
    limit?: number;
    /**
     * Namespace of the SCF function to be queried.
     */
    namespace?: string;
    /**
     * Log offset, default is `0`, offset+limit cannot be greater than 10000.
     */
    offset?: number;
    /**
     * Order to sort the log, optional values `desc` and `asc`, default `desc`.
     */
    order?: string;
    /**
     * Sort the logs according to the following fields: `functionName`, `duration`, `memUsage`, `startTime`, default `startTime`.
     */
    orderBy?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Use to filter log, optional value: `not0` only returns the error log. `is0` only returns the correct log. `TimeLimitExceeded` returns the log of the function call timeout. `ResourceLimitExceeded` returns the function call generation resource overrun log. `UserCodeException` returns logs of the user code error that occurred in the function call. Not passing the parameter means returning all logs.
     */
    retCode?: string;
    /**
     * The start time of the query, the format is `2017-05-16 20:00:00`, which can only be within one day from `endTime`.
     */
    startTime?: string;
}

/**
 * A collection of values returned by getLogs.
 */
export interface GetLogsResult {
    readonly endTime?: string;
    /**
     * Name of the SCF function.
     */
    readonly functionName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly invokeRequestId?: string;
    readonly limit?: number;
    /**
     * An information list of logs. Each element contains the following attributes:
     */
    readonly logs: outputs.Scf.GetLogsLog[];
    readonly namespace?: string;
    readonly offset?: number;
    readonly order?: string;
    readonly orderBy?: string;
    readonly resultOutputFile?: string;
    /**
     * Execution result of function, `0` means the execution is successful, other values indicate failure.
     */
    readonly retCode?: string;
    /**
     * Point in time at which the function begins execution.
     */
    readonly startTime?: string;
}

export function getLogsOutput(args: GetLogsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLogsResult> {
    return pulumi.output(args).apply(a => getLogs(a, opts))
}

/**
 * A collection of arguments for invoking getLogs.
 */
export interface GetLogsOutputArgs {
    /**
     * The end time of the query, the format is `2017-05-16 20:00:00`, which can only be within one day from `startTime`.
     */
    endTime?: pulumi.Input<string>;
    /**
     * Name of the SCF function to be queried.
     */
    functionName: pulumi.Input<string>;
    /**
     * Corresponding requestId when executing function.
     */
    invokeRequestId?: pulumi.Input<string>;
    /**
     * Number of logs, the default is `10000`, offset+limit cannot be greater than 10000.
     */
    limit?: pulumi.Input<number>;
    /**
     * Namespace of the SCF function to be queried.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Log offset, default is `0`, offset+limit cannot be greater than 10000.
     */
    offset?: pulumi.Input<number>;
    /**
     * Order to sort the log, optional values `desc` and `asc`, default `desc`.
     */
    order?: pulumi.Input<string>;
    /**
     * Sort the logs according to the following fields: `functionName`, `duration`, `memUsage`, `startTime`, default `startTime`.
     */
    orderBy?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Use to filter log, optional value: `not0` only returns the error log. `is0` only returns the correct log. `TimeLimitExceeded` returns the log of the function call timeout. `ResourceLimitExceeded` returns the function call generation resource overrun log. `UserCodeException` returns logs of the user code error that occurred in the function call. Not passing the parameter means returning all logs.
     */
    retCode?: pulumi.Input<string>;
    /**
     * The start time of the query, the format is `2017-05-16 20:00:00`, which can only be within one day from `endTime`.
     */
    startTime?: pulumi.Input<string>;
}
