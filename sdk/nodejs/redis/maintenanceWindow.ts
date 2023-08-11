// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a redis maintenanceWindow
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const zone = tencentcloud.Redis.getZoneConfig({
 *     typeId: 7,
 * });
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     vpcId: vpc.id,
 *     availabilityZone: zone.then(zone => zone.lists?[0]?.zone),
 *     cidrBlock: "10.0.1.0/24",
 * });
 * const fooInstance = new tencentcloud.redis.Instance("fooInstance", {
 *     availabilityZone: zone.then(zone => zone.lists?[0]?.zone),
 *     typeId: zone.then(zone => zone.lists?[0]?.typeId),
 *     password: "test12345789",
 *     memSize: 8192,
 *     redisShardNum: zone.then(zone => zone.lists?[0]?.redisShardNums?[0]),
 *     redisReplicasNum: zone.then(zone => zone.lists?[0]?.redisReplicasNums?[0]),
 *     port: 6379,
 *     vpcId: vpc.id,
 *     subnetId: subnet.id,
 * });
 * const fooMaintenanceWindow = new tencentcloud.redis.MaintenanceWindow("fooMaintenanceWindow", {
 *     instanceId: fooInstance.id,
 *     startTime: "17:00",
 *     endTime: "19:00",
 * });
 * ```
 *
 * ## Import
 *
 * redis maintenance_window can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Redis/maintenanceWindow:MaintenanceWindow foo instance_id
 * ```
 */
export class MaintenanceWindow extends pulumi.CustomResource {
    /**
     * Get an existing MaintenanceWindow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MaintenanceWindowState, opts?: pulumi.CustomResourceOptions): MaintenanceWindow {
        return new MaintenanceWindow(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Redis/maintenanceWindow:MaintenanceWindow';

    /**
     * Returns true if the given object is an instance of MaintenanceWindow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MaintenanceWindow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MaintenanceWindow.__pulumiType;
    }

    /**
     * The end time of the maintenance window, e.g. 19:00.
     */
    public readonly endTime!: pulumi.Output<string>;
    /**
     * The ID of instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Maintenance window start time, e.g. 17:00.
     */
    public readonly startTime!: pulumi.Output<string>;

    /**
     * Create a MaintenanceWindow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MaintenanceWindowArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MaintenanceWindowArgs | MaintenanceWindowState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MaintenanceWindowState | undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
        } else {
            const args = argsOrState as MaintenanceWindowArgs | undefined;
            if ((!args || args.endTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endTime'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.startTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startTime'");
            }
            resourceInputs["endTime"] = args ? args.endTime : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MaintenanceWindow.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MaintenanceWindow resources.
 */
export interface MaintenanceWindowState {
    /**
     * The end time of the maintenance window, e.g. 19:00.
     */
    endTime?: pulumi.Input<string>;
    /**
     * The ID of instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Maintenance window start time, e.g. 17:00.
     */
    startTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MaintenanceWindow resource.
 */
export interface MaintenanceWindowArgs {
    /**
     * The end time of the maintenance window, e.g. 19:00.
     */
    endTime: pulumi.Input<string>;
    /**
     * The ID of instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Maintenance window start time, e.g. 17:00.
     */
    startTime: pulumi.Input<string>;
}
