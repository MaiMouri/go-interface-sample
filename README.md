# go-interface-sample
## 定義



### producerの定義

```go
type Dryer struct{}
func (d *Dryer) HasTwoPlugs() bool {
	return true
}
func (d *Dryer) LengthOfPlag() bool {
	return true
}
func (d *Dryer) DistanceOfPlugs() bool {
	return true
}
func (d *Dryer) SwitchOn() bool {
	fmt.Println("ドライヤースタート💇‍♀️")
	d.drying()
	return true
}
func (d *Dryer) drying() bool {
	time.Sleep(1 * time.Second)
	fmt.Println("........乾かし中")
	time.Sleep(1 * time.Second)
	fmt.Println("........乾いた！")
	return true
}
```


## consumerの定義

producerの機能を使う側。 構造体にinterfaceを実装し、どんなものかを知らなくてもinterfaceに適合するものを受け取れるようにする。この例では、同じinterfaceを二つ使用する実装にしたのでフィールド名をfirstOutletとsecondOutletのように区別している。

`NewBathroomOutlet()`で構造体を生成する際、二つのinterfaceを引数とする。それぞれで使用する機能のproducer構造体を生成して渡すことでこの関数が実行可能。

```go

type IBathroomOutlet interface {
	HasTwoPlugs() bool
	LengthOfPlag() bool
	DistanceOfPlugs() bool
	SwitchOn() bool
}

type BathroomOutlet struct {
	firstOutlet  IBathroomOutlet // interfaceを実装する
	secondOutlet IBathroomOutlet // interfaceを実装する
}

func NewBathroomOutlet(one IBathroomOutlet, two IBathroomOutlet) *BathroomOutlet {
	return &BathroomOutlet{
		firstOutlet:  one,
		secondOutlet: two,
	}
}

func (b *BathroomOutlet) HairDrying() {
	fmt.Println("First outlet gets connected.")
	b.firstOutlet.SwitchOn()
}

func (b *BathroomOutlet) Laundry() {
	fmt.Println("Second outlet gets connected.")
	b.secondOutlet.SwitchOn()
}
```

### interfaceの定義

consumer側で使用するにはproducerの機能をラップしている実装である必要がある。
※ この時パブリック関数から呼ばれるプライベート関数については実装不要。

```go
type IBathroomOutlet interface {
	HasTwoPlugs() bool
	LengthOfPlag() bool
	DistanceOfPlugs() bool
	SwitchOn() bool
}
```

## テスト

### Mockの構造体
元の構造体を実装したMockを定義する。mockを作りたいメソッドは`Fake`等わかりやすくした名称で実装する。

```go
type MockBathroomOutlet struct {
	BathroomOutlet        // BathroomOutletの持つ要素をカバーするために元の構造体を実装
	FakeHairDrying func() // HairDryingメソッドの結果を差し替えるための新メソッド追加
}

```

### Fakeメソッドを呼び出すように元の関数を変更する


Mock構造体に対して呼ぶメソッドの中身を変更する。Mock構造体を通して呼ばれたHairDrying()は、FakeHairDrying()を呼ぶ。
```go
func (mb *MockBathroomOutlet) HairDrying() { // レシーバをMockに変更
	mb.FakeHairDrying() // Fakeバージョンのメソッドを呼ぶように中身を変更
}

```



### Testコードの書き方

```go
func TestHairDrying_mockを使用(t *testing.T) {

	// MockBathroomOutletを生成する
	m := &MockBathroomOutlet{
		FakeHairDrying: func() { // Mockの構造体にFakeHairDrying()の振る舞いを定義して生成
			fmt.Println("実は歯を磨いている")
		},
	}

	// MockBathroomOutletにメソッドを呼ぶ
	m.HairDrying() // FakeHairDryingが呼ばれる()
}

```