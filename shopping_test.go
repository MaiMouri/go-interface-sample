package main

import (
	"fmt"
	"testing"

	"github.com/MaiMouri/go-interface-sample/domain/service"
)

// 1. 差し替え用にMockの構造体を用意
type MockCard struct {
	service.PointCard            // Shoppingの持つ要素をカバーするために元の構造体を実装
	FakeReadCard      func() int // ReadCardメソッドの結果を差し替えるための新メソッド追加
}

// 2. Fakeメソッドを呼ぶためのmock関数を定義
func (mb *MockCard) ReadCard() int { // レシーバをMockに変更
	return mb.FakeReadCard() // Fakeバージョンのメソッドを呼ぶように中身を変更
}

// MockCard構造体を用意し、ReadCardを呼ぶとFakeReadCard()が実行されるようにする
func TestReadCard_mockを使用(t *testing.T) {

	// 3. MockCardを生成する
	c := &MockCard{
		FakeReadCard: func() int { // Mockの構造体にFakeReadCard()の振る舞いを定義して生成
			fmt.Println("MockCard が呼ばれた　ポイントが少なく加算される")
			return 1
		},
	}
	s := &Shopping{
		card: c,
	}
	// 4. Shoppingにメソッドを呼ぶ
	want := s.ShowCard() // 2.で定義したようにFakeReadCardが呼ばれる()

	// 5. 本来のメソッドを呼ぶ
	cc := &service.PointCard{}
	ss := &Shopping{
		card: cc,
	}
	got := ss.ShowCard()

	// 6. 結果の照合
	if got == want {
		t.Errorf("ShowCard() = %v should not be eqaul to %v", got, want)
	}

}
