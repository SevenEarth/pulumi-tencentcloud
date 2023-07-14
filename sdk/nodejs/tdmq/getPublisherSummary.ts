// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tdmq publisherSummary
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const publisherSummary = pulumi.output(tencentcloud.Tdmq.getPublisherSummary({
 *     clusterId: "pulsar-9n95ax58b9vn",
 *     namespace: "keep-ns",
 *     topic: "keep-topic",
 * }));
 * ```
 */
export function getPublisherSummary(args: GetPublisherSummaryArgs, opts?: pulumi.InvokeOptions): Promise<GetPublisherSummaryResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Tdmq/getPublisherSummary:getPublisherSummary", {
        "clusterId": args.clusterId,
        "namespace": args.namespace,
        "resultOutputFile": args.resultOutputFile,
        "topic": args.topic,
    }, opts);
}

/**
 * A collection of arguments for invoking getPublisherSummary.
 */
export interface GetPublisherSummaryArgs {
    /**
     * Cluster ID.
     */
    clusterId: string;
    /**
     * namespace name.
     */
    namespace: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * subject name.
     */
    topic: string;
}

/**
 * A collection of values returned by getPublisherSummary.
 */
export interface GetPublisherSummaryResult {
    readonly clusterId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Production rate (units per second)Note: This field may return null, indicating that no valid value can be obtained.
     */
    readonly msgRateIn: number;
    /**
     * Production rate (bytes per second)Note: This field may return null, indicating that no valid value can be obtained.
     */
    readonly msgThroughputIn: number;
    readonly namespace: string;
    /**
     * number of producersNote: This field may return null, indicating that no valid value can be obtained.
     */
    readonly publisherCount: number;
    readonly resultOutputFile?: string;
    /**
     * Message store size in bytesNote: This field may return null, indicating that no valid value can be obtained.
     */
    readonly storageSize: number;
    readonly topic: string;
}

export function getPublisherSummaryOutput(args: GetPublisherSummaryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPublisherSummaryResult> {
    return pulumi.output(args).apply(a => getPublisherSummary(a, opts))
}

/**
 * A collection of arguments for invoking getPublisherSummary.
 */
export interface GetPublisherSummaryOutputArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * namespace name.
     */
    namespace: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * subject name.
     */
    topic: pulumi.Input<string>;
}
