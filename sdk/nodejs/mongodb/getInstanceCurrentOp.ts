// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of mongodb instanceCurrentOp
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const instanceCurrentOp = tencentcloud.Mongodb.getInstanceCurrentOp({
 *     instanceId: "cmgo-b43i3wkj",
 *     op: "command",
 *     orderByType: "desc",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstanceCurrentOp(args: GetInstanceCurrentOpArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceCurrentOpResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Mongodb/getInstanceCurrentOp:getInstanceCurrentOp", {
        "instanceId": args.instanceId,
        "millisecondRunning": args.millisecondRunning,
        "ns": args.ns,
        "op": args.op,
        "orderBy": args.orderBy,
        "orderByType": args.orderByType,
        "replicaSetName": args.replicaSetName,
        "resultOutputFile": args.resultOutputFile,
        "state": args.state,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceCurrentOp.
 */
export interface GetInstanceCurrentOpArgs {
    /**
     * Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
     */
    instanceId: string;
    /**
     * Filter condition, the time that the operation has been executed (unit: millisecond),the result will return the operation that exceeds the set time, the default value is 0,and the value range is [0, 3600000].
     */
    millisecondRunning?: number;
    /**
     * Filter condition, the namespace namespace to which the operation belongs, in the format of db.collection.
     */
    ns?: string;
    /**
     * Filter condition, operation type, possible values: none, update, insert, query, command, getmore,remove and killcursors.
     */
    op?: string;
    /**
     * Returns the sorted field of the result set, currently supports: MicrosecsRunning/microsecsrunning,the default is ascending sort.
     */
    orderBy?: string;
    /**
     * Returns the sorting method of the result set, possible values: ASC/asc or DESC/desc.
     */
    orderByType?: string;
    /**
     * filter condition, shard name.
     */
    replicaSetName?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Filter condition, node status, possible value: primary, secondary.
     */
    state?: string;
}

/**
 * A collection of values returned by getInstanceCurrentOp.
 */
export interface GetInstanceCurrentOpResult {
    /**
     * current operation list.
     */
    readonly currentOps: outputs.Mongodb.GetInstanceCurrentOpCurrentOp[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly millisecondRunning?: number;
    /**
     * operation namespace.
     */
    readonly ns?: string;
    /**
     * operation value.
     */
    readonly op?: string;
    readonly orderBy?: string;
    readonly orderByType?: string;
    /**
     * Replication name.
     */
    readonly replicaSetName?: string;
    readonly resultOutputFile?: string;
    /**
     * operation state.
     */
    readonly state?: string;
}
/**
 * Use this data source to query detailed information of mongodb instanceCurrentOp
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const instanceCurrentOp = tencentcloud.Mongodb.getInstanceCurrentOp({
 *     instanceId: "cmgo-b43i3wkj",
 *     op: "command",
 *     orderByType: "desc",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstanceCurrentOpOutput(args: GetInstanceCurrentOpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceCurrentOpResult> {
    return pulumi.output(args).apply((a: any) => getInstanceCurrentOp(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceCurrentOp.
 */
export interface GetInstanceCurrentOpOutputArgs {
    /**
     * Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Filter condition, the time that the operation has been executed (unit: millisecond),the result will return the operation that exceeds the set time, the default value is 0,and the value range is [0, 3600000].
     */
    millisecondRunning?: pulumi.Input<number>;
    /**
     * Filter condition, the namespace namespace to which the operation belongs, in the format of db.collection.
     */
    ns?: pulumi.Input<string>;
    /**
     * Filter condition, operation type, possible values: none, update, insert, query, command, getmore,remove and killcursors.
     */
    op?: pulumi.Input<string>;
    /**
     * Returns the sorted field of the result set, currently supports: MicrosecsRunning/microsecsrunning,the default is ascending sort.
     */
    orderBy?: pulumi.Input<string>;
    /**
     * Returns the sorting method of the result set, possible values: ASC/asc or DESC/desc.
     */
    orderByType?: pulumi.Input<string>;
    /**
     * filter condition, shard name.
     */
    replicaSetName?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Filter condition, node status, possible value: primary, secondary.
     */
    state?: pulumi.Input<string>;
}
