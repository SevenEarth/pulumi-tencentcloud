// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ses verifyDomain
 *
 * > **NOTE:** Please add the `attributes` information returned by `tencentcloud.Ses.Domain` to the domain name resolution record through `tencentcloud.Dnspod.Record`, and then verify it.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const verifyDomain = new tencentcloud.Ses.VerifyDomain("verify_domain", {
 *     emailIdentity: "example.com",
 * });
 * ```
 */
export class VerifyDomain extends pulumi.CustomResource {
    /**
     * Get an existing VerifyDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VerifyDomainState, opts?: pulumi.CustomResourceOptions): VerifyDomain {
        return new VerifyDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ses/verifyDomain:VerifyDomain';

    /**
     * Returns true if the given object is an instance of VerifyDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VerifyDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VerifyDomain.__pulumiType;
    }

    /**
     * Domain name requested for verification.
     */
    public readonly emailIdentity!: pulumi.Output<string>;

    /**
     * Create a VerifyDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VerifyDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VerifyDomainArgs | VerifyDomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VerifyDomainState | undefined;
            resourceInputs["emailIdentity"] = state ? state.emailIdentity : undefined;
        } else {
            const args = argsOrState as VerifyDomainArgs | undefined;
            if ((!args || args.emailIdentity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'emailIdentity'");
            }
            resourceInputs["emailIdentity"] = args ? args.emailIdentity : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VerifyDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VerifyDomain resources.
 */
export interface VerifyDomainState {
    /**
     * Domain name requested for verification.
     */
    emailIdentity?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VerifyDomain resource.
 */
export interface VerifyDomainArgs {
    /**
     * Domain name requested for verification.
     */
    emailIdentity: pulumi.Input<string>;
}