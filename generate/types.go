package generate

type GoPackages struct {
	Pkgs         []*goPkg2
	MaxPkgLen    int
	MaxRelDirLen int
}
type ProtocIt struct {
	goPkgs GoPackages
	//	Pkgs   []*goPkg3
}

type goPkg2 struct {
	RelDir  string
	Pkg     string
	Files   []string
	Imports []string
	// pkgs    []*goPkg2
}

// type goPkg3 struct {
// 	goPkg2
// 	absOut     string
// 	outBit     string
// 	dirtyFiles map[string][]string // generate path -> []files
// 	imps       []*goPkg3
// }

// type goPkg3By []*goPkg3

// func (a goPkg3By) Len() int      { return len(a) }
// func (a goPkg3By) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
// func (a goPkg3By) Less(i, j int) bool {
// 	x, y := a[i], a[j]
// 	for k, _ := range x.imps {
// 		if y.goPkg2.Pkg == x.imps[k].goPkg2.Pkg {
// 			return false
// 		}
// 	}
// 	return true
// }
