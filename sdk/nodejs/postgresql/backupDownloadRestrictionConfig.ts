// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a postgresql backupDownloadRestrictionConfig
 *
 * ## Example Usage
 * ### Unlimit the restriction of the backup file download.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const backupDownloadRestrictionConfig = new tencentcloud.Postgresql.BackupDownloadRestrictionConfig("backup_download_restriction_config", {
 *     restrictionType: "NONE",
 * });
 * ```
 * ### Set the download only to allow the intranet downloads.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const backupDownloadRestrictionConfig = new tencentcloud.Postgresql.BackupDownloadRestrictionConfig("backup_download_restriction_config", {
 *     restrictionType: "INTRANET",
 * });
 * ```
 * ### Restrict the backup file download by customizing.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const pgVpc = new tencentcloud.vpc.Instance("pgVpc", {cidrBlock: _var.vpc_cidr});
 * const backupDownloadRestrictionConfig = new tencentcloud.postgresql.BackupDownloadRestrictionConfig("backupDownloadRestrictionConfig", {
 *     restrictionType: "CUSTOMIZE",
 *     vpcRestrictionEffect: "DENY",
 *     vpcIdSets: [tencentcloud_vpc.pg_vpc2.id],
 *     ipRestrictionEffect: "DENY",
 *     ipSets: ["192.168.0.0"],
 * });
 * ```
 *
 * ## Import
 *
 * postgresql backup_download_restriction_config can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Postgresql/backupDownloadRestrictionConfig:BackupDownloadRestrictionConfig backup_download_restriction_config backup_download_restriction_config_id
 * ```
 */
export class BackupDownloadRestrictionConfig extends pulumi.CustomResource {
    /**
     * Get an existing BackupDownloadRestrictionConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackupDownloadRestrictionConfigState, opts?: pulumi.CustomResourceOptions): BackupDownloadRestrictionConfig {
        return new BackupDownloadRestrictionConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Postgresql/backupDownloadRestrictionConfig:BackupDownloadRestrictionConfig';

    /**
     * Returns true if the given object is an instance of BackupDownloadRestrictionConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackupDownloadRestrictionConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackupDownloadRestrictionConfig.__pulumiType;
    }

    /**
     * ip limit Strategy: ALLOW, DENY.
     */
    public readonly ipRestrictionEffect!: pulumi.Output<string | undefined>;
    /**
     * The list of ips that are allowed or denied to download backup files.
     */
    public readonly ipSets!: pulumi.Output<string[] | undefined>;
    /**
     * Backup file download restriction type: NONE:Unlimited, both internal and external networks can be downloaded. INTRANET:Only intranet downloads are allowed. CUSTOMIZE:Customize the vpc or ip that limits downloads.
     */
    public readonly restrictionType!: pulumi.Output<string>;
    /**
     * The list of vpcIds that allow or deny downloading of backup files.
     */
    public readonly vpcIdSets!: pulumi.Output<string[] | undefined>;
    /**
     * vpc limit Strategy: ALLOW, DENY.
     */
    public readonly vpcRestrictionEffect!: pulumi.Output<string | undefined>;

    /**
     * Create a BackupDownloadRestrictionConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupDownloadRestrictionConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackupDownloadRestrictionConfigArgs | BackupDownloadRestrictionConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackupDownloadRestrictionConfigState | undefined;
            resourceInputs["ipRestrictionEffect"] = state ? state.ipRestrictionEffect : undefined;
            resourceInputs["ipSets"] = state ? state.ipSets : undefined;
            resourceInputs["restrictionType"] = state ? state.restrictionType : undefined;
            resourceInputs["vpcIdSets"] = state ? state.vpcIdSets : undefined;
            resourceInputs["vpcRestrictionEffect"] = state ? state.vpcRestrictionEffect : undefined;
        } else {
            const args = argsOrState as BackupDownloadRestrictionConfigArgs | undefined;
            if ((!args || args.restrictionType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restrictionType'");
            }
            resourceInputs["ipRestrictionEffect"] = args ? args.ipRestrictionEffect : undefined;
            resourceInputs["ipSets"] = args ? args.ipSets : undefined;
            resourceInputs["restrictionType"] = args ? args.restrictionType : undefined;
            resourceInputs["vpcIdSets"] = args ? args.vpcIdSets : undefined;
            resourceInputs["vpcRestrictionEffect"] = args ? args.vpcRestrictionEffect : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BackupDownloadRestrictionConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackupDownloadRestrictionConfig resources.
 */
export interface BackupDownloadRestrictionConfigState {
    /**
     * ip limit Strategy: ALLOW, DENY.
     */
    ipRestrictionEffect?: pulumi.Input<string>;
    /**
     * The list of ips that are allowed or denied to download backup files.
     */
    ipSets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Backup file download restriction type: NONE:Unlimited, both internal and external networks can be downloaded. INTRANET:Only intranet downloads are allowed. CUSTOMIZE:Customize the vpc or ip that limits downloads.
     */
    restrictionType?: pulumi.Input<string>;
    /**
     * The list of vpcIds that allow or deny downloading of backup files.
     */
    vpcIdSets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * vpc limit Strategy: ALLOW, DENY.
     */
    vpcRestrictionEffect?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackupDownloadRestrictionConfig resource.
 */
export interface BackupDownloadRestrictionConfigArgs {
    /**
     * ip limit Strategy: ALLOW, DENY.
     */
    ipRestrictionEffect?: pulumi.Input<string>;
    /**
     * The list of ips that are allowed or denied to download backup files.
     */
    ipSets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Backup file download restriction type: NONE:Unlimited, both internal and external networks can be downloaded. INTRANET:Only intranet downloads are allowed. CUSTOMIZE:Customize the vpc or ip that limits downloads.
     */
    restrictionType: pulumi.Input<string>;
    /**
     * The list of vpcIds that allow or deny downloading of backup files.
     */
    vpcIdSets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * vpc limit Strategy: ALLOW, DENY.
     */
    vpcRestrictionEffect?: pulumi.Input<string>;
}
