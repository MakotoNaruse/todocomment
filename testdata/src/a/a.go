package a

func f() {
	// TODO: hoge // want "todo comment must contains issue's link"

	// TODO: https://github.com/test/test/issues/
	// コメントだよ

	// コメント

	/* TODO: これからやること githubのリンクはこれから作るよ */ // want "todo comment must contains issue's link"
}

