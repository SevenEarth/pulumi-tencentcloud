// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts.Outputs
{

    [OutputType]
    public sealed class SyncConfigObjects
    {
        /// <summary>
        /// For advanced object types, such as function and procedure, when an advanced object needs to be synchronized, the initialization type must include the structure initialization type, that is, the value of the Options.InitType field is Structure or Full. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly ImmutableArray<string> AdvancedObjects;
        /// <summary>
        /// Synchronization object, not null when Mode is Partial. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.SyncConfigObjectsDatabase> Databases;
        /// <summary>
        /// Migration object type Partial (partial object). Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// OnlineDDL type. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly Outputs.SyncConfigObjectsOnlineDdl? OnlineDdl;

        [OutputConstructor]
        private SyncConfigObjects(
            ImmutableArray<string> advancedObjects,

            ImmutableArray<Outputs.SyncConfigObjectsDatabase> databases,

            string? mode,

            Outputs.SyncConfigObjectsOnlineDdl? onlineDdl)
        {
            AdvancedObjects = advancedObjects;
            Databases = databases;
            Mode = mode;
            OnlineDdl = onlineDdl;
        }
    }
}