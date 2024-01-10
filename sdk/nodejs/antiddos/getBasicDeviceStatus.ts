// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of antiddos basicDeviceStatus
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const basicDeviceStatus = pulumi.output(tencentcloud.Antiddos.getBasicDeviceStatus({
 *     filterRegion: 1,
 *     ipLists: ["127.0.0.1"],
 * }));
 * ```
 */
export function getBasicDeviceStatus(args?: GetBasicDeviceStatusArgs, opts?: pulumi.InvokeOptions): Promise<GetBasicDeviceStatusResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Antiddos/getBasicDeviceStatus:getBasicDeviceStatus", {
        "filterRegion": args.filterRegion,
        "idLists": args.idLists,
        "ipLists": args.ipLists,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getBasicDeviceStatus.
 */
export interface GetBasicDeviceStatusArgs {
    /**
     * Region Id.
     */
    filterRegion?: number;
    /**
     * Named resource transfer ID.
     */
    idLists?: string[];
    /**
     * Ip resource list.
     */
    ipLists?: string[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getBasicDeviceStatus.
 */
export interface GetBasicDeviceStatusResult {
    /**
     * Note: This field may return null, indicating that a valid value cannot be obtained.
     */
    readonly clbDatas: outputs.Antiddos.GetBasicDeviceStatusClbData[];
    /**
     * Return resources and status, status code: 1- Blocking status 2- Normal status 3- Attack status.
     */
    readonly datas: outputs.Antiddos.GetBasicDeviceStatusData[];
    readonly filterRegion?: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly idLists?: string[];
    readonly ipLists?: string[];
    readonly resultOutputFile?: string;
}

export function getBasicDeviceStatusOutput(args?: GetBasicDeviceStatusOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBasicDeviceStatusResult> {
    return pulumi.output(args).apply(a => getBasicDeviceStatus(a, opts))
}

/**
 * A collection of arguments for invoking getBasicDeviceStatus.
 */
export interface GetBasicDeviceStatusOutputArgs {
    /**
     * Region Id.
     */
    filterRegion?: pulumi.Input<number>;
    /**
     * Named resource transfer ID.
     */
    idLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Ip resource list.
     */
    ipLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
