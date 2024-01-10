// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a elasticsearch index
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const index = new tencentcloud.Elasticsearch.Index("index", {
 *     indexMetaJson: "{\"mappings\":{},\"settings\":{\"index.number_of_replicas\":1,\"index.number_of_shards\":1,\"index.refresh_interval\":\"30s\"}}",
 *     indexName: "test-es-index",
 *     indexType: "normal",
 *     instanceId: "es-xxxxxx",
 * });
 * ```
 *
 * ## Import
 *
 * elasticsearch index can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Elasticsearch/index:Index index index_id
 * ```
 */
export class Index extends pulumi.CustomResource {
    /**
     * Get an existing Index resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IndexState, opts?: pulumi.CustomResourceOptions): Index {
        return new Index(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Elasticsearch/index:Index';

    /**
     * Returns true if the given object is an instance of Index.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Index {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Index.__pulumiType;
    }

    /**
     * Create index metadata JSON, such as mappings, settings.
     */
    public readonly indexMetaJson!: pulumi.Output<string | undefined>;
    /**
     * index name to create.
     */
    public readonly indexName!: pulumi.Output<string>;
    /**
     * type of the index to be created. auto: autonomous index. normal: indicates a common index.
     */
    public readonly indexType!: pulumi.Output<string>;
    /**
     * es instance id.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a Index resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IndexArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IndexArgs | IndexState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IndexState | undefined;
            resourceInputs["indexMetaJson"] = state ? state.indexMetaJson : undefined;
            resourceInputs["indexName"] = state ? state.indexName : undefined;
            resourceInputs["indexType"] = state ? state.indexType : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as IndexArgs | undefined;
            if ((!args || args.indexName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'indexName'");
            }
            if ((!args || args.indexType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'indexType'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["indexMetaJson"] = args ? args.indexMetaJson : undefined;
            resourceInputs["indexName"] = args ? args.indexName : undefined;
            resourceInputs["indexType"] = args ? args.indexType : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Index.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Index resources.
 */
export interface IndexState {
    /**
     * Create index metadata JSON, such as mappings, settings.
     */
    indexMetaJson?: pulumi.Input<string>;
    /**
     * index name to create.
     */
    indexName?: pulumi.Input<string>;
    /**
     * type of the index to be created. auto: autonomous index. normal: indicates a common index.
     */
    indexType?: pulumi.Input<string>;
    /**
     * es instance id.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Index resource.
 */
export interface IndexArgs {
    /**
     * Create index metadata JSON, such as mappings, settings.
     */
    indexMetaJson?: pulumi.Input<string>;
    /**
     * index name to create.
     */
    indexName: pulumi.Input<string>;
    /**
     * type of the index to be created. auto: autonomous index. normal: indicates a common index.
     */
    indexType: pulumi.Input<string>;
    /**
     * es instance id.
     */
    instanceId: pulumi.Input<string>;
}
