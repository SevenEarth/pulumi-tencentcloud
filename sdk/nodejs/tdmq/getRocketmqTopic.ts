// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tdmqRocketmq topic
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const exampleRocketmqCluster = new tencentcloud.tdmq.RocketmqCluster("exampleRocketmqCluster", {
 *     clusterName: "tf_example",
 *     remark: "remark.",
 * });
 * const exampleRocketmqNamespace = new tencentcloud.tdmq.RocketmqNamespace("exampleRocketmqNamespace", {
 *     clusterId: exampleRocketmqCluster.clusterId,
 *     namespaceName: "tf_example",
 *     remark: "remark.",
 * });
 * const exampleRocketmqTopic = tencentcloud.Tdmq.getRocketmqTopicOutput({
 *     clusterId: exampleRocketmqCluster.clusterId,
 *     namespaceId: exampleRocketmqNamespace.namespaceName,
 *     filterName: exampleTdmq / rocketmqTopicRocketmqTopic.topicName,
 * });
 * const exampleTdmq_rocketmqTopicRocketmqTopic = new tencentcloud.tdmq.RocketmqTopic("exampleTdmq/rocketmqTopicRocketmqTopic", {
 *     topicName: "tf_example",
 *     namespaceName: exampleRocketmqNamespace.namespaceName,
 *     clusterId: exampleRocketmqCluster.clusterId,
 *     type: "Normal",
 *     remark: "remark.",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRocketmqTopic(args: GetRocketmqTopicArgs, opts?: pulumi.InvokeOptions): Promise<GetRocketmqTopicResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Tdmq/getRocketmqTopic:getRocketmqTopic", {
        "clusterId": args.clusterId,
        "filterName": args.filterName,
        "filterTypes": args.filterTypes,
        "namespaceId": args.namespaceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getRocketmqTopic.
 */
export interface GetRocketmqTopicArgs {
    /**
     * Cluster ID.
     */
    clusterId: string;
    /**
     * Search by topic name. Fuzzy query is supported.
     */
    filterName?: string;
    /**
     * Filter by topic type. Valid values: `Normal`, `GlobalOrder`, `PartitionedOrder`, `Transaction`.
     */
    filterTypes?: string[];
    /**
     * Namespace.
     */
    namespaceId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getRocketmqTopic.
 */
export interface GetRocketmqTopicResult {
    readonly clusterId: string;
    readonly filterName?: string;
    readonly filterTypes?: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly namespaceId: string;
    readonly resultOutputFile?: string;
    /**
     * List of topic information.
     */
    readonly topics: outputs.Tdmq.GetRocketmqTopicTopic[];
}
/**
 * Use this data source to query detailed information of tdmqRocketmq topic
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const exampleRocketmqCluster = new tencentcloud.tdmq.RocketmqCluster("exampleRocketmqCluster", {
 *     clusterName: "tf_example",
 *     remark: "remark.",
 * });
 * const exampleRocketmqNamespace = new tencentcloud.tdmq.RocketmqNamespace("exampleRocketmqNamespace", {
 *     clusterId: exampleRocketmqCluster.clusterId,
 *     namespaceName: "tf_example",
 *     remark: "remark.",
 * });
 * const exampleRocketmqTopic = tencentcloud.Tdmq.getRocketmqTopicOutput({
 *     clusterId: exampleRocketmqCluster.clusterId,
 *     namespaceId: exampleRocketmqNamespace.namespaceName,
 *     filterName: exampleTdmq / rocketmqTopicRocketmqTopic.topicName,
 * });
 * const exampleTdmq_rocketmqTopicRocketmqTopic = new tencentcloud.tdmq.RocketmqTopic("exampleTdmq/rocketmqTopicRocketmqTopic", {
 *     topicName: "tf_example",
 *     namespaceName: exampleRocketmqNamespace.namespaceName,
 *     clusterId: exampleRocketmqCluster.clusterId,
 *     type: "Normal",
 *     remark: "remark.",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRocketmqTopicOutput(args: GetRocketmqTopicOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRocketmqTopicResult> {
    return pulumi.output(args).apply((a: any) => getRocketmqTopic(a, opts))
}

/**
 * A collection of arguments for invoking getRocketmqTopic.
 */
export interface GetRocketmqTopicOutputArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Search by topic name. Fuzzy query is supported.
     */
    filterName?: pulumi.Input<string>;
    /**
     * Filter by topic type. Valid values: `Normal`, `GlobalOrder`, `PartitionedOrder`, `Transaction`.
     */
    filterTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Namespace.
     */
    namespaceId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
