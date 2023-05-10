// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mariadb account
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const account = new tencentcloud.Mariadb.Account("account", {
 *     description: "desc",
 *     host: "10.101.202.22",
 *     instanceId: "tdsql-4pzs5b67",
 *     password: "Password123.",
 *     readOnly: 0,
 *     userName: "account-test",
 * });
 * ```
 *
 * ## Import
 *
 * mariadb account can be imported using the instance_id#user_name#host, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Mariadb/account:Account account tdsql-4pzs5b67#account-test#10.101.202.22
 * ```
 */
export class Account extends pulumi.CustomResource {
    /**
     * Get an existing Account resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountState, opts?: pulumi.CustomResourceOptions): Account {
        return new Account(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mariadb/account:Account';

    /**
     * Returns true if the given object is an instance of Account.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Account {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Account.__pulumiType;
    }

    /**
     * account description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * host.
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * instance id.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * account password.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * wether account is read only, 0 means not a read only account.
     */
    public readonly readOnly!: pulumi.Output<number | undefined>;
    /**
     * user name.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a Account resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountArgs | AccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["readOnly"] = state ? state.readOnly : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            if ((!args || args.host === undefined) && !opts.urn) {
                throw new Error("Missing required property 'host'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["readOnly"] = args ? args.readOnly : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Account.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Account resources.
 */
export interface AccountState {
    /**
     * account description.
     */
    description?: pulumi.Input<string>;
    /**
     * host.
     */
    host?: pulumi.Input<string>;
    /**
     * instance id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * account password.
     */
    password?: pulumi.Input<string>;
    /**
     * wether account is read only, 0 means not a read only account.
     */
    readOnly?: pulumi.Input<number>;
    /**
     * user name.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * account description.
     */
    description?: pulumi.Input<string>;
    /**
     * host.
     */
    host: pulumi.Input<string>;
    /**
     * instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * account password.
     */
    password: pulumi.Input<string>;
    /**
     * wether account is read only, 0 means not a read only account.
     */
    readOnly?: pulumi.Input<number>;
    /**
     * user name.
     */
    userName: pulumi.Input<string>;
}
