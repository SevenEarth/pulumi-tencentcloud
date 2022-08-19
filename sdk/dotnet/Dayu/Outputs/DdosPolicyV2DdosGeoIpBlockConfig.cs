// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dayu.Outputs
{

    [OutputType]
    public sealed class DdosPolicyV2DdosGeoIpBlockConfig
    {
        /// <summary>
        /// Block action, take the value [`drop`, `trans`].
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// When the RegionType is customized, the AreaList must be filled in, and a maximum of 128 must be filled in.
        /// </summary>
        public readonly ImmutableArray<int> AreaLists;
        /// <summary>
        /// Zone type, value [oversea (overseas),china (domestic),customized (custom region)].
        /// </summary>
        public readonly string RegionType;

        [OutputConstructor]
        private DdosPolicyV2DdosGeoIpBlockConfig(
            string action,

            ImmutableArray<int> areaLists,

            string regionType)
        {
            Action = action;
            AreaLists = areaLists;
            RegionType = regionType;
        }
    }
}
