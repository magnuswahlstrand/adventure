package death

import (
	"github.com/kyeett/single-player-game/internal/comp"
)

func (s *Death) Update(_ float64) {

	// Find entities that have 0 or negative hit points
	var ids []comp.ID
	for id, s := range s.entities {
		if s.Hitpoints.Amount <= 0 {

			// Mark for removing
			ids = append(ids, id)
		}
	}

	// Remove
	for _, id := range ids {
		s.logger.Info("remove entity " + string(id))
		s.lifeCycler.Remove(id)
	}
}
