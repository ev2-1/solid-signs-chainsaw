package main

import (
	"github.com/anon55555/mt"

	"fmt"
	"image/color"
	"strings"
)

func uint16String(n uint16) string {
	return fmt.Sprintf("%d", n)
}

func boolString(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func boxString(b mt.Box) string {
	return " [0] " + vecString(b[0]) + "\n [1] " + vecString(b[1])
}

func vecString(v mt.Vec) string {
	return fmt.Sprintf("%f %f %f", v[0], v[1], v[2])
}

func strString(s string) string {
	return "'" + s + "'"
}

func textureString(ts []mt.Texture) (s string) {
	for _, t := range ts {
		s += "\n " + strString(string(t)) + ""
	}

	return
}

func float32String(f float32) string {
	return fmt.Sprintf("%f", f)
}

func int16StringS2(is [2]int16) string {
	return fmt.Sprintf(strings.Repeat("%d ", len(is)), is[0], is[1])
}

func colorStringS(cs []color.NRGBA) (s string) {
	for _, c := range cs {
		s += colorString(c) + "; "
	}
	return
}

func colorString(c color.NRGBA) string {
	return fmt.Sprintf("#%d %d %d %d", c.R, c.G, c.B, c.A)
}

func int8String(i int8) string {
	return fmt.Sprintf("%d", i)
}

func AOString(p mt.AOProps) (s string) {
	s += "MaxHP                " + uint16String(p.MaxHP) + "\n"
	s += "CollideWithNodes     " + boolString(p.CollideWithNodes) + "\n"
	s += "ColBox:\n" + boxString(p.ColBox) + "\n"
	s += "SelBox:\n" + boxString(p.SelBox) + "\n"
	s += "Pointable            " + boolString(p.Pointable) + "\n"
	s += "Visual               " + strString(p.Visual) + "\n"
	s += "VisualSize           " + vecString(p.VisualSize) + "\n"
	s += "Textures:" + textureString(p.Textures) + "\n"
	s += "UseTextureAlpha      " + boolString(p.UseTextureAlpha) + "\n"
	s += "DmgTextureMod  (sufx)" + strString(string(p.DmgTextureMod)) + "\n"
	s += "Glow                 " + int8String(p.Glow) + "\n"
	s += "Shaded               " + boolString(p.Shaded) + "\n"
	s += "SpriteSheetSize      " + int16StringS2(p.SpriteSheetSize) + "\n"
	s += "SpritePos            " + int16StringS2(p.SpritePos) + "\n"
	s += "Visible              " + boolString(p.Visible) + "\n"
	s += "MakeFootstepSnds     " + boolString(p.MakeFootstepSnds) + "\n"
	s += "RotateSpeed (r/s )   " + float32String(p.RotateSpeed) + "\n"
	s += "Mesh                 " + strString(p.Mesh) + "\n"
	s += "Colors               " + colorStringS(p.Colors) + "\n"
	s += "CollideWithAOs       " + boolString(p.CollideWithAOs) + "\n"
	s += "StepHeight           " + float32String(p.StepHeight) + "\n"
	s += "FaceRotateDir        " + boolString(p.FaceRotateDir) + "\n"
	s += "FaceRotateDirOff     " + float32String(p.FaceRotateDirOff) + "\n"
	s += "BackfaceCull         " + boolString(p.BackfaceCull) + "\n"
	s += "Nametag              " + strString(p.Nametag) + "\n"
	s += "NametagColor         " + colorString(p.NametagColor) + "\n"
	s += "NametagBG            " + colorString(p.NametagBG) + "\n"
	s += "FaceRotateSpeed (°/s)" + float32String(p.FaceRotateSpeed) + "\n"
	s += "Infotext             " + strString(p.Infotext) + "\n"
	s += "Itemstring           " + strString(p.Itemstring) + "\n"
	s += "ShowOnMinimap        " + boolString(p.ShowOnMinimap) + "\n"

	s += "(p) MaxBreath        " + uint16String(p.MaxBreath) + "\n"
	s += "(p) EyeHeight        " + float32String(p.EyeHeight) + "\n"
	s += "(p) ZoomFOV       (°)" + float32String(p.ZoomFOV) + "\n"
	s += "Weight (depricated)  " + float32String(p.Weight)
	return
}

func armorGroupsString(ags []mt.Group) (s string) {
	for _, g := range ags {
		s += fmt.Sprintf("%s:%d", g.Name, g.Rating)
	}

	return
}

func attachString(a mt.AOAttach) string {
	return fmt.Sprintf("%d b:%s pos: %s, rot: %s; visible %t", a.ParentID, a.Bone, vecString(a.Pos), vecString(a.Rot), a.ForceVisible)
}
