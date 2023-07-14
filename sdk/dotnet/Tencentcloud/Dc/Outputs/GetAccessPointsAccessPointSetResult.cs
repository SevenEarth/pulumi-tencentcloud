// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dc.Outputs
{

    [OutputType]
    public sealed class GetAccessPointsAccessPointSetResult
    {
        /// <summary>
        /// Unique access point ID.
        /// </summary>
        public readonly string AccessPointId;
        /// <summary>
        /// Access point name.
        /// </summary>
        public readonly string AccessPointName;
        /// <summary>
        /// Access point type. Valid values: `VXLAN`, `QCPL`, and `QCAR`.Note: this field may return `null`, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string AccessPointType;
        /// <summary>
        /// Access point regionNote: this field may return `null`, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string Area;
        /// <summary>
        /// Available port type at the access point. Valid values: 1000BASE-T: gigabit electrical port; 1000BASE-LX: 10 km gigabit single-mode optical port; 1000BASE-ZX: 80 km gigabit single-mode optical port; 10GBASE-LR: 10 km 10-gigabit single-mode optical port; 10GBASE-ZR: 80 km 10-gigabit single-mode optical port; 10GBASE-LH: 40 km 10-gigabit single-mode optical port; 100GBASE-LR4: 10 km 100-gigabit single-mode optical portfiber optic port.Note: this field may return `null`, indicating that no valid value is obtained.
        /// </summary>
        public readonly ImmutableArray<string> AvailablePortTypes;
        /// <summary>
        /// City where the access point is locatedNote: this field may return `null`, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string City;
        /// <summary>
        /// Latitude and longitude of the access pointNote: this field may return `null`, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccessPointsAccessPointSetCoordinateResult> Coordinates;
        /// <summary>
        /// List of ISPs supported by access point.
        /// </summary>
        public readonly ImmutableArray<string> LineOperators;
        /// <summary>
        /// Access point location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Access point region, which can be queried through `DescribeRegions`.You can call `DescribeRegions` to get the region ID.
        /// </summary>
        public readonly string RegionId;
        /// <summary>
        /// Access point status. Valid values: available, unavailable.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetAccessPointsAccessPointSetResult(
            string accessPointId,

            string accessPointName,

            string accessPointType,

            string area,

            ImmutableArray<string> availablePortTypes,

            string city,

            ImmutableArray<Outputs.GetAccessPointsAccessPointSetCoordinateResult> coordinates,

            ImmutableArray<string> lineOperators,

            string location,

            string regionId,

            string state)
        {
            AccessPointId = accessPointId;
            AccessPointName = accessPointName;
            AccessPointType = accessPointType;
            Area = area;
            AvailablePortTypes = availablePortTypes;
            City = city;
            Coordinates = coordinates;
            LineOperators = lineOperators;
            Location = location;
            RegionId = regionId;
            State = state;
        }
    }
}
