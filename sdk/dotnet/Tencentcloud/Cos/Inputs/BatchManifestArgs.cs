// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cos.Inputs
{

    public sealed class BatchManifestArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The location information of the list of objects.
        /// </summary>
        [Input("location", required: true)]
        public Input<Inputs.BatchManifestLocationArgs> Location { get; set; } = null!;

        /// <summary>
        /// Format information that describes the list of objects. If it is a CSV file, this element describes the fields contained in the manifest.
        /// </summary>
        [Input("spec", required: true)]
        public Input<Inputs.BatchManifestSpecArgs> Spec { get; set; } = null!;

        public BatchManifestArgs()
        {
        }
    }
}
