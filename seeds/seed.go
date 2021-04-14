package seeds

import "go.uber.org/zap"

// Seed : struct
type Seed struct {
	log *zap.Logger
}

// Load : Load all seed
func Load(log *zap.Logger) {
	seed := Seed{log: log}

	seed.LoadCompanySeed()
	seed.LoadUserSeed()
	seed.LoadInformationSeed()
	seed.LoadCardSeed()
}
