// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dbbrain diagEvents
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const diagEvents = tencentcloud.Dbbrain.getDiagEvents({
 *     endTime: "%s",
 *     instanceIds: ["%s"],
 *     severities: [
 *         1,
 *         4,
 *         5,
 *     ],
 *     startTime: "%s",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDiagEvents(args: GetDiagEventsArgs, opts?: pulumi.InvokeOptions): Promise<GetDiagEventsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dbbrain/getDiagEvents:getDiagEvents", {
        "endTime": args.endTime,
        "instanceIds": args.instanceIds,
        "resultOutputFile": args.resultOutputFile,
        "severities": args.severities,
        "startTime": args.startTime,
    }, opts);
}

/**
 * A collection of arguments for invoking getDiagEvents.
 */
export interface GetDiagEventsArgs {
    /**
     * end time.
     */
    endTime: string;
    /**
     * instance id list.
     */
    instanceIds?: string[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * severity list, optional value is 1-fatal, 2-severity, 3-warning, 4-tips, 5-health.
     */
    severities?: number[];
    /**
     * start time.
     */
    startTime: string;
}

/**
 * A collection of values returned by getDiagEvents.
 */
export interface GetDiagEventsResult {
    /**
     * end time.
     */
    readonly endTime: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceIds?: string[];
    /**
     * diag event list.
     */
    readonly lists: outputs.Dbbrain.GetDiagEventsList[];
    readonly resultOutputFile?: string;
    readonly severities?: number[];
    /**
     * start time.
     */
    readonly startTime: string;
}
/**
 * Use this data source to query detailed information of dbbrain diagEvents
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const diagEvents = tencentcloud.Dbbrain.getDiagEvents({
 *     endTime: "%s",
 *     instanceIds: ["%s"],
 *     severities: [
 *         1,
 *         4,
 *         5,
 *     ],
 *     startTime: "%s",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDiagEventsOutput(args: GetDiagEventsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDiagEventsResult> {
    return pulumi.output(args).apply((a: any) => getDiagEvents(a, opts))
}

/**
 * A collection of arguments for invoking getDiagEvents.
 */
export interface GetDiagEventsOutputArgs {
    /**
     * end time.
     */
    endTime: pulumi.Input<string>;
    /**
     * instance id list.
     */
    instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * severity list, optional value is 1-fatal, 2-severity, 3-warning, 4-tips, 5-health.
     */
    severities?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * start time.
     */
    startTime: pulumi.Input<string>;
}
