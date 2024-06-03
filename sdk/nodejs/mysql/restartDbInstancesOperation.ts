// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mysql restartDbInstancesOperation
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
 *     product: "cdb",
 * });
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     availabilityZone: zones.then(zones => zones.zones?.[0]?.name),
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
 *     slaveDeployMode: 0,
 *     availabilityZone: zones.then(zones => zones.zones?.[0]?.name),
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
 * const exampleRestartDbInstancesOperation = new tencentcloud.mysql.RestartDbInstancesOperation("exampleRestartDbInstancesOperation", {instanceId: exampleInstance.id});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * mysql restart_db_instances_operation can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Mysql/restartDbInstancesOperation:RestartDbInstancesOperation restart_db_instances_operation restart_db_instances_operation_id
 * ```
 */
export class RestartDbInstancesOperation extends pulumi.CustomResource {
    /**
     * Get an existing RestartDbInstancesOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RestartDbInstancesOperationState, opts?: pulumi.CustomResourceOptions): RestartDbInstancesOperation {
        return new RestartDbInstancesOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mysql/restartDbInstancesOperation:RestartDbInstancesOperation';

    /**
     * Returns true if the given object is an instance of RestartDbInstancesOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RestartDbInstancesOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RestartDbInstancesOperation.__pulumiType;
    }

    /**
     * An array of instance ID in the format: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Instance status.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;

    /**
     * Create a RestartDbInstancesOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RestartDbInstancesOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RestartDbInstancesOperationArgs | RestartDbInstancesOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RestartDbInstancesOperationState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as RestartDbInstancesOperationArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RestartDbInstancesOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RestartDbInstancesOperation resources.
 */
export interface RestartDbInstancesOperationState {
    /**
     * An array of instance ID in the format: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Instance status.
     */
    status?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RestartDbInstancesOperation resource.
 */
export interface RestartDbInstancesOperationArgs {
    /**
     * An array of instance ID in the format: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page.
     */
    instanceId: pulumi.Input<string>;
}
