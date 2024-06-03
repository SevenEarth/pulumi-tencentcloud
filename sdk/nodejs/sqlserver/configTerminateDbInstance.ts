// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a sqlserver configTerminateDbInstance
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const zones = tencentcloud.Availability.getZonesByProduct({
 *     product: "sqlserver",
 * });
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     availabilityZone: zones.then(zones => zones.zones?.[4]?.name),
 *     vpcId: vpc.id,
 *     cidrBlock: "10.0.0.0/16",
 *     isMulticast: false,
 * });
 * const securityGroup = new tencentcloud.security.Group("securityGroup", {description: "desc."});
 * const exampleBasicInstance = new tencentcloud.sqlserver.BasicInstance("exampleBasicInstance", {
 *     availabilityZone: zones.then(zones => zones.zones?.[4]?.name),
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
 * const exampleConfigTerminateDbInstance = new tencentcloud.sqlserver.ConfigTerminateDbInstance("exampleConfigTerminateDbInstance", {instanceId: exampleBasicInstance.id});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * sqlserver config_terminate_db_instance can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Sqlserver/configTerminateDbInstance:ConfigTerminateDbInstance example mssql-i9ma6oy7
 * ```
 */
export class ConfigTerminateDbInstance extends pulumi.CustomResource {
    /**
     * Get an existing ConfigTerminateDbInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigTerminateDbInstanceState, opts?: pulumi.CustomResourceOptions): ConfigTerminateDbInstance {
        return new ConfigTerminateDbInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Sqlserver/configTerminateDbInstance:ConfigTerminateDbInstance';

    /**
     * Returns true if the given object is an instance of ConfigTerminateDbInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigTerminateDbInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigTerminateDbInstance.__pulumiType;
    }

    /**
     * Instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a ConfigTerminateDbInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigTerminateDbInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigTerminateDbInstanceArgs | ConfigTerminateDbInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigTerminateDbInstanceState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as ConfigTerminateDbInstanceArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConfigTerminateDbInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigTerminateDbInstance resources.
 */
export interface ConfigTerminateDbInstanceState {
    /**
     * Instance ID.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConfigTerminateDbInstance resource.
 */
export interface ConfigTerminateDbInstanceArgs {
    /**
     * Instance ID.
     */
    instanceId: pulumi.Input<string>;
}
