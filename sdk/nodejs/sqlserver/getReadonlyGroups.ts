// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query the list of SQL Server readonly groups.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = pulumi.output(tencentcloud.Sqlserver.getDbs({
 *     instanceId: "mssql-ds1xhnt9",
 * }));
 * ```
 */
export function getReadonlyGroups(args?: GetReadonlyGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetReadonlyGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Sqlserver/getReadonlyGroups:getReadonlyGroups", {
        "masterInstanceId": args.masterInstanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getReadonlyGroups.
 */
export interface GetReadonlyGroupsArgs {
    /**
     * Master SQL Server instance ID.
     */
    masterInstanceId?: string;
    /**
     * Used to store results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getReadonlyGroups.
 */
export interface GetReadonlyGroupsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of SQL Server readonly group. Each element contains the following attributes:
     */
    readonly lists: outputs.Sqlserver.GetReadonlyGroupsList[];
    /**
     * Master instance id.
     */
    readonly masterInstanceId?: string;
    readonly resultOutputFile?: string;
}

export function getReadonlyGroupsOutput(args?: GetReadonlyGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReadonlyGroupsResult> {
    return pulumi.output(args).apply(a => getReadonlyGroups(a, opts))
}

/**
 * A collection of arguments for invoking getReadonlyGroups.
 */
export interface GetReadonlyGroupsOutputArgs {
    /**
     * Master SQL Server instance ID.
     */
    masterInstanceId?: pulumi.Input<string>;
    /**
     * Used to store results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
