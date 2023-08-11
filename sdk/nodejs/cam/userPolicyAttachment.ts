// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a CAM user policy attachment.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const config = new pulumi.Config();
 * const camUserBasic = config.get("camUserBasic") || "keep-cam-user";
 * const policyBasic = new tencentcloud.cam.Policy("policyBasic", {
 *     document: JSON.stringify({
 *         version: "2.0",
 *         statement: [
 *             {
 *                 action: ["cos:*"],
 *                 resource: ["*"],
 *                 effect: "allow",
 *             },
 *             {
 *                 effect: "allow",
 *                 action: [
 *                     "monitor:*",
 *                     "cam:ListUsersForGroup",
 *                     "cam:ListGroups",
 *                     "cam:GetGroup",
 *                 ],
 *                 resource: ["*"],
 *             },
 *         ],
 *     }),
 *     description: "tf_test",
 * });
 * const users = tencentcloud.Cam.getUsers({
 *     name: camUserBasic,
 * });
 * const userPolicyAttachmentBasic = new tencentcloud.cam.UserPolicyAttachment("userPolicyAttachmentBasic", {
 *     userName: users.then(users => users.userLists?[0]?.userId),
 *     policyId: policyBasic.id,
 * });
 * ```
 *
 * ## Import
 *
 * CAM user policy attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Cam/userPolicyAttachment:UserPolicyAttachment foo cam-test#26800353
 * ```
 */
export class UserPolicyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing UserPolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserPolicyAttachmentState, opts?: pulumi.CustomResourceOptions): UserPolicyAttachment {
        return new UserPolicyAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cam/userPolicyAttachment:UserPolicyAttachment';

    /**
     * Returns true if the given object is an instance of UserPolicyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPolicyAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPolicyAttachment.__pulumiType;
    }

    /**
     * Mode of Creation of the CAM user policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
     */
    public /*out*/ readonly createMode!: pulumi.Output<number>;
    /**
     * Create time of the CAM user policy attachment.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * ID of the policy.
     */
    public readonly policyId!: pulumi.Output<string>;
    /**
     * Name of the policy.
     */
    public /*out*/ readonly policyName!: pulumi.Output<string>;
    /**
     * Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
     */
    public /*out*/ readonly policyType!: pulumi.Output<string>;
    /**
     * It has been deprecated from version 1.59.5. Use `userName` instead. ID of the attached CAM user.
     *
     * @deprecated It has been deprecated from version 1.59.5. Use `user_name` instead.
     */
    public readonly userId!: pulumi.Output<string | undefined>;
    /**
     * Name of the attached CAM user as uniq key.
     */
    public readonly userName!: pulumi.Output<string | undefined>;

    /**
     * Create a UserPolicyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserPolicyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserPolicyAttachmentArgs | UserPolicyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserPolicyAttachmentState | undefined;
            resourceInputs["createMode"] = state ? state.createMode : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["policyId"] = state ? state.policyId : undefined;
            resourceInputs["policyName"] = state ? state.policyName : undefined;
            resourceInputs["policyType"] = state ? state.policyType : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as UserPolicyAttachmentArgs | undefined;
            if ((!args || args.policyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyId'");
            }
            resourceInputs["policyId"] = args ? args.policyId : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["createMode"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["policyName"] = undefined /*out*/;
            resourceInputs["policyType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserPolicyAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserPolicyAttachment resources.
 */
export interface UserPolicyAttachmentState {
    /**
     * Mode of Creation of the CAM user policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
     */
    createMode?: pulumi.Input<number>;
    /**
     * Create time of the CAM user policy attachment.
     */
    createTime?: pulumi.Input<string>;
    /**
     * ID of the policy.
     */
    policyId?: pulumi.Input<string>;
    /**
     * Name of the policy.
     */
    policyName?: pulumi.Input<string>;
    /**
     * Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
     */
    policyType?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.59.5. Use `userName` instead. ID of the attached CAM user.
     *
     * @deprecated It has been deprecated from version 1.59.5. Use `user_name` instead.
     */
    userId?: pulumi.Input<string>;
    /**
     * Name of the attached CAM user as uniq key.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserPolicyAttachment resource.
 */
export interface UserPolicyAttachmentArgs {
    /**
     * ID of the policy.
     */
    policyId: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.59.5. Use `userName` instead. ID of the attached CAM user.
     *
     * @deprecated It has been deprecated from version 1.59.5. Use `user_name` instead.
     */
    userId?: pulumi.Input<string>;
    /**
     * Name of the attached CAM user as uniq key.
     */
    userName?: pulumi.Input<string>;
}
