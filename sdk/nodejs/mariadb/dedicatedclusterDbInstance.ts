// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mariadb dedicatedclusterDbInstance
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const dedicatedclusterDbInstance = new tencentcloud.Mariadb.DedicatedclusterDbInstance("dedicatedcluster_db_instance", {
 *     clusterId: "dbdc-24odnuhr",
 *     dbVersionId: "8.0",
 *     goodsNum: 1,
 *     instanceName: "cluster-mariadb-test-1",
 *     memory: 2,
 *     storage: 10,
 *     subnetId: "subnet-3ku415by",
 *     vpcId: "vpc-ii1jfbhl",
 * });
 * ```
 *
 * ## Import
 *
 * mariadb dedicatedcluster_db_instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Mariadb/dedicatedclusterDbInstance:DedicatedclusterDbInstance dedicatedcluster_db_instance tdsql-050g3fmv
 * ```
 */
export class DedicatedclusterDbInstance extends pulumi.CustomResource {
    /**
     * Get an existing DedicatedclusterDbInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DedicatedclusterDbInstanceState, opts?: pulumi.CustomResourceOptions): DedicatedclusterDbInstance {
        return new DedicatedclusterDbInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mariadb/dedicatedclusterDbInstance:DedicatedclusterDbInstance';

    /**
     * Returns true if the given object is an instance of DedicatedclusterDbInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DedicatedclusterDbInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DedicatedclusterDbInstance.__pulumiType;
    }

    /**
     * dedicated cluster id.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * db engine version, default to 0.
     */
    public readonly dbVersionId!: pulumi.Output<string>;
    /**
     * number of instance.
     */
    public readonly goodsNum!: pulumi.Output<number>;
    /**
     * name of this instance.
     */
    public readonly instanceName!: pulumi.Output<string | undefined>;
    /**
     * instance memory.
     */
    public readonly memory!: pulumi.Output<number>;
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
     * Create a DedicatedclusterDbInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DedicatedclusterDbInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DedicatedclusterDbInstanceArgs | DedicatedclusterDbInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DedicatedclusterDbInstanceState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["dbVersionId"] = state ? state.dbVersionId : undefined;
            resourceInputs["goodsNum"] = state ? state.goodsNum : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["memory"] = state ? state.memory : undefined;
            resourceInputs["storage"] = state ? state.storage : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as DedicatedclusterDbInstanceArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.goodsNum === undefined) && !opts.urn) {
                throw new Error("Missing required property 'goodsNum'");
            }
            if ((!args || args.memory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memory'");
            }
            if ((!args || args.storage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storage'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["dbVersionId"] = args ? args.dbVersionId : undefined;
            resourceInputs["goodsNum"] = args ? args.goodsNum : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["storage"] = args ? args.storage : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DedicatedclusterDbInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DedicatedclusterDbInstance resources.
 */
export interface DedicatedclusterDbInstanceState {
    /**
     * dedicated cluster id.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * db engine version, default to 0.
     */
    dbVersionId?: pulumi.Input<string>;
    /**
     * number of instance.
     */
    goodsNum?: pulumi.Input<number>;
    /**
     * name of this instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * instance memory.
     */
    memory?: pulumi.Input<number>;
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
}

/**
 * The set of arguments for constructing a DedicatedclusterDbInstance resource.
 */
export interface DedicatedclusterDbInstanceArgs {
    /**
     * dedicated cluster id.
     */
    clusterId: pulumi.Input<string>;
    /**
     * db engine version, default to 0.
     */
    dbVersionId?: pulumi.Input<string>;
    /**
     * number of instance.
     */
    goodsNum: pulumi.Input<number>;
    /**
     * name of this instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * instance memory.
     */
    memory: pulumi.Input<number>;
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
}
