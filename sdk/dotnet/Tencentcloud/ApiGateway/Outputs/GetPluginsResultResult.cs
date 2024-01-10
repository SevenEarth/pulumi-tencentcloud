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
    public sealed class GetPluginsResultResult
    {
        /// <summary>
        /// API ID.
        /// </summary>
        public readonly string ApiId;
        /// <summary>
        /// API name.
        /// </summary>
        public readonly string ApiName;
        /// <summary>
        /// API type.
        /// </summary>
        public readonly string ApiType;
        /// <summary>
        /// Whether the API is bound to other plugins.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly bool AttachedOtherPlugin;
        /// <summary>
        /// Whether the API is bound to the current plugin.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly bool IsAttached;
        /// <summary>
        /// API method.
        /// </summary>
        public readonly string Method;
        /// <summary>
        /// API path.
        /// </summary>
        public readonly string Path;

        [OutputConstructor]
        private GetPluginsResultResult(
            string apiId,

            string apiName,

            string apiType,

            bool attachedOtherPlugin,

            bool isAttached,

            string method,

            string path)
        {
            ApiId = apiId;
            ApiName = apiName;
            ApiType = apiType;
            AttachedOtherPlugin = attachedOtherPlugin;
            IsAttached = isAttached;
            Method = method;
            Path = path;
        }
    }
}
