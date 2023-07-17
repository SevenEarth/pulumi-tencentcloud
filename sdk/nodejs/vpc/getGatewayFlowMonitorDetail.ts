// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vpc gatewayFlowMonitorDetail
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const gatewayFlowMonitorDetail = pulumi.output(tencentcloud.Vpc.getGatewayFlowMonitorDetail({
 *     orderDirection: "DESC",
 *     orderField: "OutTraffic",
 *     timePoint: "2023-06-02 12:15:20",
 *     vpnId: "vpngw-gt8bianl",
 * }));
 * ```
 */
export function getGatewayFlowMonitorDetail(args: GetGatewayFlowMonitorDetailArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayFlowMonitorDetailResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Vpc/getGatewayFlowMonitorDetail:getGatewayFlowMonitorDetail", {
        "directConnectGatewayId": args.directConnectGatewayId,
        "natId": args.natId,
        "orderDirection": args.orderDirection,
        "orderField": args.orderField,
        "peeringConnectionId": args.peeringConnectionId,
        "resultOutputFile": args.resultOutputFile,
        "timePoint": args.timePoint,
        "vpnId": args.vpnId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGatewayFlowMonitorDetail.
 */
export interface GetGatewayFlowMonitorDetailArgs {
    /**
     * The instance ID of the Direct Connect gateway, such as `dcg-ltjahce6`.
     */
    directConnectGatewayId?: string;
    /**
     * The instance ID of the NAT gateway, such as `nat-ltjahce6`.
     */
    natId?: string;
    /**
     * Order methods. Ascending: `ASC`, Descending: `DESC`.
     */
    orderDirection?: string;
    /**
     * The order field supports `InPkg`, `OutPkg`, `InTraffic`, and `OutTraffic`.
     */
    orderField?: string;
    /**
     * The instance ID of the peering connection, such as `pcx-ltjahce6`.
     */
    peeringConnectionId?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * The point in time. This indicates details of this minute will be queried. For example, in `2019-02-28 18:15:20`, details at `18:15` will be queried.
     */
    timePoint: string;
    /**
     * The instance ID of the VPN gateway, such as `vpn-ltjahce6`.
     */
    vpnId?: string;
}

/**
 * A collection of values returned by getGatewayFlowMonitorDetail.
 */
export interface GetGatewayFlowMonitorDetailResult {
    readonly directConnectGatewayId?: string;
    /**
     * The gateway traffic monitoring details.
     */
    readonly gatewayFlowMonitorDetailSets: outputs.Vpc.GetGatewayFlowMonitorDetailGatewayFlowMonitorDetailSet[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly natId?: string;
    readonly orderDirection?: string;
    readonly orderField?: string;
    readonly peeringConnectionId?: string;
    readonly resultOutputFile?: string;
    readonly timePoint: string;
    readonly vpnId?: string;
}

export function getGatewayFlowMonitorDetailOutput(args: GetGatewayFlowMonitorDetailOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayFlowMonitorDetailResult> {
    return pulumi.output(args).apply(a => getGatewayFlowMonitorDetail(a, opts))
}

/**
 * A collection of arguments for invoking getGatewayFlowMonitorDetail.
 */
export interface GetGatewayFlowMonitorDetailOutputArgs {
    /**
     * The instance ID of the Direct Connect gateway, such as `dcg-ltjahce6`.
     */
    directConnectGatewayId?: pulumi.Input<string>;
    /**
     * The instance ID of the NAT gateway, such as `nat-ltjahce6`.
     */
    natId?: pulumi.Input<string>;
    /**
     * Order methods. Ascending: `ASC`, Descending: `DESC`.
     */
    orderDirection?: pulumi.Input<string>;
    /**
     * The order field supports `InPkg`, `OutPkg`, `InTraffic`, and `OutTraffic`.
     */
    orderField?: pulumi.Input<string>;
    /**
     * The instance ID of the peering connection, such as `pcx-ltjahce6`.
     */
    peeringConnectionId?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * The point in time. This indicates details of this minute will be queried. For example, in `2019-02-28 18:15:20`, details at `18:15` will be queried.
     */
    timePoint: pulumi.Input<string>;
    /**
     * The instance ID of the VPN gateway, such as `vpn-ltjahce6`.
     */
    vpnId?: pulumi.Input<string>;
}
