package provider

import (
    "fmt"

    "github.com/ackers-bud/terraform-provider-http2b64/client"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
    return &schema.Resource{
        Create: Create,
        Update: Update,
        Read:   ReadUrl,
        Delete: Delete,

        Schema: map[string]*schema.Schema{
            "url": {
                Type:     schema.TypeString,
                Required: true,
            },

            "id": {
                Description: "The ID of this resource.",
                Type:        schema.TypeString,
                Computed:    true,
            },

            "status_code": {
                Description: "the returned http status code",
                Type:        schema.TypeInt,
                Computed:    true,
            },

            "response_body_base64": {
                Description: "The Returned body base64 encoded",
                Type:        schema.TypeString,
                Computed:    true,
            },
        },
    }
}

func Create(d *schema.ResourceData, meta interface{}) error {
    CreateDiagnostics := diag.Diagnostics{}
    url := d.Get("url").(string)

    responseBody, statusCode, err := client.GetFile(url)
    if err != nil {
        CreateDiagnostics = append(CreateDiagnostics, diag.Diagnostic{
            Summary: fmt.Sprintf("Error Getting %v", err),
            Detail:  "Error",
        })
    }

    err = d.Set("status_code", statusCode)
    if err != nil {
        CreateDiagnostics = append(CreateDiagnostics, diag.Diagnostic{
            Summary: fmt.Sprintf("Error Getting %v", err),
            Detail:  "Error",
        })
    }

    err = d.Set("response_body_base64", responseBody)
    if err != nil {
        CreateDiagnostics = append(CreateDiagnostics, diag.Diagnostic{
            Summary: fmt.Sprintf("Error Getting %v", err),
            Detail:  "Error",
        })
    }

    d.SetId(url)
    return nil
}

func Update(d *schema.ResourceData, meta interface{}) error {

    return ReadUrl(d, meta)
}

func ReadUrl(d *schema.ResourceData, meta interface{}) error {

    return nil
}

func Delete(d *schema.ResourceData, meta interface{}) error {

    d.SetId("")

    return nil
}
