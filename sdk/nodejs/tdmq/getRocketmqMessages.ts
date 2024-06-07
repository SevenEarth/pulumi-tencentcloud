// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tdmq message
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const message = tencentcloud.Tdmq.getRocketmqMessages({
 *     clusterId: "rocketmq-rkrbm52djmro",
 *     environmentId: "keep_ns",
 *     msgId: "A9FE8D0567FE15DB97425FC08EEF0000",
 *     queryDlqMsg: false,
 *     topicName: "keep-topic",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRocketmqMessages(args: GetRocketmqMessagesArgs, opts?: pulumi.InvokeOptions): Promise<GetRocketmqMessagesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Tdmq/getRocketmqMessages:getRocketmqMessages", {
        "clusterId": args.clusterId,
        "environmentId": args.environmentId,
        "msgId": args.msgId,
        "queryDlqMsg": args.queryDlqMsg,
        "resultOutputFile": args.resultOutputFile,
        "topicName": args.topicName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRocketmqMessages.
 */
export interface GetRocketmqMessagesArgs {
    /**
     * Cluster id.
     */
    clusterId: string;
    /**
     * Environment.
     */
    environmentId: string;
    /**
     * Message ID.
     */
    msgId: string;
    /**
     * The value is true when querying dead letters, only valid for Rocketmq.
     */
    queryDlqMsg?: boolean;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Topic, groupId is passed when querying dead letters.
     */
    topicName: string;
}

/**
 * A collection of values returned by getRocketmqMessages.
 */
export interface GetRocketmqMessagesResult {
    /**
     * Message body.
     */
    readonly body: string;
    readonly clusterId: string;
    readonly environmentId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Consumer Group ConsumptionNote: This field may return null, indicating that no valid value can be obtained.
     */
    readonly messageTracks: outputs.Tdmq.GetRocketmqMessagesMessageTrack[];
    readonly msgId: string;
    /**
     * Production time.
     */
    readonly produceTime: string;
    /**
     * Producer address.
     */
    readonly producerAddr: string;
    /**
     * Detailed parameters.
     */
    readonly properties: string;
    readonly queryDlqMsg?: boolean;
    readonly resultOutputFile?: string;
    /**
     * The topic name displayed on the details pageNote: This field may return null, indicating that no valid value can be obtained.
     */
    readonly showTopicName: string;
    readonly topicName: string;
}
/**
 * Use this data source to query detailed information of tdmq message
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const message = tencentcloud.Tdmq.getRocketmqMessages({
 *     clusterId: "rocketmq-rkrbm52djmro",
 *     environmentId: "keep_ns",
 *     msgId: "A9FE8D0567FE15DB97425FC08EEF0000",
 *     queryDlqMsg: false,
 *     topicName: "keep-topic",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRocketmqMessagesOutput(args: GetRocketmqMessagesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRocketmqMessagesResult> {
    return pulumi.output(args).apply((a: any) => getRocketmqMessages(a, opts))
}

/**
 * A collection of arguments for invoking getRocketmqMessages.
 */
export interface GetRocketmqMessagesOutputArgs {
    /**
     * Cluster id.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Environment.
     */
    environmentId: pulumi.Input<string>;
    /**
     * Message ID.
     */
    msgId: pulumi.Input<string>;
    /**
     * The value is true when querying dead letters, only valid for Rocketmq.
     */
    queryDlqMsg?: pulumi.Input<boolean>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Topic, groupId is passed when querying dead letters.
     */
    topicName: pulumi.Input<string>;
}
