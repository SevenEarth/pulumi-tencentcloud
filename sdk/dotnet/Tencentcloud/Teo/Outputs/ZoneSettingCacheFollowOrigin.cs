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
    public sealed class ZoneSettingCacheFollowOrigin
    {
        /// <summary>
        /// Specifies whether to follow the origin server configuration.
        /// </summary>
        public readonly string? Switch;

        [OutputConstructor]
        private ZoneSettingCacheFollowOrigin(string? @switch)
        {
            Switch = @switch;
        }
    }
}
