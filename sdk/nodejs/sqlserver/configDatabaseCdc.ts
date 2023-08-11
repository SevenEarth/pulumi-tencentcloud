// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a sqlserver configDatabaseCdc
 *
 * ## Example Usage
 * ### Turn off database data change capture (CDC)
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
 * const exampleBasicInstance = new tencentcloud.sqlserver.BasicInstance("exampleBasicInstance", {
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
 * const exampleDb = new tencentcloud.sqlserver.Db("exampleDb", {
 *     instanceId: exampleBasicInstance.id,
 *     charset: "Chinese_PRC_BIN",
 *     remark: "test-remark",
 * });
 * const exampleConfigDatabaseCdc = new tencentcloud.sqlserver.ConfigDatabaseCdc("exampleConfigDatabaseCdc", {
 *     instanceId: exampleBasicInstance.id,
 *     dbName: exampleDb.name,
 *     modifyType: "disable",
 * });
 * ```
 * ### Enable Database Data Change Capture (CDC)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.sqlserver.ConfigDatabaseCdc("example", {
 *     instanceId: tencentcloud_sqlserver_basic_instance.example.id,
 *     dbName: tencentcloud_sqlserver_db.example.name,
 *     modifyType: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * sqlserver config_database_cdc can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Sqlserver/configDatabaseCdc:ConfigDatabaseCdc example mssql-i9ma6oy7#tf_example_db
 * ```
 */
export class ConfigDatabaseCdc extends pulumi.CustomResource {
    /**
     * Get an existing ConfigDatabaseCdc resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigDatabaseCdcState, opts?: pulumi.CustomResourceOptions): ConfigDatabaseCdc {
        return new ConfigDatabaseCdc(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Sqlserver/configDatabaseCdc:ConfigDatabaseCdc';

    /**
     * Returns true if the given object is an instance of ConfigDatabaseCdc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigDatabaseCdc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigDatabaseCdc.__pulumiType;
    }

    /**
     * database name.
     */
    public readonly dbName!: pulumi.Output<string>;
    /**
     * Instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Enable or disable CDC. Valid values: enable, disable.
     */
    public readonly modifyType!: pulumi.Output<string>;

    /**
     * Create a ConfigDatabaseCdc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigDatabaseCdcArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigDatabaseCdcArgs | ConfigDatabaseCdcState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigDatabaseCdcState | undefined;
            resourceInputs["dbName"] = state ? state.dbName : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["modifyType"] = state ? state.modifyType : undefined;
        } else {
            const args = argsOrState as ConfigDatabaseCdcArgs | undefined;
            if ((!args || args.dbName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbName'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.modifyType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'modifyType'");
            }
            resourceInputs["dbName"] = args ? args.dbName : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["modifyType"] = args ? args.modifyType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConfigDatabaseCdc.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigDatabaseCdc resources.
 */
export interface ConfigDatabaseCdcState {
    /**
     * database name.
     */
    dbName?: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Enable or disable CDC. Valid values: enable, disable.
     */
    modifyType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConfigDatabaseCdc resource.
 */
export interface ConfigDatabaseCdcArgs {
    /**
     * database name.
     */
    dbName: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Enable or disable CDC. Valid values: enable, disable.
     */
    modifyType: pulumi.Input<string>;
}
