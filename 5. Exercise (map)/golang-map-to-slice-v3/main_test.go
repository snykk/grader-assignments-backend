package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TestData struct {
	Test     map[string]string
	Expected [][]string
}

type TestDataSlice []TestData

var _ = Describe("Main", func() {
	Describe("MapToSlice", func() {
		It(`should change map[k1: d1, k2: d2 ...] to slice multidimension of string [[k1, d1], [k2, d2]... ]`, func() {
			testDataSlice := TestDataSlice{
				{
					Test: map[string]string{
						"hello": "world",
					},
					Expected: [][]string{
						{"hello", "world"},
					},
				},
				{
					Test: map[string]string{
						"hello": "world",
						"John":  "Doe",
						"Age":   "14",
					},
					Expected: [][]string{
						{"hello", "world"},
						{"John", "Doe"},
						{"Age", "14"},
					},
				},
				{
					Test: map[string]string{
						"foo": "33",
						"bar": "44",
						"baz": "55",
					},
					Expected: [][]string{
						{"foo", "33"},
						{"bar", "44"},
						{"baz", "55"},
					},
				},
			}

			for i := 0; i < len(testDataSlice); i++ {
				result := main.MapToSlice(testDataSlice[i].Test)

				Expect(len(result)).To(Equal(len(testDataSlice[i].Expected)))
				Expect(result).To(ContainElements(testDataSlice[i].Expected))
			}
		})

		It("should change big map data to slice multidimension of string", func() {
			testDataSlice := TestDataSlice{
				{
					Test: map[string]string{
						"2AAmyxdNoH": "1Tyd4YMWS9YtNDl",
						"6Y9TOxmnRW": "I55yYyPyLdNTihp",
						"6wxZPqx1w9": "ONcMRrQkf75W4Rp",
						"FJV7VgatQ1": "a24OaFlDFUkxIlu",
						"RsvJZqcAAg": "sC2uoBmIU5giCT0",
						"UjO4vEuYCy": "XapD94SxiqyC95e",
						"apv840AE9y": "1yFfXJAPMJlw5Zz",
						"qcaY750D1A": "4EXtN4Z2dnmX3sp",
						"tb5RPg3twh": "I9FFuQnxRmtcgZq",
						"wDIHXkA2fI": "65zsfdNYPr0wsTv",
					},
					Expected: [][]string{
						{"2AAmyxdNoH", "1Tyd4YMWS9YtNDl"},
						{"6Y9TOxmnRW", "I55yYyPyLdNTihp"},
						{"6wxZPqx1w9", "ONcMRrQkf75W4Rp"},
						{"FJV7VgatQ1", "a24OaFlDFUkxIlu"},
						{"RsvJZqcAAg", "sC2uoBmIU5giCT0"},
						{"UjO4vEuYCy", "XapD94SxiqyC95e"},
						{"apv840AE9y", "1yFfXJAPMJlw5Zz"},
						{"qcaY750D1A", "4EXtN4Z2dnmX3sp"},
						{"tb5RPg3twh", "I9FFuQnxRmtcgZq"},
						{"wDIHXkA2fI", "65zsfdNYPr0wsTv"},
					},
				},
				{
					Test: map[string]string{
						"2LQLu6KnLEwfkTi6MYcLC4uUNuR3BsL3Z4Y": "N119gFilVWOSncFBttZlQLQaP5CyxPccIxhOl0kidYPObB6Yzv",
						"3STooo4WxU5E1botaTDRk3fTFx9YUA6acpA": "GFjDN4F9mb1w0IUAbonzcXRh5YmWJPx8GDcnHkeypsxJi4HsyX",
						"3upuIaux8dSCynLdTtb4HLdhhIMBWBemLP2": "tovrkxkDz67Av2xm1ZrkDjqAVdPe481EuapLZn3iNlA0KkYF8o",
						"5Hm5RLJiLL54h1qoFq7vM3IfMMV23aUatnf": "9T5H5yDztNpSyfohadDksLIelTunqtgv66oFJU881z8CA6VAiq",
						"5mBMd7ubPLAeYamBpsEjnUsTKGY0xXZXPQQ": "WBM37DboUpbLqbKf8iHDlZ5k93u9vfxwbHTwrgpwbg2OFyWbwm",
						"7NRywHKgLxQtmjJksINuH5MqSpPbKZ8cMMp": "xXKBze2k7RRxRFqVwEaVKbYxYaDFF2naqm2qygU2XZFm8NsHOr",
						"8COVJ4b1ki4Te8ckS4fSyEBBwXv8nL5Sfsx": "jfQzRFxLaAh6pJUTRkafB3R09OuKWWrvEP2O8Mh8Nrl54KDBx5",
						"8FsAWIjdKy53SgrxKZCfno2dtDZ9J1S3u01": "tYm4rWBOu85cJxplIzqBWkAdetxmvxgeDywgm3YFpSNI118ES5",
						"8mkXpPa5rm4FJPuuDOvnDk3pIJPcoL0aE5C": "4YsjOK4Ifs2SQIEeJ8D23IBfUXXHUOhHw4cUkMVCTcmUiyvHJs",
						"CnDONxiBSDZL3vfrnnQfqVzpVO7WzUpJRSF": "8Wgbsr6sTDglgUMwKQMaGJiTT5oSQrytGUFJI73QpDqhjAxV57",
						"DicETR9Doa0JsrnKkE44UE0331Gbt6MX8UK": "U9SefTxFOo3nEtYpCuZaxCmdx6uK6scqeXCU91fzc0ZHmk9yF2",
						"DuDt77YP7oYnj7oOj2ma4QVYfjbAKNHlUgR": "BaOMQVqLMAWzCMIHDegQ21sTYgFazmTX0Yn2umRuX8B9fDoUKN",
						"EZ5H5p0gB1qJoojsyiJ9nywadnNkV3CZ6rB": "r1ZSsmKKAidm7mKmPqKMV69H5BotOMzYeH9lie0Q9P9XG8sXlz",
						"FFRqNsLvmq57k4qgmZ6wNC9E1EDXFhGOXfS": "qQmnmdSCEODBJ7f2JYFTKp47uGFFVXpBogzfTp5Ubc06KWWtpC",
						"HluOhN50IUs2dvjVv2q2p4KsLFtpG3lRvjK": "LwqDCrk6ySX1AVkqmZssaUgmi95BFpCnqhfk2QQEDXpoiHYF1W",
						"ICEtjqhLmQO4E6sLf8rQkxue5QK7bEs1kn4": "Utt4feX3gAGERGk2iQG9pDdQkQLlLBOd6qAZ31PdLFbbSEWl2D",
						"KRqpVt702eqVJHNZ3GVeKt0aE4nSbcFSKxS": "i0bvR4LnQzcqlxC8zE3oHg4K6XrX1Uu0CxZuG924Ed2lLPrVDF",
						"KSTVn0HEnGwFTtFE85Jrglx5agF89YscqLW": "ob5nup3p18KO378V5382Qhmk0G8aj4MiyAeMaJ3C1J5rj8jpGF",
						"Kiyr0A6DVgThuFaTOb6X4y1RwWBpjbA5Urg": "pH4INkGNFCkewwyyF4LW7mDU3RX6KTdZQH05awFkGVKveag90V",
						"Kq3e51SYrzRCRE7yyADmVkBm4sNk8B61zBo": "MB20bhAu94hAKywuGrKE51sN7EQQ03ONl8I05dZrD5rbqSCi6x",
						"Krh6zE5c52ceNoBnwPTKti9jlre3y2G3RQC": "b1eXTRc7U29FkJmDa3gv5q5KqRv20oI7mTFxMJGrjwNxtAVwfI",
						"Log7TRblcU284BvQl9JuT7ionbwEmAlmYyN": "44v0xSp6A5Ag9J9DegHTBtz044BVHDlNpvLc8qOzRV5HidcIjS",
						"LxYz50VHayfkoFvwWauX6HGJc3gnhAQZjA3": "vmavrYVR5PdcXNT3tXNl8Eced6qElGr5GIO4z5Mu5ORK8nnr0g",
						"M9SWUyPsG945Jlqo8apsQsCNwOlccY5d3FY": "xiwEbKvm9j2leeI5F1sugASb0k2Cos5Qy1W2X6kV5MuLnGGBkc",
						"MDGxzNXY5mtKXJPUgNs3oFAczIQVJeOTPn2": "oiLvHzOQNtiDWrXXkmI7W7leXLh9VZ4bBDws72yWu23jn07Wj4",
						"PU9FzpfejyW8tRrkLilvbvieMItyW41BEQn": "Wp5iE9NiVmnG9E3tHwRDzUs4S5OvgbjGwUH7OypxSuwD9tH3Nn",
						"PWioDCN3126h8MSm6dyPZb3cP4i1YkGiFft": "mpKneAM2WBMN66HprEuHtMF03LUHVnUczPkFg5HAg5k25is57k",
						"Q90MoNMUSaFM5vyjerRYIjx9Vs5UQiechOq": "T1YalcooLk6AnJ58uKG8EvkAfeRd16wTyDcP505YRXq6U8T84j",
						"QJ7CtZlDnBJkizOtDeYa79mb5syws65WMbO": "48rN4rv6LPeunwjPiwXJ5ikUqSjz9v6ZCsvNdUIBEq7faLpmWR",
						"RA9xcSAxt54UyiZgTwQ77Z0MIKBxDajx2EW": "qJMXUWYI0mZHlj43x0ruFmUnj3P5zQjTv5ZnnoB5QWrabAGn9k",
						"SNfRTso5adT1IvJLJA7c0XxLJHiIThwc1tv": "pYEr7ZtZQI6t94WuvqpV6eqwxiENLeDoYQkKtYYsOcp47WYGvS",
						"TOiSoZj56DPy6nVT1H5hdgo2yaIHg1aEsue": "wekWQ9JLs905EYjjJzo55O42TVI2B5lWWviRu61rlIHSuchVsE",
						"XjA64e1XbVNlp73ie9byhIRqDjjnxIVZJAX": "Uh0HDCMg3ChBrGHxp8M43NvOglzi55bdFtKQnf5MyTqyB1OX5w",
						"YU9bsB9BFvr2U9k6kQHbElyVCobNh3AYmUF": "xqZkTLNLIojEVcZ5q5SK6CplmHvDdHGebBDEECHQ013Zq3E85M",
						"aQfy5jAYlKVeXsW2eU3FHUT6rwfrSNo3LbM": "KqVmjEt2kY43SJLn18KlEz5vM8igQlh4NMxZHE0V6j2HYMFW23",
						"c9N055HtPYGXZ91bSCiXRm3XuvuS2esZtog": "0LOlxsBkXKGFk4owx5Ec3UB5hlX3tYayIlL0V9z51l8mUD9Wpl",
						"cOI9dsM5Ik7fMdOOpsWIRa0AYOfQ8GGszbx": "SsNIoBKebg8JsxHustaVdNL0FRajLSnXFE7VRaNjZFXtn5Layb",
						"fBjR9chSHL3t0a9cZyGpyF9Z1KMi28aohBN": "7aSLRxrj5l2NH0YetvoxTki2WXf1SNtnB7J23VI59usf7EFImO",
						"gKmEwpbtEM52pFuDk39t44Vg0W2kJUwaX9J": "rt1hNu4x8eelHX5P1tNr3X2PzPdisiDkqhdXaQv5CJFiOPQf0G",
						"gsbDgJmVNok8nznEOuT2TrJ47oMAJDNCNk5": "rn0160PeendrEcAZU0ATE0QY2rjT3DxsBUBCqgmYWv2p6jn1ld",
						"iz04Ohk1tMGKcrhsqhZOy8nk5MjC0w8FwMl": "aGrXrejT9iRgcctcC9HPf3W74ygE0dq3Z8SfVKE4xj8GmpGmm8",
						"lD9rK8DgcDED9ZZa6kNLCpirU0uI7Mauncg": "HeGVEQAhef5Y2arYeq0O6jZM2X7H1XQiXbMbfgdsgqC0d1BE40",
						"lasJWKszLdb1kaY84430CemgGAhYyDdfzCZ": "1cvio9XYCWJrLwig5gIBKC9IJ5h2vNY1a29eY3S0hL8lCWsDzA",
						"m9YMejYa1QQB3YiOcdHkEMDCj513Sewup9M": "JTGiBPPZG6jTBo5u2DlpIBOO1gNyqTACac27kxIgGXNx3UKllm",
						"sZ36SnOKV5rkHbMB75s0SN0FQmnU1BxAMzf": "scF3slrCtDAXo7rb9j0anFVQPCH2UDUixCEEc4hpNh2APu3UkS",
						"ssPqs4i8oOKj2tqlpBFzWV5SfisYSR4VAZG": "wQGGtwAQejZpkIbOIsVGF4L576qlnmVtKfnofHCyhBMpCjdQOh",
						"tpli58mmdxpaz5bA4Xig5BskpQavH5C5HCF": "uyhwdXHsoAQ25QJGqkcYdIoG1Ss8z3UcinpCu8Wap1W44IQoTX",
						"uYvnkPtDhhWNJcGM5Mz4lQ3isB8JQY4VqzI": "78e5MaicejeBGl0bRz7q5XBYn42US3V0zA9anA8rv8VFo0bfR1",
						"v2m7m1FZLO325Oi5TYQp8vGAF0y33iXfzgj": "wIWlESBYfN2T4fGG9TkPIrc86QyIlhgnqPFLaqK2QauCA0n53j",
						"wBM22TIPlu7Ed5foHhOPqf5SEm9OiyH2jzU": "dfXSEAbx2bkhnsgGOHiOxcP8WUq7xA6ZZwR7LOki2X9l5IOANk",
					},
					Expected: [][]string{
						{"2LQLu6KnLEwfkTi6MYcLC4uUNuR3BsL3Z4Y", "N119gFilVWOSncFBttZlQLQaP5CyxPccIxhOl0kidYPObB6Yzv"},
						{"3STooo4WxU5E1botaTDRk3fTFx9YUA6acpA", "GFjDN4F9mb1w0IUAbonzcXRh5YmWJPx8GDcnHkeypsxJi4HsyX"},
						{"3upuIaux8dSCynLdTtb4HLdhhIMBWBemLP2", "tovrkxkDz67Av2xm1ZrkDjqAVdPe481EuapLZn3iNlA0KkYF8o"},
						{"5Hm5RLJiLL54h1qoFq7vM3IfMMV23aUatnf", "9T5H5yDztNpSyfohadDksLIelTunqtgv66oFJU881z8CA6VAiq"},
						{"5mBMd7ubPLAeYamBpsEjnUsTKGY0xXZXPQQ", "WBM37DboUpbLqbKf8iHDlZ5k93u9vfxwbHTwrgpwbg2OFyWbwm"},
						{"7NRywHKgLxQtmjJksINuH5MqSpPbKZ8cMMp", "xXKBze2k7RRxRFqVwEaVKbYxYaDFF2naqm2qygU2XZFm8NsHOr"},
						{"8COVJ4b1ki4Te8ckS4fSyEBBwXv8nL5Sfsx", "jfQzRFxLaAh6pJUTRkafB3R09OuKWWrvEP2O8Mh8Nrl54KDBx5"},
						{"8FsAWIjdKy53SgrxKZCfno2dtDZ9J1S3u01", "tYm4rWBOu85cJxplIzqBWkAdetxmvxgeDywgm3YFpSNI118ES5"},
						{"8mkXpPa5rm4FJPuuDOvnDk3pIJPcoL0aE5C", "4YsjOK4Ifs2SQIEeJ8D23IBfUXXHUOhHw4cUkMVCTcmUiyvHJs"},
						{"CnDONxiBSDZL3vfrnnQfqVzpVO7WzUpJRSF", "8Wgbsr6sTDglgUMwKQMaGJiTT5oSQrytGUFJI73QpDqhjAxV57"},
						{"DicETR9Doa0JsrnKkE44UE0331Gbt6MX8UK", "U9SefTxFOo3nEtYpCuZaxCmdx6uK6scqeXCU91fzc0ZHmk9yF2"},
						{"DuDt77YP7oYnj7oOj2ma4QVYfjbAKNHlUgR", "BaOMQVqLMAWzCMIHDegQ21sTYgFazmTX0Yn2umRuX8B9fDoUKN"},
						{"EZ5H5p0gB1qJoojsyiJ9nywadnNkV3CZ6rB", "r1ZSsmKKAidm7mKmPqKMV69H5BotOMzYeH9lie0Q9P9XG8sXlz"},
						{"FFRqNsLvmq57k4qgmZ6wNC9E1EDXFhGOXfS", "qQmnmdSCEODBJ7f2JYFTKp47uGFFVXpBogzfTp5Ubc06KWWtpC"},
						{"HluOhN50IUs2dvjVv2q2p4KsLFtpG3lRvjK", "LwqDCrk6ySX1AVkqmZssaUgmi95BFpCnqhfk2QQEDXpoiHYF1W"},
						{"ICEtjqhLmQO4E6sLf8rQkxue5QK7bEs1kn4", "Utt4feX3gAGERGk2iQG9pDdQkQLlLBOd6qAZ31PdLFbbSEWl2D"},
						{"KRqpVt702eqVJHNZ3GVeKt0aE4nSbcFSKxS", "i0bvR4LnQzcqlxC8zE3oHg4K6XrX1Uu0CxZuG924Ed2lLPrVDF"},
						{"KSTVn0HEnGwFTtFE85Jrglx5agF89YscqLW", "ob5nup3p18KO378V5382Qhmk0G8aj4MiyAeMaJ3C1J5rj8jpGF"},
						{"Kiyr0A6DVgThuFaTOb6X4y1RwWBpjbA5Urg", "pH4INkGNFCkewwyyF4LW7mDU3RX6KTdZQH05awFkGVKveag90V"},
						{"Kq3e51SYrzRCRE7yyADmVkBm4sNk8B61zBo", "MB20bhAu94hAKywuGrKE51sN7EQQ03ONl8I05dZrD5rbqSCi6x"},
						{"Krh6zE5c52ceNoBnwPTKti9jlre3y2G3RQC", "b1eXTRc7U29FkJmDa3gv5q5KqRv20oI7mTFxMJGrjwNxtAVwfI"},
						{"Log7TRblcU284BvQl9JuT7ionbwEmAlmYyN", "44v0xSp6A5Ag9J9DegHTBtz044BVHDlNpvLc8qOzRV5HidcIjS"},
						{"LxYz50VHayfkoFvwWauX6HGJc3gnhAQZjA3", "vmavrYVR5PdcXNT3tXNl8Eced6qElGr5GIO4z5Mu5ORK8nnr0g"},
						{"M9SWUyPsG945Jlqo8apsQsCNwOlccY5d3FY", "xiwEbKvm9j2leeI5F1sugASb0k2Cos5Qy1W2X6kV5MuLnGGBkc"},
						{"MDGxzNXY5mtKXJPUgNs3oFAczIQVJeOTPn2", "oiLvHzOQNtiDWrXXkmI7W7leXLh9VZ4bBDws72yWu23jn07Wj4"},
						{"PU9FzpfejyW8tRrkLilvbvieMItyW41BEQn", "Wp5iE9NiVmnG9E3tHwRDzUs4S5OvgbjGwUH7OypxSuwD9tH3Nn"},
						{"PWioDCN3126h8MSm6dyPZb3cP4i1YkGiFft", "mpKneAM2WBMN66HprEuHtMF03LUHVnUczPkFg5HAg5k25is57k"},
						{"Q90MoNMUSaFM5vyjerRYIjx9Vs5UQiechOq", "T1YalcooLk6AnJ58uKG8EvkAfeRd16wTyDcP505YRXq6U8T84j"},
						{"QJ7CtZlDnBJkizOtDeYa79mb5syws65WMbO", "48rN4rv6LPeunwjPiwXJ5ikUqSjz9v6ZCsvNdUIBEq7faLpmWR"},
						{"RA9xcSAxt54UyiZgTwQ77Z0MIKBxDajx2EW", "qJMXUWYI0mZHlj43x0ruFmUnj3P5zQjTv5ZnnoB5QWrabAGn9k"},
						{"SNfRTso5adT1IvJLJA7c0XxLJHiIThwc1tv", "pYEr7ZtZQI6t94WuvqpV6eqwxiENLeDoYQkKtYYsOcp47WYGvS"},
						{"TOiSoZj56DPy6nVT1H5hdgo2yaIHg1aEsue", "wekWQ9JLs905EYjjJzo55O42TVI2B5lWWviRu61rlIHSuchVsE"},
						{"XjA64e1XbVNlp73ie9byhIRqDjjnxIVZJAX", "Uh0HDCMg3ChBrGHxp8M43NvOglzi55bdFtKQnf5MyTqyB1OX5w"},
						{"YU9bsB9BFvr2U9k6kQHbElyVCobNh3AYmUF", "xqZkTLNLIojEVcZ5q5SK6CplmHvDdHGebBDEECHQ013Zq3E85M"},
						{"aQfy5jAYlKVeXsW2eU3FHUT6rwfrSNo3LbM", "KqVmjEt2kY43SJLn18KlEz5vM8igQlh4NMxZHE0V6j2HYMFW23"},
						{"c9N055HtPYGXZ91bSCiXRm3XuvuS2esZtog", "0LOlxsBkXKGFk4owx5Ec3UB5hlX3tYayIlL0V9z51l8mUD9Wpl"},
						{"cOI9dsM5Ik7fMdOOpsWIRa0AYOfQ8GGszbx", "SsNIoBKebg8JsxHustaVdNL0FRajLSnXFE7VRaNjZFXtn5Layb"},
						{"fBjR9chSHL3t0a9cZyGpyF9Z1KMi28aohBN", "7aSLRxrj5l2NH0YetvoxTki2WXf1SNtnB7J23VI59usf7EFImO"},
						{"gKmEwpbtEM52pFuDk39t44Vg0W2kJUwaX9J", "rt1hNu4x8eelHX5P1tNr3X2PzPdisiDkqhdXaQv5CJFiOPQf0G"},
						{"gsbDgJmVNok8nznEOuT2TrJ47oMAJDNCNk5", "rn0160PeendrEcAZU0ATE0QY2rjT3DxsBUBCqgmYWv2p6jn1ld"},
						{"iz04Ohk1tMGKcrhsqhZOy8nk5MjC0w8FwMl", "aGrXrejT9iRgcctcC9HPf3W74ygE0dq3Z8SfVKE4xj8GmpGmm8"},
						{"lD9rK8DgcDED9ZZa6kNLCpirU0uI7Mauncg", "HeGVEQAhef5Y2arYeq0O6jZM2X7H1XQiXbMbfgdsgqC0d1BE40"},
						{"lasJWKszLdb1kaY84430CemgGAhYyDdfzCZ", "1cvio9XYCWJrLwig5gIBKC9IJ5h2vNY1a29eY3S0hL8lCWsDzA"},
						{"m9YMejYa1QQB3YiOcdHkEMDCj513Sewup9M", "JTGiBPPZG6jTBo5u2DlpIBOO1gNyqTACac27kxIgGXNx3UKllm"},
						{"sZ36SnOKV5rkHbMB75s0SN0FQmnU1BxAMzf", "scF3slrCtDAXo7rb9j0anFVQPCH2UDUixCEEc4hpNh2APu3UkS"},
						{"ssPqs4i8oOKj2tqlpBFzWV5SfisYSR4VAZG", "wQGGtwAQejZpkIbOIsVGF4L576qlnmVtKfnofHCyhBMpCjdQOh"},
						{"tpli58mmdxpaz5bA4Xig5BskpQavH5C5HCF", "uyhwdXHsoAQ25QJGqkcYdIoG1Ss8z3UcinpCu8Wap1W44IQoTX"},
						{"uYvnkPtDhhWNJcGM5Mz4lQ3isB8JQY4VqzI", "78e5MaicejeBGl0bRz7q5XBYn42US3V0zA9anA8rv8VFo0bfR1"},
						{"v2m7m1FZLO325Oi5TYQp8vGAF0y33iXfzgj", "wIWlESBYfN2T4fGG9TkPIrc86QyIlhgnqPFLaqK2QauCA0n53j"},
						{"wBM22TIPlu7Ed5foHhOPqf5SEm9OiyH2jzU", "dfXSEAbx2bkhnsgGOHiOxcP8WUq7xA6ZZwR7LOki2X9l5IOANk"},
					},
				},
			}

			for i := 0; i < len(testDataSlice); i++ {
				res := main.MapToSlice(testDataSlice[i].Test)

				Expect(len(res)).To(Equal(len(testDataSlice[i].Expected)))
				Expect(res).To(ContainElements(testDataSlice[i].Expected))
			}
		})
	})
})
