import (
	"fmt"
	"time"
	"io/ioutil"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/berfinsari/metricbeat/lsbeat/config"
)

type Lsbeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
	lastIndexTime time.Time
}


func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Lsbeat{
		done: make(chan struct{}),
		config: config,
	}
	return bt, nil
}


func (bt *Lsbeat) Run(b *beat.Beat) error {
    logp.Info("lsbeat is running! Hit CTRL-C to stop it.")

    bt.client = b.Publisher.Connect()
    ticker := time.NewTicker(bt.config.Period)
    counter := 1
    for {

        select {
        case <-bt.done:
            return nil
        case <-ticker.C:
        }

        bt.listDir(bt.config.Path, b.Name, true) 
        bt.lastIndexTime = time.Now() 

        logp.Info("Event sent")
        counter++
    }

}

func (bt *Lsbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func (bt *Lsbeat) listDir(dirFile string, beatname string, init bool) {
    files, _ := ioutil.ReadDir(dirFile)
    for _, f := range files {
        t := f.ModTime()


        event := common.MapStr{
            "@timestamp": common.Time(time.Now()),
            "type":       beatname,
            "modtime":    common.Time(t),
            "filename":   f.Name(),
            "path":       dirFile + "/" + f.Name(),
            "directory":  f.IsDir(),
            "filesize":   f.Size(),
        }
        if init {
            bt.client.PublishEvent(event)
        } else {
            if t.After(bt.lastIndexTime) {
                bt.client.PublishEvent(event)
            }
        }

        if f.IsDir() {
            bt.listDir(dirFile+"/"+f.Name(), beatname, init)
        }
    }
}
