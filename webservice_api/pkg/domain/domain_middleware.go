package domain

import (
	"github.com/gofiber/fiber/v2"
)

// Struct Middleware HTTP
type MiddlewareHTTP struct {
	Next                      func(*fiber.Ctx) bool
	XSSProtection             string
	ContentTypeNosniff        string
	XFrameOptions             string
	HSTSMaxAge                int
	HSTSExcludeSubdomains     bool
	ContentSecurityPolicy     string
	CSPReportOnly             bool
	HSTSPreloadEnabled        bool
	ReferrerPolicy            string
	PermissionPolicy          string
	CrossOriginEmbedderPolicy string
	CrossOriginOpenerPolicy   string
	CrossOriginResourcePolicy string
	OriginAgentCluster        string
	XDNSPrefetchControl       string
	XDownloadOptions          string
	XPermittedCrossDomain     string
}

// Struct Middleware CORS
type MiddlewareCORS struct {
	Next             func(c *fiber.Ctx) bool
	AllowOriginsFunc func(origin string) bool
	AllowOrigins     string
	AllowMethods     string
	AllowHeaders     string
	AllowCredentials bool
	ExposeHeaders    string
	MaxAge           int
}
