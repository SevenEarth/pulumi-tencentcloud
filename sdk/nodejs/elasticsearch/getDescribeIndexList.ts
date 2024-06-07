// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of elasticsearch index list
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const describeIndexList = tencentcloud.Elasticsearch.getDescribeIndexList({
 *     indexType: "normal",
 *     instanceId: "es-nni6pm4s",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDescribeIndexList(args: GetDescribeIndexListArgs, opts?: pulumi.InvokeOptions): Promise<GetDescribeIndexListResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Elasticsearch/getDescribeIndexList:getDescribeIndexList", {
        "indexName": args.indexName,
        "indexStatusLists": args.indexStatusLists,
        "indexType": args.indexType,
        "instanceId": args.instanceId,
        "order": args.order,
        "orderBy": args.orderBy,
        "password": args.password,
        "resultOutputFile": args.resultOutputFile,
        "username": args.username,
    }, opts);
}

/**
 * A collection of arguments for invoking getDescribeIndexList.
 */
export interface GetDescribeIndexListArgs {
    /**
     * Index name. If you fill in the blanks, get all indexes.
     */
    indexName?: string;
    /**
     * Index status list.
     */
    indexStatusLists?: string[];
    /**
     * Index type. `auto`: Autonomous index; `normal`: General index.
     */
    indexType: string;
    /**
     * ES cluster id.
     */
    instanceId?: string;
    /**
     * Sort order, which supports asc and desc. The default is desc data format asc,desc.
     */
    order?: string;
    /**
     * Sort field. Support index name: IndexName, index storage: IndexStorage, index creation time: IndexCreateTime.
     */
    orderBy?: string;
    /**
     * Cluster access password.
     */
    password?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Cluster access user name.
     */
    username?: string;
}

/**
 * A collection of values returned by getDescribeIndexList.
 */
export interface GetDescribeIndexListResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Index metadata field.
     */
    readonly indexMetaFields: outputs.Elasticsearch.GetDescribeIndexListIndexMetaField[];
    /**
     * Index name.
     */
    readonly indexName?: string;
    readonly indexStatusLists?: string[];
    /**
     * Index type.
     */
    readonly indexType: string;
    readonly instanceId?: string;
    readonly order?: string;
    readonly orderBy?: string;
    readonly password?: string;
    readonly resultOutputFile?: string;
    readonly username?: string;
}
/**
 * Use this data source to query detailed information of elasticsearch index list
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const describeIndexList = tencentcloud.Elasticsearch.getDescribeIndexList({
 *     indexType: "normal",
 *     instanceId: "es-nni6pm4s",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDescribeIndexListOutput(args: GetDescribeIndexListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDescribeIndexListResult> {
    return pulumi.output(args).apply((a: any) => getDescribeIndexList(a, opts))
}

/**
 * A collection of arguments for invoking getDescribeIndexList.
 */
export interface GetDescribeIndexListOutputArgs {
    /**
     * Index name. If you fill in the blanks, get all indexes.
     */
    indexName?: pulumi.Input<string>;
    /**
     * Index status list.
     */
    indexStatusLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Index type. `auto`: Autonomous index; `normal`: General index.
     */
    indexType: pulumi.Input<string>;
    /**
     * ES cluster id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Sort order, which supports asc and desc. The default is desc data format asc,desc.
     */
    order?: pulumi.Input<string>;
    /**
     * Sort field. Support index name: IndexName, index storage: IndexStorage, index creation time: IndexCreateTime.
     */
    orderBy?: pulumi.Input<string>;
    /**
     * Cluster access password.
     */
    password?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Cluster access user name.
     */
    username?: pulumi.Input<string>;
}
