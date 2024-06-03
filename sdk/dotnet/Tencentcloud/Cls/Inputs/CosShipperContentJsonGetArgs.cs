// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cls.Inputs
{

    public sealed class CosShipperContentJsonGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enablement flag.
        /// </summary>
        [Input("enableTag", required: true)]
        public Input<bool> EnableTag { get; set; } = null!;

        [Input("metaFields", required: true)]
        private InputList<string>? _metaFields;

        /// <summary>
        /// Metadata information list
        /// Note: this field may return null, indicating that no valid values can be obtained..
        /// </summary>
        public InputList<string> MetaFields
        {
            get => _metaFields ?? (_metaFields = new InputList<string>());
            set => _metaFields = value;
        }

        public CosShipperContentJsonGetArgs()
        {
        }
        public static new CosShipperContentJsonGetArgs Empty => new CosShipperContentJsonGetArgs();
    }
}
