// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dbbrain diagDbInstances
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const diagDbInstances = pulumi.output(tencentcloud.Dbbrain.getDiagDbInstances({
 *     instanceNames: ["keep_preset_mysql"],
 *     isSupported: true,
 *     product: "mysql",
 * }));
 * ```
 */
export function getDiagDbInstances(args: GetDiagDbInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetDiagDbInstancesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Dbbrain/getDiagDbInstances:getDiagDbInstances", {
        "instanceIds": args.instanceIds,
        "instanceNames": args.instanceNames,
        "isSupported": args.isSupported,
        "product": args.product,
        "regions": args.regions,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDiagDbInstances.
 */
export interface GetDiagDbInstancesArgs {
    /**
     * query based on the instance ID condition.
     */
    instanceIds?: string[];
    /**
     * query based on the instance name condition.
     */
    instanceNames?: string[];
    /**
     * whether it is an instance supported by DBbrain, always pass `true`.
     */
    isSupported: boolean;
    /**
     * service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database TDSQL-C for MySQL, the default is `mysql`.
     */
    product: string;
    /**
     * query based on geographical conditions.
     */
    regions?: string[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDiagDbInstances.
 */
export interface GetDiagDbInstancesResult {
    /**
     * all-instance inspection status. `0`: All-instance inspection is enabled; `1`: All-instance inspection is not enabled.
     */
    readonly dbScanStatus: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceIds?: string[];
    readonly instanceNames?: string[];
    /**
     * whether it is an instance supported by DBbrain.
     */
    readonly isSupported: boolean;
    /**
     * information about the instance.
     */
    readonly items: outputs.Dbbrain.GetDiagDbInstancesItem[];
    /**
     * belongs to the product.
     */
    readonly product: string;
    readonly regions?: string[];
    readonly resultOutputFile?: string;
}

export function getDiagDbInstancesOutput(args: GetDiagDbInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDiagDbInstancesResult> {
    return pulumi.output(args).apply(a => getDiagDbInstances(a, opts))
}

/**
 * A collection of arguments for invoking getDiagDbInstances.
 */
export interface GetDiagDbInstancesOutputArgs {
    /**
     * query based on the instance ID condition.
     */
    instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * query based on the instance name condition.
     */
    instanceNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * whether it is an instance supported by DBbrain, always pass `true`.
     */
    isSupported: pulumi.Input<boolean>;
    /**
     * service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database TDSQL-C for MySQL, the default is `mysql`.
     */
    product: pulumi.Input<string>;
    /**
     * query based on geographical conditions.
     */
    regions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
