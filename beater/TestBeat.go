package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/jackcloudman/testbeat/config"
)

// testbeat configuration.
type testbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of testbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &testbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts testbeat.
func (bt *testbeat) Run(b *beat.Beat) error {
	logp.Info("testbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	currentID := loadState()
	logp.Info( fmt.Sprintf("CurrentID: %d",currentID))
	startConection()
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
		// Read employees
		persons, err := ReadEmployees(currentID)
		if err != nil {
			logp.Error(err)
		}
		if persons != nil {
			currentID = persons[len(persons)-1].id
			for _,p:= range persons{
				event := beat.Event{
					Timestamp: time.Now(),
					Fields: common.MapStr{
						"type":    b.Info.Name,
						"name": p.name,
						"id": p.id,
						"location": p.location,
					},
				}
				bt.client.Publish(event)
				logp.Info(fmt.Sprintf("Registro (%s) enviado",p.name))
			}
		} else {
			logp.Info("Estas al dia!")
		}
		storeState(currentID)
	}
}

// Stop stops testbeat.
func (bt *testbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
