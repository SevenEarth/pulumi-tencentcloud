// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway.Outputs
{

    [OutputType]
    public sealed class ImportOpenApiServiceConfigCosConfig
    {
        /// <summary>
        /// The API calls the backend COS method, and the optional values for the front-end request method and Action are:GET: GetObjectPUT: PutObjectPOST: PostObject, AppendObjectHEAD: HeadObjectDELETE: DeleteObject.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// The API calls the signature switch of the backend COS, which defaults to false.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly bool? Authorization;
        /// <summary>
        /// The bucket name of the API backend COS.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string? BucketName;
        /// <summary>
        /// Path matching mode for API backend COS, optional values:BackEndPath: Backend path matchingFullPath: Full Path MatchingThe default value is: BackEndPathNote: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string? PathMatchMode;

        [OutputConstructor]
        private ImportOpenApiServiceConfigCosConfig(
            string? action,

            bool? authorization,

            string? bucketName,

            string? pathMatchMode)
        {
            Action = action;
            Authorization = authorization;
            BucketName = bucketName;
            PathMatchMode = pathMatchMode;
        }
    }
}
