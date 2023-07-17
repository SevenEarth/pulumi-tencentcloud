// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dbbrain dbSpaceStatus
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const dbSpaceStatus = pulumi.output(tencentcloud.Dbbrain.getDbSpaceStatus({
 *     instanceId: "%s",
 *     product: "mysql",
 *     rangeDays: 7,
 * }));
 * ```
 */
export function getDbSpaceStatus(args: GetDbSpaceStatusArgs, opts?: pulumi.InvokeOptions): Promise<GetDbSpaceStatusResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Dbbrain/getDbSpaceStatus:getDbSpaceStatus", {
        "instanceId": args.instanceId,
        "product": args.product,
        "rangeDays": args.rangeDays,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDbSpaceStatus.
 */
export interface GetDbSpaceStatusArgs {
    /**
     * instance id.
     */
    instanceId: string;
    /**
     * Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
     */
    product?: string;
    /**
     * The number of days in the time period, the deadline is the current day, and the default is 7 days.
     */
    rangeDays?: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDbSpaceStatus.
 */
export interface GetDbSpaceStatusResult {
    /**
     * Estimated number of days available.
     */
    readonly availableDays: number;
    /**
     * Disk growth (MB).
     */
    readonly growth: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly product?: string;
    readonly rangeDays?: number;
    /**
     * Disk remaining (MB).
     */
    readonly remain: number;
    readonly resultOutputFile?: string;
    /**
     * Total disk size (MB).
     */
    readonly total: number;
}

export function getDbSpaceStatusOutput(args: GetDbSpaceStatusOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDbSpaceStatusResult> {
    return pulumi.output(args).apply(a => getDbSpaceStatus(a, opts))
}

/**
 * A collection of arguments for invoking getDbSpaceStatus.
 */
export interface GetDbSpaceStatusOutputArgs {
    /**
     * instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
     */
    product?: pulumi.Input<string>;
    /**
     * The number of days in the time period, the deadline is the current day, and the default is 7 days.
     */
    rangeDays?: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
