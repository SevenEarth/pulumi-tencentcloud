// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tsf deliveryConfigByGroupId
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const deliveryConfigByGroupId = pulumi.output(tencentcloud.Tsf.getDeliveryConfigByGroupId({
 *     groupId: "group-yrjkln9v",
 * }));
 * ```
 */
export function getDeliveryConfigByGroupId(args: GetDeliveryConfigByGroupIdArgs, opts?: pulumi.InvokeOptions): Promise<GetDeliveryConfigByGroupIdResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Tsf/getDeliveryConfigByGroupId:getDeliveryConfigByGroupId", {
        "groupId": args.groupId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDeliveryConfigByGroupId.
 */
export interface GetDeliveryConfigByGroupIdArgs {
    /**
     * groupId.
     */
    groupId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDeliveryConfigByGroupId.
 */
export interface GetDeliveryConfigByGroupIdResult {
    readonly groupId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * configuration item for deliver to a Kafka.
     */
    readonly results: outputs.Tsf.GetDeliveryConfigByGroupIdResult[];
}

export function getDeliveryConfigByGroupIdOutput(args: GetDeliveryConfigByGroupIdOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeliveryConfigByGroupIdResult> {
    return pulumi.output(args).apply(a => getDeliveryConfigByGroupId(a, opts))
}

/**
 * A collection of arguments for invoking getDeliveryConfigByGroupId.
 */
export interface GetDeliveryConfigByGroupIdOutputArgs {
    /**
     * groupId.
     */
    groupId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
