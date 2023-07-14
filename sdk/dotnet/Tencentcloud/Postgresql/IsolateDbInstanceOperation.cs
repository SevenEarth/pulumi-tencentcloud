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
    /// Provides a resource to create a postgresql isolate_db_instance_operation
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
    ///         var isolateDbInstanceOperation = new Tencentcloud.Postgresql.IsolateDbInstanceOperation("isolateDbInstanceOperation", new Tencentcloud.Postgresql.IsolateDbInstanceOperationArgs
    ///         {
    ///             DbInstanceIdSets = 
    ///             {
    ///                 local.Pgsql_id,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Postgresql/isolateDbInstanceOperation:IsolateDbInstanceOperation")]
    public partial class IsolateDbInstanceOperation : Pulumi.CustomResource
    {
        /// <summary>
        /// List of resource IDs. Note that currently you cannot isolate multiple instances at the same time. Only one instance ID can be passed in here.
        /// </summary>
        [Output("dbInstanceIdSets")]
        public Output<ImmutableArray<string>> DbInstanceIdSets { get; private set; } = null!;


        /// <summary>
        /// Create a IsolateDbInstanceOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IsolateDbInstanceOperation(string name, IsolateDbInstanceOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/isolateDbInstanceOperation:IsolateDbInstanceOperation", name, args ?? new IsolateDbInstanceOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IsolateDbInstanceOperation(string name, Input<string> id, IsolateDbInstanceOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/isolateDbInstanceOperation:IsolateDbInstanceOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IsolateDbInstanceOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IsolateDbInstanceOperation Get(string name, Input<string> id, IsolateDbInstanceOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new IsolateDbInstanceOperation(name, id, state, options);
        }
    }

    public sealed class IsolateDbInstanceOperationArgs : Pulumi.ResourceArgs
    {
        [Input("dbInstanceIdSets", required: true)]
        private InputList<string>? _dbInstanceIdSets;

        /// <summary>
        /// List of resource IDs. Note that currently you cannot isolate multiple instances at the same time. Only one instance ID can be passed in here.
        /// </summary>
        public InputList<string> DbInstanceIdSets
        {
            get => _dbInstanceIdSets ?? (_dbInstanceIdSets = new InputList<string>());
            set => _dbInstanceIdSets = value;
        }

        public IsolateDbInstanceOperationArgs()
        {
        }
    }

    public sealed class IsolateDbInstanceOperationState : Pulumi.ResourceArgs
    {
        [Input("dbInstanceIdSets")]
        private InputList<string>? _dbInstanceIdSets;

        /// <summary>
        /// List of resource IDs. Note that currently you cannot isolate multiple instances at the same time. Only one instance ID can be passed in here.
        /// </summary>
        public InputList<string> DbInstanceIdSets
        {
            get => _dbInstanceIdSets ?? (_dbInstanceIdSets = new InputList<string>());
            set => _dbInstanceIdSets = value;
        }

        public IsolateDbInstanceOperationState()
        {
        }
    }
}
