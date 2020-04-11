package domain

// MergeOverrides creates the union of two sets by Key with the override values taking precedence
func MergeOverrides(base, overrides []Limit) []Limit {
	merged := map[string]struct{}{}
	res := make([]Limit, 0)

	for _, e := range base {
		merged[e.Key] = struct{}{}
		if b := findByKey(overrides, e.Key); b != nil {
			res = append(res, *b)
			continue
		}

		res = append(res, e)
	}

	for _, e := range overrides {
		if _, ok := merged[e.Key]; ok {
			continue
		}

		res = append(res, e)
	}

	return res
}

func findByKey(haystack []Limit, needle string) *Limit {
	for _, e := range haystack {
		if e.Key == needle {
			return &e
		}
	}

	return nil
}
