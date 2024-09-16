// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package fomantic

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Init() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_Init_24bf`,
		Function: `function __templ_Init_24bf(){user = {
        "user_create" : "/v4/account",
        "user_login"  : "/v4/account",
        "user_update" : "/v4/account",
        "user_delete" : "/v4/account",
    };

    $.api.settings.api = {
        ...user,
    };
    
    $.api.settings.serializeForm = true;

    $.api.settings.contentType = "application/json; charset=UTF-8";

    $.api.settings.beforeSend = function(settings) {
        settings.data = JSON.stringify(settings.data);
        return settings;
    };

    $.api.settings.successTest = function(response) {
        if(response && response.Success) {
            return response.Success;
        }

        return false;
    };
}`,
		Call:       templ.SafeScript(`__templ_Init_24bf`),
		CallInline: templ.SafeScriptInline(`__templ_Init_24bf`),
	}
}

func Fomantic() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
