// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mariadb hourDbInstance
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const basic = new tencentcloud.Mariadb.HourDbInstance("basic", {
 *     dbVersionId: "8.0",
 *     instanceName: "db-test-2",
 *     memory: 2,
 *     nodeCount: 2,
 *     storage: 10,
 *     subnetId: "subnet-jdi5xn22",
 *     tags: {
 *         createdBy: "terraform",
 *     },
 *     vpcId: "vpc-k1t8ickr",
 *     zones: [
 *         "ap-guangzhou-7",
 *         "ap-guangzhou-7",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * mariadb hour_db_instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Mariadb/hourDbInstance:HourDbInstance hour_db_instance tdsql-kjqih9nn
 * ```
 */
export class HourDbInstance extends pulumi.CustomResource {
    /**
     * Get an existing HourDbInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HourDbInstanceState, opts?: pulumi.CustomResourceOptions): HourDbInstance {
        return new HourDbInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mariadb/hourDbInstance:HourDbInstance';

    /**
     * Returns true if the given object is an instance of HourDbInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HourDbInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HourDbInstance.__pulumiType;
    }

    /**
     * db engine version, default to 10.1.9.
     */
    public readonly dbVersionId!: pulumi.Output<string>;
    /**
     * name of this instance.
     */
    public readonly instanceName!: pulumi.Output<string | undefined>;
    /**
     * instance memory.
     */
    public readonly memory!: pulumi.Output<number>;
    /**
     * number of node for instance.
     */
    public readonly nodeCount!: pulumi.Output<number>;
    /**
     * instance disk storage.
     */
    public readonly storage!: pulumi.Output<number>;
    /**
     * subnet id, it&amp;#39;s required when vpcId is set.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * Tag description list.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * vpc id.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * available zone of instance.
     */
    public readonly zones!: pulumi.Output<string[]>;

    /**
     * Create a HourDbInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HourDbInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HourDbInstanceArgs | HourDbInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HourDbInstanceState | undefined;
            resourceInputs["dbVersionId"] = state ? state.dbVersionId : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["memory"] = state ? state.memory : undefined;
            resourceInputs["nodeCount"] = state ? state.nodeCount : undefined;
            resourceInputs["storage"] = state ? state.storage : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as HourDbInstanceArgs | undefined;
            if ((!args || args.memory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memory'");
            }
            if ((!args || args.nodeCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeCount'");
            }
            if ((!args || args.storage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storage'");
            }
            if ((!args || args.zones === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zones'");
            }
            resourceInputs["dbVersionId"] = args ? args.dbVersionId : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["nodeCount"] = args ? args.nodeCount : undefined;
            resourceInputs["storage"] = args ? args.storage : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["zones"] = args ? args.zones : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HourDbInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HourDbInstance resources.
 */
export interface HourDbInstanceState {
    /**
     * db engine version, default to 10.1.9.
     */
    dbVersionId?: pulumi.Input<string>;
    /**
     * name of this instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * instance memory.
     */
    memory?: pulumi.Input<number>;
    /**
     * number of node for instance.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * instance disk storage.
     */
    storage?: pulumi.Input<number>;
    /**
     * subnet id, it&amp;#39;s required when vpcId is set.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * vpc id.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * available zone of instance.
     */
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a HourDbInstance resource.
 */
export interface HourDbInstanceArgs {
    /**
     * db engine version, default to 10.1.9.
     */
    dbVersionId?: pulumi.Input<string>;
    /**
     * name of this instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * instance memory.
     */
    memory: pulumi.Input<number>;
    /**
     * number of node for instance.
     */
    nodeCount: pulumi.Input<number>;
    /**
     * instance disk storage.
     */
    storage: pulumi.Input<number>;
    /**
     * subnet id, it&amp;#39;s required when vpcId is set.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * vpc id.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * available zone of instance.
     */
    zones: pulumi.Input<pulumi.Input<string>[]>;
}