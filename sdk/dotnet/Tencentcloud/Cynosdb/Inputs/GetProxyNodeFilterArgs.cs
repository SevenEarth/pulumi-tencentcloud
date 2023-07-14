// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb.Inputs
{

    public sealed class GetProxyNodeFilterInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Exact match or not.
        /// </summary>
        [Input("exactMatch")]
        public Input<bool>? ExactMatch { get; set; }

        /// <summary>
        /// Search Fields. Supported: Status, ProxyNodeId, ClusterId.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("names", required: true)]
        private InputList<string>? _names;

        /// <summary>
        /// Search String.
        /// </summary>
        public InputList<string> Names
        {
            get => _names ?? (_names = new InputList<string>());
            set => _names = value;
        }

        /// <summary>
        /// Operator.
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// Search String.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public GetProxyNodeFilterInputArgs()
        {
        }
    }
}
