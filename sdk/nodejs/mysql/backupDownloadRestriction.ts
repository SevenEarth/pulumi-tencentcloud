// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mysql backupDownloadRestriction
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const backupDownloadRestriction = new tencentcloud.Mysql.BackupDownloadRestriction("backup_download_restriction", {
 *     ipComparisonSymbol: "In",
 *     limitIps: ["127.0.0.1"],
 *     limitType: "Customize",
 *     limitVpcs: [{
 *         region: "ap-guangzhou",
 *         vpcLists: ["vpc-4owdpnwr"],
 *     }],
 *     vpcComparisonSymbol: "In",
 * });
 * ```
 *
 * ## Import
 *
 * mysql backup_download_restriction can be imported using the "BackupDownloadRestriction", as follows.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Mysql/backupDownloadRestriction:BackupDownloadRestriction backup_download_restriction BackupDownloadRestriction
 * ```
 */
export class BackupDownloadRestriction extends pulumi.CustomResource {
    /**
     * Get an existing BackupDownloadRestriction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackupDownloadRestrictionState, opts?: pulumi.CustomResourceOptions): BackupDownloadRestriction {
        return new BackupDownloadRestriction(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mysql/backupDownloadRestriction:BackupDownloadRestriction';

    /**
     * Returns true if the given object is an instance of BackupDownloadRestriction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackupDownloadRestriction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackupDownloadRestriction.__pulumiType;
    }

    /**
     * In: The specified ip can be downloaded; NotIn: The specified ip cannot be downloaded. The default is In.
     */
    public readonly ipComparisonSymbol!: pulumi.Output<string | undefined>;
    /**
     * ip settings to limit downloads.
     */
    public readonly limitIps!: pulumi.Output<string[] | undefined>;
    /**
     * NoLimit No limit, both internal and external networks can be downloaded; LimitOnlyIntranet Only intranet can be downloaded; Customize user-defined vpc:ip can be downloaded. LimitVpc and LimitIp can be set only when the value is Customize.
     */
    public readonly limitType!: pulumi.Output<string>;
    /**
     * vpc settings to limit downloads.
     */
    public readonly limitVpcs!: pulumi.Output<outputs.Mysql.BackupDownloadRestrictionLimitVpc[] | undefined>;
    /**
     * This parameter only supports In, which means that the vpc specified by LimitVpc can be downloaded. The default is In.
     */
    public readonly vpcComparisonSymbol!: pulumi.Output<string | undefined>;

    /**
     * Create a BackupDownloadRestriction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupDownloadRestrictionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackupDownloadRestrictionArgs | BackupDownloadRestrictionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackupDownloadRestrictionState | undefined;
            resourceInputs["ipComparisonSymbol"] = state ? state.ipComparisonSymbol : undefined;
            resourceInputs["limitIps"] = state ? state.limitIps : undefined;
            resourceInputs["limitType"] = state ? state.limitType : undefined;
            resourceInputs["limitVpcs"] = state ? state.limitVpcs : undefined;
            resourceInputs["vpcComparisonSymbol"] = state ? state.vpcComparisonSymbol : undefined;
        } else {
            const args = argsOrState as BackupDownloadRestrictionArgs | undefined;
            if ((!args || args.limitType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'limitType'");
            }
            resourceInputs["ipComparisonSymbol"] = args ? args.ipComparisonSymbol : undefined;
            resourceInputs["limitIps"] = args ? args.limitIps : undefined;
            resourceInputs["limitType"] = args ? args.limitType : undefined;
            resourceInputs["limitVpcs"] = args ? args.limitVpcs : undefined;
            resourceInputs["vpcComparisonSymbol"] = args ? args.vpcComparisonSymbol : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BackupDownloadRestriction.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackupDownloadRestriction resources.
 */
export interface BackupDownloadRestrictionState {
    /**
     * In: The specified ip can be downloaded; NotIn: The specified ip cannot be downloaded. The default is In.
     */
    ipComparisonSymbol?: pulumi.Input<string>;
    /**
     * ip settings to limit downloads.
     */
    limitIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * NoLimit No limit, both internal and external networks can be downloaded; LimitOnlyIntranet Only intranet can be downloaded; Customize user-defined vpc:ip can be downloaded. LimitVpc and LimitIp can be set only when the value is Customize.
     */
    limitType?: pulumi.Input<string>;
    /**
     * vpc settings to limit downloads.
     */
    limitVpcs?: pulumi.Input<pulumi.Input<inputs.Mysql.BackupDownloadRestrictionLimitVpc>[]>;
    /**
     * This parameter only supports In, which means that the vpc specified by LimitVpc can be downloaded. The default is In.
     */
    vpcComparisonSymbol?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackupDownloadRestriction resource.
 */
export interface BackupDownloadRestrictionArgs {
    /**
     * In: The specified ip can be downloaded; NotIn: The specified ip cannot be downloaded. The default is In.
     */
    ipComparisonSymbol?: pulumi.Input<string>;
    /**
     * ip settings to limit downloads.
     */
    limitIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * NoLimit No limit, both internal and external networks can be downloaded; LimitOnlyIntranet Only intranet can be downloaded; Customize user-defined vpc:ip can be downloaded. LimitVpc and LimitIp can be set only when the value is Customize.
     */
    limitType: pulumi.Input<string>;
    /**
     * vpc settings to limit downloads.
     */
    limitVpcs?: pulumi.Input<pulumi.Input<inputs.Mysql.BackupDownloadRestrictionLimitVpc>[]>;
    /**
     * This parameter only supports In, which means that the vpc specified by LimitVpc can be downloaded. The default is In.
     */
    vpcComparisonSymbol?: pulumi.Input<string>;
}
