package main

import (
	"encoding/json"
	"image"
	"log"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

type Icon struct {
	Src   string `json:"src"`
	Sizes string `json:"sizes"`
	Type  string `json:"type"`
}

type Icons []Icon

type WebManifest struct {
	Name            string `json:"name"`
	ShortName       string `json:"short_name"`
	Icons           Icons  `json:"icons"`
	ThemeColor      string `json:"theme_color"`
	BackgroundColor string `json:"background_color"`
	Display         string `json:"display"`
}

func generateWebManifest(dstPath, siteName, themeColor, displayMode string) {
	icon192 := Icon{Src: "/android-chrome-192x192.png", Sizes: "192x192", Type: "image/png"}
	icon512 := Icon{Src: "/android-chrome-512x512.png", Sizes: "512x512", Type: "image/png"}

	var icons Icons
	icons = append(icons, icon192)
	icons = append(icons, icon512)

	webManifestJson := WebManifest{
		Name:            siteName,
		ShortName:       siteName,
		Icons:           icons,
		ThemeColor:      themeColor,
		BackgroundColor: themeColor,
		Display:         displayMode,
	}

	f, err := os.Create(dstPath)
	if err != nil {
		log.Printf("failed to create WebManifestJson: %v", err)
	}
	defer f.Close()

	data, err2 := json.Marshal(webManifestJson)
	if err2 != nil {
		log.Printf("failed to json marshal: %v", err)
		os.Exit(1)
	}

	_, err3 := f.Write(data)
	if err3 != nil {
		log.Printf("failed to write file: %v", err)
		os.Exit(1)
	}
}

func generateBrowserConfigXML(dstPath, tileColor string) {
	configXml := `<?xml version="1.0" encoding="utf-8"?>
<browserconfig>
    <msapplication>
        <tile>
            <square70x70logo src="/mstile-70x70.png"/>
            <square150x150logo src="/mstile-150x150.png"/>
            <wide310x150logo src="/mstile-310x150.png"/>
            <square310x310logo src="/mstile-310x310.png"/>
            <TileColor>` + tileColor + `</TileColor>
        </tile>
    </msapplication>
</browserconfig>`

	f, err := os.Create(dstPath)
	if err != nil {
		log.Printf("failed to create BrowserConfigXML: %v", err)
	}
	defer f.Close()

	data := []byte(configXml)
	_, err2 := f.Write(data)
	if err2 != nil {
		log.Printf("failed to write file: %v", err)
		os.Exit(1)
	}
}

func generateImage(img *image.Image, dstPath string, width, height int) {
	newImg := imaging.Clone(*img)
	resizeImg := imaging.Resize(newImg, width, height, imaging.Lanczos)

	err := imaging.Save(resizeImg, dstPath)
	if err != nil {
		log.Printf("failed to save image(%s): %v", filepath.Base(dstPath), err)
		os.Exit(1)
	}
}

func generateFaviconImages(targetImg, outputDir string) {
	src, err := imaging.Open(targetImg)
	if err != nil {
		log.Printf("failed to open image: %v", err)
		os.Exit(1)
	}

	generateImage(&src, filepath.Join(outputDir, "android-chrome-192x192.png"), 192, 192)
	generateImage(&src, filepath.Join(outputDir, "android-chrome-512x512.png"), 512, 512)
	generateImage(&src, filepath.Join(outputDir, "apple-touch-icon.png"), 180, 180)
	generateImage(&src, filepath.Join(outputDir, "favicon-16x16.png"), 16, 16)
	generateImage(&src, filepath.Join(outputDir, "favicon-32x32.png"), 32, 32)
	generateImage(&src, filepath.Join(outputDir, "favicon.png"), 48, 48)
	generateImage(&src, filepath.Join(outputDir, "mstile-70x70.png"), 70, 70)
	generateImage(&src, filepath.Join(outputDir, "mstile-150x150.png"), 150, 150)
	generateImage(&src, filepath.Join(outputDir, "mstile-310x150.png"), 310, 150)
	generateImage(&src, filepath.Join(outputDir, "mstile-310x310.png"), 310, 310)
}

func main() {
	// Specify target image.
	targetImg := "test.png"

	// Specify output directory.
	outputDir := "public"

	// Specify your site name.
	sitename := "test"

	// Specify tile color.
	tileColor := "#da532c"

	// Specify theme color.
	themeColor := "#ffffff"

	// Specify display mode.
	displayMode := "standalone"

	generateFaviconImages(targetImg, outputDir)
	generateBrowserConfigXML(filepath.Join(outputDir, "browserconfig.xml"), tileColor)
	generateWebManifest(filepath.Join(outputDir, "site.webmanifest"), sitename, themeColor, displayMode)

	log.Print("gosfg: : SUCCESS")
}
