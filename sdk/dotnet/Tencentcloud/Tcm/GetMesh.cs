// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcm
{
    public static class GetMesh
    {
        /// <summary>
        /// Use this data source to query detailed information of tcm mesh
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mesh = Tencentcloud.Tcm.GetMesh.Invoke(new()
        ///     {
        ///         MeshClusters = new[]
        ///         {
        ///             "cls-xxxx",
        ///         },
        ///         MeshIds = new[]
        ///         {
        ///             "mesh-xxxxxx",
        ///         },
        ///         MeshNames = new[]
        ///         {
        ///             "KEEP_MASH",
        ///         },
        ///         Tags = new[]
        ///         {
        ///             "key",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetMeshResult> InvokeAsync(GetMeshArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMeshResult>("tencentcloud:Tcm/getMesh:getMesh", args ?? new GetMeshArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tcm mesh
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mesh = Tencentcloud.Tcm.GetMesh.Invoke(new()
        ///     {
        ///         MeshClusters = new[]
        ///         {
        ///             "cls-xxxx",
        ///         },
        ///         MeshIds = new[]
        ///         {
        ///             "mesh-xxxxxx",
        ///         },
        ///         MeshNames = new[]
        ///         {
        ///             "KEEP_MASH",
        ///         },
        ///         Tags = new[]
        ///         {
        ///             "key",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetMeshResult> Invoke(GetMeshInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMeshResult>("tencentcloud:Tcm/getMesh:getMesh", args ?? new GetMeshInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMeshArgs : global::Pulumi.InvokeArgs
    {
        [Input("meshClusters")]
        private List<string>? _meshClusters;

        /// <summary>
        /// Mesh name.
        /// </summary>
        public List<string> MeshClusters
        {
            get => _meshClusters ?? (_meshClusters = new List<string>());
            set => _meshClusters = value;
        }

        [Input("meshIds")]
        private List<string>? _meshIds;

        /// <summary>
        /// Mesh instance Id.
        /// </summary>
        public List<string> MeshIds
        {
            get => _meshIds ?? (_meshIds = new List<string>());
            set => _meshIds = value;
        }

        [Input("meshNames")]
        private List<string>? _meshNames;

        /// <summary>
        /// Display name.
        /// </summary>
        public List<string> MeshNames
        {
            get => _meshNames ?? (_meshNames = new List<string>());
            set => _meshNames = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// tag key.
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        public GetMeshArgs()
        {
        }
        public static new GetMeshArgs Empty => new GetMeshArgs();
    }

    public sealed class GetMeshInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("meshClusters")]
        private InputList<string>? _meshClusters;

        /// <summary>
        /// Mesh name.
        /// </summary>
        public InputList<string> MeshClusters
        {
            get => _meshClusters ?? (_meshClusters = new InputList<string>());
            set => _meshClusters = value;
        }

        [Input("meshIds")]
        private InputList<string>? _meshIds;

        /// <summary>
        /// Mesh instance Id.
        /// </summary>
        public InputList<string> MeshIds
        {
            get => _meshIds ?? (_meshIds = new InputList<string>());
            set => _meshIds = value;
        }

        [Input("meshNames")]
        private InputList<string>? _meshNames;

        /// <summary>
        /// Display name.
        /// </summary>
        public InputList<string> MeshNames
        {
            get => _meshNames ?? (_meshNames = new InputList<string>());
            set => _meshNames = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// tag key.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public GetMeshInvokeArgs()
        {
        }
        public static new GetMeshInvokeArgs Empty => new GetMeshInvokeArgs();
    }


    [OutputType]
    public sealed class GetMeshResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> MeshClusters;
        /// <summary>
        /// Mesh instance Id.
        /// </summary>
        public readonly ImmutableArray<string> MeshIds;
        /// <summary>
        /// The mesh information is queriedNote: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMeshMeshListResult> MeshLists;
        public readonly ImmutableArray<string> MeshNames;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private GetMeshResult(
            string id,

            ImmutableArray<string> meshClusters,

            ImmutableArray<string> meshIds,

            ImmutableArray<Outputs.GetMeshMeshListResult> meshLists,

            ImmutableArray<string> meshNames,

            string? resultOutputFile,

            ImmutableArray<string> tags)
        {
            Id = id;
            MeshClusters = meshClusters;
            MeshIds = meshIds;
            MeshLists = meshLists;
            MeshNames = meshNames;
            ResultOutputFile = resultOutputFile;
            Tags = tags;
        }
    }
}
