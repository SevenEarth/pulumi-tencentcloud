// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getVipInstance(args: GetVipInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetVipInstanceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Tdmq/getVipInstance:getVipInstance", {
        "clusterId": args.clusterId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getVipInstance.
 */
export interface GetVipInstanceArgs {
    clusterId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getVipInstance.
 */
export interface GetVipInstanceResult {
    readonly clusterId: string;
    readonly clusterInfos: outputs.Tdmq.GetVipInstanceClusterInfo[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceConfigs: outputs.Tdmq.GetVipInstanceInstanceConfig[];
    readonly resultOutputFile?: string;
}
export function getVipInstanceOutput(args: GetVipInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVipInstanceResult> {
    return pulumi.output(args).apply((a: any) => getVipInstance(a, opts))
}

/**
 * A collection of arguments for invoking getVipInstance.
 */
export interface GetVipInstanceOutputArgs {
    clusterId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
