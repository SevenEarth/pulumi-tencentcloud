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

    public sealed class CompareTaskObjectsObjectItemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// database mode.
        /// </summary>
        [Input("dbMode")]
        public Input<string>? DbMode { get; set; }

        /// <summary>
        /// database name.
        /// </summary>
        [Input("dbName")]
        public Input<string>? DbName { get; set; }

        /// <summary>
        /// schema name.
        /// </summary>
        [Input("schemaName")]
        public Input<string>? SchemaName { get; set; }

        /// <summary>
        /// table mode.
        /// </summary>
        [Input("tableMode")]
        public Input<string>? TableMode { get; set; }

        [Input("tables")]
        private InputList<Inputs.CompareTaskObjectsObjectItemTableArgs>? _tables;

        /// <summary>
        /// table list.
        /// </summary>
        public InputList<Inputs.CompareTaskObjectsObjectItemTableArgs> Tables
        {
            get => _tables ?? (_tables = new InputList<Inputs.CompareTaskObjectsObjectItemTableArgs>());
            set => _tables = value;
        }

        /// <summary>
        /// view mode.
        /// </summary>
        [Input("viewMode")]
        public Input<string>? ViewMode { get; set; }

        [Input("views")]
        private InputList<Inputs.CompareTaskObjectsObjectItemViewArgs>? _views;

        /// <summary>
        /// view list.
        /// </summary>
        public InputList<Inputs.CompareTaskObjectsObjectItemViewArgs> Views
        {
            get => _views ?? (_views = new InputList<Inputs.CompareTaskObjectsObjectItemViewArgs>());
            set => _views = value;
        }

        public CompareTaskObjectsObjectItemArgs()
        {
        }
    }
}
