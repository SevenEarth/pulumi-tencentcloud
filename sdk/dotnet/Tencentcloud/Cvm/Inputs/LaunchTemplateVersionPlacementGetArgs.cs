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

    public sealed class LaunchTemplateVersionPlacementGetArgs : Pulumi.ResourceArgs
    {
        [Input("hostIds")]
        private InputList<string>? _hostIds;

        /// <summary>
        /// ID list of CDHs from which the instance can be created. If you have purchased CDHs and specify this parameter, the instances you purchase will be randomly deployed on the CDHs.
        /// </summary>
        public InputList<string> HostIds
        {
            get => _hostIds ?? (_hostIds = new InputList<string>());
            set => _hostIds = value;
        }

        [Input("hostIps")]
        private InputList<string>? _hostIps;

        /// <summary>
        /// IPs of the hosts to create CVMs.
        /// </summary>
        public InputList<string> HostIps
        {
            get => _hostIps ?? (_hostIps = new InputList<string>());
            set => _hostIps = value;
        }

        /// <summary>
        /// ID of the project to which the instance belongs. This parameter can be obtained from the projectId returned by DescribeProject. If this is left empty, the default project is used.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// ID of the availability zone where the instance resides. You can call the DescribeZones API and obtain the ID in the returned Zone field.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public LaunchTemplateVersionPlacementGetArgs()
        {
        }
    }
}
