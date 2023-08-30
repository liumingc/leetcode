package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	eN := len(equations)
	varMap := make(map[string]*vari)
	for i := 0; i < eN; i++ {
		a, b := equations[i][0], equations[i][1]

		va := varMap[a]
		if va == nil {
			va = &vari{
				sym:  a,
				neib: make(map[string]*neighbor),
			}
			varMap[a] = va
		}
		vb := varMap[b]
		if vb == nil {
			vb = &vari{
				sym:  b,
				neib: make(map[string]*neighbor),
			}
			varMap[b] = vb
		}
		va.neib[b] = &neighbor{val: values[i], x: vb}
		vb.neib[a] = &neighbor{val: 1 / values[i], x: va}
	}

	printVarMap(varMap)

	qN := len(queries)
	res := make([]float64, qN)
	for i := 0; i < qN; i++ {
		x := -1.0
		c, d := queries[i][0], queries[i][1]
		vc := varMap[c]
		vd := varMap[d]
		if vc != nil && vd != nil {
			// is there a path from vc to vd?
			neibs := map[string]*neighbor{vc.sym: {val: 1.0, x: vc}}

		Louter:
			for len(neibs) > 0 {
				nextNeibs := make(map[string]*neighbor)
				for _, nb := range neibs {
					val := nb.val
					if nb.x.sym == vd.sym { // found
						x = val
						break Louter
					}
					if nb.x.q == i+1 { // visited before
						continue
					}
					nb.x.q = i + 1
					for _, nnb := range nb.x.neib {
						nextNeibs[nnb.x.sym] = &neighbor{
							val: val * nnb.val,
							x:   nnb.x,
						}
					}
				}
				neibs = nextNeibs
			}

		}

		res[i] = x
	}
	return res
}

type vari struct {
	sym  string
	neib map[string]*neighbor
	q    int
}

type neighbor struct {
	val float64
	x   *vari
}
