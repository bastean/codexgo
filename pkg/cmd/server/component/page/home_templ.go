// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package page

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Home() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex h-screen w-full items-center justify-center\"><div role=\"tablist\" class=\"tabs tabs-lifted\"><input type=\"radio\" name=\"tabs\" role=\"tab\" class=\"tab\" aria-label=\"Register\"><div role=\"tabpanel\" class=\"tab-content bg-base-100 border-base-300 rounded-box p-6\"><form class=\"flex h-full w-full flex-col items-center justify-center gap-5\" hx-put=\"/\" hx-target-*=\"#alert\"><input name=\"email\" type=\"email\" placeholder=\"Email\" class=\"input input-bordered w-full max-w-xs\"> <input name=\"username\" type=\"text\" placeholder=\"Username\" class=\"input input-bordered w-full max-w-xs\"> <input name=\"password\" type=\"password\" placeholder=\"Password\" class=\"input input-bordered w-full max-w-xs\"> <button class=\"btn btn-primary\">Register</button></form></div><input type=\"radio\" name=\"tabs\" role=\"tab\" class=\"tab\" aria-label=\"Login\" checked><div role=\"tabpanel\" class=\"tab-content bg-base-100 border-base-300 rounded-box p-6\"><form class=\"flex h-full w-full flex-col items-center justify-center gap-5\" hx-post=\"/\" hx-target=\"#main\" hx-target-error=\"#alert\"><input name=\"email\" type=\"email\" placeholder=\"Email\" class=\"input input-bordered w-full max-w-xs\"> <input name=\"password\" type=\"password\" placeholder=\"Password\" class=\"input input-bordered w-full max-w-xs\"> <button class=\"btn btn-primary\">Login</button></form></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}