// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Key
{
    /// <summary>
    /// Provides a key pair resource.
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
    ///         var foo = new Tencentcloud.Key.Pair("foo", new Tencentcloud.Key.PairArgs
    ///         {
    ///             KeyName = "terraform_test",
    ///         });
    ///         var foo1 = new Tencentcloud.Key.Pair("foo1", new Tencentcloud.Key.PairArgs
    ///         {
    ///             KeyName = "terraform_test",
    ///             PublicKey = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDjd8fTnp7Dcuj4mLaQxf9Zs/ORgUL9fQxRCNKkPgP1paTy1I513maMX126i36Lxxl3+FUB52oVbo/FgwlIfX8hyCnv8MCxqnuSDozf1CD0/wRYHcTWAtgHQHBPCC2nJtod6cVC3kB18KeV4U7zsxmwFeBIxojMOOmcOBuh7+trRw==",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Key pair can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Key/pair:Pair foo skey-17634f05
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Key/pair:Pair")]
    public partial class Pair : Pulumi.CustomResource
    {
        /// <summary>
        /// The key pair's name. It is the only in one TencentCloud account.
        /// </summary>
        [Output("keyName")]
        public Output<string> KeyName { get; private set; } = null!;

        /// <summary>
        /// Specifys to which project the key pair belongs.
        /// </summary>
        [Output("projectId")]
        public Output<int?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// You can import an existing public key and using TencentCloud key pair to manage it.
        /// </summary>
        [Output("publicKey")]
        public Output<string> PublicKey { get; private set; } = null!;

        /// <summary>
        /// Tags of the key pair.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Pair resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Pair(string name, PairArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Key/pair:Pair", name, args ?? new PairArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Pair(string name, Input<string> id, PairState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Key/pair:Pair", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Pair resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Pair Get(string name, Input<string> id, PairState? state = null, CustomResourceOptions? options = null)
        {
            return new Pair(name, id, state, options);
        }
    }

    public sealed class PairArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key pair's name. It is the only in one TencentCloud account.
        /// </summary>
        [Input("keyName", required: true)]
        public Input<string> KeyName { get; set; } = null!;

        /// <summary>
        /// Specifys to which project the key pair belongs.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// You can import an existing public key and using TencentCloud key pair to manage it.
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags of the key pair.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public PairArgs()
        {
        }
    }

    public sealed class PairState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key pair's name. It is the only in one TencentCloud account.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// Specifys to which project the key pair belongs.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// You can import an existing public key and using TencentCloud key pair to manage it.
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags of the key pair.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public PairState()
        {
        }
    }
}
