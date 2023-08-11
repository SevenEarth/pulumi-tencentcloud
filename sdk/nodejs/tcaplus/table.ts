// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this resource to create TcaplusDB table.
 *
 * ## Example Usage
 * ### Create a tcaplus database table
 *
 * The tcaplus database table should be pre-defined in the idl file.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const config = new pulumi.Config();
 * const availabilityZone = config.get("availabilityZone") || "ap-guangzhou-3";
 * const vpc = tencentcloud.Vpc.getSubnets({
 *     isDefault: true,
 *     availabilityZone: availabilityZone,
 * });
 * const vpcId = vpc.then(vpc => vpc.instanceLists?[0]?.vpcId);
 * const subnetId = vpc.then(vpc => vpc.instanceLists?[0]?.subnetId);
 * const exampleCluster = new tencentcloud.tcaplus.Cluster("exampleCluster", {
 *     idlType: "PROTO",
 *     clusterName: "tf_example_tcaplus_cluster",
 *     vpcId: vpcId,
 *     subnetId: subnetId,
 *     password: "your_pw_123111",
 *     oldPasswordExpireLast: 3600,
 * });
 * const exampleTablegroup = new tencentcloud.tcaplus.Tablegroup("exampleTablegroup", {
 *     clusterId: exampleCluster.id,
 *     tablegroupName: "tf_example_group_name",
 * });
 * const exampleIdl = new tencentcloud.tcaplus.Idl("exampleIdl", {
 *     clusterId: exampleCluster.id,
 *     tablegroupId: exampleTablegroup.id,
 *     fileName: "tf_example_tcaplus_idl",
 *     fileType: "PROTO",
 *     fileExtType: "proto",
 *     fileContent: `    syntax = "proto2";
 *     package myTcaplusTable;
 *     import "tcaplusservice.optionv1.proto";
 *     message example_table { # refer the table name
 *         option(tcaplusservice.tcaplus_primary_key) = "uin,name,region";
 *         required int64 uin = 1;
 *         required string name = 2;
 *         required int32 region = 3;
 *         required int32 gamesvrid = 4;
 *         optional int32 logintime = 5 [default = 1];
 *         repeated int64 lockid = 6 [packed = true];
 *         optional bool is_available = 7 [default = false];
 *         optional pay_info pay = 8;
 *     }
 *
 *     message pay_info {
 *         required int64 pay_id = 1;
 *         optional uint64 total_money = 2;
 *         optional uint64 pay_times = 3;
 *         optional pay_auth_info auth = 4;
 *         message pay_auth_info {
 *             required string pay_keys = 1;
 *             optional int64 update_time = 2;
 *         }
 *     }
 * `,
 * });
 * const exampleTable = new tencentcloud.tcaplus.Table("exampleTable", {
 *     clusterId: exampleCluster.id,
 *     tablegroupId: exampleTablegroup.id,
 *     tableName: "example_table",
 *     tableType: "GENERIC",
 *     description: "test",
 *     idlId: exampleIdl.id,
 *     tableIdlType: "PROTO",
 *     reservedReadCu: 1000,
 *     reservedWriteCu: 20,
 *     reservedVolume: 1,
 * });
 * ```
 */
export class Table extends pulumi.CustomResource {
    /**
     * Get an existing Table resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TableState, opts?: pulumi.CustomResourceOptions): Table {
        return new Table(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tcaplus/table:Table';

    /**
     * Returns true if the given object is an instance of Table.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Table {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Table.__pulumiType;
    }

    /**
     * ID of the TcaplusDB cluster to which the table belongs.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Create time of the TcaplusDB table.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Description of the TcaplusDB table.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Error messages for creating TcaplusDB table.
     */
    public /*out*/ readonly error!: pulumi.Output<string>;
    /**
     * ID of the IDL File.
     */
    public readonly idlId!: pulumi.Output<string>;
    /**
     * Reserved read capacity units of the TcaplusDB table.
     */
    public readonly reservedReadCu!: pulumi.Output<number>;
    /**
     * Reserved storage capacity of the TcaplusDB table (unit: GB).
     */
    public readonly reservedVolume!: pulumi.Output<number>;
    /**
     * Reserved write capacity units of the TcaplusDB table.
     */
    public readonly reservedWriteCu!: pulumi.Output<number>;
    /**
     * Status of the TcaplusDB table.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * IDL type of the TcaplusDB table. Valid values: `PROTO` and `TDR`.
     */
    public readonly tableIdlType!: pulumi.Output<string>;
    /**
     * Name of the TcaplusDB table.
     */
    public readonly tableName!: pulumi.Output<string>;
    /**
     * Size of the TcaplusDB table.
     */
    public /*out*/ readonly tableSize!: pulumi.Output<number>;
    /**
     * Type of the TcaplusDB table. Valid values are `GENERIC` and `LIST`.
     */
    public readonly tableType!: pulumi.Output<string>;
    /**
     * ID of the table group to which the table belongs.
     */
    public readonly tablegroupId!: pulumi.Output<string>;

    /**
     * Create a Table resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TableArgs | TableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TableState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["error"] = state ? state.error : undefined;
            resourceInputs["idlId"] = state ? state.idlId : undefined;
            resourceInputs["reservedReadCu"] = state ? state.reservedReadCu : undefined;
            resourceInputs["reservedVolume"] = state ? state.reservedVolume : undefined;
            resourceInputs["reservedWriteCu"] = state ? state.reservedWriteCu : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tableIdlType"] = state ? state.tableIdlType : undefined;
            resourceInputs["tableName"] = state ? state.tableName : undefined;
            resourceInputs["tableSize"] = state ? state.tableSize : undefined;
            resourceInputs["tableType"] = state ? state.tableType : undefined;
            resourceInputs["tablegroupId"] = state ? state.tablegroupId : undefined;
        } else {
            const args = argsOrState as TableArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.idlId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'idlId'");
            }
            if ((!args || args.reservedReadCu === undefined) && !opts.urn) {
                throw new Error("Missing required property 'reservedReadCu'");
            }
            if ((!args || args.reservedVolume === undefined) && !opts.urn) {
                throw new Error("Missing required property 'reservedVolume'");
            }
            if ((!args || args.reservedWriteCu === undefined) && !opts.urn) {
                throw new Error("Missing required property 'reservedWriteCu'");
            }
            if ((!args || args.tableIdlType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tableIdlType'");
            }
            if ((!args || args.tableName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tableName'");
            }
            if ((!args || args.tableType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tableType'");
            }
            if ((!args || args.tablegroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tablegroupId'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["idlId"] = args ? args.idlId : undefined;
            resourceInputs["reservedReadCu"] = args ? args.reservedReadCu : undefined;
            resourceInputs["reservedVolume"] = args ? args.reservedVolume : undefined;
            resourceInputs["reservedWriteCu"] = args ? args.reservedWriteCu : undefined;
            resourceInputs["tableIdlType"] = args ? args.tableIdlType : undefined;
            resourceInputs["tableName"] = args ? args.tableName : undefined;
            resourceInputs["tableType"] = args ? args.tableType : undefined;
            resourceInputs["tablegroupId"] = args ? args.tablegroupId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tableSize"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Table.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Table resources.
 */
export interface TableState {
    /**
     * ID of the TcaplusDB cluster to which the table belongs.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Create time of the TcaplusDB table.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Description of the TcaplusDB table.
     */
    description?: pulumi.Input<string>;
    /**
     * Error messages for creating TcaplusDB table.
     */
    error?: pulumi.Input<string>;
    /**
     * ID of the IDL File.
     */
    idlId?: pulumi.Input<string>;
    /**
     * Reserved read capacity units of the TcaplusDB table.
     */
    reservedReadCu?: pulumi.Input<number>;
    /**
     * Reserved storage capacity of the TcaplusDB table (unit: GB).
     */
    reservedVolume?: pulumi.Input<number>;
    /**
     * Reserved write capacity units of the TcaplusDB table.
     */
    reservedWriteCu?: pulumi.Input<number>;
    /**
     * Status of the TcaplusDB table.
     */
    status?: pulumi.Input<string>;
    /**
     * IDL type of the TcaplusDB table. Valid values: `PROTO` and `TDR`.
     */
    tableIdlType?: pulumi.Input<string>;
    /**
     * Name of the TcaplusDB table.
     */
    tableName?: pulumi.Input<string>;
    /**
     * Size of the TcaplusDB table.
     */
    tableSize?: pulumi.Input<number>;
    /**
     * Type of the TcaplusDB table. Valid values are `GENERIC` and `LIST`.
     */
    tableType?: pulumi.Input<string>;
    /**
     * ID of the table group to which the table belongs.
     */
    tablegroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Table resource.
 */
export interface TableArgs {
    /**
     * ID of the TcaplusDB cluster to which the table belongs.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Description of the TcaplusDB table.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of the IDL File.
     */
    idlId: pulumi.Input<string>;
    /**
     * Reserved read capacity units of the TcaplusDB table.
     */
    reservedReadCu: pulumi.Input<number>;
    /**
     * Reserved storage capacity of the TcaplusDB table (unit: GB).
     */
    reservedVolume: pulumi.Input<number>;
    /**
     * Reserved write capacity units of the TcaplusDB table.
     */
    reservedWriteCu: pulumi.Input<number>;
    /**
     * IDL type of the TcaplusDB table. Valid values: `PROTO` and `TDR`.
     */
    tableIdlType: pulumi.Input<string>;
    /**
     * Name of the TcaplusDB table.
     */
    tableName: pulumi.Input<string>;
    /**
     * Type of the TcaplusDB table. Valid values are `GENERIC` and `LIST`.
     */
    tableType: pulumi.Input<string>;
    /**
     * ID of the table group to which the table belongs.
     */
    tablegroupId: pulumi.Input<string>;
}
