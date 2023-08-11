// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tag
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const tag = new tencentcloud.Tag.Instance("tag", {
 *     tagKey: "test",
 *     tagValue: "Terraform",
 * });
 * ```
 *
 * ## Import
 *
 * tag tag can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Tag/instance:Instance tag tag_id
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
    public static readonly __pulumiType = 'tencentcloud:Tag/instance:Instance';

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
     * tag key.
     */
    public readonly tagKey!: pulumi.Output<string>;
    /**
     * tag value.
     */
    public readonly tagValue!: pulumi.Output<string>;

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
            resourceInputs["tagKey"] = state ? state.tagKey : undefined;
            resourceInputs["tagValue"] = state ? state.tagValue : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.tagKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tagKey'");
            }
            if ((!args || args.tagValue === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tagValue'");
            }
            resourceInputs["tagKey"] = args ? args.tagKey : undefined;
            resourceInputs["tagValue"] = args ? args.tagValue : undefined;
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
     * tag key.
     */
    tagKey?: pulumi.Input<string>;
    /**
     * tag value.
     */
    tagValue?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * tag key.
     */
    tagKey: pulumi.Input<string>;
    /**
     * tag value.
     */
    tagValue: pulumi.Input<string>;
}
