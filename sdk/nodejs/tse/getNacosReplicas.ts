// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tse nacosReplicas
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const nacosReplicas = pulumi.output(tencentcloud.Tse.getNacosReplicas({
 *     instanceId: "ins-8078da86",
 * }));
 * ```
 */
export function getNacosReplicas(args: GetNacosReplicasArgs, opts?: pulumi.InvokeOptions): Promise<GetNacosReplicasResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Tse/getNacosReplicas:getNacosReplicas", {
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getNacosReplicas.
 */
export interface GetNacosReplicasArgs {
    /**
     * engine instance ID.
     */
    instanceId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getNacosReplicas.
 */
export interface GetNacosReplicasResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    /**
     * Engine instance replica information.
     */
    readonly replicas: outputs.Tse.GetNacosReplicasReplica[];
    readonly resultOutputFile?: string;
}

export function getNacosReplicasOutput(args: GetNacosReplicasOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNacosReplicasResult> {
    return pulumi.output(args).apply(a => getNacosReplicas(a, opts))
}

/**
 * A collection of arguments for invoking getNacosReplicas.
 */
export interface GetNacosReplicasOutputArgs {
    /**
     * engine instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
