// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm.Inputs
{

    public sealed class GetInstancesModificationFilterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Fields to be filtered.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;

        /// <summary>
        /// Value of the field.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetInstancesModificationFilterArgs()
        {
        }
        public static new GetInstancesModificationFilterArgs Empty => new GetInstancesModificationFilterArgs();
    }
}
