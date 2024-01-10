// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of ssm sshKeyPairValue
 *
 * > **NOTE:** Must set at least one of `secretName` or `sshKeyId`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = pulumi.output(tencentcloud.Ssm.getSshKeyPairValue({
 *     secretName: "keep_terraform",
 *     sshKeyId: "skey-2ae2snwd",
 * }));
 * ```
 * ### Or
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = pulumi.output(tencentcloud.Ssm.getSshKeyPairValue({
 *     secretName: "keep_terraform",
 * }));
 * ```
 * ### Or
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = pulumi.output(tencentcloud.Ssm.getSshKeyPairValue({
 *     sshKeyId: "skey-2ae2snwd",
 * }));
 * ```
 */
export function getSshKeyPairValue(args?: GetSshKeyPairValueArgs, opts?: pulumi.InvokeOptions): Promise<GetSshKeyPairValueResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Ssm/getSshKeyPairValue:getSshKeyPairValue", {
        "resultOutputFile": args.resultOutputFile,
        "secretName": args.secretName,
        "sshKeyId": args.sshKeyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSshKeyPairValue.
 */
export interface GetSshKeyPairValueArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Secret name.
     */
    secretName?: string;
    /**
     * The key pair ID is the unique identifier of the key pair in the cloud server.
     */
    sshKeyId?: string;
}

/**
 * A collection of values returned by getSshKeyPairValue.
 */
export interface GetSshKeyPairValueResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Private key plain text, encoded using base64.
     */
    readonly privateKey: string;
    /**
     * The project ID to which this key pair belongs.
     */
    readonly projectId: number;
    /**
     * Public key plain text, encoded using base64.
     */
    readonly publicKey: string;
    readonly resultOutputFile?: string;
    readonly secretName: string;
    /**
     * Description of the SSH key pair. Users can modify the description information of the key pair in the CVM console.
     */
    readonly sshKeyDescription: string;
    readonly sshKeyId: string;
    /**
     * SSH key name.
     */
    readonly sshKeyName: string;
}

export function getSshKeyPairValueOutput(args?: GetSshKeyPairValueOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSshKeyPairValueResult> {
    return pulumi.output(args).apply(a => getSshKeyPairValue(a, opts))
}

/**
 * A collection of arguments for invoking getSshKeyPairValue.
 */
export interface GetSshKeyPairValueOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Secret name.
     */
    secretName?: pulumi.Input<string>;
    /**
     * The key pair ID is the unique identifier of the key pair in the cloud server.
     */
    sshKeyId?: pulumi.Input<string>;
}
