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

    public sealed class DomainMaxAgeMaxAgeRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to follow origin, values: `on`/`off`, if set to `on`, the `max_age_time` will be ignored.
        /// </summary>
        [Input("followOrigin")]
        public Input<string>? FollowOrigin { get; set; }

        [Input("maxAgeContents", required: true)]
        private InputList<string>? _maxAgeContents;

        /// <summary>
        /// List of rule paths for each `max_age_type`: `*` for `all`, file ext like `jpg` for `file`, `/dir/like/` for `directory` and `/path/index.html` for `path`.
        /// </summary>
        public InputList<string> MaxAgeContents
        {
            get => _maxAgeContents ?? (_maxAgeContents = new InputList<string>());
            set => _maxAgeContents = value;
        }

        /// <summary>
        /// Max Age time in seconds, this can set to `0` that stands for no cache.
        /// </summary>
        [Input("maxAgeTime", required: true)]
        public Input<int> MaxAgeTime { get; set; } = null!;

        /// <summary>
        /// The following types are supported: `all`: all documents take effect, `file`: the specified file suffix takes effect, `directory`: the specified path takes effect, `path`: specify the absolute path to take effect, `index`: home page.
        /// </summary>
        [Input("maxAgeType", required: true)]
        public Input<string> MaxAgeType { get; set; } = null!;

        public DomainMaxAgeMaxAgeRuleGetArgs()
        {
        }
        public static new DomainMaxAgeMaxAgeRuleGetArgs Empty => new DomainMaxAgeMaxAgeRuleGetArgs();
    }
}
