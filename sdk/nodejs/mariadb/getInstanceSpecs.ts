// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of mariadb instanceSpecs
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const instanceSpecs = tencentcloud.Mariadb.getInstanceSpecs({});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstanceSpecs(args?: GetInstanceSpecsArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceSpecsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Mariadb/getInstanceSpecs:getInstanceSpecs", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceSpecs.
 */
export interface GetInstanceSpecsArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getInstanceSpecs.
 */
export interface GetInstanceSpecsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * list of instance specifications.
     */
    readonly specs: outputs.Mariadb.GetInstanceSpecsSpec[];
}
/**
 * Use this data source to query detailed information of mariadb instanceSpecs
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const instanceSpecs = tencentcloud.Mariadb.getInstanceSpecs({});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstanceSpecsOutput(args?: GetInstanceSpecsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceSpecsResult> {
    return pulumi.output(args).apply((a: any) => getInstanceSpecs(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceSpecs.
 */
export interface GetInstanceSpecsOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
