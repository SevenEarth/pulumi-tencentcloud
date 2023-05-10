// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cdn.Outputs
{

    [OutputType]
    public sealed class DomainCacheKeyQueryString
    {
        /// <summary>
        /// Specify key rule QS action, values: `includeCustom`, `excludeCustom`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Whether to sort again, values `on`, `off` (Default).
        /// </summary>
        public readonly string? Reorder;
        /// <summary>
        /// Whether to use QueryString as part of CacheKey, values `on`, `off` (Default).
        /// </summary>
        public readonly string? Switch;
        /// <summary>
        /// Array of included/excluded query strings (separated by `;`).
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private DomainCacheKeyQueryString(
            string? action,

            string? reorder,

            string? @switch,

            string? value)
        {
            Action = action;
            Reorder = reorder;
            Switch = @switch;
            Value = value;
        }
    }
}
