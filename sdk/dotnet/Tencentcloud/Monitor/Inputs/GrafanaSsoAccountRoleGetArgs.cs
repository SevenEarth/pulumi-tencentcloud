// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor.Inputs
{

    public sealed class GrafanaSsoAccountRoleGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Grafana organization id string.
        /// </summary>
        [Input("organization", required: true)]
        public Input<string> Organization { get; set; } = null!;

        /// <summary>
        /// Grafana role, one of {Admin,Editor,Viewer}.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public GrafanaSsoAccountRoleGetArgs()
        {
        }
    }
}
