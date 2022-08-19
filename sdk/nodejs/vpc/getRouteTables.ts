// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query vpc route tables information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const config = new pulumi.Config();
 * const availabilityZone = config.get("availabilityZone") || "ap-guangzhou-3";
 * const foo = new tencentcloud.vpc.Instance("foo", {cidrBlock: "10.0.0.0/16"});
 * const routeTable = new tencentcloud.route.Table("routeTable", {
 *     vpcId: foo.id,
 *     tags: {
 *         test: "test",
 *     },
 * });
 * const idInstances = tencentcloud.Vpc.getRouteTablesOutput({
 *     routeTableId: routeTable.id,
 * });
 * const nameInstances = tencentcloud.Vpc.getRouteTablesOutput({
 *     name: routeTable.name,
 * });
 * const vpcDefaultInstance = tencentcloud.Vpc.getRouteTablesOutput({
 *     vpcId: foo.id,
 *     associationMain: true,
 * });
 * const tagsInstances = routeTable.tags.apply(tags => tencentcloud.Vpc.getRouteTablesOutput({
 *     tags: tags,
 * }));
 * ```
 */
export function getRouteTables(args?: GetRouteTablesArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteTablesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Vpc/getRouteTables:getRouteTables", {
        "associationMain": args.associationMain,
        "name": args.name,
        "resultOutputFile": args.resultOutputFile,
        "routeTableId": args.routeTableId,
        "tagKey": args.tagKey,
        "tags": args.tags,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouteTables.
 */
export interface GetRouteTablesArgs {
    /**
     * Filter the main routing table.
     */
    associationMain?: boolean;
    /**
     * Name of the routing table to be queried.
     */
    name?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * ID of the routing table to be queried.
     */
    routeTableId?: string;
    /**
     * Filter if routing table has this tag.
     */
    tagKey?: string;
    /**
     * Tags of the routing table to be queried.
     */
    tags?: {[key: string]: any};
    /**
     * ID of the VPC to be queried.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getRouteTables.
 */
export interface GetRouteTablesResult {
    readonly associationMain?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The information list of the VPC route table.
     */
    readonly instanceLists: outputs.Vpc.GetRouteTablesInstanceList[];
    /**
     * Name of the routing table.
     */
    readonly name?: string;
    readonly resultOutputFile?: string;
    /**
     * ID of the routing table.
     */
    readonly routeTableId?: string;
    readonly tagKey?: string;
    /**
     * Tags of the routing table.
     */
    readonly tags?: {[key: string]: any};
    /**
     * ID of the VPC.
     */
    readonly vpcId?: string;
}

export function getRouteTablesOutput(args?: GetRouteTablesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouteTablesResult> {
    return pulumi.output(args).apply(a => getRouteTables(a, opts))
}

/**
 * A collection of arguments for invoking getRouteTables.
 */
export interface GetRouteTablesOutputArgs {
    /**
     * Filter the main routing table.
     */
    associationMain?: pulumi.Input<boolean>;
    /**
     * Name of the routing table to be queried.
     */
    name?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * ID of the routing table to be queried.
     */
    routeTableId?: pulumi.Input<string>;
    /**
     * Filter if routing table has this tag.
     */
    tagKey?: pulumi.Input<string>;
    /**
     * Tags of the routing table to be queried.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * ID of the VPC to be queried.
     */
    vpcId?: pulumi.Input<string>;
}
