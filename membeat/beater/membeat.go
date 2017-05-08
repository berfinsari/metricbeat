package beater

import (
	"fmt"
	"time"
	"bufio"
	"net"
	"strings"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/berfinsari/membeat/config"
)

type Membeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Membeat{
		done: make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Membeat) Run(b *beat.Beat) error {
	logp.Info("membeat is running! Hit CTRL-C to stop it.")
	bt.CollectStats("stats slabs",b.Name)
}


func (*bt Membeat) CollectStats(statsCommand string, beatname string ){
        command := statsCommand
        hostName := "localhost"
        portNum := "11211"
        var message, atama

        conn, err := net.Dial("tcp",hostName+":"+portNum)
        if err != nil{
                fmt.Println("Error %v",err)
                return
        } else{
                defer conn.Close()
                fmt.Fprintf(conn,command+"\n")
                message,_ = bufio.NewReader(conn).ReadString('E')
        }

        message = strings.SplitAfter(message, "\n")
        for i:=0;i<len(veri)-1;i++ {
                gecici := strings.Fields(message[i])
                atama[i] =gecici[2]
        }

        bt.client = b.Publisher.Connect()
        ticker := time.NewTicker(bt.config.Period)

        for {
                select {
                case <-bt.done:
                        return nil
                case <-ticker.C:
                }


	        activeslabs := atama[0]
		totalmalloced := atama[1]

	        event := common.MapStr{
		        "@timestamp":           common.Time(time.Now()),
		        "type":                 beatname,
			"active_slabs":         activeslabs,
			"total_malloced":       totalmalloced,
		}
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
        }

}

func (bt *Membeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
