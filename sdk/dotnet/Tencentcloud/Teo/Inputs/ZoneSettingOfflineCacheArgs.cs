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

    public sealed class ZoneSettingOfflineCacheArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable offline cache.
        /// </summary>
        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        public ZoneSettingOfflineCacheArgs()
        {
        }
        public static new ZoneSettingOfflineCacheArgs Empty => new ZoneSettingOfflineCacheArgs();
    }
}
