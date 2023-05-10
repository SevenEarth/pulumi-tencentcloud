// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf
{
    /// <summary>
    /// Provides a resource to create a tsf application
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
    ///         var application = new Tencentcloud.Tsf.Application("application", new Tencentcloud.Tsf.ApplicationArgs
    ///         {
    ///             ApplicationDesc = "This is my application",
    ///             ApplicationName = "my-app",
    ///             ApplicationRuntimeType = "Java",
    ///             ApplicationType = "C",
    ///             IgnoreCreateImageRepository = true,
    ///             MicroserviceType = "M",
    ///             ServiceConfigLists = 
    ///             {
    ///                 new Tencentcloud.Tsf.Inputs.ApplicationServiceConfigListArgs
    ///                 {
    ///                     HealthCheck = new Tencentcloud.Tsf.Inputs.ApplicationServiceConfigListHealthCheckArgs
    ///                     {
    ///                         Path = "/health",
    ///                     },
    ///                     Name = "my-service",
    ///                     Ports = 
    ///                     {
    ///                         new Tencentcloud.Tsf.Inputs.ApplicationServiceConfigListPortArgs
    ///                         {
    ///                             Protocol = "HTTP",
    ///                             TargetPort = 8080,
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tsf/application:Application")]
    public partial class Application : Pulumi.CustomResource
    {
        /// <summary>
        /// Application description.
        /// </summary>
        [Output("applicationDesc")]
        public Output<string?> ApplicationDesc { get; private set; } = null!;

        /// <summary>
        /// Application log configuration, deprecated parameter.
        /// </summary>
        [Output("applicationLogConfig")]
        public Output<string?> ApplicationLogConfig { get; private set; } = null!;

        /// <summary>
        /// Application name.
        /// </summary>
        [Output("applicationName")]
        public Output<string> ApplicationName { get; private set; } = null!;

        /// <summary>
        /// Application resource type, deprecated parameter.
        /// </summary>
        [Output("applicationResourceType")]
        public Output<string?> ApplicationResourceType { get; private set; } = null!;

        /// <summary>
        /// Application runtime type.
        /// </summary>
        [Output("applicationRuntimeType")]
        public Output<string?> ApplicationRuntimeType { get; private set; } = null!;

        /// <summary>
        /// Application type: V for virtual machine, C for container, S for serverless.
        /// </summary>
        [Output("applicationType")]
        public Output<string> ApplicationType { get; private set; } = null!;

        /// <summary>
        /// Ignore creating image repository.
        /// </summary>
        [Output("ignoreCreateImageRepository")]
        public Output<bool?> IgnoreCreateImageRepository { get; private set; } = null!;

        /// <summary>
        /// Application microservice type: M for service mesh, N for normal application, G for gateway application.
        /// </summary>
        [Output("microserviceType")]
        public Output<string> MicroserviceType { get; private set; } = null!;

        /// <summary>
        /// ID of the dataset to be bound.
        /// </summary>
        [Output("programId")]
        public Output<string?> ProgramId { get; private set; } = null!;

        /// <summary>
        /// N/A.
        /// </summary>
        [Output("programIdLists")]
        public Output<ImmutableArray<string>> ProgramIdLists { get; private set; } = null!;

        /// <summary>
        /// List of service configuration information.
        /// </summary>
        [Output("serviceConfigLists")]
        public Output<ImmutableArray<Outputs.ApplicationServiceConfigList>> ServiceConfigLists { get; private set; } = null!;


        /// <summary>
        /// Create a Application resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Application(string name, ApplicationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/application:Application", name, args ?? new ApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Application(string name, Input<string> id, ApplicationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/application:Application", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Application resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Application Get(string name, Input<string> id, ApplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new Application(name, id, state, options);
        }
    }

    public sealed class ApplicationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application description.
        /// </summary>
        [Input("applicationDesc")]
        public Input<string>? ApplicationDesc { get; set; }

        /// <summary>
        /// Application log configuration, deprecated parameter.
        /// </summary>
        [Input("applicationLogConfig")]
        public Input<string>? ApplicationLogConfig { get; set; }

        /// <summary>
        /// Application name.
        /// </summary>
        [Input("applicationName", required: true)]
        public Input<string> ApplicationName { get; set; } = null!;

        /// <summary>
        /// Application resource type, deprecated parameter.
        /// </summary>
        [Input("applicationResourceType")]
        public Input<string>? ApplicationResourceType { get; set; }

        /// <summary>
        /// Application runtime type.
        /// </summary>
        [Input("applicationRuntimeType")]
        public Input<string>? ApplicationRuntimeType { get; set; }

        /// <summary>
        /// Application type: V for virtual machine, C for container, S for serverless.
        /// </summary>
        [Input("applicationType", required: true)]
        public Input<string> ApplicationType { get; set; } = null!;

        /// <summary>
        /// Ignore creating image repository.
        /// </summary>
        [Input("ignoreCreateImageRepository")]
        public Input<bool>? IgnoreCreateImageRepository { get; set; }

        /// <summary>
        /// Application microservice type: M for service mesh, N for normal application, G for gateway application.
        /// </summary>
        [Input("microserviceType", required: true)]
        public Input<string> MicroserviceType { get; set; } = null!;

        /// <summary>
        /// ID of the dataset to be bound.
        /// </summary>
        [Input("programId")]
        public Input<string>? ProgramId { get; set; }

        [Input("programIdLists")]
        private InputList<string>? _programIdLists;

        /// <summary>
        /// N/A.
        /// </summary>
        public InputList<string> ProgramIdLists
        {
            get => _programIdLists ?? (_programIdLists = new InputList<string>());
            set => _programIdLists = value;
        }

        [Input("serviceConfigLists")]
        private InputList<Inputs.ApplicationServiceConfigListArgs>? _serviceConfigLists;

        /// <summary>
        /// List of service configuration information.
        /// </summary>
        public InputList<Inputs.ApplicationServiceConfigListArgs> ServiceConfigLists
        {
            get => _serviceConfigLists ?? (_serviceConfigLists = new InputList<Inputs.ApplicationServiceConfigListArgs>());
            set => _serviceConfigLists = value;
        }

        public ApplicationArgs()
        {
        }
    }

    public sealed class ApplicationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application description.
        /// </summary>
        [Input("applicationDesc")]
        public Input<string>? ApplicationDesc { get; set; }

        /// <summary>
        /// Application log configuration, deprecated parameter.
        /// </summary>
        [Input("applicationLogConfig")]
        public Input<string>? ApplicationLogConfig { get; set; }

        /// <summary>
        /// Application name.
        /// </summary>
        [Input("applicationName")]
        public Input<string>? ApplicationName { get; set; }

        /// <summary>
        /// Application resource type, deprecated parameter.
        /// </summary>
        [Input("applicationResourceType")]
        public Input<string>? ApplicationResourceType { get; set; }

        /// <summary>
        /// Application runtime type.
        /// </summary>
        [Input("applicationRuntimeType")]
        public Input<string>? ApplicationRuntimeType { get; set; }

        /// <summary>
        /// Application type: V for virtual machine, C for container, S for serverless.
        /// </summary>
        [Input("applicationType")]
        public Input<string>? ApplicationType { get; set; }

        /// <summary>
        /// Ignore creating image repository.
        /// </summary>
        [Input("ignoreCreateImageRepository")]
        public Input<bool>? IgnoreCreateImageRepository { get; set; }

        /// <summary>
        /// Application microservice type: M for service mesh, N for normal application, G for gateway application.
        /// </summary>
        [Input("microserviceType")]
        public Input<string>? MicroserviceType { get; set; }

        /// <summary>
        /// ID of the dataset to be bound.
        /// </summary>
        [Input("programId")]
        public Input<string>? ProgramId { get; set; }

        [Input("programIdLists")]
        private InputList<string>? _programIdLists;

        /// <summary>
        /// N/A.
        /// </summary>
        public InputList<string> ProgramIdLists
        {
            get => _programIdLists ?? (_programIdLists = new InputList<string>());
            set => _programIdLists = value;
        }

        [Input("serviceConfigLists")]
        private InputList<Inputs.ApplicationServiceConfigListGetArgs>? _serviceConfigLists;

        /// <summary>
        /// List of service configuration information.
        /// </summary>
        public InputList<Inputs.ApplicationServiceConfigListGetArgs> ServiceConfigLists
        {
            get => _serviceConfigLists ?? (_serviceConfigLists = new InputList<Inputs.ApplicationServiceConfigListGetArgs>());
            set => _serviceConfigLists = value;
        }

        public ApplicationState()
        {
        }
    }
}