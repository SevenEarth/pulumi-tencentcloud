// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts.Inputs
{

    public sealed class SyncConfigObjectsArgs : Pulumi.ResourceArgs
    {
        [Input("advancedObjects")]
        private InputList<string>? _advancedObjects;

        /// <summary>
        /// For advanced object types, such as function and procedure, when an advanced object needs to be synchronized, the initialization type must include the structure initialization type, that is, the value of the Options.InitType field is Structure or Full. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public InputList<string> AdvancedObjects
        {
            get => _advancedObjects ?? (_advancedObjects = new InputList<string>());
            set => _advancedObjects = value;
        }

        [Input("databases")]
        private InputList<Inputs.SyncConfigObjectsDatabaseArgs>? _databases;

        /// <summary>
        /// Synchronization object, not null when Mode is Partial. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public InputList<Inputs.SyncConfigObjectsDatabaseArgs> Databases
        {
            get => _databases ?? (_databases = new InputList<Inputs.SyncConfigObjectsDatabaseArgs>());
            set => _databases = value;
        }

        /// <summary>
        /// Migration object type Partial (partial object). Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// OnlineDDL type. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("onlineDdl")]
        public Input<Inputs.SyncConfigObjectsOnlineDdlArgs>? OnlineDdl { get; set; }

        public SyncConfigObjectsArgs()
        {
        }
    }
}
