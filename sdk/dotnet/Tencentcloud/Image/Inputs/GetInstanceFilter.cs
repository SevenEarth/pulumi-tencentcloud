// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Image.Inputs
{

    public sealed class GetInstanceFilterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Key of the filter, valid keys: `image-id`, `image-type`, `image-name`.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;

        /// <summary>
        /// Values of the filter.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetInstanceFilterArgs()
        {
        }
        public static new GetInstanceFilterArgs Empty => new GetInstanceFilterArgs();
    }
}
