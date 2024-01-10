// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Eb.Outputs
{

    [OutputType]
    public sealed class GetPlatformProductsPlatformProductResult
    {
        /// <summary>
        /// Platform product name.
        /// </summary>
        public readonly string ProductName;
        /// <summary>
        /// Platform product type.
        /// </summary>
        public readonly string ProductType;

        [OutputConstructor]
        private GetPlatformProductsPlatformProductResult(
            string productName,

            string productType)
        {
            ProductName = productName;
            ProductType = productType;
        }
    }
}
