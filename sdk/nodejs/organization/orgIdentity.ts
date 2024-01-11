// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a organization orgIdentity
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const orgIdentity = new tencentcloud.Organization.OrgIdentity("org_identity", {
 *     description: "iac-test",
 *     identityAliasName: "example-iac-test",
 *     identityPolicies: [{
 *         policyId: 1,
 *         policyName: "AdministratorAccess",
 *         policyType: 2,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * organization org_identity can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Organization/orgIdentity:OrgIdentity org_identity org_identity_id
 * ```
 */
export class OrgIdentity extends pulumi.CustomResource {
    /**
     * Get an existing OrgIdentity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgIdentityState, opts?: pulumi.CustomResourceOptions): OrgIdentity {
        return new OrgIdentity(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Organization/orgIdentity:OrgIdentity';

    /**
     * Returns true if the given object is an instance of OrgIdentity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgIdentity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgIdentity.__pulumiType;
    }

    /**
     * Identity description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Identity name.Supports English letters and numbers, the length cannot exceed 40 characters.
     */
    public readonly identityAliasName!: pulumi.Output<string>;
    /**
     * Identity policy list.
     */
    public readonly identityPolicies!: pulumi.Output<outputs.Organization.OrgIdentityIdentityPolicy[]>;

    /**
     * Create a OrgIdentity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrgIdentityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgIdentityArgs | OrgIdentityState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgIdentityState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["identityAliasName"] = state ? state.identityAliasName : undefined;
            resourceInputs["identityPolicies"] = state ? state.identityPolicies : undefined;
        } else {
            const args = argsOrState as OrgIdentityArgs | undefined;
            if ((!args || args.identityAliasName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identityAliasName'");
            }
            if ((!args || args.identityPolicies === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identityPolicies'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["identityAliasName"] = args ? args.identityAliasName : undefined;
            resourceInputs["identityPolicies"] = args ? args.identityPolicies : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrgIdentity.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgIdentity resources.
 */
export interface OrgIdentityState {
    /**
     * Identity description.
     */
    description?: pulumi.Input<string>;
    /**
     * Identity name.Supports English letters and numbers, the length cannot exceed 40 characters.
     */
    identityAliasName?: pulumi.Input<string>;
    /**
     * Identity policy list.
     */
    identityPolicies?: pulumi.Input<pulumi.Input<inputs.Organization.OrgIdentityIdentityPolicy>[]>;
}

/**
 * The set of arguments for constructing a OrgIdentity resource.
 */
export interface OrgIdentityArgs {
    /**
     * Identity description.
     */
    description?: pulumi.Input<string>;
    /**
     * Identity name.Supports English letters and numbers, the length cannot exceed 40 characters.
     */
    identityAliasName: pulumi.Input<string>;
    /**
     * Identity policy list.
     */
    identityPolicies: pulumi.Input<pulumi.Input<inputs.Organization.OrgIdentityIdentityPolicy>[]>;
}