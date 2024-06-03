// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a bi userRole
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const userRole = new tencentcloud.bi.UserRole("userRole", {
 *     areaCode: "+83",
 *     email: "1055000000@qq.com",
 *     phoneNumber: "13470010000",
 *     roleIdLists: [10629359],
 *     userId: "100032767426",
 *     userName: "keep-iac-test",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * bi user_role can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Bi/userRole:UserRole user_role user_id
 * ```
 */
export class UserRole extends pulumi.CustomResource {
    /**
     * Get an existing UserRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserRoleState, opts?: pulumi.CustomResourceOptions): UserRole {
        return new UserRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Bi/userRole:UserRole';

    /**
     * Returns true if the given object is an instance of UserRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserRole.__pulumiType;
    }

    /**
     * Mobile area code(Note: This field may return null, indicating that no valid value can be obtained).
     */
    public readonly areaCode!: pulumi.Output<string>;
    /**
     * E-mail(Note: This field may return null, indicating that no valid value can be obtained).
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Phone number(Note: This field may return null, indicating that no valid value can be obtained).
     */
    public readonly phoneNumber!: pulumi.Output<string>;
    /**
     * Role id list.
     */
    public readonly roleIdLists!: pulumi.Output<number[]>;
    /**
     * User id.
     */
    public readonly userId!: pulumi.Output<string>;
    /**
     * Username.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a UserRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserRoleArgs | UserRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserRoleState | undefined;
            resourceInputs["areaCode"] = state ? state.areaCode : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["phoneNumber"] = state ? state.phoneNumber : undefined;
            resourceInputs["roleIdLists"] = state ? state.roleIdLists : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as UserRoleArgs | undefined;
            if ((!args || args.areaCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'areaCode'");
            }
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.phoneNumber === undefined) && !opts.urn) {
                throw new Error("Missing required property 'phoneNumber'");
            }
            if ((!args || args.roleIdLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleIdLists'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["areaCode"] = args ? args.areaCode : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["phoneNumber"] = args ? args.phoneNumber : undefined;
            resourceInputs["roleIdLists"] = args ? args.roleIdLists : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserRole resources.
 */
export interface UserRoleState {
    /**
     * Mobile area code(Note: This field may return null, indicating that no valid value can be obtained).
     */
    areaCode?: pulumi.Input<string>;
    /**
     * E-mail(Note: This field may return null, indicating that no valid value can be obtained).
     */
    email?: pulumi.Input<string>;
    /**
     * Phone number(Note: This field may return null, indicating that no valid value can be obtained).
     */
    phoneNumber?: pulumi.Input<string>;
    /**
     * Role id list.
     */
    roleIdLists?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * User id.
     */
    userId?: pulumi.Input<string>;
    /**
     * Username.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserRole resource.
 */
export interface UserRoleArgs {
    /**
     * Mobile area code(Note: This field may return null, indicating that no valid value can be obtained).
     */
    areaCode: pulumi.Input<string>;
    /**
     * E-mail(Note: This field may return null, indicating that no valid value can be obtained).
     */
    email: pulumi.Input<string>;
    /**
     * Phone number(Note: This field may return null, indicating that no valid value can be obtained).
     */
    phoneNumber: pulumi.Input<string>;
    /**
     * Role id list.
     */
    roleIdLists: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * User id.
     */
    userId: pulumi.Input<string>;
    /**
     * Username.
     */
    userName: pulumi.Input<string>;
}
