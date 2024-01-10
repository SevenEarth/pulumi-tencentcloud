// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of es elasticsearchInstanceLogs
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const elasticsearchInstanceLogs = pulumi.output(tencentcloud.Elasticsearch.getInstanceLogs({
 *     instanceId: "es-xxxxxx",
 * }));
 * ```
 */
export function getInstanceLogs(args: GetInstanceLogsArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceLogsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Elasticsearch/getInstanceLogs:getInstanceLogs", {
        "endTime": args.endTime,
        "instanceId": args.instanceId,
        "logType": args.logType,
        "orderByType": args.orderByType,
        "resultOutputFile": args.resultOutputFile,
        "searchKey": args.searchKey,
        "startTime": args.startTime,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceLogs.
 */
export interface GetInstanceLogsArgs {
    /**
     * End time. The format is YYYY-MM-DD HH:MM:SS, such as 2019-01-22 20:15:53.
     */
    endTime?: string;
    /**
     * Instance id.
     */
    instanceId: string;
    /**
     * Log type. Log type, default is 1, Valid values:
     * - 1: master log
     * - 2: Search slow log
     * - 3: Index slow log
     * - 4: GC log.
     */
    logType?: number;
    /**
     * Order type. Time sort method. Default is 0, valid values:
     * - 0: descending;
     * - 1: ascending order.
     */
    orderByType?: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Search key. Support LUCENE syntax, such as level:WARN, ip:1.1.1.1, message:test-index, etc.
     */
    searchKey?: string;
    /**
     * Start time. The format is YYYY-MM-DD HH:MM:SS, such as 2019-01-22 20:15:53.
     */
    startTime?: string;
}

/**
 * A collection of values returned by getInstanceLogs.
 */
export interface GetInstanceLogsResult {
    readonly endTime?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    /**
     * List of log details.
     */
    readonly instanceLogLists: outputs.Elasticsearch.GetInstanceLogsInstanceLogList[];
    readonly logType?: number;
    readonly orderByType?: number;
    readonly resultOutputFile?: string;
    readonly searchKey?: string;
    readonly startTime?: string;
}

export function getInstanceLogsOutput(args: GetInstanceLogsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceLogsResult> {
    return pulumi.output(args).apply(a => getInstanceLogs(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceLogs.
 */
export interface GetInstanceLogsOutputArgs {
    /**
     * End time. The format is YYYY-MM-DD HH:MM:SS, such as 2019-01-22 20:15:53.
     */
    endTime?: pulumi.Input<string>;
    /**
     * Instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Log type. Log type, default is 1, Valid values:
     * - 1: master log
     * - 2: Search slow log
     * - 3: Index slow log
     * - 4: GC log.
     */
    logType?: pulumi.Input<number>;
    /**
     * Order type. Time sort method. Default is 0, valid values:
     * - 0: descending;
     * - 1: ascending order.
     */
    orderByType?: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Search key. Support LUCENE syntax, such as level:WARN, ip:1.1.1.1, message:test-index, etc.
     */
    searchKey?: pulumi.Input<string>;
    /**
     * Start time. The format is YYYY-MM-DD HH:MM:SS, such as 2019-01-22 20:15:53.
     */
    startTime?: pulumi.Input<string>;
}
