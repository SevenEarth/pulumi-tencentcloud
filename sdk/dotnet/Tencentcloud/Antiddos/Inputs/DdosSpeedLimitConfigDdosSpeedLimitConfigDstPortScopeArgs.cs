// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Antiddos.Inputs
{

    public sealed class DdosSpeedLimitConfigDdosSpeedLimitConfigDstPortScopeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Starting port, ranging from 1 to 65535.
        /// </summary>
        [Input("beginPort", required: true)]
        public Input<int> BeginPort { get; set; } = null!;

        /// <summary>
        /// end  port, ranging from 1 to 65535.
        /// </summary>
        [Input("endPort", required: true)]
        public Input<int> EndPort { get; set; } = null!;

        public DdosSpeedLimitConfigDdosSpeedLimitConfigDstPortScopeArgs()
        {
        }
    }
}
