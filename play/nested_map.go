package main

// var coordinates = map[string]interface{
// 	"asdf": {
// 		"UpiHandle1": (12, 3),
// 		"UpiHandle2": (11, 3),
// 		"UpiHandle3": (10, 3),
// 		"MerchantName": (9, 3),
// 	},
// }

func getCoordinates(businessCategory string, field string) (int, int) {
	switch businessCategory {
	case "A":
		switch field {
		case "UpiHandle1":
			return 12, 3
		}
	}
	return 0, 0
}
