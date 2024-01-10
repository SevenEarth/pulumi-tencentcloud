// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ses blackList
 *
 * > **NOTE:** Used to remove email addresses from blacklists.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const blackList = new tencentcloud.Ses.BlackListDelete("black_list", {
 *     emailAddress: "terraform-tf@gmail.com",
 * });
 * ```
 */
export class BlackListDelete extends pulumi.CustomResource {
    /**
     * Get an existing BlackListDelete resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BlackListDeleteState, opts?: pulumi.CustomResourceOptions): BlackListDelete {
        return new BlackListDelete(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ses/blackListDelete:BlackListDelete';

    /**
     * Returns true if the given object is an instance of BlackListDelete.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BlackListDelete {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BlackListDelete.__pulumiType;
    }

    /**
     * Email addresses to be unblocklisted.
     */
    public readonly emailAddress!: pulumi.Output<string>;

    /**
     * Create a BlackListDelete resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BlackListDeleteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BlackListDeleteArgs | BlackListDeleteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BlackListDeleteState | undefined;
            resourceInputs["emailAddress"] = state ? state.emailAddress : undefined;
        } else {
            const args = argsOrState as BlackListDeleteArgs | undefined;
            if ((!args || args.emailAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'emailAddress'");
            }
            resourceInputs["emailAddress"] = args ? args.emailAddress : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BlackListDelete.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BlackListDelete resources.
 */
export interface BlackListDeleteState {
    /**
     * Email addresses to be unblocklisted.
     */
    emailAddress?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BlackListDelete resource.
 */
export interface BlackListDeleteArgs {
    /**
     * Email addresses to be unblocklisted.
     */
    emailAddress: pulumi.Input<string>;
}
