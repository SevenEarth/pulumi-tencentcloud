// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a organization orgMemberEmail
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const orgMemberEmail = new tencentcloud.Organization.OrgMemberEmail("org_member_email", {
 *     countryCode: "86",
 *     email: "iac-example@qq.com",
 *     memberUin: 100033704327,
 *     phone: "12345678901",
 * });
 * ```
 *
 * ## Import
 *
 * organization org_member_email can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Organization/orgMemberEmail:OrgMemberEmail org_member_email org_member_email_id
 * ```
 */
export class OrgMemberEmail extends pulumi.CustomResource {
    /**
     * Get an existing OrgMemberEmail resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgMemberEmailState, opts?: pulumi.CustomResourceOptions): OrgMemberEmail {
        return new OrgMemberEmail(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Organization/orgMemberEmail:OrgMemberEmail';

    /**
     * Returns true if the given object is an instance of OrgMemberEmail.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgMemberEmail {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgMemberEmail.__pulumiType;
    }

    /**
     * Application timeNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    public /*out*/ readonly applyTime!: pulumi.Output<string>;
    /**
     * Binding IDNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    public /*out*/ readonly bindId!: pulumi.Output<number>;
    /**
     * Binding status is not binding: unbound, to be activated: value, successful binding: success, binding failure: failedNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    public /*out*/ readonly bindStatus!: pulumi.Output<string>;
    /**
     * Binding timeNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    public /*out*/ readonly bindTime!: pulumi.Output<string>;
    /**
     * International region.
     */
    public readonly countryCode!: pulumi.Output<string>;
    /**
     * FailedNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * Email address.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Member Uin.
     */
    public readonly memberUin!: pulumi.Output<number>;
    /**
     * Phone number.
     */
    public readonly phone!: pulumi.Output<string>;
    /**
     * Safe mobile phone binding state is not bound: 0, has been binded: 1Note: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    public /*out*/ readonly phoneBind!: pulumi.Output<number>;

    /**
     * Create a OrgMemberEmail resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrgMemberEmailArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgMemberEmailArgs | OrgMemberEmailState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgMemberEmailState | undefined;
            resourceInputs["applyTime"] = state ? state.applyTime : undefined;
            resourceInputs["bindId"] = state ? state.bindId : undefined;
            resourceInputs["bindStatus"] = state ? state.bindStatus : undefined;
            resourceInputs["bindTime"] = state ? state.bindTime : undefined;
            resourceInputs["countryCode"] = state ? state.countryCode : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["memberUin"] = state ? state.memberUin : undefined;
            resourceInputs["phone"] = state ? state.phone : undefined;
            resourceInputs["phoneBind"] = state ? state.phoneBind : undefined;
        } else {
            const args = argsOrState as OrgMemberEmailArgs | undefined;
            if ((!args || args.countryCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'countryCode'");
            }
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.memberUin === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memberUin'");
            }
            if ((!args || args.phone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'phone'");
            }
            resourceInputs["countryCode"] = args ? args.countryCode : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["memberUin"] = args ? args.memberUin : undefined;
            resourceInputs["phone"] = args ? args.phone : undefined;
            resourceInputs["applyTime"] = undefined /*out*/;
            resourceInputs["bindId"] = undefined /*out*/;
            resourceInputs["bindStatus"] = undefined /*out*/;
            resourceInputs["bindTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["phoneBind"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrgMemberEmail.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgMemberEmail resources.
 */
export interface OrgMemberEmailState {
    /**
     * Application timeNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    applyTime?: pulumi.Input<string>;
    /**
     * Binding IDNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    bindId?: pulumi.Input<number>;
    /**
     * Binding status is not binding: unbound, to be activated: value, successful binding: success, binding failure: failedNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    bindStatus?: pulumi.Input<string>;
    /**
     * Binding timeNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    bindTime?: pulumi.Input<string>;
    /**
     * International region.
     */
    countryCode?: pulumi.Input<string>;
    /**
     * FailedNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    description?: pulumi.Input<string>;
    /**
     * Email address.
     */
    email?: pulumi.Input<string>;
    /**
     * Member Uin.
     */
    memberUin?: pulumi.Input<number>;
    /**
     * Phone number.
     */
    phone?: pulumi.Input<string>;
    /**
     * Safe mobile phone binding state is not bound: 0, has been binded: 1Note: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    phoneBind?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a OrgMemberEmail resource.
 */
export interface OrgMemberEmailArgs {
    /**
     * International region.
     */
    countryCode: pulumi.Input<string>;
    /**
     * Email address.
     */
    email: pulumi.Input<string>;
    /**
     * Member Uin.
     */
    memberUin: pulumi.Input<number>;
    /**
     * Phone number.
     */
    phone: pulumi.Input<string>;
}
