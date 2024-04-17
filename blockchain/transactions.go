package blockchain

import (
	"time"

	"github.com/fullgukbap/coin/utils"
)

const (
	// minerReward 변수는 광부에게 주는 포상금 입니다.
	minerReward int = 50
)

// Tx 구조체는 거래를 하기 위해 사용되는 정보로 구성됩니다.
type Tx struct {
	// Id 변수는 각각의 거래를 식별하기 위해 사용됩니다.
	Id string `json:"id"`

	// Timestamp는 거래가 발생한 시간을 기록하기 위해 사용됩니다.
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

// getId 함수는 Id 값을 자동적으로 생성하여 대입해줍니다.
func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

// TxIn 구조체는 거래의 Input의 구성요소가 존재합니다.
type TxIn struct {
	// Owner는 소유주를 나타냅니다.
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

// TxOuts 구조체는 거래의 Output의 구성요소가 존재합니다.
type TxOut struct {
	// Owner는 소유주를 나타냅니다.
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

// makeCoinbaseTx 함수는 채굴자를 주소로 삼는 코인베이스 거래내역을 생성해 *Tx 포인터를 반환합니다.
func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"COINBASE", minerReward},
	}

	txOuts := []*TxOut{
		{address, minerReward},
	}

	tx := Tx{
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return &tx
}