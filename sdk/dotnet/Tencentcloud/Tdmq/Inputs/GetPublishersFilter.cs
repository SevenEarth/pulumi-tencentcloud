// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tdmq.Inputs
{

    public sealed class GetPublishersFilterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the filter parameter.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("values")]
        private List<string>? _values;

        /// <summary>
        /// value.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetPublishersFilterArgs()
        {
        }
        public static new GetPublishersFilterArgs Empty => new GetPublishersFilterArgs();
    }
}
