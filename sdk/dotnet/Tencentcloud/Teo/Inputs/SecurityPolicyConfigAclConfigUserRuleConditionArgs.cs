// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Inputs
{

    public sealed class SecurityPolicyConfigAclConfigUserRuleConditionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Content to match.
        /// </summary>
        [Input("matchContent", required: true)]
        public Input<string> MatchContent { get; set; } = null!;

        /// <summary>
        /// Items to match. Valid values:- `host`: Host of the request.- `sip`: Client IP.- `ua`: User-Agent.- `cookie`: Session cookie.- `cgi`: CGI script.- `xff`: XFF extension header.- `url`: URL of the request.- `accept`: Accept encoding of the request.- `method`: HTTP method of the request.- `header`: HTTP header of the request.- `sip_proto`: Network protocol of the request.
        /// </summary>
        [Input("matchFrom", required: true)]
        public Input<string> MatchFrom { get; set; } = null!;

        /// <summary>
        /// Parameter for match item. For example, when match from header, match parameter can be set to a header key.
        /// </summary>
        [Input("matchParam", required: true)]
        public Input<string> MatchParam { get; set; } = null!;

        /// <summary>
        /// Valid values:- `equal`: string equal.- `not_equal`: string not equal.- `include`: string include.- `not_include`: string not include.- `match`: ip match.- `not_match`: ip not match.- `include_area`: area include.- `is_empty`: field existed but empty.- `not_exists`: field is not existed.- `regexp`: regex match.- `len_gt`: value greater than.- `len_lt`: value less than.- `len_eq`: value equal.- `match_prefix`: string prefix match.- `match_suffix`: string suffix match.- `wildcard`: wildcard match.
        /// </summary>
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        public SecurityPolicyConfigAclConfigUserRuleConditionArgs()
        {
        }
    }
}
