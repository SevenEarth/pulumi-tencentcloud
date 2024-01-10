// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ssl updateCertificateRecordRollback
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const updateCertificateRecordRollback = new tencentcloud.Ssl.UpdateCertificateRecordRollbackOperation("update_certificate_record_rollback", {
 *     deployRecordId: "1603",
 * });
 * ```
 *
 * ## Import
 *
 * ssl update_certificate_record_rollback can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Ssl/updateCertificateRecordRollbackOperation:UpdateCertificateRecordRollbackOperation update_certificate_record_rollback update_certificate_record_rollback_id
 * ```
 */
export class UpdateCertificateRecordRollbackOperation extends pulumi.CustomResource {
    /**
     * Get an existing UpdateCertificateRecordRollbackOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UpdateCertificateRecordRollbackOperationState, opts?: pulumi.CustomResourceOptions): UpdateCertificateRecordRollbackOperation {
        return new UpdateCertificateRecordRollbackOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ssl/updateCertificateRecordRollbackOperation:UpdateCertificateRecordRollbackOperation';

    /**
     * Returns true if the given object is an instance of UpdateCertificateRecordRollbackOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UpdateCertificateRecordRollbackOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UpdateCertificateRecordRollbackOperation.__pulumiType;
    }

    /**
     * Deployment record ID to be rolled back.
     */
    public readonly deployRecordId!: pulumi.Output<string | undefined>;

    /**
     * Create a UpdateCertificateRecordRollbackOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UpdateCertificateRecordRollbackOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UpdateCertificateRecordRollbackOperationArgs | UpdateCertificateRecordRollbackOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UpdateCertificateRecordRollbackOperationState | undefined;
            resourceInputs["deployRecordId"] = state ? state.deployRecordId : undefined;
        } else {
            const args = argsOrState as UpdateCertificateRecordRollbackOperationArgs | undefined;
            resourceInputs["deployRecordId"] = args ? args.deployRecordId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UpdateCertificateRecordRollbackOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UpdateCertificateRecordRollbackOperation resources.
 */
export interface UpdateCertificateRecordRollbackOperationState {
    /**
     * Deployment record ID to be rolled back.
     */
    deployRecordId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UpdateCertificateRecordRollbackOperation resource.
 */
export interface UpdateCertificateRecordRollbackOperationArgs {
    /**
     * Deployment record ID to be rolled back.
     */
    deployRecordId?: pulumi.Input<string>;
}
