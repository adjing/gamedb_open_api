package gamedb_open_api

func GetAdsByteDeleteByIndex(ad []byte, indexnum int) []byte {

	var b []byte
	var index = 0
	for i := 0; i < len(ad); i++ {
		index = index + 1
		if index > indexnum {
			a := ad[i]
			b = append(b, a)

		}

	}
	return b

}
