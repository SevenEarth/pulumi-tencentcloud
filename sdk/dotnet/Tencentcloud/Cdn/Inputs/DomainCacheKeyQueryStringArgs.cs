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

    public sealed class DomainCacheKeyQueryStringArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify key rule QS action, values: `includeCustom`, `excludeCustom`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Whether to sort again, values `on`, `off` (Default).
        /// </summary>
        [Input("reorder")]
        public Input<string>? Reorder { get; set; }

        /// <summary>
        /// Whether to use QueryString as part of CacheKey, values `on`, `off` (Default).
        /// </summary>
        [Input("switch")]
        public Input<string>? Switch { get; set; }

        /// <summary>
        /// Array of included/excluded query strings (separated by `;`).
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public DomainCacheKeyQueryStringArgs()
        {
        }
    }
}
