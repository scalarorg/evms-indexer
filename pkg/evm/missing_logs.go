package evm

import (
	"sync"
	"sync/atomic"

	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

type MissingLogs struct {
	chainId   string
	mutex     sync.Mutex
	logs      []ethTypes.Log
	Recovered atomic.Bool //True if logs are recovered
}

func (m *MissingLogs) IsRecovered() bool {
	return m.Recovered.Load()
}
func (m *MissingLogs) SetRecovered(recovered bool) {
	m.Recovered.Store(recovered)
}

func (m *MissingLogs) AppendLogs(logs []ethTypes.Log) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	log.Info().Str("Chain", m.chainId).Int("Number of event logs", len(logs)).Msgf("[EvmClient] [AppendLogs] appending logs")
	m.logs = append(m.logs, logs...)
}

func (m *MissingLogs) GetLogs(count int) []ethTypes.Log {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var logs []ethTypes.Log
	if len(m.logs) <= count {
		logs = m.logs
		m.logs = []ethTypes.Log{}
	} else {
		logs = m.logs[:count]
		m.logs = m.logs[count:]
	}
	log.Info().Str("Chain", m.chainId).Int("Number of logs", len(logs)).
		Int("Remaining logs", len(m.logs)).
		Msgf("[MissingLogs] [GetLogs] returned logs")
	return logs
}
