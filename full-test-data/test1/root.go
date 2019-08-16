package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"

import "log"

type Root struct {
	Something int
	Success   bool
}

func (c *Root) OnClickRun(event *vugu.DOMEvent, n int) {
	c.Success = true
	log.Printf("HEY, GOT HERE!")
}
func (c *Root) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut, vgreterr error) {

	vgout = &vugu.BuildOut{}

	var vgn *vugu.VGNode
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "span", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "test-span"}, vugu.VGAttribute{Namespace: "", Key: "id", Val: "test_span_id"}}}
	vgout.Doc = vgn // Doc root for output
	{
		vgparent := vgn
		_ = vgparent
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    This is a test.\n    ", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "id", Val: "run"}}}
		vgparent.AppendChild(vgn)
		vgn.Attr = append(vgn.Attr, vugu.VGAttribute{Key: "data-whatever", Val: fmt.Sprint(c.Something)})
		vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
			EventType: "click",
			Func:      func(event *vugu.DOMEvent) { c.OnClickRun(event, 7) },
			// TODO: implement capture, etc.
		})
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "run", Attr: []vugu.VGAttribute(nil)}
			vgparent.AppendChild(vgn)
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    ", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
		if c.Success {
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "id", Val: "success"}}}
			vgparent.AppendChild(vgn)
		}
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "success", Attr: []vugu.VGAttribute(nil)}
			vgparent.AppendChild(vgn)
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    ", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
		{
			vghtml := "Some <strong>content</strong> here"
			vgn.InnerHTML = &vghtml
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
	}
	return vgout, nil
}

// 'fix' unused imports
var _ = fmt.Sprintf
var _ = reflect.New
var _ = js.ValueOf
