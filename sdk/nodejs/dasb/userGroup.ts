// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a dasb userGroup
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = new tencentcloud.Dasb.UserGroup("example", {});
 * ```
 * ### Or
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = new tencentcloud.Dasb.UserGroup("example", {
 *     departmentId: "1.2",
 * });
 * ```
 *
 * ## Import
 *
 * dasb user_group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Dasb/userGroup:UserGroup example 16
 * ```
 */
export class UserGroup extends pulumi.CustomResource {
    /**
     * Get an existing UserGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserGroupState, opts?: pulumi.CustomResourceOptions): UserGroup {
        return new UserGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dasb/userGroup:UserGroup';

    /**
     * Returns true if the given object is an instance of UserGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserGroup.__pulumiType;
    }

    /**
     * ID of the department to which the user group belongs, such as: 1.2.3.
     */
    public readonly departmentId!: pulumi.Output<string | undefined>;
    /**
     * User group name, maximum length 32 characters.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a UserGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserGroupArgs | UserGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserGroupState | undefined;
            resourceInputs["departmentId"] = state ? state.departmentId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as UserGroupArgs | undefined;
            resourceInputs["departmentId"] = args ? args.departmentId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserGroup resources.
 */
export interface UserGroupState {
    /**
     * ID of the department to which the user group belongs, such as: 1.2.3.
     */
    departmentId?: pulumi.Input<string>;
    /**
     * User group name, maximum length 32 characters.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserGroup resource.
 */
export interface UserGroupArgs {
    /**
     * ID of the department to which the user group belongs, such as: 1.2.3.
     */
    departmentId?: pulumi.Input<string>;
    /**
     * User group name, maximum length 32 characters.
     */
    name?: pulumi.Input<string>;
}
