// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.As.Inputs
{

    public sealed class ScalingConfigInstanceNameSettingsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CVM instance name.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// Type of CVM instance name. Valid values: `ORIGINAL` and `UNIQUE`. Default is `ORIGINAL`.
        /// </summary>
        [Input("instanceNameStyle")]
        public Input<string>? InstanceNameStyle { get; set; }

        public ScalingConfigInstanceNameSettingsGetArgs()
        {
        }
    }
}
