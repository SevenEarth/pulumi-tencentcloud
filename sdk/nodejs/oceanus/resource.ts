// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a oceanus resource
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.oceanus.Resource("example", {
 *     folderId: "folder-7ctl246z",
 *     remark: "remark.",
 *     resourceConfigRemark: "config remark.",
 *     resourceLoc: {
 *         param: {
 *             bucket: "keep-terraform-1257058945",
 *             path: "OceanusResource/junit-4.13.2.jar",
 *             region: "ap-guangzhou",
 *         },
 *         storageType: 1,
 *     },
 *     resourceType: 1,
 *     workSpaceId: "space-2idq8wbr",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class Resource extends pulumi.CustomResource {
    /**
     * Get an existing Resource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceState, opts?: pulumi.CustomResourceOptions): Resource {
        return new Resource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Oceanus/resource:Resource';

    /**
     * Returns true if the given object is an instance of Resource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Resource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Resource.__pulumiType;
    }

    /**
     * Folder id.
     */
    public readonly folderId!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Resource description.
     */
    public readonly remark!: pulumi.Output<string | undefined>;
    /**
     * Resource version description.
     */
    public readonly resourceConfigRemark!: pulumi.Output<string | undefined>;
    /**
     * Resource ID.
     */
    public /*out*/ readonly resourceId!: pulumi.Output<string>;
    /**
     * Resource location.
     */
    public readonly resourceLoc!: pulumi.Output<outputs.Oceanus.ResourceResourceLoc>;
    /**
     * Resource type, only support JAR now, value is 1.
     */
    public readonly resourceType!: pulumi.Output<number>;
    /**
     * Resource Version.
     */
    public /*out*/ readonly version!: pulumi.Output<number>;
    /**
     * Workspace serialId.
     */
    public readonly workSpaceId!: pulumi.Output<string | undefined>;

    /**
     * Create a Resource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceArgs | ResourceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResourceState | undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
            resourceInputs["resourceConfigRemark"] = state ? state.resourceConfigRemark : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["resourceLoc"] = state ? state.resourceLoc : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["workSpaceId"] = state ? state.workSpaceId : undefined;
        } else {
            const args = argsOrState as ResourceArgs | undefined;
            if ((!args || args.resourceLoc === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceLoc'");
            }
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
            resourceInputs["resourceConfigRemark"] = args ? args.resourceConfigRemark : undefined;
            resourceInputs["resourceLoc"] = args ? args.resourceLoc : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["workSpaceId"] = args ? args.workSpaceId : undefined;
            resourceInputs["resourceId"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Resource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Resource resources.
 */
export interface ResourceState {
    /**
     * Folder id.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Resource description.
     */
    remark?: pulumi.Input<string>;
    /**
     * Resource version description.
     */
    resourceConfigRemark?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    resourceLoc?: pulumi.Input<inputs.Oceanus.ResourceResourceLoc>;
    /**
     * Resource type, only support JAR now, value is 1.
     */
    resourceType?: pulumi.Input<number>;
    /**
     * Resource Version.
     */
    version?: pulumi.Input<number>;
    /**
     * Workspace serialId.
     */
    workSpaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Resource resource.
 */
export interface ResourceArgs {
    /**
     * Folder id.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Resource description.
     */
    remark?: pulumi.Input<string>;
    /**
     * Resource version description.
     */
    resourceConfigRemark?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    resourceLoc: pulumi.Input<inputs.Oceanus.ResourceResourceLoc>;
    /**
     * Resource type, only support JAR now, value is 1.
     */
    resourceType: pulumi.Input<number>;
    /**
     * Workspace serialId.
     */
    workSpaceId?: pulumi.Input<string>;
}
