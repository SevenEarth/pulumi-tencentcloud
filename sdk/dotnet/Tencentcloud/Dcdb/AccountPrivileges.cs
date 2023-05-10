// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb
{
    /// <summary>
    /// Provides a resource to create a dcdb account_privileges
    /// 
    /// ## Import
    /// 
    /// dcdb account_privileges can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Dcdb/accountPrivileges:AccountPrivileges account_privileges instanceId#userName#host#dbName#tabName#viewName#colName
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dcdb/accountPrivileges:AccountPrivileges")]
    public partial class AccountPrivileges : Pulumi.CustomResource
    {
        /// <summary>
        /// The account of the database, including username and host.
        /// </summary>
        [Output("account")]
        public Output<Outputs.AccountPrivilegesAccount> Account { get; private set; } = null!;

        /// <summary>
        /// &amp;amp;quot;Permissions for columns in database tables. Optional values for the Privileges permission are:&amp;amp;quot;&amp;amp;quot;SELECT, INSERT, UPDATE, REFERENCES.&amp;amp;quot;&amp;amp;quot;Note that if this parameter is not passed, the existing privileges are reserved. If you need to clear them, please pass an empty array in the complex type Privileges field.&amp;amp;quot;.
        /// </summary>
        [Output("columnPrivileges")]
        public Output<Outputs.AccountPrivilegesColumnPrivileges?> ColumnPrivileges { get; private set; } = null!;

        /// <summary>
        /// &amp;amp;quot;Database permissions. Optional values for the Privileges permission are:&amp;amp;quot;&amp;amp;quot;SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE TEMPORARY TABLES,&amp;amp;quot;&amp;amp;quot;LOCK TABLES, EXECUTE, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER.&amp;amp;quot;&amp;amp;quot;Note that if this parameter is not passed, the existing privileges are reserved. If you need to clear them, please pass an empty array in the complex type Privileges field.&amp;amp;quot;.
        /// </summary>
        [Output("databasePrivileges")]
        public Output<Outputs.AccountPrivilegesDatabasePrivileges?> DatabasePrivileges { get; private set; } = null!;

        /// <summary>
        /// &amp;amp;quot;Global permissions. Among them, the optional value of the permission in GlobalPrivileges is:&amp;amp;quot;&amp;amp;quot;SELECT, INSERT, UPDATE, DELETE, CREATE, PROCESS, DROP, REFERENCES, INDEX, ALTER, SHOW DATABASES,&amp;amp;quot;&amp;amp;quot;CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER.&amp;amp;quot;&amp;amp;quot;Note that if this parameter is not passed, it means that the existing permissions are reserved. If it needs to be cleared, pass an empty array in this field.&amp;amp;quot;.
        /// </summary>
        [Output("globalPrivileges")]
        public Output<ImmutableArray<string>> GlobalPrivileges { get; private set; } = null!;

        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// &amp;amp;quot;Permissions for tables in the database. Optional values for the Privileges permission are:&amp;amp;quot;&amp;amp;quot;SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE VIEW, SHOW VIEW, TRIGGER.&amp;amp;quot;&amp;amp;quot;Note that if this parameter is not passed, the existing privileges are reserved. If you need to clear them, please pass an empty array in the complex type Privileges field.&amp;amp;quot;.
        /// </summary>
        [Output("tablePrivileges")]
        public Output<Outputs.AccountPrivilegesTablePrivileges?> TablePrivileges { get; private set; } = null!;

        /// <summary>
        /// &amp;amp;quot;Permissions for database views. Optional values for the Privileges permission are:&amp;amp;quot;&amp;amp;quot;SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE VIEW, SHOW VIEW, TRIGGER.&amp;amp;quot;&amp;amp;quot;Note that if this parameter is not passed, the existing privileges are reserved. If you need to clear them, please pass an empty array in the complex type Privileges field.&amp;amp;quot;.
        /// </summary>
        [Output("viewPrivileges")]
        public Output<Outputs.AccountPrivilegesViewPrivileges?> ViewPrivileges { get; private set; } = null!;


        /// <summary>
        /// Create a AccountPrivileges resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountPrivileges(string name, AccountPrivilegesArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dcdb/accountPrivileges:AccountPrivileges", name, args ?? new AccountPrivilegesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountPrivileges(string name, Input<string> id, AccountPrivilegesState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dcdb/accountPrivileges:AccountPrivileges", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccountPrivileges resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountPrivileges Get(string name, Input<string> id, AccountPrivilegesState? state = null, CustomResourceOptions? options = null)
        {
            return new AccountPrivileges(name, id, state, options);
        }
    }

    public sealed class AccountPrivilegesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account of the database, including username and host.
        /// </summary>
        [Input("account", required: true)]
        public Input<Inputs.AccountPrivilegesAccountArgs> Account { get; set; } = null!;

        /// <summary>
        /// &amp;amp;quot;Permissions for columns in database tables. Optional values for the Privileges permission are:&amp;amp;quot;&amp;amp;quot;SELECT, INSERT, UPDATE, REFERENCES.&amp;amp;quot;&amp;amp;quot;Note that if this parameter is not passed, the existing privileges are reserved. If you need to clear them, please pass an empty array in the complex type Privileges field.&amp;amp;quot;.
        /// </summary>
        [Input("columnPrivileges")]
        public Input<Inputs.AccountPrivilegesColumnPrivilegesArgs>? ColumnPrivileges { get; set; }

        /// <summary>
        /// &amp;amp;quot;Database permissions. Optional values for the Privileges permission are:&amp;amp;quot;&amp;amp;quot;SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE TEMPORARY TABLES,&amp;amp;quot;&amp;amp;quot;LOCK TABLES, EXECUTE, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER.&amp;amp;quot;&amp;amp;quot;Note that if this parameter is not passed, the existing privileges are reserved. If you need to clear them, please pass an empty array in the complex type Privileges field.&amp;amp;quot;.
        /// </summary>
        [Input("databasePrivileges")]
        public Input<Inputs.AccountPrivilegesDatabasePrivilegesArgs>? DatabasePrivileges { get; set; }

        [Input("globalPrivileges")]
        private InputList<string>? _globalPrivileges;

        /// <summary>
        /// &amp;amp;quot;Global permissions. Among them, the optional value of the permission in GlobalPrivileges is:&amp;amp;quot;&amp;amp;quot;SELECT, INSERT, UPDATE, DELETE, CREATE, PROCESS, DROP, REFERENCES, INDEX, ALTER, SHOW DATABASES,&amp;amp;quot;&amp;amp;quot;CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER.&amp;amp;quot;&amp;amp;quot;Note that if this parameter is not passed, it means that the existing permissions are reserved. If it needs to be cleared, pass an empty array in this field.&amp;amp;quot;.
        /// </summary>
        public InputList<string> GlobalPrivileges
        {
            get => _globalPrivileges ?? (_globalPrivileges = new InputList<string>());
            set => _globalPrivileges = value;
        }

        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// &amp;amp;quot;Permissions for tables in the database. Optional values for the Privileges permission are:&amp;amp;quot;&amp;amp;quot;SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE VIEW, SHOW VIEW, TRIGGER.&amp;amp;quot;&amp;amp;quot;Note that if this parameter is not passed, the existing privileges are reserved. If you need to clear them, please pass an empty array in the complex type Privileges field.&amp;amp;quot;.
        /// </summary>
        [Input("tablePrivileges")]
        public Input<Inputs.AccountPrivilegesTablePrivilegesArgs>? TablePrivileges { get; set; }

        /// <summary>
        /// &amp;amp;quot;Permissions for database views. Optional values for the Privileges permission are:&amp;amp;quot;&amp;amp;quot;SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE VIEW, SHOW VIEW, TRIGGER.&amp;amp;quot;&amp;amp;quot;Note that if this parameter is not passed, the existing privileges are reserved. If you need to clear them, please pass an empty array in the complex type Privileges field.&amp;amp;quot;.
        /// </summary>
        [Input("viewPrivileges")]
        public Input<Inputs.AccountPrivilegesViewPrivilegesArgs>? ViewPrivileges { get; set; }

        public AccountPrivilegesArgs()
        {
        }
    }

    public sealed class AccountPrivilegesState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account of the database, including username and host.
        /// </summary>
        [Input("account")]
        public Input<Inputs.AccountPrivilegesAccountGetArgs>? Account { get; set; }

        /// <summary>
        /// &amp;amp;quot;Permissions for columns in database tables. Optional values for the Privileges permission are:&amp;amp;quot;&amp;amp;quot;SELECT, INSERT, UPDATE, REFERENCES.&amp;amp;quot;&amp;amp;quot;Note that if this parameter is not passed, the existing privileges are reserved. If you need to clear them, please pass an empty array in the complex type Privileges field.&amp;amp;quot;.
        /// </summary>
        [Input("columnPrivileges")]
        public Input<Inputs.AccountPrivilegesColumnPrivilegesGetArgs>? ColumnPrivileges { get; set; }

        /// <summary>
        /// &amp;amp;quot;Database permissions. Optional values for the Privileges permission are:&amp;amp;quot;&amp;amp;quot;SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE TEMPORARY TABLES,&amp;amp;quot;&amp;amp;quot;LOCK TABLES, EXECUTE, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER.&amp;amp;quot;&amp;amp;quot;Note that if this parameter is not passed, the existing privileges are reserved. If you need to clear them, please pass an empty array in the complex type Privileges field.&amp;amp;quot;.
        /// </summary>
        [Input("databasePrivileges")]
        public Input<Inputs.AccountPrivilegesDatabasePrivilegesGetArgs>? DatabasePrivileges { get; set; }

        [Input("globalPrivileges")]
        private InputList<string>? _globalPrivileges;

        /// <summary>
        /// &amp;amp;quot;Global permissions. Among them, the optional value of the permission in GlobalPrivileges is:&amp;amp;quot;&amp;amp;quot;SELECT, INSERT, UPDATE, DELETE, CREATE, PROCESS, DROP, REFERENCES, INDEX, ALTER, SHOW DATABASES,&amp;amp;quot;&amp;amp;quot;CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER.&amp;amp;quot;&amp;amp;quot;Note that if this parameter is not passed, it means that the existing permissions are reserved. If it needs to be cleared, pass an empty array in this field.&amp;amp;quot;.
        /// </summary>
        public InputList<string> GlobalPrivileges
        {
            get => _globalPrivileges ?? (_globalPrivileges = new InputList<string>());
            set => _globalPrivileges = value;
        }

        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// &amp;amp;quot;Permissions for tables in the database. Optional values for the Privileges permission are:&amp;amp;quot;&amp;amp;quot;SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE VIEW, SHOW VIEW, TRIGGER.&amp;amp;quot;&amp;amp;quot;Note that if this parameter is not passed, the existing privileges are reserved. If you need to clear them, please pass an empty array in the complex type Privileges field.&amp;amp;quot;.
        /// </summary>
        [Input("tablePrivileges")]
        public Input<Inputs.AccountPrivilegesTablePrivilegesGetArgs>? TablePrivileges { get; set; }

        /// <summary>
        /// &amp;amp;quot;Permissions for database views. Optional values for the Privileges permission are:&amp;amp;quot;&amp;amp;quot;SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE VIEW, SHOW VIEW, TRIGGER.&amp;amp;quot;&amp;amp;quot;Note that if this parameter is not passed, the existing privileges are reserved. If you need to clear them, please pass an empty array in the complex type Privileges field.&amp;amp;quot;.
        /// </summary>
        [Input("viewPrivileges")]
        public Input<Inputs.AccountPrivilegesViewPrivilegesGetArgs>? ViewPrivileges { get; set; }

        public AccountPrivilegesState()
        {
        }
    }
}
