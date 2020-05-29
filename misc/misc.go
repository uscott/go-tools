package misc

func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}
	return cp
}

func SlcRmStr(slc *[]string, s string) {
	n, nrm := len(*slc), 0
	for i := 0; i < n; i++ {
		j := i - nrm
		t := (*slc)[j]
		if t == s {
			nrm++
			switch {
			case j == 0:
				*slc = (*slc)[1:]
			case j == len(*slc)-1:
				*slc = (*slc)[:len(*slc)-1]
			default:
				*slc = append((*slc)[:j], (*slc)[j+1:]...)
			}
		}
	}
}
