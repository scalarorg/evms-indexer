package btc

import (
	"encoding/json"
	"time"
)

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
	VaultTag             string        `json:"vault_tag"`
	VaultVersion         uint8         `json:"vault_version"`
	DatabaseURL          string        `json:"database_url"` // Separate DB for electrum
	DialTimeout          time.Duration `json:"dial_timeout"`
	MethodTimeout        time.Duration `json:"method_timeout"`
	PingInterval         time.Duration `json:"ping_interval"`
	MaxReconnectAttempts int           `json:"max_reconnect_attempts"`
	ReconnectDelay       time.Duration `json:"reconnect_delay"`
	EnableAutoReconnect  bool          `json:"enable_auto_reconnect"`
	StartHeight          int64         `json:"start_height"`
	EndHeight            int64         `json:"end_height"`
	BatchSize            int           `json:"batch_size"`
	Confirmations        int           `json:"confirmations"`
	SourceChain          string        `json:"source_chain"`
	FetchThread          int           `json:"fetch_thread"` // Number of worker threads for fetching blocks
}

// UnmarshalJSON implements custom JSON unmarshaling to handle string-to-time.Duration conversion
func (c *BtcConfig) UnmarshalJSON(data []byte) error {
	// Create a temporary struct with string fields for time.Duration
	type Alias BtcConfig
	aux := &struct {
		DialTimeout    string `json:"dial_timeout"`
		MethodTimeout  string `json:"method_timeout"`
		PingInterval   string `json:"ping_interval"`
		ReconnectDelay string `json:"reconnect_delay"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	// Unmarshal into the temporary struct
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Convert string durations to time.Duration
	if aux.DialTimeout != "" {
		duration, err := time.ParseDuration(aux.DialTimeout)
		if err != nil {
			return err
		}
		c.DialTimeout = duration
	}

	if aux.MethodTimeout != "" {
		duration, err := time.ParseDuration(aux.MethodTimeout)
		if err != nil {
			return err
		}
		c.MethodTimeout = duration
	}

	if aux.PingInterval != "" {
		duration, err := time.ParseDuration(aux.PingInterval)
		if err != nil {
			return err
		}
		c.PingInterval = duration
	}

	if aux.ReconnectDelay != "" {
		duration, err := time.ParseDuration(aux.ReconnectDelay)
		if err != nil {
			return err
		}
		c.ReconnectDelay = duration
	}

	return nil
}
