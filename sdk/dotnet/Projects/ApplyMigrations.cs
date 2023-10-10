// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Supabase.Projects
{
    public static class ApplyMigrations
    {
        public static Task<ApplyMigrationsResult> InvokeAsync(ApplyMigrationsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ApplyMigrationsResult>("supabase:projects:applyMigrations", args ?? new ApplyMigrationsArgs(), options.WithDefaults());

        public static Output<ApplyMigrationsResult> Invoke(ApplyMigrationsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ApplyMigrationsResult>("supabase:projects:applyMigrations", args ?? new ApplyMigrationsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ApplyMigrationsArgs : global::Pulumi.InvokeArgs
    {
        [Input("db_host", required: true)]
        public string Db_host { get; set; } = null!;

        [Input("db_pass", required: true)]
        public string Db_pass { get; set; } = null!;

        [Input("include_all")]
        public bool? Include_all { get; set; }

        [Input("include_roles")]
        public bool? Include_roles { get; set; }

        [Input("include_seed")]
        public bool? Include_seed { get; set; }

        public ApplyMigrationsArgs()
        {
        }
        public static new ApplyMigrationsArgs Empty => new ApplyMigrationsArgs();
    }

    public sealed class ApplyMigrationsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("db_host", required: true)]
        public Input<string> Db_host { get; set; } = null!;

        [Input("db_pass", required: true)]
        public Input<string> Db_pass { get; set; } = null!;

        [Input("include_all")]
        public Input<bool>? Include_all { get; set; }

        [Input("include_roles")]
        public Input<bool>? Include_roles { get; set; }

        [Input("include_seed")]
        public Input<bool>? Include_seed { get; set; }

        public ApplyMigrationsInvokeArgs()
        {
        }
        public static new ApplyMigrationsInvokeArgs Empty => new ApplyMigrationsInvokeArgs();
    }


    [OutputType]
    public sealed class ApplyMigrationsResult
    {
        public readonly bool Success;

        [OutputConstructor]
        private ApplyMigrationsResult(bool success)
        {
            Success = success;
        }
    }
}
