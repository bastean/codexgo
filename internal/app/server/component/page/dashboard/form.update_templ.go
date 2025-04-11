// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package dashboard

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

const (
	UpdateFormTagID = "update"
)

func UpdateFormInit(formTagID string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_UpdateFormInit_023b`,
		Function: `function __templ_UpdateFormInit_023b(formTagID){$(` + "`" + `#${formTagID}` + "`" + `)
        .form({
            on: "blur",
            inline: true,
            preventLeaving: true,
            keyboardShortcuts: false,
            fields: {
                Email: {
                    optional: true,
                    rules: [
                        {
                            type: "email"
                        }
                    ]
                },
                Username: {
                    optional: true,
                    rules: [
                        {
                            type: "size[2..20]"
                        },
                        {
                            type: "regExp[/^[A-Za-z0-9]+$/]",
                            prompt: "{name} must be alphanumeric only"
                        },
                        {
                            type: "regExp[/^.*[^0-9].*$/]",
                            prompt: "{name} cannot be only numbers"
                        }
                    ]
                },
                UpdatedPassword: {
                    optional: true,
                    rules: [
                        {
                            type: "size[8..64]"
                        },
                        {
                            type: "regExp[/^.*[^0-9].*$/]",
                            prompt: "{name} cannot be only numbers"
                        }
                    ]
                },
                ConfirmPassword: {
                    depends: "UpdatedPassword",
                    rules: [
                        {
                            type: "match[UpdatedPassword]"
                        }
                    ]
                },
                Password: {
                    rules: [
                        {
                            type: "notEmpty",
                            prompt: "{name} is required to update the account settings"
                        },
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
            action: "user_update", 
            method: "PATCH",
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
		Call:       templ.SafeScript(`__templ_UpdateFormInit_023b`, formTagID),
		CallInline: templ.SafeScriptInline(`__templ_UpdateFormInit_023b`, formTagID),
	}
}

func UpdateForm(email, username string) templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<form id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(UpdateFormTagID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/app/server/component/page/dashboard/form.update.templ`, Line: 102, Col: 27}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\" class=\"ui inverted form\"><h1 class=\"ui dividing inverted header\">Account settings</h1><div class=\"required field\"><label>Email</label><div class=\"ui inverted transparent left icon input\"><i class=\"envelope icon\"></i> <input type=\"text\" placeholder=\"Email\" name=\"Email\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(email)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/app/server/component/page/dashboard/form.update.templ`, Line: 110, Col: 69}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\"></div></div><div class=\"required field\"><label>Username</label><div class=\"ui inverted transparent left icon input\"><i class=\"user icon\"></i> <input type=\"text\" placeholder=\"Username\" name=\"Username\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(username)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/app/server/component/page/dashboard/form.update.templ`, Line: 117, Col: 78}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\"></div></div><div class=\"required field\"><label>New Password</label><div class=\"ui inverted transparent left icon input\"><i class=\"lock icon\"></i> <input type=\"password\" placeholder=\"New Password\" name=\"UpdatedPassword\"></div></div><div class=\"required field\"><label>Confirm Password</label><div class=\"ui inverted transparent left icon input\"><i class=\"redo icon\"></i> <input type=\"password\" placeholder=\"Confirm Password\" name=\"ConfirmPassword\"></div></div><div class=\"ui divider\"></div><div class=\"required field\"><label>Current Password</label><div class=\"ui inverted transparent left icon input\"><i class=\"lock icon\"></i> <input type=\"password\" placeholder=\"Current Password\" name=\"Password\"></div></div><div class=\"ui divider\"></div><div class=\"ui grid\"><div class=\"two column row\"><div class=\"ten wide column\"><button class=\"ui fluid primary submit button\">Update</button></div><div class=\"six wide column\"><button class=\"ui animated fluid negative reset button\"><div class=\"visible content\">Reset</div><div class=\"hidden content\"><i class=\"trash icon\"></i></div></button></div></div></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = UpdateFormInit(UpdateFormTagID).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
