// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provide a resource to create a CynosDB cluster.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.Cynosdb.Cluster("foo", {
 *     availableZone: "ap-guangzhou-4",
 *     clusterName: "tf-cynosdb",
 *     dbType: "MYSQL",
 *     dbVersion: "5.7",
 *     forceDelete: false,
 *     instanceCpuCore: 1,
 *     instanceMaintainDuration: 7200,
 *     instanceMaintainStartTime: 10800,
 *     instanceMaintainWeekdays: [
 *         "Fri",
 *         "Mon",
 *         "Sat",
 *         "Sun",
 *         "Thu",
 *         "Wed",
 *         "Tue",
 *     ],
 *     instanceMemorySize: 2,
 *     paramItem: [{
 *         currentValue: "utf8mb4",
 *         name: "character_set_server",
 *     }],
 *     password: "cynos@123",
 *     roGroupSgs: ["sg-ibyjkl6r"],
 *     rwGroupSgs: ["sg-ibyjkl6r"],
 *     storageLimit: 1000,
 *     subnetId: "subnet-q6fhy1mi",
 *     tags: {
 *         test: "test",
 *     },
 *     vpcId: "vpc-h70b6b49",
 * });
 * ```
 *
 * ## Import
 *
 * CynosDB cluster can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Cynosdb/cluster:Cluster foo cynosdbmysql-dzj5l8gz
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cynosdb/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * Auto renew flag. Valid values are `0`(MANUAL_RENEW), `1`(AUTO_RENEW). Default value is `0`. Only works for PREPAID cluster.
     */
    public readonly autoRenewFlag!: pulumi.Output<number | undefined>;
    /**
     * The available zone of the CynosDB Cluster.
     */
    public readonly availableZone!: pulumi.Output<string>;
    /**
     * The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`.
     */
    public readonly chargeType!: pulumi.Output<string | undefined>;
    /**
     * Charset used by CynosDB cluster.
     */
    public /*out*/ readonly charset!: pulumi.Output<string>;
    /**
     * Name of CynosDB cluster.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * Status of the Cynosdb cluster.
     */
    public /*out*/ readonly clusterStatus!: pulumi.Output<string>;
    /**
     * Creation time of the CynosDB cluster.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Type of CynosDB, and available values include `MYSQL`.
     */
    public readonly dbType!: pulumi.Output<string>;
    /**
     * Version of CynosDB, which is related to `dbType`. For `MYSQL`, available value is `5.7`.
     */
    public readonly dbVersion!: pulumi.Output<string>;
    /**
     * Indicate whether to delete cluster instance directly or not. Default is false. If set true, the cluster and its `All RELATED INSTANCES` will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
     */
    public readonly forceDelete!: pulumi.Output<boolean | undefined>;
    /**
     * The number of CPU cores of read-write type instance in the CynosDB cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
     */
    public readonly instanceCpuCore!: pulumi.Output<number>;
    /**
     * ID of instance.
     */
    public /*out*/ readonly instanceId!: pulumi.Output<string>;
    /**
     * Duration time for maintenance, unit in second. `3600` by default.
     */
    public readonly instanceMaintainDuration!: pulumi.Output<number | undefined>;
    /**
     * Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
     */
    public readonly instanceMaintainStartTime!: pulumi.Output<number | undefined>;
    /**
     * Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
     */
    public readonly instanceMaintainWeekdays!: pulumi.Output<string[]>;
    /**
     * Memory capacity of read-write type instance, unit in GB. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
     */
    public readonly instanceMemorySize!: pulumi.Output<number>;
    /**
     * Name of instance.
     */
    public /*out*/ readonly instanceName!: pulumi.Output<string>;
    /**
     * Status of the instance.
     */
    public /*out*/ readonly instanceStatus!: pulumi.Output<string>;
    /**
     * Storage size of the instance, unit in GB.
     */
    public /*out*/ readonly instanceStorageSize!: pulumi.Output<number>;
    /**
     * Specify parameter list of database. Use `data.tencentcloud_mysql_default_params` to query available parameter details.
     */
    public readonly paramItems!: pulumi.Output<outputs.Cynosdb.ClusterParamItem[] | undefined>;
    /**
     * Password of `root` account.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Port of CynosDB cluster.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The tenancy (time unit is month) of the prepaid instance. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. NOTE: it only works when chargeType is set to `PREPAID`.
     */
    public readonly prepaidPeriod!: pulumi.Output<number | undefined>;
    /**
     * ID of the project. `0` by default.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * Readonly addresses. Each element contains the following attributes:
     */
    public /*out*/ readonly roGroupAddrs!: pulumi.Output<outputs.Cynosdb.ClusterRoGroupAddr[]>;
    /**
     * ID of read-only instance group.
     */
    public /*out*/ readonly roGroupId!: pulumi.Output<string>;
    /**
     * List of instances in the read-only instance group.
     */
    public /*out*/ readonly roGroupInstances!: pulumi.Output<outputs.Cynosdb.ClusterRoGroupInstance[]>;
    /**
     * IDs of security group for `roGroup`.
     */
    public readonly roGroupSgs!: pulumi.Output<string[] | undefined>;
    /**
     * Read-write addresses. Each element contains the following attributes:
     */
    public /*out*/ readonly rwGroupAddrs!: pulumi.Output<outputs.Cynosdb.ClusterRwGroupAddr[]>;
    /**
     * ID of read-write instance group.
     */
    public /*out*/ readonly rwGroupId!: pulumi.Output<string>;
    /**
     * List of instances in the read-write instance group.
     */
    public /*out*/ readonly rwGroupInstances!: pulumi.Output<outputs.Cynosdb.ClusterRwGroupInstance[]>;
    /**
     * IDs of security group for `rwGroup`.
     */
    public readonly rwGroupSgs!: pulumi.Output<string[] | undefined>;
    /**
     * Storage limit of CynosDB cluster instance, unit in GB. The maximum storage of a non-serverless instance in GB. NOTE: If dbType is `MYSQL` and chargeType is `PREPAID`, the value cannot exceed the maximum storage corresponding to the CPU and memory specifications, when chargeType is `POSTPAID_BY_HOUR`, this argument is unnecessary.
     */
    public readonly storageLimit!: pulumi.Output<number | undefined>;
    /**
     * Used storage of CynosDB cluster, unit in MB.
     */
    public /*out*/ readonly storageUsed!: pulumi.Output<number>;
    /**
     * ID of the subnet within this VPC.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * The tags of the CynosDB cluster.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * ID of the VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            resourceInputs["autoRenewFlag"] = state ? state.autoRenewFlag : undefined;
            resourceInputs["availableZone"] = state ? state.availableZone : undefined;
            resourceInputs["chargeType"] = state ? state.chargeType : undefined;
            resourceInputs["charset"] = state ? state.charset : undefined;
            resourceInputs["clusterName"] = state ? state.clusterName : undefined;
            resourceInputs["clusterStatus"] = state ? state.clusterStatus : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["dbType"] = state ? state.dbType : undefined;
            resourceInputs["dbVersion"] = state ? state.dbVersion : undefined;
            resourceInputs["forceDelete"] = state ? state.forceDelete : undefined;
            resourceInputs["instanceCpuCore"] = state ? state.instanceCpuCore : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["instanceMaintainDuration"] = state ? state.instanceMaintainDuration : undefined;
            resourceInputs["instanceMaintainStartTime"] = state ? state.instanceMaintainStartTime : undefined;
            resourceInputs["instanceMaintainWeekdays"] = state ? state.instanceMaintainWeekdays : undefined;
            resourceInputs["instanceMemorySize"] = state ? state.instanceMemorySize : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["instanceStatus"] = state ? state.instanceStatus : undefined;
            resourceInputs["instanceStorageSize"] = state ? state.instanceStorageSize : undefined;
            resourceInputs["paramItems"] = state ? state.paramItems : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["prepaidPeriod"] = state ? state.prepaidPeriod : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["roGroupAddrs"] = state ? state.roGroupAddrs : undefined;
            resourceInputs["roGroupId"] = state ? state.roGroupId : undefined;
            resourceInputs["roGroupInstances"] = state ? state.roGroupInstances : undefined;
            resourceInputs["roGroupSgs"] = state ? state.roGroupSgs : undefined;
            resourceInputs["rwGroupAddrs"] = state ? state.rwGroupAddrs : undefined;
            resourceInputs["rwGroupId"] = state ? state.rwGroupId : undefined;
            resourceInputs["rwGroupInstances"] = state ? state.rwGroupInstances : undefined;
            resourceInputs["rwGroupSgs"] = state ? state.rwGroupSgs : undefined;
            resourceInputs["storageLimit"] = state ? state.storageLimit : undefined;
            resourceInputs["storageUsed"] = state ? state.storageUsed : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.availableZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availableZone'");
            }
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            if ((!args || args.dbType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbType'");
            }
            if ((!args || args.dbVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbVersion'");
            }
            if ((!args || args.instanceCpuCore === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceCpuCore'");
            }
            if ((!args || args.instanceMemorySize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceMemorySize'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["autoRenewFlag"] = args ? args.autoRenewFlag : undefined;
            resourceInputs["availableZone"] = args ? args.availableZone : undefined;
            resourceInputs["chargeType"] = args ? args.chargeType : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["dbType"] = args ? args.dbType : undefined;
            resourceInputs["dbVersion"] = args ? args.dbVersion : undefined;
            resourceInputs["forceDelete"] = args ? args.forceDelete : undefined;
            resourceInputs["instanceCpuCore"] = args ? args.instanceCpuCore : undefined;
            resourceInputs["instanceMaintainDuration"] = args ? args.instanceMaintainDuration : undefined;
            resourceInputs["instanceMaintainStartTime"] = args ? args.instanceMaintainStartTime : undefined;
            resourceInputs["instanceMaintainWeekdays"] = args ? args.instanceMaintainWeekdays : undefined;
            resourceInputs["instanceMemorySize"] = args ? args.instanceMemorySize : undefined;
            resourceInputs["paramItems"] = args ? args.paramItems : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["prepaidPeriod"] = args ? args.prepaidPeriod : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["roGroupSgs"] = args ? args.roGroupSgs : undefined;
            resourceInputs["rwGroupSgs"] = args ? args.rwGroupSgs : undefined;
            resourceInputs["storageLimit"] = args ? args.storageLimit : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["charset"] = undefined /*out*/;
            resourceInputs["clusterStatus"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["instanceName"] = undefined /*out*/;
            resourceInputs["instanceStatus"] = undefined /*out*/;
            resourceInputs["instanceStorageSize"] = undefined /*out*/;
            resourceInputs["roGroupAddrs"] = undefined /*out*/;
            resourceInputs["roGroupId"] = undefined /*out*/;
            resourceInputs["roGroupInstances"] = undefined /*out*/;
            resourceInputs["rwGroupAddrs"] = undefined /*out*/;
            resourceInputs["rwGroupId"] = undefined /*out*/;
            resourceInputs["rwGroupInstances"] = undefined /*out*/;
            resourceInputs["storageUsed"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * Auto renew flag. Valid values are `0`(MANUAL_RENEW), `1`(AUTO_RENEW). Default value is `0`. Only works for PREPAID cluster.
     */
    autoRenewFlag?: pulumi.Input<number>;
    /**
     * The available zone of the CynosDB Cluster.
     */
    availableZone?: pulumi.Input<string>;
    /**
     * The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Charset used by CynosDB cluster.
     */
    charset?: pulumi.Input<string>;
    /**
     * Name of CynosDB cluster.
     */
    clusterName?: pulumi.Input<string>;
    /**
     * Status of the Cynosdb cluster.
     */
    clusterStatus?: pulumi.Input<string>;
    /**
     * Creation time of the CynosDB cluster.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Type of CynosDB, and available values include `MYSQL`.
     */
    dbType?: pulumi.Input<string>;
    /**
     * Version of CynosDB, which is related to `dbType`. For `MYSQL`, available value is `5.7`.
     */
    dbVersion?: pulumi.Input<string>;
    /**
     * Indicate whether to delete cluster instance directly or not. Default is false. If set true, the cluster and its `All RELATED INSTANCES` will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * The number of CPU cores of read-write type instance in the CynosDB cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
     */
    instanceCpuCore?: pulumi.Input<number>;
    /**
     * ID of instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Duration time for maintenance, unit in second. `3600` by default.
     */
    instanceMaintainDuration?: pulumi.Input<number>;
    /**
     * Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
     */
    instanceMaintainStartTime?: pulumi.Input<number>;
    /**
     * Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
     */
    instanceMaintainWeekdays?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Memory capacity of read-write type instance, unit in GB. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
     */
    instanceMemorySize?: pulumi.Input<number>;
    /**
     * Name of instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Status of the instance.
     */
    instanceStatus?: pulumi.Input<string>;
    /**
     * Storage size of the instance, unit in GB.
     */
    instanceStorageSize?: pulumi.Input<number>;
    /**
     * Specify parameter list of database. Use `data.tencentcloud_mysql_default_params` to query available parameter details.
     */
    paramItems?: pulumi.Input<pulumi.Input<inputs.Cynosdb.ClusterParamItem>[]>;
    /**
     * Password of `root` account.
     */
    password?: pulumi.Input<string>;
    /**
     * Port of CynosDB cluster.
     */
    port?: pulumi.Input<number>;
    /**
     * The tenancy (time unit is month) of the prepaid instance. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. NOTE: it only works when chargeType is set to `PREPAID`.
     */
    prepaidPeriod?: pulumi.Input<number>;
    /**
     * ID of the project. `0` by default.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Readonly addresses. Each element contains the following attributes:
     */
    roGroupAddrs?: pulumi.Input<pulumi.Input<inputs.Cynosdb.ClusterRoGroupAddr>[]>;
    /**
     * ID of read-only instance group.
     */
    roGroupId?: pulumi.Input<string>;
    /**
     * List of instances in the read-only instance group.
     */
    roGroupInstances?: pulumi.Input<pulumi.Input<inputs.Cynosdb.ClusterRoGroupInstance>[]>;
    /**
     * IDs of security group for `roGroup`.
     */
    roGroupSgs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Read-write addresses. Each element contains the following attributes:
     */
    rwGroupAddrs?: pulumi.Input<pulumi.Input<inputs.Cynosdb.ClusterRwGroupAddr>[]>;
    /**
     * ID of read-write instance group.
     */
    rwGroupId?: pulumi.Input<string>;
    /**
     * List of instances in the read-write instance group.
     */
    rwGroupInstances?: pulumi.Input<pulumi.Input<inputs.Cynosdb.ClusterRwGroupInstance>[]>;
    /**
     * IDs of security group for `rwGroup`.
     */
    rwGroupSgs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Storage limit of CynosDB cluster instance, unit in GB. The maximum storage of a non-serverless instance in GB. NOTE: If dbType is `MYSQL` and chargeType is `PREPAID`, the value cannot exceed the maximum storage corresponding to the CPU and memory specifications, when chargeType is `POSTPAID_BY_HOUR`, this argument is unnecessary.
     */
    storageLimit?: pulumi.Input<number>;
    /**
     * Used storage of CynosDB cluster, unit in MB.
     */
    storageUsed?: pulumi.Input<number>;
    /**
     * ID of the subnet within this VPC.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The tags of the CynosDB cluster.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * ID of the VPC.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Auto renew flag. Valid values are `0`(MANUAL_RENEW), `1`(AUTO_RENEW). Default value is `0`. Only works for PREPAID cluster.
     */
    autoRenewFlag?: pulumi.Input<number>;
    /**
     * The available zone of the CynosDB Cluster.
     */
    availableZone: pulumi.Input<string>;
    /**
     * The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Name of CynosDB cluster.
     */
    clusterName: pulumi.Input<string>;
    /**
     * Type of CynosDB, and available values include `MYSQL`.
     */
    dbType: pulumi.Input<string>;
    /**
     * Version of CynosDB, which is related to `dbType`. For `MYSQL`, available value is `5.7`.
     */
    dbVersion: pulumi.Input<string>;
    /**
     * Indicate whether to delete cluster instance directly or not. Default is false. If set true, the cluster and its `All RELATED INSTANCES` will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * The number of CPU cores of read-write type instance in the CynosDB cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
     */
    instanceCpuCore: pulumi.Input<number>;
    /**
     * Duration time for maintenance, unit in second. `3600` by default.
     */
    instanceMaintainDuration?: pulumi.Input<number>;
    /**
     * Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
     */
    instanceMaintainStartTime?: pulumi.Input<number>;
    /**
     * Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
     */
    instanceMaintainWeekdays?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Memory capacity of read-write type instance, unit in GB. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
     */
    instanceMemorySize: pulumi.Input<number>;
    /**
     * Specify parameter list of database. Use `data.tencentcloud_mysql_default_params` to query available parameter details.
     */
    paramItems?: pulumi.Input<pulumi.Input<inputs.Cynosdb.ClusterParamItem>[]>;
    /**
     * Password of `root` account.
     */
    password: pulumi.Input<string>;
    /**
     * Port of CynosDB cluster.
     */
    port?: pulumi.Input<number>;
    /**
     * The tenancy (time unit is month) of the prepaid instance. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. NOTE: it only works when chargeType is set to `PREPAID`.
     */
    prepaidPeriod?: pulumi.Input<number>;
    /**
     * ID of the project. `0` by default.
     */
    projectId?: pulumi.Input<number>;
    /**
     * IDs of security group for `roGroup`.
     */
    roGroupSgs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * IDs of security group for `rwGroup`.
     */
    rwGroupSgs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Storage limit of CynosDB cluster instance, unit in GB. The maximum storage of a non-serverless instance in GB. NOTE: If dbType is `MYSQL` and chargeType is `PREPAID`, the value cannot exceed the maximum storage corresponding to the CPU and memory specifications, when chargeType is `POSTPAID_BY_HOUR`, this argument is unnecessary.
     */
    storageLimit?: pulumi.Input<number>;
    /**
     * ID of the subnet within this VPC.
     */
    subnetId: pulumi.Input<string>;
    /**
     * The tags of the CynosDB cluster.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * ID of the VPC.
     */
    vpcId: pulumi.Input<string>;
}
