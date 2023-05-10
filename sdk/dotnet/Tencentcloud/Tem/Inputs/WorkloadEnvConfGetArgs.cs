// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tem.Inputs
{

    public sealed class WorkloadEnvConfGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// referenced config name when type=referenced.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// env key.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// referenced secret name when type=referenced.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        /// <summary>
        /// env type, support default, referenced.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// env value.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public WorkloadEnvConfGetArgs()
        {
        }
    }
}
