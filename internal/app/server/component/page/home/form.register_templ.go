// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/bastean/codexgo/v4/internal/app/server/service/captcha"

const (
	RegisterFormTagID   = "register-form"
	RegisterSubmitTagID = "register-submit"
)

func RegisterFormInit(submitTagID, formTagID, loginTabTagID string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_RegisterFormInit_e044`,
		Function: `function __templ_RegisterFormInit_e044(submitTagID, formTagID, loginTabTagID){$(` + "`" + `#${submitTagID}` + "`" + `)
        .popup({
            position: "top center",
            hoverable: true
        })
    ;

    $(` + "`" + `#${formTagID}` + "`" + `)
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
                Username: {
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
                },
                ConfirmPassword: {
                    rules: [
                        {
                            type: "match[Password]"
                        }
                    ]
                },
                Terms: {
                    rules: [
                        {
                            type: "checked",
                            prompt: "Terms and Conditions must be checked"
                        }
                    ]
                },
                CaptchaAnswer: {
                    rules: [
                        {
                            type: "notEmpty"
                        }
                    ]
                }
            }
        })
        .api({
            action: "user_create", 
            method: "PUT",
            beforeSend: function(settings) {
                settings.data.CaptchaAnswer = _.toString(settings.data.CaptchaAnswer);

                settings.data = JSON.stringify(settings.data);

                return settings;
            },
            onSuccess: function(response, element, xhr) {
                $.toast({
                    class: "success",
                    message: response.Message,
                    showProgress: "top"
                });

                _.delay(function() {
                    $.tab("change tab", loginTabTagID);
                    $(` + "`" + `#${formTagID}` + "`" + `).form("reset");
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
		Call:       templ.SafeScript(`__templ_RegisterFormInit_e044`, submitTagID, formTagID, loginTabTagID),
		CallInline: templ.SafeScriptInline(`__templ_RegisterFormInit_e044`, submitTagID, formTagID, loginTabTagID),
	}
}

func RegisterForm(captcha *captcha.Captcha) templ.Component {
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
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(RegisterFormTagID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/app/server/component/page/home/form.register.templ`, Line: 116, Col: 29}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\" class=\"ui inverted form\"><h1 class=\"ui dividing inverted header\">Create an account<div class=\"sub header\">Already have an account? ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, ShowTab(LoginTabTagID))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<a style=\"cursor: pointer;\" onclick=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 templ.ComponentScript = ShowTab(LoginTabTagID)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\">Sign in</a></div></h1><div class=\"required field\"><label>Email</label><div class=\"ui inverted transparent left icon input\"><i class=\"envelope icon\"></i> <input type=\"text\" placeholder=\"Email\" name=\"Email\"></div></div><div class=\"required field\"><label>Username</label><div class=\"ui inverted transparent left icon input\"><i class=\"user icon\"></i> <input type=\"text\" placeholder=\"Username\" name=\"Username\"></div></div><div class=\"required field\"><label>Password</label><div class=\"ui inverted transparent left icon input\"><i class=\"lock icon\"></i> <input type=\"password\" placeholder=\"Password\" name=\"Password\"></div></div><div class=\"required field\"><label>Confirm Password</label><div class=\"ui inverted transparent left icon input\"><i class=\"redo icon\"></i> <input type=\"password\" placeholder=\"Confirm Password\" name=\"ConfirmPassword\"></div></div><div class=\"ui divider\"></div><div class=\"ui middle aligned center aligned grid\"><div class=\"column\"><div class=\"inline required field\"><div class=\"ui inverted checkbox\"><input type=\"checkbox\" name=\"Terms\"> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, ShowModal(TermsModalTagID))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<label style=\"cursor: pointer;\" onclick=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 templ.ComponentScript = ShowModal(TermsModalTagID)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\">I agree to the <a><b>Terms and Conditions</b></a></label></div></div><button id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(RegisterSubmitTagID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/app/server/component/page/home/form.register.templ`, Line: 158, Col: 36}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\" class=\"ui fluid primary submit button\">Sign up</button><div class=\"ui inverted basic popup\"><div class=\"header\">Verify Captcha</div><div class=\"ui divider\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CaptchaForm(captcha).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "</div></div></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = RegisterFormInit(RegisterSubmitTagID, RegisterFormTagID, LoginTabTagID).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
