package conversion

import "strconv"

func StringsToFloats(strs []string) ([]float64, error) {
	var floats []float64
	for _, str := range strs {
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, f)
	}

	return floats, nil
}
