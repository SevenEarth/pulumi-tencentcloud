// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Scf.Outputs
{

    [OutputType]
    public sealed class GetLayerVersionsLayerVersionResult
    {
        /// <summary>
        /// Creation time.
        /// </summary>
        public readonly string AddTime;
        /// <summary>
        /// Runtime applicable to a versionNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<string> CompatibleRuntimes;
        /// <summary>
        /// Version descriptionNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Layer name.
        /// </summary>
        public readonly string LayerName;
        /// <summary>
        /// Version number.
        /// </summary>
        public readonly int LayerVersion;
        /// <summary>
        /// License informationNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string LicenseInfo;
        /// <summary>
        /// StampNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string Stamp;
        /// <summary>
        /// Current status of specific layer version. For valid values, please see [here](https://intl.cloud.tencent.com/document/product/583/47175?from_cn_redirect=1#.E5.B1.82.EF.BC.88layer.EF.BC.89.E7.8A.B6.E6.80.81).
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetLayerVersionsLayerVersionResult(
            string addTime,

            ImmutableArray<string> compatibleRuntimes,

            string description,

            string layerName,

            int layerVersion,

            string licenseInfo,

            string stamp,

            string status)
        {
            AddTime = addTime;
            CompatibleRuntimes = compatibleRuntimes;
            Description = description;
            LayerName = layerName;
            LayerVersion = layerVersion;
            LicenseInfo = licenseInfo;
            Stamp = stamp;
            Status = status;
        }
    }
}
