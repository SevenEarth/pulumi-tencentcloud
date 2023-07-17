// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a vpc ipv6CidrBlock
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const cidr_block = new tencentcloud.vpc.Instance("cidr-block", {
 *     cidrBlock: "10.0.0.0/16",
 *     isMulticast: false,
 * });
 * const ipv6CidrBlock = new tencentcloud.vpc.Ipv6CidrBlock("ipv6CidrBlock", {vpcId: cidr_block.id});
 * ```
 *
 * ## Import
 *
 * vpc ipv6_cidr_block can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Vpc/ipv6CidrBlock:Ipv6CidrBlock ipv6_cidr_block vpc_id
 * ```
 */
export class Ipv6CidrBlock extends pulumi.CustomResource {
    /**
     * Get an existing Ipv6CidrBlock resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Ipv6CidrBlockState, opts?: pulumi.CustomResourceOptions): Ipv6CidrBlock {
        return new Ipv6CidrBlock(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Vpc/ipv6CidrBlock:Ipv6CidrBlock';

    /**
     * Returns true if the given object is an instance of Ipv6CidrBlock.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ipv6CidrBlock {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ipv6CidrBlock.__pulumiType;
    }

    /**
     * ipv6 cidr block.
     */
    public /*out*/ readonly ipv6CidrBlock!: pulumi.Output<string>;
    /**
     * `VPC` instance `ID`, in the form of `vpc-f49l6u0z`.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Ipv6CidrBlock resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Ipv6CidrBlockArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Ipv6CidrBlockArgs | Ipv6CidrBlockState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Ipv6CidrBlockState | undefined;
            resourceInputs["ipv6CidrBlock"] = state ? state.ipv6CidrBlock : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as Ipv6CidrBlockArgs | undefined;
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["ipv6CidrBlock"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ipv6CidrBlock.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ipv6CidrBlock resources.
 */
export interface Ipv6CidrBlockState {
    /**
     * ipv6 cidr block.
     */
    ipv6CidrBlock?: pulumi.Input<string>;
    /**
     * `VPC` instance `ID`, in the form of `vpc-f49l6u0z`.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ipv6CidrBlock resource.
 */
export interface Ipv6CidrBlockArgs {
    /**
     * `VPC` instance `ID`, in the form of `vpc-f49l6u0z`.
     */
    vpcId: pulumi.Input<string>;
}
