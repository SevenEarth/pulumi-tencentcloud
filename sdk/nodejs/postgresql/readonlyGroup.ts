// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this resource to create postgresql readonly group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const group = new tencentcloud.Postgresql.ReadonlyGroup("group", {
 *     masterDbInstanceId: "postgres-gzg9jb2n",
 *     maxReplayLag: 100,
 *     maxReplayLatency: 512,
 *     minDelayEliminateReserve: 1,
 *     projectId: 0,
 *     replayLagEliminate: 1,
 *     replayLatencyEliminate: 1,
 *     subnetId: "subnet-enm92y0m",
 *     vpcId: "vpc-86v957zb",
 * });
 * ```
 */
export class ReadonlyGroup extends pulumi.CustomResource {
    /**
     * Get an existing ReadonlyGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReadonlyGroupState, opts?: pulumi.CustomResourceOptions): ReadonlyGroup {
        return new ReadonlyGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Postgresql/readonlyGroup:ReadonlyGroup';

    /**
     * Returns true if the given object is an instance of ReadonlyGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReadonlyGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReadonlyGroup.__pulumiType;
    }

    /**
     * Create time of the postgresql instance.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Primary instance ID.
     */
    public readonly masterDbInstanceId!: pulumi.Output<string>;
    /**
     * Delay threshold in ms.
     */
    public readonly maxReplayLag!: pulumi.Output<number>;
    /**
     * Delayed log size threshold in MB.
     */
    public readonly maxReplayLatency!: pulumi.Output<number>;
    /**
     * The minimum number of read-only replicas that must be retained in an RO group.
     */
    public readonly minDelayEliminateReserve!: pulumi.Output<number>;
    /**
     * RO group name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Project ID.
     */
    public readonly projectId!: pulumi.Output<number>;
    /**
     * Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
     */
    public readonly replayLagEliminate!: pulumi.Output<number>;
    /**
     * Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
     */
    public readonly replayLatencyEliminate!: pulumi.Output<number>;
    /**
     * ID of security group. If both vpcId and subnetId are not set, this argument should not be set either.
     */
    public readonly securityGroupsIds!: pulumi.Output<string[] | undefined>;
    /**
     * VPC subnet ID.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * VPC ID.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a ReadonlyGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReadonlyGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReadonlyGroupArgs | ReadonlyGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReadonlyGroupState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["masterDbInstanceId"] = state ? state.masterDbInstanceId : undefined;
            resourceInputs["maxReplayLag"] = state ? state.maxReplayLag : undefined;
            resourceInputs["maxReplayLatency"] = state ? state.maxReplayLatency : undefined;
            resourceInputs["minDelayEliminateReserve"] = state ? state.minDelayEliminateReserve : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["replayLagEliminate"] = state ? state.replayLagEliminate : undefined;
            resourceInputs["replayLatencyEliminate"] = state ? state.replayLatencyEliminate : undefined;
            resourceInputs["securityGroupsIds"] = state ? state.securityGroupsIds : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as ReadonlyGroupArgs | undefined;
            if ((!args || args.masterDbInstanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'masterDbInstanceId'");
            }
            if ((!args || args.maxReplayLag === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxReplayLag'");
            }
            if ((!args || args.maxReplayLatency === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxReplayLatency'");
            }
            if ((!args || args.minDelayEliminateReserve === undefined) && !opts.urn) {
                throw new Error("Missing required property 'minDelayEliminateReserve'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.replayLagEliminate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'replayLagEliminate'");
            }
            if ((!args || args.replayLatencyEliminate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'replayLatencyEliminate'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["masterDbInstanceId"] = args ? args.masterDbInstanceId : undefined;
            resourceInputs["maxReplayLag"] = args ? args.maxReplayLag : undefined;
            resourceInputs["maxReplayLatency"] = args ? args.maxReplayLatency : undefined;
            resourceInputs["minDelayEliminateReserve"] = args ? args.minDelayEliminateReserve : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["replayLagEliminate"] = args ? args.replayLagEliminate : undefined;
            resourceInputs["replayLatencyEliminate"] = args ? args.replayLatencyEliminate : undefined;
            resourceInputs["securityGroupsIds"] = args ? args.securityGroupsIds : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReadonlyGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReadonlyGroup resources.
 */
export interface ReadonlyGroupState {
    /**
     * Create time of the postgresql instance.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Primary instance ID.
     */
    masterDbInstanceId?: pulumi.Input<string>;
    /**
     * Delay threshold in ms.
     */
    maxReplayLag?: pulumi.Input<number>;
    /**
     * Delayed log size threshold in MB.
     */
    maxReplayLatency?: pulumi.Input<number>;
    /**
     * The minimum number of read-only replicas that must be retained in an RO group.
     */
    minDelayEliminateReserve?: pulumi.Input<number>;
    /**
     * RO group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Project ID.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
     */
    replayLagEliminate?: pulumi.Input<number>;
    /**
     * Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
     */
    replayLatencyEliminate?: pulumi.Input<number>;
    /**
     * ID of security group. If both vpcId and subnetId are not set, this argument should not be set either.
     */
    securityGroupsIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * VPC subnet ID.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * VPC ID.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReadonlyGroup resource.
 */
export interface ReadonlyGroupArgs {
    /**
     * Primary instance ID.
     */
    masterDbInstanceId: pulumi.Input<string>;
    /**
     * Delay threshold in ms.
     */
    maxReplayLag: pulumi.Input<number>;
    /**
     * Delayed log size threshold in MB.
     */
    maxReplayLatency: pulumi.Input<number>;
    /**
     * The minimum number of read-only replicas that must be retained in an RO group.
     */
    minDelayEliminateReserve: pulumi.Input<number>;
    /**
     * RO group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Project ID.
     */
    projectId: pulumi.Input<number>;
    /**
     * Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
     */
    replayLagEliminate: pulumi.Input<number>;
    /**
     * Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
     */
    replayLatencyEliminate: pulumi.Input<number>;
    /**
     * ID of security group. If both vpcId and subnetId are not set, this argument should not be set either.
     */
    securityGroupsIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * VPC subnet ID.
     */
    subnetId: pulumi.Input<string>;
    /**
     * VPC ID.
     */
    vpcId: pulumi.Input<string>;
}
