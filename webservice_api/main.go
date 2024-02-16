package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	domain "web/pkg/domain"
	http "web/pkg/http"
)

var setupMiddlewareHTTP = domain.MiddlewareHTTP{
	Next: nil,

	// ป้องกันการโจมตีแบบ Cross Site Scripting (CSP)
	XSSProtection:         "1; mode=block", // กำหนดการ กรอง Cross-site scripting (XSS)
	ContentSecurityPolicy: "",              // กำหนดการ สอบข้อมูลหน้าเว็บหรือ script ที่เป็นอันตรายก่อนที่จะส่งมาถึง Server
	CSPReportOnly:         false,           // กำหนดการ แจ้งเตือนการป้องกันการโจมตีแบบ script

	// สำหรับ HTTPS
	HSTSMaxAge:            0,     // กำหนดการ จำกัดเวลาที่ เข้าเว็บผ่านช่องทางเข้ารหัส HTTPS (สำหรับ HTTPS)
	HSTSExcludeSubdomains: false, // กำหนดการ ประกาศว่าควรเข้าถึงได้โดยใช้การเชื่อมต่อที่ปลอดภัย (HTTPS) เท่านั้น (สำหรับ HTTPS)
	HSTSPreloadEnabled:    false, // กำหนดการ ป้องกันเรื่อง connection ตอนใช้งาน HSTS (สำหรับ HTTPS)

	ReferrerPolicy:   "no-referrer", // กำหนด สั่งให้บราวเซอร์ที่ request เปิดเผยข้อมูลตามเส้นทางที่มา
	PermissionPolicy: "",            // กำหนด เส้นทางการเข้าถึง

	// อธิบาย Referrer-Policy
	// no-referrer                        = ส่วน Referrer จะถูกละเว้น คำขอที่ส่งจะไม่มีข้อมูลผู้อ้างอิงใดๆ
	// no-referrer-when-downgrade         = ส่งที่มา Origin path, และquerystring เมื่อระดับความปลอดภัยของโปรโตคอลคงเดิมหรือปรับปรุง (HTTP→HTTP, HTTP→HTTPS, HTTPS→HTTPS) ไม่ส่ง Referer สำหรับคำขอไปยังปลายทางที่มีความปลอดภัยน้อย (HTTPS→HTTP, HTTPS→file)
	// origin                             = ส่งที่มา Origin เท่านั้น เช่น https://example.com/page.html จะส่ง https://example.com/ หากคุณต้องการให้เว็บไซต์ของคุณทราบที่มาของคำขอโดยให้ใช้ origin ตัวอย่างเช่น google ,facebook
	// origin-when-cross-origin           = เมื่อดำเนิน การคำขอ  same-origin ในระดับโปรโตคอลเดียวกัน (HTTP→HTTP, HTTPS→HTTPS) ให้ส่งที่มา Origin path, และquerystring ส่งเฉพาะ Origin สำหรับคำขอ cross origin และคำขอไปยังปลายทางที่มีความปลอดภัยน้อย (HTTPS→HTTP)
	// same-origin                        = ส่งที่มา Origin และquery string ไม่ส่งส่วนหัวสำหรับคำขอข้าม Origin
	// strict-origin                      = ส่งที่มา Origin เมื่อระดับความปลอดภัยของโปรโตคอลเท่าเดิม (HTTPS→HTTPS) ไม่ส่ง Referrer ไปยังปลายทางที่มีความปลอดภัยน้อย เช่น (HTTPS→HTTP)
	// strict-origin-when-cross-origin    = ส่งที่มา Origin path, และquerystring เมื่อดำเนินการคำขอ สำหรับคำขอ cross-origin ให้ส่ง origin (เท่านั้น) เมื่อระดับความปลอดภัยของโปรโตคอลยังคงเหมือนเดิม (HTTPS→HTTPS) ไม่ส่ง Referer ไปยังปลายทางที่มีความปลอดภัยน้อย เช่น (HTTPS→HTTP)
	// unsafe-url                         = ส่งที่มา Origin path, และquerystring เมื่อดำเนินการคำขอใดๆ โดยไม่คำนึงถึงความปลอดภัย

	// default tag html  <a href="" rel=""</a> จะส่งแบบ referer origin เลยต้องกำหนด noreferrer เพื่อไม่ระบุ referer origin
	// <a href="https://meta.wikimedia.org/wiki/Wikimedia_News" rel="noreferrer">an external link</a> = ไม้ระบุ referer origin
	// <a href="https://meta.wikimedia.org/wiki/Wikimedia_News" rel="tag">an external link</a> = ระบุ referer origin
	// <meta name="https://google.com" content="origin" /> = ระบุ referer origin

	// สำหรับ Same-Origin Policy
	CrossOriginEmbedderPolicy: "unsafe-none", // กำหนดการเข้าถึงข้อมูลทรัพยากร
	CrossOriginOpenerPolicy:   "unsafe-none", // กำหนดการเข้าถึงข้อมูลทรัพยากรในรูปแบบหน้าต่างป็อปอัป
	CrossOriginResourcePolicy: "cross-site",  // กำหนดการเข้าถึงข้อมูลทรัพยากรเฉพาะ

	// ชื่อโฮสต์ และพอร์ตเดียวกันจะถือเป็น " Origin เดียวกัน" สิ่งอื่นๆ ถือเป็น "ข้าม Origin "
	// อธิบาย Cross-Origin-Embedder-Policy (COEP)
	// require-corp                 = อนุญาติให้ดึงเอกสารเฉพาะ Origin เดียวกันเท่านั้น *AllowCredentials: true,
	// unsafe-none                  = อนุญาตให้ดึงเอกสารข้อมูลข้าม Origin
	// credentialless               = อนุญาตให้ตอบกลับได้โดยไม่ต้องได้รับอนุญาต แต่คุกกี้จะถูกละเว้นจากคำขอหมายถึง ต้องมี cookie เข้ามาได้ *AllowCredentials: true,

	// อธิบาย Cross-Origin-Opener-Policy (COOP)
	// unsafe-none                  = อนุญาตให้ดึงเอกสารข้อมูลข้าม Origin
	// same-origin-allow-popups     = ส่งสัญญาณไปยังหน้าต่างที่เพิ่งเปิดใหม่ โดยที่หน้าต่างเหล่านั้นไม่ได้ตั้งค่า (COOP)
	// same-origin                  = อนุญาติให้ดึงเอกสารเฉพาะ Origin เดียวกันเท่านั้น ไม่อนุญาติดึงเอกสารข้าม Origin

	// อธิบาย Cross-Origin-Resource-Policy (CORP)
	// same-site                    = อนุญาติให้ดึงเอกสารเฉพาะไซต์เดียวกันเท่านั้น > https://web.dev/same-site-same-origin/
	// same-origin                  = อนุญาติให้ดึงเอกสารเฉพาะ Origin เดียวกันเท่านั้น ไม่อนุญาติดึงเอกสารข้าม Origin
	// cross-origin                 = อนุญาติให้ดึงเอกสารไซต์เดียวกัน และ Origin เดียวกันเท่านั้น ใช้กับ (COEP)
	// cross-site                   = อนุญาตให้ผู้ใช้ร้องขอจากไซต์ต่าง ๆ สามารถอ่านทรัพยากรได้

	// เมื่อดึงข้อมูลทรัพยากรรูปภาพ หากคุณไม่ระบุ crossOrigin รูปภาพจะถูกดึงโดยไม่มี CORS
	// crossorigin บนแท็ก <img> ระบุว่า CORS ได้รับการสนับสนุนเมื่อโหลดรูปภาพจากเซิร์ฟเวอร์หรือโดเมนของบุคคลที่สาม
	// <img crossorigin = "ไม่ระบุชื่อ | ใช้ข้อมูลรับรอง" >
	// <img crossorigin="anonymous | use-credentials">
	// <img src="http://192.168.100.231:8085/image/2.png" crossorigin="use-credentials">
	// anonymous       = คำขอข้าม Origin จะถูกส่งโดยไม่มีข้อมูลประจำตัวและดำเนินการตรวจสอบสิทธิ์ HTTP พื้นฐาน
	// use-credentials = คำขอข้าม Origin จะถูกส่งไปพร้อมข้อมูลรับรอง คุกกี้ และใบรับรอง
	// ถ้า crossOrigin เป็นสตริงว่าง ( "") anonymous โหมดจะถูกเลือก

	// สำหรับ ป้องกัน Scripting
	OriginAgentCluster: "?1", //  ป้องกันการเข้าถึงสคริปต์แบบซิงโครนัสระหว่างเพจข้าม Origin ของไซต์เดียวกัน
	//

	// สำหรับ Download
	XDownloadOptions: "noopen", // กำหนดการดาวน์โหลดโดยตรงในเบราว์เซอร์

	// สำหรับ ป้องกัน เว็บไซต์
	ContentTypeNosniff:    "nosniff",     // ป้องกัน การอัปโหลดไฟล์ ข้างในไฟล์อาจซ่อนไฟล์ที่เป็นอันตรายกับเซิฟเวอร์ได้
	XFrameOptions:         "DENY",        // ป้องกัน การโจมตีแบบแสดงผลหน้าเว็บของเราผ่านทาง tag HTML iFrame จากเว็บไซต์อื่น ๆ
	XPermittedCrossDomain: "master-only", // ป้องกัน ไม่ให้ผู้อื่นฝังเว็บไซต์
	XDNSPrefetchControl:   "off",         // ป้องกัน การดึงข้อมูล DNS ล่วงหน้า ( สำหรับ DNS)

}
var setupMiddlewareCORS = domain.MiddlewareCORS{
	Next:             nil,
	AllowOriginsFunc: nil,
	AllowOrigins:     "http://192.168.100.231:5500", // เป็น Header กำหนดว่าให้ใครเข้าถึง Resource ของ Server ได้บ้าง
	// *ระวังการใช้ 127.0.0.1 และ localhost เพราะมันไม่สามารบล็อคได้ ถ้าเครื่อง client ใช้ 127.0.0.1 และ localhost ในการ req api
	AllowHeaders:     "Origin, Content-Type, Accept",  // กำหนด Header ที่ใช้เรียกข้อมูลได้
	AllowMethods:     "GET",                           // กำหนด HTTP Method ที่สามารถ เรียน Server ได้
	AllowCredentials: true,                            // โดยปกติ Cookie จะไม่ถูกส่งใน CORS Request แต่หากำหนดค่านี้เป็น True ก็จะกลายเป็นว่าจะนำ Cookie เข้ามาไว้ใน CORS Requests ได้
	ExposeHeaders:    "Content-Encoding, X-User-Addr", // ระบุว่าส่วนหัว HTTP ใด ที่ไคลเอ็นต์จะเปิดเผย
	MaxAge:           0,                               // กำหนดเวลาที่ browser จะ cache response ของ request ที่เป็นการขออนุญาตต่างๆ
}

func main() {

	// Setup Fiber API
	app := fiber.New()
	app.Use(cors.New())
	app.Use(etag.New())
	mount := app.Group("/api/v1/dmk")

	// Config save bandwidth
	app.Use(etag.New(etag.Config{
		Weak: true,
		Next: nil,
	}))

	// Config Middleware HTTP
	app.Use(helmet.New(helmet.Config{
		setupMiddlewareHTTP.Next,
		setupMiddlewareHTTP.XSSProtection,
		setupMiddlewareHTTP.ContentTypeNosniff,
		setupMiddlewareHTTP.XFrameOptions,
		setupMiddlewareHTTP.HSTSMaxAge,
		setupMiddlewareHTTP.HSTSExcludeSubdomains,
		setupMiddlewareHTTP.ContentSecurityPolicy,
		setupMiddlewareHTTP.CSPReportOnly,
		setupMiddlewareHTTP.HSTSPreloadEnabled,
		setupMiddlewareHTTP.ReferrerPolicy,
		setupMiddlewareHTTP.PermissionPolicy,
		setupMiddlewareHTTP.CrossOriginEmbedderPolicy,
		setupMiddlewareHTTP.CrossOriginOpenerPolicy,
		setupMiddlewareHTTP.CrossOriginResourcePolicy,
		setupMiddlewareHTTP.OriginAgentCluster,
		setupMiddlewareHTTP.XDNSPrefetchControl,
		setupMiddlewareHTTP.XDownloadOptions,
		setupMiddlewareHTTP.XPermittedCrossDomain,
	}))

	// Config Middleware CORS
	// CORS คือ
	// GET HEAD POST จะใช้งานได้ทันทีเมื่อ allow origin
	// เมื่อ Client req HTTP Methods อะไรมา มันจะลองส่ง HTTP Methods OPTIONS (Preflight request) มาเป็นอันดับแรกเพื่อดูว่า server allow HTTP Methods อะไรบ้าง
	// และ server ก็จะส่งมาว่า allow อะไรบ้าง
	// GET HEAD POST = ไม่ต้องเพิ่มใน Config HTTPS Methods เพราะเป็น CORS Safelisted methods *แต่ถ้าใส่ ต้องเพิ่ม OPTIONS ไปด้วยเพื่อให้มี Preflight request
	// เมื่อมี PUT PATCH DELETE ต้องใส่ OPTIONS ไปด้วยเพื่อให้มี Preflight request
	app.Use(cors.New(cors.Config{
		setupMiddlewareCORS.Next,
		setupMiddlewareCORS.AllowOriginsFunc,
		setupMiddlewareCORS.AllowOrigins,
		setupMiddlewareCORS.AllowMethods,
		setupMiddlewareCORS.AllowHeaders,
		setupMiddlewareCORS.AllowCredentials,
		setupMiddlewareCORS.ExposeHeaders,
		setupMiddlewareCORS.MaxAge,
	}))

	// Config Logger
	app.Use(logger.New(logger.Config{
		Format: "${red}[${ip}]:${port} ${status} ${time} - ${method} ${yellow}${path} \n",
	}))

	// API Dashboard
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics RPTS"}))

	// Link
	// http://127.0.0.1:8085/api/v1/dmk/place_T1_DepartureSecurity
	// http://127.0.0.1:8085/api/v1/dmk/place_T2_DepartureSecurity

	// API Place_T1_DepartureSecurity
	http.Place_T1_DepartureSecurity(mount)

	// API Place_T2_DepartureSecurity
	http.Place_T2_DepartureSecurity(mount)

	// app.Use("/image", filesystem.New(filesystem.Config{
	// 	Root: pkger.Dir("/image"),
	// }))

	// Start server
	log.Fatal(app.Listen(":8085"))
}
