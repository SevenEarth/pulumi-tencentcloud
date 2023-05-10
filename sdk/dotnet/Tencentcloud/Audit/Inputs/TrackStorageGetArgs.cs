// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Audit.Inputs
{

    public sealed class TrackStorageGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Track Storage name:- when StorageType is `cls`, StorageName is cls topicId- when StorageType is `cos`, StorageName is cos bucket name that does not contain `-APPID`.
        /// </summary>
        [Input("storageName", required: true)]
        public Input<string> StorageName { get; set; } = null!;

        /// <summary>
        /// Storage path prefix.
        /// </summary>
        [Input("storagePrefix", required: true)]
        public Input<string> StoragePrefix { get; set; } = null!;

        /// <summary>
        /// Storage region.
        /// </summary>
        [Input("storageRegion", required: true)]
        public Input<string> StorageRegion { get; set; } = null!;

        /// <summary>
        /// Track Storage type, optional:- `cos`- `cls`.
        /// </summary>
        [Input("storageType", required: true)]
        public Input<string> StorageType { get; set; } = null!;

        public TrackStorageGetArgs()
        {
        }
    }
}
