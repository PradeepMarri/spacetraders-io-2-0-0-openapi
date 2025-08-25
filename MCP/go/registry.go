package main

import (
	"github.com/spacetraders-api/mcp-server/config"
	"github.com/spacetraders-api/mcp-server/models"
	tools_systems "github.com/spacetraders-api/mcp-server/tools/systems"
	tools_agents "github.com/spacetraders-api/mcp-server/tools/agents"
	tools_fleet "github.com/spacetraders-api/mcp-server/tools/fleet"
	tools_contracts "github.com/spacetraders-api/mcp-server/tools/contracts"
	tools_register "github.com/spacetraders-api/mcp-server/tools/register"
	tools_factions "github.com/spacetraders-api/mcp-server/tools/factions"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_systems.CreateGet_waypointTool(cfg),
		tools_agents.CreateGet_my_agentTool(cfg),
		tools_fleet.CreateShip_refineTool(cfg),
		tools_systems.CreateGet_system_waypointsTool(cfg),
		tools_fleet.CreateDock_shipTool(cfg),
		tools_contracts.CreateGet_contractsTool(cfg),
		tools_fleet.CreateCreate_ship_waypoint_scanTool(cfg),
		tools_fleet.CreateWarp_shipTool(cfg),
		tools_contracts.CreateAccept_contractTool(cfg),
		tools_fleet.CreatePurchase_cargoTool(cfg),
		tools_systems.CreateGet_shipyardTool(cfg),
		tools_fleet.CreateOrbit_shipTool(cfg),
		tools_fleet.CreateGet_ship_cooldownTool(cfg),
		tools_fleet.CreateGet_my_shipTool(cfg),
		tools_systems.CreateGet_systemsTool(cfg),
		tools_fleet.CreateCreate_chartTool(cfg),
		tools_contracts.CreateFulfill_contractTool(cfg),
		tools_fleet.CreateCreate_surveyTool(cfg),
		tools_register.CreateRegisterTool(cfg),
		tools_fleet.CreateJump_shipTool(cfg),
		tools_fleet.CreateSell_cargoTool(cfg),
		tools_contracts.CreateDeliver_contractTool(cfg),
		tools_systems.CreateGet_jump_gateTool(cfg),
		tools_contracts.CreateGet_contractTool(cfg),
		tools_fleet.CreateGet_ship_navTool(cfg),
		tools_fleet.CreatePatch_ship_navTool(cfg),
		tools_fleet.CreateTransfer_cargoTool(cfg),
		tools_fleet.CreateGet_my_ship_cargoTool(cfg),
		tools_fleet.CreateGet_my_shipsTool(cfg),
		tools_fleet.CreatePurchase_shipTool(cfg),
		tools_systems.CreateGet_marketTool(cfg),
		tools_factions.CreateGet_factionTool(cfg),
		tools_fleet.CreateCreate_ship_ship_scanTool(cfg),
		tools_fleet.CreateCreate_ship_system_scanTool(cfg),
		tools_fleet.CreateJettisonTool(cfg),
		tools_fleet.CreateExtract_resourcesTool(cfg),
		tools_factions.CreateGet_factionsTool(cfg),
		tools_fleet.CreateNavigate_shipTool(cfg),
		tools_fleet.CreateRefuel_shipTool(cfg),
		tools_systems.CreateGet_systemTool(cfg),
	}
}
