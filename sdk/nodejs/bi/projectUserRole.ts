// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a bi projectUserRole
 *
 * > **NOTE:** You cannot use `tencentcloud.Bi.UserRole` and `tencentcloud.Bi.ProjectUserRole` at the same time to modify the `phoneNumber` and `email` of the same user.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const projectUserRole = new tencentcloud.bi.ProjectUserRole("projectUserRole", {
 *     areaCode: "+86",
 *     email: "123456@qq.com",
 *     phoneNumber: "13130001000",
 *     projectId: 11015030,
 *     roleIdLists: [10629453],
 *     userId: "100024664626",
 *     userName: "keep-cam-user",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * bi project_user_role can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Bi/projectUserRole:ProjectUserRole project_user_role projectId#userId
 * ```
 */
export class ProjectUserRole extends pulumi.CustomResource {
    /**
     * Get an existing ProjectUserRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectUserRoleState, opts?: pulumi.CustomResourceOptions): ProjectUserRole {
        return new ProjectUserRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Bi/projectUserRole:ProjectUserRole';

    /**
     * Returns true if the given object is an instance of ProjectUserRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectUserRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectUserRole.__pulumiType;
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
     * Project id.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * Role id list.
     */
    public readonly roleIdLists!: pulumi.Output<number[] | undefined>;
    /**
     * User id.
     */
    public readonly userId!: pulumi.Output<string>;
    /**
     * Username.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a ProjectUserRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectUserRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectUserRoleArgs | ProjectUserRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectUserRoleState | undefined;
            resourceInputs["areaCode"] = state ? state.areaCode : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["phoneNumber"] = state ? state.phoneNumber : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["roleIdLists"] = state ? state.roleIdLists : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as ProjectUserRoleArgs | undefined;
            if ((!args || args.areaCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'areaCode'");
            }
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.phoneNumber === undefined) && !opts.urn) {
                throw new Error("Missing required property 'phoneNumber'");
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
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["roleIdLists"] = args ? args.roleIdLists : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectUserRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectUserRole resources.
 */
export interface ProjectUserRoleState {
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
     * Project id.
     */
    projectId?: pulumi.Input<number>;
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
 * The set of arguments for constructing a ProjectUserRole resource.
 */
export interface ProjectUserRoleArgs {
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
     * Project id.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Role id list.
     */
    roleIdLists?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * User id.
     */
    userId: pulumi.Input<string>;
    /**
     * Username.
     */
    userName: pulumi.Input<string>;
}
