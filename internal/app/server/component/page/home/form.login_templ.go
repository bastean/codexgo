// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

const (
	LoginFormTagId = "login"
)

func LoginFormInit(formTagId string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_LoginFormInit_a643`,
		Function: `function __templ_LoginFormInit_a643(formTagId){$(` + "`" + `#${formTagId}` + "`" + `)
        .form({
            on: "blur",
            inline: true,
            preventLeaving: true,
            keyboardShortcuts: false,
            fields: {
                Email: {
                    rules: [
                        {
                            type: "email"
                        }
                    ]
                },
                Password: {
                    rules: [
                        {
                            type: "size[8..64]"
                        },
                        {
                            type: "regExp[/^.*[^0-9].*$/]",
                            prompt: "{name} cannot be only numbers"
                        }
                    ]
                }
            }
        })
        .api({
            action: "user_login", 
            method: "POST",
            onSuccess: function(response, element, xhr) {
                $.toast({
                    class: "success",
                    message: response.Message,
                    showProgress: "top"
                });

                _.delay(function() {
                    window.location.replace("/dashboard");
                }, 1000);
            },
            onFailure: function(response, element, xhr) {
                $.toast({
                    class: "error",
                    message: response.Message,
                    showProgress: "top"
                });
            }
        })
    ;
}`,
		Call:       templ.SafeScript(`__templ_LoginFormInit_a643`, formTagId),
		CallInline: templ.SafeScriptInline(`__templ_LoginFormInit_a643`, formTagId),
	}
}

func LoginForm() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(LoginFormTagId)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/app/server/component/page/home/form.login.templ`, Line: 61, Col: 26}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"ui inverted form\"><h1 class=\"ui dividing inverted header\">Sign in to your account<div class=\"sub header\">Don't have an account? ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, ShowTab(RegisterTabTagId))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a style=\"cursor: pointer;\" onclick=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 templ.ComponentScript = ShowTab(RegisterTabTagId)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Sign up</a></div></h1><div class=\"required field\"><label>Email</label><div class=\"ui inverted transparent left icon input\"><i class=\"envelope icon\"></i> <input type=\"text\" placeholder=\"Email\" name=\"Email\"></div></div><div class=\"required field\"><label>Password</label><div class=\"ui inverted transparent left icon input\"><i class=\"lock icon\"></i> <input type=\"password\" placeholder=\"Password\" name=\"Password\"></div></div><div class=\"ui divider\"></div><button class=\"ui fluid primary submit button\">Sign in</button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = LoginFormInit(LoginFormTagId).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
