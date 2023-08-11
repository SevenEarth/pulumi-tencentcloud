// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mysql roStopReplication
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const zones = tencentcloud.Availability.getZonesByProduct({
 *     product: "cdb",
 * });
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     availabilityZone: zones.then(zones => zones.zones?[0]?.name),
 *     vpcId: vpc.id,
 *     cidrBlock: "10.0.0.0/16",
 *     isMulticast: false,
 * });
 * const securityGroup = new tencentcloud.security.Group("securityGroup", {description: "mysql test"});
 * const exampleInstance = new tencentcloud.mysql.Instance("exampleInstance", {
 *     internetService: 1,
 *     engineVersion: "5.7",
 *     chargeType: "POSTPAID",
 *     rootPassword: "PassWord123",
 *     slaveDeployMode: 1,
 *     availabilityZone: zones.then(zones => zones.zones?[0]?.name),
 *     firstSlaveZone: zones.then(zones => zones.zones?[1]?.name),
 *     slaveSyncMode: 1,
 *     instanceName: "tf-example-mysql",
 *     memSize: 4000,
 *     volumeSize: 200,
 *     vpcId: vpc.id,
 *     subnetId: subnet.id,
 *     intranetPort: 3306,
 *     securityGroups: [securityGroup.id],
 *     tags: {
 *         name: "test",
 *     },
 *     parameters: {
 *         character_set_server: "utf8",
 *         max_connections: "1000",
 *     },
 * });
 * const exampleProxy = new tencentcloud.mysql.Proxy("exampleProxy", {
 *     instanceId: exampleInstance.id,
 *     uniqVpcId: vpc.id,
 *     uniqSubnetId: subnet.id,
 *     proxyNodeCustoms: [{
 *         nodeCount: 1,
 *         cpu: 2,
 *         mem: 4000,
 *         region: "ap-guangzhou",
 *         zone: "ap-guangzhou-3",
 *     }],
 *     securityGroups: [securityGroup.id],
 *     desc: "desc.",
 *     connectionPoolLimit: 2,
 *     vip: "10.0.0.120",
 *     vport: 3306,
 * });
 * const exampleRoStopReplication = new tencentcloud.mysql.RoStopReplication("exampleRoStopReplication", {instanceId: exampleProxy.id});
 * ```
 */
export class RoStopReplication extends pulumi.CustomResource {
    /**
     * Get an existing RoStopReplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RoStopReplicationState, opts?: pulumi.CustomResourceOptions): RoStopReplication {
        return new RoStopReplication(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mysql/roStopReplication:RoStopReplication';

    /**
     * Returns true if the given object is an instance of RoStopReplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RoStopReplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RoStopReplication.__pulumiType;
    }

    /**
     * Read-Only instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a RoStopReplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoStopReplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RoStopReplicationArgs | RoStopReplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RoStopReplicationState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as RoStopReplicationArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RoStopReplication.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RoStopReplication resources.
 */
export interface RoStopReplicationState {
    /**
     * Read-Only instance ID.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RoStopReplication resource.
 */
export interface RoStopReplicationArgs {
    /**
     * Read-Only instance ID.
     */
    instanceId: pulumi.Input<string>;
}
