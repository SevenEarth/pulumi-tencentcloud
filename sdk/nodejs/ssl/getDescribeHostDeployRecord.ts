// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of ssl describeHostDeployRecord
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const describeHostDeployRecord = tencentcloud.Ssl.getDescribeHostDeployRecord({
 *     certificateId: "8u8DII0l",
 *     resourceType: "ddos",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDescribeHostDeployRecord(args: GetDescribeHostDeployRecordArgs, opts?: pulumi.InvokeOptions): Promise<GetDescribeHostDeployRecordResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Ssl/getDescribeHostDeployRecord:getDescribeHostDeployRecord", {
        "certificateId": args.certificateId,
        "resourceType": args.resourceType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDescribeHostDeployRecord.
 */
export interface GetDescribeHostDeployRecordArgs {
    /**
     * Certificate ID to be deployed.
     */
    certificateId: string;
    /**
     * Resource Type.
     */
    resourceType?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDescribeHostDeployRecord.
 */
export interface GetDescribeHostDeployRecordResult {
    readonly certificateId: string;
    /**
     * Certificate deployment record listNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    readonly deployRecordLists: outputs.Ssl.GetDescribeHostDeployRecordDeployRecordList[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Deploy resource type.
     */
    readonly resourceType?: string;
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of ssl describeHostDeployRecord
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const describeHostDeployRecord = tencentcloud.Ssl.getDescribeHostDeployRecord({
 *     certificateId: "8u8DII0l",
 *     resourceType: "ddos",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDescribeHostDeployRecordOutput(args: GetDescribeHostDeployRecordOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDescribeHostDeployRecordResult> {
    return pulumi.output(args).apply((a: any) => getDescribeHostDeployRecord(a, opts))
}

/**
 * A collection of arguments for invoking getDescribeHostDeployRecord.
 */
export interface GetDescribeHostDeployRecordOutputArgs {
    /**
     * Certificate ID to be deployed.
     */
    certificateId: pulumi.Input<string>;
    /**
     * Resource Type.
     */
    resourceType?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
