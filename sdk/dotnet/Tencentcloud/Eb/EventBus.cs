// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Eb
{
    /// <summary>
    /// Provides a resource to create a eb event_bus
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
    ///         var foo = new Tencentcloud.Eb.EventBus("foo", new Tencentcloud.Eb.EventBusArgs
    ///         {
    ///             Description = "event bus desc",
    ///             EnableStore = false,
    ///             EventBusName = "tf-event_bus",
    ///             SaveDays = 1,
    ///             Tags = 
    ///             {
    ///                 { "createdBy", "terraform" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// eb event_bus can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Eb/eventBus:EventBus event_bus event_bus_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Eb/eventBus:EventBus")]
    public partial class EventBus : Pulumi.CustomResource
    {
        /// <summary>
        /// Event set description, unlimited character type, description within 200 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether the EB storage is enabled.
        /// </summary>
        [Output("enableStore")]
        public Output<bool?> EnableStore { get; private set; } = null!;

        /// <summary>
        /// Event set name, which can only contain letters, numbers, underscores, hyphens, starts with a letter and ends with a number or letter, 2~60 characters.
        /// </summary>
        [Output("eventBusName")]
        public Output<string> EventBusName { get; private set; } = null!;

        /// <summary>
        /// EB storage duration.
        /// </summary>
        [Output("saveDays")]
        public Output<int?> SaveDays { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a EventBus resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventBus(string name, EventBusArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Eb/eventBus:EventBus", name, args ?? new EventBusArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventBus(string name, Input<string> id, EventBusState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Eb/eventBus:EventBus", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventBus resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventBus Get(string name, Input<string> id, EventBusState? state = null, CustomResourceOptions? options = null)
        {
            return new EventBus(name, id, state, options);
        }
    }

    public sealed class EventBusArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Event set description, unlimited character type, description within 200 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the EB storage is enabled.
        /// </summary>
        [Input("enableStore")]
        public Input<bool>? EnableStore { get; set; }

        /// <summary>
        /// Event set name, which can only contain letters, numbers, underscores, hyphens, starts with a letter and ends with a number or letter, 2~60 characters.
        /// </summary>
        [Input("eventBusName", required: true)]
        public Input<string> EventBusName { get; set; } = null!;

        /// <summary>
        /// EB storage duration.
        /// </summary>
        [Input("saveDays")]
        public Input<int>? SaveDays { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public EventBusArgs()
        {
        }
    }

    public sealed class EventBusState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Event set description, unlimited character type, description within 200 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the EB storage is enabled.
        /// </summary>
        [Input("enableStore")]
        public Input<bool>? EnableStore { get; set; }

        /// <summary>
        /// Event set name, which can only contain letters, numbers, underscores, hyphens, starts with a letter and ends with a number or letter, 2~60 characters.
        /// </summary>
        [Input("eventBusName")]
        public Input<string>? EventBusName { get; set; }

        /// <summary>
        /// EB storage duration.
        /// </summary>
        [Input("saveDays")]
        public Input<int>? SaveDays { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public EventBusState()
        {
        }
    }
}
