package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_term "github.com/input-api/mcp-server/tools/term"
	tools_api "github.com/input-api/mcp-server/tools/api"
	tools_cache "github.com/input-api/mcp-server/tools/cache"
	tools_p3 "github.com/input-api/mcp-server/tools/p3"
	tools_servicevalidate "github.com/input-api/mcp-server/tools/servicevalidate"
	tools_expire "github.com/input-api/mcp-server/tools/expire"
	tools_general "github.com/input-api/mcp-server/tools/general"
	tools_clean "github.com/input-api/mcp-server/tools/clean"
	tools_key "github.com/input-api/mcp-server/tools/key"
	tools_authzen "github.com/input-api/mcp-server/tools/authzen"
	tools_encrypt "github.com/input-api/mcp-server/tools/encrypt"
	tools_upn "github.com/input-api/mcp-server/tools/upn"
	tools_validate "github.com/input-api/mcp-server/tools/validate"
	tools_v1 "github.com/input-api/mcp-server/tools/v1"
	tools_username "github.com/input-api/mcp-server/tools/username"
	tools_authorize "github.com/input-api/mcp-server/tools/authorize"
	tools_well_known "github.com/input-api/mcp-server/tools/well_known"
	tools_logout "github.com/input-api/mcp-server/tools/logout"
	tools_repositories "github.com/input-api/mcp-server/tools/repositories"
	tools_decrypt "github.com/input-api/mcp-server/tools/decrypt"
	tools_ticketid "github.com/input-api/mcp-server/tools/ticketid"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_term.CreateGet_termTool(cfg),
		tools_api.CreateGetTool(cfg),
		tools_cache.CreateDelete_cache_usernameTool(cfg),
		tools_cache.CreateGet_cache_usernameTool(cfg),
		tools_p3.CreatePost_p3_servicevalidateTool(cfg),
		tools_servicevalidate.CreatePost_servicevalidateTool(cfg),
		tools_expire.CreateDelete_expireTool(cfg),
		tools_general.CreateGet_Tool(cfg),
		tools_general.CreatePost_Tool(cfg),
		tools_clean.CreateDelete_cleanTool(cfg),
		tools_key.CreateDelete_keyTool(cfg),
		tools_authzen.CreatePost_authzenTool(cfg),
		tools_encrypt.CreatePost_encryptTool(cfg),
		tools_upn.CreateGet_upnTool(cfg),
		tools_validate.CreatePost_validateTool(cfg),
		tools_v1.CreatePost_v1_servicesTool(cfg),
		tools_username.CreateGet_usernameTool(cfg),
		tools_authorize.CreatePost_authorizeTool(cfg),
		tools_well_known.CreateGet_well_known_acme_challenge_tokenTool(cfg),
		tools_logout.CreatePost_logout_postTool(cfg),
		tools_repositories.CreateGet_repositoriesTool(cfg),
		tools_decrypt.CreatePost_decryptTool(cfg),
		tools_username.CreateDelete_username_providerid_keyTool(cfg),
		tools_ticketid.CreateDelete_ticketidTool(cfg),
		tools_ticketid.CreateGet_ticketidTool(cfg),
	}
}
