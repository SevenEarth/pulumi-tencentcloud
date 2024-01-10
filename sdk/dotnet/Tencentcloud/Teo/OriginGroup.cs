// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo
{
    /// <summary>
    /// Provides a resource to create a teo origin_group
    /// 
    /// ## Example Usage
    /// ### Self origin group
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var originGroup = new Tencentcloud.Teo.OriginGroup("originGroup", new Tencentcloud.Teo.OriginGroupArgs
    ///         {
    ///             ConfigurationType = "weight",
    ///             OriginGroupName = "test-group",
    ///             OriginRecords = 
    ///             {
    ///                 new Tencentcloud.Teo.Inputs.OriginGroupOriginRecordArgs
    ///                 {
    ///                     Areas = {},
    ///                     Port = 8080,
    ///                     Private = false,
    ///                     Record = "150.109.8.1",
    ///                     Weight = 100,
    ///                 },
    ///             },
    ///             OriginType = "self",
    ///             ZoneId = "zone-297z8rf93cfw",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Cos origin group
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var originGroup = new Tencentcloud.Teo.OriginGroup("originGroup", new Tencentcloud.Teo.OriginGroupArgs
    ///         {
    ///             ConfigurationType = "weight",
    ///             OriginGroupName = "test",
    ///             OriginRecords = 
    ///             {
    ///                 new Tencentcloud.Teo.Inputs.OriginGroupOriginRecordArgs
    ///                 {
    ///                     Areas = {},
    ///                     Port = 0,
    ///                     Private = true,
    ///                     Record = "test-ruichaolin-1310708577.cos.ap-nanjing.myqcloud.com",
    ///                     Weight = 100,
    ///                 },
    ///             },
    ///             OriginType = "cos",
    ///             ZoneId = "zone-2o3h21ed8bpu",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// teo origin_group can be imported using the zone_id#originGroup_id, e.g. `
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Teo/originGroup:OriginGroup origin_group zone-297z8rf93cfw#origin-4f8a30b2-3720-11ed-b66b-525400dceb86
    /// ```
    /// 
    ///  `
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Teo/originGroup:OriginGroup")]
    public partial class OriginGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Type of the origin group, this field should be set when `OriginType` is self, otherwise leave it empty. Valid values: `area`: select an origin by using Geo info of the client IP and `Area` field in Records; `weight`: weighted select an origin by using `Weight` field in Records; `proto`: config by HTTP protocol.
        /// </summary>
        [Output("configurationType")]
        public Output<string> ConfigurationType { get; private set; } = null!;

        /// <summary>
        /// OriginGroup ID.
        /// </summary>
        [Output("originGroupId")]
        public Output<string> OriginGroupId { get; private set; } = null!;

        /// <summary>
        /// OriginGroup Name.
        /// </summary>
        [Output("originGroupName")]
        public Output<string> OriginGroupName { get; private set; } = null!;

        /// <summary>
        /// Origin site records.
        /// </summary>
        [Output("originRecords")]
        public Output<ImmutableArray<Outputs.OriginGroupOriginRecord>> OriginRecords { get; private set; } = null!;

        /// <summary>
        /// Type of the origin site. Valid values: `self`: self-build website; `cos`: tencent cos; `third_party`: third party cos.
        /// </summary>
        [Output("originType")]
        public Output<string> OriginType { get; private set; } = null!;

        /// <summary>
        /// Last modification date.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Site ID.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a OriginGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OriginGroup(string name, OriginGroupArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Teo/originGroup:OriginGroup", name, args ?? new OriginGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OriginGroup(string name, Input<string> id, OriginGroupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Teo/originGroup:OriginGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OriginGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OriginGroup Get(string name, Input<string> id, OriginGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new OriginGroup(name, id, state, options);
        }
    }

    public sealed class OriginGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of the origin group, this field should be set when `OriginType` is self, otherwise leave it empty. Valid values: `area`: select an origin by using Geo info of the client IP and `Area` field in Records; `weight`: weighted select an origin by using `Weight` field in Records; `proto`: config by HTTP protocol.
        /// </summary>
        [Input("configurationType", required: true)]
        public Input<string> ConfigurationType { get; set; } = null!;

        /// <summary>
        /// OriginGroup Name.
        /// </summary>
        [Input("originGroupName", required: true)]
        public Input<string> OriginGroupName { get; set; } = null!;

        [Input("originRecords", required: true)]
        private InputList<Inputs.OriginGroupOriginRecordArgs>? _originRecords;

        /// <summary>
        /// Origin site records.
        /// </summary>
        public InputList<Inputs.OriginGroupOriginRecordArgs> OriginRecords
        {
            get => _originRecords ?? (_originRecords = new InputList<Inputs.OriginGroupOriginRecordArgs>());
            set => _originRecords = value;
        }

        /// <summary>
        /// Type of the origin site. Valid values: `self`: self-build website; `cos`: tencent cos; `third_party`: third party cos.
        /// </summary>
        [Input("originType", required: true)]
        public Input<string> OriginType { get; set; } = null!;

        /// <summary>
        /// Site ID.
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public OriginGroupArgs()
        {
        }
    }

    public sealed class OriginGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of the origin group, this field should be set when `OriginType` is self, otherwise leave it empty. Valid values: `area`: select an origin by using Geo info of the client IP and `Area` field in Records; `weight`: weighted select an origin by using `Weight` field in Records; `proto`: config by HTTP protocol.
        /// </summary>
        [Input("configurationType")]
        public Input<string>? ConfigurationType { get; set; }

        /// <summary>
        /// OriginGroup ID.
        /// </summary>
        [Input("originGroupId")]
        public Input<string>? OriginGroupId { get; set; }

        /// <summary>
        /// OriginGroup Name.
        /// </summary>
        [Input("originGroupName")]
        public Input<string>? OriginGroupName { get; set; }

        [Input("originRecords")]
        private InputList<Inputs.OriginGroupOriginRecordGetArgs>? _originRecords;

        /// <summary>
        /// Origin site records.
        /// </summary>
        public InputList<Inputs.OriginGroupOriginRecordGetArgs> OriginRecords
        {
            get => _originRecords ?? (_originRecords = new InputList<Inputs.OriginGroupOriginRecordGetArgs>());
            set => _originRecords = value;
        }

        /// <summary>
        /// Type of the origin site. Valid values: `self`: self-build website; `cos`: tencent cos; `third_party`: third party cos.
        /// </summary>
        [Input("originType")]
        public Input<string>? OriginType { get; set; }

        /// <summary>
        /// Last modification date.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// Site ID.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public OriginGroupState()
        {
        }
    }
}
