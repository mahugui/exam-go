package utils

func ArrBytesToInterface(slice1 [][]byte) (slice2 []interface{}) {
	slice2 = make([]interface{}, len(slice1))
	for i, v := range slice1 {
		slice2[i] = v
	}
	return
}

func ArrInterfaceByteToArrString(slice1 []interface{}) (slice2 []string) {
	slice2 = make([]string, len(slice1))
	for i, v := range slice1 {
		slice2[i] = string(v.([]byte))
	}
	return
}