package btc

import "time"

// ElectrumIndexerConfig holds configuration for the electrum indexer
type BtcConfig struct {
	Enable               bool          `json:"enable"`
	BtcHost              string        `json:"btc_host"`
	BtcPort              int           `json:"btc_port"`
	BtcUser              string        `json:"btc_user"`
	BtcPassword          string        `json:"btc_password"`
	BtcSSL               *bool         `json:"btc_ssl"`
	BtcNetwork           string        `json:"btc_network"`
	ElectrumHost         string        `json:"electrum_host"`
	ElectrumPort         int           `json:"electrum_port"`
	DatabaseURL          string        `json:"database_url"` // Separate DB for electrum
	DialTimeout          time.Duration `json:"dial_timeout"`
	MethodTimeout        time.Duration `json:"method_timeout"`
	PingInterval         time.Duration `json:"ping_interval"`
	MaxReconnectAttempts int           `json:"max_reconnect_attempts"`
	ReconnectDelay       time.Duration `json:"reconnect_delay"`
	EnableAutoReconnect  bool          `json:"enable_auto_reconnect"`
	BatchSize            int           `json:"batch_size"`
	Confirmations        int           `json:"confirmations"`
	SourceChain          string        `json:"source_chain"`
}
