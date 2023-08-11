// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this resource to create SQL Server account DB attachment
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const zones = tencentcloud.Availability.getZonesByProduct({
 *     product: "sqlserver",
 * });
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     availabilityZone: zones.then(zones => zones.zones?[4]?.name),
 *     vpcId: vpc.id,
 *     cidrBlock: "10.0.0.0/16",
 *     isMulticast: false,
 * });
 * const securityGroup = new tencentcloud.security.Group("securityGroup", {description: "desc."});
 * const exampleBasicInstance = new tencentcloud.sqlserver.BasicInstance("exampleBasicInstance", {
 *     availabilityZone: zones.then(zones => zones.zones?[4]?.name),
 *     chargeType: "POSTPAID_BY_HOUR",
 *     vpcId: vpc.id,
 *     subnetId: subnet.id,
 *     projectId: 0,
 *     memory: 4,
 *     storage: 100,
 *     cpu: 2,
 *     machineType: "CLOUD_PREMIUM",
 *     maintenanceWeekSets: [
 *         1,
 *         2,
 *         3,
 *     ],
 *     maintenanceStartTime: "09:00",
 *     maintenanceTimeSpan: 3,
 *     securityGroups: [securityGroup.id],
 *     tags: {
 *         test: "test",
 *     },
 * });
 * const exampleDb = new tencentcloud.sqlserver.Db("exampleDb", {
 *     instanceId: exampleBasicInstance.id,
 *     charset: "Chinese_PRC_BIN",
 *     remark: "test-remark",
 * });
 * const exampleAccount = new tencentcloud.sqlserver.Account("exampleAccount", {
 *     instanceId: exampleBasicInstance.id,
 *     password: "Qwer@234",
 *     remark: "test-remark",
 * });
 * const exampleAccountDbAttachment = new tencentcloud.sqlserver.AccountDbAttachment("exampleAccountDbAttachment", {
 *     instanceId: exampleBasicInstance.id,
 *     accountName: exampleAccount.name,
 *     dbName: exampleDb.name,
 *     privilege: "ReadWrite",
 * });
 * ```
 *
 * ## Import
 *
 * SQL Server account DB attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Sqlserver/accountDbAttachment:AccountDbAttachment example mssql-3cdq7kx5#tf_example_account#tf_example_db
 * ```
 */
export class AccountDbAttachment extends pulumi.CustomResource {
    /**
     * Get an existing AccountDbAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountDbAttachmentState, opts?: pulumi.CustomResourceOptions): AccountDbAttachment {
        return new AccountDbAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Sqlserver/accountDbAttachment:AccountDbAttachment';

    /**
     * Returns true if the given object is an instance of AccountDbAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountDbAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountDbAttachment.__pulumiType;
    }

    /**
     * SQL Server account name.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * SQL Server DB name.
     */
    public readonly dbName!: pulumi.Output<string>;
    /**
     * SQL Server instance ID that the account belongs to.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Privilege of the account on DB. Valid values: `ReadOnly`, `ReadWrite`.
     */
    public readonly privilege!: pulumi.Output<string>;

    /**
     * Create a AccountDbAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountDbAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountDbAttachmentArgs | AccountDbAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountDbAttachmentState | undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["dbName"] = state ? state.dbName : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["privilege"] = state ? state.privilege : undefined;
        } else {
            const args = argsOrState as AccountDbAttachmentArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.dbName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbName'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.privilege === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privilege'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["dbName"] = args ? args.dbName : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["privilege"] = args ? args.privilege : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccountDbAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccountDbAttachment resources.
 */
export interface AccountDbAttachmentState {
    /**
     * SQL Server account name.
     */
    accountName?: pulumi.Input<string>;
    /**
     * SQL Server DB name.
     */
    dbName?: pulumi.Input<string>;
    /**
     * SQL Server instance ID that the account belongs to.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Privilege of the account on DB. Valid values: `ReadOnly`, `ReadWrite`.
     */
    privilege?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccountDbAttachment resource.
 */
export interface AccountDbAttachmentArgs {
    /**
     * SQL Server account name.
     */
    accountName: pulumi.Input<string>;
    /**
     * SQL Server DB name.
     */
    dbName: pulumi.Input<string>;
    /**
     * SQL Server instance ID that the account belongs to.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Privilege of the account on DB. Valid values: `ReadOnly`, `ReadWrite`.
     */
    privilege: pulumi.Input<string>;
}
