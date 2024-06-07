// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cam mfaFlag
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const info = tencentcloud.User.getInfo({});
 * const mfaFlag = new tencentcloud.cam.MfaFlag("mfaFlag", {
 *     opUin: info.then(info => info.uin),
 *     loginFlag: {
 *         phone: 0,
 *         stoken: 1,
 *         wechat: 0,
 *     },
 *     actionFlag: {
 *         phone: 0,
 *         stoken: 1,
 *         wechat: 0,
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * cam mfa_flag can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Cam/mfaFlag:MfaFlag mfa_flag mfa_flag_id
 * ```
 */
export class MfaFlag extends pulumi.CustomResource {
    /**
     * Get an existing MfaFlag resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MfaFlagState, opts?: pulumi.CustomResourceOptions): MfaFlag {
        return new MfaFlag(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cam/mfaFlag:MfaFlag';

    /**
     * Returns true if the given object is an instance of MfaFlag.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MfaFlag {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MfaFlag.__pulumiType;
    }

    /**
     * Action flag setting.
     */
    public readonly actionFlag!: pulumi.Output<outputs.Cam.MfaFlagActionFlag | undefined>;
    /**
     * Login flag setting.
     */
    public readonly loginFlag!: pulumi.Output<outputs.Cam.MfaFlagLoginFlag | undefined>;
    /**
     * Operate uin.
     */
    public readonly opUin!: pulumi.Output<number>;

    /**
     * Create a MfaFlag resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MfaFlagArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MfaFlagArgs | MfaFlagState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MfaFlagState | undefined;
            resourceInputs["actionFlag"] = state ? state.actionFlag : undefined;
            resourceInputs["loginFlag"] = state ? state.loginFlag : undefined;
            resourceInputs["opUin"] = state ? state.opUin : undefined;
        } else {
            const args = argsOrState as MfaFlagArgs | undefined;
            if ((!args || args.opUin === undefined) && !opts.urn) {
                throw new Error("Missing required property 'opUin'");
            }
            resourceInputs["actionFlag"] = args ? args.actionFlag : undefined;
            resourceInputs["loginFlag"] = args ? args.loginFlag : undefined;
            resourceInputs["opUin"] = args ? args.opUin : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MfaFlag.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MfaFlag resources.
 */
export interface MfaFlagState {
    /**
     * Action flag setting.
     */
    actionFlag?: pulumi.Input<inputs.Cam.MfaFlagActionFlag>;
    /**
     * Login flag setting.
     */
    loginFlag?: pulumi.Input<inputs.Cam.MfaFlagLoginFlag>;
    /**
     * Operate uin.
     */
    opUin?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a MfaFlag resource.
 */
export interface MfaFlagArgs {
    /**
     * Action flag setting.
     */
    actionFlag?: pulumi.Input<inputs.Cam.MfaFlagActionFlag>;
    /**
     * Login flag setting.
     */
    loginFlag?: pulumi.Input<inputs.Cam.MfaFlagLoginFlag>;
    /**
     * Operate uin.
     */
    opUin: pulumi.Input<number>;
}
