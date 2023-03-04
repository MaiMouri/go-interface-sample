package main

import (
	"fmt"
	"testing"
)

// 1. 差し替え用にMockの構造体を用意
type MockBathroomOutlet struct {
	BathroomOutlet        // BathroomOutletの持つ要素をカバーするために元の構造体を実装
	FakeHairDrying func() // HairDryingメソッドの結果を差し替えるための新メソッド追加
}

// 2. Fakeメソッドを呼ぶためのmock関数を定義
func (mb *MockBathroomOutlet) HairDrying() { // レシーバをMockに変更
	mb.FakeHairDrying() // Fakeバージョンのメソッドを呼ぶように中身を変更
}

func TestHairDrying_mockを使用(t *testing.T) {

	// 3. MockBathroomOutletを生成する
	m := &MockBathroomOutlet{
		FakeHairDrying: func() { // Mockの構造体にFakeHairDrying()の振る舞いを定義して生成
			fmt.Println("実は歯を磨いている")
		},
	}

	// 4. MockBathroomOutletにメソッドを呼ぶ
	m.HairDrying() // 2.で定義したようにFakeHairDryingが呼ばれる()
}
