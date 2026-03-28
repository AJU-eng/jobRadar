package main

import "log"

func main() {

	DBConfig := dbConfig{
		addr:           "postgresql://neondb_owner:npg_ZBXq8RCh2jVW@ep-mute-field-a4kr7fz6-pooler.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require",
		maxOpenConnecs: 30,
		maxIdleConnecs: 30,
		maxIdleTime:    "15m",
	}
	cnfg := Config{
		addr:     ":8080",
		dbConfig: DBConfig,
	}

	app := application{
		config: cnfg,
	}

	log.Println("server started running on :8080")

	mux := app.Mount()

	err := app.Run(mux)

	if err != nil {
		log.Println(err)
	}

}
