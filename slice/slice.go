package slice

// StrSlcRm removes string s from *s in place
func StrSlcRm(slc *[]string, s string) {
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

// StrSlcEquals tests whether the two string slices are
// equal by value
func StrSlcEquals(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, w := range s1 {
		if w != s2[i] {
			return false
		}
	}
	return true
}
