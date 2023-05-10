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

    public sealed class ZoneSettingHttpsHstsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to include subdomain names. Valid values: `on` and `off`. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("includeSubDomains")]
        public Input<string>? IncludeSubDomains { get; set; }

        /// <summary>
        /// MaxAge value in seconds, should be no more than 1 day. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("maxAge")]
        public Input<int>? MaxAge { get; set; }

        /// <summary>
        /// Specifies whether to preload. Valid values: `on` and `off`. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("preload")]
        public Input<string>? Preload { get; set; }

        /// <summary>
        /// - `on`: Enable.- `off`: Disable.
        /// </summary>
        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        public ZoneSettingHttpsHstsArgs()
        {
        }
    }
}
