// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a redis backupDownloadRestriction
 *
 * ## Example Usage
 * ### Modify the network information and address of the current region backup file download
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const foo = new tencentcloud.redis.BackupDownloadRestriction("foo", {
 *     limitType: "Customize",
 *     vpcComparisonSymbol: "In",
 *     ipComparisonSymbol: "In",
 *     limitVpcs: [{
 *         region: "ap-guangzhou",
 *         vpcLists: [_var.vpc_id],
 *     }],
 *     limitIps: [
 *         "10.1.1.12",
 *         "10.1.1.13",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * redis backup_download_restriction can be imported using the region, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Redis/backupDownloadRestriction:BackupDownloadRestriction foo ap-guangzhou
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
    public static readonly __pulumiType = 'tencentcloud:Redis/backupDownloadRestriction:BackupDownloadRestriction';

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
     * Identifies whether the customized LimitIP address can download the backup file.- In: Custom IP addresses are available for download.- NotIn: Custom IPs are not available for download.
     */
    public readonly ipComparisonSymbol!: pulumi.Output<string | undefined>;
    /**
     * A custom VPC IP address for downloadable backup files.If the parameter LimitType is **Customize**, you need to configure this parameter.
     */
    public readonly limitIps!: pulumi.Output<string[] | undefined>;
    /**
     * Types of network restrictions for downloading backup files:- NoLimit: There is no limit, and backup files can be downloaded from both Tencent Cloud and internal and external networks.- LimitOnlyIntranet: Only intranet addresses automatically assigned by Tencent Cloud can download backup files.- Customize: refers to a user-defined private network downloadable backup file.
     */
    public readonly limitType!: pulumi.Output<string>;
    /**
     * A custom VPC ID for a downloadable backup file.If the parameter LimitType is **Customize**, you need to configure this parameter.
     */
    public readonly limitVpcs!: pulumi.Output<outputs.Redis.BackupDownloadRestrictionLimitVpc[] | undefined>;
    /**
     * This parameter only supports entering In, which means that the custom LimitVpc can download the backup file.
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
     * Identifies whether the customized LimitIP address can download the backup file.- In: Custom IP addresses are available for download.- NotIn: Custom IPs are not available for download.
     */
    ipComparisonSymbol?: pulumi.Input<string>;
    /**
     * A custom VPC IP address for downloadable backup files.If the parameter LimitType is **Customize**, you need to configure this parameter.
     */
    limitIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Types of network restrictions for downloading backup files:- NoLimit: There is no limit, and backup files can be downloaded from both Tencent Cloud and internal and external networks.- LimitOnlyIntranet: Only intranet addresses automatically assigned by Tencent Cloud can download backup files.- Customize: refers to a user-defined private network downloadable backup file.
     */
    limitType?: pulumi.Input<string>;
    /**
     * A custom VPC ID for a downloadable backup file.If the parameter LimitType is **Customize**, you need to configure this parameter.
     */
    limitVpcs?: pulumi.Input<pulumi.Input<inputs.Redis.BackupDownloadRestrictionLimitVpc>[]>;
    /**
     * This parameter only supports entering In, which means that the custom LimitVpc can download the backup file.
     */
    vpcComparisonSymbol?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackupDownloadRestriction resource.
 */
export interface BackupDownloadRestrictionArgs {
    /**
     * Identifies whether the customized LimitIP address can download the backup file.- In: Custom IP addresses are available for download.- NotIn: Custom IPs are not available for download.
     */
    ipComparisonSymbol?: pulumi.Input<string>;
    /**
     * A custom VPC IP address for downloadable backup files.If the parameter LimitType is **Customize**, you need to configure this parameter.
     */
    limitIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Types of network restrictions for downloading backup files:- NoLimit: There is no limit, and backup files can be downloaded from both Tencent Cloud and internal and external networks.- LimitOnlyIntranet: Only intranet addresses automatically assigned by Tencent Cloud can download backup files.- Customize: refers to a user-defined private network downloadable backup file.
     */
    limitType: pulumi.Input<string>;
    /**
     * A custom VPC ID for a downloadable backup file.If the parameter LimitType is **Customize**, you need to configure this parameter.
     */
    limitVpcs?: pulumi.Input<pulumi.Input<inputs.Redis.BackupDownloadRestrictionLimitVpc>[]>;
    /**
     * This parameter only supports entering In, which means that the custom LimitVpc can download the backup file.
     */
    vpcComparisonSymbol?: pulumi.Input<string>;
}
