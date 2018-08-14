<<<<<<< HEAD
// +build codecgen.exec

=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
// Copyright (c) 2012-2015 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

package codec

// DO NOT EDIT. THIS FILE IS AUTO-GENERATED FROM gen-dec-(map|array).go.tmpl

const genDecMapTmpl = `
{{var "v"}} := *{{ .Varname }}
{{var "l"}} := r.ReadMapStart()
{{var "bh"}} := z.DecBasicHandle()
if {{var "v"}} == nil {
<<<<<<< HEAD
	{{var "rl"}} := z.DecInferLen({{var "l"}}, {{var "bh"}}.MaxInitLen, {{ .Size }})
=======
	{{var "rl"}}, _ := z.DecInferLen({{var "l"}}, {{var "bh"}}.MaxInitLen, {{ .Size }})
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	{{var "v"}} = make(map[{{ .KTyp }}]{{ .Typ }}, {{var "rl"}})
	*{{ .Varname }} = {{var "v"}}
}
var {{var "mk"}} {{ .KTyp }}
var {{var "mv"}} {{ .Typ }}
<<<<<<< HEAD
var {{var "mg"}}, {{var "mdn"}} {{if decElemKindPtr}}, {{var "ms"}}, {{var "mok"}}{{end}} bool
=======
var {{var "mg"}} {{if decElemKindPtr}}, {{var "ms"}}, {{var "mok"}}{{end}} bool
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
if {{var "bh"}}.MapValueReset {
	{{if decElemKindPtr}}{{var "mg"}} = true
	{{else if decElemKindIntf}}if !{{var "bh"}}.InterfaceReset { {{var "mg"}} = true }
	{{else if not decElemKindImmutable}}{{var "mg"}} = true
	{{end}} }
<<<<<<< HEAD
if {{var "l"}} != 0 {
{{var "hl"}} := {{var "l"}} > 0 
	for {{var "j"}} := 0; ({{var "hl"}} && {{var "j"}} < {{var "l"}}) || !({{var "hl"}} || r.CheckBreak()); {{var "j"}}++ {
	r.ReadMapElemKey() {{/* z.DecSendContainerState(codecSelfer_containerMapKey{{ .Sfx }}) */}}
=======
if {{var "l"}} > 0  {
for {{var "j"}} := 0; {{var "j"}} < {{var "l"}}; {{var "j"}}++ {
	z.DecSendContainerState(codecSelfer_containerMapKey{{ .Sfx }})
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	{{ $x := printf "%vmk%v" .TempVar .Rand }}{{ decLineVarK $x }}
{{ if eq .KTyp "interface{}" }}{{/* // special case if a byte array. */}}if {{var "bv"}}, {{var "bok"}} := {{var "mk"}}.([]byte); {{var "bok"}} {
		{{var "mk"}} = string({{var "bv"}})
	}{{ end }}{{if decElemKindPtr}}
	{{var "ms"}} = true{{end}}
	if {{var "mg"}} {
		{{if decElemKindPtr}}{{var "mv"}}, {{var "mok"}} = {{var "v"}}[{{var "mk"}}] 
		if {{var "mok"}} {
			{{var "ms"}} = false
		} {{else}}{{var "mv"}} = {{var "v"}}[{{var "mk"}}] {{end}}
	} {{if not decElemKindImmutable}}else { {{var "mv"}} = {{decElemZero}} }{{end}}
<<<<<<< HEAD
	r.ReadMapElemValue() {{/* z.DecSendContainerState(codecSelfer_containerMapValue{{ .Sfx }}) */}}
	{{var "mdn"}} = false
	{{ $x := printf "%vmv%v" .TempVar .Rand }}{{ $y := printf "%vmdn%v" .TempVar .Rand }}{{ decLineVar $x $y }}
	if {{var "mdn"}} {
		if {{ var "bh" }}.DeleteOnNilMapValue { delete({{var "v"}}, {{var "mk"}}) } else { {{var "v"}}[{{var "mk"}}] = {{decElemZero}} }
	} else if {{if decElemKindPtr}} {{var "ms"}} && {{end}} {{var "v"}} != nil {
=======
	z.DecSendContainerState(codecSelfer_containerMapValue{{ .Sfx }})
	{{ $x := printf "%vmv%v" .TempVar .Rand }}{{ decLineVar $x }}
	if {{if decElemKindPtr}} {{var "ms"}} && {{end}} {{var "v"}} != nil {
		{{var "v"}}[{{var "mk"}}] = {{var "mv"}}
	}
}
} else if {{var "l"}} < 0  {
for {{var "j"}} := 0; !r.CheckBreak(); {{var "j"}}++ {
	z.DecSendContainerState(codecSelfer_containerMapKey{{ .Sfx }})
	{{ $x := printf "%vmk%v" .TempVar .Rand }}{{ decLineVarK $x }}
{{ if eq .KTyp "interface{}" }}{{/* // special case if a byte array. */}}if {{var "bv"}}, {{var "bok"}} := {{var "mk"}}.([]byte); {{var "bok"}} {
		{{var "mk"}} = string({{var "bv"}})
	}{{ end }}{{if decElemKindPtr}}
	{{var "ms"}} = true {{ end }}
	if {{var "mg"}} {
		{{if decElemKindPtr}}{{var "mv"}}, {{var "mok"}} = {{var "v"}}[{{var "mk"}}] 
		if {{var "mok"}} {
			{{var "ms"}} = false
		} {{else}}{{var "mv"}} = {{var "v"}}[{{var "mk"}}] {{end}}
	} {{if not decElemKindImmutable}}else { {{var "mv"}} = {{decElemZero}} }{{end}}
	z.DecSendContainerState(codecSelfer_containerMapValue{{ .Sfx }})
	{{ $x := printf "%vmv%v" .TempVar .Rand }}{{ decLineVar $x }}
	if {{if decElemKindPtr}} {{var "ms"}} && {{end}} {{var "v"}} != nil {
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
		{{var "v"}}[{{var "mk"}}] = {{var "mv"}}
	}
}
} // else len==0: TODO: Should we clear map entries?
<<<<<<< HEAD
r.ReadMapEnd() {{/* z.DecSendContainerState(codecSelfer_containerMapEnd{{ .Sfx }}) */}}
=======
z.DecSendContainerState(codecSelfer_containerMapEnd{{ .Sfx }})
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
`

const genDecListTmpl = `
{{var "v"}} := {{if not isArray}}*{{end}}{{ .Varname }}
{{var "h"}}, {{var "l"}} := z.DecSliceHelperStart() {{/* // helper, containerLenS */}}{{if not isArray}}
var {{var "c"}} bool {{/* // changed */}}
_ = {{var "c"}}{{end}}
if {{var "l"}} == 0 {
	{{if isSlice }}if {{var "v"}} == nil {
		{{var "v"}} = []{{ .Typ }}{}
		{{var "c"}} = true
	} else if len({{var "v"}}) != 0 {
		{{var "v"}} = {{var "v"}}[:0]
		{{var "c"}} = true
<<<<<<< HEAD
	} {{else if isChan }}if {{var "v"}} == nil {
		{{var "v"}} = make({{ .CTyp }}, 0)
		{{var "c"}} = true
	} {{end}}
} else {
	{{var "hl"}} := {{var "l"}} > 0
	var {{var "rl"}} int
	_ =  {{var "rl"}}
	{{if isSlice }} if {{var "hl"}} {
	if {{var "l"}} > cap({{var "v"}}) {
		{{var "rl"}} = z.DecInferLen({{var "l"}}, z.DecBasicHandle().MaxInitLen, {{ .Size }})
		if {{var "rl"}} <= cap({{var "v"}}) {
			{{var "v"}} = {{var "v"}}[:{{var "rl"}}]
=======
	} {{end}} {{if isChan }}if {{var "v"}} == nil {
		{{var "v"}} = make({{ .CTyp }}, 0)
		{{var "c"}} = true
	} {{end}}
} else if {{var "l"}} > 0 {
	{{if isChan }}if {{var "v"}} == nil {
		{{var "rl"}}, _ = z.DecInferLen({{var "l"}}, z.DecBasicHandle().MaxInitLen, {{ .Size }})
		{{var "v"}} = make({{ .CTyp }}, {{var "rl"}})
		{{var "c"}} = true
	}
	for {{var "r"}} := 0; {{var "r"}} < {{var "l"}}; {{var "r"}}++ {
		{{var "h"}}.ElemContainerState({{var "r"}})
		var {{var "t"}} {{ .Typ }}
		{{ $x := printf "%st%s" .TempVar .Rand }}{{ decLineVar $x }}
		{{var "v"}} <- {{var "t"}}
	}
	{{ else }}	var {{var "rr"}}, {{var "rl"}} int {{/* // num2read, length of slice/array/chan */}}
	var {{var "rt"}} bool {{/* truncated */}}
	_, _ = {{var "rl"}}, {{var "rt"}}
	{{var "rr"}} = {{var "l"}} // len({{var "v"}})
	if {{var "l"}} > cap({{var "v"}}) {
		{{if isArray }}z.DecArrayCannotExpand(len({{var "v"}}), {{var "l"}})
		{{ else }}{{if not .Immutable }}
		{{var "rg"}} := len({{var "v"}}) > 0
		{{var "v2"}} := {{var "v"}} {{end}}
		{{var "rl"}}, {{var "rt"}} = z.DecInferLen({{var "l"}}, z.DecBasicHandle().MaxInitLen, {{ .Size }})
		if {{var "rt"}} {
			if {{var "rl"}} <= cap({{var "v"}}) {
				{{var "v"}} = {{var "v"}}[:{{var "rl"}}]
			} else {
				{{var "v"}} = make([]{{ .Typ }}, {{var "rl"}})
			}
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
		} else {
			{{var "v"}} = make([]{{ .Typ }}, {{var "rl"}})
		}
		{{var "c"}} = true
<<<<<<< HEAD
	} else if {{var "l"}} != len({{var "v"}}) {
		{{var "v"}} = {{var "v"}}[:{{var "l"}}]
		{{var "c"}} = true
	}
	} {{end}}
	var {{var "j"}} int 
    // var {{var "dn"}} bool 
	for ; ({{var "hl"}} && {{var "j"}} < {{var "l"}}) || !({{var "hl"}} || r.CheckBreak()); {{var "j"}}++ {
		{{if not isArray}} if {{var "j"}} == 0 && {{var "v"}} == nil {
			if {{var "hl"}} {
				{{var "rl"}} = z.DecInferLen({{var "l"}}, z.DecBasicHandle().MaxInitLen, {{ .Size }})
			} else {
				{{var "rl"}} = {{if isSlice}}8{{else if isChan}}64{{end}}
			}
			{{var "v"}} = make({{if isSlice}}[]{{ .Typ }}{{else if isChan}}{{.CTyp}}{{end}}, {{var "rl"}})
			{{var "c"}} = true 
		}{{end}}
		{{var "h"}}.ElemContainerState({{var "j"}})
        {{/* {{var "dn"}} = r.TryDecodeAsNil() */}}{{/* commented out, as decLineVar handles this already each time */}}
        {{if isChan}}{{ $x := printf "%[1]vvcx%[2]v" .TempVar .Rand }}var {{$x}} {{ .Typ }}
		{{ decLineVar $x }}
		{{var "v"}} <- {{ $x }}
        // println(">>>> sending ", {{ $x }}, " into ", {{var "v"}}) // TODO: remove this
        {{else}}{{/* // if indefinite, etc, then expand the slice if necessary */}}
		var {{var "db"}} bool
		if {{var "j"}} >= len({{var "v"}}) {
			{{if isSlice }} {{var "v"}} = append({{var "v"}}, {{ zero }})
			{{var "c"}} = true
			{{else}} z.DecArrayCannotExpand(len(v), {{var "j"}}+1); {{var "db"}} = true
			{{end}}
		}
		if {{var "db"}} {
			z.DecSwallow()
		} else {
			{{ $x := printf "%[1]vv%[2]v[%[1]vj%[2]v]" .TempVar .Rand }}{{ decLineVar $x }}
		}
        {{end}}
	}
	{{if isSlice}} if {{var "j"}} < len({{var "v"}}) {
		{{var "v"}} = {{var "v"}}[:{{var "j"}}]
		{{var "c"}} = true
	} else if {{var "j"}} == 0 && {{var "v"}} == nil {
		{{var "v"}} = make([]{{ .Typ }}, 0)
		{{var "c"}} = true
	} {{end}}
=======
		{{var "rr"}} = len({{var "v"}}) {{if not .Immutable }}
			if {{var "rg"}} { copy({{var "v"}}, {{var "v2"}}) } {{end}} {{end}}{{/* end not Immutable, isArray */}}
	} {{if isSlice }} else if {{var "l"}} != len({{var "v"}}) {
		{{var "v"}} = {{var "v"}}[:{{var "l"}}]
		{{var "c"}} = true
	} {{end}}	{{/* end isSlice:47 */}} 
	{{var "j"}} := 0
	for ; {{var "j"}} < {{var "rr"}} ; {{var "j"}}++ {
		{{var "h"}}.ElemContainerState({{var "j"}})
		{{ $x := printf "%[1]vv%[2]v[%[1]vj%[2]v]" .TempVar .Rand }}{{ decLineVar $x }}
	}
	{{if isArray }}for ; {{var "j"}} < {{var "l"}} ; {{var "j"}}++ {
		{{var "h"}}.ElemContainerState({{var "j"}})
		z.DecSwallow()
	}
	{{ else }}if {{var "rt"}} {
		for ; {{var "j"}} < {{var "l"}} ; {{var "j"}}++ {
			{{var "v"}} = append({{var "v"}}, {{ zero}})
			{{var "h"}}.ElemContainerState({{var "j"}})
			{{ $x := printf "%[1]vv%[2]v[%[1]vj%[2]v]" .TempVar .Rand }}{{ decLineVar $x }}
		}
	} {{end}} {{/* end isArray:56 */}}
	{{end}} {{/* end isChan:16 */}}
} else { {{/* len < 0 */}}
	{{var "j"}} := 0
	for ; !r.CheckBreak(); {{var "j"}}++ {
		{{if isChan }}
		{{var "h"}}.ElemContainerState({{var "j"}})
		var {{var "t"}} {{ .Typ }}
		{{ $x := printf "%st%s" .TempVar .Rand }}{{ decLineVar $x }}
		{{var "v"}} <- {{var "t"}} 
		{{ else }}
		if {{var "j"}} >= len({{var "v"}}) {
			{{if isArray }}z.DecArrayCannotExpand(len({{var "v"}}), {{var "j"}}+1)
			{{ else }}{{var "v"}} = append({{var "v"}}, {{zero}})// var {{var "z"}} {{ .Typ }}
			{{var "c"}} = true {{end}}
		}
		{{var "h"}}.ElemContainerState({{var "j"}})
		if {{var "j"}} < len({{var "v"}}) {
			{{ $x := printf "%[1]vv%[2]v[%[1]vj%[2]v]" .TempVar .Rand }}{{ decLineVar $x }}
		} else {
			z.DecSwallow()
		}
		{{end}}
	}
	{{if isSlice }}if {{var "j"}} < len({{var "v"}}) {
		{{var "v"}} = {{var "v"}}[:{{var "j"}}]
		{{var "c"}} = true
	} else if {{var "j"}} == 0 && {{var "v"}} == nil {
		{{var "v"}} = []{{ .Typ }}{}
		{{var "c"}} = true
	}{{end}}
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
}
{{var "h"}}.End()
{{if not isArray }}if {{var "c"}} { 
	*{{ .Varname }} = {{var "v"}}
}{{end}}
`

<<<<<<< HEAD
const genEncChanTmpl = `
{{.Label}}:
switch timeout{{.Sfx}} :=  z.EncBasicHandle().ChanRecvTimeout; {
case timeout{{.Sfx}} == 0: // only consume available
	for {
		select {
		case b{{.Sfx}} := <-{{.Chan}}:
			{{ .Slice }} = append({{.Slice}}, b{{.Sfx}})
		default:
			break {{.Label}}
		}
	}
case timeout{{.Sfx}} > 0: // consume until timeout
	tt{{.Sfx}} := time.NewTimer(timeout{{.Sfx}})
	for {
		select {
		case b{{.Sfx}} := <-{{.Chan}}:
			{{.Slice}} = append({{.Slice}}, b{{.Sfx}})
		case <-tt{{.Sfx}}.C:
			// close(tt.C)
			break {{.Label}}
		}
	}
default: // consume until close
	for b{{.Sfx}} := range {{.Chan}} {
		{{.Slice}} = append({{.Slice}}, b{{.Sfx}})
	}
}
`
=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
