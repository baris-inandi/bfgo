package optimizer

func Optimize(code string) string {
	return canonicalise(
		removeUnusedLeading(
			code,
		),
	)
}
