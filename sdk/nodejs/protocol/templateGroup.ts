// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage protocol template group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.Protocol.TemplateGroup("foo", {
 *     templateIds: [
 *         "ipl-axaf24151",
 *         "ipl-axaf24152",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Protocol template group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Protocol/templateGroup:TemplateGroup foo ppmg-0np3u974
 * ```
 */
export class TemplateGroup extends pulumi.CustomResource {
    /**
     * Get an existing TemplateGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TemplateGroupState, opts?: pulumi.CustomResourceOptions): TemplateGroup {
        return new TemplateGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Protocol/templateGroup:TemplateGroup';

    /**
     * Returns true if the given object is an instance of TemplateGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TemplateGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TemplateGroup.__pulumiType;
    }

    /**
     * Name of the protocol template group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Service template ID list.
     */
    public readonly templateIds!: pulumi.Output<string[]>;

    /**
     * Create a TemplateGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TemplateGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TemplateGroupArgs | TemplateGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TemplateGroupState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["templateIds"] = state ? state.templateIds : undefined;
        } else {
            const args = argsOrState as TemplateGroupArgs | undefined;
            if ((!args || args.templateIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateIds'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["templateIds"] = args ? args.templateIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TemplateGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TemplateGroup resources.
 */
export interface TemplateGroupState {
    /**
     * Name of the protocol template group.
     */
    name?: pulumi.Input<string>;
    /**
     * Service template ID list.
     */
    templateIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a TemplateGroup resource.
 */
export interface TemplateGroupArgs {
    /**
     * Name of the protocol template group.
     */
    name?: pulumi.Input<string>;
    /**
     * Service template ID list.
     */
    templateIds: pulumi.Input<pulumi.Input<string>[]>;
}
