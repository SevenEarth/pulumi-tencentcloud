// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to operate a tcr image signature.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const exampleInstance = new tencentcloud.tcr.Instance("exampleInstance", {
 *     instanceType: "premium",
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * const exampleNamespace = new tencentcloud.tcr.Namespace("exampleNamespace", {
 *     instanceId: exampleInstance.id,
 *     isPublic: true,
 *     isAutoScan: true,
 *     isPreventVul: true,
 *     severity: "medium",
 *     cveWhitelistItems: [{
 *         cveId: "cve-xxxxx",
 *     }],
 * });
 * const exampleRepository = new tencentcloud.tcr.Repository("exampleRepository", {
 *     instanceId: exampleInstance.id,
 *     namespaceName: exampleNamespace.name,
 *     briefDesc: "111",
 *     description: "111111111111111111111111111111111111",
 * });
 * const exampleCreateImageSignatureOperation = new tencentcloud.tcr.CreateImageSignatureOperation("exampleCreateImageSignatureOperation", {
 *     registryId: exampleInstance.id,
 *     namespaceName: exampleNamespace.name,
 *     repositoryName: exampleRepository.name,
 *     imageVersion: "v1",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * tcr image_signature_operation can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Tcr/createImageSignatureOperation:CreateImageSignatureOperation image_signature_operation image_signature_operation_id
 * ```
 */
export class CreateImageSignatureOperation extends pulumi.CustomResource {
    /**
     * Get an existing CreateImageSignatureOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CreateImageSignatureOperationState, opts?: pulumi.CustomResourceOptions): CreateImageSignatureOperation {
        return new CreateImageSignatureOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tcr/createImageSignatureOperation:CreateImageSignatureOperation';

    /**
     * Returns true if the given object is an instance of CreateImageSignatureOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CreateImageSignatureOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CreateImageSignatureOperation.__pulumiType;
    }

    /**
     * image version name.
     */
    public readonly imageVersion!: pulumi.Output<string>;
    /**
     * namespace name.
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * instance id.
     */
    public readonly registryId!: pulumi.Output<string>;
    /**
     * repository name.
     */
    public readonly repositoryName!: pulumi.Output<string>;

    /**
     * Create a CreateImageSignatureOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CreateImageSignatureOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CreateImageSignatureOperationArgs | CreateImageSignatureOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CreateImageSignatureOperationState | undefined;
            resourceInputs["imageVersion"] = state ? state.imageVersion : undefined;
            resourceInputs["namespaceName"] = state ? state.namespaceName : undefined;
            resourceInputs["registryId"] = state ? state.registryId : undefined;
            resourceInputs["repositoryName"] = state ? state.repositoryName : undefined;
        } else {
            const args = argsOrState as CreateImageSignatureOperationArgs | undefined;
            if ((!args || args.imageVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageVersion'");
            }
            if ((!args || args.namespaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if ((!args || args.registryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registryId'");
            }
            if ((!args || args.repositoryName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repositoryName'");
            }
            resourceInputs["imageVersion"] = args ? args.imageVersion : undefined;
            resourceInputs["namespaceName"] = args ? args.namespaceName : undefined;
            resourceInputs["registryId"] = args ? args.registryId : undefined;
            resourceInputs["repositoryName"] = args ? args.repositoryName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CreateImageSignatureOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CreateImageSignatureOperation resources.
 */
export interface CreateImageSignatureOperationState {
    /**
     * image version name.
     */
    imageVersion?: pulumi.Input<string>;
    /**
     * namespace name.
     */
    namespaceName?: pulumi.Input<string>;
    /**
     * instance id.
     */
    registryId?: pulumi.Input<string>;
    /**
     * repository name.
     */
    repositoryName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CreateImageSignatureOperation resource.
 */
export interface CreateImageSignatureOperationArgs {
    /**
     * image version name.
     */
    imageVersion: pulumi.Input<string>;
    /**
     * namespace name.
     */
    namespaceName: pulumi.Input<string>;
    /**
     * instance id.
     */
    registryId: pulumi.Input<string>;
    /**
     * repository name.
     */
    repositoryName: pulumi.Input<string>;
}
