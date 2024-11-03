package class_file

import (
	"fmt"
	"io"

	"github.com/takattty/make-jvm/support"
)

type (
	ClassFile struct{}
)

// JVM仕様で定義されているclassファイルの固定マジックナンバー
const magicNumber = 0xCAFEBABE

func ReadClassFile(cfReader io.Reader) (*ClassFile, error) {
	r, err := support.NewByteSeq(cfReader)
	if err != nil {
		return nil, err
	}

	// マジックナンバー（u4=非負4バイト整数）を検証
	if r.ReadUnit32() != magicNumber {
		return nil, fmt.Errorf("not java class file")
	}

	// マイナー・メジャーバージョン（u2=非負2バイト整数）は単にスキップ（合計4バイト）
	r.Skip(4)

	// TODO: major_versionより後の、constant_pool_count以降の内容のデコード
	return nil, nil
}
