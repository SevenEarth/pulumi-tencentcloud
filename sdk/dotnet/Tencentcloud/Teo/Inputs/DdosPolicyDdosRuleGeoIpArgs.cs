// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Inputs
{

    public sealed class DdosPolicyDdosRuleGeoIpArgs : Pulumi.ResourceArgs
    {
        [Input("regionIds")]
        private InputList<int>? _regionIds;
        public InputList<int> RegionIds
        {
            get => _regionIds ?? (_regionIds = new InputList<int>());
            set => _regionIds = value;
        }

        [Input("switch")]
        public Input<string>? Switch { get; set; }

        public DdosPolicyDdosRuleGeoIpArgs()
        {
        }
    }
}
