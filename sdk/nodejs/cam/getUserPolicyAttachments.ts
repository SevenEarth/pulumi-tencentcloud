// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of CAM user policy attachments
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = tencentcloud.Cam.getUserPolicyAttachments({
 *     userId: tencentcloud_cam_user.foo.id,
 * });
 * const bar = tencentcloud.Cam.getUserPolicyAttachments({
 *     userId: tencentcloud_cam_user.foo.id,
 *     policyId: tencentcloud_cam_policy.foo.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getUserPolicyAttachments(args?: GetUserPolicyAttachmentsArgs, opts?: pulumi.InvokeOptions): Promise<GetUserPolicyAttachmentsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cam/getUserPolicyAttachments:getUserPolicyAttachments", {
        "createMode": args.createMode,
        "policyId": args.policyId,
        "policyType": args.policyType,
        "resultOutputFile": args.resultOutputFile,
        "userId": args.userId,
        "userName": args.userName,
    }, opts);
}

/**
 * A collection of arguments for invoking getUserPolicyAttachments.
 */
export interface GetUserPolicyAttachmentsArgs {
    /**
     * Mode of Creation of the CAM user policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
     */
    createMode?: number;
    /**
     * ID of CAM policy to be queried.
     */
    policyId?: string;
    /**
     * Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
     */
    policyType?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * It has been deprecated from version 1.59.6. Use `userName` instead. ID of the attached CAM user to be queried.
     *
     * @deprecated It has been deprecated from version 1.59.6. Use `userName` instead.
     */
    userId?: string;
    /**
     * Name of the attached CAM user as unique key to be queried.
     */
    userName?: string;
}

/**
 * A collection of values returned by getUserPolicyAttachments.
 */
export interface GetUserPolicyAttachmentsResult {
    /**
     * Mode of Creation of the CAM user policy attachment. `1` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
     */
    readonly createMode?: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of CAM user.
     */
    readonly policyId?: string;
    /**
     * Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
     */
    readonly policyType?: string;
    readonly resultOutputFile?: string;
    /**
     * (**Deprecated**) It has been deprecated from version 1.59.6. Use `userName` instead. ID of CAM user.
     *
     * @deprecated It has been deprecated from version 1.59.6. Use `userName` instead.
     */
    readonly userId?: string;
    /**
     * Name of CAM user as unique key.
     */
    readonly userName?: string;
    /**
     * A list of CAM user policy attachments. Each element contains the following attributes:
     */
    readonly userPolicyAttachmentLists: outputs.Cam.GetUserPolicyAttachmentsUserPolicyAttachmentList[];
}
/**
 * Use this data source to query detailed information of CAM user policy attachments
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = tencentcloud.Cam.getUserPolicyAttachments({
 *     userId: tencentcloud_cam_user.foo.id,
 * });
 * const bar = tencentcloud.Cam.getUserPolicyAttachments({
 *     userId: tencentcloud_cam_user.foo.id,
 *     policyId: tencentcloud_cam_policy.foo.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getUserPolicyAttachmentsOutput(args?: GetUserPolicyAttachmentsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserPolicyAttachmentsResult> {
    return pulumi.output(args).apply((a: any) => getUserPolicyAttachments(a, opts))
}

/**
 * A collection of arguments for invoking getUserPolicyAttachments.
 */
export interface GetUserPolicyAttachmentsOutputArgs {
    /**
     * Mode of Creation of the CAM user policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
     */
    createMode?: pulumi.Input<number>;
    /**
     * ID of CAM policy to be queried.
     */
    policyId?: pulumi.Input<string>;
    /**
     * Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
     */
    policyType?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.59.6. Use `userName` instead. ID of the attached CAM user to be queried.
     *
     * @deprecated It has been deprecated from version 1.59.6. Use `userName` instead.
     */
    userId?: pulumi.Input<string>;
    /**
     * Name of the attached CAM user as unique key to be queried.
     */
    userName?: pulumi.Input<string>;
}
