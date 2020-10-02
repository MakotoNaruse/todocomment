package a

func f() {
	// TODO: hoge // want "todo comment must contains issue's link"

	// TODO: hoge
	// nolint: todocomment

	// TODO: https://github.com/test/test/issues/
	// コメントだよ

	// コメント

	/* TODO これからやること githubのリンクはこれから作るよ */ // want "todo comment must contains issue's link"

	/*
		TODO これからやります issueには起こしていません
		nolint: todocomment
	 */
}

// TODO is a stuct name
// nolint: todocomment
type TODO struct {
}
