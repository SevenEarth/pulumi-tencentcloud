// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a dcdb dbInstance
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const dbInstance = new tencentcloud.dcdb.DbInstance("dbInstance", {
 *     instanceName: "test_dcdb_db_instance",
 *     zones: ["ap-guangzhou-5"],
 *     period: 1,
 *     shardMemory: 2,
 *     shardStorage: 10,
 *     shardNodeCount: 2,
 *     shardCount: 2,
 *     vpcId: local.vpc_id,
 *     subnetId: local.subnet_id,
 *     dbVersionId: "8.0",
 *     resourceTags: [{
 *         tagKey: "aaa",
 *         tagValue: "bbb",
 *     }],
 *     initParams: [
 *         {
 *             param: "character_set_server",
 *             value: "utf8mb4",
 *         },
 *         {
 *             param: "lower_case_table_names",
 *             value: "1",
 *         },
 *         {
 *             param: "sync_mode",
 *             value: "2",
 *         },
 *         {
 *             param: "innodb_page_size",
 *             value: "16384",
 *         },
 *     ],
 *     securityGroupIds: [local.sg_id],
 * });
 * ```
 *
 * ## Import
 *
 * dcdb db_instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Dcdb/dbInstance:DbInstance db_instance db_instance_id
 * ```
 */
export class DbInstance extends pulumi.CustomResource {
    /**
     * Get an existing DbInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DbInstanceState, opts?: pulumi.CustomResourceOptions): DbInstance {
        return new DbInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dcdb/dbInstance:DbInstance';

    /**
     * Returns true if the given object is an instance of DbInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DbInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DbInstance.__pulumiType;
    }

    /**
     * &amp;quot;Automatic renewal flag, 0 means the default state (the user has not set it, that is, the initial state is manual renewal, and the user has activated the prepaid non-stop privilege and will also perform automatic renewal).&amp;quot;&amp;quot;1 means automatic renewal, 2 means no automatic renewal (user setting).&amp;quot;&amp;quot;if the business has no concept of renewal or automatic renewal is not required, it needs to be set to 0.&amp;quot;.
     */
    public readonly autoRenewFlag!: pulumi.Output<number | undefined>;
    /**
     * Whether to automatically use vouchers for payment, not used by default.
     */
    public readonly autoVoucher!: pulumi.Output<boolean | undefined>;
    /**
     * &amp;quot;Database engine version, currently available: 8.0.18, 10.1.9, 5.7.17.&amp;quot;&amp;quot;8.0.18 - MySQL 8.0.18;&amp;quot;&amp;quot;10.1.9 - Mariadb 10.1.9;&amp;quot;&amp;quot;5.7.17 - Percona 5.7.17&amp;quot;&amp;quot;If not filled, the default is 5.7.17, which means Percona 5.7.17.&amp;quot;.
     */
    public readonly dbVersionId!: pulumi.Output<string | undefined>;
    /**
     * DCN source instance ID.
     */
    public readonly dcnInstanceId!: pulumi.Output<string | undefined>;
    /**
     * DCN source region.
     */
    public readonly dcnRegion!: pulumi.Output<string | undefined>;
    /**
     * &amp;quot;parameter list. The optional values of this interface are:&amp;quot;&amp;quot;character_set_server (character set, must be passed),&amp;quot;&amp;quot;lower_case_table_names (table name is case sensitive, must be passed, 0 - sensitive; 1 - insensitive),&amp;quot;&amp;quot;innodb_page_size (innodb data page, default 16K),&amp;quot;&amp;quot;sync_mode ( Synchronous mode: 0 - asynchronous; 1 - strong synchronous; 2 - strong synchronous degenerate. The default is strong synchronous degenerate)&amp;quot;.
     */
    public readonly initParams!: pulumi.Output<outputs.Dcdb.DbInstanceInitParam[] | undefined>;
    /**
     * Instance name, you can set the name of the instance independently through this field.
     */
    public readonly instanceName!: pulumi.Output<string | undefined>;
    /**
     * Whether to support IPv6.
     */
    public readonly ipv6Flag!: pulumi.Output<number | undefined>;
    /**
     * The length of time you want to buy, unit: month.
     */
    public readonly period!: pulumi.Output<number>;
    /**
     * Project ID, which can be obtained by viewing the project list, if not passed, it will be associated with the default project.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * Array of tag key-value pairs.
     */
    public readonly resourceTags!: pulumi.Output<outputs.Dcdb.DbInstanceResourceTag[] | undefined>;
    /**
     * Security group ids, the security group can be passed in the form of an array, compatible with the previous SecurityGroupId parameter.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The number of instance fragments, the optional range is 2-8, and new fragments can be added to a maximum of 64 fragments by upgrading the instance.
     */
    public readonly shardCount!: pulumi.Output<number>;
    /**
     * &amp;quot;Shard memory size, unit: GB, can pass DescribeShardSpec&amp;quot;&amp;quot;Query the instance specification to obtain.&amp;quot;.
     */
    public readonly shardMemory!: pulumi.Output<number>;
    /**
     * &amp;quot;Number of single shard nodes, can pass DescribeShardSpec&amp;quot;&amp;quot;Query the instance specification to obtain.&amp;quot;.
     */
    public readonly shardNodeCount!: pulumi.Output<number>;
    /**
     * &amp;quot;Shard storage size, unit: GB, can pass DescribeShardSpec&amp;quot;&amp;quot;Query the instance specification to obtain.&amp;quot;.
     */
    public readonly shardStorage!: pulumi.Output<number>;
    /**
     * Virtual private network subnet ID, required when VpcId is not empty.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * Voucher ID list, currently only supports specifying one voucher.
     */
    public readonly voucherIds!: pulumi.Output<string[] | undefined>;
    /**
     * Virtual private network ID, if not passed or passed empty, it means that it is created as a basic network.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;
    /**
     * &amp;quot;The availability zone distribution of shard nodes can be filled with up to two availability zones. When the shard specification is one master and two slaves, two of the nodes are in the first availability zone.&amp;quot;&amp;quot;Note that the current availability zone that can be sold needs to be pulled through the DescribeDCDBSaleInfo interface.&amp;quot;.
     */
    public readonly zones!: pulumi.Output<string[]>;

    /**
     * Create a DbInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DbInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DbInstanceArgs | DbInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DbInstanceState | undefined;
            resourceInputs["autoRenewFlag"] = state ? state.autoRenewFlag : undefined;
            resourceInputs["autoVoucher"] = state ? state.autoVoucher : undefined;
            resourceInputs["dbVersionId"] = state ? state.dbVersionId : undefined;
            resourceInputs["dcnInstanceId"] = state ? state.dcnInstanceId : undefined;
            resourceInputs["dcnRegion"] = state ? state.dcnRegion : undefined;
            resourceInputs["initParams"] = state ? state.initParams : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["ipv6Flag"] = state ? state.ipv6Flag : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["resourceTags"] = state ? state.resourceTags : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["shardCount"] = state ? state.shardCount : undefined;
            resourceInputs["shardMemory"] = state ? state.shardMemory : undefined;
            resourceInputs["shardNodeCount"] = state ? state.shardNodeCount : undefined;
            resourceInputs["shardStorage"] = state ? state.shardStorage : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["voucherIds"] = state ? state.voucherIds : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as DbInstanceArgs | undefined;
            if ((!args || args.period === undefined) && !opts.urn) {
                throw new Error("Missing required property 'period'");
            }
            if ((!args || args.shardCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shardCount'");
            }
            if ((!args || args.shardMemory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shardMemory'");
            }
            if ((!args || args.shardNodeCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shardNodeCount'");
            }
            if ((!args || args.shardStorage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shardStorage'");
            }
            if ((!args || args.zones === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zones'");
            }
            resourceInputs["autoRenewFlag"] = args ? args.autoRenewFlag : undefined;
            resourceInputs["autoVoucher"] = args ? args.autoVoucher : undefined;
            resourceInputs["dbVersionId"] = args ? args.dbVersionId : undefined;
            resourceInputs["dcnInstanceId"] = args ? args.dcnInstanceId : undefined;
            resourceInputs["dcnRegion"] = args ? args.dcnRegion : undefined;
            resourceInputs["initParams"] = args ? args.initParams : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["ipv6Flag"] = args ? args.ipv6Flag : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["resourceTags"] = args ? args.resourceTags : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["shardCount"] = args ? args.shardCount : undefined;
            resourceInputs["shardMemory"] = args ? args.shardMemory : undefined;
            resourceInputs["shardNodeCount"] = args ? args.shardNodeCount : undefined;
            resourceInputs["shardStorage"] = args ? args.shardStorage : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["voucherIds"] = args ? args.voucherIds : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["zones"] = args ? args.zones : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DbInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DbInstance resources.
 */
export interface DbInstanceState {
    /**
     * &amp;quot;Automatic renewal flag, 0 means the default state (the user has not set it, that is, the initial state is manual renewal, and the user has activated the prepaid non-stop privilege and will also perform automatic renewal).&amp;quot;&amp;quot;1 means automatic renewal, 2 means no automatic renewal (user setting).&amp;quot;&amp;quot;if the business has no concept of renewal or automatic renewal is not required, it needs to be set to 0.&amp;quot;.
     */
    autoRenewFlag?: pulumi.Input<number>;
    /**
     * Whether to automatically use vouchers for payment, not used by default.
     */
    autoVoucher?: pulumi.Input<boolean>;
    /**
     * &amp;quot;Database engine version, currently available: 8.0.18, 10.1.9, 5.7.17.&amp;quot;&amp;quot;8.0.18 - MySQL 8.0.18;&amp;quot;&amp;quot;10.1.9 - Mariadb 10.1.9;&amp;quot;&amp;quot;5.7.17 - Percona 5.7.17&amp;quot;&amp;quot;If not filled, the default is 5.7.17, which means Percona 5.7.17.&amp;quot;.
     */
    dbVersionId?: pulumi.Input<string>;
    /**
     * DCN source instance ID.
     */
    dcnInstanceId?: pulumi.Input<string>;
    /**
     * DCN source region.
     */
    dcnRegion?: pulumi.Input<string>;
    /**
     * &amp;quot;parameter list. The optional values of this interface are:&amp;quot;&amp;quot;character_set_server (character set, must be passed),&amp;quot;&amp;quot;lower_case_table_names (table name is case sensitive, must be passed, 0 - sensitive; 1 - insensitive),&amp;quot;&amp;quot;innodb_page_size (innodb data page, default 16K),&amp;quot;&amp;quot;sync_mode ( Synchronous mode: 0 - asynchronous; 1 - strong synchronous; 2 - strong synchronous degenerate. The default is strong synchronous degenerate)&amp;quot;.
     */
    initParams?: pulumi.Input<pulumi.Input<inputs.Dcdb.DbInstanceInitParam>[]>;
    /**
     * Instance name, you can set the name of the instance independently through this field.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Whether to support IPv6.
     */
    ipv6Flag?: pulumi.Input<number>;
    /**
     * The length of time you want to buy, unit: month.
     */
    period?: pulumi.Input<number>;
    /**
     * Project ID, which can be obtained by viewing the project list, if not passed, it will be associated with the default project.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Array of tag key-value pairs.
     */
    resourceTags?: pulumi.Input<pulumi.Input<inputs.Dcdb.DbInstanceResourceTag>[]>;
    /**
     * Security group ids, the security group can be passed in the form of an array, compatible with the previous SecurityGroupId parameter.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of instance fragments, the optional range is 2-8, and new fragments can be added to a maximum of 64 fragments by upgrading the instance.
     */
    shardCount?: pulumi.Input<number>;
    /**
     * &amp;quot;Shard memory size, unit: GB, can pass DescribeShardSpec&amp;quot;&amp;quot;Query the instance specification to obtain.&amp;quot;.
     */
    shardMemory?: pulumi.Input<number>;
    /**
     * &amp;quot;Number of single shard nodes, can pass DescribeShardSpec&amp;quot;&amp;quot;Query the instance specification to obtain.&amp;quot;.
     */
    shardNodeCount?: pulumi.Input<number>;
    /**
     * &amp;quot;Shard storage size, unit: GB, can pass DescribeShardSpec&amp;quot;&amp;quot;Query the instance specification to obtain.&amp;quot;.
     */
    shardStorage?: pulumi.Input<number>;
    /**
     * Virtual private network subnet ID, required when VpcId is not empty.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Voucher ID list, currently only supports specifying one voucher.
     */
    voucherIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Virtual private network ID, if not passed or passed empty, it means that it is created as a basic network.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * &amp;quot;The availability zone distribution of shard nodes can be filled with up to two availability zones. When the shard specification is one master and two slaves, two of the nodes are in the first availability zone.&amp;quot;&amp;quot;Note that the current availability zone that can be sold needs to be pulled through the DescribeDCDBSaleInfo interface.&amp;quot;.
     */
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a DbInstance resource.
 */
export interface DbInstanceArgs {
    /**
     * &amp;quot;Automatic renewal flag, 0 means the default state (the user has not set it, that is, the initial state is manual renewal, and the user has activated the prepaid non-stop privilege and will also perform automatic renewal).&amp;quot;&amp;quot;1 means automatic renewal, 2 means no automatic renewal (user setting).&amp;quot;&amp;quot;if the business has no concept of renewal or automatic renewal is not required, it needs to be set to 0.&amp;quot;.
     */
    autoRenewFlag?: pulumi.Input<number>;
    /**
     * Whether to automatically use vouchers for payment, not used by default.
     */
    autoVoucher?: pulumi.Input<boolean>;
    /**
     * &amp;quot;Database engine version, currently available: 8.0.18, 10.1.9, 5.7.17.&amp;quot;&amp;quot;8.0.18 - MySQL 8.0.18;&amp;quot;&amp;quot;10.1.9 - Mariadb 10.1.9;&amp;quot;&amp;quot;5.7.17 - Percona 5.7.17&amp;quot;&amp;quot;If not filled, the default is 5.7.17, which means Percona 5.7.17.&amp;quot;.
     */
    dbVersionId?: pulumi.Input<string>;
    /**
     * DCN source instance ID.
     */
    dcnInstanceId?: pulumi.Input<string>;
    /**
     * DCN source region.
     */
    dcnRegion?: pulumi.Input<string>;
    /**
     * &amp;quot;parameter list. The optional values of this interface are:&amp;quot;&amp;quot;character_set_server (character set, must be passed),&amp;quot;&amp;quot;lower_case_table_names (table name is case sensitive, must be passed, 0 - sensitive; 1 - insensitive),&amp;quot;&amp;quot;innodb_page_size (innodb data page, default 16K),&amp;quot;&amp;quot;sync_mode ( Synchronous mode: 0 - asynchronous; 1 - strong synchronous; 2 - strong synchronous degenerate. The default is strong synchronous degenerate)&amp;quot;.
     */
    initParams?: pulumi.Input<pulumi.Input<inputs.Dcdb.DbInstanceInitParam>[]>;
    /**
     * Instance name, you can set the name of the instance independently through this field.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Whether to support IPv6.
     */
    ipv6Flag?: pulumi.Input<number>;
    /**
     * The length of time you want to buy, unit: month.
     */
    period: pulumi.Input<number>;
    /**
     * Project ID, which can be obtained by viewing the project list, if not passed, it will be associated with the default project.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Array of tag key-value pairs.
     */
    resourceTags?: pulumi.Input<pulumi.Input<inputs.Dcdb.DbInstanceResourceTag>[]>;
    /**
     * Security group ids, the security group can be passed in the form of an array, compatible with the previous SecurityGroupId parameter.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of instance fragments, the optional range is 2-8, and new fragments can be added to a maximum of 64 fragments by upgrading the instance.
     */
    shardCount: pulumi.Input<number>;
    /**
     * &amp;quot;Shard memory size, unit: GB, can pass DescribeShardSpec&amp;quot;&amp;quot;Query the instance specification to obtain.&amp;quot;.
     */
    shardMemory: pulumi.Input<number>;
    /**
     * &amp;quot;Number of single shard nodes, can pass DescribeShardSpec&amp;quot;&amp;quot;Query the instance specification to obtain.&amp;quot;.
     */
    shardNodeCount: pulumi.Input<number>;
    /**
     * &amp;quot;Shard storage size, unit: GB, can pass DescribeShardSpec&amp;quot;&amp;quot;Query the instance specification to obtain.&amp;quot;.
     */
    shardStorage: pulumi.Input<number>;
    /**
     * Virtual private network subnet ID, required when VpcId is not empty.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Voucher ID list, currently only supports specifying one voucher.
     */
    voucherIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Virtual private network ID, if not passed or passed empty, it means that it is created as a basic network.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * &amp;quot;The availability zone distribution of shard nodes can be filled with up to two availability zones. When the shard specification is one master and two slaves, two of the nodes are in the first availability zone.&amp;quot;&amp;quot;Note that the current availability zone that can be sold needs to be pulled through the DescribeDCDBSaleInfo interface.&amp;quot;.
     */
    zones: pulumi.Input<pulumi.Input<string>[]>;
}