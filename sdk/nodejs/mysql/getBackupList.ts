// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query the list of backup databases.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const defaultBackupList = pulumi.output(tencentcloud.Mysql.getBackupList({
 *     maxNumber: 10,
 *     mysqlId: "terraform-test-local-database",
 *     resultOutputFile: "mytestpath",
 * }));
 * ```
 */
export function getBackupList(args: GetBackupListArgs, opts?: pulumi.InvokeOptions): Promise<GetBackupListResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Mysql/getBackupList:getBackupList", {
        "maxNumber": args.maxNumber,
        "mysqlId": args.mysqlId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackupList.
 */
export interface GetBackupListArgs {
    /**
     * The latest files to list, rang from 1 to 10000. And the default value is `10`.
     */
    maxNumber?: number;
    /**
     * Instance ID, such as `cdb-c1nl9rpv`. It is identical to the instance ID displayed in the database console page.
     */
    mysqlId: string;
    /**
     * Used to store results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getBackupList.
 */
export interface GetBackupListResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of MySQL backup. Each element contains the following attributes:
     */
    readonly lists: outputs.Mysql.GetBackupListList[];
    readonly maxNumber?: number;
    readonly mysqlId: string;
    readonly resultOutputFile?: string;
}

export function getBackupListOutput(args: GetBackupListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackupListResult> {
    return pulumi.output(args).apply(a => getBackupList(a, opts))
}

/**
 * A collection of arguments for invoking getBackupList.
 */
export interface GetBackupListOutputArgs {
    /**
     * The latest files to list, rang from 1 to 10000. And the default value is `10`.
     */
    maxNumber?: pulumi.Input<number>;
    /**
     * Instance ID, such as `cdb-c1nl9rpv`. It is identical to the instance ID displayed in the database console page.
     */
    mysqlId: pulumi.Input<string>;
    /**
     * Used to store results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
