// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a SQL Server instance resource to create basic database instances.
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
 * const securityGroup = new tencentcloud.security.Group("securityGroup", {description: "desc."});
 * const example = new tencentcloud.sqlserver.BasicInstance("example", {
 *     availabilityZone: zones.then(zones => zones.zones?[4]?.name),
 *     chargeType: "POSTPAID_BY_HOUR",
 *     vpcId: vpc.id,
 *     subnetId: subnet.id,
 *     projectId: 0,
 *     memory: 4,
 *     storage: 100,
 *     cpu: 2,
 *     machineType: "CLOUD_PREMIUM",
 *     maintenanceWeekSets: [
 *         1,
 *         2,
 *         3,
 *     ],
 *     maintenanceStartTime: "09:00",
 *     maintenanceTimeSpan: 3,
 *     securityGroups: [securityGroup.id],
 *     tags: {
 *         test: "test",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * SQL Server basic instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Sqlserver/basicInstance:BasicInstance example mssql-3cdq7kx5
 * ```
 */
export class BasicInstance extends pulumi.CustomResource {
    /**
     * Get an existing BasicInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BasicInstanceState, opts?: pulumi.CustomResourceOptions): BasicInstance {
        return new BasicInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Sqlserver/basicInstance:BasicInstance';

    /**
     * Returns true if the given object is an instance of BasicInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BasicInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BasicInstance.__pulumiType;
    }

    /**
     * Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal, the default is 1 automatic renewal. Only valid when purchasing a prepaid instance.
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
     * Pay type of the SQL Server basic instance. For now, only `POSTPAID_BY_HOUR` is valid.
     */
    public readonly chargeType!: pulumi.Output<string | undefined>;
    /**
     * The CPU number of the SQL Server basic instance.
     */
    public readonly cpu!: pulumi.Output<number>;
    /**
     * Create time of the SQL Server basic instance.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Version of the SQL Server basic database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
     */
    public readonly engineVersion!: pulumi.Output<string | undefined>;
    /**
     * The host type of the purchased instance, `CLOUD_PREMIUM` for virtual machine high-performance cloud disk, `CLOUD_SSD` for virtual machine SSD cloud disk, `CLOUD_HSSD` for virtual machine enhanced cloud disk, `CLOUD_BSSD` for virtual machine general purpose SSD cloud disk.
     */
    public readonly machineType!: pulumi.Output<string>;
    /**
     * Start time of the maintenance in one day, format like `HH:mm`.
     */
    public readonly maintenanceStartTime!: pulumi.Output<string>;
    /**
     * The timespan of maintenance in one day, unit is hour.
     */
    public readonly maintenanceTimeSpan!: pulumi.Output<number>;
    /**
     * A list of integer indicates weekly maintenance. For example, [1,7] presents do weekly maintenance on every Monday and Sunday.
     */
    public readonly maintenanceWeekSets!: pulumi.Output<number[]>;
    /**
     * Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloudSqlserverSpecinfos` provides.
     */
    public readonly memory!: pulumi.Output<number>;
    /**
     * Name of the SQL Server basic instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Purchase instance period, the default value is 1, which means one month. The value does not exceed 48.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Project ID, default value is 0.
     */
    public readonly projectId!: pulumi.Output<number>;
    /**
     * Security group bound to the instance.
     */
    public readonly securityGroups!: pulumi.Output<string[] | undefined>;
    /**
     * Status of the SQL Server basic instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.
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
     * The tags of the SQL Server basic instance.
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
     * Create a BasicInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BasicInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BasicInstanceArgs | BasicInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BasicInstanceState | undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["autoVoucher"] = state ? state.autoVoucher : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["chargeType"] = state ? state.chargeType : undefined;
            resourceInputs["cpu"] = state ? state.cpu : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["engineVersion"] = state ? state.engineVersion : undefined;
            resourceInputs["machineType"] = state ? state.machineType : undefined;
            resourceInputs["maintenanceStartTime"] = state ? state.maintenanceStartTime : undefined;
            resourceInputs["maintenanceTimeSpan"] = state ? state.maintenanceTimeSpan : undefined;
            resourceInputs["maintenanceWeekSets"] = state ? state.maintenanceWeekSets : undefined;
            resourceInputs["memory"] = state ? state.memory : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["storage"] = state ? state.storage : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vip"] = state ? state.vip : undefined;
            resourceInputs["voucherIds"] = state ? state.voucherIds : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vport"] = state ? state.vport : undefined;
        } else {
            const args = argsOrState as BasicInstanceArgs | undefined;
            if ((!args || args.cpu === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cpu'");
            }
            if ((!args || args.machineType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'machineType'");
            }
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
            resourceInputs["cpu"] = args ? args.cpu : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["machineType"] = args ? args.machineType : undefined;
            resourceInputs["maintenanceStartTime"] = args ? args.maintenanceStartTime : undefined;
            resourceInputs["maintenanceTimeSpan"] = args ? args.maintenanceTimeSpan : undefined;
            resourceInputs["maintenanceWeekSets"] = args ? args.maintenanceWeekSets : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["storage"] = args ? args.storage : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["voucherIds"] = args ? args.voucherIds : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["vip"] = undefined /*out*/;
            resourceInputs["vport"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BasicInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BasicInstance resources.
 */
export interface BasicInstanceState {
    /**
     * Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal, the default is 1 automatic renewal. Only valid when purchasing a prepaid instance.
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
     * Pay type of the SQL Server basic instance. For now, only `POSTPAID_BY_HOUR` is valid.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * The CPU number of the SQL Server basic instance.
     */
    cpu?: pulumi.Input<number>;
    /**
     * Create time of the SQL Server basic instance.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Version of the SQL Server basic database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * The host type of the purchased instance, `CLOUD_PREMIUM` for virtual machine high-performance cloud disk, `CLOUD_SSD` for virtual machine SSD cloud disk, `CLOUD_HSSD` for virtual machine enhanced cloud disk, `CLOUD_BSSD` for virtual machine general purpose SSD cloud disk.
     */
    machineType?: pulumi.Input<string>;
    /**
     * Start time of the maintenance in one day, format like `HH:mm`.
     */
    maintenanceStartTime?: pulumi.Input<string>;
    /**
     * The timespan of maintenance in one day, unit is hour.
     */
    maintenanceTimeSpan?: pulumi.Input<number>;
    /**
     * A list of integer indicates weekly maintenance. For example, [1,7] presents do weekly maintenance on every Monday and Sunday.
     */
    maintenanceWeekSets?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloudSqlserverSpecinfos` provides.
     */
    memory?: pulumi.Input<number>;
    /**
     * Name of the SQL Server basic instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Purchase instance period, the default value is 1, which means one month. The value does not exceed 48.
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
     * Status of the SQL Server basic instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.
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
     * The tags of the SQL Server basic instance.
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
}

/**
 * The set of arguments for constructing a BasicInstance resource.
 */
export interface BasicInstanceArgs {
    /**
     * Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal, the default is 1 automatic renewal. Only valid when purchasing a prepaid instance.
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
     * Pay type of the SQL Server basic instance. For now, only `POSTPAID_BY_HOUR` is valid.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * The CPU number of the SQL Server basic instance.
     */
    cpu: pulumi.Input<number>;
    /**
     * Version of the SQL Server basic database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * The host type of the purchased instance, `CLOUD_PREMIUM` for virtual machine high-performance cloud disk, `CLOUD_SSD` for virtual machine SSD cloud disk, `CLOUD_HSSD` for virtual machine enhanced cloud disk, `CLOUD_BSSD` for virtual machine general purpose SSD cloud disk.
     */
    machineType: pulumi.Input<string>;
    /**
     * Start time of the maintenance in one day, format like `HH:mm`.
     */
    maintenanceStartTime?: pulumi.Input<string>;
    /**
     * The timespan of maintenance in one day, unit is hour.
     */
    maintenanceTimeSpan?: pulumi.Input<number>;
    /**
     * A list of integer indicates weekly maintenance. For example, [1,7] presents do weekly maintenance on every Monday and Sunday.
     */
    maintenanceWeekSets?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloudSqlserverSpecinfos` provides.
     */
    memory: pulumi.Input<number>;
    /**
     * Name of the SQL Server basic instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Purchase instance period, the default value is 1, which means one month. The value does not exceed 48.
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
     * The tags of the SQL Server basic instance.
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
}
