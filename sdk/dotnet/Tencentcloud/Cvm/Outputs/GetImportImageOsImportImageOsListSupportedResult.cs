// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm.Outputs
{

    [OutputType]
    public sealed class GetImportImageOsImportImageOsListSupportedResult
    {
        /// <summary>
        /// Supported Linux OS Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<string> Linuxes;
        /// <summary>
        /// Supported Windows OS Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<string> Windows;

        [OutputConstructor]
        private GetImportImageOsImportImageOsListSupportedResult(
            ImmutableArray<string> linuxes,

            ImmutableArray<string> windows)
        {
            Linuxes = linuxes;
            Windows = windows;
        }
    }
}
