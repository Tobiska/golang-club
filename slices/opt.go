package main

type A struct {
	ss string
}

func toInf(strArr []string) []*A {
	aArr := make([]*A, len(strArr))
	aArrVal := make([]A, len(strArr))
	for i := 0; i < len(strArr); i++ {
		aArrVal[i].ss = strArr[i]
		aArr[i] = &aArrVal[i]
	}
	return aArr
}

func toInf2(strArr []string) []*A {
	aArr := make([]*A, 0, len(strArr))
	for s := range strArr {
		aArr = append(aArr, &A{
			ss: string(s),
		})
	}
	return aArr
}

func main34() {
	sl := make([]int)
}
