package server

import (
	"SOMAS2023/internal/clients/team2"
	"SOMAS2023/internal/common/objects"

	baseserver "github.com/MattSScott/basePlatformSOMAS/BaseServer"
	"github.com/google/uuid"
)

const BikerAgentCount = 6

func GetAgentGenerators() []baseserver.AgentGeneratorCountPair[objects.IBaseBiker] {
	return []baseserver.AgentGeneratorCountPair[objects.IBaseBiker]{
		baseserver.MakeAgentGeneratorCountPair[objects.IBaseBiker](BikerAgentGenerator, BikerAgentCount),
	}
}

func BikerAgentGenerator() objects.IBaseBiker {
	// return objects.GetIBaseBiker(utils.GenerateRandomColour(), uuid.New())
	return team2.NewBaseTeam2Biker(uuid.New())
}

func (s *Server) spawnLootBox() {
	lootBox := objects.GetLootBox()
	s.lootBoxes[lootBox.GetID()] = lootBox
}

func (s *Server) replenishLootBoxes() {
	count := LootBoxCount - len(s.lootBoxes)
	for i := 0; i < count; i++ {
		s.spawnLootBox()
	}
}

func (s *Server) spawnMegaBike() {
	megaBike := objects.GetMegaBike()
	s.megaBikes[megaBike.GetID()] = megaBike
}

func (s *Server) replenishMegaBikes() {
	for i := 0; i < MegaBikeCount-len(s.megaBikes); i++ {
		s.spawnMegaBike()
	}
}
