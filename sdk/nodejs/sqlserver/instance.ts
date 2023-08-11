// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this resource to create SQL Server instance
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const zones = tencentcloud.Availability.getZonesByProduct({
 *     product: "sqlserver",
 * });
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     availabilityZone: zones.then(zones => zones.zones?[4]?.name),
 *     vpcId: vpc.id,
 *     cidrBlock: "10.0.0.0/16",
 *     isMulticast: false,
 * });
 * const example = new tencentcloud.sqlserver.Instance("example", {
 *     availabilityZone: zones.then(zones => zones.zones?[4]?.name),
 *     chargeType: "POSTPAID_BY_HOUR",
 *     vpcId: vpc.id,
 *     subnetId: subnet.id,
 *     projectId: 0,
 *     memory: 16,
 *     storage: 100,
 * });
 * ```
 *
 * ## Import
 *
 * SQL Server instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Sqlserver/instance:Instance example mssql-3cdq7kx5
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Sqlserver/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal (Default). Only valid when purchasing a prepaid instance.
     */
    public readonly autoRenew!: pulumi.Output<number | undefined>;
    /**
     * Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
     */
    public readonly autoVoucher!: pulumi.Output<number | undefined>;
    /**
     * Availability zone.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * Pay type of the SQL Server instance. Available values `PREPAID`, `POSTPAID_BY_HOUR`.
     */
    public readonly chargeType!: pulumi.Output<string | undefined>;
    /**
     * Create time of the SQL Server instance.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Version of the SQL Server database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
     */
    public readonly engineVersion!: pulumi.Output<string | undefined>;
    /**
     * Instance type. `DUAL` (dual-server high availability), `CLUSTER` (cluster). Default is `DUAL`.
     */
    public readonly haType!: pulumi.Output<string | undefined>;
    /**
     * Start time of the maintenance in one day, format like `HH:mm`.
     */
    public readonly maintenanceStartTime!: pulumi.Output<string>;
    /**
     * The timespan of maintenance in one day, unit is hour.
     */
    public readonly maintenanceTimeSpan!: pulumi.Output<number>;
    /**
     * A list of integer indicates weekly maintenance. For example, [2,7] presents do weekly maintenance on every Tuesday and Sunday.
     */
    public readonly maintenanceWeekSets!: pulumi.Output<number[]>;
    /**
     * Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloudSqlserverSpecinfos` provides.
     */
    public readonly memory!: pulumi.Output<number>;
    /**
     * Indicate whether to deploy across availability zones.
     */
    public readonly multiZones!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the SQL Server instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Purchase instance period in month. The value does not exceed 48.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Project ID, default value is 0.
     */
    public readonly projectId!: pulumi.Output<number>;
    /**
     * Readonly flag. `RO` (read-only instance), `MASTER` (primary instance with read-only instances). If it is left empty, it refers to an instance which is not read-only and has no RO group.
     */
    public /*out*/ readonly roFlag!: pulumi.Output<string>;
    /**
     * Security group bound to the instance.
     */
    public readonly securityGroups!: pulumi.Output<string[] | undefined>;
    /**
     * Status of the SQL Server instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storageMin` and `storageMax` which data source `tencentcloudSqlserverSpecinfos` provides.
     */
    public readonly storage!: pulumi.Output<number>;
    /**
     * ID of subnet.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * The tags of the SQL Server.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * IP for private access.
     */
    public /*out*/ readonly vip!: pulumi.Output<string>;
    /**
     * An array of voucher IDs, currently only one can be used for a single order.
     */
    public readonly voucherIds!: pulumi.Output<string[] | undefined>;
    /**
     * ID of VPC.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;
    /**
     * Port for private access.
     */
    public /*out*/ readonly vport!: pulumi.Output<number>;
    /**
     * It has been deprecated from version 1.81.2. The way to execute the allocation. Supported values include: 0 - execute immediately, 1 - execute in maintenance window.
     *
     * @deprecated It has been deprecated from version 1.81.2.
     */
    public readonly waitSwitch!: pulumi.Output<number | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["autoVoucher"] = state ? state.autoVoucher : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["chargeType"] = state ? state.chargeType : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["engineVersion"] = state ? state.engineVersion : undefined;
            resourceInputs["haType"] = state ? state.haType : undefined;
            resourceInputs["maintenanceStartTime"] = state ? state.maintenanceStartTime : undefined;
            resourceInputs["maintenanceTimeSpan"] = state ? state.maintenanceTimeSpan : undefined;
            resourceInputs["maintenanceWeekSets"] = state ? state.maintenanceWeekSets : undefined;
            resourceInputs["memory"] = state ? state.memory : undefined;
            resourceInputs["multiZones"] = state ? state.multiZones : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["roFlag"] = state ? state.roFlag : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["storage"] = state ? state.storage : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vip"] = state ? state.vip : undefined;
            resourceInputs["voucherIds"] = state ? state.voucherIds : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vport"] = state ? state.vport : undefined;
            resourceInputs["waitSwitch"] = state ? state.waitSwitch : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.memory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memory'");
            }
            if ((!args || args.storage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storage'");
            }
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["autoVoucher"] = args ? args.autoVoucher : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["chargeType"] = args ? args.chargeType : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["haType"] = args ? args.haType : undefined;
            resourceInputs["maintenanceStartTime"] = args ? args.maintenanceStartTime : undefined;
            resourceInputs["maintenanceTimeSpan"] = args ? args.maintenanceTimeSpan : undefined;
            resourceInputs["maintenanceWeekSets"] = args ? args.maintenanceWeekSets : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["multiZones"] = args ? args.multiZones : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["storage"] = args ? args.storage : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["voucherIds"] = args ? args.voucherIds : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["waitSwitch"] = args ? args.waitSwitch : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["roFlag"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["vip"] = undefined /*out*/;
            resourceInputs["vport"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal (Default). Only valid when purchasing a prepaid instance.
     */
    autoRenew?: pulumi.Input<number>;
    /**
     * Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
     */
    autoVoucher?: pulumi.Input<number>;
    /**
     * Availability zone.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Pay type of the SQL Server instance. Available values `PREPAID`, `POSTPAID_BY_HOUR`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Create time of the SQL Server instance.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Version of the SQL Server database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Instance type. `DUAL` (dual-server high availability), `CLUSTER` (cluster). Default is `DUAL`.
     */
    haType?: pulumi.Input<string>;
    /**
     * Start time of the maintenance in one day, format like `HH:mm`.
     */
    maintenanceStartTime?: pulumi.Input<string>;
    /**
     * The timespan of maintenance in one day, unit is hour.
     */
    maintenanceTimeSpan?: pulumi.Input<number>;
    /**
     * A list of integer indicates weekly maintenance. For example, [2,7] presents do weekly maintenance on every Tuesday and Sunday.
     */
    maintenanceWeekSets?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloudSqlserverSpecinfos` provides.
     */
    memory?: pulumi.Input<number>;
    /**
     * Indicate whether to deploy across availability zones.
     */
    multiZones?: pulumi.Input<boolean>;
    /**
     * Name of the SQL Server instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Purchase instance period in month. The value does not exceed 48.
     */
    period?: pulumi.Input<number>;
    /**
     * Project ID, default value is 0.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Readonly flag. `RO` (read-only instance), `MASTER` (primary instance with read-only instances). If it is left empty, it refers to an instance which is not read-only and has no RO group.
     */
    roFlag?: pulumi.Input<string>;
    /**
     * Security group bound to the instance.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Status of the SQL Server instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.
     */
    status?: pulumi.Input<number>;
    /**
     * Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storageMin` and `storageMax` which data source `tencentcloudSqlserverSpecinfos` provides.
     */
    storage?: pulumi.Input<number>;
    /**
     * ID of subnet.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The tags of the SQL Server.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * IP for private access.
     */
    vip?: pulumi.Input<string>;
    /**
     * An array of voucher IDs, currently only one can be used for a single order.
     */
    voucherIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of VPC.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Port for private access.
     */
    vport?: pulumi.Input<number>;
    /**
     * It has been deprecated from version 1.81.2. The way to execute the allocation. Supported values include: 0 - execute immediately, 1 - execute in maintenance window.
     *
     * @deprecated It has been deprecated from version 1.81.2.
     */
    waitSwitch?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal (Default). Only valid when purchasing a prepaid instance.
     */
    autoRenew?: pulumi.Input<number>;
    /**
     * Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
     */
    autoVoucher?: pulumi.Input<number>;
    /**
     * Availability zone.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Pay type of the SQL Server instance. Available values `PREPAID`, `POSTPAID_BY_HOUR`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Version of the SQL Server database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Instance type. `DUAL` (dual-server high availability), `CLUSTER` (cluster). Default is `DUAL`.
     */
    haType?: pulumi.Input<string>;
    /**
     * Start time of the maintenance in one day, format like `HH:mm`.
     */
    maintenanceStartTime?: pulumi.Input<string>;
    /**
     * The timespan of maintenance in one day, unit is hour.
     */
    maintenanceTimeSpan?: pulumi.Input<number>;
    /**
     * A list of integer indicates weekly maintenance. For example, [2,7] presents do weekly maintenance on every Tuesday and Sunday.
     */
    maintenanceWeekSets?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloudSqlserverSpecinfos` provides.
     */
    memory: pulumi.Input<number>;
    /**
     * Indicate whether to deploy across availability zones.
     */
    multiZones?: pulumi.Input<boolean>;
    /**
     * Name of the SQL Server instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Purchase instance period in month. The value does not exceed 48.
     */
    period?: pulumi.Input<number>;
    /**
     * Project ID, default value is 0.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Security group bound to the instance.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storageMin` and `storageMax` which data source `tencentcloudSqlserverSpecinfos` provides.
     */
    storage: pulumi.Input<number>;
    /**
     * ID of subnet.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The tags of the SQL Server.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * An array of voucher IDs, currently only one can be used for a single order.
     */
    voucherIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of VPC.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.81.2. The way to execute the allocation. Supported values include: 0 - execute immediately, 1 - execute in maintenance window.
     *
     * @deprecated It has been deprecated from version 1.81.2.
     */
    waitSwitch?: pulumi.Input<number>;
}
