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

    public sealed class MigrateJobMigrateOptionDatabaseTableDatabaseViewGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// NewViewName.
        /// </summary>
        [Input("newViewName")]
        public Input<string>? NewViewName { get; set; }

        /// <summary>
        /// ViewName.
        /// </summary>
        [Input("viewName")]
        public Input<string>? ViewName { get; set; }

        public MigrateJobMigrateOptionDatabaseTableDatabaseViewGetArgs()
        {
        }
    }
}
