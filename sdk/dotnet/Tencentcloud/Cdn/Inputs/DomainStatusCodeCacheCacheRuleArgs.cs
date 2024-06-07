// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cdn.Inputs
{

    public sealed class DomainStatusCodeCacheCacheRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Status code cache expiration time (in seconds).
        /// </summary>
        [Input("cacheTime", required: true)]
        public Input<int> CacheTime { get; set; } = null!;

        /// <summary>
        /// Code of status cache. available values: `403`, `404`.
        /// </summary>
        [Input("statusCode", required: true)]
        public Input<string> StatusCode { get; set; } = null!;

        public DomainStatusCodeCacheCacheRuleArgs()
        {
        }
        public static new DomainStatusCodeCacheCacheRuleArgs Empty => new DomainStatusCodeCacheCacheRuleArgs();
    }
}
