// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query API gateway IP strategy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const service = new tencentcloud.apigateway.Service("service", {
 *     serviceName: "ck",
 *     protocol: "http&https",
 *     serviceDesc: "your nice service",
 *     netTypes: [
 *         "INNER",
 *         "OUTER",
 *     ],
 *     ipVersion: "IPv4",
 * });
 * const test = new tencentcloud.apigateway.IpStrategy("test", {
 *     serviceId: service.id,
 *     strategyName: "tf_test",
 *     strategyType: "BLACK",
 *     strategyData: "9.9.9.9",
 * });
 * const id = tencentcloud.ApiGateway.getIpStrategiesOutput({
 *     serviceId: test.serviceId,
 * });
 * const name = tencentcloud.ApiGateway.getIpStrategiesOutput({
 *     serviceId: test.serviceId,
 *     strategyName: test.strategyName,
 * });
 * ```
 */
export function getIpStrategies(args: GetIpStrategiesArgs, opts?: pulumi.InvokeOptions): Promise<GetIpStrategiesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:ApiGateway/getIpStrategies:getIpStrategies", {
        "resultOutputFile": args.resultOutputFile,
        "serviceId": args.serviceId,
        "strategyName": args.strategyName,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpStrategies.
 */
export interface GetIpStrategiesArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * The service ID to be queried.
     */
    serviceId: string;
    /**
     * Name of IP policy.
     */
    strategyName?: string;
}

/**
 * A collection of values returned by getIpStrategies.
 */
export interface GetIpStrategiesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of strategy.
     */
    readonly lists: outputs.ApiGateway.GetIpStrategiesList[];
    readonly resultOutputFile?: string;
    /**
     * The service ID.
     */
    readonly serviceId: string;
    /**
     * Name of the strategy.
     */
    readonly strategyName?: string;
}

export function getIpStrategiesOutput(args: GetIpStrategiesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIpStrategiesResult> {
    return pulumi.output(args).apply(a => getIpStrategies(a, opts))
}

/**
 * A collection of arguments for invoking getIpStrategies.
 */
export interface GetIpStrategiesOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * The service ID to be queried.
     */
    serviceId: pulumi.Input<string>;
    /**
     * Name of IP policy.
     */
    strategyName?: pulumi.Input<string>;
}
