// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cynosdb proxyNode
 */
export function getProxyNode(args?: GetProxyNodeArgs, opts?: pulumi.InvokeOptions): Promise<GetProxyNodeResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cynosdb/getProxyNode:getProxyNode", {
        "filters": args.filters,
        "orderBy": args.orderBy,
        "orderByType": args.orderByType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getProxyNode.
 */
export interface GetProxyNodeArgs {
    /**
     * Search criteria, if there are multiple filters, the relationship between the filters is a logical AND relationship.
     */
    filters?: inputs.Cynosdb.GetProxyNodeFilter[];
    /**
     * Sort field, value range:CREATETIME: creation time; PRIODENDTIME: expiration time.
     */
    orderBy?: string;
    /**
     * Sort type, value range:ASC: ascending sort; DESC: descending sort.
     */
    orderByType?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getProxyNode.
 */
export interface GetProxyNodeResult {
    readonly filters?: outputs.Cynosdb.GetProxyNodeFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly orderBy?: string;
    readonly orderByType?: string;
    /**
     * Database Agent Node List.
     */
    readonly proxyNodeInfos: outputs.Cynosdb.GetProxyNodeProxyNodeInfo[];
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of cynosdb proxyNode
 */
export function getProxyNodeOutput(args?: GetProxyNodeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProxyNodeResult> {
    return pulumi.output(args).apply((a: any) => getProxyNode(a, opts))
}

/**
 * A collection of arguments for invoking getProxyNode.
 */
export interface GetProxyNodeOutputArgs {
    /**
     * Search criteria, if there are multiple filters, the relationship between the filters is a logical AND relationship.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.Cynosdb.GetProxyNodeFilterArgs>[]>;
    /**
     * Sort field, value range:CREATETIME: creation time; PRIODENDTIME: expiration time.
     */
    orderBy?: pulumi.Input<string>;
    /**
     * Sort type, value range:ASC: ascending sort; DESC: descending sort.
     */
    orderByType?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
