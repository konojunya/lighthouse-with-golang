package main

import (
	"encoding/json"
	"os/exec"

	frame "github.com/konojunya/go-frame"
)

type Lighthouse struct {
	IsOnHTTPS                     float32 `json:"is-on-https"`
	RedirectsHTTP                 float32 `json:"redirects-http"`
	ServiceWorker                 float32 `json:"service-worker"`
	WorksOffline                  float32 `json:"works-offline"`
	Viewport                      float32 `json:"viewport"`
	WithoutJavascript             float32 `json:"without-javascript"`
	FirstContentfulPaint          float32 `json:"first-contentful-paint"`
	FirstMeaningfulPaint          float32 `json:"first-meaningful-paint"`
	LoadFastEnoughForPwa          float32 `json:"load-fast-enough-for-pwa"`
	SpeedIndex                    float32 `json:"speed-index"`
	EstimatedInputLatency         float32 `json:"estimated-input-latency"`
	ErrorsInConsole               float32 `json:"errors-in-console"`
	TimeToFirstByte               float32 `json:"time-to-first-byte"`
	FirstCPUIdle                  float32 `json:"first-cpu-idle"`
	Interactive                   float32 `json:"interactive"`
	Redirects                     float32 `json:"redirects"`
	WebappInstallBanner           float32 `json:"webapp-install-banner"`
	SplashScreen                  float32 `json:"splash-screen"`
	ThemedOmnibox                 float32 `json:"themed-omnibox"`
	ManifestShortNameLength       float32 `json:"manifest-short-name-length"`
	ContentWidth                  float32 `json:"content-width"`
	ImageAspectRatio              float32 `json:"image-aspect-ratio"`
	Deprecations                  float32 `json:"deprecations"`
	MainthreadWorkBreakdown       float32 `json:"mainthread-work-breakdown"`
	BootupTime                    float32 `json:"bootup-time"`
	UsesRelPreload                float32 `json:"uses-rel-preload"`
	UsesRelPreconnect             float32 `json:"uses-rel-preconnect"`
	FontDisplay                   float32 `json:"font-display"`
	Bypass                        float32 `json:"bypass"`
	ColorContrast                 float32 `json:"color-contrast"`
	DocumentTitle                 float32 `json:"document-title"`
	DuplicateID                   float32 `json:"duplicate-id"`
	HTMLHasLang                   float32 `json:"html-has-lang"`
	HTMLLangValid                 float32 `json:"html-lang-valid"`
	ImageAlt                      float32 `json:"image-alt"`
	LinkName                      float32 `json:"link-name"`
	List                          float32 `json:"list"`
	Listitem                      float32 `json:"listitem"`
	MetaViewport                  float32 `json:"meta-viewport"`
	UsesLongCacheTTL              float32 `json:"uses-long-cache-ttl"`
	TotalByteWeight               float32 `json:"total-byte-weight"`
	OffscreenImages               float32 `json:"offscreen-images"`
	RenderBlockingResources       float32 `json:"render-blocking-resources"`
	UnminifiedCSS                 float32 `json:"unminified-css"`
	UnminifiedJavascript          float32 `json:"unminified-javascript"`
	UnusedCSSRules                float32 `json:"unused-css-rules"`
	UsesWebpImages                float32 `json:"uses-webp-images"`
	UsesOptimizedImages           float32 `json:"uses-optimized-images"`
	UsesTextCompression           float32 `json:"uses-text-compression"`
	UsesResponsiveImages          float32 `json:"uses-responsive-images"`
	EfficientAnimatedContent      float32 `json:"efficient-animated-content"`
	AppcacheManifest              float32 `json:"appcache-manifest"`
	Doctype                       float32 `json:"doctype"`
	DomSize                       float32 `json:"dom-size"`
	ExternalAnchorsUseRelNoopener float32 `json:"external-anchors-use-rel-noopener"`
	GeolocationOnStart            float32 `json:"geolocation-on-start"`
	NoDocumentWrite               float32 `json:"no-document-write"`
	NoVulnerableLibraries         float32 `json:"no-vulnerable-libraries"`
	NoWebsql                      float32 `json:"no-websql"`
	NotificationOnStart           float32 `json:"notification-on-start"`
	PasswordInputsCanBePastedInto float32 `json:"password-inputs-can-be-pasted-into"`
	UsesHTTP2                     float32 `json:"uses-http2"`
	UsesPassiveEventListeners     float32 `json:"uses-passive-event-listeners"`
	MetaDescription               float32 `json:"meta-description"`
	HTTPStatusCode                float32 `json:"http-status-code"`
	FontSize                      float32 `json:"font-size"`
	LinkText                      float32 `json:"link-text"`
	IsCrawlable                   float32 `json:"is-crawlable"`
	RobotsTxt                     float32 `json:"robots-txt"`
	Hreflang                      float32 `json:"hreflang"`
	Plugins                       float32 `json:"plugins"`
}

func main() {
	// ここのURLは任意で変えれるようにする
	// light-house-format.js は、結果をjsonの形に綺麗にするためのものなので必要
	commandStr := "lighthouse https://konojunya.com --output json --quiet | node lighthouse-format.js"
	out, err := exec.Command("sh", "-c", commandStr).Output()
	if err != nil {
		panic(err)
	}

	var lighthouseScore = Lighthouse{}
	err = json.Unmarshal(out, &lighthouseScore)
	if err != nil {
		panic(err)
	}
	frame.Print(lighthouseScore)
}
