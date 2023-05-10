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
    public sealed class ZoneSettingForceRedirect
    {
        /// <summary>
        /// Redirection status code.- 301- 302 Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly int? RedirectStatusCode;
        /// <summary>
        /// Whether to enable force redirect.- `on`: Enable.- `off`: Disable.
        /// </summary>
        public readonly string Switch;

        [OutputConstructor]
        private ZoneSettingForceRedirect(
            int? redirectStatusCode,

            string @switch)
        {
            RedirectStatusCode = redirectStatusCode;
            Switch = @switch;
        }
    }
}
