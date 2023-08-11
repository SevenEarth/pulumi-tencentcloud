// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this resource to create TcaplusDB IDL file.
 *
 * ## Example Usage
 * ### Create a tcaplus database idl file
 *
 * The file will be with a specified cluster and tablegroup.
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
 * const main = new tencentcloud.tcaplus.Idl("main", {
 *     clusterId: exampleCluster.id,
 *     tablegroupId: exampleTablegroup.id,
 *     fileName: "tf_example_tcaplus_idl",
 *     fileType: "PROTO",
 *     fileExtType: "proto",
 *     fileContent: `    syntax = "proto2";
 *     package myTcaplusTable;
 *     import "tcaplusservice.optionv1.proto";
 *     message tb_online {
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
 * ```
 */
export class Idl extends pulumi.CustomResource {
    /**
     * Get an existing Idl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IdlState, opts?: pulumi.CustomResourceOptions): Idl {
        return new Idl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tcaplus/idl:Idl';

    /**
     * Returns true if the given object is an instance of Idl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Idl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Idl.__pulumiType;
    }

    /**
     * ID of the TcaplusDB cluster to which the table group belongs.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * IDL file content of the TcaplusDB table.
     */
    public readonly fileContent!: pulumi.Output<string>;
    /**
     * File ext type of the IDL file. If `fileType` is `PROTO`, `fileExtType` must be 'proto'; If `fileType` is `TDR`, `fileExtType` must be 'xml'.
     */
    public readonly fileExtType!: pulumi.Output<string>;
    /**
     * Name of the IDL file.
     */
    public readonly fileName!: pulumi.Output<string>;
    /**
     * Type of the IDL file. Valid values are PROTO and TDR.
     */
    public readonly fileType!: pulumi.Output<string>;
    /**
     * Table info of the IDL.
     */
    public /*out*/ readonly tableInfos!: pulumi.Output<outputs.Tcaplus.IdlTableInfo[]>;
    /**
     * ID of the table group to which the IDL file belongs.
     */
    public readonly tablegroupId!: pulumi.Output<string>;

    /**
     * Create a Idl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IdlArgs | IdlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IdlState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["fileContent"] = state ? state.fileContent : undefined;
            resourceInputs["fileExtType"] = state ? state.fileExtType : undefined;
            resourceInputs["fileName"] = state ? state.fileName : undefined;
            resourceInputs["fileType"] = state ? state.fileType : undefined;
            resourceInputs["tableInfos"] = state ? state.tableInfos : undefined;
            resourceInputs["tablegroupId"] = state ? state.tablegroupId : undefined;
        } else {
            const args = argsOrState as IdlArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.fileContent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileContent'");
            }
            if ((!args || args.fileExtType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileExtType'");
            }
            if ((!args || args.fileName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileName'");
            }
            if ((!args || args.fileType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileType'");
            }
            if ((!args || args.tablegroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tablegroupId'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["fileContent"] = args ? args.fileContent : undefined;
            resourceInputs["fileExtType"] = args ? args.fileExtType : undefined;
            resourceInputs["fileName"] = args ? args.fileName : undefined;
            resourceInputs["fileType"] = args ? args.fileType : undefined;
            resourceInputs["tablegroupId"] = args ? args.tablegroupId : undefined;
            resourceInputs["tableInfos"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Idl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Idl resources.
 */
export interface IdlState {
    /**
     * ID of the TcaplusDB cluster to which the table group belongs.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * IDL file content of the TcaplusDB table.
     */
    fileContent?: pulumi.Input<string>;
    /**
     * File ext type of the IDL file. If `fileType` is `PROTO`, `fileExtType` must be 'proto'; If `fileType` is `TDR`, `fileExtType` must be 'xml'.
     */
    fileExtType?: pulumi.Input<string>;
    /**
     * Name of the IDL file.
     */
    fileName?: pulumi.Input<string>;
    /**
     * Type of the IDL file. Valid values are PROTO and TDR.
     */
    fileType?: pulumi.Input<string>;
    /**
     * Table info of the IDL.
     */
    tableInfos?: pulumi.Input<pulumi.Input<inputs.Tcaplus.IdlTableInfo>[]>;
    /**
     * ID of the table group to which the IDL file belongs.
     */
    tablegroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Idl resource.
 */
export interface IdlArgs {
    /**
     * ID of the TcaplusDB cluster to which the table group belongs.
     */
    clusterId: pulumi.Input<string>;
    /**
     * IDL file content of the TcaplusDB table.
     */
    fileContent: pulumi.Input<string>;
    /**
     * File ext type of the IDL file. If `fileType` is `PROTO`, `fileExtType` must be 'proto'; If `fileType` is `TDR`, `fileExtType` must be 'xml'.
     */
    fileExtType: pulumi.Input<string>;
    /**
     * Name of the IDL file.
     */
    fileName: pulumi.Input<string>;
    /**
     * Type of the IDL file. Valid values are PROTO and TDR.
     */
    fileType: pulumi.Input<string>;
    /**
     * ID of the table group to which the IDL file belongs.
     */
    tablegroupId: pulumi.Input<string>;
}
