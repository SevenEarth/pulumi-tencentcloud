// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a lighthouse blueprint
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const blueprint = new tencentcloud.lighthouse.Blueprint("blueprint", {
 *     blueprintName: "blueprint_name_test",
 *     description: "blueprint_description_test",
 *     instanceId: "lhins-xxxxxx",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * lighthouse blueprint can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Lighthouse/blueprint:Blueprint blueprint blueprint_id
 * ```
 */
export class Blueprint extends pulumi.CustomResource {
    /**
     * Get an existing Blueprint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BlueprintState, opts?: pulumi.CustomResourceOptions): Blueprint {
        return new Blueprint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Lighthouse/blueprint:Blueprint';

    /**
     * Returns true if the given object is an instance of Blueprint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Blueprint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Blueprint.__pulumiType;
    }

    /**
     * Blueprint name, which can contain up to 60 characters.
     */
    public readonly blueprintName!: pulumi.Output<string>;
    /**
     * Blueprint description, which can contain up to 60 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ID of the instance for which to make a blueprint.
     */
    public readonly instanceId!: pulumi.Output<string | undefined>;

    /**
     * Create a Blueprint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BlueprintArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BlueprintArgs | BlueprintState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BlueprintState | undefined;
            resourceInputs["blueprintName"] = state ? state.blueprintName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as BlueprintArgs | undefined;
            if ((!args || args.blueprintName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'blueprintName'");
            }
            resourceInputs["blueprintName"] = args ? args.blueprintName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Blueprint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Blueprint resources.
 */
export interface BlueprintState {
    /**
     * Blueprint name, which can contain up to 60 characters.
     */
    blueprintName?: pulumi.Input<string>;
    /**
     * Blueprint description, which can contain up to 60 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of the instance for which to make a blueprint.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Blueprint resource.
 */
export interface BlueprintArgs {
    /**
     * Blueprint name, which can contain up to 60 characters.
     */
    blueprintName: pulumi.Input<string>;
    /**
     * Blueprint description, which can contain up to 60 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of the instance for which to make a blueprint.
     */
    instanceId?: pulumi.Input<string>;
}
