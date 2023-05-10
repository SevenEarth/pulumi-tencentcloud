// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dts syncJobs
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const job = new tencentcloud.dts.SyncJob("job", {
 *     jobName: "tf_dts_test",
 *     payMode: "PostPay",
 *     srcDatabaseType: "mysql",
 *     srcRegion: "ap-guangzhou",
 *     dstDatabaseType: "cynosdbmysql",
 *     dstRegion: "ap-guangzhou",
 *     tags: [{
 *         tagKey: "aaa",
 *         tagValue: "bbb",
 *     }],
 *     autoRenew: 0,
 *     instanceClass: "micro",
 * });
 * const syncJobs = tencentcloud.Dts.getSyncJobsOutput({
 *     jobId: job.id,
 *     jobName: "tf_dts_test",
 * });
 * ```
 */
export function getSyncJobs(args?: GetSyncJobsArgs, opts?: pulumi.InvokeOptions): Promise<GetSyncJobsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Dts/getSyncJobs:getSyncJobs", {
        "jobId": args.jobId,
        "jobName": args.jobName,
        "jobType": args.jobType,
        "order": args.order,
        "orderSeq": args.orderSeq,
        "payMode": args.payMode,
        "resultOutputFile": args.resultOutputFile,
        "runMode": args.runMode,
        "statuses": args.statuses,
        "tagFilters": args.tagFilters,
    }, opts);
}

/**
 * A collection of arguments for invoking getSyncJobs.
 */
export interface GetSyncJobsArgs {
    /**
     * job id.
     */
    jobId?: string;
    /**
     * job name.
     */
    jobName?: string;
    /**
     * job type.
     */
    jobType?: string;
    /**
     * order field.
     */
    order?: string;
    /**
     * order way, optional value is DESC or ASC.
     */
    orderSeq?: string;
    /**
     * pay mode, optional value is PrePay or PostPay.
     */
    payMode?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * run mode, optional value is mmediate or Timed.
     */
    runMode?: string;
    /**
     * status.
     */
    statuses?: string[];
    /**
     * tag filters.
     */
    tagFilters?: inputs.Dts.GetSyncJobsTagFilter[];
}

/**
 * A collection of values returned by getSyncJobs.
 */
export interface GetSyncJobsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * job id.
     */
    readonly jobId?: string;
    /**
     * job name.
     */
    readonly jobName?: string;
    readonly jobType?: string;
    /**
     * sync job list.
     */
    readonly lists: outputs.Dts.GetSyncJobsList[];
    readonly order?: string;
    readonly orderSeq?: string;
    /**
     * pay mode.
     */
    readonly payMode?: string;
    readonly resultOutputFile?: string;
    /**
     * run mode.
     */
    readonly runMode?: string;
    /**
     * status.
     */
    readonly statuses?: string[];
    readonly tagFilters?: outputs.Dts.GetSyncJobsTagFilter[];
}

export function getSyncJobsOutput(args?: GetSyncJobsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSyncJobsResult> {
    return pulumi.output(args).apply(a => getSyncJobs(a, opts))
}

/**
 * A collection of arguments for invoking getSyncJobs.
 */
export interface GetSyncJobsOutputArgs {
    /**
     * job id.
     */
    jobId?: pulumi.Input<string>;
    /**
     * job name.
     */
    jobName?: pulumi.Input<string>;
    /**
     * job type.
     */
    jobType?: pulumi.Input<string>;
    /**
     * order field.
     */
    order?: pulumi.Input<string>;
    /**
     * order way, optional value is DESC or ASC.
     */
    orderSeq?: pulumi.Input<string>;
    /**
     * pay mode, optional value is PrePay or PostPay.
     */
    payMode?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * run mode, optional value is mmediate or Timed.
     */
    runMode?: pulumi.Input<string>;
    /**
     * status.
     */
    statuses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * tag filters.
     */
    tagFilters?: pulumi.Input<pulumi.Input<inputs.Dts.GetSyncJobsTagFilterArgs>[]>;
}
