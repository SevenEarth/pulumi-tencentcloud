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

    public sealed class ZoneSettingUpstreamHttp2Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to enable cache prefresh.- `on`: Enable.- `off`: Disable.
        /// </summary>
        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        public ZoneSettingUpstreamHttp2Args()
        {
        }
    }
}
