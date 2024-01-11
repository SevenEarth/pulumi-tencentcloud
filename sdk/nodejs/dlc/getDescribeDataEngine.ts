// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dlc describeDataEngine
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const describeDataEngine = pulumi.output(tencentcloud.Dlc.getDescribeDataEngine({
 *     dataEngineName: "testSpark",
 * }));
 * ```
 */
export function getDescribeDataEngine(args: GetDescribeDataEngineArgs, opts?: pulumi.InvokeOptions): Promise<GetDescribeDataEngineResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Dlc/getDescribeDataEngine:getDescribeDataEngine", {
        "dataEngineName": args.dataEngineName,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDescribeDataEngine.
 */
export interface GetDescribeDataEngineArgs {
    /**
     * Engine name.
     */
    dataEngineName: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDescribeDataEngine.
 */
export interface GetDescribeDataEngineResult {
    /**
     * Engine name.
     */
    readonly dataEngineName: string;
    /**
     * Data engine details.
     */
    readonly dataEngines: outputs.Dlc.GetDescribeDataEngineDataEngine[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
}

export function getDescribeDataEngineOutput(args: GetDescribeDataEngineOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDescribeDataEngineResult> {
    return pulumi.output(args).apply(a => getDescribeDataEngine(a, opts))
}

/**
 * A collection of arguments for invoking getDescribeDataEngine.
 */
export interface GetDescribeDataEngineOutputArgs {
    /**
     * Engine name.
     */
    dataEngineName: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}