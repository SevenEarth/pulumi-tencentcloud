// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dbbrain
{
    /// <summary>
    /// Provides a resource to create a dbbrain modify_diag_db_instance_conf
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @on = new Tencentcloud.Dbbrain.ModifyDiagDbInstanceOperation("on", new Tencentcloud.Dbbrain.ModifyDiagDbInstanceOperationArgs
    ///         {
    ///             InstanceConfs = new Tencentcloud.Dbbrain.Inputs.ModifyDiagDbInstanceOperationInstanceConfsArgs
    ///             {
    ///                 DailyInspection = "Yes",
    ///                 OverviewDisplay = "Yes",
    ///             },
    ///             InstanceIds = 
    ///             {
    ///                 "%s",
    ///             },
    ///             Product = "mysql",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var off = new Tencentcloud.Dbbrain.ModifyDiagDbInstanceOperation("off", new Tencentcloud.Dbbrain.ModifyDiagDbInstanceOperationArgs
    ///         {
    ///             InstanceConfs = new Tencentcloud.Dbbrain.Inputs.ModifyDiagDbInstanceOperationInstanceConfsArgs
    ///             {
    ///                 DailyInspection = "No",
    ///                 OverviewDisplay = "No",
    ///             },
    ///             InstanceIds = 
    ///             {
    ///                 "%s",
    ///             },
    ///             Product = "mysql",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dbbrain/modifyDiagDbInstanceOperation:ModifyDiagDbInstanceOperation")]
    public partial class ModifyDiagDbInstanceOperation : Pulumi.CustomResource
    {
        /// <summary>
        /// Instance configuration, including inspection, overview switch, etc.
        /// </summary>
        [Output("instanceConfs")]
        public Output<Outputs.ModifyDiagDbInstanceOperationInstanceConfs> InstanceConfs { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of the instance whose inspection status is changed.
        /// </summary>
        [Output("instanceIds")]
        public Output<ImmutableArray<string>> InstanceIds { get; private set; } = null!;

        /// <summary>
        /// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL.
        /// </summary>
        [Output("product")]
        public Output<string> Product { get; private set; } = null!;

        /// <summary>
        /// Effective instance region, the value is All, which means all regions.
        /// </summary>
        [Output("regions")]
        public Output<string?> Regions { get; private set; } = null!;


        /// <summary>
        /// Create a ModifyDiagDbInstanceOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ModifyDiagDbInstanceOperation(string name, ModifyDiagDbInstanceOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dbbrain/modifyDiagDbInstanceOperation:ModifyDiagDbInstanceOperation", name, args ?? new ModifyDiagDbInstanceOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ModifyDiagDbInstanceOperation(string name, Input<string> id, ModifyDiagDbInstanceOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dbbrain/modifyDiagDbInstanceOperation:ModifyDiagDbInstanceOperation", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ModifyDiagDbInstanceOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ModifyDiagDbInstanceOperation Get(string name, Input<string> id, ModifyDiagDbInstanceOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new ModifyDiagDbInstanceOperation(name, id, state, options);
        }
    }

    public sealed class ModifyDiagDbInstanceOperationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance configuration, including inspection, overview switch, etc.
        /// </summary>
        [Input("instanceConfs", required: true)]
        public Input<Inputs.ModifyDiagDbInstanceOperationInstanceConfsArgs> InstanceConfs { get; set; } = null!;

        [Input("instanceIds")]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// Specifies the ID of the instance whose inspection status is changed.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL.
        /// </summary>
        [Input("product", required: true)]
        public Input<string> Product { get; set; } = null!;

        /// <summary>
        /// Effective instance region, the value is All, which means all regions.
        /// </summary>
        [Input("regions")]
        public Input<string>? Regions { get; set; }

        public ModifyDiagDbInstanceOperationArgs()
        {
        }
    }

    public sealed class ModifyDiagDbInstanceOperationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance configuration, including inspection, overview switch, etc.
        /// </summary>
        [Input("instanceConfs")]
        public Input<Inputs.ModifyDiagDbInstanceOperationInstanceConfsGetArgs>? InstanceConfs { get; set; }

        [Input("instanceIds")]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// Specifies the ID of the instance whose inspection status is changed.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL.
        /// </summary>
        [Input("product")]
        public Input<string>? Product { get; set; }

        /// <summary>
        /// Effective instance region, the value is All, which means all regions.
        /// </summary>
        [Input("regions")]
        public Input<string>? Regions { get; set; }

        public ModifyDiagDbInstanceOperationState()
        {
        }
    }
}
