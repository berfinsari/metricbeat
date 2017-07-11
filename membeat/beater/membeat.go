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

	"github.com/berfinsari/metricbeat/membeat/config"
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
        bt.client = b.Publisher.Connect()
        ticker := time.NewTicker(bt.config.Period)
	counter := 1

        for {
                select {
                case <-bt.done:
                        return nil
                case <-ticker.C:
                }
		bt.collectStats(b.Name)
		logp.Info("Event sent")
		counter ++
	}

}
func (bt *Membeat) collectStats(beatname string) {

	command := "stats slabs"
        hostName := "localhost"
        portNum := "11211"
	var value [55]string
	var message string

	conn, err := net.Dial("tcp",hostName+":"+portNum)
	if err != nil{
                logp.Err("Error reading stats: %v",err)
                return
        } else{
                defer conn.Close()
                fmt.Fprintf(conn,command+"\n")
                message,_ = bufio.NewReader(conn).ReadString('E')
        }

        veri := strings.SplitAfter(message, "\n")
        for i:=0;i<len(veri)-1;i++ {
                gecici := strings.Fields(veri[i])
                value[i] = gecici[2]
	 }
	now := time.Now()

	pid :=	value[0]
	uptime := value[1]
	time :=	value[2]
	version := value[3]
        libevent := value[4]
        pointer_size := value[5]
        rusage_user :=	value[6]
        rusage_system :=value[7]
        curr_connections := value[8]
        total_connections := value[9]
        reserved_fds := value[10]
        cmd_get := value[11]
        cmd_set := value[12]
        cmd_flush := value[13]
        cmd_touch := value[14]
        get_hits := value[15]
        get_misses := value[16]
        delete_misses := value[17]
        delete_hits := value[18]
        incr_misses := value[19]
        incr_hits := value[20]
        decr_misses := value[21]
        decr_hits := value[22]
        cas_misses := value[23]
        cas_hits := value[24]
        cas_badval := value[25]
        touch_hits := value[26]
        touch_misses := value[27]
        auth_cmds := value[28]
        auth_errors := value[29]
        bytes_read := value[30]
        bytes_written := value[31]
        limit_maxbytes := value[32]
        accepting_conns := value[33]
        listen_disable_num := value[34]
	time_in_listen_disable_us := value[35]
        threads := value[36]
        conn_yields := value[37]
        hash_power_level := value[38]
        hash_bytes := value[39]
        hash_is_expanding := value[40]
        malloc_fails := value[41]
        bytes := value[42]
        curr_items := value[43]
        total_items := value[44]
        expired_unfetched := value[45]
        evicted_unfetched := value[46]
        evictions := value[47]
        reclaimed := value[48]
        crawler_reclaimed := value[49]
        crawler_items_checked := value[50]
        lrutail_reflocked := value[51]

        event := common.MapStr{
		"@timestamp":			common.Time(now),
	        "type":				beatname,
		"pid":				pid,
		"uptime":			uptime,
		"time":				time,
		"version":			version,
		"libevent":			libevent,
		"pointer_size":			pointer_size,
		"rusage_user":			rusage_user,
		"rusage_system":		rusage_system,
		"curr_connections":		curr_connections,
		"total_connections":		total_connections,
		"reserved_fds":			reserved_fds,
		"cmd_get":			cmd_get,
		"cmd_set":			cmd_set,
		"cmd_flush":			cmd_flush,
		"cmd_touch":			cmd_touch,
		"get_hits":			get_hits,
		"get_misses":			get_misses,
		"delete_misses":		delete_misses,
		"delete_hits":			delete_hits,
		"incr_misses":			incr_misses,
		"incr_hits":			incr_hits,
		"decr_misses":			decr_misses,
		"decr_hits":			decr_hits,
		"cas_misses":			cas_misses,
		"cas_hits":			cas_hits,
		"cas_badval":			cas_badval,
		"touch_hits":			touch_hits,
		"touch_misses":			touch_misses,
		"auth_cmds":			auth_cmds,
		"auth_errors":			auth_errors,
		"bytes_read":			bytes_read,
		"bytes_written":		bytes_written,
		"limit_maxbytes":		limit_maxbytes,
		"accepting_conns":		accepting_conns,
		"listen_disable_num":		listen_disable_num,
		"time_in_listen_disable_us":	time_in_listen_disable_us,
		"threads":			threads,
		"conn_yields":			conn_yields,
		"hash_power_level":		hash_power_level,
		"hash_bytes":			hash_bytes,
		"hash_is_expanding":		hash_is_expanding,
		"malloc_fails":			malloc_fails,
		"bytes":			bytes,
		"curr_items":			curr_items,
		"total_items":			total_items,
		"expired_unfetched":		expired_unfetched,
		"evicted_unfetched":		evicted_unfetched,
		"evictions":			evictions,
		"reclaimed":			reclaimed,
		"crawler_reclaimed":		crawler_reclaimed,
		"crawler_items_checked":	crawler_items_checked,
		"lrutail_reflocked":		lrutail_reflocked,
	}
	bt.client.PublishEvent(event)
}

func (bt *Membeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
