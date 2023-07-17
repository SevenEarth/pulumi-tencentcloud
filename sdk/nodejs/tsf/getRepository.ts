// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tsf repository
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const repository = pulumi.output(tencentcloud.Tsf.getRepository({
 *     repositoryType: "default",
 *     searchWord: "test",
 * }));
 * ```
 */
export function getRepository(args?: GetRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Tsf/getRepository:getRepository", {
        "repositoryType": args.repositoryType,
        "resultOutputFile": args.resultOutputFile,
        "searchWord": args.searchWord,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepository.
 */
export interface GetRepositoryArgs {
    /**
     * Repository type (default Repository: default, private Repository: private).
     */
    repositoryType?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Query keywords (search by Repository name).
     */
    searchWord?: string;
}

/**
 * A collection of values returned by getRepository.
 */
export interface GetRepositoryResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Repository type (default Repository: default, private Repository: private).
     */
    readonly repositoryType?: string;
    readonly resultOutputFile?: string;
    /**
     * A list of Repository information that meets the query criteria.
     */
    readonly results: outputs.Tsf.GetRepositoryResult[];
    readonly searchWord?: string;
}

export function getRepositoryOutput(args?: GetRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRepositoryResult> {
    return pulumi.output(args).apply(a => getRepository(a, opts))
}

/**
 * A collection of arguments for invoking getRepository.
 */
export interface GetRepositoryOutputArgs {
    /**
     * Repository type (default Repository: default, private Repository: private).
     */
    repositoryType?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Query keywords (search by Repository name).
     */
    searchWord?: pulumi.Input<string>;
}
