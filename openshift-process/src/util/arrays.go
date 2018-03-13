package util

type MapFunc func(string) string

func MapString(list []string, function MapFunc) []string {
	if len(list) < 1 {
		return []string{}
	}

	output := make([]string, len(list))
	for i,element := range list {
		output[i] = function(element)
	}

	return output
}