// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a PrivateDns zoneVpcAttachment
 *
 * > **NOTE:**  If you need to bind account A to account B's VPC resources, you need to first grant role authorization to account A.
 *
 * ## Example Usage
 * ### Append VPC associated with private dns zone
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const exampleZone = new tencentcloud.privatedns.Zone("exampleZone", {
 *     domain: "domain.com",
 *     remark: "remark.",
 *     dnsForwardStatus: "DISABLED",
 *     cnameSpeedupStatus: "ENABLED",
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * const exampleZoneVpcAttachment = new tencentcloud.privatedns.ZoneVpcAttachment("exampleZoneVpcAttachment", {
 *     zoneId: exampleZone.id,
 *     vpcSet: {
 *         uniqVpcId: vpc.id,
 *         region: "ap-guangzhou",
 *     },
 * });
 * ```
 * ### Add VPC information for associated accounts in the private dns zone
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.privatedns.ZoneVpcAttachment("example", {
 *     zoneId: tencentcloud_private_dns_zone.example.id,
 *     accountVpcSet: {
 *         uniqVpcId: "vpc-82znjzn3",
 *         region: "ap-guangzhou",
 *         uin: "100017155920",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * PrivateDns zone_vpc_attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:PrivateDns/zoneVpcAttachment:ZoneVpcAttachment example zone-6t11lof0#vpc-jdx11z0t
 * ```
 */
export class ZoneVpcAttachment extends pulumi.CustomResource {
    /**
     * Get an existing ZoneVpcAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneVpcAttachmentState, opts?: pulumi.CustomResourceOptions): ZoneVpcAttachment {
        return new ZoneVpcAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:PrivateDns/zoneVpcAttachment:ZoneVpcAttachment';

    /**
     * Returns true if the given object is an instance of ZoneVpcAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ZoneVpcAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ZoneVpcAttachment.__pulumiType;
    }

    /**
     * New add account vpc info.
     */
    public readonly accountVpcSet!: pulumi.Output<outputs.PrivateDns.ZoneVpcAttachmentAccountVpcSet | undefined>;
    /**
     * New add vpc info.
     */
    public readonly vpcSet!: pulumi.Output<outputs.PrivateDns.ZoneVpcAttachmentVpcSet | undefined>;
    /**
     * PrivateZone ID.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a ZoneVpcAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZoneVpcAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZoneVpcAttachmentArgs | ZoneVpcAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ZoneVpcAttachmentState | undefined;
            resourceInputs["accountVpcSet"] = state ? state.accountVpcSet : undefined;
            resourceInputs["vpcSet"] = state ? state.vpcSet : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ZoneVpcAttachmentArgs | undefined;
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["accountVpcSet"] = args ? args.accountVpcSet : undefined;
            resourceInputs["vpcSet"] = args ? args.vpcSet : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ZoneVpcAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ZoneVpcAttachment resources.
 */
export interface ZoneVpcAttachmentState {
    /**
     * New add account vpc info.
     */
    accountVpcSet?: pulumi.Input<inputs.PrivateDns.ZoneVpcAttachmentAccountVpcSet>;
    /**
     * New add vpc info.
     */
    vpcSet?: pulumi.Input<inputs.PrivateDns.ZoneVpcAttachmentVpcSet>;
    /**
     * PrivateZone ID.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ZoneVpcAttachment resource.
 */
export interface ZoneVpcAttachmentArgs {
    /**
     * New add account vpc info.
     */
    accountVpcSet?: pulumi.Input<inputs.PrivateDns.ZoneVpcAttachmentAccountVpcSet>;
    /**
     * New add vpc info.
     */
    vpcSet?: pulumi.Input<inputs.PrivateDns.ZoneVpcAttachmentVpcSet>;
    /**
     * PrivateZone ID.
     */
    zoneId: pulumi.Input<string>;
}
