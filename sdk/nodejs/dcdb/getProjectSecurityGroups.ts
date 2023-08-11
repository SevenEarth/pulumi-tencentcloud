// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dcdb projectSecurityGroups
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const projectSecurityGroups = pulumi.output(tencentcloud.Dcdb.getProjectSecurityGroups({
 *     product: "dcdb",
 *     projectId: 0,
 * }));
 * ```
 */
export function getProjectSecurityGroups(args: GetProjectSecurityGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectSecurityGroupsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Dcdb/getProjectSecurityGroups:getProjectSecurityGroups", {
        "product": args.product,
        "projectId": args.projectId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectSecurityGroups.
 */
export interface GetProjectSecurityGroupsArgs {
    /**
     * Database engine name. Valid value: `dcdb`.
     */
    product: string;
    /**
     * Project ID.
     */
    projectId?: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getProjectSecurityGroups.
 */
export interface GetProjectSecurityGroupsResult {
    /**
     * Security group details.
     */
    readonly groups: outputs.Dcdb.GetProjectSecurityGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly product: string;
    /**
     * Project ID.
     */
    readonly projectId?: number;
    readonly resultOutputFile?: string;
}

export function getProjectSecurityGroupsOutput(args: GetProjectSecurityGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectSecurityGroupsResult> {
    return pulumi.output(args).apply(a => getProjectSecurityGroups(a, opts))
}

/**
 * A collection of arguments for invoking getProjectSecurityGroups.
 */
export interface GetProjectSecurityGroupsOutputArgs {
    /**
     * Database engine name. Valid value: `dcdb`.
     */
    product: pulumi.Input<string>;
    /**
     * Project ID.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}