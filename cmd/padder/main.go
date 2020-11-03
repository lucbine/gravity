package main

// main is the bootstrap.
//func main() {
//	cfg := configs.NewConfig()
//	err := cfg.ParseCmd(os.Args[1:])
//	switch errors.Cause(err) {
//	case nil:
//	case flag.ErrHelp:
//		os.Exit(0)
//	default:
//		log.Fatalf("parse cmd flags errors: %s\n", err)
//	}
//
//	if cfg.ConfigFile == "" {
//		log.Fatalf("configs file is required")
//	}
//
//	if err := cfg.CreateConfigFromFile(cfg.ConfigFile); err != nil {
//		log.Fatalf("failed to load configs from file. %v", err)
//	}
//	if err = configs.Validate(cfg.PadderConfig); err != nil {
//		log.Fatalf("configs validation failed: %v", err)
//	}
//	logutil.MustInitLogger(&cfg.Log)
//	utils.LogRawInfo("padder")
//	if cfg.PreviewMode {
//		stats, err := padder.Preview(cfg.PadderConfig)
//		if err != nil {
//			log.Fatalf("pad preview bin log failed: %v", err)
//		}
//		statsJson, err := json.MarshalIndent(stats, "", "    ")
//		if err != nil {
//			log.Fatalf("parse json failed: %v", err)
//		}
//		err = ioutil.WriteFile("stats.json", statsJson, 0644)
//		if err != nil {
//			log.Fatalf("export preview statistic failed: %v", err)
//		}
//	} else {
//		if err := padder.Pad(cfg.PadderConfig); err != nil {
//			log.Fatalf("pad bin log failed: %v", err)
//		}
//	}
//}
