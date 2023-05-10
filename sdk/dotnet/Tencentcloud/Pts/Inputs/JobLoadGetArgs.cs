// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Pts.Inputs
{

    public sealed class JobLoadGetArgs : Pulumi.ResourceArgs
    {
        [Input("geoRegionsLoadDistributions")]
        private InputList<Inputs.JobLoadGeoRegionsLoadDistributionGetArgs>? _geoRegionsLoadDistributions;
        public InputList<Inputs.JobLoadGeoRegionsLoadDistributionGetArgs> GeoRegionsLoadDistributions
        {
            get => _geoRegionsLoadDistributions ?? (_geoRegionsLoadDistributions = new InputList<Inputs.JobLoadGeoRegionsLoadDistributionGetArgs>());
            set => _geoRegionsLoadDistributions = value;
        }

        [Input("loadSpec")]
        public Input<Inputs.JobLoadLoadSpecGetArgs>? LoadSpec { get; set; }

        [Input("vpcLoadDistribution")]
        public Input<Inputs.JobLoadVpcLoadDistributionGetArgs>? VpcLoadDistribution { get; set; }

        public JobLoadGetArgs()
        {
        }
    }
}
