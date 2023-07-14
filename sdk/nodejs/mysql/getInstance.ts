// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a MySQL instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const database = pulumi.output(tencentcloud.Mysql.getInstance({
 *     mysqlId: "terraform-test-local-database",
 *     resultOutputFile: "mytestpath",
 * }));
 * ```
 */
export function getInstance(args?: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Mysql/getInstance:getInstance", {
        "chargeType": args.chargeType,
        "engineVersion": args.engineVersion,
        "initFlag": args.initFlag,
        "instanceName": args.instanceName,
        "instanceRole": args.instanceRole,
        "limit": args.limit,
        "mysqlId": args.mysqlId,
        "offset": args.offset,
        "payType": args.payType,
        "resultOutputFile": args.resultOutputFile,
        "securityGroupId": args.securityGroupId,
        "status": args.status,
        "withDr": args.withDr,
        "withMaster": args.withMaster,
        "withRo": args.withRo,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceArgs {
    /**
     * Pay type of instance, valid values are `PREPAID` and `POSTPAID`.
     */
    chargeType?: string;
    /**
     * The version number of the database engine to use. Supported versions include 5.5/5.6/5.7/8.0.
     */
    engineVersion?: string;
    /**
     * Initialization mark. Available values: `0` - Uninitialized; `1` - Initialized.
     */
    initFlag?: number;
    /**
     * Name of mysql instance.
     */
    instanceName?: string;
    /**
     * Instance type. Supported values include: `master` - master instance, `dr` - disaster recovery instance, and `ro` - read-only instance.
     */
    instanceRole?: string;
    /**
     * Number of results returned for a single request. Default is `20`, and maximum is 2000.
     */
    limit?: number;
    /**
     * Instance ID, such as `cdb-c1nl9rpv`. It is identical to the instance ID displayed in the database console page.
     */
    mysqlId?: string;
    /**
     * Record offset. Default is 0.
     */
    offset?: number;
    /**
     * It has been deprecated from version 1.36.0. Please use `chargeType` instead. Pay type of instance, `0`: prepay, `1`: postpaid.
     *
     * @deprecated It has been deprecated from version 1.36.0. Please use `charge_type` instead.
     */
    payType?: number;
    /**
     * Used to store results.
     */
    resultOutputFile?: string;
    /**
     * Security groups ID of instance.
     */
    securityGroupId?: string;
    /**
     * Instance status. Available values: `0` - Creating; `1` - Running; `4` - Isolating; `5` - Isolated.
     */
    status?: number;
    /**
     * Indicates whether to query disaster recovery instances.
     */
    withDr?: number;
    /**
     * Indicates whether to query master instances.
     */
    withMaster?: number;
    /**
     * Indicates whether to query read-only instances.
     */
    withRo?: number;
}

/**
 * A collection of values returned by getInstance.
 */
export interface GetInstanceResult {
    /**
     * Pay type of instance.
     */
    readonly chargeType?: string;
    /**
     * The version number of the database engine to use. Supported versions include `5.5`/`5.6`/`5.7`/`8.0`.
     */
    readonly engineVersion?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Initialization mark. Available values: `0` - Uninitialized; `1` - Initialized.
     */
    readonly initFlag?: number;
    /**
     * A list of instances. Each element contains the following attributes:
     */
    readonly instanceLists: outputs.Mysql.GetInstanceInstanceList[];
    /**
     * Name of mysql instance.
     */
    readonly instanceName?: string;
    /**
     * Instance type. Supported values include: `master` - master instance, `dr` - disaster recovery instance, and `ro` - read-only instance.
     */
    readonly instanceRole?: string;
    readonly limit?: number;
    /**
     * Instance ID, such as `cdb-c1nl9rpv`. It is identical to the instance ID displayed in the database console page.
     */
    readonly mysqlId?: string;
    readonly offset?: number;
    /**
     * Pay type of instance, `0`: prepaid, `1`: postpaid.
     *
     * @deprecated It has been deprecated from version 1.36.0. Please use `charge_type` instead.
     */
    readonly payType?: number;
    readonly resultOutputFile?: string;
    readonly securityGroupId?: string;
    /**
     * Instance status. Available values: `0` - Creating; `1` - Running; `4` - Isolating; `5` - Isolated.
     */
    readonly status?: number;
    readonly withDr?: number;
    readonly withMaster?: number;
    readonly withRo?: number;
}

export function getInstanceOutput(args?: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply(a => getInstance(a, opts))
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceOutputArgs {
    /**
     * Pay type of instance, valid values are `PREPAID` and `POSTPAID`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * The version number of the database engine to use. Supported versions include 5.5/5.6/5.7/8.0.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Initialization mark. Available values: `0` - Uninitialized; `1` - Initialized.
     */
    initFlag?: pulumi.Input<number>;
    /**
     * Name of mysql instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Instance type. Supported values include: `master` - master instance, `dr` - disaster recovery instance, and `ro` - read-only instance.
     */
    instanceRole?: pulumi.Input<string>;
    /**
     * Number of results returned for a single request. Default is `20`, and maximum is 2000.
     */
    limit?: pulumi.Input<number>;
    /**
     * Instance ID, such as `cdb-c1nl9rpv`. It is identical to the instance ID displayed in the database console page.
     */
    mysqlId?: pulumi.Input<string>;
    /**
     * Record offset. Default is 0.
     */
    offset?: pulumi.Input<number>;
    /**
     * It has been deprecated from version 1.36.0. Please use `chargeType` instead. Pay type of instance, `0`: prepay, `1`: postpaid.
     *
     * @deprecated It has been deprecated from version 1.36.0. Please use `charge_type` instead.
     */
    payType?: pulumi.Input<number>;
    /**
     * Used to store results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Security groups ID of instance.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * Instance status. Available values: `0` - Creating; `1` - Running; `4` - Isolating; `5` - Isolated.
     */
    status?: pulumi.Input<number>;
    /**
     * Indicates whether to query disaster recovery instances.
     */
    withDr?: pulumi.Input<number>;
    /**
     * Indicates whether to query master instances.
     */
    withMaster?: pulumi.Input<number>;
    /**
     * Indicates whether to query read-only instances.
     */
    withRo?: pulumi.Input<number>;
}
