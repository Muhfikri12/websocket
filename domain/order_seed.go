package domain

func OrderSeed() []Order {
	return []Order{
		{
			CustomerID:    3,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 85, Quantity: 19, UnitPrice: 364},
				{VariantID: 195, Quantity: 7, UnitPrice: 357},
				{VariantID: 124, Quantity: 8, UnitPrice: 476},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "bank transfer",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 14, Quantity: 13, UnitPrice: 196},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "debit",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 57, Quantity: 29, UnitPrice: 236},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "qris",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 191, Quantity: 28, UnitPrice: 264},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 71, Quantity: 25, UnitPrice: 564},
				{VariantID: 32, Quantity: 27, UnitPrice: 474},
				{VariantID: 37, Quantity: 6, UnitPrice: 127},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "bank transfer",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 78, Quantity: 9, UnitPrice: 593},
				{VariantID: 75, Quantity: 27, UnitPrice: 558},
				{VariantID: 199, Quantity: 7, UnitPrice: 129},
				{VariantID: 194, Quantity: 26, UnitPrice: 131},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "qris",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 76, Quantity: 25, UnitPrice: 104},
				{VariantID: 62, Quantity: 14, UnitPrice: 187},
				{VariantID: 13, Quantity: 16, UnitPrice: 346},
				{VariantID: 8, Quantity: 11, UnitPrice: 730},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 52, Quantity: 11, UnitPrice: 289},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 104, Quantity: 20, UnitPrice: 273},
				{VariantID: 127, Quantity: 8, UnitPrice: 419},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "credit card",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 156, Quantity: 7, UnitPrice: 583},
				{VariantID: 195, Quantity: 28, UnitPrice: 292},
				{VariantID: 120, Quantity: 11, UnitPrice: 309},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "bank transfer",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 137, Quantity: 4, UnitPrice: 461},
				{VariantID: 94, Quantity: 11, UnitPrice: 281},
				{VariantID: 187, Quantity: 11, UnitPrice: 362},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "qris",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 91, Quantity: 26, UnitPrice: 709},
				{VariantID: 8, Quantity: 10, UnitPrice: 175},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "bank transfer",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 98, Quantity: 10, UnitPrice: 145},
				{VariantID: 36, Quantity: 9, UnitPrice: 150},
			},
		},
		{
			CustomerID:    8,
			PaymentMethod: "qris",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 156, Quantity: 8, UnitPrice: 248},
				{VariantID: 143, Quantity: 16, UnitPrice: 104},
				{VariantID: 82, Quantity: 24, UnitPrice: 179},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "qris",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 53, Quantity: 18, UnitPrice: 336},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 60, Quantity: 30, UnitPrice: 179},
				{VariantID: 63, Quantity: 23, UnitPrice: 544},
				{VariantID: 132, Quantity: 20, UnitPrice: 609},
				{VariantID: 170, Quantity: 16, UnitPrice: 421},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "bank transfer",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 141, Quantity: 28, UnitPrice: 299},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "debit",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 84, Quantity: 26, UnitPrice: 705},
			},
		},
		{
			CustomerID:    8,
			PaymentMethod: "debit",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 99, Quantity: 4, UnitPrice: 143},
				{VariantID: 109, Quantity: 10, UnitPrice: 664},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "debit",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 151, Quantity: 6, UnitPrice: 364},
				{VariantID: 97, Quantity: 11, UnitPrice: 445},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 35, Quantity: 28, UnitPrice: 715},
				{VariantID: 170, Quantity: 5, UnitPrice: 300},
				{VariantID: 200, Quantity: 28, UnitPrice: 196},
				{VariantID: 169, Quantity: 28, UnitPrice: 159},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "credit card",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 53, Quantity: 4, UnitPrice: 380},
				{VariantID: 180, Quantity: 2, UnitPrice: 617},
				{VariantID: 109, Quantity: 5, UnitPrice: 323},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "bank transfer",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 5, Quantity: 14, UnitPrice: 119},
				{VariantID: 158, Quantity: 8, UnitPrice: 567},
				{VariantID: 160, Quantity: 16, UnitPrice: 527},
			},
		},
		{
			CustomerID:    8,
			PaymentMethod: "bank transfer",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 87, Quantity: 10, UnitPrice: 243},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 74, Quantity: 29, UnitPrice: 338},
			},
		},
		{
			CustomerID:    3,
			PaymentMethod: "cod",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 21, Quantity: 3, UnitPrice: 485},
				{VariantID: 173, Quantity: 28, UnitPrice: 112},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "cod",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 68, Quantity: 6, UnitPrice: 575},
				{VariantID: 42, Quantity: 12, UnitPrice: 239},
				{VariantID: 180, Quantity: 5, UnitPrice: 693},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "qris",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 14, Quantity: 9, UnitPrice: 746},
				{VariantID: 122, Quantity: 20, UnitPrice: 106},
				{VariantID: 142, Quantity: 28, UnitPrice: 530},
				{VariantID: 200, Quantity: 6, UnitPrice: 734},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "credit card",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 19, Quantity: 13, UnitPrice: 306},
				{VariantID: 39, Quantity: 20, UnitPrice: 129},
				{VariantID: 120, Quantity: 1, UnitPrice: 420},
				{VariantID: 189, Quantity: 30, UnitPrice: 576},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "credit card",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 123, Quantity: 2, UnitPrice: 383},
				{VariantID: 174, Quantity: 2, UnitPrice: 365},
				{VariantID: 170, Quantity: 27, UnitPrice: 533},
				{VariantID: 82, Quantity: 19, UnitPrice: 180},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "credit card",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 78, Quantity: 26, UnitPrice: 212},
				{VariantID: 125, Quantity: 4, UnitPrice: 453},
				{VariantID: 26, Quantity: 20, UnitPrice: 650},
			},
		},
		{
			CustomerID:    8,
			PaymentMethod: "credit card",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 37, Quantity: 8, UnitPrice: 289},
			},
		},
		{
			CustomerID:    3,
			PaymentMethod: "debit",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 68, Quantity: 3, UnitPrice: 117},
				{VariantID: 17, Quantity: 9, UnitPrice: 121},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "cod",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 74, Quantity: 27, UnitPrice: 619},
				{VariantID: 55, Quantity: 11, UnitPrice: 473},
			},
		},
		{
			CustomerID:    3,
			PaymentMethod: "qris",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 133, Quantity: 28, UnitPrice: 158},
				{VariantID: 176, Quantity: 4, UnitPrice: 366},
			},
		},
		{
			CustomerID:    3,
			PaymentMethod: "debit",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 65, Quantity: 18, UnitPrice: 676},
				{VariantID: 97, Quantity: 5, UnitPrice: 181},
				{VariantID: 66, Quantity: 13, UnitPrice: 682},
			},
		},
		{
			CustomerID:    8,
			PaymentMethod: "bank transfer",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 161, Quantity: 6, UnitPrice: 192},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 149, Quantity: 24, UnitPrice: 649},
				{VariantID: 193, Quantity: 27, UnitPrice: 321},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 173, Quantity: 22, UnitPrice: 506},
				{VariantID: 200, Quantity: 23, UnitPrice: 161},
				{VariantID: 9, Quantity: 27, UnitPrice: 364},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "qris",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 198, Quantity: 22, UnitPrice: 582},
				{VariantID: 28, Quantity: 24, UnitPrice: 734},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "qris",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 104, Quantity: 30, UnitPrice: 410},
				{VariantID: 193, Quantity: 2, UnitPrice: 486},
				{VariantID: 113, Quantity: 28, UnitPrice: 149},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "cod",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 52, Quantity: 23, UnitPrice: 735},
				{VariantID: 50, Quantity: 26, UnitPrice: 502},
				{VariantID: 21, Quantity: 25, UnitPrice: 293},
			},
		},
		{
			CustomerID:    8,
			PaymentMethod: "bank transfer",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 98, Quantity: 6, UnitPrice: 192},
				{VariantID: 96, Quantity: 26, UnitPrice: 603},
			},
		},
		{
			CustomerID:    3,
			PaymentMethod: "credit card",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 85, Quantity: 15, UnitPrice: 128},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "credit card",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 178, Quantity: 18, UnitPrice: 425},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "bank transfer",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 152, Quantity: 8, UnitPrice: 742},
				{VariantID: 51, Quantity: 22, UnitPrice: 596},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "bank transfer",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 26, Quantity: 23, UnitPrice: 561},
				{VariantID: 196, Quantity: 3, UnitPrice: 664},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "bank transfer",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 162, Quantity: 7, UnitPrice: 504},
				{VariantID: 77, Quantity: 30, UnitPrice: 341},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "bank transfer",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 24, Quantity: 26, UnitPrice: 737},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "cod",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 48, Quantity: 7, UnitPrice: 713},
				{VariantID: 40, Quantity: 23, UnitPrice: 370},
				{VariantID: 12, Quantity: 29, UnitPrice: 749},
				{VariantID: 61, Quantity: 10, UnitPrice: 186},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "qris",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 142, Quantity: 17, UnitPrice: 157},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 11, Quantity: 28, UnitPrice: 296},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "bank transfer",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 42, Quantity: 17, UnitPrice: 725},
				{VariantID: 88, Quantity: 30, UnitPrice: 214},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "bank transfer",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 38, Quantity: 29, UnitPrice: 640},
				{VariantID: 20, Quantity: 1, UnitPrice: 697},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "bank transfer",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 86, Quantity: 2, UnitPrice: 576},
				{VariantID: 59, Quantity: 13, UnitPrice: 701},
				{VariantID: 175, Quantity: 28, UnitPrice: 648},
				{VariantID: 70, Quantity: 26, UnitPrice: 197},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 31, Quantity: 27, UnitPrice: 415},
				{VariantID: 54, Quantity: 22, UnitPrice: 169},
				{VariantID: 130, Quantity: 29, UnitPrice: 316},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "credit card",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 17, Quantity: 21, UnitPrice: 687},
				{VariantID: 70, Quantity: 17, UnitPrice: 341},
				{VariantID: 116, Quantity: 4, UnitPrice: 554},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "credit card",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 40, Quantity: 14, UnitPrice: 108},
				{VariantID: 84, Quantity: 16, UnitPrice: 527},
				{VariantID: 187, Quantity: 29, UnitPrice: 601},
				{VariantID: 135, Quantity: 16, UnitPrice: 275},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "qris",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 83, Quantity: 18, UnitPrice: 308},
				{VariantID: 112, Quantity: 22, UnitPrice: 398},
				{VariantID: 195, Quantity: 13, UnitPrice: 273},
				{VariantID: 109, Quantity: 19, UnitPrice: 294},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "credit card",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 157, Quantity: 12, UnitPrice: 299},
				{VariantID: 86, Quantity: 10, UnitPrice: 556},
				{VariantID: 193, Quantity: 28, UnitPrice: 662},
				{VariantID: 33, Quantity: 12, UnitPrice: 309},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "credit card",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 150, Quantity: 21, UnitPrice: 486},
				{VariantID: 60, Quantity: 13, UnitPrice: 308},
				{VariantID: 80, Quantity: 18, UnitPrice: 583},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "credit card",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 127, Quantity: 20, UnitPrice: 188},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "bank transfer",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 89, Quantity: 16, UnitPrice: 446},
				{VariantID: 55, Quantity: 12, UnitPrice: 703},
				{VariantID: 78, Quantity: 3, UnitPrice: 173},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 35, Quantity: 23, UnitPrice: 384},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "qris",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 175, Quantity: 27, UnitPrice: 742},
				{VariantID: 87, Quantity: 18, UnitPrice: 398},
				{VariantID: 28, Quantity: 13, UnitPrice: 115},
				{VariantID: 117, Quantity: 3, UnitPrice: 723},
			},
		},
		{
			CustomerID:    3,
			PaymentMethod: "bank transfer",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 127, Quantity: 16, UnitPrice: 650},
				{VariantID: 1, Quantity: 29, UnitPrice: 537},
				{VariantID: 91, Quantity: 30, UnitPrice: 563},
				{VariantID: 58, Quantity: 24, UnitPrice: 252},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 143, Quantity: 21, UnitPrice: 283},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "cod",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 174, Quantity: 12, UnitPrice: 492},
				{VariantID: 131, Quantity: 8, UnitPrice: 655},
				{VariantID: 17, Quantity: 8, UnitPrice: 426},
				{VariantID: 45, Quantity: 9, UnitPrice: 165},
			},
		},
		{
			CustomerID:    3,
			PaymentMethod: "qris",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 160, Quantity: 11, UnitPrice: 614},
				{VariantID: 163, Quantity: 26, UnitPrice: 710},
				{VariantID: 181, Quantity: 27, UnitPrice: 231},
				{VariantID: 18, Quantity: 14, UnitPrice: 475},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 63, Quantity: 2, UnitPrice: 735},
				{VariantID: 199, Quantity: 11, UnitPrice: 609},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "debit",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 110, Quantity: 25, UnitPrice: 518},
				{VariantID: 42, Quantity: 21, UnitPrice: 444},
				{VariantID: 89, Quantity: 22, UnitPrice: 719},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "credit card",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 48, Quantity: 15, UnitPrice: 518},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "cod",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 28, Quantity: 16, UnitPrice: 562},
				{VariantID: 190, Quantity: 27, UnitPrice: 148},
				{VariantID: 59, Quantity: 1, UnitPrice: 433},
				{VariantID: 179, Quantity: 30, UnitPrice: 202},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "credit card",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 85, Quantity: 4, UnitPrice: 541},
				{VariantID: 133, Quantity: 20, UnitPrice: 677},
				{VariantID: 183, Quantity: 28, UnitPrice: 668},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "credit card",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 71, Quantity: 1, UnitPrice: 115},
				{VariantID: 190, Quantity: 7, UnitPrice: 608},
				{VariantID: 177, Quantity: 7, UnitPrice: 531},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "cod",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 147, Quantity: 3, UnitPrice: 437},
				{VariantID: 83, Quantity: 23, UnitPrice: 642},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "credit card",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 105, Quantity: 17, UnitPrice: 189},
				{VariantID: 134, Quantity: 13, UnitPrice: 551},
				{VariantID: 21, Quantity: 24, UnitPrice: 634},
				{VariantID: 161, Quantity: 28, UnitPrice: 505},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 71, Quantity: 3, UnitPrice: 265},
				{VariantID: 128, Quantity: 8, UnitPrice: 316},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 146, Quantity: 30, UnitPrice: 421},
				{VariantID: 109, Quantity: 12, UnitPrice: 645},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 189, Quantity: 11, UnitPrice: 706},
				{VariantID: 59, Quantity: 1, UnitPrice: 474},
				{VariantID: 88, Quantity: 15, UnitPrice: 151},
				{VariantID: 13, Quantity: 7, UnitPrice: 418},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "credit card",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 84, Quantity: 3, UnitPrice: 534},
				{VariantID: 34, Quantity: 2, UnitPrice: 176},
				{VariantID: 31, Quantity: 10, UnitPrice: 334},
				{VariantID: 60, Quantity: 3, UnitPrice: 357},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "bank transfer",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 43, Quantity: 5, UnitPrice: 132},
				{VariantID: 25, Quantity: 22, UnitPrice: 107},
				{VariantID: 75, Quantity: 2, UnitPrice: 598},
			},
		},
		{
			CustomerID:    3,
			PaymentMethod: "debit",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 125, Quantity: 18, UnitPrice: 465},
			},
		},
		{
			CustomerID:    8,
			PaymentMethod: "credit card",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 20, Quantity: 26, UnitPrice: 272},
				{VariantID: 155, Quantity: 16, UnitPrice: 123},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "bank transfer",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 104, Quantity: 25, UnitPrice: 431},
				{VariantID: 189, Quantity: 24, UnitPrice: 331},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "debit",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 25, Quantity: 14, UnitPrice: 501},
				{VariantID: 61, Quantity: 19, UnitPrice: 725},
				{VariantID: 167, Quantity: 5, UnitPrice: 423},
				{VariantID: 87, Quantity: 19, UnitPrice: 462},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "bank transfer",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 171, Quantity: 11, UnitPrice: 549},
				{VariantID: 102, Quantity: 19, UnitPrice: 180},
				{VariantID: 118, Quantity: 15, UnitPrice: 376},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "bank transfer",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 122, Quantity: 10, UnitPrice: 381},
				{VariantID: 27, Quantity: 13, UnitPrice: 727},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "credit card",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 94, Quantity: 24, UnitPrice: 329},
				{VariantID: 163, Quantity: 26, UnitPrice: 654},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "debit",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 25, Quantity: 16, UnitPrice: 245},
				{VariantID: 76, Quantity: 21, UnitPrice: 485},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "qris",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 103, Quantity: 26, UnitPrice: 312},
				{VariantID: 29, Quantity: 2, UnitPrice: 434},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "credit card",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 40, Quantity: 8, UnitPrice: 184},
				{VariantID: 114, Quantity: 29, UnitPrice: 703},
				{VariantID: 128, Quantity: 26, UnitPrice: 389},
				{VariantID: 82, Quantity: 12, UnitPrice: 224},
			},
		},
		{
			CustomerID:    8,
			PaymentMethod: "bank transfer",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 118, Quantity: 7, UnitPrice: 318},
				{VariantID: 181, Quantity: 28, UnitPrice: 598},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "debit",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 178, Quantity: 30, UnitPrice: 714},
				{VariantID: 132, Quantity: 5, UnitPrice: 282},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "cod",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 34, Quantity: 7, UnitPrice: 345},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 117, Quantity: 30, UnitPrice: 362},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "credit card",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 60, Quantity: 21, UnitPrice: 591},
				{VariantID: 64, Quantity: 27, UnitPrice: 696},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "bank transfer",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 114, Quantity: 6, UnitPrice: 173},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 136, Quantity: 11, UnitPrice: 504},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "qris",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 194, Quantity: 22, UnitPrice: 373},
				{VariantID: 42, Quantity: 10, UnitPrice: 451},
				{VariantID: 160, Quantity: 23, UnitPrice: 327},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "debit",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 6, Quantity: 17, UnitPrice: 636},
				{VariantID: 100, Quantity: 4, UnitPrice: 287},
				{VariantID: 134, Quantity: 4, UnitPrice: 378},
				{VariantID: 164, Quantity: 4, UnitPrice: 586},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "bank transfer",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 167, Quantity: 25, UnitPrice: 643},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "bank transfer",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 51, Quantity: 27, UnitPrice: 268},
				{VariantID: 143, Quantity: 4, UnitPrice: 475},
				{VariantID: 23, Quantity: 24, UnitPrice: 435},
				{VariantID: 168, Quantity: 5, UnitPrice: 187},
			},
		},
		{
			CustomerID:    8,
			PaymentMethod: "bank transfer",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 50, Quantity: 30, UnitPrice: 164},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "credit card",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 74, Quantity: 5, UnitPrice: 434},
				{VariantID: 83, Quantity: 19, UnitPrice: 433},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 62, Quantity: 12, UnitPrice: 160},
				{VariantID: 156, Quantity: 18, UnitPrice: 634},
				{VariantID: 107, Quantity: 27, UnitPrice: 490},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "bank transfer",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 99, Quantity: 23, UnitPrice: 529},
				{VariantID: 178, Quantity: 25, UnitPrice: 625},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 133, Quantity: 13, UnitPrice: 530},
				{VariantID: 126, Quantity: 7, UnitPrice: 184},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "qris",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 154, Quantity: 28, UnitPrice: 127},
				{VariantID: 62, Quantity: 12, UnitPrice: 399},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "debit",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 74, Quantity: 24, UnitPrice: 238},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 52, Quantity: 13, UnitPrice: 296},
				{VariantID: 17, Quantity: 11, UnitPrice: 425},
				{VariantID: 200, Quantity: 20, UnitPrice: 485},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "credit card",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 180, Quantity: 10, UnitPrice: 121},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "credit card",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 53, Quantity: 30, UnitPrice: 582},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 130, Quantity: 15, UnitPrice: 724},
				{VariantID: 71, Quantity: 4, UnitPrice: 411},
				{VariantID: 115, Quantity: 25, UnitPrice: 490},
				{VariantID: 52, Quantity: 25, UnitPrice: 552},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "qris",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 139, Quantity: 20, UnitPrice: 424},
				{VariantID: 131, Quantity: 10, UnitPrice: 500},
				{VariantID: 132, Quantity: 20, UnitPrice: 606},
				{VariantID: 82, Quantity: 12, UnitPrice: 447},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "cod",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 22, Quantity: 7, UnitPrice: 314},
				{VariantID: 167, Quantity: 26, UnitPrice: 267},
			},
		},
		{
			CustomerID:    8,
			PaymentMethod: "bank transfer",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 50, Quantity: 15, UnitPrice: 182},
				{VariantID: 152, Quantity: 29, UnitPrice: 632},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "qris",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 73, Quantity: 27, UnitPrice: 643},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 131, Quantity: 22, UnitPrice: 581},
				{VariantID: 26, Quantity: 29, UnitPrice: 470},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "debit",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 73, Quantity: 4, UnitPrice: 727},
				{VariantID: 68, Quantity: 25, UnitPrice: 562},
				{VariantID: 164, Quantity: 19, UnitPrice: 269},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "bank transfer",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 58, Quantity: 8, UnitPrice: 636},
				{VariantID: 140, Quantity: 30, UnitPrice: 359},
				{VariantID: 123, Quantity: 4, UnitPrice: 292},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "bank transfer",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 37, Quantity: 30, UnitPrice: 217},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "qris",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 144, Quantity: 24, UnitPrice: 615},
				{VariantID: 146, Quantity: 10, UnitPrice: 632},
				{VariantID: 128, Quantity: 28, UnitPrice: 650},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "credit card",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 195, Quantity: 25, UnitPrice: 420},
				{VariantID: 88, Quantity: 17, UnitPrice: 569},
				{VariantID: 93, Quantity: 24, UnitPrice: 343},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "credit card",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 195, Quantity: 18, UnitPrice: 306},
				{VariantID: 107, Quantity: 14, UnitPrice: 644},
				{VariantID: 93, Quantity: 18, UnitPrice: 256},
				{VariantID: 99, Quantity: 12, UnitPrice: 603},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "cod",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 72, Quantity: 24, UnitPrice: 632},
				{VariantID: 64, Quantity: 24, UnitPrice: 660},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "credit card",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 36, Quantity: 21, UnitPrice: 450},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "cod",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 141, Quantity: 12, UnitPrice: 378},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "qris",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 60, Quantity: 25, UnitPrice: 569},
				{VariantID: 96, Quantity: 16, UnitPrice: 257},
			},
		},
		{
			CustomerID:    3,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 9, Quantity: 4, UnitPrice: 216},
				{VariantID: 197, Quantity: 4, UnitPrice: 196},
				{VariantID: 157, Quantity: 16, UnitPrice: 483},
				{VariantID: 47, Quantity: 12, UnitPrice: 144},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "bank transfer",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 105, Quantity: 16, UnitPrice: 385},
				{VariantID: 186, Quantity: 6, UnitPrice: 152},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "cod",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 55, Quantity: 14, UnitPrice: 172},
				{VariantID: 107, Quantity: 12, UnitPrice: 623},
				{VariantID: 172, Quantity: 21, UnitPrice: 480},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "bank transfer",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 18, Quantity: 30, UnitPrice: 207},
				{VariantID: 143, Quantity: 2, UnitPrice: 619},
				{VariantID: 35, Quantity: 28, UnitPrice: 732},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "qris",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 75, Quantity: 9, UnitPrice: 167},
				{VariantID: 148, Quantity: 2, UnitPrice: 227},
				{VariantID: 4, Quantity: 25, UnitPrice: 331},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "qris",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 136, Quantity: 17, UnitPrice: 543},
				{VariantID: 175, Quantity: 20, UnitPrice: 356},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 85, Quantity: 23, UnitPrice: 521},
				{VariantID: 80, Quantity: 14, UnitPrice: 401},
				{VariantID: 187, Quantity: 8, UnitPrice: 528},
				{VariantID: 58, Quantity: 22, UnitPrice: 558},
			},
		},
		{
			CustomerID:    3,
			PaymentMethod: "credit card",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 164, Quantity: 8, UnitPrice: 579},
				{VariantID: 149, Quantity: 2, UnitPrice: 192},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "bank transfer",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 173, Quantity: 14, UnitPrice: 579},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "cod",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 82, Quantity: 13, UnitPrice: 116},
				{VariantID: 62, Quantity: 3, UnitPrice: 555},
				{VariantID: 87, Quantity: 3, UnitPrice: 230},
				{VariantID: 80, Quantity: 13, UnitPrice: 313},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "debit",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 77, Quantity: 29, UnitPrice: 414},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "credit card",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 2, Quantity: 1, UnitPrice: 520},
				{VariantID: 154, Quantity: 19, UnitPrice: 249},
				{VariantID: 195, Quantity: 5, UnitPrice: 745},
				{VariantID: 165, Quantity: 23, UnitPrice: 750},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "debit",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 84, Quantity: 23, UnitPrice: 653},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "bank transfer",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 114, Quantity: 15, UnitPrice: 350},
				{VariantID: 99, Quantity: 11, UnitPrice: 638},
				{VariantID: 54, Quantity: 10, UnitPrice: 614},
			},
		},
		{
			CustomerID:    3,
			PaymentMethod: "credit card",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 119, Quantity: 27, UnitPrice: 408},
				{VariantID: 37, Quantity: 30, UnitPrice: 571},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "qris",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 14, Quantity: 10, UnitPrice: 514},
				{VariantID: 65, Quantity: 10, UnitPrice: 293},
				{VariantID: 83, Quantity: 17, UnitPrice: 142},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "debit",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 36, Quantity: 3, UnitPrice: 456},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "debit",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 134, Quantity: 25, UnitPrice: 156},
				{VariantID: 91, Quantity: 20, UnitPrice: 447},
				{VariantID: 4, Quantity: 29, UnitPrice: 157},
				{VariantID: 65, Quantity: 4, UnitPrice: 437},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "bank transfer",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 96, Quantity: 23, UnitPrice: 738},
				{VariantID: 174, Quantity: 6, UnitPrice: 117},
				{VariantID: 87, Quantity: 28, UnitPrice: 384},
			},
		},
		{
			CustomerID:    8,
			PaymentMethod: "debit",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 107, Quantity: 27, UnitPrice: 589},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 82, Quantity: 10, UnitPrice: 212},
				{VariantID: 184, Quantity: 20, UnitPrice: 320},
				{VariantID: 47, Quantity: 10, UnitPrice: 199},
				{VariantID: 166, Quantity: 28, UnitPrice: 549},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 147, Quantity: 9, UnitPrice: 692},
				{VariantID: 35, Quantity: 4, UnitPrice: 226},
				{VariantID: 191, Quantity: 22, UnitPrice: 265},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "bank transfer",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 58, Quantity: 24, UnitPrice: 133},
				{VariantID: 16, Quantity: 3, UnitPrice: 701},
				{VariantID: 95, Quantity: 24, UnitPrice: 696},
				{VariantID: 98, Quantity: 13, UnitPrice: 116},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "credit card",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 47, Quantity: 20, UnitPrice: 153},
				{VariantID: 138, Quantity: 10, UnitPrice: 668},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "credit card",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 71, Quantity: 4, UnitPrice: 707},
				{VariantID: 22, Quantity: 12, UnitPrice: 532},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "credit card",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 46, Quantity: 2, UnitPrice: 281},
				{VariantID: 144, Quantity: 28, UnitPrice: 397},
				{VariantID: 147, Quantity: 24, UnitPrice: 250},
			},
		},
		{
			CustomerID:    3,
			PaymentMethod: "qris",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 151, Quantity: 26, UnitPrice: 520},
				{VariantID: 174, Quantity: 5, UnitPrice: 654},
				{VariantID: 1, Quantity: 4, UnitPrice: 556},
			},
		},
		{
			CustomerID:    8,
			PaymentMethod: "credit card",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 188, Quantity: 1, UnitPrice: 156},
				{VariantID: 141, Quantity: 6, UnitPrice: 210},
				{VariantID: 135, Quantity: 21, UnitPrice: 341},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "qris",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 49, Quantity: 24, UnitPrice: 513},
			},
		},
		{
			CustomerID:    3,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 107, Quantity: 29, UnitPrice: 444},
				{VariantID: 44, Quantity: 29, UnitPrice: 396},
				{VariantID: 200, Quantity: 13, UnitPrice: 564},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "cod",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 47, Quantity: 10, UnitPrice: 127},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "debit",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 29, Quantity: 16, UnitPrice: 631},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "credit card",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 2, Quantity: 19, UnitPrice: 621},
				{VariantID: 111, Quantity: 6, UnitPrice: 525},
				{VariantID: 172, Quantity: 9, UnitPrice: 293},
				{VariantID: 112, Quantity: 30, UnitPrice: 551},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "cod",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 121, Quantity: 30, UnitPrice: 350},
				{VariantID: 2, Quantity: 18, UnitPrice: 320},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "bank transfer",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 188, Quantity: 12, UnitPrice: 396},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "bank transfer",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 91, Quantity: 12, UnitPrice: 604},
				{VariantID: 185, Quantity: 21, UnitPrice: 265},
				{VariantID: 115, Quantity: 22, UnitPrice: 633},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "credit card",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 90, Quantity: 17, UnitPrice: 309},
				{VariantID: 139, Quantity: 20, UnitPrice: 456},
				{VariantID: 128, Quantity: 26, UnitPrice: 580},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "bank transfer",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 103, Quantity: 25, UnitPrice: 687},
				{VariantID: 171, Quantity: 6, UnitPrice: 301},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "credit card",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 1, Quantity: 14, UnitPrice: 202},
				{VariantID: 19, Quantity: 22, UnitPrice: 683},
				{VariantID: 102, Quantity: 3, UnitPrice: 457},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "debit",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 20, Quantity: 21, UnitPrice: 747},
				{VariantID: 153, Quantity: 30, UnitPrice: 487},
				{VariantID: 181, Quantity: 6, UnitPrice: 372},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 89, Quantity: 20, UnitPrice: 565},
				{VariantID: 181, Quantity: 2, UnitPrice: 289},
				{VariantID: 65, Quantity: 9, UnitPrice: 593},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "qris",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 2, Quantity: 13, UnitPrice: 418},
				{VariantID: 197, Quantity: 15, UnitPrice: 364},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "bank transfer",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 82, Quantity: 11, UnitPrice: 116},
				{VariantID: 17, Quantity: 19, UnitPrice: 731},
				{VariantID: 138, Quantity: 20, UnitPrice: 195},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "qris",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 9, Quantity: 1, UnitPrice: 729},
				{VariantID: 181, Quantity: 22, UnitPrice: 464},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "credit card",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 2, Quantity: 22, UnitPrice: 457},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "cod",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 30, Quantity: 6, UnitPrice: 314},
				{VariantID: 75, Quantity: 16, UnitPrice: 586},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "cod",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 65, Quantity: 27, UnitPrice: 493},
				{VariantID: 122, Quantity: 18, UnitPrice: 377},
				{VariantID: 30, Quantity: 4, UnitPrice: 570},
			},
		},
		{
			CustomerID:    8,
			PaymentMethod: "debit",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 119, Quantity: 2, UnitPrice: 280},
				{VariantID: 123, Quantity: 1, UnitPrice: 222},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "cod",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 54, Quantity: 11, UnitPrice: 278},
				{VariantID: 69, Quantity: 15, UnitPrice: 332},
				{VariantID: 135, Quantity: 16, UnitPrice: 325},
				{VariantID: 181, Quantity: 19, UnitPrice: 654},
			},
		},
		{
			CustomerID:    8,
			PaymentMethod: "credit card",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 112, Quantity: 2, UnitPrice: 704},
				{VariantID: 79, Quantity: 24, UnitPrice: 737},
				{VariantID: 99, Quantity: 10, UnitPrice: 275},
			},
		},
		{
			CustomerID:    2,
			PaymentMethod: "cod",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 68, Quantity: 26, UnitPrice: 601},
				{VariantID: 103, Quantity: 20, UnitPrice: 297},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "bank transfer",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 32, Quantity: 9, UnitPrice: 571},
				{VariantID: 72, Quantity: 19, UnitPrice: 559},
				{VariantID: 37, Quantity: 1, UnitPrice: 650},
				{VariantID: 92, Quantity: 25, UnitPrice: 650},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "credit card",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 62, Quantity: 23, UnitPrice: 458},
				{VariantID: 92, Quantity: 1, UnitPrice: 383},
				{VariantID: 22, Quantity: 25, UnitPrice: 123},
				{VariantID: 200, Quantity: 7, UnitPrice: 746},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "bank transfer",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 176, Quantity: 1, UnitPrice: 272},
				{VariantID: 142, Quantity: 23, UnitPrice: 383},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "credit card",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 91, Quantity: 26, UnitPrice: 492},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "bank transfer",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 57, Quantity: 21, UnitPrice: 650},
				{VariantID: 123, Quantity: 28, UnitPrice: 324},
				{VariantID: 106, Quantity: 25, UnitPrice: 257},
				{VariantID: 27, Quantity: 12, UnitPrice: 709},
			},
		},
		{
			CustomerID:    5,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 3, Quantity: 26, UnitPrice: 600},
				{VariantID: 49, Quantity: 9, UnitPrice: 475},
				{VariantID: 51, Quantity: 22, UnitPrice: 130},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "bank transfer",
			Status:        "processed",
			Items: []OrderItem{
				{VariantID: 199, Quantity: 28, UnitPrice: 394},
				{VariantID: 179, Quantity: 17, UnitPrice: 155},
				{VariantID: 65, Quantity: 11, UnitPrice: 109},
				{VariantID: 108, Quantity: 12, UnitPrice: 383},
			},
		},
		{
			CustomerID:    3,
			PaymentMethod: "qris",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 74, Quantity: 20, UnitPrice: 350},
				{VariantID: 180, Quantity: 7, UnitPrice: 274},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "bank transfer",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 187, Quantity: 9, UnitPrice: 160},
			},
		},
		{
			CustomerID:    3,
			PaymentMethod: "credit card",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 45, Quantity: 23, UnitPrice: 289},
				{VariantID: 141, Quantity: 1, UnitPrice: 598},
				{VariantID: 118, Quantity: 28, UnitPrice: 501},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "debit",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 195, Quantity: 26, UnitPrice: 396},
				{VariantID: 64, Quantity: 30, UnitPrice: 471},
				{VariantID: 89, Quantity: 4, UnitPrice: 542},
				{VariantID: 31, Quantity: 3, UnitPrice: 215},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 102, Quantity: 28, UnitPrice: 288},
				{VariantID: 45, Quantity: 11, UnitPrice: 729},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "qris",
			Status:        "canceled",
			Items: []OrderItem{
				{VariantID: 82, Quantity: 8, UnitPrice: 133},
				{VariantID: 71, Quantity: 21, UnitPrice: 502},
				{VariantID: 200, Quantity: 24, UnitPrice: 216},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "qris",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 189, Quantity: 18, UnitPrice: 350},
				{VariantID: 11, Quantity: 20, UnitPrice: 474},
				{VariantID: 80, Quantity: 4, UnitPrice: 452},
			},
		},
		{
			CustomerID:    10,
			PaymentMethod: "debit",
			Status:        "created",
			Items: []OrderItem{
				{VariantID: 188, Quantity: 29, UnitPrice: 585},
				{VariantID: 189, Quantity: 3, UnitPrice: 438},
			},
		},
		{
			CustomerID:    9,
			PaymentMethod: "qris",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 17, Quantity: 20, UnitPrice: 353},
				{VariantID: 68, Quantity: 12, UnitPrice: 652},
			},
		},
		{
			CustomerID:    6,
			PaymentMethod: "bank transfer",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 191, Quantity: 11, UnitPrice: 474},
			},
		},
		{
			CustomerID:    7,
			PaymentMethod: "qris",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 30, Quantity: 19, UnitPrice: 228},
			},
		},
		{
			CustomerID:    4,
			PaymentMethod: "credit card",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 20, Quantity: 9, UnitPrice: 399},
			},
		},
		{
			CustomerID:    1,
			PaymentMethod: "debit",
			Status:        "completed",
			Items: []OrderItem{
				{VariantID: 138, Quantity: 25, UnitPrice: 285},
				{VariantID: 9, Quantity: 10, UnitPrice: 517},
			},
		},
	}
}
