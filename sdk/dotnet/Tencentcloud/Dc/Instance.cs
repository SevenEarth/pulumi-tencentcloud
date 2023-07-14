// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dc
{
    /// <summary>
    /// Provides a resource to create a dc instance
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
    ///         var instance = new Tencentcloud.Dc.Instance("instance", new Tencentcloud.Dc.InstanceArgs
    ///         {
    ///             AccessPointId = "ap-shenzhen-b-ft",
    ///             Bandwidth = 10,
    ///             CustomerContactNumber = "0",
    ///             DirectConnectName = "terraform-for-test",
    ///             LineOperator = "In-houseWiring",
    ///             PortType = "10GBase-LR",
    ///             SignLaw = true,
    ///             Vlan = -1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// dc instance can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Dc/instance:Instance instance dc_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dc/instance:Instance")]
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// Access point of connection.You can call `DescribeAccessPoints` to get the region ID. The selected access point must exist and be available.
        /// </summary>
        [Output("accessPointId")]
        public Output<string> AccessPointId { get; private set; } = null!;

        /// <summary>
        /// Connection port bandwidth in Mbps. Value range: [2,10240]. Default value: 1000.
        /// </summary>
        [Output("bandwidth")]
        public Output<int?> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// Circuit code of a connection, which is provided by the ISP or connection provider.
        /// </summary>
        [Output("circuitCode")]
        public Output<string?> CircuitCode { get; private set; } = null!;

        /// <summary>
        /// User-side IP address for connection debugging, which is automatically assigned by default.
        /// </summary>
        [Output("customerAddress")]
        public Output<string?> CustomerAddress { get; private set; } = null!;

        /// <summary>
        /// Email address of connection applicant, which is obtained from the account system by default.
        /// </summary>
        [Output("customerContactMail")]
        public Output<string?> CustomerContactMail { get; private set; } = null!;

        /// <summary>
        /// Contact number of connection applicant, which is obtained from the account system by default.
        /// </summary>
        [Output("customerContactNumber")]
        public Output<string?> CustomerContactNumber { get; private set; } = null!;

        /// <summary>
        /// Name of connection applicant, which is obtained from the account system by default.
        /// </summary>
        [Output("customerName")]
        public Output<string?> CustomerName { get; private set; } = null!;

        /// <summary>
        /// Connection name.
        /// </summary>
        [Output("directConnectName")]
        public Output<string> DirectConnectName { get; private set; } = null!;

        /// <summary>
        /// Fault reporting contact number.
        /// </summary>
        [Output("faultReportContactNumber")]
        public Output<string?> FaultReportContactNumber { get; private set; } = null!;

        /// <summary>
        /// Fault reporting contact person.
        /// </summary>
        [Output("faultReportContactPerson")]
        public Output<string?> FaultReportContactPerson { get; private set; } = null!;

        /// <summary>
        /// ISP that provides connections. Valid values: ChinaTelecom (China Telecom), ChinaMobile (China Mobile), ChinaUnicom (China Unicom), In-houseWiring (in-house wiring), ChinaOther (other Chinese ISPs), InternationalOperator (international ISPs).
        /// </summary>
        [Output("lineOperator")]
        public Output<string> LineOperator { get; private set; } = null!;

        /// <summary>
        /// Local IDC location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Port type of connection. Valid values: 100Base-T (100-Megabit electrical Ethernet interface), 1000Base-T (1-Gigabit electrical Ethernet interface), 1000Base-LX (1-Gigabit single-module optical Ethernet interface; 10 KM), 10GBase-T (10-Gigabit electrical Ethernet interface), 10GBase-LR (10-Gigabit single-module optical Ethernet interface; 10 KM). Default value: 1000Base-LX.
        /// </summary>
        [Output("portType")]
        public Output<string> PortType { get; private set; } = null!;

        /// <summary>
        /// ID of redundant connection.
        /// </summary>
        [Output("redundantDirectConnectId")]
        public Output<string?> RedundantDirectConnectId { get; private set; } = null!;

        /// <summary>
        /// Whether the connection applicant has signed the service agreement. Default value: true.
        /// </summary>
        [Output("signLaw")]
        public Output<bool?> SignLaw { get; private set; } = null!;

        /// <summary>
        /// Tencent-side IP address for connection debugging, which is automatically assigned by default.
        /// </summary>
        [Output("tencentAddress")]
        public Output<string?> TencentAddress { get; private set; } = null!;

        /// <summary>
        /// VLAN for connection debugging, which is enabled and automatically assigned by default.
        /// </summary>
        [Output("vlan")]
        public Output<int?> Vlan { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dc/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dc/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access point of connection.You can call `DescribeAccessPoints` to get the region ID. The selected access point must exist and be available.
        /// </summary>
        [Input("accessPointId", required: true)]
        public Input<string> AccessPointId { get; set; } = null!;

        /// <summary>
        /// Connection port bandwidth in Mbps. Value range: [2,10240]. Default value: 1000.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// Circuit code of a connection, which is provided by the ISP or connection provider.
        /// </summary>
        [Input("circuitCode")]
        public Input<string>? CircuitCode { get; set; }

        /// <summary>
        /// User-side IP address for connection debugging, which is automatically assigned by default.
        /// </summary>
        [Input("customerAddress")]
        public Input<string>? CustomerAddress { get; set; }

        /// <summary>
        /// Email address of connection applicant, which is obtained from the account system by default.
        /// </summary>
        [Input("customerContactMail")]
        public Input<string>? CustomerContactMail { get; set; }

        /// <summary>
        /// Contact number of connection applicant, which is obtained from the account system by default.
        /// </summary>
        [Input("customerContactNumber")]
        public Input<string>? CustomerContactNumber { get; set; }

        /// <summary>
        /// Name of connection applicant, which is obtained from the account system by default.
        /// </summary>
        [Input("customerName")]
        public Input<string>? CustomerName { get; set; }

        /// <summary>
        /// Connection name.
        /// </summary>
        [Input("directConnectName", required: true)]
        public Input<string> DirectConnectName { get; set; } = null!;

        /// <summary>
        /// Fault reporting contact number.
        /// </summary>
        [Input("faultReportContactNumber")]
        public Input<string>? FaultReportContactNumber { get; set; }

        /// <summary>
        /// Fault reporting contact person.
        /// </summary>
        [Input("faultReportContactPerson")]
        public Input<string>? FaultReportContactPerson { get; set; }

        /// <summary>
        /// ISP that provides connections. Valid values: ChinaTelecom (China Telecom), ChinaMobile (China Mobile), ChinaUnicom (China Unicom), In-houseWiring (in-house wiring), ChinaOther (other Chinese ISPs), InternationalOperator (international ISPs).
        /// </summary>
        [Input("lineOperator", required: true)]
        public Input<string> LineOperator { get; set; } = null!;

        /// <summary>
        /// Local IDC location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Port type of connection. Valid values: 100Base-T (100-Megabit electrical Ethernet interface), 1000Base-T (1-Gigabit electrical Ethernet interface), 1000Base-LX (1-Gigabit single-module optical Ethernet interface; 10 KM), 10GBase-T (10-Gigabit electrical Ethernet interface), 10GBase-LR (10-Gigabit single-module optical Ethernet interface; 10 KM). Default value: 1000Base-LX.
        /// </summary>
        [Input("portType", required: true)]
        public Input<string> PortType { get; set; } = null!;

        /// <summary>
        /// ID of redundant connection.
        /// </summary>
        [Input("redundantDirectConnectId")]
        public Input<string>? RedundantDirectConnectId { get; set; }

        /// <summary>
        /// Whether the connection applicant has signed the service agreement. Default value: true.
        /// </summary>
        [Input("signLaw")]
        public Input<bool>? SignLaw { get; set; }

        /// <summary>
        /// Tencent-side IP address for connection debugging, which is automatically assigned by default.
        /// </summary>
        [Input("tencentAddress")]
        public Input<string>? TencentAddress { get; set; }

        /// <summary>
        /// VLAN for connection debugging, which is enabled and automatically assigned by default.
        /// </summary>
        [Input("vlan")]
        public Input<int>? Vlan { get; set; }

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access point of connection.You can call `DescribeAccessPoints` to get the region ID. The selected access point must exist and be available.
        /// </summary>
        [Input("accessPointId")]
        public Input<string>? AccessPointId { get; set; }

        /// <summary>
        /// Connection port bandwidth in Mbps. Value range: [2,10240]. Default value: 1000.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// Circuit code of a connection, which is provided by the ISP or connection provider.
        /// </summary>
        [Input("circuitCode")]
        public Input<string>? CircuitCode { get; set; }

        /// <summary>
        /// User-side IP address for connection debugging, which is automatically assigned by default.
        /// </summary>
        [Input("customerAddress")]
        public Input<string>? CustomerAddress { get; set; }

        /// <summary>
        /// Email address of connection applicant, which is obtained from the account system by default.
        /// </summary>
        [Input("customerContactMail")]
        public Input<string>? CustomerContactMail { get; set; }

        /// <summary>
        /// Contact number of connection applicant, which is obtained from the account system by default.
        /// </summary>
        [Input("customerContactNumber")]
        public Input<string>? CustomerContactNumber { get; set; }

        /// <summary>
        /// Name of connection applicant, which is obtained from the account system by default.
        /// </summary>
        [Input("customerName")]
        public Input<string>? CustomerName { get; set; }

        /// <summary>
        /// Connection name.
        /// </summary>
        [Input("directConnectName")]
        public Input<string>? DirectConnectName { get; set; }

        /// <summary>
        /// Fault reporting contact number.
        /// </summary>
        [Input("faultReportContactNumber")]
        public Input<string>? FaultReportContactNumber { get; set; }

        /// <summary>
        /// Fault reporting contact person.
        /// </summary>
        [Input("faultReportContactPerson")]
        public Input<string>? FaultReportContactPerson { get; set; }

        /// <summary>
        /// ISP that provides connections. Valid values: ChinaTelecom (China Telecom), ChinaMobile (China Mobile), ChinaUnicom (China Unicom), In-houseWiring (in-house wiring), ChinaOther (other Chinese ISPs), InternationalOperator (international ISPs).
        /// </summary>
        [Input("lineOperator")]
        public Input<string>? LineOperator { get; set; }

        /// <summary>
        /// Local IDC location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Port type of connection. Valid values: 100Base-T (100-Megabit electrical Ethernet interface), 1000Base-T (1-Gigabit electrical Ethernet interface), 1000Base-LX (1-Gigabit single-module optical Ethernet interface; 10 KM), 10GBase-T (10-Gigabit electrical Ethernet interface), 10GBase-LR (10-Gigabit single-module optical Ethernet interface; 10 KM). Default value: 1000Base-LX.
        /// </summary>
        [Input("portType")]
        public Input<string>? PortType { get; set; }

        /// <summary>
        /// ID of redundant connection.
        /// </summary>
        [Input("redundantDirectConnectId")]
        public Input<string>? RedundantDirectConnectId { get; set; }

        /// <summary>
        /// Whether the connection applicant has signed the service agreement. Default value: true.
        /// </summary>
        [Input("signLaw")]
        public Input<bool>? SignLaw { get; set; }

        /// <summary>
        /// Tencent-side IP address for connection debugging, which is automatically assigned by default.
        /// </summary>
        [Input("tencentAddress")]
        public Input<string>? TencentAddress { get; set; }

        /// <summary>
        /// VLAN for connection debugging, which is enabled and automatically assigned by default.
        /// </summary>
        [Input("vlan")]
        public Input<int>? Vlan { get; set; }

        public InstanceState()
        {
        }
    }
}
