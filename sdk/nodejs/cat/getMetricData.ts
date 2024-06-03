// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cat metricData
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const metricData = tencentcloud.Cat.getMetricData({
 *     analyzeTaskType: "AnalyzeTaskType_Network",
 *     field: "avg(\"ping_time\")",
 *     filters: [
 *         "\"host\" = 'www.qq.com'",
 *         "time >= now()-1h",
 *     ],
 *     metricType: "gauge",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getMetricData(args: GetMetricDataArgs, opts?: pulumi.InvokeOptions): Promise<GetMetricDataResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cat/getMetricData:getMetricData", {
        "analyzeTaskType": args.analyzeTaskType,
        "field": args.field,
        "filter": args.filter,
        "filters": args.filters,
        "groupBy": args.groupBy,
        "metricType": args.metricType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getMetricData.
 */
export interface GetMetricDataArgs {
    /**
     * Analysis of task type, supported types: `AnalyzeTaskType_Network`: network quality, `AnalyzeTaskType_Browse`: page performance, `AnalyzeTaskType_Transport`: port performance, `AnalyzeTaskType_UploadDownload`: file transport, `AnalyzeTaskType_MediaStream`: audiovisual experience.
     */
    analyzeTaskType: string;
    /**
     * Detailed fields of metrics, specified metrics can be passed or aggregate metrics, such as avg(ping_time) means entire delay.
     */
    field: string;
    /**
     * Filter conditions can be passed as a single filter or multiple parameters concatenated together.
     */
    filter?: string;
    /**
     * Multiple condition filtering, supports combining multiple filtering conditions for query.
     */
    filters: string[];
    /**
     * Aggregation time, such as 1m, 1d, 30d, and so on.
     */
    groupBy?: string;
    /**
     * Metric type, metrics queries are passed with gauge by default.
     */
    metricType: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getMetricData.
 */
export interface GetMetricDataResult {
    readonly analyzeTaskType: string;
    readonly field: string;
    readonly filter?: string;
    readonly filters: string[];
    readonly groupBy?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Return JSON string.
     */
    readonly metricSet: string;
    readonly metricType: string;
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of cat metricData
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const metricData = tencentcloud.Cat.getMetricData({
 *     analyzeTaskType: "AnalyzeTaskType_Network",
 *     field: "avg(\"ping_time\")",
 *     filters: [
 *         "\"host\" = 'www.qq.com'",
 *         "time >= now()-1h",
 *     ],
 *     metricType: "gauge",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getMetricDataOutput(args: GetMetricDataOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMetricDataResult> {
    return pulumi.output(args).apply((a: any) => getMetricData(a, opts))
}

/**
 * A collection of arguments for invoking getMetricData.
 */
export interface GetMetricDataOutputArgs {
    /**
     * Analysis of task type, supported types: `AnalyzeTaskType_Network`: network quality, `AnalyzeTaskType_Browse`: page performance, `AnalyzeTaskType_Transport`: port performance, `AnalyzeTaskType_UploadDownload`: file transport, `AnalyzeTaskType_MediaStream`: audiovisual experience.
     */
    analyzeTaskType: pulumi.Input<string>;
    /**
     * Detailed fields of metrics, specified metrics can be passed or aggregate metrics, such as avg(ping_time) means entire delay.
     */
    field: pulumi.Input<string>;
    /**
     * Filter conditions can be passed as a single filter or multiple parameters concatenated together.
     */
    filter?: pulumi.Input<string>;
    /**
     * Multiple condition filtering, supports combining multiple filtering conditions for query.
     */
    filters: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Aggregation time, such as 1m, 1d, 30d, and so on.
     */
    groupBy?: pulumi.Input<string>;
    /**
     * Metric type, metrics queries are passed with gauge by default.
     */
    metricType: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
