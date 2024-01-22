package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-pdf/fpdf"
)

func main() {
	// Définir les arguments de la ligne de commande
	title := flag.String("title", "", "Le titre du PDF")
	text := flag.String("text", "", "Le texte du PDF")
	flag.Parse()

	// Créer un nouveau PDF
	pdf := fpdf.New("P", "mm", "A4", "")

	// Ajouter le numéro de page en bas à droite
	pdf.AliasNbPages("")
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d/{nb}", pdf.PageNo()), "", 0, "R", false, 0, "")
	})

	// Ajouter une page
	pdf.AddPage()

	// Ajouter le logo en haut à gauche et plus grand
	pdf.Image("logo.png", 10, 10, 40, 40, false, "", 0, "")

	// Définir la police pour le titre plus grande
	pdf.SetFont("Arial", "B", 30)

	// Définir la couleur pour le titre
	pdf.SetTextColor(0, 0, 0)

	// Ajouter le titre centré une ligne en dessous du logo
	width, _ := pdf.GetPageSize()
	pdf.SetY(60) // Déplacer le Y à une position en dessous du logo
	pdf.CellFormat(width-20, 10, *title, "0", 0, "C", false, 0, "")

	// Souligner le titre
	pdf.SetDrawColor(0, 0, 0) // Noir
	pdf.Line(10, 70, width-10, 70)

	// Définir la marge gauche
	pdf.SetLeftMargin(10)

	// Définir la police pour le texte avec une taille plus petite
	pdf.SetFont("Arial", "", 12)

	// Déplacer le Y cinq lignes en dessous du titre
	pdf.SetY(100)

	// Ajouter des espaces au début du texte pour l'indentation
	*text = "    " + *text

	// Remplacer les retours à la ligne après 'Mon' par des espaces
	*text = strings.ReplaceAll(*text, "Mon\n", "Mon ")

	// Ajouter le texte// 10mm de marge à gauche pour l'indentation
	*text = "    " + *text // Ajouter quatre espaces au début pour l'indentation
	pdf.MultiCell(0, 5, *text, "", "", false)

	// Ajouter la date
	_, pageHeight := pdf.GetPageSize()
	dateY := pageHeight - 60 // 60 est la hauteur totale de la signature et de la date
	pdf.SetXY(10, dateY)
	pdf.SetFont("Arial", "", 12) // Même taille de police que le texte
	pdf.Cell(50, 10, "Date: "+time.Now().Format("02-01-2006"))

	// Ajouter la signature électronique
	signatureY := dateY + 10                              // La signature est placée juste en dessous de la date
	pdf.SetFillColor(200, 200, 200)                       // Gris clair
	pdf.RoundedRect(10, signatureY, 50, 50, 5, "D", "DF") // Rectangle avec bords arrondis
	pdf.SetXY(10, signatureY)
	pdf.Cell(50, 10, "Signature")
	pdf.Image("signature.png", 10, signatureY+10, 30, 30, false, "", 0, "")

	// Générer le PDF
	err := pdf.OutputFileAndClose("output.pdf")
	if err != nil {
		fmt.Println("Erreur lors de la génération du PDF :", err)
		os.Exit(1)
	}
}
