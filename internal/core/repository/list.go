package repository

func List() []string {
	res := make([]string, 0)

	for _, repo := range repositories {
		res = append(res, repo.URL)
	}

	return res
}
