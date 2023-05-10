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

    public sealed class MigrateJobMigrateOptionDatabaseTableArgs : Pulumi.ResourceArgs
    {
        [Input("advancedObjects")]
        private InputList<string>? _advancedObjects;

        /// <summary>
        /// AdvancedObjects.
        /// </summary>
        public InputList<string> AdvancedObjects
        {
            get => _advancedObjects ?? (_advancedObjects = new InputList<string>());
            set => _advancedObjects = value;
        }

        [Input("databases")]
        private InputList<Inputs.MigrateJobMigrateOptionDatabaseTableDatabaseArgs>? _databases;

        /// <summary>
        /// The database list.
        /// </summary>
        public InputList<Inputs.MigrateJobMigrateOptionDatabaseTableDatabaseArgs> Databases
        {
            get => _databases ?? (_databases = new InputList<Inputs.MigrateJobMigrateOptionDatabaseTableDatabaseArgs>());
            set => _databases = value;
        }

        /// <summary>
        /// Object mode. eg:all,partial.
        /// </summary>
        [Input("objectMode", required: true)]
        public Input<string> ObjectMode { get; set; } = null!;

        public MigrateJobMigrateOptionDatabaseTableArgs()
        {
        }
    }
}
