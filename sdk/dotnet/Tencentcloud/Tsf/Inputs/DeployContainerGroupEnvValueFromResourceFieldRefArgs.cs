// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Inputs
{

    public sealed class DeployContainerGroupEnvValueFromResourceFieldRefArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Resource configuration of Kubernetes. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("resource")]
        public Input<string>? Resource { get; set; }

        public DeployContainerGroupEnvValueFromResourceFieldRefArgs()
        {
        }
        public static new DeployContainerGroupEnvValueFromResourceFieldRefArgs Empty => new DeployContainerGroupEnvValueFromResourceFieldRefArgs();
    }
}
