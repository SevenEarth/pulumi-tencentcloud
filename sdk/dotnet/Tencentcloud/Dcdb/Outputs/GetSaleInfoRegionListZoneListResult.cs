// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb.Outputs
{

    [OutputType]
    public sealed class GetSaleInfoRegionListZoneListResult
    {
        /// <summary>
        /// is zone on sale.
        /// </summary>
        public readonly bool OnSale;
        /// <summary>
        /// zone name(en).
        /// </summary>
        public readonly string Zone;
        /// <summary>
        /// zone id.
        /// </summary>
        public readonly int ZoneId;
        /// <summary>
        /// zone name(zh).
        /// </summary>
        public readonly string ZoneName;

        [OutputConstructor]
        private GetSaleInfoRegionListZoneListResult(
            bool onSale,

            string zone,

            int zoneId,

            string zoneName)
        {
            OnSale = onSale;
            Zone = zone;
            ZoneId = zoneId;
            ZoneName = zoneName;
        }
    }
}
