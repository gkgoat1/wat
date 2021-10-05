package src

import (
	"fmt"
	"math/rand"
)

type Ir struct {
	Tymap map[string]int
	Opmap map[string]map[int]int
}

func (i *Ir) Init(tys []string) {
	for _, x := range tys {
		i.Tymap[x] = rand.Int()
		i.Opmap[x] = map[int]int{}
		for _, y := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15} {
			i.Opmap[x][y] = rand.Int()
		}
	}
}
func (i *Ir) Gen(stack int, prog func(i *Ir) string) string {
	x := "int done = 0;"
	x += fmt.Sprintf("void *r1;int** prog{%s};int pidx = 0;void* stack[%d];int sidx = 0;while(!done && (pidx++))switch(prog[pidx][0]){", prog(i), stack)
	for t, ty := range i.Tymap {
		switch t {
		case "push":
			x += fmt.Sprintf("case %d: stack[sidx++] = prog[pidx][%d];break", ty, i.Opmap[t][0])
		case "dup":
			x += fmt.Sprintf("case %d: r1 = stack[sidx]; stack[sidx++] = r1; break;", ty)
		case "rot":
			x += fmt.Sprintf("case %d: int n = stack[sidx--];int i;r1 = stack[n];for(i = n; i < sidx;i++)stack[i]=stack[i+1];stack[sidx] = r1;break;", ty)

		default:
			x += fmt.Sprintf("case %d: %s(...stack);break;", ty, t)
		}
	}
	x += "}"
	return x
}
