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

    public sealed class BatchManifestSpecGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("fields")]
        private InputList<string>? _fields;

        /// <summary>
        /// Describes the fields contained in the listing, which you need to use to specify CSV file fields when Format is COSBatchOperations_CSV_V1. Legal fields are: Ignore, Bucket, Key, VersionId.
        /// </summary>
        public InputList<string> Fields
        {
            get => _fields ?? (_fields = new InputList<string>());
            set => _fields = value;
        }

        /// <summary>
        /// Specifies the format information for the list of objects. Legal fields are: COSBatchOperations_CSV_V1, COSInventoryReport_CSV_V1.
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        public BatchManifestSpecGetArgs()
        {
        }
        public static new BatchManifestSpecGetArgs Empty => new BatchManifestSpecGetArgs();
    }
}
