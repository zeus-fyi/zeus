package filepaths

import (
	"path"

	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

type Path struct {
	PackageName string
	DirIn       string
	DirOut      string
	FnIn        string
	FnOut       string
	Env         string
	FilterFiles strings_filter.FilterOpts
}

type Paths struct {
	Slice []Path
}

func (ps *Paths) AddPathToSlice(p Path) {
	ps.Slice = append(ps.Slice, p)
}

func (p *Path) FileOutPath() string {
	return path.Join(p.DirOut, p.FnOut)
}

func (p *Path) FileInPath() string {
	return path.Join(p.DirIn, p.FnIn)
}

func (p *Path) LeftExtendDirInPath(dirExtend string) string {
	p.DirIn = path.Join(dirExtend, p.DirIn)
	return p.DirIn
}

func (p *Path) RightExtendDirInPath(dirExtend string) string {
	p.DirIn = path.Join(p.DirIn, dirExtend)
	return p.DirIn
}

func (p *Path) LeftExtendDirOutPath(dirExtend string) string {
	p.DirOut = path.Join(dirExtend, p.DirOut)
	return p.DirOut
}

func (p *Path) RightExtendDirOutPath(dirExtend string) string {
	p.DirOut = path.Join(p.DirOut, dirExtend)
	return p.DirOut
}
