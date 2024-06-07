// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const customLine = new tencentcloud.dnspod.CustomLine("customLine", {
 *     area: "6.6.6.1-6.6.6.2",
 *     domain: "dnspod.com",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * dnspod custom_line can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Dnspod/customLine:CustomLine custom_line domain#name
 * ```
 */
export class CustomLine extends pulumi.CustomResource {
    /**
     * Get an existing CustomLine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomLineState, opts?: pulumi.CustomResourceOptions): CustomLine {
        return new CustomLine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dnspod/customLine:CustomLine';

    /**
     * Returns true if the given object is an instance of CustomLine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomLine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomLine.__pulumiType;
    }

    /**
     * The IP segment of custom line, split with `-`.
     */
    public readonly area!: pulumi.Output<string>;
    /**
     * Domain.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * The Name of custom line.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a CustomLine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomLineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomLineArgs | CustomLineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomLineState | undefined;
            resourceInputs["area"] = state ? state.area : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as CustomLineArgs | undefined;
            if ((!args || args.area === undefined) && !opts.urn) {
                throw new Error("Missing required property 'area'");
            }
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            resourceInputs["area"] = args ? args.area : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomLine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomLine resources.
 */
export interface CustomLineState {
    /**
     * The IP segment of custom line, split with `-`.
     */
    area?: pulumi.Input<string>;
    /**
     * Domain.
     */
    domain?: pulumi.Input<string>;
    /**
     * The Name of custom line.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomLine resource.
 */
export interface CustomLineArgs {
    /**
     * The IP segment of custom line, split with `-`.
     */
    area: pulumi.Input<string>;
    /**
     * Domain.
     */
    domain: pulumi.Input<string>;
    /**
     * The Name of custom line.
     */
    name?: pulumi.Input<string>;
}
