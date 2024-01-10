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
    public sealed class GetApiAppApiResultServiceConfigResult
    {
        /// <summary>
        /// Load balancing method.
        /// </summary>
        public readonly string Method;
        /// <summary>
        /// API backend service path, such as /path. If ServiceType is HTTP, this parameter is required. The front-end and back-end paths can be different.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Backend type. It takes effect when vpc is enabled. Currently supported types are clb, cvm and upstream.
        /// </summary>
        public readonly string Product;
        /// <summary>
        /// The unique ID of the vpc.
        /// </summary>
        public readonly string UniqVpcId;
        /// <summary>
        /// Only required when binding vpc channel.
        /// </summary>
        public readonly string UpstreamId;
        /// <summary>
        /// API&amp;amp;#39;s backend service url. If ServiceType is HTTP, this parameter must be passed.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetApiAppApiResultServiceConfigResult(
            string method,

            string path,

            string product,

            string uniqVpcId,

            string upstreamId,

            string url)
        {
            Method = method;
            Path = path;
            Product = product;
            UniqVpcId = uniqVpcId;
            UpstreamId = upstreamId;
            Url = url;
        }
    }
}
