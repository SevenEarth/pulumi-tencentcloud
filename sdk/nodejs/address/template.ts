// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage address template.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const foo = new tencentcloud.address.Template("foo", {addresses: [
 *     "10.0.0.1",
 *     "10.0.1.0/24",
 *     "10.0.0.1-10.0.0.100",
 * ]});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Address template can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Address/template:Template foo ipm-makf7k9e"
 * ```
 */
export class Template extends pulumi.CustomResource {
    /**
     * Get an existing Template resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TemplateState, opts?: pulumi.CustomResourceOptions): Template {
        return new Template(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Address/template:Template';

    /**
     * Returns true if the given object is an instance of Template.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Template {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Template.__pulumiType;
    }

    /**
     * Address list. IP(`10.0.0.1`), CIDR(`10.0.1.0/24`), IP range(`10.0.0.1-10.0.0.100`) format are supported.
     */
    public readonly addresses!: pulumi.Output<string[]>;
    /**
     * Name of the address template.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Template resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TemplateArgs | TemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TemplateState | undefined;
            resourceInputs["addresses"] = state ? state.addresses : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as TemplateArgs | undefined;
            if ((!args || args.addresses === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addresses'");
            }
            resourceInputs["addresses"] = args ? args.addresses : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Template.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Template resources.
 */
export interface TemplateState {
    /**
     * Address list. IP(`10.0.0.1`), CIDR(`10.0.1.0/24`), IP range(`10.0.0.1-10.0.0.100`) format are supported.
     */
    addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the address template.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Template resource.
 */
export interface TemplateArgs {
    /**
     * Address list. IP(`10.0.0.1`), CIDR(`10.0.1.0/24`), IP range(`10.0.0.1-10.0.0.100`) format are supported.
     */
    addresses: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the address template.
     */
    name?: pulumi.Input<string>;
}
