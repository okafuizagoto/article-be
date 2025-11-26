package boot

import (
	"article-be/internal/config"
	"article-be/internal/consumer"
	"article-be/internal/registry"
	"log"
)

func StartKafkaConsumers(cfg config.KafkaConfig, reg *registry.Registry) {
	if cfg.Topics.LocalToProd != "" {
		// Local → Prod
		go func() {
			consumer.ConsumeLoop(
				cfg.Brokers,
				cfg.Topics.LocalToProd,
				cfg.GroupID+"-local-to-prod",
				reg,
			)
		}()
	}
	// // Prod → Local
	// go func() {
	// 	consumer.ConsumeLoop(
	// 		cfg.Brokers,
	// 		cfg.Topics.ProdToLocal,
	// 		cfg.GroupID+"-prod-to-local",
	// 		reg,
	// 	)
	// }()

	log.Println("[BOOT] Kafka consumers started")
}
