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

    public sealed class MigrateJobMigrateOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Consistency.
        /// </summary>
        [Input("consistency")]
        public Input<Inputs.MigrateJobMigrateOptionConsistencyArgs>? Consistency { get; set; }

        /// <summary>
        /// Migration object option, you need to tell the migration service which library table objects to migrate.
        /// </summary>
        [Input("databaseTable", required: true)]
        public Input<Inputs.MigrateJobMigrateOptionDatabaseTableArgs> DatabaseTable { get; set; } = null!;

        [Input("extraAttrs")]
        private InputList<Inputs.MigrateJobMigrateOptionExtraAttrArgs>? _extraAttrs;

        /// <summary>
        /// ExtraAttr.
        /// </summary>
        public InputList<Inputs.MigrateJobMigrateOptionExtraAttrArgs> ExtraAttrs
        {
            get => _extraAttrs ?? (_extraAttrs = new InputList<Inputs.MigrateJobMigrateOptionExtraAttrArgs>());
            set => _extraAttrs = value;
        }

        /// <summary>
        /// IsDstReadOnly.
        /// </summary>
        [Input("isDstReadOnly")]
        public Input<bool>? IsDstReadOnly { get; set; }

        /// <summary>
        /// IsMigrateAccount.
        /// </summary>
        [Input("isMigrateAccount")]
        public Input<bool>? IsMigrateAccount { get; set; }

        /// <summary>
        /// IsOverrideRoot.
        /// </summary>
        [Input("isOverrideRoot")]
        public Input<bool>? IsOverrideRoot { get; set; }

        /// <summary>
        /// MigrateType.
        /// </summary>
        [Input("migrateType")]
        public Input<string>? MigrateType { get; set; }

        public MigrateJobMigrateOptionArgs()
        {
        }
        public static new MigrateJobMigrateOptionArgs Empty => new MigrateJobMigrateOptionArgs();
    }
}
