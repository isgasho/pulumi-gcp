// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The queryParameterMatch matches if the value of the parameter exactly matches
        /// the contents of exactMatch. Only one of presentMatch, exactMatch and regexMatch
        /// must be set.
        /// </summary>
        [Input("exactMatch")]
        public Input<string>? ExactMatch { get; set; }

        /// <summary>
        /// The name of the header.
        /// </summary>
        [Input("headerName", required: true)]
        public Input<string> HeaderName { get; set; } = null!;

        /// <summary>
        /// If set to false, the headerMatch is considered a match if the match criteria
        /// above are met. If set to true, the headerMatch is considered a match if the
        /// match criteria above are NOT met. Defaults to false.
        /// </summary>
        [Input("invertMatch")]
        public Input<bool>? InvertMatch { get; set; }

        /// <summary>
        /// The value of the header must start with the contents of prefixMatch. Only one of
        /// exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch
        /// must be set.
        /// </summary>
        [Input("prefixMatch")]
        public Input<string>? PrefixMatch { get; set; }

        /// <summary>
        /// Specifies that the queryParameterMatch matches if the request contains the query
        /// parameter, irrespective of whether the parameter has a value or not. Only one of
        /// presentMatch, exactMatch and regexMatch must be set.
        /// </summary>
        [Input("presentMatch")]
        public Input<bool>? PresentMatch { get; set; }

        /// <summary>
        /// The header value must be an integer and its value must be in the range specified
        /// in rangeMatch. If the header does not contain an integer, number or is empty,
        /// the match fails. For example for a range [-5, 0]
        /// * -3 will match
        /// * 0 will not match
        /// * 0.25 will not match
        /// * -3someString will not match.
        /// Only one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or
        /// rangeMatch must be set.
        /// Structure is documented below.
        /// </summary>
        [Input("rangeMatch")]
        public Input<Inputs.RegionUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchArgs>? RangeMatch { get; set; }

        /// <summary>
        /// The queryParameterMatch matches if the value of the parameter matches the
        /// regular expression specified by regexMatch. For the regular expression grammar,
        /// please see en.cppreference.com/w/cpp/regex/ecmascript  Only one of presentMatch,
        /// exactMatch and regexMatch must be set.
        /// </summary>
        [Input("regexMatch")]
        public Input<string>? RegexMatch { get; set; }

        /// <summary>
        /// The value of the header must end with the contents of suffixMatch. Only one of
        /// exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch
        /// must be set.
        /// </summary>
        [Input("suffixMatch")]
        public Input<string>? SuffixMatch { get; set; }

        public RegionUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchArgs()
        {
        }
    }
}
