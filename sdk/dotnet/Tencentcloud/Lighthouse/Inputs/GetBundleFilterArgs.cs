// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Lighthouse.Inputs
{

    public sealed class GetBundleFilterInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Field to be filtered.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// Filter value of field.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public GetBundleFilterInputArgs()
        {
        }
        public static new GetBundleFilterInputArgs Empty => new GetBundleFilterInputArgs();
    }
}
