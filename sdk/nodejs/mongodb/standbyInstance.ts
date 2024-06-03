// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provide a resource to create a Mongodb standby instance.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const shanghai = new tencentcloud.Provider("shanghai", {region: "ap-shanghai"});
 * const mongodbInstance = new tencentcloud.mongodb.Instance("mongodbInstance", {
 *     instanceName: "tf-mongodb-test",
 *     memory: 4,
 *     volume: 100,
 *     engineVersion: "MONGO_40_WT",
 *     machineType: "HIO10G",
 *     availableZone: _var.availability_zone,
 *     projectId: 0,
 *     password: "test1234",
 *     tags: {
 *         test: "test",
 *     },
 * });
 * const mongodbStandbyInstance = new tencentcloud.mongodb.StandbyInstance("mongodbStandbyInstance", {
 *     instanceName: "tf-mongodb-standby-test",
 *     memory: 4,
 *     volume: 100,
 *     availableZone: "ap-shanghai-2",
 *     projectId: 0,
 *     fatherInstanceId: mongodbInstance.id,
 *     fatherInstanceRegion: "ap-guangzhou",
 *     tags: {
 *         test: "test",
 *     },
 * }, {
 *     provider: tencentcloud.shanghai,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Mongodb instance can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Mongodb/standbyInstance:StandbyInstance mongodb cmgo-41s6jwy4
 * ```
 */
export class StandbyInstance extends pulumi.CustomResource {
    /**
     * Get an existing StandbyInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StandbyInstanceState, opts?: pulumi.CustomResourceOptions): StandbyInstance {
        return new StandbyInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mongodb/standbyInstance:StandbyInstance';

    /**
     * Returns true if the given object is an instance of StandbyInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StandbyInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StandbyInstance.__pulumiType;
    }

    /**
     * Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
     */
    public readonly autoRenewFlag!: pulumi.Output<number | undefined>;
    /**
     * The available zone of the Mongodb standby instance. NOTE: must not be same with father instance's.
     */
    public readonly availableZone!: pulumi.Output<string>;
    /**
     * The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
     */
    public readonly chargeType!: pulumi.Output<string | undefined>;
    /**
     * Creation time of the Mongodb instance.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Version of the standby Mongodb instance and must be same as the version of main instance.
     */
    public /*out*/ readonly engineVersion!: pulumi.Output<string>;
    /**
     * Indicates the main instance ID of standby instances.
     */
    public readonly fatherInstanceId!: pulumi.Output<string>;
    /**
     * Indicates the region of main instance.
     */
    public readonly fatherInstanceRegion!: pulumi.Output<string>;
    /**
     * Name of the Mongodb instance.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * Type of standby Mongodb instance and must be same as the type of main instance.
     */
    public /*out*/ readonly machineType!: pulumi.Output<string>;
    /**
     * Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
     */
    public readonly memory!: pulumi.Output<number>;
    /**
     * The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when chargeType is set to `PREPAID`.
     */
    public readonly prepaidPeriod!: pulumi.Output<number | undefined>;
    /**
     * ID of the project which the instance belongs.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * ID of the security group.
     */
    public readonly securityGroups!: pulumi.Output<string[] | undefined>;
    /**
     * Status of the Mongodb instance, and available values include pending initialization(expressed with 0),  processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * ID of the subnet within this VPC. The value is required if `vpcId` is set.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * The tags of the Mongodb. Key name `project` is system reserved and can't be used.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * IP of the Mongodb instance.
     */
    public /*out*/ readonly vip!: pulumi.Output<string>;
    /**
     * Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
     */
    public readonly volume!: pulumi.Output<number>;
    /**
     * ID of the VPC.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;
    /**
     * IP port of the Mongodb instance.
     */
    public /*out*/ readonly vport!: pulumi.Output<number>;

    /**
     * Create a StandbyInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StandbyInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StandbyInstanceArgs | StandbyInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StandbyInstanceState | undefined;
            resourceInputs["autoRenewFlag"] = state ? state.autoRenewFlag : undefined;
            resourceInputs["availableZone"] = state ? state.availableZone : undefined;
            resourceInputs["chargeType"] = state ? state.chargeType : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["engineVersion"] = state ? state.engineVersion : undefined;
            resourceInputs["fatherInstanceId"] = state ? state.fatherInstanceId : undefined;
            resourceInputs["fatherInstanceRegion"] = state ? state.fatherInstanceRegion : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["machineType"] = state ? state.machineType : undefined;
            resourceInputs["memory"] = state ? state.memory : undefined;
            resourceInputs["prepaidPeriod"] = state ? state.prepaidPeriod : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vip"] = state ? state.vip : undefined;
            resourceInputs["volume"] = state ? state.volume : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vport"] = state ? state.vport : undefined;
        } else {
            const args = argsOrState as StandbyInstanceArgs | undefined;
            if ((!args || args.availableZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availableZone'");
            }
            if ((!args || args.fatherInstanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fatherInstanceId'");
            }
            if ((!args || args.fatherInstanceRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fatherInstanceRegion'");
            }
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            if ((!args || args.memory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memory'");
            }
            if ((!args || args.volume === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volume'");
            }
            resourceInputs["autoRenewFlag"] = args ? args.autoRenewFlag : undefined;
            resourceInputs["availableZone"] = args ? args.availableZone : undefined;
            resourceInputs["chargeType"] = args ? args.chargeType : undefined;
            resourceInputs["fatherInstanceId"] = args ? args.fatherInstanceId : undefined;
            resourceInputs["fatherInstanceRegion"] = args ? args.fatherInstanceRegion : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["prepaidPeriod"] = args ? args.prepaidPeriod : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["volume"] = args ? args.volume : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["engineVersion"] = undefined /*out*/;
            resourceInputs["machineType"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["vip"] = undefined /*out*/;
            resourceInputs["vport"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StandbyInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StandbyInstance resources.
 */
export interface StandbyInstanceState {
    /**
     * Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
     */
    autoRenewFlag?: pulumi.Input<number>;
    /**
     * The available zone of the Mongodb standby instance. NOTE: must not be same with father instance's.
     */
    availableZone?: pulumi.Input<string>;
    /**
     * The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Creation time of the Mongodb instance.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Version of the standby Mongodb instance and must be same as the version of main instance.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Indicates the main instance ID of standby instances.
     */
    fatherInstanceId?: pulumi.Input<string>;
    /**
     * Indicates the region of main instance.
     */
    fatherInstanceRegion?: pulumi.Input<string>;
    /**
     * Name of the Mongodb instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Type of standby Mongodb instance and must be same as the type of main instance.
     */
    machineType?: pulumi.Input<string>;
    /**
     * Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
     */
    memory?: pulumi.Input<number>;
    /**
     * The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when chargeType is set to `PREPAID`.
     */
    prepaidPeriod?: pulumi.Input<number>;
    /**
     * ID of the project which the instance belongs.
     */
    projectId?: pulumi.Input<number>;
    /**
     * ID of the security group.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Status of the Mongodb instance, and available values include pending initialization(expressed with 0),  processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).
     */
    status?: pulumi.Input<number>;
    /**
     * ID of the subnet within this VPC. The value is required if `vpcId` is set.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The tags of the Mongodb. Key name `project` is system reserved and can't be used.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * IP of the Mongodb instance.
     */
    vip?: pulumi.Input<string>;
    /**
     * Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
     */
    volume?: pulumi.Input<number>;
    /**
     * ID of the VPC.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * IP port of the Mongodb instance.
     */
    vport?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a StandbyInstance resource.
 */
export interface StandbyInstanceArgs {
    /**
     * Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
     */
    autoRenewFlag?: pulumi.Input<number>;
    /**
     * The available zone of the Mongodb standby instance. NOTE: must not be same with father instance's.
     */
    availableZone: pulumi.Input<string>;
    /**
     * The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Indicates the main instance ID of standby instances.
     */
    fatherInstanceId: pulumi.Input<string>;
    /**
     * Indicates the region of main instance.
     */
    fatherInstanceRegion: pulumi.Input<string>;
    /**
     * Name of the Mongodb instance.
     */
    instanceName: pulumi.Input<string>;
    /**
     * Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
     */
    memory: pulumi.Input<number>;
    /**
     * The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when chargeType is set to `PREPAID`.
     */
    prepaidPeriod?: pulumi.Input<number>;
    /**
     * ID of the project which the instance belongs.
     */
    projectId?: pulumi.Input<number>;
    /**
     * ID of the security group.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the subnet within this VPC. The value is required if `vpcId` is set.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The tags of the Mongodb. Key name `project` is system reserved and can't be used.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
     */
    volume: pulumi.Input<number>;
    /**
     * ID of the VPC.
     */
    vpcId?: pulumi.Input<string>;
}
