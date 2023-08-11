// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage CAM user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.Cam.User("foo", {
 *     consoleLogin: true,
 *     countryCode: "86",
 *     email: "hello@test.com",
 *     forceDelete: true,
 *     needResetPassword: true,
 *     password: "Gail@1234",
 *     phoneNum: "12345678910",
 *     remark: "tf_user_test",
 *     tags: {
 *         test: "tf_cam_user",
 *     },
 *     useApi: true,
 * });
 * ```
 *
 * ## Import
 *
 * CAM user can be imported using the user name, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Cam/user:User foo cam-user-test
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cam/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * Indicate whether the CAM user can login to the web console or not.
     */
    public readonly consoleLogin!: pulumi.Output<boolean | undefined>;
    /**
     * Country code of the phone number, for example: '86'.
     */
    public readonly countryCode!: pulumi.Output<string>;
    /**
     * Email of the CAM user.
     */
    public readonly email!: pulumi.Output<string | undefined>;
    /**
     * Indicate whether to force deletes the CAM user. If set false, the API secret key will be checked and failed when exists; otherwise the user will be deleted directly. Default is false.
     */
    public readonly forceDelete!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the CAM user.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Indicate whether the CAM user need to reset the password when first logins.
     */
    public readonly needResetPassword!: pulumi.Output<boolean | undefined>;
    /**
     * The password of the CAM user. Password should be at least 8 characters and no more than 32 characters, includes uppercase letters, lowercase letters, numbers and special characters. Only required when `consoleLogin` is true. If not set, a random password will be automatically generated.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Phone number of the CAM user.
     */
    public readonly phoneNum!: pulumi.Output<string | undefined>;
    /**
     * Remark of the CAM user.
     */
    public readonly remark!: pulumi.Output<string | undefined>;
    /**
     * Secret ID of the CAM user.
     */
    public /*out*/ readonly secretId!: pulumi.Output<string>;
    /**
     * Secret key of the CAM user.
     */
    public /*out*/ readonly secretKey!: pulumi.Output<string>;
    /**
     * A list of tags used to associate different resources.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * ID of the CAM user.
     */
    public /*out*/ readonly uid!: pulumi.Output<number>;
    /**
     * Uin of the CAM User.
     */
    public /*out*/ readonly uin!: pulumi.Output<number>;
    /**
     * Indicate whether to generate the API secret key or not.
     */
    public readonly useApi!: pulumi.Output<boolean | undefined>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["consoleLogin"] = state ? state.consoleLogin : undefined;
            resourceInputs["countryCode"] = state ? state.countryCode : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["forceDelete"] = state ? state.forceDelete : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["needResetPassword"] = state ? state.needResetPassword : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["phoneNum"] = state ? state.phoneNum : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
            resourceInputs["secretId"] = state ? state.secretId : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["uin"] = state ? state.uin : undefined;
            resourceInputs["useApi"] = state ? state.useApi : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            resourceInputs["consoleLogin"] = args ? args.consoleLogin : undefined;
            resourceInputs["countryCode"] = args ? args.countryCode : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["forceDelete"] = args ? args.forceDelete : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["needResetPassword"] = args ? args.needResetPassword : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["phoneNum"] = args ? args.phoneNum : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["useApi"] = args ? args.useApi : undefined;
            resourceInputs["secretId"] = undefined /*out*/;
            resourceInputs["secretKey"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["uin"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * Indicate whether the CAM user can login to the web console or not.
     */
    consoleLogin?: pulumi.Input<boolean>;
    /**
     * Country code of the phone number, for example: '86'.
     */
    countryCode?: pulumi.Input<string>;
    /**
     * Email of the CAM user.
     */
    email?: pulumi.Input<string>;
    /**
     * Indicate whether to force deletes the CAM user. If set false, the API secret key will be checked and failed when exists; otherwise the user will be deleted directly. Default is false.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * Name of the CAM user.
     */
    name?: pulumi.Input<string>;
    /**
     * Indicate whether the CAM user need to reset the password when first logins.
     */
    needResetPassword?: pulumi.Input<boolean>;
    /**
     * The password of the CAM user. Password should be at least 8 characters and no more than 32 characters, includes uppercase letters, lowercase letters, numbers and special characters. Only required when `consoleLogin` is true. If not set, a random password will be automatically generated.
     */
    password?: pulumi.Input<string>;
    /**
     * Phone number of the CAM user.
     */
    phoneNum?: pulumi.Input<string>;
    /**
     * Remark of the CAM user.
     */
    remark?: pulumi.Input<string>;
    /**
     * Secret ID of the CAM user.
     */
    secretId?: pulumi.Input<string>;
    /**
     * Secret key of the CAM user.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * A list of tags used to associate different resources.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * ID of the CAM user.
     */
    uid?: pulumi.Input<number>;
    /**
     * Uin of the CAM User.
     */
    uin?: pulumi.Input<number>;
    /**
     * Indicate whether to generate the API secret key or not.
     */
    useApi?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Indicate whether the CAM user can login to the web console or not.
     */
    consoleLogin?: pulumi.Input<boolean>;
    /**
     * Country code of the phone number, for example: '86'.
     */
    countryCode?: pulumi.Input<string>;
    /**
     * Email of the CAM user.
     */
    email?: pulumi.Input<string>;
    /**
     * Indicate whether to force deletes the CAM user. If set false, the API secret key will be checked and failed when exists; otherwise the user will be deleted directly. Default is false.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * Name of the CAM user.
     */
    name?: pulumi.Input<string>;
    /**
     * Indicate whether the CAM user need to reset the password when first logins.
     */
    needResetPassword?: pulumi.Input<boolean>;
    /**
     * The password of the CAM user. Password should be at least 8 characters and no more than 32 characters, includes uppercase letters, lowercase letters, numbers and special characters. Only required when `consoleLogin` is true. If not set, a random password will be automatically generated.
     */
    password?: pulumi.Input<string>;
    /**
     * Phone number of the CAM user.
     */
    phoneNum?: pulumi.Input<string>;
    /**
     * Remark of the CAM user.
     */
    remark?: pulumi.Input<string>;
    /**
     * A list of tags used to associate different resources.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Indicate whether to generate the API secret key or not.
     */
    useApi?: pulumi.Input<boolean>;
}
