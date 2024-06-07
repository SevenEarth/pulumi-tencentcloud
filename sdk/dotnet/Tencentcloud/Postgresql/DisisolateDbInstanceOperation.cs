// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Postgresql
{
    /// <summary>
    /// Provides a resource to create a postgresql disisolate_db_instance_operation
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var disisolateDbInstanceOperation = new Tencentcloud.Postgresql.DisisolateDbInstanceOperation("disisolateDbInstanceOperation", new()
    ///     {
    ///         DbInstanceIdSets = new[]
    ///         {
    ///             local.Pgsql_id,
    ///         },
    ///         Period = 1,
    ///         AutoVoucher = false,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Postgresql/disisolateDbInstanceOperation:DisisolateDbInstanceOperation")]
    public partial class DisisolateDbInstanceOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to use vouchers. Valid values:true (yes), false (no). Default value:false.
        /// </summary>
        [Output("autoVoucher")]
        public Output<bool?> AutoVoucher { get; private set; } = null!;

        /// <summary>
        /// List of resource IDs. Note that currently you cannot remove multiple instances from isolation at the same time. Only one instance ID can be passed in here.
        /// </summary>
        [Output("dbInstanceIdSets")]
        public Output<ImmutableArray<string>> DbInstanceIdSets { get; private set; } = null!;

        /// <summary>
        /// The valid period (in months) of the monthly-subscribed instance when removing it from isolation.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// Voucher ID list.
        /// </summary>
        [Output("voucherIds")]
        public Output<ImmutableArray<string>> VoucherIds { get; private set; } = null!;


        /// <summary>
        /// Create a DisisolateDbInstanceOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DisisolateDbInstanceOperation(string name, DisisolateDbInstanceOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/disisolateDbInstanceOperation:DisisolateDbInstanceOperation", name, args ?? new DisisolateDbInstanceOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DisisolateDbInstanceOperation(string name, Input<string> id, DisisolateDbInstanceOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/disisolateDbInstanceOperation:DisisolateDbInstanceOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DisisolateDbInstanceOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DisisolateDbInstanceOperation Get(string name, Input<string> id, DisisolateDbInstanceOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new DisisolateDbInstanceOperation(name, id, state, options);
        }
    }

    public sealed class DisisolateDbInstanceOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to use vouchers. Valid values:true (yes), false (no). Default value:false.
        /// </summary>
        [Input("autoVoucher")]
        public Input<bool>? AutoVoucher { get; set; }

        [Input("dbInstanceIdSets", required: true)]
        private InputList<string>? _dbInstanceIdSets;

        /// <summary>
        /// List of resource IDs. Note that currently you cannot remove multiple instances from isolation at the same time. Only one instance ID can be passed in here.
        /// </summary>
        public InputList<string> DbInstanceIdSets
        {
            get => _dbInstanceIdSets ?? (_dbInstanceIdSets = new InputList<string>());
            set => _dbInstanceIdSets = value;
        }

        /// <summary>
        /// The valid period (in months) of the monthly-subscribed instance when removing it from isolation.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("voucherIds")]
        private InputList<string>? _voucherIds;

        /// <summary>
        /// Voucher ID list.
        /// </summary>
        public InputList<string> VoucherIds
        {
            get => _voucherIds ?? (_voucherIds = new InputList<string>());
            set => _voucherIds = value;
        }

        public DisisolateDbInstanceOperationArgs()
        {
        }
        public static new DisisolateDbInstanceOperationArgs Empty => new DisisolateDbInstanceOperationArgs();
    }

    public sealed class DisisolateDbInstanceOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to use vouchers. Valid values:true (yes), false (no). Default value:false.
        /// </summary>
        [Input("autoVoucher")]
        public Input<bool>? AutoVoucher { get; set; }

        [Input("dbInstanceIdSets")]
        private InputList<string>? _dbInstanceIdSets;

        /// <summary>
        /// List of resource IDs. Note that currently you cannot remove multiple instances from isolation at the same time. Only one instance ID can be passed in here.
        /// </summary>
        public InputList<string> DbInstanceIdSets
        {
            get => _dbInstanceIdSets ?? (_dbInstanceIdSets = new InputList<string>());
            set => _dbInstanceIdSets = value;
        }

        /// <summary>
        /// The valid period (in months) of the monthly-subscribed instance when removing it from isolation.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("voucherIds")]
        private InputList<string>? _voucherIds;

        /// <summary>
        /// Voucher ID list.
        /// </summary>
        public InputList<string> VoucherIds
        {
            get => _voucherIds ?? (_voucherIds = new InputList<string>());
            set => _voucherIds = value;
        }

        public DisisolateDbInstanceOperationState()
        {
        }
        public static new DisisolateDbInstanceOperationState Empty => new DisisolateDbInstanceOperationState();
    }
}
