// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cvm launchTemplateVersion
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.Cvm.LaunchTemplateVersion("foo", {
 *     disableApiTermination: false,
 *     imageId: "img-9qrfy1xt",
 *     instanceType: "S5.MEDIUM4",
 *     launchTemplateId: "lt-r9ajalbi",
 *     launchTemplateVersionDescription: "version description",
 *     placement: {
 *         projectId: 0,
 *         zone: "ap-guangzhou-6",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * cvm launch_template_version can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Cvm/launchTemplateVersion:LaunchTemplateVersion launch_template_version ${launch_template_id}#${launch_template_version}
 * ```
 */
export class LaunchTemplateVersion extends pulumi.CustomResource {
    /**
     * Get an existing LaunchTemplateVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LaunchTemplateVersionState, opts?: pulumi.CustomResourceOptions): LaunchTemplateVersion {
        return new LaunchTemplateVersion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cvm/launchTemplateVersion:LaunchTemplateVersion';

    /**
     * Returns true if the given object is an instance of LaunchTemplateVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LaunchTemplateVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LaunchTemplateVersion.__pulumiType;
    }

    /**
     * Scheduled tasks.
     */
    public readonly actionTimer!: pulumi.Output<outputs.Cvm.LaunchTemplateVersionActionTimer>;
    /**
     * The role name of CAM.
     */
    public readonly camRoleName!: pulumi.Output<string>;
    /**
     * A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
     */
    public readonly clientToken!: pulumi.Output<string>;
    /**
     * The configuration information of instance data disks. If this parameter is not specified, no data disk will be purchased by default.
     */
    public readonly dataDisks!: pulumi.Output<outputs.Cvm.LaunchTemplateVersionDataDisk[]>;
    /**
     * Whether the termination protection is enabled. `TRUE`: Enable instance protection, which means that this instance can not be deleted by an API action.`FALSE`: Do not enable the instance protection. Default value: `FALSE`.
     */
    public readonly disableApiTermination!: pulumi.Output<boolean>;
    /**
     * Placement group ID. You can only specify one.
     */
    public readonly disasterRecoverGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Whether the request is a dry run only.
     */
    public readonly dryRun!: pulumi.Output<boolean>;
    /**
     * Enhanced service. You can use this parameter to specify whether to enable services such as Anti-DDoS and Cloud Monitor. If this parameter is not specified, Cloud Monitor and Anti-DDoS are enabled for public images by default.
     */
    public readonly enhancedService!: pulumi.Output<outputs.Cvm.LaunchTemplateVersionEnhancedService>;
    /**
     * Hostname of a CVM.
     */
    public readonly hostName!: pulumi.Output<string>;
    /**
     * HPC cluster ID. The HPC cluster must and can only be specified for a high-performance computing instance.
     */
    public readonly hpcClusterId!: pulumi.Output<string>;
    /**
     * Image ID.
     */
    public readonly imageId!: pulumi.Output<string>;
    /**
     * Describes the billing method of an instance.
     */
    public readonly instanceChargePrepaid!: pulumi.Output<outputs.Cvm.LaunchTemplateVersionInstanceChargePrepaid>;
    /**
     * The charge type of instance.
     */
    public readonly instanceChargeType!: pulumi.Output<string>;
    /**
     * The number of instances to be purchased.
     */
    public readonly instanceCount!: pulumi.Output<number>;
    /**
     * Options related to bidding requests.
     */
    public readonly instanceMarketOptions!: pulumi.Output<outputs.Cvm.LaunchTemplateVersionInstanceMarketOptions>;
    /**
     * Instance name to be displayed.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * The type of the instance. If this parameter is not specified, the system will dynamically specify the default model according to the resource sales in the current region.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * Describes the accessibility of an instance in the public network, including its network billing method, maximum bandwidth, etc.
     */
    public readonly internetAccessible!: pulumi.Output<outputs.Cvm.LaunchTemplateVersionInternetAccessible>;
    /**
     * Instance launch template ID. This parameter is used as a basis for creating new template versions.
     */
    public readonly launchTemplateId!: pulumi.Output<string>;
    /**
     * This parameter, when specified, is used to create instance launch templates. If this parameter is not specified, the default version will be used.
     */
    public readonly launchTemplateVersion!: pulumi.Output<number>;
    /**
     * Description of instance launch template versions. This parameter can contain 2-256 characters.
     */
    public readonly launchTemplateVersionDescription!: pulumi.Output<string>;
    /**
     * Describes login settings of an instance.
     */
    public readonly loginSettings!: pulumi.Output<outputs.Cvm.LaunchTemplateVersionLoginSettings>;
    /**
     * Location of the instance. You can use this parameter to specify the attributes of the instance, such as its availability zone, project, and CDH (for dedicated CVMs).
     */
    public readonly placement!: pulumi.Output<outputs.Cvm.LaunchTemplateVersionPlacement>;
    /**
     * Security groups to which the instance belongs. If this parameter is not specified, the instance will be associated with default security groups.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * System disk configuration information of the instance. If this parameter is not specified, it is assigned according to the system default.
     */
    public readonly systemDisk!: pulumi.Output<outputs.Cvm.LaunchTemplateVersionSystemDisk>;
    /**
     * Description of tags associated with resource instances during instance creation.
     */
    public readonly tagSpecifications!: pulumi.Output<outputs.Cvm.LaunchTemplateVersionTagSpecification[]>;
    /**
     * User data provided to the instance. This parameter needs to be encoded in base64 format with the maximum size of 16 KB.
     */
    public readonly userData!: pulumi.Output<string>;
    /**
     * Describes information on VPC, including subnets, IP addresses, etc.
     */
    public readonly virtualPrivateCloud!: pulumi.Output<outputs.Cvm.LaunchTemplateVersionVirtualPrivateCloud>;

    /**
     * Create a LaunchTemplateVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LaunchTemplateVersionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LaunchTemplateVersionArgs | LaunchTemplateVersionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LaunchTemplateVersionState | undefined;
            resourceInputs["actionTimer"] = state ? state.actionTimer : undefined;
            resourceInputs["camRoleName"] = state ? state.camRoleName : undefined;
            resourceInputs["clientToken"] = state ? state.clientToken : undefined;
            resourceInputs["dataDisks"] = state ? state.dataDisks : undefined;
            resourceInputs["disableApiTermination"] = state ? state.disableApiTermination : undefined;
            resourceInputs["disasterRecoverGroupIds"] = state ? state.disasterRecoverGroupIds : undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["enhancedService"] = state ? state.enhancedService : undefined;
            resourceInputs["hostName"] = state ? state.hostName : undefined;
            resourceInputs["hpcClusterId"] = state ? state.hpcClusterId : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["instanceChargePrepaid"] = state ? state.instanceChargePrepaid : undefined;
            resourceInputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            resourceInputs["instanceCount"] = state ? state.instanceCount : undefined;
            resourceInputs["instanceMarketOptions"] = state ? state.instanceMarketOptions : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["internetAccessible"] = state ? state.internetAccessible : undefined;
            resourceInputs["launchTemplateId"] = state ? state.launchTemplateId : undefined;
            resourceInputs["launchTemplateVersion"] = state ? state.launchTemplateVersion : undefined;
            resourceInputs["launchTemplateVersionDescription"] = state ? state.launchTemplateVersionDescription : undefined;
            resourceInputs["loginSettings"] = state ? state.loginSettings : undefined;
            resourceInputs["placement"] = state ? state.placement : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["systemDisk"] = state ? state.systemDisk : undefined;
            resourceInputs["tagSpecifications"] = state ? state.tagSpecifications : undefined;
            resourceInputs["userData"] = state ? state.userData : undefined;
            resourceInputs["virtualPrivateCloud"] = state ? state.virtualPrivateCloud : undefined;
        } else {
            const args = argsOrState as LaunchTemplateVersionArgs | undefined;
            if ((!args || args.launchTemplateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'launchTemplateId'");
            }
            if ((!args || args.placement === undefined) && !opts.urn) {
                throw new Error("Missing required property 'placement'");
            }
            resourceInputs["actionTimer"] = args ? args.actionTimer : undefined;
            resourceInputs["camRoleName"] = args ? args.camRoleName : undefined;
            resourceInputs["clientToken"] = args ? args.clientToken : undefined;
            resourceInputs["dataDisks"] = args ? args.dataDisks : undefined;
            resourceInputs["disableApiTermination"] = args ? args.disableApiTermination : undefined;
            resourceInputs["disasterRecoverGroupIds"] = args ? args.disasterRecoverGroupIds : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["enhancedService"] = args ? args.enhancedService : undefined;
            resourceInputs["hostName"] = args ? args.hostName : undefined;
            resourceInputs["hpcClusterId"] = args ? args.hpcClusterId : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["instanceChargePrepaid"] = args ? args.instanceChargePrepaid : undefined;
            resourceInputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            resourceInputs["instanceCount"] = args ? args.instanceCount : undefined;
            resourceInputs["instanceMarketOptions"] = args ? args.instanceMarketOptions : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["internetAccessible"] = args ? args.internetAccessible : undefined;
            resourceInputs["launchTemplateId"] = args ? args.launchTemplateId : undefined;
            resourceInputs["launchTemplateVersion"] = args ? args.launchTemplateVersion : undefined;
            resourceInputs["launchTemplateVersionDescription"] = args ? args.launchTemplateVersionDescription : undefined;
            resourceInputs["loginSettings"] = args ? args.loginSettings : undefined;
            resourceInputs["placement"] = args ? args.placement : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["systemDisk"] = args ? args.systemDisk : undefined;
            resourceInputs["tagSpecifications"] = args ? args.tagSpecifications : undefined;
            resourceInputs["userData"] = args ? args.userData : undefined;
            resourceInputs["virtualPrivateCloud"] = args ? args.virtualPrivateCloud : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LaunchTemplateVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LaunchTemplateVersion resources.
 */
export interface LaunchTemplateVersionState {
    /**
     * Scheduled tasks.
     */
    actionTimer?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionActionTimer>;
    /**
     * The role name of CAM.
     */
    camRoleName?: pulumi.Input<string>;
    /**
     * A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
     */
    clientToken?: pulumi.Input<string>;
    /**
     * The configuration information of instance data disks. If this parameter is not specified, no data disk will be purchased by default.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.Cvm.LaunchTemplateVersionDataDisk>[]>;
    /**
     * Whether the termination protection is enabled. `TRUE`: Enable instance protection, which means that this instance can not be deleted by an API action.`FALSE`: Do not enable the instance protection. Default value: `FALSE`.
     */
    disableApiTermination?: pulumi.Input<boolean>;
    /**
     * Placement group ID. You can only specify one.
     */
    disasterRecoverGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the request is a dry run only.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Enhanced service. You can use this parameter to specify whether to enable services such as Anti-DDoS and Cloud Monitor. If this parameter is not specified, Cloud Monitor and Anti-DDoS are enabled for public images by default.
     */
    enhancedService?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionEnhancedService>;
    /**
     * Hostname of a CVM.
     */
    hostName?: pulumi.Input<string>;
    /**
     * HPC cluster ID. The HPC cluster must and can only be specified for a high-performance computing instance.
     */
    hpcClusterId?: pulumi.Input<string>;
    /**
     * Image ID.
     */
    imageId?: pulumi.Input<string>;
    /**
     * Describes the billing method of an instance.
     */
    instanceChargePrepaid?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionInstanceChargePrepaid>;
    /**
     * The charge type of instance.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The number of instances to be purchased.
     */
    instanceCount?: pulumi.Input<number>;
    /**
     * Options related to bidding requests.
     */
    instanceMarketOptions?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionInstanceMarketOptions>;
    /**
     * Instance name to be displayed.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The type of the instance. If this parameter is not specified, the system will dynamically specify the default model according to the resource sales in the current region.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Describes the accessibility of an instance in the public network, including its network billing method, maximum bandwidth, etc.
     */
    internetAccessible?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionInternetAccessible>;
    /**
     * Instance launch template ID. This parameter is used as a basis for creating new template versions.
     */
    launchTemplateId?: pulumi.Input<string>;
    /**
     * This parameter, when specified, is used to create instance launch templates. If this parameter is not specified, the default version will be used.
     */
    launchTemplateVersion?: pulumi.Input<number>;
    /**
     * Description of instance launch template versions. This parameter can contain 2-256 characters.
     */
    launchTemplateVersionDescription?: pulumi.Input<string>;
    /**
     * Describes login settings of an instance.
     */
    loginSettings?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionLoginSettings>;
    /**
     * Location of the instance. You can use this parameter to specify the attributes of the instance, such as its availability zone, project, and CDH (for dedicated CVMs).
     */
    placement?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionPlacement>;
    /**
     * Security groups to which the instance belongs. If this parameter is not specified, the instance will be associated with default security groups.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * System disk configuration information of the instance. If this parameter is not specified, it is assigned according to the system default.
     */
    systemDisk?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionSystemDisk>;
    /**
     * Description of tags associated with resource instances during instance creation.
     */
    tagSpecifications?: pulumi.Input<pulumi.Input<inputs.Cvm.LaunchTemplateVersionTagSpecification>[]>;
    /**
     * User data provided to the instance. This parameter needs to be encoded in base64 format with the maximum size of 16 KB.
     */
    userData?: pulumi.Input<string>;
    /**
     * Describes information on VPC, including subnets, IP addresses, etc.
     */
    virtualPrivateCloud?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionVirtualPrivateCloud>;
}

/**
 * The set of arguments for constructing a LaunchTemplateVersion resource.
 */
export interface LaunchTemplateVersionArgs {
    /**
     * Scheduled tasks.
     */
    actionTimer?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionActionTimer>;
    /**
     * The role name of CAM.
     */
    camRoleName?: pulumi.Input<string>;
    /**
     * A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
     */
    clientToken?: pulumi.Input<string>;
    /**
     * The configuration information of instance data disks. If this parameter is not specified, no data disk will be purchased by default.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.Cvm.LaunchTemplateVersionDataDisk>[]>;
    /**
     * Whether the termination protection is enabled. `TRUE`: Enable instance protection, which means that this instance can not be deleted by an API action.`FALSE`: Do not enable the instance protection. Default value: `FALSE`.
     */
    disableApiTermination?: pulumi.Input<boolean>;
    /**
     * Placement group ID. You can only specify one.
     */
    disasterRecoverGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the request is a dry run only.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Enhanced service. You can use this parameter to specify whether to enable services such as Anti-DDoS and Cloud Monitor. If this parameter is not specified, Cloud Monitor and Anti-DDoS are enabled for public images by default.
     */
    enhancedService?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionEnhancedService>;
    /**
     * Hostname of a CVM.
     */
    hostName?: pulumi.Input<string>;
    /**
     * HPC cluster ID. The HPC cluster must and can only be specified for a high-performance computing instance.
     */
    hpcClusterId?: pulumi.Input<string>;
    /**
     * Image ID.
     */
    imageId?: pulumi.Input<string>;
    /**
     * Describes the billing method of an instance.
     */
    instanceChargePrepaid?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionInstanceChargePrepaid>;
    /**
     * The charge type of instance.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The number of instances to be purchased.
     */
    instanceCount?: pulumi.Input<number>;
    /**
     * Options related to bidding requests.
     */
    instanceMarketOptions?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionInstanceMarketOptions>;
    /**
     * Instance name to be displayed.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The type of the instance. If this parameter is not specified, the system will dynamically specify the default model according to the resource sales in the current region.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Describes the accessibility of an instance in the public network, including its network billing method, maximum bandwidth, etc.
     */
    internetAccessible?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionInternetAccessible>;
    /**
     * Instance launch template ID. This parameter is used as a basis for creating new template versions.
     */
    launchTemplateId: pulumi.Input<string>;
    /**
     * This parameter, when specified, is used to create instance launch templates. If this parameter is not specified, the default version will be used.
     */
    launchTemplateVersion?: pulumi.Input<number>;
    /**
     * Description of instance launch template versions. This parameter can contain 2-256 characters.
     */
    launchTemplateVersionDescription?: pulumi.Input<string>;
    /**
     * Describes login settings of an instance.
     */
    loginSettings?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionLoginSettings>;
    /**
     * Location of the instance. You can use this parameter to specify the attributes of the instance, such as its availability zone, project, and CDH (for dedicated CVMs).
     */
    placement: pulumi.Input<inputs.Cvm.LaunchTemplateVersionPlacement>;
    /**
     * Security groups to which the instance belongs. If this parameter is not specified, the instance will be associated with default security groups.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * System disk configuration information of the instance. If this parameter is not specified, it is assigned according to the system default.
     */
    systemDisk?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionSystemDisk>;
    /**
     * Description of tags associated with resource instances during instance creation.
     */
    tagSpecifications?: pulumi.Input<pulumi.Input<inputs.Cvm.LaunchTemplateVersionTagSpecification>[]>;
    /**
     * User data provided to the instance. This parameter needs to be encoded in base64 format with the maximum size of 16 KB.
     */
    userData?: pulumi.Input<string>;
    /**
     * Describes information on VPC, including subnets, IP addresses, etc.
     */
    virtualPrivateCloud?: pulumi.Input<inputs.Cvm.LaunchTemplateVersionVirtualPrivateCloud>;
}
