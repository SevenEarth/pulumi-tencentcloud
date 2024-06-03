// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of HA VIPs.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const havips = tencentcloud.Ha.getVips({
 *     addressIp: "10.0.4.16",
 *     id: "havip-kjqwe4ba",
 *     name: "test",
 *     subnetId: "subnet-4d4m4cd4",
 *     vpcId: "vpc-gzea3dd7",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getVips(args?: GetVipsArgs, opts?: pulumi.InvokeOptions): Promise<GetVipsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Ha/getVips:getVips", {
        "addressIp": args.addressIp,
        "id": args.id,
        "name": args.name,
        "resultOutputFile": args.resultOutputFile,
        "subnetId": args.subnetId,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVips.
 */
export interface GetVipsArgs {
    /**
     * EIP of the HA VIP to be queried.
     */
    addressIp?: string;
    /**
     * ID of the HA VIP to be queried.
     */
    id?: string;
    /**
     * Name of the HA VIP. The length of character is limited to 1-60.
     */
    name?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Subnet id of the HA VIP to be queried.
     */
    subnetId?: string;
    /**
     * VPC id of the HA VIP to be queried.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getVips.
 */
export interface GetVipsResult {
    /**
     * EIP that is associated.
     */
    readonly addressIp?: string;
    /**
     * Information list of the dedicated HA VIPs.
     */
    readonly haVipLists: outputs.Ha.GetVipsHaVipList[];
    /**
     * ID of the HA VIP.
     */
    readonly id?: string;
    /**
     * Name of the HA VIP.
     */
    readonly name?: string;
    readonly resultOutputFile?: string;
    /**
     * Subnet id.
     */
    readonly subnetId?: string;
    /**
     * VPC id.
     */
    readonly vpcId?: string;
}
/**
 * Use this data source to query detailed information of HA VIPs.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const havips = tencentcloud.Ha.getVips({
 *     addressIp: "10.0.4.16",
 *     id: "havip-kjqwe4ba",
 *     name: "test",
 *     subnetId: "subnet-4d4m4cd4",
 *     vpcId: "vpc-gzea3dd7",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getVipsOutput(args?: GetVipsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVipsResult> {
    return pulumi.output(args).apply((a: any) => getVips(a, opts))
}

/**
 * A collection of arguments for invoking getVips.
 */
export interface GetVipsOutputArgs {
    /**
     * EIP of the HA VIP to be queried.
     */
    addressIp?: pulumi.Input<string>;
    /**
     * ID of the HA VIP to be queried.
     */
    id?: pulumi.Input<string>;
    /**
     * Name of the HA VIP. The length of character is limited to 1-60.
     */
    name?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Subnet id of the HA VIP to be queried.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * VPC id of the HA VIP to be queried.
     */
    vpcId?: pulumi.Input<string>;
}
