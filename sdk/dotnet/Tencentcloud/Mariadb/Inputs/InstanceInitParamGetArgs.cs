// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb.Inputs
{

    public sealed class InstanceInitParamGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// parameter name.
        /// </summary>
        [Input("param", required: true)]
        public Input<string> Param { get; set; } = null!;

        /// <summary>
        /// parameter value.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public InstanceInitParamGetArgs()
        {
        }
    }
}
