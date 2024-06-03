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

    public sealed class ClusterParamItemGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Param expected value to set.
        /// </summary>
        [Input("currentValue", required: true)]
        public Input<string> CurrentValue { get; set; } = null!;

        /// <summary>
        /// Name of param, e.g. `character_set_server`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Param old value, indicates the value which already set, this value is required when modifying current_value.
        /// </summary>
        [Input("oldValue")]
        public Input<string>? OldValue { get; set; }

        public ClusterParamItemGetArgs()
        {
        }
        public static new ClusterParamItemGetArgs Empty => new ClusterParamItemGetArgs();
    }
}
