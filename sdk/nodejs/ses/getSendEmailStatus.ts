// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of ses sendEmailStatus
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const sendEmailStatus = pulumi.output(tencentcloud.Ses.getSendEmailStatus({
 *     messageId: "qcloudses-30-4123414323-date-20210101094334-syNARhMTbKI1",
 *     requestDate: "2020-09-22",
 *     toEmailAddress: "example@cloud.com",
 * }));
 * ```
 */
export function getSendEmailStatus(args: GetSendEmailStatusArgs, opts?: pulumi.InvokeOptions): Promise<GetSendEmailStatusResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Ses/getSendEmailStatus:getSendEmailStatus", {
        "messageId": args.messageId,
        "requestDate": args.requestDate,
        "resultOutputFile": args.resultOutputFile,
        "toEmailAddress": args.toEmailAddress,
    }, opts);
}

/**
 * A collection of arguments for invoking getSendEmailStatus.
 */
export interface GetSendEmailStatusArgs {
    /**
     * The MessageId field returned by the SendMail API.
     */
    messageId?: string;
    /**
     * Date sent. This parameter is required. You can only query the sending status for a single date at a time.
     */
    requestDate: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Recipient email address.
     */
    toEmailAddress?: string;
}

/**
 * A collection of values returned by getSendEmailStatus.
 */
export interface GetSendEmailStatusResult {
    /**
     * Status of sent emails.
     */
    readonly emailStatusLists: outputs.Ses.GetSendEmailStatusEmailStatusList[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The MessageId field returned by the SendEmail API.
     */
    readonly messageId?: string;
    readonly requestDate: string;
    readonly resultOutputFile?: string;
    /**
     * Recipient email address.
     */
    readonly toEmailAddress?: string;
}

export function getSendEmailStatusOutput(args: GetSendEmailStatusOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSendEmailStatusResult> {
    return pulumi.output(args).apply(a => getSendEmailStatus(a, opts))
}

/**
 * A collection of arguments for invoking getSendEmailStatus.
 */
export interface GetSendEmailStatusOutputArgs {
    /**
     * The MessageId field returned by the SendMail API.
     */
    messageId?: pulumi.Input<string>;
    /**
     * Date sent. This parameter is required. You can only query the sending status for a single date at a time.
     */
    requestDate: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Recipient email address.
     */
    toEmailAddress?: pulumi.Input<string>;
}
