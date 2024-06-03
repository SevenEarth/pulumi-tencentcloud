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

    public sealed class ZoneSettingCacheKeyQueryStringArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// `includeCustom`: Include the specified query strings.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Whether to use QueryString as part of CacheKey.
        /// </summary>
        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// Array of query strings used/excluded. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public ZoneSettingCacheKeyQueryStringArgs()
        {
        }
        public static new ZoneSettingCacheKeyQueryStringArgs Empty => new ZoneSettingCacheKeyQueryStringArgs();
    }
}
