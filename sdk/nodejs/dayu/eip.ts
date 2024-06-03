// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this resource to create dayu eip rule
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const test = new tencentcloud.dayu.Eip("test", {
 *     bindResourceId: "ins-4m0jvxic",
 *     bindResourceRegion: "hk",
 *     bindResourceType: "cvm",
 *     eip: "162.62.163.50",
 *     resourceId: "bgpip-000004xg",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class Eip extends pulumi.CustomResource {
    /**
     * Get an existing Eip resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EipState, opts?: pulumi.CustomResourceOptions): Eip {
        return new Eip(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dayu/eip:Eip';

    /**
     * Returns true if the given object is an instance of Eip.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Eip {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Eip.__pulumiType;
    }

    /**
     * Resource id to bind.
     */
    public readonly bindResourceId!: pulumi.Output<string>;
    /**
     * Resource region to bind.
     */
    public readonly bindResourceRegion!: pulumi.Output<string>;
    /**
     * Resource type to bind, value range [`clb`, `cvm`].
     */
    public readonly bindResourceType!: pulumi.Output<string>;
    /**
     * Created time of the resource instance.
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * Eip of the resource.
     */
    public readonly eip!: pulumi.Output<string>;
    /**
     * Eip address status of the resource instance.
     */
    public /*out*/ readonly eipAddressStatus!: pulumi.Output<string>;
    /**
     * Eip bound rsc eni of the resource instance.
     */
    public /*out*/ readonly eipBoundRscEni!: pulumi.Output<string>;
    /**
     * Eip bound rsc ins of the resource instance.
     */
    public /*out*/ readonly eipBoundRscIns!: pulumi.Output<string>;
    /**
     * Eip bound rsc vip of the resource instance.
     */
    public /*out*/ readonly eipBoundRscVip!: pulumi.Output<string>;
    /**
     * Expired time of the resource instance.
     */
    public /*out*/ readonly expiredTime!: pulumi.Output<string>;
    /**
     * Modify time of the resource instance.
     */
    public /*out*/ readonly modifyTime!: pulumi.Output<string>;
    /**
     * Protection status of the resource instance.
     */
    public /*out*/ readonly protectionStatus!: pulumi.Output<string>;
    /**
     * ID of the resource.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * Region of the resource instance.
     */
    public /*out*/ readonly resourceRegion!: pulumi.Output<string>;

    /**
     * Create a Eip resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EipArgs | EipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EipState | undefined;
            resourceInputs["bindResourceId"] = state ? state.bindResourceId : undefined;
            resourceInputs["bindResourceRegion"] = state ? state.bindResourceRegion : undefined;
            resourceInputs["bindResourceType"] = state ? state.bindResourceType : undefined;
            resourceInputs["createdTime"] = state ? state.createdTime : undefined;
            resourceInputs["eip"] = state ? state.eip : undefined;
            resourceInputs["eipAddressStatus"] = state ? state.eipAddressStatus : undefined;
            resourceInputs["eipBoundRscEni"] = state ? state.eipBoundRscEni : undefined;
            resourceInputs["eipBoundRscIns"] = state ? state.eipBoundRscIns : undefined;
            resourceInputs["eipBoundRscVip"] = state ? state.eipBoundRscVip : undefined;
            resourceInputs["expiredTime"] = state ? state.expiredTime : undefined;
            resourceInputs["modifyTime"] = state ? state.modifyTime : undefined;
            resourceInputs["protectionStatus"] = state ? state.protectionStatus : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["resourceRegion"] = state ? state.resourceRegion : undefined;
        } else {
            const args = argsOrState as EipArgs | undefined;
            if ((!args || args.bindResourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindResourceId'");
            }
            if ((!args || args.bindResourceRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindResourceRegion'");
            }
            if ((!args || args.bindResourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindResourceType'");
            }
            if ((!args || args.eip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eip'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            resourceInputs["bindResourceId"] = args ? args.bindResourceId : undefined;
            resourceInputs["bindResourceRegion"] = args ? args.bindResourceRegion : undefined;
            resourceInputs["bindResourceType"] = args ? args.bindResourceType : undefined;
            resourceInputs["eip"] = args ? args.eip : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["eipAddressStatus"] = undefined /*out*/;
            resourceInputs["eipBoundRscEni"] = undefined /*out*/;
            resourceInputs["eipBoundRscIns"] = undefined /*out*/;
            resourceInputs["eipBoundRscVip"] = undefined /*out*/;
            resourceInputs["expiredTime"] = undefined /*out*/;
            resourceInputs["modifyTime"] = undefined /*out*/;
            resourceInputs["protectionStatus"] = undefined /*out*/;
            resourceInputs["resourceRegion"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Eip.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Eip resources.
 */
export interface EipState {
    /**
     * Resource id to bind.
     */
    bindResourceId?: pulumi.Input<string>;
    /**
     * Resource region to bind.
     */
    bindResourceRegion?: pulumi.Input<string>;
    /**
     * Resource type to bind, value range [`clb`, `cvm`].
     */
    bindResourceType?: pulumi.Input<string>;
    /**
     * Created time of the resource instance.
     */
    createdTime?: pulumi.Input<string>;
    /**
     * Eip of the resource.
     */
    eip?: pulumi.Input<string>;
    /**
     * Eip address status of the resource instance.
     */
    eipAddressStatus?: pulumi.Input<string>;
    /**
     * Eip bound rsc eni of the resource instance.
     */
    eipBoundRscEni?: pulumi.Input<string>;
    /**
     * Eip bound rsc ins of the resource instance.
     */
    eipBoundRscIns?: pulumi.Input<string>;
    /**
     * Eip bound rsc vip of the resource instance.
     */
    eipBoundRscVip?: pulumi.Input<string>;
    /**
     * Expired time of the resource instance.
     */
    expiredTime?: pulumi.Input<string>;
    /**
     * Modify time of the resource instance.
     */
    modifyTime?: pulumi.Input<string>;
    /**
     * Protection status of the resource instance.
     */
    protectionStatus?: pulumi.Input<string>;
    /**
     * ID of the resource.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * Region of the resource instance.
     */
    resourceRegion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Eip resource.
 */
export interface EipArgs {
    /**
     * Resource id to bind.
     */
    bindResourceId: pulumi.Input<string>;
    /**
     * Resource region to bind.
     */
    bindResourceRegion: pulumi.Input<string>;
    /**
     * Resource type to bind, value range [`clb`, `cvm`].
     */
    bindResourceType: pulumi.Input<string>;
    /**
     * Eip of the resource.
     */
    eip: pulumi.Input<string>;
    /**
     * ID of the resource.
     */
    resourceId: pulumi.Input<string>;
}
