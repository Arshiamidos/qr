package utils

import (
	"fmt"
	"strconv"
	"regexp"
)

//https://www.thonky.com/qr-code-tutorial/data-encoding
var Mode = map[rune]string{
	'N': "0001",
	'A': "0010",
	'B': "0100",
	'K': "1000",
}
func AlignmentPatternLocation(v int) []int{
	return map[int][]int{
		2	:{6	,18	,0	,0	,0		,0		,0	},
		3	:{6	,22	,0	,0	,0		,0		,0	},
		4	:{6	,26	,0	,0	,0		,0		,0	},
		5	:{6	,30	,0	,0	,0		,0		,0	},
		6	:{6	,34	,0	,0	,0		,0		,0	},
		7	:{6	,22	,38	,0	,0		,0		,0	},
		8	:{6	,24	,42	,0	,0		,0		,0	},
		9	:{6	,26	,46	,0	,0		,0		,0	},
		10	:{6	,28	,50	,0	,0		,0		,0	},
		11	:{6	,30	,54	,0	,0		,0		,0	},
		12	:{6	,32	,58	,0	,0		,0		,0	},
		13	:{6	,34	,62	,0	,0		,0		,0	},
		14	:{6	,26	,46	,66	,0		,0		,0	},
		15	:{6	,26	,48	,70	,0		,0		,0	},
		16	:{6	,26	,50	,74	,0		,0		,0	},
		17	:{6	,30	,54	,78	,0		,0		,0	},
		18	:{6	,30	,56	,82	,0		,0		,0	},
		19	:{6	,30	,58	,86	,0		,0		,0	},
		20	:{6	,34	,62	,90	,0		,0		,0	},
		21	:{6	,28	,50	,72	,94		,0		,0	},
		22	:{6	,26	,50	,74	,98		,0		,0	},
		23	:{6	,30	,54	,78	,102	,0		,0	},
		24	:{6	,28	,54	,80	,106	,0		,0	},
		25	:{6	,32	,58	,84	,110	,0		,0	},
		26	:{6	,30	,58	,86	,114	,0		,0	},
		27	:{6	,34	,62	,90	,118	,0		,0	},
		28	:{6	,26	,50	,74	,98		,122	,0	},
		29	:{6	,30	,54	,78	,102	,126	,0	},
		30	:{6	,26	,52	,78	,104	,130	,0	},
		31	:{6	,30	,56	,82	,108	,134	,0	},
		32	:{6	,34	,60	,86	,112	,138	,0	},
		33	:{6	,30	,58	,86	,114	,142	,0	},
		34	:{6	,34	,62	,90	,118	,146	,0	},
		35	:{6	,30	,54	,78	,102	,126	,150},
		36	:{6	,24	,50	,76	,102	,128	,154},
		37	:{6	,28	,54	,80	,106	,132	,158},
		38	:{6	,32	,58	,84	,110	,136	,162},
		39	:{6	,26	,54	,82	,110	,138	,166},
		40	:{6	,30	,58	,86	,114	,142	,170},
	}[v]
}

func ReminderBits(v int) int{
	return map[int]int{
		1	:0,
		2	:7,
		3	:7,
		4	:7,
		5	:7,
		6	:7,
		7	:0,
		8	:0,
		9	:0,
		10	:0,
		11	:0,
		12	:0,
		13	:0,
		14	:3,
		15	:3,
		16	:3,
		17	:3,
		18	:3,
		19	:3,
		20	:3,
		21	:4,
		22	:4,
		23	:4,
		24	:4,
		25	:4,
		26	:4,
		27	:4,
		28	:3,
		29	:3,
		30	:3,
		31	:3,
		32	:3,
		33	:3,
		34	:3,
		35	:0,
		36	:0,
		37	:0,
		38	:0,
		39	:0,
		40	:0,
	}[v]
}
func CharacterCountIndicator(v int, t rune) int {
	if v >= 1 && v <= 9 {
		return map[rune]int{
			'N': 10,
			'A': 9,
			'B': 8,
			'K': 8,
		}[t]
	} else if v >= 10 && v <= 26 {
		return map[rune]int{
			'N': 12,
			'A': 11,
			'B': 16,
			'K': 10,
		}[t]
	} else if v >= 27 && v <= 40 {
		return map[rune]int{
			'N': 14,
			'A': 13,
			'B': 16,
			'K': 12,
		}[t]
	} else {
		panic("not found matching version vs type")
	}
}
func AlphaNumericValue(c byte) int {
	return map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'A': 10,
		'B': 11,
		'C': 12,
		'D': 13,
		'E': 14,
		'F': 15,
		'G': 16,
		'H': 17,
		'I': 18,
		'J': 19,
		'K': 20,
		'L': 21,
		'M': 22,
		'N': 23,
		'O': 24,
		'P': 25,
		'Q': 26,
		'R': 27,
		'S': 28,
		'T': 29,
		'U': 30,
		'V': 31,
		'W': 32,
		'X': 33,
		'Y': 34,
		'Z': 35,
		' ': 36,
		'$': 37,
		'%': 38,
		'*': 39,
		'+': 40,
		'-': 41,
		'.': 42,
		'/': 43,
		':': 44,
	}[c]
}
func QrSize(v int) int{
	return ((v-1)*4)+21
}
func BreakUp8Bit(orgi string, s string, v int, e rune, t rune) string {

	s = Mode[t] + fmt.Sprintf("%0*b", CharacterCountIndicator(v, t), len(orgi)) + s
	//will be reach capacity with <=4 zeros?so add it add <=4 zeros
	//not?add 0000
	//still not ? reach multipication of 8 by adding 0 nex add  11101100 00010001
	eccw := GetECCWBytesNeed(v, e)
	if len(s)+4 >= eccw {
		rem := eccw - len(s)
		for i := 0; i < rem; i++ {
			s = s + "0"
		}
		return s
	} else {
		s = s + "0000"
		if len(s)%8 != 0 {
			til8 := 8 - len(s)%8
			for i := 0; i < til8; i++ {
				s = s + "0"
			}
		}
		if len(s) != eccw {
			tilLast := (eccw - len(s)) / 8
			for i := 0; i < tilLast; i++ {
				//11101100 00010001
				if i%2 == 0 {
					s = s + "11101100"
				} else {
					s = s + "00010001"
				}
			}
		}
	}
	if len(s) > eccw {
		panic("wrong size of bitcodes")
	}
	return s
}
func ConvertToMessagePoly(s string) map[int]int{
	r, _ := regexp.Compile("........")
	datas := r.FindAllString(s, -1)
	ints:=[]int{}
	for i := 0; i < len(datas); i++ {
		d,_:=strconv.ParseInt(datas[i],2,16)
		ints=append(ints,int(d))
	}
	poly:=map[int]int{}

	for i := len(ints)-1; i >=0 ; i-- {
		poly[len(ints)-1-i]=ints[i]
	}
	return poly
}
