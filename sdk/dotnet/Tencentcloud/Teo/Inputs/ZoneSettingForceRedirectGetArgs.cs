// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Inputs
{

    public sealed class ZoneSettingForceRedirectGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Redirection status code.- 301- 302 Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("redirectStatusCode")]
        public Input<int>? RedirectStatusCode { get; set; }

        /// <summary>
        /// Whether to enable force redirect.- `on`: Enable.- `off`: Disable.
        /// </summary>
        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        public ZoneSettingForceRedirectGetArgs()
        {
        }
    }
}
