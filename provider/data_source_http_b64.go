package provider

import (
    "context"
    "fmt"

    "github.com/hashicorp/terraform-plugin-log/tflog"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceScaffolding() *schema.Resource {
    return &schema.Resource{

        // This description is used by the documentation generator and the language server.
        Description: `
        TBD
        `,

        ReadContext: Read,

        Schema: map[string]*schema.Schema{
            "url": {
                Description: "The URL for the request. Supported schemes are `http` and `https`.",
                Type:        schema.TypeString,
                Required:    true,
            },

            "response_body_base64": {
                Description: "The response body returned as a string.",
                Type:        schema.TypeString,
                Computed:    true,
            },

            "status_code": {
                Description: `The HTTP response status code.`,
                Type:        schema.TypeInt,
                Computed:    true,
            },

            "id": {
                Description: "The ID of this resource.",
                Type:        schema.TypeString,
                Computed:    true,
            },
        },
    }
}

func Read(ctx context.Context, req *schema.ResourceData, meta interface{}) diag.Diagnostics {
    // d := diag.Diagnostics{}
    url := req.Get("url").(string)

    // d = append(d, diag.Diagnostic{
    //    Summary: "Error reading response body",
    //    Detail:  fmt.Sprintf("Error reading response body: %s", ""),
    // })
    // return d

    tflog.Info(ctx, fmt.Sprintf("\nStarting.. requesting URL [%s] \n", url))

    return nil
}
