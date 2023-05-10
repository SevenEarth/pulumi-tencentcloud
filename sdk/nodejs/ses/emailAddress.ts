// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ses emailAddress
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const emailAddress = new tencentcloud.Ses.EmailAddress("email_address", {
 *     emailAddress: "aaa@iac-tf.cloud",
 *     emailSenderName: "aaa",
 * });
 * ```
 *
 * ## Import
 *
 * ses email_address can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Ses/emailAddress:EmailAddress email_address aaa@iac-tf.cloud
 * ```
 */
export class EmailAddress extends pulumi.CustomResource {
    /**
     * Get an existing EmailAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EmailAddressState, opts?: pulumi.CustomResourceOptions): EmailAddress {
        return new EmailAddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ses/emailAddress:EmailAddress';

    /**
     * Returns true if the given object is an instance of EmailAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EmailAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EmailAddress.__pulumiType;
    }

    /**
     * Your sender address. (You can create up to 10 sender addresses for each domain.).
     */
    public readonly emailAddress!: pulumi.Output<string>;
    /**
     * Sender name.
     */
    public readonly emailSenderName!: pulumi.Output<string | undefined>;

    /**
     * Create a EmailAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EmailAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EmailAddressArgs | EmailAddressState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EmailAddressState | undefined;
            resourceInputs["emailAddress"] = state ? state.emailAddress : undefined;
            resourceInputs["emailSenderName"] = state ? state.emailSenderName : undefined;
        } else {
            const args = argsOrState as EmailAddressArgs | undefined;
            if ((!args || args.emailAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'emailAddress'");
            }
            resourceInputs["emailAddress"] = args ? args.emailAddress : undefined;
            resourceInputs["emailSenderName"] = args ? args.emailSenderName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EmailAddress.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EmailAddress resources.
 */
export interface EmailAddressState {
    /**
     * Your sender address. (You can create up to 10 sender addresses for each domain.).
     */
    emailAddress?: pulumi.Input<string>;
    /**
     * Sender name.
     */
    emailSenderName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EmailAddress resource.
 */
export interface EmailAddressArgs {
    /**
     * Your sender address. (You can create up to 10 sender addresses for each domain.).
     */
    emailAddress: pulumi.Input<string>;
    /**
     * Sender name.
     */
    emailSenderName?: pulumi.Input<string>;
}
