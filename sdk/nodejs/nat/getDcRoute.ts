// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vpc natDcRoute
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const natDcRoute = pulumi.output(tencentcloud.Nat.getDcRoute({
 *     natGatewayId: "nat-gnxkey2e",
 *     vpcId: "vpc-pyyv5k3v",
 * }));
 * ```
 */
export function getDcRoute(args: GetDcRouteArgs, opts?: pulumi.InvokeOptions): Promise<GetDcRouteResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Nat/getDcRoute:getDcRoute", {
        "natGatewayId": args.natGatewayId,
        "resultOutputFile": args.resultOutputFile,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDcRoute.
 */
export interface GetDcRouteArgs {
    /**
     * Unique identifier of Nat Gateway.
     */
    natGatewayId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Unique identifier of Vpc.
     */
    vpcId: string;
}

/**
 * A collection of values returned by getDcRoute.
 */
export interface GetDcRouteResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Data of route.
     */
    readonly natDirectConnectGatewayRouteSets: outputs.Nat.GetDcRouteNatDirectConnectGatewayRouteSet[];
    readonly natGatewayId: string;
    readonly resultOutputFile?: string;
    readonly vpcId: string;
}

export function getDcRouteOutput(args: GetDcRouteOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDcRouteResult> {
    return pulumi.output(args).apply(a => getDcRoute(a, opts))
}

/**
 * A collection of arguments for invoking getDcRoute.
 */
export interface GetDcRouteOutputArgs {
    /**
     * Unique identifier of Nat Gateway.
     */
    natGatewayId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Unique identifier of Vpc.
     */
    vpcId: pulumi.Input<string>;
}