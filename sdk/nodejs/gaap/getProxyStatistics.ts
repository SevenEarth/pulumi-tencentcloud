// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of gaap proxy statistics
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const proxyStatistics = pulumi.output(tencentcloud.Gaap.getProxyStatistics({
 *     endTime: "2023-10-09 23:59:59",
 *     granularity: 300,
 *     metricNames: [
 *         "InBandwidth",
 *         "OutBandwidth",
 *         "InFlow",
 *         "OutFlow",
 *         "InPackets",
 *         "OutPackets",
 *         "Concurrent",
 *         "HttpQPS",
 *         "HttpsQPS",
 *         "Latency",
 *         "PacketLoss",
 *     ],
 *     proxyId: "link-8lpyo88p",
 *     startTime: "2023-10-09 00:00:00",
 * }));
 * ```
 */
export function getProxyStatistics(args: GetProxyStatisticsArgs, opts?: pulumi.InvokeOptions): Promise<GetProxyStatisticsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Gaap/getProxyStatistics:getProxyStatistics", {
        "endTime": args.endTime,
        "granularity": args.granularity,
        "isp": args.isp,
        "metricNames": args.metricNames,
        "proxyId": args.proxyId,
        "resultOutputFile": args.resultOutputFile,
        "startTime": args.startTime,
    }, opts);
}

/**
 * A collection of arguments for invoking getProxyStatistics.
 */
export interface GetProxyStatisticsArgs {
    /**
     * End Time(2019-03-25 12:00:00).
     */
    endTime: string;
    /**
     * Monitoring granularity, currently supporting 60 300 3600 86400, in seconds.When the time range does not exceed 3 days, support a minimum granularity of 60 seconds;When the time range does not exceed 7 days, support a minimum granularity of 300 seconds;When the time range does not exceed 30 days, the minimum granularity supported is 3600 seconds.
     */
    granularity: number;
    /**
     * Operator (valid when the proxy is a three network proxy), supports CMCC, CUCC, CTCC, and merges data from the three operators if null values are passed or not passed.
     */
    isp?: string;
    /**
     * Metric Names. Valid values: InBandwidth,OutBandwidth, Concurrent, InPackets, OutPackets, PacketLoss, Latency, HttpQPS, HttpsQPS.
     */
    metricNames: string[];
    /**
     * Proxy Id.
     */
    proxyId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Start Time(2019-03-25 12:00:00).
     */
    startTime: string;
}

/**
 * A collection of values returned by getProxyStatistics.
 */
export interface GetProxyStatisticsResult {
    readonly endTime: string;
    readonly granularity: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isp?: string;
    readonly metricNames: string[];
    readonly proxyId: string;
    readonly resultOutputFile?: string;
    readonly startTime: string;
    /**
     * proxy Statistics.
     */
    readonly statisticsDatas: outputs.Gaap.GetProxyStatisticsStatisticsData[];
}

export function getProxyStatisticsOutput(args: GetProxyStatisticsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProxyStatisticsResult> {
    return pulumi.output(args).apply(a => getProxyStatistics(a, opts))
}

/**
 * A collection of arguments for invoking getProxyStatistics.
 */
export interface GetProxyStatisticsOutputArgs {
    /**
     * End Time(2019-03-25 12:00:00).
     */
    endTime: pulumi.Input<string>;
    /**
     * Monitoring granularity, currently supporting 60 300 3600 86400, in seconds.When the time range does not exceed 3 days, support a minimum granularity of 60 seconds;When the time range does not exceed 7 days, support a minimum granularity of 300 seconds;When the time range does not exceed 30 days, the minimum granularity supported is 3600 seconds.
     */
    granularity: pulumi.Input<number>;
    /**
     * Operator (valid when the proxy is a three network proxy), supports CMCC, CUCC, CTCC, and merges data from the three operators if null values are passed or not passed.
     */
    isp?: pulumi.Input<string>;
    /**
     * Metric Names. Valid values: InBandwidth,OutBandwidth, Concurrent, InPackets, OutPackets, PacketLoss, Latency, HttpQPS, HttpsQPS.
     */
    metricNames: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Proxy Id.
     */
    proxyId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Start Time(2019-03-25 12:00:00).
     */
    startTime: pulumi.Input<string>;
}
