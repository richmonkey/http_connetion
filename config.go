
package main

import "strconv"
import "log"
import "github.com/richmonkey/cfg"

const DEFAULT_GROUP_DELIVER_COUNT = 4

type Config struct {
	port                int

	redis_address       string
	redis_password      string
	redis_db            int

	http_listen_address string

}

func get_int(app_cfg map[string]string, key string) int {
	concurrency, present := app_cfg[key]
	if !present {
		log.Fatalf("key:%s non exist", key)
	}
	n, err := strconv.ParseInt(concurrency, 10, 64)
	if err != nil {
		log.Fatalf("key:%s is't integer", key)
	}
	return int(n)
}

func get_opt_int(app_cfg map[string]string, key string) int64 {
	concurrency, present := app_cfg[key]
	if !present {
		return 0
	}
	n, err := strconv.ParseInt(concurrency, 10, 64)
	if err != nil {
		log.Fatalf("key:%s is't integer", key)
	}
	return n
}


func get_string(app_cfg map[string]string, key string) string {
	concurrency, present := app_cfg[key]
	if !present {
		log.Fatalf("key:%s non exist", key)
	}
	return concurrency
}

func get_opt_string(app_cfg map[string]string, key string) string {
	concurrency, present := app_cfg[key]
	if !present {
		return ""
	}
	return concurrency
}

func read_cfg(cfg_path string) *Config {
	config := new(Config)
	app_cfg := make(map[string]string)
	err := cfg.Load(cfg_path, app_cfg)
	if err != nil {
		log.Fatal(err)
	}

	config.port = get_int(app_cfg, "port")
	config.http_listen_address = get_string(app_cfg, "http_listen_address")
	config.redis_address = get_string(app_cfg, "redis_address")
	config.redis_password = get_opt_string(app_cfg, "redis_password")
	db := get_opt_int(app_cfg, "redis_db")
	config.redis_db = int(db)
	
	return config
}
