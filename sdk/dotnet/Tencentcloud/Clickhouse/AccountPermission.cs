// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Clickhouse
{
    /// <summary>
    /// Provides a resource to create a clickhouse account_permission
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
    ///     var accountPermissionAllDatabase = new Tencentcloud.Clickhouse.AccountPermission("accountPermissionAllDatabase", new()
    ///     {
    ///         AllDatabase = true,
    ///         Cluster = "default_cluster",
    ///         GlobalPrivileges = new[]
    ///         {
    ///             "SELECT",
    ///             "ALTER",
    ///         },
    ///         InstanceId = "cdwch-xxxxxx",
    ///         UserName = "user1",
    ///     });
    /// 
    ///     var accountPermissionNotAllDatabase = new Tencentcloud.Clickhouse.AccountPermission("accountPermissionNotAllDatabase", new()
    ///     {
    ///         AllDatabase = false,
    ///         Cluster = "default_cluster",
    ///         DatabasePrivilegeLists = new[]
    ///         {
    ///             new Tencentcloud.Clickhouse.Inputs.AccountPermissionDatabasePrivilegeListArgs
    ///             {
    ///                 DatabaseName = "xxxxxx",
    ///                 DatabasePrivileges = new[]
    ///                 {
    ///                     "SELECT",
    ///                     "ALTER",
    ///                 },
    ///             },
    ///         },
    ///         InstanceId = "cdwch-xxxxxx",
    ///         UserName = "user2",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// clickhouse account_permission can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Clickhouse/accountPermission:AccountPermission account_permission ${instanceId}#${cluster}#${userName}
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Clickhouse/accountPermission:AccountPermission")]
    public partial class AccountPermission : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether all database tables.
        /// </summary>
        [Output("allDatabase")]
        public Output<bool> AllDatabase { get; private set; } = null!;

        /// <summary>
        /// Cluster name.
        /// </summary>
        [Output("cluster")]
        public Output<string> Cluster { get; private set; } = null!;

        /// <summary>
        /// Database privilege list.
        /// </summary>
        [Output("databasePrivilegeLists")]
        public Output<ImmutableArray<Outputs.AccountPermissionDatabasePrivilegeList>> DatabasePrivilegeLists { get; private set; } = null!;

        /// <summary>
        /// Global privileges.
        /// </summary>
        [Output("globalPrivileges")]
        public Output<ImmutableArray<string>> GlobalPrivileges { get; private set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// User name.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a AccountPermission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountPermission(string name, AccountPermissionArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Clickhouse/accountPermission:AccountPermission", name, args ?? new AccountPermissionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountPermission(string name, Input<string> id, AccountPermissionState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Clickhouse/accountPermission:AccountPermission", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccountPermission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountPermission Get(string name, Input<string> id, AccountPermissionState? state = null, CustomResourceOptions? options = null)
        {
            return new AccountPermission(name, id, state, options);
        }
    }

    public sealed class AccountPermissionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether all database tables.
        /// </summary>
        [Input("allDatabase", required: true)]
        public Input<bool> AllDatabase { get; set; } = null!;

        /// <summary>
        /// Cluster name.
        /// </summary>
        [Input("cluster", required: true)]
        public Input<string> Cluster { get; set; } = null!;

        [Input("databasePrivilegeLists")]
        private InputList<Inputs.AccountPermissionDatabasePrivilegeListArgs>? _databasePrivilegeLists;

        /// <summary>
        /// Database privilege list.
        /// </summary>
        public InputList<Inputs.AccountPermissionDatabasePrivilegeListArgs> DatabasePrivilegeLists
        {
            get => _databasePrivilegeLists ?? (_databasePrivilegeLists = new InputList<Inputs.AccountPermissionDatabasePrivilegeListArgs>());
            set => _databasePrivilegeLists = value;
        }

        [Input("globalPrivileges")]
        private InputList<string>? _globalPrivileges;

        /// <summary>
        /// Global privileges.
        /// </summary>
        public InputList<string> GlobalPrivileges
        {
            get => _globalPrivileges ?? (_globalPrivileges = new InputList<string>());
            set => _globalPrivileges = value;
        }

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// User name.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public AccountPermissionArgs()
        {
        }
        public static new AccountPermissionArgs Empty => new AccountPermissionArgs();
    }

    public sealed class AccountPermissionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether all database tables.
        /// </summary>
        [Input("allDatabase")]
        public Input<bool>? AllDatabase { get; set; }

        /// <summary>
        /// Cluster name.
        /// </summary>
        [Input("cluster")]
        public Input<string>? Cluster { get; set; }

        [Input("databasePrivilegeLists")]
        private InputList<Inputs.AccountPermissionDatabasePrivilegeListGetArgs>? _databasePrivilegeLists;

        /// <summary>
        /// Database privilege list.
        /// </summary>
        public InputList<Inputs.AccountPermissionDatabasePrivilegeListGetArgs> DatabasePrivilegeLists
        {
            get => _databasePrivilegeLists ?? (_databasePrivilegeLists = new InputList<Inputs.AccountPermissionDatabasePrivilegeListGetArgs>());
            set => _databasePrivilegeLists = value;
        }

        [Input("globalPrivileges")]
        private InputList<string>? _globalPrivileges;

        /// <summary>
        /// Global privileges.
        /// </summary>
        public InputList<string> GlobalPrivileges
        {
            get => _globalPrivileges ?? (_globalPrivileges = new InputList<string>());
            set => _globalPrivileges = value;
        }

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// User name.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public AccountPermissionState()
        {
        }
        public static new AccountPermissionState Empty => new AccountPermissionState();
    }
}
