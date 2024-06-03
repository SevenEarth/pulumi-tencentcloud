// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Outputs
{

    [OutputType]
    public sealed class ZoneSettingHttpsHsts
    {
        /// <summary>
        /// Specifies whether to include subdomain names. Valid values: `on` and `off`. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? IncludeSubDomains;
        /// <summary>
        /// MaxAge value in seconds, should be no more than 1 day. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly int? MaxAge;
        /// <summary>
        /// Specifies whether to preload. Valid values: `on` and `off`. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? Preload;
        /// <summary>
        /// `on`: Enable.
        /// </summary>
        public readonly string Switch;

        [OutputConstructor]
        private ZoneSettingHttpsHsts(
            string? includeSubDomains,

            int? maxAge,

            string? preload,

            string @switch)
        {
            IncludeSubDomains = includeSubDomains;
            MaxAge = maxAge;
            Preload = preload;
            Switch = @switch;
        }
    }
}
